package resolver

import (
	"strings"
	"testing"

	"github.com/vultisig/recipes/types"
)

// TestNativeBridgeAddressConstants verifies all bridge addresses are correctly formatted
func TestNativeBridgeAddressConstants(t *testing.T) {
	validateAddress := func(t *testing.T, name, addr string) {
		t.Helper()
		if !strings.HasPrefix(addr, "0x") {
			t.Errorf("%s: address should start with 0x: %s", name, addr)
		}
		if len(addr) != 42 {
			t.Errorf("%s: address should be 42 chars: %s has %d", name, addr, len(addr))
		}
	}

	// Verify L1 addresses
	validateAddress(t, "ArbitrumL1GatewayRouter", ArbitrumL1GatewayRouter)
	validateAddress(t, "OptimismL1StandardBridge", OptimismL1StandardBridge)
	validateAddress(t, "BaseL1StandardBridge", BaseL1StandardBridge)

	// Verify L2 addresses
	validateAddress(t, "ArbitrumL2GatewayRouter", ArbitrumL2GatewayRouter)
	validateAddress(t, "OptimismL2StandardBridge", OptimismL2StandardBridge)
	validateAddress(t, "BaseL2StandardBridge", BaseL2StandardBridge)
}

// TestNativeBridgeResolverSupports verifies the resolver supports all bridge constants
func TestNativeBridgeResolverSupports(t *testing.T) {
	resolver := NewNativeBridgeResolver()

	supportedConstants := []types.MagicConstant{
		types.MagicConstant_ARBITRUM_L1_GATEWAY,
		types.MagicConstant_OPTIMISM_L1_BRIDGE,
		types.MagicConstant_BASE_L1_BRIDGE,
		types.MagicConstant_ARBITRUM_L2_GATEWAY,
		types.MagicConstant_OPTIMISM_L2_BRIDGE,
		types.MagicConstant_BASE_L2_BRIDGE,
	}

	for _, constant := range supportedConstants {
		if !resolver.Supports(constant) {
			t.Errorf("Resolver should support %v", constant)
		}
	}

	// Verify it doesn't support unrelated constants
	unsupportedConstants := []types.MagicConstant{
		types.MagicConstant_UNSPECIFIED,
		types.MagicConstant_VULTISIG_TREASURY,
		types.MagicConstant_THORCHAIN_VAULT,
	}

	for _, constant := range unsupportedConstants {
		if resolver.Supports(constant) {
			t.Errorf("Resolver should not support %v", constant)
		}
	}
}

// TestNativeBridgeResolverResolve verifies addresses are correctly resolved
func TestNativeBridgeResolverResolve(t *testing.T) {
	resolver := NewNativeBridgeResolver()

	testCases := []struct {
		constant types.MagicConstant
		expected string
	}{
		{types.MagicConstant_ARBITRUM_L1_GATEWAY, ArbitrumL1GatewayRouter},
		{types.MagicConstant_OPTIMISM_L1_BRIDGE, OptimismL1StandardBridge},
		{types.MagicConstant_BASE_L1_BRIDGE, BaseL1StandardBridge},
		{types.MagicConstant_ARBITRUM_L2_GATEWAY, ArbitrumL2GatewayRouter},
		{types.MagicConstant_OPTIMISM_L2_BRIDGE, OptimismL2StandardBridge},
		{types.MagicConstant_BASE_L2_BRIDGE, BaseL2StandardBridge},
	}

	for _, tc := range testCases {
		addr, memo, err := resolver.Resolve(tc.constant, "", "")
		if err != nil {
			t.Errorf("Resolve(%v) failed: %v", tc.constant, err)
			continue
		}
		if memo != "" {
			t.Errorf("Resolve(%v) returned non-empty memo: %s", tc.constant, memo)
		}
		if addr != tc.expected {
			t.Errorf("Resolve(%v) = %s, want %s", tc.constant, addr, tc.expected)
		}
	}
}

// TestNativeBridgeAddressValues verifies the specific addresses match official docs
func TestNativeBridgeAddressValues(t *testing.T) {
	// These are the official addresses from documentation
	// DO NOT MODIFY without verifying against official sources
	testCases := []struct {
		name     string
		actual   string
		expected string
	}{
		// L1 Addresses (on Ethereum)
		{
			name:     "Arbitrum L1 Gateway",
			actual:   ArbitrumL1GatewayRouter,
			expected: "0x72Ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef",
		},
		{
			name:     "Optimism L1 Bridge",
			actual:   OptimismL1StandardBridge,
			expected: "0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1",
		},
		{
			name:     "Base L1 Bridge",
			actual:   BaseL1StandardBridge,
			expected: "0x3154Cf16ccdb4C6d922629664174b904d80F2C35",
		},
		// L2 Addresses
		{
			name:     "Arbitrum L2 Gateway",
			actual:   ArbitrumL2GatewayRouter,
			expected: "0x5288c571Fd7aD117beA99bF60FE0846C4E84F933",
		},
		{
			name:     "Optimism L2 Bridge (OP Stack predeploy)",
			actual:   OptimismL2StandardBridge,
			expected: "0x4200000000000000000000000000000000000010",
		},
		{
			name:     "Base L2 Bridge (OP Stack predeploy)",
			actual:   BaseL2StandardBridge,
			expected: "0x4200000000000000000000000000000000000010",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !strings.EqualFold(tc.actual, tc.expected) {
				t.Errorf("Address mismatch for %s:\n  got:  %s\n  want: %s",
					tc.name, tc.actual, tc.expected)
			}
		})
	}
}

// TestLiFiRouterResolverSupports verifies LiFi resolver supports LIFI_ROUTER
func TestLiFiRouterResolverSupports(t *testing.T) {
	resolver := NewLiFiRouterResolver()

	if !resolver.Supports(types.MagicConstant_LIFI_ROUTER) {
		t.Error("LiFi resolver should support LIFI_ROUTER")
	}

	if resolver.Supports(types.MagicConstant_THORCHAIN_VAULT) {
		t.Error("LiFi resolver should not support THORCHAIN_VAULT")
	}
}

// TestLiFiRouterResolverResolve verifies LiFi addresses are correctly resolved
func TestLiFiRouterResolverResolve(t *testing.T) {
	resolver := NewLiFiRouterResolver()

	chains := []string{"Ethereum", "Arbitrum", "Optimism", "Base", "Polygon", "BSC", "Avalanche"}
	expectedMainnet := "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE"

	for _, chain := range chains {
		addr, memo, err := resolver.Resolve(types.MagicConstant_LIFI_ROUTER, chain, "")
		if err != nil {
			t.Errorf("Resolve(LIFI_ROUTER, %s) failed: %v", chain, err)
			continue
		}
		if memo != "" {
			t.Errorf("Resolve(LIFI_ROUTER, %s) returned non-empty memo: %s", chain, memo)
		}
		if !strings.EqualFold(addr, expectedMainnet) {
			t.Errorf("Resolve(LIFI_ROUTER, %s) = %s, want %s", chain, addr, expectedMainnet)
		}
	}

	// zkSync has a different address
	addr, _, err := resolver.Resolve(types.MagicConstant_LIFI_ROUTER, "Zksync", "")
	if err != nil {
		t.Errorf("Resolve(LIFI_ROUTER, Zksync) failed: %v", err)
	}
	if strings.EqualFold(addr, expectedMainnet) {
		t.Error("zkSync LiFi address should be different from mainnet")
	}
}

// TestRegistryContainsBridgeResolvers verifies the registry includes bridge resolvers
func TestRegistryContainsBridgeResolvers(t *testing.T) {
	registry := NewMagicConstantRegistry()

	// Test LiFi resolver is registered
	lifiResolver, err := registry.GetResolver(types.MagicConstant_LIFI_ROUTER)
	if err != nil {
		t.Errorf("Registry missing LiFi resolver: %v", err)
	}
	if lifiResolver == nil {
		t.Error("LiFi resolver is nil")
	}

	// Test native bridge resolver is registered
	for _, constant := range []types.MagicConstant{
		types.MagicConstant_ARBITRUM_L1_GATEWAY,
		types.MagicConstant_OPTIMISM_L1_BRIDGE,
		types.MagicConstant_BASE_L1_BRIDGE,
	} {
		resolver, err := registry.GetResolver(constant)
		if err != nil {
			t.Errorf("Registry missing resolver for %v: %v", constant, err)
		}
		if resolver == nil {
			t.Errorf("Resolver for %v is nil", constant)
		}
	}
}
