// Package bitcoincash provides the Bitcoin Cash UTXO engine with CashAddr support.
package bitcoincash

import (
	bchchain "github.com/vultisig/recipes/chain/utxo/bitcoincash"
	"github.com/vultisig/recipes/engine/utxo"
	"github.com/vultisig/vultisig-go/common"
)

// NewBitcoinCash creates a new Bitcoin Cash engine using the generic UTXO engine
// with a custom address extractor for CashAddr format.
func NewBitcoinCash() *utxo.Engine {
	return utxo.NewEngine(utxo.Config{
		ChainID:         "bitcoin-cash",
		SupportedChains: []common.Chain{common.BitcoinCash},
		NetworkParams:   nil, // Not used - we use custom ExtractAddress
		ExtractAddress:  bchchain.ExtractCashAddress,
	})
}
