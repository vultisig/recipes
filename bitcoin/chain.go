package bitcoin

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// Bitcoin implements the Chain interface for the Bitcoin blockchain
type Bitcoin struct{}

// ID returns the unique identifier for the Bitcoin chain
func (b *Bitcoin) ID() string {
	return "bitcoin"
}

// Name returns a human-readable name for the Bitcoin chain
func (b *Bitcoin) Name() string {
	return "Bitcoin"
}

// SupportedProtocols returns the list of protocol IDs supported by the Bitcoin chain
func (b *Bitcoin) SupportedProtocols() []string {
	return []string{"btc"}
}

// Description returns a human-readable description for the Bitcoin chain
func (b *Bitcoin) Description() string {
	return "Bitcoin is a digital currency that is not controlled by any government or financial institution."
}

func (b *Bitcoin) GetProtocol(id string) (types.Protocol, error) {
	return nil, fmt.Errorf("protocol %q not found or not supported on Bitcoin", id)
}

func (b *Bitcoin) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	return nil, fmt.Errorf("transaction parsing not supported on Bitcoin")
}

// NewBitcoin creates a new Bitcoin chain instance
func NewChain() types.Chain {
	return &Bitcoin{}
}
