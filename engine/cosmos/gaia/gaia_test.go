package gaia_test

import (
	"strings"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vultisig/recipes/engine/cosmos/gaia"
	"github.com/vultisig/recipes/types"
)

// Valid bech32 test addresses (deterministic, not real keys).
// Generated with bech32.ConvertAndEncode over fixed 20-byte payloads.
const (
	// Account (delegator) addresses — HRP "cosmos".
	testDelegator1 = "cosmos1uzqrd7nsnmvwjflffw48gs6u8a3n5ugrvs2dlk"
	testDelegator2 = "cosmos1e5mzm3769s7ksr64lw4533yc88sgne68gt3nuy"

	// Validator operator addresses — HRP "cosmosvaloper".
	testValSrc = "cosmosvaloper14dpelgwgj5pnhkqwf05sgeqdjas52k2qqsptjp"
	testValDst = "cosmosvaloper18w90vkkuhm2a39m8u63uh6c89kj00v6ukg4jqc"
	testValAbc = "cosmosvaloper1mn7wzsk6h3aqcxeuyl9vye4w08rnnstnmk4j66"

	// Wrong-HRP addresses used in negative tests.
	// These are valid bech32 but carry the wrong prefix for the field they are
	// placed in — proving that the extractor rejects them.
	testAccountAsValidator   = "cosmos1gvll0hewp46cvm8wm6hqysyh2cve38j5k7npf7"        // "cosmos" HRP where "cosmosvaloper" is required
	testValidatorAsDelegator = "cosmosvaloper1d8dg727hwgyqq6h34nwysvsm3fqkvcs3fvndue" // "cosmosvaloper" HRP where "cosmos" is required
)

// buildTestCodec creates a codec that can pack Any messages for test tx construction.
// We need our own codec here because gaia.Engine's codec is unexported.
func buildTestCodec() codec.Codec {
	ir := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(ir)
	banktypes.RegisterInterfaces(ir)
	stakingtypes.RegisterInterfaces(ir)
	distributiontypes.RegisterInterfaces(ir)
	return codec.NewProtoCodec(ir)
}

// marshalTx serialises a single-message tx to proto bytes.
func marshalTx(t *testing.T, cdc codec.Codec, msg cosmostypes.Msg) []byte {
	t.Helper()
	any, err := codectypes.NewAnyWithValue(msg)
	require.NoError(t, err)
	tx := &sdktx.Tx{
		Body:     &sdktx.TxBody{Messages: []*codectypes.Any{any}},
		AuthInfo: &sdktx.AuthInfo{},
	}
	raw, err := cdc.Marshal(tx)
	require.NoError(t, err)
	return raw
}

// TestNewGaia_MsgBeginRedelegate_Dispatch verifies that gaia.NewGaia().Evaluate
// does NOT return "unsupported TypeUrl" for MsgBeginRedelegate — the core of B1.
func TestNewGaia_MsgBeginRedelegate_Dispatch(t *testing.T) {
	g := gaia.NewGaia()
	cdc := buildTestCodec()

	msg := &stakingtypes.MsgBeginRedelegate{
		DelegatorAddress:    testDelegator1,
		ValidatorSrcAddress: testValSrc,
		ValidatorDstAddress: testValDst,
		Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
	}
	txBytes := marshalTx(t, cdc, msg)

	rule := &types.Rule{
		Resource: "cosmos.staking_redelegate.redelegate",
		Effect:   types.Effect_EFFECT_ALLOW,
	}

	err := g.Evaluate(rule, txBytes)
	// The dispatch gate must pass (no "unsupported TypeUrl" or "unsupported protocol").
	// Parameter constraints are empty so validation stops at dispatch — that is enough
	// to prove the wiring is correct.
	require.NoError(t, err, "MsgBeginRedelegate must not hit unsupported TypeUrl in gaia.NewGaia()")
}

// TestNewGaia_MsgWithdrawDelegatorReward_Dispatch verifies the same for
// MsgWithdrawDelegatorReward.
func TestNewGaia_MsgWithdrawDelegatorReward_Dispatch(t *testing.T) {
	g := gaia.NewGaia()
	cdc := buildTestCodec()

	msg := &distributiontypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: testDelegator1,
		ValidatorAddress: testValAbc,
	}
	txBytes := marshalTx(t, cdc, msg)

	rule := &types.Rule{
		Resource: "cosmos.staking_withdraw_rewards.withdraw_rewards",
		Effect:   types.Effect_EFFECT_ALLOW,
	}

	err := g.Evaluate(rule, txBytes)
	require.NoError(t, err, "MsgWithdrawDelegatorReward must not hit unsupported TypeUrl in gaia.NewGaia()")
}

// TestNewGaia_MsgBeginRedelegate_WithConstraint exercises the full dispatch path
// including a parameter constraint on validator_dst_address, confirming that
// NewGaia() routes redelegate params all the way through extractParameterValue.
func TestNewGaia_MsgBeginRedelegate_WithConstraint(t *testing.T) {
	g := gaia.NewGaia()
	cdc := buildTestCodec()

	msg := &stakingtypes.MsgBeginRedelegate{
		DelegatorAddress:    testDelegator1,
		ValidatorSrcAddress: testValSrc,
		ValidatorDstAddress: testValDst,
		Amount:              cosmostypes.NewCoin("uatom", math.NewInt(500_000)),
	}
	txBytes := marshalTx(t, cdc, msg)

	rule := &types.Rule{
		Resource: "cosmos.staking_redelegate.redelegate",
		Effect:   types.Effect_EFFECT_ALLOW,
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "validator_dst_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testValDst,
					},
				},
			},
		},
	}

	err := g.Evaluate(rule, txBytes)
	require.NoError(t, err, "constraint on validator_dst_address must pass through gaia.NewGaia()")
}

// TestNewGaia_MsgWithdrawDelegatorReward_WithConstraint exercises the full dispatch
// path for withdraw_rewards with a validator_address constraint.
func TestNewGaia_MsgWithdrawDelegatorReward_WithConstraint(t *testing.T) {
	g := gaia.NewGaia()
	cdc := buildTestCodec()

	msg := &distributiontypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: testDelegator1,
		ValidatorAddress: testValAbc,
	}
	txBytes := marshalTx(t, cdc, msg)

	rule := &types.Rule{
		Resource: "cosmos.staking_withdraw_rewards.withdraw_rewards",
		Effect:   types.Effect_EFFECT_ALLOW,
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "validator_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testValAbc,
					},
				},
			},
		},
	}

	err := g.Evaluate(rule, txBytes)
	require.NoError(t, err, "constraint on validator_address must pass through gaia.NewGaia()")
}

// TestNewGaia_UnknownProtocol_Fails confirms that an unknown protocol ID still
// returns an error, so we haven't broken the gate by being too permissive.
func TestNewGaia_UnknownProtocol_Fails(t *testing.T) {
	g := gaia.NewGaia()
	cdc := buildTestCodec()

	msg := &banktypes.MsgSend{
		FromAddress: testDelegator1,
		ToAddress:   testDelegator2,
		Amount:      cosmostypes.NewCoins(cosmostypes.NewCoin("uatom", math.NewInt(1))),
	}
	txBytes := marshalTx(t, cdc, msg)

	rule := &types.Rule{
		Resource: "cosmos.unknown_protocol.transfer",
		Effect:   types.Effect_EFFECT_ALLOW,
	}

	err := g.Evaluate(rule, txBytes)
	assert.Error(t, err, "unknown protocol must still return an error")
}

// --- Negative tests (C3): verify that wrong-HRP addresses are REJECTED ---

// TestNewGaia_RedelegateRejectsAccountHRPInSrcValidatorField proves that a
// cosmos1... (account HRP) address in validator_src_address is rejected.
// A rule constraining validator_src_address to cosmosvaloper1... would otherwise
// silently accept a cosmos1... address — a fail-open bypass (codex C1).
func TestNewGaia_RedelegateRejectsAccountHRPInSrcValidatorField(t *testing.T) {
	g := gaia.NewGaia()
	cdc := buildTestCodec()

	msg := &stakingtypes.MsgBeginRedelegate{
		DelegatorAddress:    testDelegator1,
		ValidatorSrcAddress: testAccountAsValidator, // cosmos1... where cosmosvaloper1... required
		ValidatorDstAddress: testValDst,
		Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
	}
	txBytes := marshalTx(t, cdc, msg)

	rule := &types.Rule{
		Resource: "cosmos.staking_redelegate.redelegate",
		Effect:   types.Effect_EFFECT_ALLOW,
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "validator_src_address",
				Constraint: &types.Constraint{
					Type:  types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{FixedValue: testValSrc},
				},
			},
		},
	}

	err := g.Evaluate(rule, txBytes)
	require.Error(t, err, "account HRP in validator_src_address must be rejected")
	assert.True(t, strings.Contains(err.Error(), "validator_src_address") || strings.Contains(err.Error(), "HRP"),
		"error should mention the field or HRP mismatch, got: %s", err)
}

// TestNewGaia_RedelegateRejectsAccountHRPInDstValidatorField proves the same
// for validator_dst_address.
func TestNewGaia_RedelegateRejectsAccountHRPInDstValidatorField(t *testing.T) {
	g := gaia.NewGaia()
	cdc := buildTestCodec()

	msg := &stakingtypes.MsgBeginRedelegate{
		DelegatorAddress:    testDelegator1,
		ValidatorSrcAddress: testValSrc,
		ValidatorDstAddress: testAccountAsValidator, // cosmos1... where cosmosvaloper1... required
		Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
	}
	txBytes := marshalTx(t, cdc, msg)

	rule := &types.Rule{
		Resource: "cosmos.staking_redelegate.redelegate",
		Effect:   types.Effect_EFFECT_ALLOW,
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "validator_dst_address",
				Constraint: &types.Constraint{
					Type:  types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{FixedValue: testValDst},
				},
			},
		},
	}

	err := g.Evaluate(rule, txBytes)
	require.Error(t, err, "account HRP in validator_dst_address must be rejected")
	assert.True(t, strings.Contains(err.Error(), "validator_dst_address") || strings.Contains(err.Error(), "HRP"),
		"error should mention the field or HRP mismatch, got: %s", err)
}

// TestNewGaia_WithdrawRejectsAccountHRPInValidatorField proves that a cosmos1...
// address in MsgWithdrawDelegatorReward.ValidatorAddress is rejected.
func TestNewGaia_WithdrawRejectsAccountHRPInValidatorField(t *testing.T) {
	g := gaia.NewGaia()
	cdc := buildTestCodec()

	msg := &distributiontypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: testDelegator1,
		ValidatorAddress: testAccountAsValidator, // cosmos1... where cosmosvaloper1... required
	}
	txBytes := marshalTx(t, cdc, msg)

	rule := &types.Rule{
		Resource: "cosmos.staking_withdraw_rewards.withdraw_rewards",
		Effect:   types.Effect_EFFECT_ALLOW,
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "validator_address",
				Constraint: &types.Constraint{
					Type:  types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{FixedValue: testValAbc},
				},
			},
		},
	}

	err := g.Evaluate(rule, txBytes)
	require.Error(t, err, "account HRP in validator_address must be rejected")
	assert.True(t, strings.Contains(err.Error(), "validator_address") || strings.Contains(err.Error(), "HRP"),
		"error should mention the field or HRP mismatch, got: %s", err)
}

// TestNewGaia_RedelegateRejectsValidatorHRPInDelegatorField is a symmetric
// defensive test: a cosmosvaloper1... address in DelegatorAddress must also be
// rejected. This closes the reverse confusion direction.
func TestNewGaia_RedelegateRejectsValidatorHRPInDelegatorField(t *testing.T) {
	g := gaia.NewGaia()
	cdc := buildTestCodec()

	msg := &stakingtypes.MsgBeginRedelegate{
		DelegatorAddress:    testValidatorAsDelegator, // cosmosvaloper1... where cosmos1... required
		ValidatorSrcAddress: testValSrc,
		ValidatorDstAddress: testValDst,
		Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
	}
	txBytes := marshalTx(t, cdc, msg)

	rule := &types.Rule{
		Resource: "cosmos.staking_redelegate.redelegate",
		Effect:   types.Effect_EFFECT_ALLOW,
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "delegator_address",
				Constraint: &types.Constraint{
					Type:  types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{FixedValue: testDelegator1},
				},
			},
		},
	}

	err := g.Evaluate(rule, txBytes)
	require.Error(t, err, "validator HRP in delegator_address must be rejected")
	assert.True(t, strings.Contains(err.Error(), "delegator_address") || strings.Contains(err.Error(), "HRP"),
		"error should mention the field or HRP mismatch, got: %s", err)
}
