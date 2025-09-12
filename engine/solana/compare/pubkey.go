package compare

import (
	"github.com/vultisig/recipes/engine/compare"
)

type Pubkey struct {
	compare.Falsy[string]
	expected string
}

func NewPubkey(raw string) (compare.Compare[string], error) {
	return &Pubkey{
		expected: raw,
	}, nil
}

func (p *Pubkey) Fixed(actual string) bool {
	return p.expected == actual
}

func (p *Pubkey) Magic(actual string) bool {
	return p.expected == actual
}
