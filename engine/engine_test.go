package engine

import (
	"log"
	"os"
	"testing"

	"github.com/vultisig/recipes/chain"
	"github.com/vultisig/recipes/types"
	"google.golang.org/protobuf/encoding/protojson"
)


var testVectors = []struct {
	policyPath string
	chainStr string
	txHex string
	shouldPass bool
}{
	{
		policyPath: "../testdata/payroll.json",
		chainStr: "ethereum",
		txHex: "0x00ec80872386f26fc10000830f424094b0b00000000000000000000000000000000000018806f05b59d3b2000080",
		shouldPass: true,
	},
	{
		policyPath: "../testdata/payroll.json",
		chainStr: "bitcoin",
		txHex: "010000000100000000000000000000000000000000000000000000000000000000000000000000000000ffffffff01404b4c00000000001976a91462e907b15cbf27d5425399ebf6f0fb50ebb88f1888ac00000000",
		shouldPass: true,
	},
	{
		policyPath: "../testdata/payroll.json",
		chainStr: "ethereum",
		txHex: "0x00ec80872386f26fc10000830f424094b1b00000000000000000000000000000000000018806f05b59d3b2000080",
		shouldPass: false,
	},
	{
		policyPath: "../testdata/payroll.json",
		chainStr: "bitcoin",
		txHex: "010000000100000000000000000000000000000000000000000000000000000000000000000000000000ffffffff01404b4c00000000001976a91462e917b15cbf27d5425399ebf6f0fb50ebb88f1888ac00000000",
		shouldPass: false,
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

			chain, err := chain.GetChain(tv.chainStr)
			if err != nil {
				t.Fatalf("Failed to get chain: %v", err)
			}

			tx, err := chain.ParseTransaction(tv.txHex)
			if err != nil {
				t.Fatalf("Failed to parse transaction: %v", err)
			}

			transactionAllowedByPolicy, matchingRule, err := engine.Evaluate(policy, chain, tx)
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