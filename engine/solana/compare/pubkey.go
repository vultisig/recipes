package compare

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/vultisig/recipes/engine/compare"
)

type PubKey struct {
	compare.Falsy[solana.PublicKey]
	expected solana.PublicKey
}

func NewPubKey(raw string) (compare.Compare[solana.PublicKey], error) {
	pk, err := solana.PublicKeyFromBase58(raw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pubkey: %w", err)
	}

	return &PubKey{
		expected: pk,
	}, nil
}

func (p *PubKey) Fixed(actual solana.PublicKey) bool {
	return p.expected.Equals(actual)
}
