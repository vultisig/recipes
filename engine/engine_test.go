package engine

import (
	"log"
	"os"
	"strings"
	"sync"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/vultisig/recipes/chain"
	"github.com/vultisig/recipes/ethereum"
	"github.com/vultisig/recipes/testdata"
	"github.com/vultisig/recipes/types"
)

var testVectors = []struct {
	policyPath string
	chainStr   string
	schemaPath string
	txHex      string
	txHexFunc  func() string
	shouldPass bool
}{
	{
		policyPath: "../testdata/payroll.json",
		chainStr:   "ethereum",
		txHex:      "0x00ec80872386f26fc10000830f424094b0b00000000000000000000000000000000000018806f05b59d3b2000080",
		shouldPass: true,
	},
	{
		policyPath: "../testdata/payroll.json",
		chainStr:   "bitcoin",
		txHex:      "010000000100000000000000000000000000000000000000000000000000000000000000000000000000ffffffff01404b4c00000000001976a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac00000000",
		shouldPass: true,
	},
	{
		policyPath: "../testdata/payroll.json",
		chainStr:   "ethereum",
		txHex:      "0x00ec80872386f26fc10000830f424094b1b00000000000000000000000000000000000018806f05b59d3b2000080",
		shouldPass: false,
	},
	{
		policyPath: "../testdata/payroll.json",
		chainStr:   "bitcoin",
		txHex:      "010000000100000000000000000000000000000000000000000000000000000000000000000000000000ffffffff01404b4c00000000001976a91462e917b15cbf27d5425399ebf6f0fb50ebb88f1888ac00000000",
		shouldPass: false,
	},
	// Uniswap test cases
	{
		policyPath: "../testdata/uniswap_policy.json",
		chainStr:   "ethereum",
		txHexFunc:  testdata.ValidSwapExactETHForTokensTxHex,
		shouldPass: true,
	},
	{
		policyPath: "../testdata/uniswap_policy.json",
		chainStr:   "ethereum",
		txHexFunc:  testdata.InvalidRecipientSwapExactETHForTokensTxHex,
		shouldPass: false,
	},
	// additional Uniswap tests
	{
		policyPath: "../testdata/uniswap_policy.json",
		chainStr:   "ethereum",
		txHexFunc:  testdata.ExceedAmountSwapExactTokensForETHTxHex,
		shouldPass: false,
	},
	{
		policyPath: "../testdata/uniswap_policy.json",
		chainStr:   "ethereum",
		txHexFunc:  testdata.ValidSwapExactTokensForETHTxHex,
		shouldPass: true,
	},
	{
		policyPath: "../testdata/uniswap_policy.json",
		chainStr:   "ethereum",
		txHexFunc:  testdata.ValidAddLiquidityTxHex,
		shouldPass: true,
	},
	{
		policyPath: "../testdata/uniswap_policy.json",
		chainStr:   "ethereum",
		txHexFunc:  testdata.InvalidTokenAddLiquidityTxHex,
		shouldPass: false,
	},
	{
		policyPath: "../testdata/uniswap_policy.json",
		chainStr:   "ethereum",
		txHexFunc:  testdata.ValidRemoveLiquidityTxHex,
		shouldPass: true,
	},
	{
		policyPath: "../testdata/uniswap_policy.json",
		chainStr:   "ethereum",
		txHexFunc:  testdata.InvalidRecipientRemoveLiquidityTxHex,
		shouldPass: false,
	},
	// Schema validation tests
	{
		policyPath: "../testdata/payroll.json",
		schemaPath: "../testdata/payroll_schema.json",
		shouldPass: true,
	},
	{
		policyPath: "../testdata/xrp_payroll.json",
		schemaPath: "../testdata/payroll_schema.json",
		shouldPass: false,
	},
	{
		policyPath: "../testdata/invalid_configuration_payroll.json",
		schemaPath: "../testdata/payroll_schema.json",
		shouldPass: false,
	},
}

var registerOnce sync.Once

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
				return
			}

			c, err := chain.GetChain(tv.chainStr)
			if err != nil {
				t.Fatalf("Failed to get chain: %v", err)
			}

			// For Uniswap tests, ensure the Uniswap ABI is loaded and protocols registered
			if strings.Contains(tv.policyPath, "uniswap_policy") && tv.chainStr == "ethereum" {
				ethChain := c.(*ethereum.Ethereum)
				abiBytes, err := os.ReadFile("../abi/uniswapV2_router.json")
				if err != nil {
					t.Fatalf("Failed to read Uniswap ABI: %v", err)
				}
				if err := ethChain.LoadABI("uniswapv2_router", abiBytes); err != nil {
					t.Fatalf("Failed to load Uniswap ABI: %v", err)
				}
				// Register protocols (no token list, no ERC20 ABI needed for this test)
				registerOnce.Do(func() {
					if err := ethereum.RegisterEthereumProtocols(ethChain, nil); err != nil {
						t.Fatalf("Failed to register Ethereum protocols: %v", err)
					}
				})
			}

			txHex := tv.txHex
			if tv.txHexFunc != nil {
				txHex = tv.txHexFunc()
			}

			tx, err := c.ParseTransaction(txHex)
			if err != nil {
				t.Fatalf("Failed to parse transaction: %v", err)
			}

			transactionAllowedByPolicy, matchingRule, err := engine.Evaluate(&policy, c, tx)
			if err != nil {
				t.Fatalf("Failed to evaluate transaction: %v", err)
			}

			if transactionAllowedByPolicy != tv.shouldPass {
				t.Fatalf("Transaction allowed by policy: %t, expected: %t", transactionAllowedByPolicy, tv.shouldPass)
			}

			if tv.shouldPass && matchingRule == nil {
				t.Fatalf("No matching rule found")
			}
		})
	}
}
