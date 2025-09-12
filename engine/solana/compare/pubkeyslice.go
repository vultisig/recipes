package compare

import (
	"strings"

	"github.com/vultisig/recipes/engine/compare"
)

type PubkeySlice struct {
	compare.Falsy[[]string]
	expected []string
}

func NewPubkeySlice(raw string) (compare.Compare[[]string], error) {
	// Parse comma-separated list of pubkeys
	expected := strings.Split(raw, ",")
	for i, pubkey := range expected {
		expected[i] = strings.TrimSpace(pubkey)
	}

	return &PubkeySlice{
		expected: expected,
	}, nil
}

func (p *PubkeySlice) Fixed(actual []string) bool {
	if len(p.expected) != len(actual) {
		return false
	}

	for i, expectedPubkey := range p.expected {
		if expectedPubkey != actual[i] {
			return false
		}
	}

	return true
}

func (p *PubkeySlice) Magic(actual []string) bool {
	return p.Fixed(actual)
}
