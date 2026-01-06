// Package bridge provides a canonical bridge router for Vultisig applications.
//
// Applications should use this package to bridge assets across chains without
// worrying about which provider to use - the router automatically selects the
// best available provider based on priority and availability.
//
// # What is Bridging?
//
// Bridging transfers the SAME asset from one chain to another:
//   - USDC on Ethereum -> USDC on Arbitrum
//   - ETH on Ethereum -> ETH on Base
//
// This is different from swapping, which exchanges DIFFERENT assets:
//   - ETH -> BTC (swap)
//   - USDC -> DAI (swap)
//
// # Priority Order
//
// Providers are tried in the following order:
//  1. Native L2 - Official L2 bridges (cheapest for ETH<->L2)
//  2. LiFi - Cross-chain bridge aggregator (best rates, most routes)
//  3. Across - Fast bridging via SpokePool (supports Hyperliquid)
//  4. deBridge - DLN protocol for EVM chains
//
// # Basic Usage
//
// For most use cases, simply use the package-level functions:
//
//	// Get a quote (router selects provider automatically)
//	quote, err := bridge.GetQuote(ctx, bridge.QuoteRequest{
//	    From:        bridge.BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
//	    To:          bridge.BridgeAsset{Chain: "Arbitrum", Symbol: "ETH"},
//	    Amount:      big.NewInt(1e18), // 1 ETH in wei
//	    Sender:      "0x...",
//	    Destination: "0x...",
//	})
//
//	// Build the transaction
//	result, err := bridge.BuildTx(ctx, bridge.BridgeRequest{
//	    Quote:       quote,
//	    Sender:      "0x...",
//	    Destination: "0x...",
//	})
//
// # Supported Routes
//
// LiFi supports bridging between most EVM chains:
//   - Ethereum, Arbitrum, Optimism, Base, Polygon, BSC, Avalanche, etc.
//
// Native L2 bridges support:
//   - Ethereum <-> Arbitrum (via Arbitrum Gateway)
//   - Ethereum <-> Optimism (via OP Standard Bridge)
//   - Ethereum <-> Base (via Base Standard Bridge)
//
// # Security
//
// All bridge addresses are triple-verified against official documentation.
// The router will reject any transaction targeting an unknown address.
package bridge

import (
	"context"
	"sync"
)

var (
	// defaultRouter is the global router instance
	defaultRouter     *Router
	defaultRouterOnce sync.Once
)

// getDefaultRouter returns the global router, creating it if necessary
func getDefaultRouter() *Router {
	defaultRouterOnce.Do(func() {
		defaultRouter = NewDefaultRouter()
	})
	return defaultRouter
}

// GetQuote gets a bridge quote using the default router.
// The router automatically selects the best available provider.
//
// Example:
//
//	quote, err := bridge.GetQuote(ctx, bridge.QuoteRequest{
//	    From:        bridge.BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
//	    To:          bridge.BridgeAsset{Chain: "Arbitrum", Symbol: "ETH"},
//	    Amount:      big.NewInt(1e18),
//	    Sender:      "0x...",
//	    Destination: "0x...",
//	})
func GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	return getDefaultRouter().GetQuote(ctx, req)
}

// BuildTx builds an unsigned transaction using the default router.
//
// Example:
//
//	result, err := bridge.BuildTx(ctx, bridge.BridgeRequest{
//	    Quote:       quote,
//	    Sender:      "0x...",
//	    Destination: "0x...",
//	})
func BuildTx(ctx context.Context, req BridgeRequest) (*BridgeResult, error) {
	return getDefaultRouter().BuildTx(ctx, req)
}

// FindRoute checks if a bridge route is available for the given assets.
func FindRoute(ctx context.Context, from, to BridgeAsset) (*RouteResult, error) {
	return getDefaultRouter().FindRoute(ctx, from, to)
}

// ListProviders returns all configured providers in priority order.
func ListProviders() []string {
	return getDefaultRouter().ListProviders()
}

// GetSupportedChains returns all chains supported by at least one provider.
func GetSupportedChains() []string {
	return getDefaultRouter().GetSupportedChains()
}

// GetProviderStatus returns the availability status for a provider and chain.
func GetProviderStatus(ctx context.Context, providerName, chain string) (*ProviderStatus, error) {
	return getDefaultRouter().GetProviderStatus(ctx, providerName, chain)
}
