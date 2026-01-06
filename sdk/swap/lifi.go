package swap

import (
	"context"
	"fmt"
)

// LiFi supported EVM chains
var lifiSupportedChains = []string{
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

// LiFiProvider implements SwapProvider for LiFi aggregator
type LiFiProvider struct {
	BaseProvider
	apiKey string
}

// NewLiFiProvider creates a new LiFi provider
func NewLiFiProvider(apiKey string) *LiFiProvider {
	return &LiFiProvider{
		BaseProvider: NewBaseProvider("LiFi", PriorityLiFi, lifiSupportedChains),
		apiKey:       apiKey,
	}
}

// SupportsRoute checks if LiFi can handle a swap between two assets
func (p *LiFiProvider) SupportsRoute(from, to Asset) bool {
	return p.SupportsChain(from.Chain) && p.SupportsChain(to.Chain)
}

// IsAvailable checks if LiFi is available for a specific chain
// LiFi is generally always available if the chain is supported
func (p *LiFiProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	return p.SupportsChain(chain), nil
}

// GetStatus returns detailed availability status for a chain
func (p *LiFiProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	return &ProviderStatus{
		Chain:     chain,
		Available: p.SupportsChain(chain),
	}, nil
}

// GetQuote gets a swap quote from LiFi
func (p *LiFiProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	// TODO: Implement LiFi quote API
	return nil, fmt.Errorf("LiFi provider not yet implemented")
}

// BuildTx builds an unsigned transaction for the swap
func (p *LiFiProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	// TODO: Implement LiFi transaction building
	return nil, fmt.Errorf("LiFi provider not yet implemented")
}

