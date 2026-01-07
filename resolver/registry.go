package resolver

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

type MagicConstantRegistry struct {
	resolvers []Resolver
}

func NewMagicConstantRegistry() *MagicConstantRegistry {
	registry := &MagicConstantRegistry{
		resolvers: make([]Resolver, 0),
	}

	// Register all resolvers
	registry.Register(NewDefaultTreasuryResolver())
	registry.Register(NewTHORChainVaultResolver())
	registry.Register(NewTHORChainRouterResolver())
	registry.Register(NewMayaChainVaultResolver())
	registry.Register(NewMayaChainRouterResolver())
	// Swap aggregator routers
	registry.Register(NewLiFiRouterResolver())
	registry.Register(NewOneInchRouterResolver())
	registry.Register(NewUniswapRouterResolver())
	// Native L2 bridge resolvers
	registry.Register(NewNativeBridgeResolver())

	return registry
}

func (r *MagicConstantRegistry) Register(resolver Resolver) {
	r.resolvers = append(r.resolvers, resolver)
}

// GetResolver returns the first resolver that supports the given magic constant
func (r *MagicConstantRegistry) GetResolver(constant types.MagicConstant) (Resolver, error) {
	for _, resolver := range r.resolvers {
		if resolver.Supports(constant) {
			return resolver, nil
		}
	}

	return nil, fmt.Errorf("no resolver found for magic constant: %v", constant)
}
