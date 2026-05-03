package gaia_test

import (
	"testing"

	"cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vultisig/recipes/engine/cosmos/gaia"
	"github.com/vultisig/recipes/types"
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
		DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ValidatorSrcAddress: "cosmosvaloper1srcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ValidatorDstAddress: "cosmosvaloper1dstxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
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
		DelegatorAddress: "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ValidatorAddress: "cosmosvaloper1abcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
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

	const dstValidator = "cosmosvaloper1dstxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	msg := &stakingtypes.MsgBeginRedelegate{
		DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ValidatorSrcAddress: "cosmosvaloper1srcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ValidatorDstAddress: dstValidator,
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
						FixedValue: dstValidator,
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

	const valAddr = "cosmosvaloper1abcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	msg := &distributiontypes.MsgWithdrawDelegatorReward{
		DelegatorAddress: "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ValidatorAddress: valAddr,
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
						FixedValue: valAddr,
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
		FromAddress: "cosmos1fromxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ToAddress:   "cosmos1toxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
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
