package base

import (
	"math/big"

	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/ethereum"
	"github.com/vultisig/recipes/types"
)

// Base embeds Ethereum and implements types.Chain
type Base struct {
	*ethereum.Ethereum
}

// ID returns the unique identifier for the Base chain
func (b *Base) ID() string {
	return b.Ethereum.ID()
}

// Name returns a human-readable name for the Base chain
func (b *Base) Name() string {
	return b.Ethereum.Name()
}

// Description returns a detailed description of the Base chain
func (b *Base) Description() string {
	return b.Ethereum.Description()
}

// SupportedProtocols returns the list of protocol IDs supported by this chain
func (b *Base) SupportedProtocols() []string {
	return b.Ethereum.SupportedProtocols()
}

// GetProtocol returns a protocol handler for the given ID
func (b *Base) GetProtocol(id string) (types.Protocol, error) {
	return b.Ethereum.GetProtocol(id)
}

// ParseTransaction delegates to the base EVM implementation
func (b *Base) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	return b.Ethereum.ParseTransaction(txHex)
}

// ComputeTxHash delegates to the base EVM implementation
func (b *Base) ComputeTxHash(proposedTxHex string, sigs []tss.KeysignResponse) (string, error) {
	return b.Ethereum.ComputeTxHash(proposedTxHex, sigs)
}

// Constructor
func NewBase() *Base {
	ethChain := ethereum.NewEthereum()

	ethChain.SetChainID(big.NewInt(8453))
	ethChain.SetID("base")
	ethChain.SetName("Base")
	ethChain.SetDescription("Base is a secure, low-cost, builder-friendly Ethereum L2 built to bring the next billion users onchain.")
	ethChain.SetSupportedABIs(map[string]bool{
		"erc20": true,
	})

	return &Base{
		Ethereum: ethChain,
	}
}
