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
				cosmos.TypeUrlCosmosMsgSend:                    cosmos.MessageTypeSend,
				cosmos.TypeUrlCosmosMsgBeginRedelegate:         cosmos.MessageTypeBeginRedelegate,
				cosmos.TypeUrlCosmosMsgWithdrawDelegatorReward: cosmos.MessageTypeWithdrawDelegatorReward,
			}),
			ProtocolMessageTypes: map[string]cosmos.MessageType{
				"atom":                     cosmos.MessageTypeSend,
				"send":                     cosmos.MessageTypeSend,
				"staking_redelegate":       cosmos.MessageTypeBeginRedelegate,
				"staking_withdraw_rewards": cosmos.MessageTypeWithdrawDelegatorReward,
			},
			// Cosmos Hub uses "cosmos" for account addresses and "cosmosvaloper" for
			// validator operator addresses (Cosmos SDK convention: <prefix>valoper).
			// These are enforced by the engine to prevent a delegator address from
			// being silently accepted in a validator field (fail-open bypass, codex C1).
			Bech32Prefix:          "cosmos",
			ValidatorBech32Prefix: "cosmosvaloper",
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

// ExtractTxBytes extracts transaction bytes from a base64-encoded Cosmos transaction.
func (g *Gaia) ExtractTxBytes(txData string) ([]byte, error) {
	return g.engine.ExtractTxBytes(txData)
}

