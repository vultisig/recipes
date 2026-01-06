package resolver

import (
	"testing"

	"github.com/vultisig/recipes/types"
)

func TestLiFiRouterResolver(t *testing.T) {
	resolver := NewLiFiRouterResolver()

	if !resolver.Supports(types.MagicConstant_LIFI_ROUTER) {
		t.Error("LiFiRouterResolver should support LIFI_ROUTER")
	}

	if resolver.Supports(types.MagicConstant_THORCHAIN_ROUTER) {
		t.Error("LiFiRouterResolver should not support THORCHAIN_ROUTER")
	}

	// Test known chains
	chains := []string{"Ethereum", "Arbitrum", "Optimism", "Polygon", "BSC", "Avalanche", "Base"}
	expectedAddress := "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE"

	for _, chain := range chains {
		addr, _, err := resolver.Resolve(types.MagicConstant_LIFI_ROUTER, chain, "")
		if err != nil {
			t.Errorf("LiFi router should be available for %s: %v", chain, err)
			continue
		}
		if addr != expectedAddress {
			t.Errorf("LiFi router for %s = %s, want %s", chain, addr, expectedAddress)
		}
	}
}

func TestOneInchRouterResolver(t *testing.T) {
	resolver := NewOneInchRouterResolver()

	if !resolver.Supports(types.MagicConstant_ONEINCH_ROUTER) {
		t.Error("OneInchRouterResolver should support ONEINCH_ROUTER")
	}

	// Test known chains
	chains := []string{"Ethereum", "BSC", "Polygon", "Avalanche", "Arbitrum", "Optimism", "Base"}
	expectedAddress := "0x111111125421cA6dc452d289314280a0f8842A65"

	for _, chain := range chains {
		addr, _, err := resolver.Resolve(types.MagicConstant_ONEINCH_ROUTER, chain, "")
		if err != nil {
			t.Errorf("1inch router should be available for %s: %v", chain, err)
			continue
		}
		if addr != expectedAddress {
			t.Errorf("1inch router for %s = %s, want %s", chain, addr, expectedAddress)
		}
	}

	// Test unsupported chain
	_, _, err := resolver.Resolve(types.MagicConstant_ONEINCH_ROUTER, "Solana", "")
	if err == nil {
		t.Error("1inch should not be available for Solana")
	}
}

func TestUniswapRouterResolver(t *testing.T) {
	resolver := NewUniswapRouterResolver()

	if !resolver.Supports(types.MagicConstant_UNISWAP_UNIVERSAL_ROUTER) {
		t.Error("UniswapRouterResolver should support UNISWAP_UNIVERSAL_ROUTER")
	}

	// Test known chains
	chains := []string{"Ethereum", "Polygon", "Arbitrum", "Optimism", "Base"}
	expectedAddress := "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"

	for _, chain := range chains {
		addr, _, err := resolver.Resolve(types.MagicConstant_UNISWAP_UNIVERSAL_ROUTER, chain, "")
		if err != nil {
			t.Errorf("Uniswap router should be available for %s: %v", chain, err)
			continue
		}
		if addr != expectedAddress {
			t.Errorf("Uniswap router for %s = %s, want %s", chain, addr, expectedAddress)
		}
	}

	// Test unsupported chain
	_, _, err := resolver.Resolve(types.MagicConstant_UNISWAP_UNIVERSAL_ROUTER, "Solana", "")
	if err == nil {
		t.Error("Uniswap should not be available for Solana")
	}
}

func TestRegistryContainsSwapResolvers(t *testing.T) {
	registry := NewMagicConstantRegistry()

	// Test that all swap resolvers are registered
	constants := []types.MagicConstant{
		types.MagicConstant_LIFI_ROUTER,
		types.MagicConstant_ONEINCH_ROUTER,
		types.MagicConstant_UNISWAP_UNIVERSAL_ROUTER,
	}

	for _, constant := range constants {
		r, err := registry.GetResolver(constant)
		if err != nil {
			t.Errorf("Registry should have resolver for %v: %v", constant, err)
			continue
		}
		if r == nil {
			t.Errorf("Resolver for %v should not be nil", constant)
		}
	}
}

