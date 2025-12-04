package gaia

import (
	"github.com/vultisig/recipes/chain/cosmos"
	cosmosengine "github.com/vultisig/recipes/engine/cosmos"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

// Gaia wraps the generic Cosmos engine for Cosmos Hub (GAIA).
type Gaia struct {
	engine *cosmosengine.Engine
}

// NewGaia creates a new Gaia engine.
func NewGaia() *Gaia {
	return &Gaia{
		engine: cosmosengine.NewEngine(cosmosengine.Config{
			ChainID:         "cosmos",
			SupportedChains: []common.Chain{common.GaiaChain},
			MessageTypeRegistry: cosmos.NewMessageTypeRegistry(map[string]cosmos.MessageType{
				cosmos.TypeUrlCosmosMsgSend: cosmos.MessageTypeSend,
			}),
			ProtocolMessageTypes: map[string]cosmos.MessageType{
				"atom": cosmos.MessageTypeSend,
				"send": cosmos.MessageTypeSend,
			},
		}),
	}
}

// Supports returns true if this engine supports the given chain.
func (g *Gaia) Supports(chain common.Chain) bool {
	return g.engine.Supports(chain)
}

// Evaluate validates a transaction against the given rule.
func (g *Gaia) Evaluate(rule *types.Rule, txBytes []byte) error {
	return g.engine.Evaluate(rule, txBytes)
}

