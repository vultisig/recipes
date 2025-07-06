package arbitrum

import (
	"math/big"

	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/ethereum"
	"github.com/vultisig/recipes/types"
)

// Arbitrum embeds EVMBaseChain and implements types.Chain
type Arbitrum struct {
	*ethereum.Ethereum
}

// ID returns the unique identifier for the Arbitrum chain
func (a *Arbitrum) ID() string {
	return a.Ethereum.ID()
}

// Name returns a human-readable name for the Arbitrum chain
func (a *Arbitrum) Name() string {
	return a.Ethereum.Name()
}

// Description returns a detailed description of the Arbitrum chain
func (a *Arbitrum) Description() string {
	return a.Ethereum.Description()
}

// Implement types.Chain interface methods that delegate to base
func (a *Arbitrum) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	return a.Ethereum.ParseTransaction(txHex)
}

func (a *Arbitrum) ComputeTxHash(proposedTxHex string, sigs []tss.KeysignResponse) (string, error) {
	return a.Ethereum.ComputeTxHash(proposedTxHex, sigs)
}

// Constructor
func NewArbitrum() types.Chain {
	ethChain := ethereum.NewEthereum()
	// Set Arbitrum-specific values after creation
	if eth, ok := ethChain.(*ethereum.Ethereum); ok {
		eth.SetChainID(big.NewInt(42161))
		eth.SetID("arbitrum")
		eth.SetName("Arbitrum One")
		eth.SetDescription("Arbitrum One is an Optimistic Rollup Layer 2 scaling solution for Ethereum.")
	}
	return &Arbitrum{
		Ethereum: ethChain.(*ethereum.Ethereum),
	}
}
