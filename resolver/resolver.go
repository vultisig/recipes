package resolver

import (
	"github.com/vultisig/recipes/types"
)

// Resolver defines the interface for magic constant resolution
type Resolver interface {
	// Supports returns true if this resolver can handle the given magic constant
	Supports(constant types.MagicConstant) bool

	// Resolve converts a magic constant to an actual address + memo
	Resolve(constant types.MagicConstant, chainID, assetID string) (string, string, error)
}
