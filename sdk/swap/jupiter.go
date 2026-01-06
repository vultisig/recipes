package swap

import (
	"context"
	"fmt"
)

// Jupiter only supports Solana
var jupiterSupportedChains = []string{
	"Solana",
}

// JupiterProvider implements SwapProvider for Jupiter aggregator
type JupiterProvider struct {
	BaseProvider
	apiURL string
}

// NewJupiterProvider creates a new Jupiter provider
func NewJupiterProvider(apiURL string) *JupiterProvider {
	if apiURL == "" {
		apiURL = "https://api.jup.ag"
	}
	return &JupiterProvider{
		BaseProvider: NewBaseProvider("Jupiter", PriorityJupiter, jupiterSupportedChains),
		apiURL:       apiURL,
	}
}

// SupportsRoute checks if Jupiter can handle a swap between two assets
// Jupiter only supports Solana-to-Solana swaps
func (p *JupiterProvider) SupportsRoute(from, to Asset) bool {
	return from.Chain == "Solana" && to.Chain == "Solana"
}

// IsAvailable checks if Jupiter is available for a specific chain
func (p *JupiterProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	return chain == "Solana", nil
}

// GetStatus returns detailed availability status for a chain
func (p *JupiterProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	return &ProviderStatus{
		Chain:     chain,
		Available: chain == "Solana",
	}, nil
}

// GetQuote gets a swap quote from Jupiter
func (p *JupiterProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	// TODO: Implement Jupiter quote API
	return nil, fmt.Errorf("Jupiter provider not yet implemented")
}

// BuildTx builds an unsigned transaction for the swap
func (p *JupiterProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	// TODO: Implement Jupiter transaction building
	return nil, fmt.Errorf("Jupiter provider not yet implemented")
}

