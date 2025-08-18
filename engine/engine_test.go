package engine

import (
	"log"
	"math/big"
	"os"
	"testing"

	ecommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/vultisig/recipes/sdk/evm/codegen/erc20"
	"github.com/vultisig/vultisig-go/common"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/vultisig/recipes/types"
)

func buildUnsignedTx(to ecommon.Address, data []byte, value *big.Int) []byte {
	unsigned := struct {
		ChainID    *big.Int
		Nonce      uint64
		GasTipCap  *big.Int
		GasFeeCap  *big.Int
		Gas        uint64
		To         *ecommon.Address `rlp:"nil"`
		Value      *big.Int
		Data       []byte
		AccessList etypes.AccessList
	}{
		ChainID:    big.NewInt(1),
		Nonce:      0,
		GasTipCap:  big.NewInt(2_000_000_000),  // 2 gwei
		GasFeeCap:  big.NewInt(20_000_000_000), // 20 gwei
		Gas:        300_000,
		To:         &to,
		Value:      value,
		Data:       data,
		AccessList: nil,
	}
	payload, err := rlp.EncodeToBytes(unsigned)
	if err != nil {
		panic(err)
	}
	return append([]byte{etypes.DynamicFeeTxType}, payload...)
}

var testVectors = []struct {
	policyPath string
	chain      common.Chain
	schemaPath string
	txHex      string
	txHexFunc  func() string
	shouldPass bool
}{
	{
		policyPath: "../testdata/payroll.json",
		chain:      common.Ethereum,
		txHexFunc: func() string {
			return hexutil.Encode(buildUnsignedTx(
				ecommon.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
				erc20.NewErc20().PackTransfer(
					ecommon.HexToAddress("0xcf0475d9B0a29975bc5132A3066010eC898d8CaB"),
					big.NewInt(1000000),
				),
				big.NewInt(0),
			))
		},
		shouldPass: true,
	},
	// Schema validation tests
	{
		policyPath: "../testdata/payroll.json",
		schemaPath: "../testdata/payroll_schema.json",
		shouldPass: true,
	},
}

func TestEngine(t *testing.T) {
	engine := NewEngine()
	engine.SetLogger(log.Default())

	for _, testVector := range testVectors {
		tv := testVector
		t.Run(tv.policyPath, func(t *testing.T) {
			policyFileBytes, err := os.ReadFile(tv.policyPath)
			if err != nil {
				t.Fatalf("Failed to read policy file: %v", err)
			}

			var policy types.Policy
			if err := protojson.Unmarshal(policyFileBytes, &policy); err != nil {
				t.Fatalf("Failed to unmarshal policy: %v", err)
			}

			var schema types.RecipeSchema
			if tv.schemaPath != "" {
				schemaFileBytes, err := os.ReadFile(tv.schemaPath)
				if err != nil {
					t.Fatalf("Failed to read schema file %s: %v", tv.schemaPath, err)
				}

				if err := protojson.Unmarshal(schemaFileBytes, &schema); err != nil {
					t.Fatalf("Failed to unmarshal schema JSON: %v", err)
				}
				t.Logf("Successfully loaded schema for plugin: %s (Version: %d)",
					schema.GetPluginName(), schema.GetPluginVersion())

				err = engine.ValidatePolicyWithSchema(&policy, &schema)
				if err != nil && tv.shouldPass {
					t.Fatalf("Failed to validate policy: %s vs. %s: %v", tv.policyPath, tv.schemaPath, err)
				}
				if err == nil && !tv.shouldPass {
					t.Fatalf("Expected validation to fail for policy: %s vs. %s, but it passed", tv.policyPath, tv.schemaPath)
				}
				return
			}

			txHex := tv.txHex
			if tv.txHexFunc != nil {
				txHex = tv.txHexFunc()
			}

			txBytes, err := hexutil.Decode(txHex)
			if err != nil && tv.shouldPass {
				t.Fatalf("Failed to decode transaction: %v", err)
			}

			matchingRule, err := engine.Evaluate(&policy, tv.chain, txBytes)
			if err != nil && tv.shouldPass {
				t.Fatalf("Failed to evaluate transaction: %v", err)
			}

			if tv.shouldPass && matchingRule == nil {
				t.Fatalf("No matching rule found")
			}
		})
	}
}
