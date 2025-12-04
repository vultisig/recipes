package thorchain

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/vultisig/recipes/chain/cosmos"
	cosmosengine "github.com/vultisig/recipes/engine/cosmos"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

// Thorchain wraps the generic Cosmos engine for THORChain.
type Thorchain struct {
	engine *cosmosengine.Engine
}

// NewThorchain creates a new Thorchain engine.
func NewThorchain() *Thorchain {
	return &Thorchain{
		engine: cosmosengine.NewEngine(cosmosengine.Config{
			ChainID:         "thorchain",
			SupportedChains: []common.Chain{common.THORChain},
			MessageTypeRegistry: cosmos.NewMessageTypeRegistry(map[string]cosmos.MessageType{
				cosmos.TypeUrlCosmosMsgSend:    cosmos.MessageTypeSend,
				cosmos.TypeUrlCustomMsgSend:    cosmos.MessageTypeSend,
				cosmos.TypeUrlCustomMsgDeposit: cosmos.MessageTypeDeposit,
			}),
			ProtocolMessageTypes: map[string]cosmos.MessageType{
				"rune":           cosmos.MessageTypeSend,
				"send":           cosmos.MessageTypeSend,
				"thorchain_swap": cosmos.MessageTypeDeposit,
			},
			RegisterExtraTypes: func(ir codectypes.InterfaceRegistry) {
				ir.RegisterImplementations((*sdk.Msg)(nil), &types.MsgDeposit{})
			},
		}),
	}
}

// Supports returns true if this engine supports the given chain.
func (t *Thorchain) Supports(chain common.Chain) bool {
	return t.engine.Supports(chain)
}

// Evaluate validates a transaction against the given rule.
func (t *Thorchain) Evaluate(rule *types.Rule, txBytes []byte) error {
	return t.engine.Evaluate(rule, txBytes)
}

