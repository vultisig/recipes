package swap

import (
	"context"
	"fmt"
)

// 1inch supported EVM chains (same-chain swaps only)
var oneInchSupportedChains = []string{
	"Ethereum",
	"BSC",
	"Polygon",
	"Avalanche",
	"Arbitrum",
	"Optimism",
	"Base",
	"Fantom",
	"Gnosis",
}

// OneInchProvider implements SwapProvider for 1inch aggregator
type OneInchProvider struct {
	BaseProvider
	apiKey string
}

// NewOneInchProvider creates a new 1inch provider
func NewOneInchProvider(apiKey string) *OneInchProvider {
	return &OneInchProvider{
		BaseProvider: NewBaseProvider("1inch", PriorityOneInch, oneInchSupportedChains),
		apiKey:       apiKey,
	}
}

// SupportsRoute checks if 1inch can handle a swap between two assets
// 1inch only supports same-chain swaps
func (p *OneInchProvider) SupportsRoute(from, to Asset) bool {
	return from.Chain == to.Chain && p.SupportsChain(from.Chain)
}

// IsAvailable checks if 1inch is available for a specific chain
func (p *OneInchProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	return p.SupportsChain(chain), nil
}

// GetStatus returns detailed availability status for a chain
func (p *OneInchProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	return &ProviderStatus{
		Chain:     chain,
		Available: p.SupportsChain(chain),
	}, nil
}

// GetQuote gets a swap quote from 1inch
func (p *OneInchProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	// TODO: Implement 1inch quote API
	return nil, fmt.Errorf("1inch provider not yet implemented")
}

// BuildTx builds an unsigned transaction for the swap
func (p *OneInchProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	// TODO: Implement 1inch transaction building
	return nil, fmt.Errorf("1inch provider not yet implemented")
}

