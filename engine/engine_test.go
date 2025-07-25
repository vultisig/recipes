package engine

import (
	"log"
	"os"
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/vultisig/recipes/chain"
	"github.com/vultisig/recipes/types"
)

func TestEngine_ValidatePolicyWithSchema(t *testing.T) {
	engine := NewEngine()
	engine.SetLogger(log.Default())

	testcases := []struct {
		name          string
		policyPath    string
		schemaPath    string
		expectedError bool
		errorContains string
	}{
		{
			name:          "Valid policy",
			policyPath:    "../testdata/payroll_erc20.json",
			schemaPath:    "../testdata/payroll_schema.json",
			expectedError: false,
		},
		{
			name:          "Policy with no rules",
			policyPath:    "../testdata/payroll_erc20.json", // Will be modified to have no rules
			schemaPath:    "../testdata/payroll_schema.json",
			expectedError: true,
			errorContains: "policy has no rules",
		},
		{
			name:          "Policy ID doesn't match schema plugin ID",
			policyPath:    "../testdata/payroll_erc20.json", // Will be modified to have different ID
			schemaPath:    "../testdata/payroll_schema.json",
			expectedError: true,
			errorContains: "does not match schema plugin ID",
		},
		{
			name:          "Invalid configuration",
			policyPath:    "../testdata/payroll_erc20.json", // Will be modified to have invalid configuration
			schemaPath:    "../testdata/payroll_schema.json",
			expectedError: true,
			errorContains: "configuration validation failed",
		},
		{
			name:          "Unsupported resource",
			policyPath:    "../testdata/payroll_erc20.json", // Will be modified to have unsupported resource
			schemaPath:    "../testdata/payroll_schema.json",
			expectedError: true,
			errorContains: "uses unsupported resource",
		},
		{
			name:          "Invalid parameter constraint",
			policyPath:    "../testdata/payroll_erc20.json", // Will be modified to have invalid parameter constraint
			schemaPath:    "../testdata/payroll_schema.json",
			expectedError: true,
			errorContains: "constraint validation failed",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Read and parse policy file
			policyFileBytes, err := os.ReadFile(tc.policyPath)
			if err != nil {
				t.Fatalf("Failed to read policy file: %v", err)
			}

			var policy types.Policy
			if err := protojson.Unmarshal(policyFileBytes, &policy); err != nil {
				t.Fatalf("Failed to unmarshal policy: %v", err)
			}

			// Read and parse schema file
			schemaFileBytes, err := os.ReadFile(tc.schemaPath)
			if err != nil {
				t.Fatalf("Failed to read schema file: %v", err)
			}

			var schema types.RecipeSchema
			if err := protojson.Unmarshal(schemaFileBytes, &schema); err != nil {
				t.Fatalf("Failed to unmarshal schema: %v", err)
			}

			// Modify policy based on test case
			switch tc.name {
			case "Policy with no rules":
				policy.Rules = nil
			case "Policy ID doesn't match schema plugin ID":
				policy.Id = "different-id"
			case "Invalid configuration":
				invalidConfig, err := structpb.NewStruct(map[string]interface{}{
					"unknown_field": "value",
				})
				if err != nil {
					t.Fatalf("Failed to create invalid configuration: %v", err)
				}
				policy.Configuration = invalidConfig
			case "Unsupported resource":
				policy.Rules[0].Resource = "unsupported.resource.path"
			case "Invalid parameter constraint":
				// Add an unsupported parameter
				policy.Rules[0].ParameterConstraints = append(policy.Rules[0].ParameterConstraints,
					&types.ParameterConstraint{
						ParameterName: "unsupported_param",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value:    &types.Constraint_FixedValue{FixedValue: "value"},
							Required: true,
						},
					})
			}

			// Validate policy against schema
			err = engine.ValidatePolicyWithSchema(&policy, &schema)

			// Check if error matches expectation
			if tc.expectedError {
				if err == nil {
					t.Errorf("Expected error but got nil")
				} else if !strings.Contains(err.Error(), tc.errorContains) {
					t.Errorf("Expected error containing '%s', got: %v", tc.errorContains, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

func TestEngine_Evaluate(t *testing.T) {
	engine := NewEngine()
	engine.SetLogger(log.Default())

	var testcases = []struct {
		policyPath string
		chainStr   string
		txHex      string
		shouldPass bool
	}{
		{
			policyPath: "../testdata/payroll_erc20.json",
			chainStr:   "ethereum",
			txHex:      "0x02f9016f018207188289fd84177aab1b8301772e94dac17f958d2ee523a2206206994597c13d831ec780b844a9059cbb0000000000000000000000009756752dc9f0a947366b6c91bb0487d6c6bf4d170000000000000000000000000000000000000000000000000000000000950858f90100f8fe94dac17f958d2ee523a2206206994597c13d831ec7f8e7a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000003a00000000000000000000000000000000000000000000000000000000000000004a0000000000000000000000000000000000000000000000000000000000000000aa043a6b5ea290b0f552490febd2e8b34a549d6929a7cc62128cb0f52d5dc6fe52fa04963d32529428ccb0324ab153dcc06505006f80d9d888e440dcaaae6e0de81c9a0cc147bba85f155948dd5950fb9848d3ba69891fbe3f53dccc4376f15220ff080",
			shouldPass: true,
		},
		{
			policyPath: "../testdata/payroll_erc20.json",
			chainStr:   "ethereum",
			// wrong amount in calldata
			txHex:      "0x02f9016f018207188289fd841b17ef788301772e94dac17f958d2ee523a2206206994597c13d831ec780b844a9059cbb0000000000000000000000009756752dc9f0a947366b6c91bb0487d6c6bf4d170000000000000000000000000000000000000000000000000000000000950859f90100f8fe94dac17f958d2ee523a2206206994597c13d831ec7f8e7a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000003a00000000000000000000000000000000000000000000000000000000000000004a0000000000000000000000000000000000000000000000000000000000000000aa043a6b5ea290b0f552490febd2e8b34a549d6929a7cc62128cb0f52d5dc6fe52fa04963d32529428ccb0324ab153dcc06505006f80d9d888e440dcaaae6e0de81c9a0cc147bba85f155948dd5950fb9848d3ba69891fbe3f53dccc4376f15220ff080",
			shouldPass: false,
		},
	}

	for _, _tc := range testcases {
		tc := _tc
		t.Run(tc.policyPath, func(t *testing.T) {
			policyFileBytes, err := os.ReadFile(tc.policyPath)
			if err != nil {
				t.Fatalf("Failed to read policy file: %v", err)
			}

			var policy types.Policy
			if err := protojson.Unmarshal(policyFileBytes, &policy); err != nil {
				t.Fatalf("Failed to unmarshal policy: %v", err)
			}

			c, err := chain.GetChain(tc.chainStr)
			if err != nil {
				t.Fatalf("Failed to get chain: %v", err)
			}

			tx, err := c.ParseTransaction(tc.txHex)
			if err != nil {
				t.Fatalf("Failed to parse transaction: %v", err)
			}

			transactionAllowedByPolicy, matchingRule, err := engine.Evaluate(&policy, c, tx)
			if err != nil {
				t.Fatalf("Failed to evaluate transaction: %v", err)
			}

			if transactionAllowedByPolicy != tc.shouldPass {
				t.Fatalf(
					"Transaction allowed by policy: %t, expected: %t",
					transactionAllowedByPolicy,
					tc.shouldPass,
				)
			}

			if tc.shouldPass && matchingRule == nil {
				t.Fatalf("No matching rule found")
			}
		})
	}
}
