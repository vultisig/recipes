package cosmos

import (
	"math/big"
	"testing"

	"cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/vultisig-go/common"

	"github.com/vultisig/recipes/chain/cosmos"
)

// newTestEngine creates a generic Cosmos engine wired for staking + distribution
// in addition to the default bank MsgSend mappings.
func newTestEngine() *Engine {
	return NewEngine(Config{
		ChainID:         "cosmos",
		SupportedChains: []common.Chain{common.GaiaChain},
		MessageTypeRegistry: cosmos.NewMessageTypeRegistry(map[string]cosmos.MessageType{
			cosmos.TypeUrlCosmosMsgSend:                     cosmos.MessageTypeSend,
			cosmos.TypeUrlCosmosMsgBeginRedelegate:          cosmos.MessageTypeBeginRedelegate,
			cosmos.TypeUrlCosmosMsgWithdrawDelegatorReward:  cosmos.MessageTypeWithdrawDelegatorReward,
		}),
		ProtocolMessageTypes: map[string]cosmos.MessageType{
			"atom":                       cosmos.MessageTypeSend,
			"staking_redelegate":         cosmos.MessageTypeBeginRedelegate,
			"staking_withdraw_rewards":   cosmos.MessageTypeWithdrawDelegatorReward,
		},
	})
}

// packAny wraps a sdk message into an Any using the engine's codec, exercising the
// real interface registry that the engine builds in NewEngine.
func packAny(t *testing.T, e *Engine, msg cosmostypes.Msg) *codectypes.Any {
	t.Helper()
	any, err := codectypes.NewAnyWithValue(msg)
	require.NoError(t, err)
	// Round-trip through the engine codec to confirm the registry knows the type.
	require.NoError(t, e.cdc.UnpackAny(any, new(cosmostypes.Msg)))
	return any
}

func TestExtractParameterFromMsgBeginRedelegate(t *testing.T) {
	e := newTestEngine()

	t.Run("happy path returns all params", func(t *testing.T) {
		msg := &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "cosmosvaloper1srcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorDstAddress: "cosmosvaloper1dstxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
		}

		got, err := e.extractParameterFromMsgBeginRedelegate("delegator_address", msg)
		require.NoError(t, err)
		assert.Equal(t, msg.DelegatorAddress, got)

		got, err = e.extractParameterFromMsgBeginRedelegate("validator_src_address", msg)
		require.NoError(t, err)
		assert.Equal(t, msg.ValidatorSrcAddress, got)

		got, err = e.extractParameterFromMsgBeginRedelegate("validator_dst_address", msg)
		require.NoError(t, err)
		assert.Equal(t, msg.ValidatorDstAddress, got)

		got, err = e.extractParameterFromMsgBeginRedelegate("amount", msg)
		require.NoError(t, err)
		assert.Equal(t, big.NewInt(1_000_000), got)

		got, err = e.extractParameterFromMsgBeginRedelegate("denom", msg)
		require.NoError(t, err)
		assert.Equal(t, "uatom", got)
	})

	t.Run("rejects same src and dst validator", func(t *testing.T) {
		msg := &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "cosmosvaloper1samexxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorDstAddress: "cosmosvaloper1samexxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
		}

		_, err := e.extractParameterFromMsgBeginRedelegate("amount", msg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "src and dst validators must differ")
	})

	// Regression: previously the equality check was guarded by `Src != ""`, which let
	// a fully empty pair ("" == "") slip past validation. Both empty addresses must
	// each be rejected on their own, before the equality comparison runs.
	t.Run("rejects both empty validator addresses", func(t *testing.T) {
		msg := &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "",
			ValidatorDstAddress: "",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
		}

		_, err := e.extractParameterFromMsgBeginRedelegate("amount", msg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "validator_src_address required")
	})

	t.Run("rejects empty src validator address", func(t *testing.T) {
		msg := &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "",
			ValidatorDstAddress: "cosmosvaloper1dstxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
		}

		_, err := e.extractParameterFromMsgBeginRedelegate("amount", msg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "validator_src_address required")
	})

	t.Run("rejects empty dst validator address", func(t *testing.T) {
		msg := &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "cosmosvaloper1srcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorDstAddress: "",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
		}

		_, err := e.extractParameterFromMsgBeginRedelegate("amount", msg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "validator_dst_address required")
	})

	t.Run("rejects zero amount", func(t *testing.T) {
		msg := &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "cosmosvaloper1srcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorDstAddress: "cosmosvaloper1dstxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(0)),
		}

		_, err := e.extractParameterFromMsgBeginRedelegate("amount", msg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "amount must be > 0")
	})

	t.Run("rejects negative amount", func(t *testing.T) {
		msg := &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "cosmosvaloper1srcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorDstAddress: "cosmosvaloper1dstxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Amount:              cosmostypes.Coin{Denom: "uatom", Amount: math.NewInt(-1)},
		}

		_, err := e.extractParameterFromMsgBeginRedelegate("amount", msg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "amount must be > 0")
	})

	t.Run("rejects unknown parameter", func(t *testing.T) {
		msg := &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "cosmosvaloper1srcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorDstAddress: "cosmosvaloper1dstxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
		}

		_, err := e.extractParameterFromMsgBeginRedelegate("memo", msg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported parameter")
	})
}

func TestExtractParameterFromMsgWithdrawDelegatorReward(t *testing.T) {
	e := newTestEngine()

	t.Run("happy path returns delegator + validator", func(t *testing.T) {
		msg := &distributiontypes.MsgWithdrawDelegatorReward{
			DelegatorAddress: "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorAddress: "cosmosvaloper1abcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		}

		got, err := e.extractParameterFromMsgWithdrawDelegatorReward("delegator_address", msg)
		require.NoError(t, err)
		assert.Equal(t, msg.DelegatorAddress, got)

		got, err = e.extractParameterFromMsgWithdrawDelegatorReward("validator_address", msg)
		require.NoError(t, err)
		assert.Equal(t, msg.ValidatorAddress, got)
	})

	t.Run("rejects unknown parameter", func(t *testing.T) {
		msg := &distributiontypes.MsgWithdrawDelegatorReward{
			DelegatorAddress: "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorAddress: "cosmosvaloper1abcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		}

		_, err := e.extractParameterFromMsgWithdrawDelegatorReward("amount", msg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported parameter")
	})

	// Regression: empty addresses must be rejected before the parameter switch
	// runs, mirroring the redelegate validator-address checks. Without these
	// guards, a policy match could silently succeed against an empty string.
	t.Run("rejects empty delegator address", func(t *testing.T) {
		msg := &distributiontypes.MsgWithdrawDelegatorReward{
			DelegatorAddress: "",
			ValidatorAddress: "cosmosvaloper1abcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		}

		_, err := e.extractParameterFromMsgWithdrawDelegatorReward("delegator_address", msg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "delegator_address required")
	})

	t.Run("rejects empty validator address", func(t *testing.T) {
		msg := &distributiontypes.MsgWithdrawDelegatorReward{
			DelegatorAddress: "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorAddress: "",
		}

		_, err := e.extractParameterFromMsgWithdrawDelegatorReward("validator_address", msg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "validator_address required")
	})

	t.Run("rejects both empty addresses", func(t *testing.T) {
		msg := &distributiontypes.MsgWithdrawDelegatorReward{
			DelegatorAddress: "",
			ValidatorAddress: "",
		}

		_, err := e.extractParameterFromMsgWithdrawDelegatorReward("delegator_address", msg)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "delegator_address required")
	})
}

func TestUnpackMsgBeginRedelegate(t *testing.T) {
	e := newTestEngine()

	t.Run("unpacks a real Any wrapped MsgBeginRedelegate", func(t *testing.T) {
		original := &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "cosmosvaloper1srcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorDstAddress: "cosmosvaloper1dstxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1_000_000)),
		}
		any := packAny(t, e, original)
		assert.Equal(t, cosmos.TypeUrlCosmosMsgBeginRedelegate, any.TypeUrl,
			"TypeUrl should match the canonical Cosmos SDK staking module url")

		got, err := e.unpackMsgBeginRedelegate(any)
		require.NoError(t, err)
		assert.Equal(t, original.DelegatorAddress, got.DelegatorAddress)
		assert.Equal(t, original.ValidatorSrcAddress, got.ValidatorSrcAddress)
		assert.Equal(t, original.ValidatorDstAddress, got.ValidatorDstAddress)
		assert.Equal(t, original.Amount.Denom, got.Amount.Denom)
		assert.True(t, original.Amount.Amount.Equal(got.Amount.Amount))
	})

	t.Run("rejects nil message", func(t *testing.T) {
		_, err := e.unpackMsgBeginRedelegate(nil)
		require.Error(t, err)
	})

	t.Run("rejects wrong message type", func(t *testing.T) {
		wrong := &distributiontypes.MsgWithdrawDelegatorReward{
			DelegatorAddress: "cosmos1xxx",
			ValidatorAddress: "cosmosvaloper1xxx",
		}
		any := packAny(t, e, wrong)
		_, err := e.unpackMsgBeginRedelegate(any)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "expected staking MsgBeginRedelegate")
	})
}

func TestUnpackMsgWithdrawDelegatorReward(t *testing.T) {
	e := newTestEngine()

	t.Run("unpacks a real Any wrapped MsgWithdrawDelegatorReward", func(t *testing.T) {
		original := &distributiontypes.MsgWithdrawDelegatorReward{
			DelegatorAddress: "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorAddress: "cosmosvaloper1abcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		}
		any := packAny(t, e, original)
		assert.Equal(t, cosmos.TypeUrlCosmosMsgWithdrawDelegatorReward, any.TypeUrl,
			"TypeUrl should match the canonical Cosmos SDK distribution module url")

		got, err := e.unpackMsgWithdrawDelegatorReward(any)
		require.NoError(t, err)
		assert.Equal(t, original.DelegatorAddress, got.DelegatorAddress)
		assert.Equal(t, original.ValidatorAddress, got.ValidatorAddress)
	})

	t.Run("rejects nil message", func(t *testing.T) {
		_, err := e.unpackMsgWithdrawDelegatorReward(nil)
		require.Error(t, err)
	})

	t.Run("rejects wrong message type", func(t *testing.T) {
		wrong := &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1xxx",
			ValidatorSrcAddress: "cosmosvaloper1srcxxx",
			ValidatorDstAddress: "cosmosvaloper1dstxxx",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1)),
		}
		any := packAny(t, e, wrong)
		_, err := e.unpackMsgWithdrawDelegatorReward(any)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "expected distribution MsgWithdrawDelegatorReward")
	})
}

func TestDetectMessageTypeForStakingAndDistribution(t *testing.T) {
	e := newTestEngine()

	t.Run("MsgBeginRedelegate", func(t *testing.T) {
		any := packAny(t, e, &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "cosmosvaloper1srcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorDstAddress: "cosmosvaloper1dstxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(1)),
		})
		mt, err := e.detectMessageType(any)
		require.NoError(t, err)
		assert.Equal(t, cosmos.MessageTypeBeginRedelegate, mt)
	})

	t.Run("MsgWithdrawDelegatorReward", func(t *testing.T) {
		any := packAny(t, e, &distributiontypes.MsgWithdrawDelegatorReward{
			DelegatorAddress: "cosmos1xxx",
			ValidatorAddress: "cosmosvaloper1xxx",
		})
		mt, err := e.detectMessageType(any)
		require.NoError(t, err)
		assert.Equal(t, cosmos.MessageTypeWithdrawDelegatorReward, mt)
	})
}

// TestExtractParameterValueDispatchesToNewMsgTypes confirms the public dispatch
// path picks up the new message types — this is the single biggest behavior gap
// the schemas close: without this, the verifier engine had no way to read params
// off a redelegate / withdraw_rewards tx.
func TestExtractParameterValueDispatchesToNewMsgTypes(t *testing.T) {
	e := newTestEngine()

	t.Run("MsgBeginRedelegate routes through new extractor", func(t *testing.T) {
		any := packAny(t, e, &stakingtypes.MsgBeginRedelegate{
			DelegatorAddress:    "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorSrcAddress: "cosmosvaloper1srcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorDstAddress: "cosmosvaloper1dstxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Amount:              cosmostypes.NewCoin("uatom", math.NewInt(2_500_000)),
		})
		txData := &tx.Tx{Body: &tx.TxBody{Messages: []*codectypes.Any{any}}}

		got, err := e.extractParameterValue("amount", txData, cosmos.MessageTypeBeginRedelegate)
		require.NoError(t, err)
		assert.Equal(t, big.NewInt(2_500_000), got)
	})

	t.Run("MsgWithdrawDelegatorReward routes through new extractor", func(t *testing.T) {
		any := packAny(t, e, &distributiontypes.MsgWithdrawDelegatorReward{
			DelegatorAddress: "cosmos1delegatorxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ValidatorAddress: "cosmosvaloper1abcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		})
		txData := &tx.Tx{Body: &tx.TxBody{Messages: []*codectypes.Any{any}}}

		got, err := e.extractParameterValue("validator_address", txData, cosmos.MessageTypeWithdrawDelegatorReward)
		require.NoError(t, err)
		assert.Equal(t, "cosmosvaloper1abcxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", got)
	})
}
