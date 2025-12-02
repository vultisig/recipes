package bitcoin

import (
	"github.com/btcsuite/btcd/chaincfg"

	"github.com/vultisig/recipes/engine/utxo"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

// Btc wraps the generic UTXO engine for Bitcoin.
type Btc struct {
	engine *utxo.Engine
}

// NewBtc creates a new Bitcoin engine.
func NewBtc() *Btc {
	return &Btc{
		engine: utxo.NewEngine(utxo.Config{
			ChainID:         "bitcoin",
			SupportedChains: []common.Chain{common.Bitcoin},
			NetworkParams:   &chaincfg.MainNetParams,
		}),
	}
}

// Supports returns true if this engine supports the given chain.
func (b *Btc) Supports(chain common.Chain) bool {
	return b.engine.Supports(chain)
}

// Evaluate validates a transaction against the given rule.
func (b *Btc) Evaluate(rule *types.Rule, txBytes []byte) error {
	return b.engine.Evaluate(rule, txBytes)
}
