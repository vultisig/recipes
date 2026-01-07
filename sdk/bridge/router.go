package bridge

import (
	"context"
	"errors"
	"fmt"
	"sort"
)

var (
	// ErrNoRouteAvailable is returned when no provider can handle the bridge route
	ErrNoRouteAvailable = errors.New("no bridge route available")

	// ErrNoProvidersConfigured is returned when the router has no providers
	ErrNoProvidersConfigured = errors.New("no bridge providers configured")

	// ErrProviderUnavailable is returned when the selected provider is unavailable
	ErrProviderUnavailable = errors.New("provider is unavailable for this chain")

	// ErrAssetMismatch is returned when trying to bridge different assets
	ErrAssetMismatch = errors.New("bridge requires same asset on both chains")
)

// Router orchestrates bridge provider selection and execution
type Router struct {
	providers []BridgeProvider
}

// RouterOption is a function that configures the router
type RouterOption func(*Router)

// WithProvider adds a provider to the router
func WithProvider(p BridgeProvider) RouterOption {
	return func(r *Router) {
		r.providers = append(r.providers, p)
	}
}

// NewRouter creates a new bridge router with the given options
func NewRouter(opts ...RouterOption) *Router {
	r := &Router{
		providers: make([]BridgeProvider, 0),
	}

	for _, opt := range opts {
		opt(r)
	}

	// Sort providers by priority (lower number = higher priority)
	sort.Slice(r.providers, func(i, j int) bool {
		return r.providers[i].Priority() < r.providers[j].Priority()
	})

	return r
}

// NewDefaultRouter creates a router with all default providers
// Providers are sorted by priority: Native L2 (1), LiFi (2), Across (3), deBridge (4)
func NewDefaultRouter() *Router {
	return NewRouter(
		WithProvider(NewNativeL2Provider()),
		WithProvider(NewLiFiProvider("")),
		WithProvider(NewAcrossProvider()),
		WithProvider(NewDeBridgeProvider()),
	)
}

// FindRoute finds the first available provider that can handle the bridge route
func (r *Router) FindRoute(ctx context.Context, from, to BridgeAsset) (*RouteResult, error) {
	if len(r.providers) == 0 {
		return nil, ErrNoProvidersConfigured
	}

	var lastErr error

	for _, provider := range r.providers {
		// Check if provider supports this route
		if !provider.SupportsRoute(from, to) {
			continue
		}

		// Check if provider is available for source chain
		available, err := provider.IsAvailable(ctx, from.Chain)
		if err != nil {
			lastErr = fmt.Errorf("provider %s availability check failed: %w", provider.Name(), err)
			continue
		}
		if !available {
			lastErr = fmt.Errorf("provider %s is not available for chain %s", provider.Name(), from.Chain)
			continue
		}

		// Check destination chain availability
		available, err = provider.IsAvailable(ctx, to.Chain)
		if err != nil {
			lastErr = fmt.Errorf("provider %s availability check failed for destination: %w", provider.Name(), err)
			continue
		}
		if !available {
			lastErr = fmt.Errorf("provider %s is not available for destination chain %s", provider.Name(), to.Chain)
			continue
		}

		return &RouteResult{
			Provider:    provider.Name(),
			IsSupported: true,
		}, nil
	}

	if lastErr != nil {
		return nil, fmt.Errorf("%w: %v", ErrNoRouteAvailable, lastErr)
	}

	return nil, ErrNoRouteAvailable
}

// GetQuote gets a bridge quote from the first available provider
func (r *Router) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	if len(r.providers) == 0 {
		return nil, ErrNoProvidersConfigured
	}

	var lastErr error

	for _, provider := range r.providers {
		// Check if provider supports this route
		if !provider.SupportsRoute(req.From, req.To) {
			continue
		}

		// Check if provider is available for source chain
		available, err := provider.IsAvailable(ctx, req.From.Chain)
		if err != nil || !available {
			if err != nil {
				lastErr = err
			}
			continue
		}

		// Check destination chain availability
		available, err = provider.IsAvailable(ctx, req.To.Chain)
		if err != nil || !available {
			if err != nil {
				lastErr = err
			}
			continue
		}

		// Get quote from this provider
		quote, err := provider.GetQuote(ctx, req)
		if err != nil {
			lastErr = fmt.Errorf("provider %s quote failed: %w", provider.Name(), err)
			continue
		}

		return quote, nil
	}

	if lastErr != nil {
		return nil, fmt.Errorf("%w: %v", ErrNoRouteAvailable, lastErr)
	}

	return nil, ErrNoRouteAvailable
}

// BuildTx builds an unsigned transaction using the quote's provider
func (r *Router) BuildTx(ctx context.Context, req BridgeRequest) (*BridgeResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	// Find the provider that issued the quote
	var provider BridgeProvider
	for _, p := range r.providers {
		if p.Name() == req.Quote.Provider {
			provider = p
			break
		}
	}

	if provider == nil {
		return nil, fmt.Errorf("provider %s not found", req.Quote.Provider)
	}

	return provider.BuildTx(ctx, req)
}

// GetProviderStatus returns the availability status for a provider and chain
func (r *Router) GetProviderStatus(ctx context.Context, providerName, chain string) (*ProviderStatus, error) {
	for _, p := range r.providers {
		if p.Name() == providerName {
			return p.GetStatus(ctx, chain)
		}
	}
	return nil, fmt.Errorf("provider %s not found", providerName)
}

// ListProviders returns all configured providers in priority order
func (r *Router) ListProviders() []string {
	names := make([]string, len(r.providers))
	for i, p := range r.providers {
		names[i] = p.Name()
	}
	return names
}

// GetSupportedChains returns all chains supported by at least one provider
func (r *Router) GetSupportedChains() []string {
	chainSet := make(map[string]struct{})
	for _, p := range r.providers {
		for _, chain := range p.SupportedChains() {
			chainSet[chain] = struct{}{}
		}
	}

	chains := make([]string, 0, len(chainSet))
	for chain := range chainSet {
		chains = append(chains, chain)
	}
	sort.Strings(chains)
	return chains
}

// RouteResult contains the result of finding a bridge route
type RouteResult struct {
	Provider    string // Provider name that can handle the route
	IsSupported bool   // Whether the route is supported
	Quote       *Quote // Optional quote if requested
}
