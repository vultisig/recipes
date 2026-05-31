package metarule

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/recipes/types"
)

const (
	// USDC on Solana (EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v)
	testUSDCMint = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"
	// PENGU on Solana (2zMMhcVQEXDtdE6vsFS7S7D5oUodfJHE8vd1gnBouauv)
	testPENGUMint = "2zMMhcVQEXDtdE6vsFS7S7D5oUodfJHE8vd1gnBouauv"
	// Jupiter DCA program address
	jupiterDCAProgram = "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M"
)

// TestTryFormat_SolanaDCA_USDCtoPENGU validates that a "solana.dca" meta-rule
// for USDC → PENGU (ETH longtail) correctly expands to jupiter_dca.openDcaV2
// rules. This is the root cause of issue #483.
func TestTryFormat_SolanaDCA_USDCtoPENGU(t *testing.T) {
	metaRule := NewMetaRule()

	// 1000 USDC total (6 decimals), 100 USDC per cycle, every 86400s (daily)
	rule := &types.Rule{
		Resource: "solana.dca",
		Effect:   types.Effect_EFFECT_ALLOW,
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testAddress, // "4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM"
					},
				},
			},
			{
				ParameterName: "from_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testUSDCMint,
					},
				},
			},
			{
				ParameterName: "in_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000000", // 1000 USDC in base units
					},
				},
			},
			{
				ParameterName: "in_amount_per_cycle",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "100000000", // 100 USDC per cycle
					},
				},
			},
			{
				ParameterName: "cycle_frequency",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "86400", // 1 day in seconds
					},
				},
			},
			{
				ParameterName: "to_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testPENGUMint,
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err, "solana.dca USDC->PENGU must not return an error (was: unsupported protocol id)")
	require.NotEmpty(t, result)

	// Must include the openDcaV2 rule targeting the Jupiter DCA program
	var dcaRule *types.Rule
	for _, r := range result {
		if r.Resource == "solana.jupiter_dca.openDcaV2" {
			dcaRule = r
			break
		}
	}
	require.NotNil(t, dcaRule, "result must contain solana.jupiter_dca.openDcaV2 rule")
	assert.Equal(t, types.Effect_EFFECT_ALLOW, dcaRule.Effect)
	assert.Equal(t, jupiterDCAProgram, dcaRule.Target.GetAddress())

	// Verify parameter constraints
	paramByName := make(map[string]*types.ParameterConstraint)
	for _, pc := range dcaRule.ParameterConstraints {
		paramByName[pc.ParameterName] = pc
	}

	// inputMint must be USDC
	inputMint, ok := paramByName["account_inputMint"]
	require.True(t, ok, "account_inputMint constraint must be present")
	assert.Equal(t, testUSDCMint, inputMint.Constraint.GetFixedValue())

	// outputMint must be PENGU
	outputMint, ok := paramByName["account_outputMint"]
	require.True(t, ok, "account_outputMint constraint must be present")
	assert.Equal(t, testPENGUMint, outputMint.Constraint.GetFixedValue())

	// inAmount must match
	inAmount, ok := paramByName["arg_inAmount"]
	require.True(t, ok, "arg_inAmount constraint must be present")
	assert.Equal(t, "1000000000", inAmount.Constraint.GetFixedValue())

	// inAmountPerCycle must match
	inAmountPerCycle, ok := paramByName["arg_inAmountPerCycle"]
	require.True(t, ok, "arg_inAmountPerCycle constraint must be present")
	assert.Equal(t, "100000000", inAmountPerCycle.Constraint.GetFixedValue())

	// cycleFrequency must match
	cycleFreq, ok := paramByName["arg_cycleFrequency"]
	require.True(t, ok, "arg_cycleFrequency constraint must be present")
	assert.Equal(t, "86400", cycleFreq.Constraint.GetFixedValue())

	// user/payer must be locked to from_address
	user, ok := paramByName["account_user"]
	require.True(t, ok)
	assert.Equal(t, testAddress, user.Constraint.GetFixedValue())

	// Must also include an ATA-create rule for the user's input (USDC) token account
	hasATARule := false
	for _, r := range result {
		if r.Resource == "solana.associated_token_account.create" {
			for _, pc := range r.ParameterConstraints {
				if pc.ParameterName == "account_mint" && pc.Constraint.GetFixedValue() == testUSDCMint {
					hasATARule = true
					break
				}
			}
		}
	}
	assert.True(t, hasATARule, "result must include an ATA-create rule for the input (USDC) mint")
}

// TestTryFormat_SolanaDCA_MissingConstraint ensures that missing required DCA
// constraints are caught early with a descriptive error.
func TestTryFormat_SolanaDCA_MissingConstraint(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.dca",
		Effect:   types.Effect_EFFECT_ALLOW,
		// from_address is intentionally missing
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "in_amount",
				Constraint: &types.Constraint{
					Type:  types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{FixedValue: "1000000000"},
				},
			},
			{
				ParameterName: "in_amount_per_cycle",
				Constraint: &types.Constraint{
					Type:  types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{FixedValue: "100000000"},
				},
			},
			{
				ParameterName: "cycle_frequency",
				Constraint: &types.Constraint{
					Type:  types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{FixedValue: "86400"},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	require.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "from_address"), "error must mention the missing constraint")
}

// TestTryFormat_SolanaDCA_UnsupportedProtocol verifies that an unknown Solana
// protocol still returns an error (regression guard for the default case).
func TestTryFormat_SolanaDCA_UnsupportedProtocol(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.stake",
		Effect:   types.Effect_EFFECT_ALLOW,
	}

	_, err := metaRule.TryFormat(rule)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol id")
}
