package maya

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/vultisig/recipes/chain/cosmos"
	cosmosengine "github.com/vultisig/recipes/engine/cosmos"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

// Maya wraps the generic Cosmos engine for MAYAChain.
type Maya struct {
	engine *cosmosengine.Engine
}

// NewMaya creates a new Maya engine.
func NewMaya() *Maya {
	return &Maya{
		engine: cosmosengine.NewEngine(cosmosengine.Config{
			ChainID:         "mayachain",
			SupportedChains: []common.Chain{common.MayaChain},
			MessageTypeRegistry: cosmos.NewMessageTypeRegistry(map[string]cosmos.MessageType{
				cosmos.TypeUrlCosmosMsgSend:    cosmos.MessageTypeSend,
				cosmos.TypeUrlCustomMsgSend:    cosmos.MessageTypeSend,
				cosmos.TypeUrlCustomMsgDeposit: cosmos.MessageTypeDeposit,
			}),
			ProtocolMessageTypes: map[string]cosmos.MessageType{
				"cacao":          cosmos.MessageTypeSend,
				"send":           cosmos.MessageTypeSend,
				"mayachain_swap": cosmos.MessageTypeDeposit,
			},
			RegisterExtraTypes: func(ir codectypes.InterfaceRegistry) {
				ir.RegisterImplementations((*sdk.Msg)(nil), &types.MsgDeposit{})
			},
		}),
	}
}

// Supports returns true if this engine supports the given chain.
func (m *Maya) Supports(chain common.Chain) bool {
	return m.engine.Supports(chain)
}

// Evaluate validates a transaction against the given rule.
func (m *Maya) Evaluate(rule *types.Rule, txBytes []byte) error {
	return m.engine.Evaluate(rule, txBytes)
}

