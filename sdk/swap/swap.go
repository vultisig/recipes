// Package swap provides a canonical swap router for Vultisig applications.
//
// Applications should use this package to swap assets without worrying about
// which provider to use - the router automatically selects the best available
// provider based on priority and availability.
//
// # Priority Order
//
// Providers are tried in the following order:
//  1. THORChain - Cross-chain DEX with deep liquidity
//  2. Mayachain - Similar to THORChain, alternative liquidity
//  3. LiFi - Cross-chain aggregator
//  4. 1inch - EVM DEX aggregator
//  5. Jupiter - Solana DEX aggregator
//  6. Uniswap - EVM DEX
//
// # Basic Usage
//
// For most use cases, simply use the package-level functions:
//
//	// Get a quote (router selects provider automatically)
//	quote, err := swap.GetQuote(ctx, swap.QuoteRequest{
//	    From:        swap.Asset{Chain: "Ethereum", Symbol: "ETH"},
//	    To:          swap.Asset{Chain: "Bitcoin", Symbol: "BTC"},
//	    Amount:      big.NewInt(1e18), // 1 ETH in wei
//	    Sender:      "0x...",
//	    Destination: "bc1q...",
//	})
//
//	// Build the transaction
//	result, err := swap.BuildTx(ctx, swap.SwapRequest{
//	    Quote:       quote,
//	    Sender:      "0x...",
//	    Destination: "bc1q...",
//	})
//
// # Advanced Usage
//
// For more control, create a custom router:
//
//	router := swap.NewRouter(
//	    swap.WithProvider(swap.NewTHORChainProvider(client)),
//	    swap.WithProvider(swap.NewMayachainProvider(client)),
//	)
package swap

import (
	"context"
	"math/big"
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

// GetQuote gets a swap quote using the default router.
// The router automatically selects the best available provider.
//
// Example:
//
//	quote, err := swap.GetQuote(ctx, swap.QuoteRequest{
//	    From:        swap.Asset{Chain: "Ethereum", Symbol: "ETH"},
//	    To:          swap.Asset{Chain: "Bitcoin", Symbol: "BTC"},
//	    Amount:      big.NewInt(1e18),
//	    Sender:      "0x...",
//	    Destination: "bc1q...",
//	})
func GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	return getDefaultRouter().GetQuote(ctx, req)
}

// BuildTx builds an unsigned swap transaction using the default router.
// The quote must have been obtained from GetQuote.
//
// Example:
//
//	result, err := swap.BuildTx(ctx, swap.SwapRequest{
//	    Quote:       quote,
//	    Sender:      "0x...",
//	    Destination: "bc1q...",
//	})
func BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	return getDefaultRouter().BuildTx(ctx, req)
}

// FindRoute finds a provider that can handle a swap between two assets.
// This is useful to check if a swap is possible before getting a quote.
//
// Example:
//
//	route, err := swap.FindRoute(ctx, ethAsset, btcAsset)
//	if route.IsSupported {
//	    // Swap is possible via route.Provider
//	}
func FindRoute(ctx context.Context, from, to Asset) (*RouteResult, error) {
	return getDefaultRouter().FindRoute(ctx, from, to)
}

// CanSwap is a convenience function that returns true if a swap is possible
// between two assets with any available provider.
//
// Example:
//
//	if swap.CanSwap(ctx, ethAsset, btcAsset) {
//	    // Proceed with swap
//	}
func CanSwap(ctx context.Context, from, to Asset) bool {
	route, err := FindRoute(ctx, from, to)
	return err == nil && route != nil && route.IsSupported
}

// GetProviderStatus returns the availability status of a specific provider
// for a given chain.
//
// Example:
//
//	status, err := swap.GetProviderStatus(ctx, "THORChain", "Ethereum")
//	if status.Halted {
//	    // THORChain is halted for Ethereum
//	}
func GetProviderStatus(ctx context.Context, provider, chain string) (*ProviderStatus, error) {
	return getDefaultRouter().GetProviderStatus(ctx, provider, chain)
}

// ListProviders returns all configured providers in priority order.
//
// Example:
//
//	providers := swap.ListProviders()
//	// ["THORChain", "Mayachain", "LiFi", "1inch", "Jupiter", "Uniswap"]
func ListProviders() []string {
	return getDefaultRouter().ListProviders()
}

// GetSupportedChains returns all chains supported by at least one provider.
//
// Example:
//
//	chains := swap.GetSupportedChains()
//	// ["Arbitrum", "Avalanche", "Base", "Bitcoin", "BSC", ...]
func GetSupportedChains() []string {
	return getDefaultRouter().GetSupportedChains()
}

// QuoteAndBuild is a convenience function that gets a quote and builds
// the transaction in a single call.
//
// Example:
//
//	quote, result, err := swap.QuoteAndBuild(ctx, swap.QuoteRequest{
//	    From:        swap.Asset{Chain: "Ethereum", Symbol: "ETH"},
//	    To:          swap.Asset{Chain: "Bitcoin", Symbol: "BTC"},
//	    Amount:      big.NewInt(1e18),
//	    Sender:      "0x...",
//	    Destination: "bc1q...",
//	})
func QuoteAndBuild(ctx context.Context, req QuoteRequest) (*Quote, *SwapResult, error) {
	quote, err := GetQuote(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	result, err := BuildTx(ctx, SwapRequest{
		Quote:       quote,
		Sender:      req.Sender,
		Destination: req.Destination,
	})
	if err != nil {
		return quote, nil, err
	}

	return quote, result, nil
}

// NewAsset creates a new Asset with the given parameters.
// This is a convenience constructor.
//
// Example:
//
//	eth := swap.NewAsset("Ethereum", "ETH", "", 18)
//	usdc := swap.NewAsset("Ethereum", "USDC", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", 6)
func NewAsset(chain, symbol, address string, decimals int) Asset {
	return Asset{
		Chain:    chain,
		Symbol:   symbol,
		Address:  address,
		Decimals: decimals,
	}
}

// NativeAsset creates a new native asset (no contract address) for a chain.
//
// Example:
//
//	eth := swap.NativeAsset("Ethereum", "ETH", 18)
//	btc := swap.NativeAsset("Bitcoin", "BTC", 8)
func NativeAsset(chain, symbol string, decimals int) Asset {
	return Asset{
		Chain:    chain,
		Symbol:   symbol,
		Address:  "",
		Decimals: decimals,
	}
}

// ToBaseUnits converts a decimal amount to base units (e.g., ETH to wei).
//
// Example:
//
//	// Convert 1.5 ETH to wei
//	amount := swap.ToBaseUnits(1.5, 18) // Returns 1500000000000000000
func ToBaseUnits(amount float64, decimals int) *big.Int {
	// Use big.Float for precision
	f := new(big.Float).SetFloat64(amount)
	multiplier := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	f.Mul(f, multiplier)
	result, _ := f.Int(nil)
	return result
}

// FromBaseUnits converts base units to a decimal amount (e.g., wei to ETH).
//
// Example:
//
//	// Convert wei to ETH
//	amount := swap.FromBaseUnits(big.NewInt(1e18), 18) // Returns 1.0
func FromBaseUnits(amount *big.Int, decimals int) float64 {
	if amount == nil {
		return 0
	}
	f := new(big.Float).SetInt(amount)
	divisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	f.Quo(f, divisor)
	result, _ := f.Float64()
	return result
}

