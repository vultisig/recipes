package resolver

import (
	"fmt"
	"strings"

	"github.com/vultisig/recipes/types"
)

// ============================================================================
// LiFi Diamond Contract Addresses
// ============================================================================
// Source: https://docs.li.fi/list-of-all-lifi-contract-addresses
// VERIFIED: These addresses are the same across most EVM chains (Diamond proxy pattern)
// ============================================================================

var lifiRouters = map[string]string{
	"Ethereum":  "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Arbitrum":  "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Optimism":  "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Polygon":   "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"BSC":       "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Avalanche": "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Base":      "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Fantom":    "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Gnosis":    "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Blast":     "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Zksync":    "0x341e94069f53234fE6DabeF707aD424830525715", // zkSync has different address
}

// LiFiRouterResolver resolves LiFi Diamond contract addresses
type LiFiRouterResolver struct{}

// NewLiFiRouterResolver creates a new LiFi router resolver
func NewLiFiRouterResolver() Resolver {
	return &LiFiRouterResolver{}
}

// Supports returns true if this resolver handles LIFI_ROUTER
func (r *LiFiRouterResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_LIFI_ROUTER
}

// Resolve returns the LiFi Diamond address for the given chain
func (r *LiFiRouterResolver) Resolve(constant types.MagicConstant, chainID, _ string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("LiFiRouterResolver does not support type: %v", constant)
	}

	// Normalize chain ID (handle case variations)
	normalizedChain := normalizeChainID(chainID)

	router, ok := lifiRouters[normalizedChain]
	if !ok {
		return "", "", fmt.Errorf("no LiFi router found for chain %s", chainID)
	}

	return router, "", nil
}

// normalizeChainID normalizes chain ID strings for lookup
func normalizeChainID(chainID string) string {
	// Map common variations to canonical names
	normalized := strings.ToLower(chainID)
	switch normalized {
	case "ethereum", "eth":
		return "Ethereum"
	case "arbitrum", "arb", "arbitrumone":
		return "Arbitrum"
	case "optimism", "op":
		return "Optimism"
	case "polygon", "matic":
		return "Polygon"
	case "bsc", "binance", "bnb":
		return "BSC"
	case "avalanche", "avax":
		return "Avalanche"
	case "base":
		return "Base"
	case "fantom", "ftm":
		return "Fantom"
	case "gnosis", "xdai":
		return "Gnosis"
	case "blast":
		return "Blast"
	case "zksync", "zksyncera":
		return "Zksync"
	default:
		return chainID
	}
}

// ResolveLiFiRouter is a convenience function to get LiFi router address
func ResolveLiFiRouter(chainID string) (string, error) {
	resolver := NewLiFiRouterResolver()
	addr, _, err := resolver.Resolve(types.MagicConstant_LIFI_ROUTER, chainID, "")
	return addr, err
}
