package litecoin

import (
	"github.com/btcsuite/btcd/chaincfg"

	"github.com/vultisig/recipes/engine/utxo"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

// Litecoin wraps the generic UTXO engine for Litecoin.
type Litecoin struct {
	engine *utxo.Engine
}

// NewLitecoin creates a new Litecoin engine.
func NewLitecoin() *Litecoin {
	ltcParams := &chaincfg.Params{
		Name:             "mainnet",
		PubKeyHashAddrID: 0x30,  // L prefix for P2PKH
		ScriptHashAddrID: 0x32,  // M prefix for P2SH
		Bech32HRPSegwit:  "ltc", // ltc1 prefix for SegWit
	}

	return &Litecoin{
		engine: utxo.NewEngine(utxo.Config{
			ChainID:         "litecoin",
			SupportedChains: []common.Chain{common.Litecoin},
			NetworkParams:   ltcParams,
		}),
	}
}

// Supports returns true if this engine supports the given chain.
func (l *Litecoin) Supports(chain common.Chain) bool {
	return l.engine.Supports(chain)
}

// Evaluate validates a transaction against the given rule.
func (l *Litecoin) Evaluate(rule *types.Rule, txBytes []byte) error {
	return l.engine.Evaluate(rule, txBytes)
}

