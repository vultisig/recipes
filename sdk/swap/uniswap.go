package swap

import (
	"context"
	"fmt"
)

// Uniswap supported chains
var uniswapSupportedChains = []string{
	"Ethereum",
	"Arbitrum",
	"Optimism",
	"Polygon",
	"Base",
}

// UniswapProvider implements SwapProvider for Uniswap
type UniswapProvider struct {
	BaseProvider
}

// NewUniswapProvider creates a new Uniswap provider
func NewUniswapProvider() *UniswapProvider {
	return &UniswapProvider{
		BaseProvider: NewBaseProvider("Uniswap", PriorityUniswap, uniswapSupportedChains),
	}
}

// SupportsRoute checks if Uniswap can handle a swap between two assets
// Uniswap only supports same-chain swaps on supported chains
func (p *UniswapProvider) SupportsRoute(from, to Asset) bool {
	return from.Chain == to.Chain && p.SupportsChain(from.Chain)
}

// IsAvailable checks if Uniswap is available for a specific chain
func (p *UniswapProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	return p.SupportsChain(chain), nil
}

// GetStatus returns detailed availability status for a chain
func (p *UniswapProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	return &ProviderStatus{
		Chain:     chain,
		Available: p.SupportsChain(chain),
	}, nil
}

// GetQuote gets a swap quote from Uniswap
func (p *UniswapProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	// TODO: Implement Uniswap quote API
	return nil, fmt.Errorf("Uniswap provider not yet implemented")
}

// BuildTx builds an unsigned transaction for the swap
func (p *UniswapProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	// TODO: Implement Uniswap transaction building
	return nil, fmt.Errorf("Uniswap provider not yet implemented")
}

