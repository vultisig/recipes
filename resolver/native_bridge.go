package resolver

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// ============================================================================
// CRITICAL: Native L2 Bridge Contract Addresses
// ============================================================================
// These addresses are VERIFIED against official documentation and Etherscan.
// DO NOT modify without triple-verification from official sources.
//
// Sources:
// - Arbitrum: https://docs.arbitrum.io/build-decentralized-apps/reference/useful-addresses
// - Optimism: https://docs.optimism.io/chain/addresses
// - Base: https://docs.base.org/docs/base-contracts
// ============================================================================

// L1 Bridge Addresses (on Ethereum Mainnet)
const (
	// Arbitrum L1 Gateway Router
	// Verified: https://etherscan.io/address/0x72Ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef
	ArbitrumL1GatewayRouter = "0x72Ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef"

	// Optimism L1 Standard Bridge (Proxy)
	// Verified: https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1
	OptimismL1StandardBridge = "0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1"

	// Base L1 Standard Bridge (Proxy)
	// Verified: https://etherscan.io/address/0x3154Cf16ccdb4C6d922629664174b904d80F2C35
	BaseL1StandardBridge = "0x3154Cf16ccdb4C6d922629664174b904d80F2C35"
)

// L2 Bridge Addresses (on L2 chains)
const (
	// Arbitrum L2 Gateway Router
	// Verified: https://arbiscan.io/address/0x5288c571Fd7aD117beA99bF60FE0846C4E84F933
	ArbitrumL2GatewayRouter = "0x5288c571Fd7aD117beA99bF60FE0846C4E84F933"

	// Optimism L2 Standard Bridge (Predeploy)
	// Verified: https://optimistic.etherscan.io/address/0x4200000000000000000000000000000000000010
	OptimismL2StandardBridge = "0x4200000000000000000000000000000000000010"

	// Base L2 Standard Bridge (Predeploy - same as Optimism, OP Stack)
	// Verified: https://basescan.org/address/0x4200000000000000000000000000000000000010
	BaseL2StandardBridge = "0x4200000000000000000000000000000000000010"
)

// NativeBridgeResolver resolves native L2 bridge addresses
type NativeBridgeResolver struct{}

// NewNativeBridgeResolver creates a new native bridge resolver
func NewNativeBridgeResolver() Resolver {
	return &NativeBridgeResolver{}
}

// Supports returns true if this resolver handles native bridge magic constants
func (r *NativeBridgeResolver) Supports(constant types.MagicConstant) bool {
	switch constant {
	case types.MagicConstant_ARBITRUM_L1_GATEWAY,
		types.MagicConstant_OPTIMISM_L1_BRIDGE,
		types.MagicConstant_BASE_L1_BRIDGE,
		types.MagicConstant_ARBITRUM_L2_GATEWAY,
		types.MagicConstant_OPTIMISM_L2_BRIDGE,
		types.MagicConstant_BASE_L2_BRIDGE:
		return true
	default:
		return false
	}
}

// Resolve returns the bridge address for the given magic constant
func (r *NativeBridgeResolver) Resolve(constant types.MagicConstant, _, _ string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("NativeBridgeResolver does not support type: %v", constant)
	}

	switch constant {
	// L1 bridges (on Ethereum)
	case types.MagicConstant_ARBITRUM_L1_GATEWAY:
		return ArbitrumL1GatewayRouter, "", nil
	case types.MagicConstant_OPTIMISM_L1_BRIDGE:
		return OptimismL1StandardBridge, "", nil
	case types.MagicConstant_BASE_L1_BRIDGE:
		return BaseL1StandardBridge, "", nil

	// L2 bridges (on respective L2 chains)
	case types.MagicConstant_ARBITRUM_L2_GATEWAY:
		return ArbitrumL2GatewayRouter, "", nil
	case types.MagicConstant_OPTIMISM_L2_BRIDGE:
		return OptimismL2StandardBridge, "", nil
	case types.MagicConstant_BASE_L2_BRIDGE:
		return BaseL2StandardBridge, "", nil

	default:
		return "", "", fmt.Errorf("unknown bridge constant: %v", constant)
	}
}

// Convenience functions for direct access

// ResolveArbitrumL1Gateway returns the Arbitrum L1 Gateway Router address
func ResolveArbitrumL1Gateway() string {
	return ArbitrumL1GatewayRouter
}

// ResolveArbitrumL2Gateway returns the Arbitrum L2 Gateway Router address
func ResolveArbitrumL2Gateway() string {
	return ArbitrumL2GatewayRouter
}

// ResolveOptimismL1Bridge returns the Optimism L1 Standard Bridge address
func ResolveOptimismL1Bridge() string {
	return OptimismL1StandardBridge
}

// ResolveOptimismL2Bridge returns the Optimism L2 Standard Bridge address
func ResolveOptimismL2Bridge() string {
	return OptimismL2StandardBridge
}

// ResolveBaseL1Bridge returns the Base L1 Standard Bridge address
func ResolveBaseL1Bridge() string {
	return BaseL1StandardBridge
}

// ResolveBaseL2Bridge returns the Base L2 Standard Bridge address
func ResolveBaseL2Bridge() string {
	return BaseL2StandardBridge
}
