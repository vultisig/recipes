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
	"github.com/stretchr/testify/require"
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

// TestDenyWins asserts that EFFECT_DENY rules are evaluated before ALLOW rules and that a
// matching DENY rule causes the evaluation to return an error even when an ALLOW rule
// would have matched the same transaction.
func TestDenyWins(t *testing.T) {
	eng, err := NewEngine()
	require.NoError(t, err)

	usdcAddr := ecommon.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	uint256Max, _ := new(big.Int).SetString(
		"115792089237316195423570985008687907853269984665640564039457584007913129639935", 10,
	)
	spender := ecommon.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")

	// Build an erc20.approve(spender, uint256Max) transaction — this should be denied.
	approveData := erc20.NewErc20().PackApprove(spender, uint256Max)
	txBytes := buildUnsignedTx(usdcAddr, approveData, big.NewInt(0))

	policy := &types.Policy{
		Rules: []*types.Rule{
			{
				Id:       "deny-unbounded-approve",
				Resource: "ethereum.erc20.approve",
				Effect:   types.Effect_EFFECT_DENY,
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: usdcAddr.Hex()},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "spender",
						Constraint:    &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_ANY},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type:  types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{FixedValue: uint256Max.String()},
						},
					},
				},
			},
			{
				Id:       "allow-usdc-approve",
				Resource: "ethereum.erc20.approve",
				Effect:   types.Effect_EFFECT_ALLOW,
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target:     &types.Target_Address{Address: usdcAddr.Hex()},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "spender",
						Constraint:    &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_ANY},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
							Value: &types.Constraint_MaxValue{
								MaxValue: uint256Max.String(), // ALLOW MAX = uint256Max
							},
						},
					},
				},
			},
		},
	}

	// uint256Max matches the DENY rule — must be rejected despite the ALLOW rule matching too.
	rule, err := eng.Evaluate(policy, common.Ethereum, txBytes)
	require.Error(t, err, "uint256.max approve must be denied")
	require.Nil(t, rule)
	require.Contains(t, err.Error(), "denied")
}

func TestEngine(t *testing.T) {
	engine, err := NewEngine()
	require.NoError(t, err)

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
