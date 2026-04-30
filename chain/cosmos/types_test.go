package cosmos

import (
	"testing"

	"cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessageTypeStringForNewTypes(t *testing.T) {
	assert.Equal(t, "MsgBeginRedelegate", MessageTypeBeginRedelegate.String())
	assert.Equal(t, "MsgWithdrawDelegatorReward", MessageTypeWithdrawDelegatorReward.String())
}

func TestStakingDistributionTypeUrlConstants(t *testing.T) {
	// These exact strings are what the verifier reads off Any.TypeUrl on the wire.
	// Drift here means recipes and chain/upstream Cosmos SDK fall out of sync —
	// pin them in a test so a rename is caught here, not at runtime.
	assert.Equal(t,
		"/cosmos.staking.v1beta1.MsgBeginRedelegate",
		TypeUrlCosmosMsgBeginRedelegate,
	)
	assert.Equal(t,
		"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
		TypeUrlCosmosMsgWithdrawDelegatorReward,
	)
}

func TestNewChainRegistersStakingAndDistributionInterfaces(t *testing.T) {
	// Build a minimal Cosmos chain — same pattern as gaia.NewChain.
	chain := NewChain(ChainConfig{
		ID:           "cosmos",
		Name:         "Cosmos",
		Description:  "test",
		Bech32Prefix: "cosmos",
		Protocols:    []string{"atom"},
		MessageTypeRegistry: NewMessageTypeRegistry(map[string]MessageType{
			TypeUrlCosmosMsgSend:                           MessageTypeSend,
			TypeUrlCosmosMsgBeginRedelegate:                MessageTypeBeginRedelegate,
			TypeUrlCosmosMsgWithdrawDelegatorReward:        MessageTypeWithdrawDelegatorReward,
		}),
	})

	t.Run("MsgBeginRedelegate round-trips through chain codec", func(t *testing.T) {
		original := &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "cosmosvaloper1srcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorDstAddress: "cosmosvaloper1dstxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
		}
		any, err := codectypes.NewAnyWithValue(original)
		require.NoError(t, err)

		var decoded cosmostypes.Msg
		err = chain.Codec().UnpackAny(any, &decoded)
		require.NoError(t, err, "chain codec must know MsgBeginRedelegate")
		_, ok := decoded.(*stakingtypes.MsgBeginRedelegate)
		assert.True(t, ok, "decoded type must be *stakingtypes.MsgBeginRedelegate")
	})

	t.Run("MsgWithdrawDelegatorReward round-trips through chain codec", func(t *testing.T) {
		original := &distributiontypes.MsgWithdrawDelegatorReward{
			DelegatorAddress: "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorAddress: "cosmosvaloper1abcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		}
		any, err := codectypes.NewAnyWithValue(original)
		require.NoError(t, err)

		var decoded cosmostypes.Msg
		err = chain.Codec().UnpackAny(any, &decoded)
		require.NoError(t, err, "chain codec must know MsgWithdrawDelegatorReward")
		_, ok := decoded.(*distributiontypes.MsgWithdrawDelegatorReward)
		assert.True(t, ok, "decoded type must be *distributiontypes.MsgWithdrawDelegatorReward")
	})
}
