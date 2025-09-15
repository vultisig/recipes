package metarule

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/recipes/solana"
	"github.com/vultisig/recipes/types"
)

const testAddress = "4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM"

func TestTryFormat_NonMetaRule(t *testing.T) {
	metaRule := NewMetaRule()

	// Test with a complete rule that has function ID (not a meta-rule)
	rule := &types.Rule{
		Resource: "solana.sol.transfer", // Already has function ID
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testAddress,
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	assert.Equal(t, rule, result) // Should return unchanged
}

func TestTryFormat_UnsupportedChain(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "bitcoin.send", // Unsupported chain
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "chain not supported")
}

func TestTryFormat_SolanaSOLTransfer(t *testing.T) {
	metaRule := NewMetaRule()

	// Test native SOL transfer (target is system program)
	rule := &types.Rule{}
	rule.Resource = "solana.send" // Meta-rule format with empty function ID
	rule.Target = &types.Target{
		TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
		Target: &types.Target_Address{
			Address: solana.SystemProgramID.String(), // Native SOL transfer
		},
	}
	rule.ParameterConstraints = []*types.ParameterConstraint{
		{
			ParameterName: "recipient",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				Value: &types.Constraint_FixedValue{
					FixedValue: testAddress,
				},
			},
		},
		{
			ParameterName: "amount",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				Value: &types.Constraint_FixedValue{
					FixedValue: "1000000",
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)

	assert.Equal(t, "solana.sol.transfer", result.GetResource())
	assert.Equal(t, types.TargetType_TARGET_TYPE_ADDRESS, result.GetTarget().GetTargetType())
	assert.Equal(t, testAddress, result.GetTarget().GetAddress())

	// Should have only amount constraint (recipient becomes the target)
	assert.Len(t, result.GetParameterConstraints(), 1)
	assert.Equal(t, "amount", result.GetParameterConstraints()[0].GetParameterName())
	assert.Equal(t, "1000000", result.GetParameterConstraints()[0].GetConstraint().GetFixedValue())
}

func TestTryFormat_SolanaSPLTokenTransfer(t *testing.T) {
	metaRule := NewMetaRule()

	// Test SPL token transfer (target is not system program)
	const tokenMintAddress = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // USDC mint
	rule := &types.Rule{
		Resource: "solana.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: tokenMintAddress, // SPL token transfer
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testAddress,
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)

	assert.Equal(t, "solana.spl_token.transfer", result.GetResource())
	assert.Equal(t, tokenMintAddress, result.GetTarget().GetAddress())

	// Should have 4 constraints for SPL token transfer
	assert.Len(t, result.GetParameterConstraints(), 4)

	constraintNames := make(map[string]*types.Constraint)
	for _, pc := range result.GetParameterConstraints() {
		constraintNames[pc.GetParameterName()] = pc.GetConstraint()
	}

	assert.Contains(t, constraintNames, "destination")
	assert.Contains(t, constraintNames, "amount")
	assert.Contains(t, constraintNames, "source")
	assert.Contains(t, constraintNames, "authority")

	// Check specific constraint values
	assert.Equal(t, testAddress, constraintNames["destination"].GetFixedValue())
	assert.Equal(t, "1000000", constraintNames["amount"].GetFixedValue())
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, constraintNames["source"].GetType())
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, constraintNames["authority"].GetType())
}

func TestTryFormat_SolanaMissingRecipientConstraint(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			// Missing recipient constraint
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse `recipient`")
}

func TestTryFormat_SolanaMissingAmountConstraint(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "RecipientAddress123",
					},
				},
			},
			// Missing amount constraint
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse `amount`")
}

func TestTryFormat_SolanaUnsupportedProtocol(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.stake", // Unsupported protocol
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol id: stake")
}

func TestTryFormat_SolanaInvalidRecipientConstraintType(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY, // Invalid for recipient
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid constraint type for `recipient`")
}
