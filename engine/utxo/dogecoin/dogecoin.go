package dogecoin

import (
	"github.com/btcsuite/btcd/chaincfg"

	"github.com/vultisig/recipes/engine/utxo"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

// Dogecoin wraps the generic UTXO engine for Dogecoin.
type Dogecoin struct {
	engine *utxo.Engine
}

// NewDogecoin creates a new Dogecoin engine.
func NewDogecoin() *Dogecoin {
	dogeParams := &chaincfg.Params{
		Name:             "mainnet",
		PubKeyHashAddrID: 0x1e, // D prefix for P2PKH
		ScriptHashAddrID: 0x16, // 9 or A prefix for P2SH
	}

	return &Dogecoin{
		engine: utxo.NewEngine(utxo.Config{
			ChainID:         "dogecoin",
			SupportedChains: []common.Chain{common.Dogecoin},
			NetworkParams:   dogeParams,
		}),
	}
}

// Supports returns true if this engine supports the given chain.
func (d *Dogecoin) Supports(chain common.Chain) bool {
	return d.engine.Supports(chain)
}

// Evaluate validates a transaction against the given rule.
func (d *Dogecoin) Evaluate(rule *types.Rule, txBytes []byte) error {
	return d.engine.Evaluate(rule, txBytes)
}


