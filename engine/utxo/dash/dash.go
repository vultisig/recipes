package dash

import (
	chaindash "github.com/vultisig/recipes/chain/utxo/dash"
	"github.com/vultisig/recipes/engine/utxo"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

// Dash wraps the generic UTXO engine for Dash.
type Dash struct {
	engine *utxo.Engine
}

// NewDash creates a new Dash engine.
func NewDash() *Dash {
	return &Dash{
		engine: utxo.NewEngine(utxo.Config{
			ChainID:         "dash",
			SupportedChains: []common.Chain{common.Dash},
			NetworkParams:   chaindash.DashMainNetParams,
		}),
	}
}

// Supports returns true if this engine supports the given chain.
func (d *Dash) Supports(chain common.Chain) bool {
	return d.engine.Supports(chain)
}

// Evaluate validates a transaction against the given rule.
func (d *Dash) Evaluate(rule *types.Rule, txBytes []byte) error {
	return d.engine.Evaluate(rule, txBytes)
}

// ExtractTxBytes extracts transaction bytes from a PSBT string.
func (d *Dash) ExtractTxBytes(txData string) ([]byte, error) {
	return d.engine.ExtractTxBytes(txData)
}
