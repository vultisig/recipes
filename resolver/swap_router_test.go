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

	// Test known chains - each chain may have its own router address
	testCases := map[string]string{
		"Ethereum":  "0x111111125421cA6dc452d289314280a0f8842A65",
		"BSC":       "0x111111125421cA6dc452d289314280a0f8842A65",
		"Polygon":   "0x111111125421cA6dc452d289314280a0f8842A65",
		"Avalanche": "0x652747cb44D5fC52799c3DaEa613c52625588AB5",
		"Arbitrum":  "0x6b0CE50D408d27ABA09F7e96Ac437011D8CDFbB8",
		"Optimism":  "0x111111125421cA6dc452d289314280a0f8842A65",
		"Base":      "0x111111125421cA6dc452d289314280a0f8842A65",
		"Gnosis":    "0xed6c1002450cbf418e96d16361cbed3a84366c43",
	}

	for chain, expectedAddress := range testCases {
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

	// Fantom removed from V6 - should error
	_, _, err = resolver.Resolve(types.MagicConstant_ONEINCH_ROUTER, "Fantom", "")
	if err == nil {
		t.Error("1inch V6 should not be available for Fantom (removed)")
	}
}

func TestUniswapRouterResolver(t *testing.T) {
	resolver := NewUniswapRouterResolver()

	if !resolver.Supports(types.MagicConstant_UNISWAP_UNIVERSAL_ROUTER) {
		t.Error("UniswapRouterResolver should support UNISWAP_UNIVERSAL_ROUTER")
	}

	// Test known chains - each chain has its own Universal Router address
	testCases := map[string]string{
		"Ethereum": "0x66a9893cc07d91d95644aedd05d03f95e1dba8af",
		"Polygon":  "0x1095692a6237d83c6a72f3f5efedb9a670c49223",
		"Arbitrum": "0x5E325eDA8064b456f4781070C0738d849c824258",
		"Optimism": "0x851116d9223fabed8e56c0e6b8ad0c31d98b3507",
		"Base":     "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD",
		"BSC":      "0x1906c1d672b88cd1b9ac7593301ca990f94eae07",
	}

	for chain, expectedAddress := range testCases {
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

