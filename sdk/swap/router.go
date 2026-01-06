package swap

import (
	"context"
	"errors"
	"fmt"
	"sort"
)

var (
	// ErrNoRouteAvailable is returned when no provider can handle the swap route
	ErrNoRouteAvailable = errors.New("no swap route available")

	// ErrNoProvidersConfigured is returned when the router has no providers
	ErrNoProvidersConfigured = errors.New("no swap providers configured")

	// ErrProviderUnavailable is returned when the selected provider is unavailable
	ErrProviderUnavailable = errors.New("provider is unavailable for this chain")
)

// Router orchestrates swap provider selection and execution
type Router struct {
	providers []SwapProvider
}

// RouterOption is a function that configures the router
type RouterOption func(*Router)

// WithProvider adds a provider to the router
func WithProvider(p SwapProvider) RouterOption {
	return func(r *Router) {
		r.providers = append(r.providers, p)
	}
}

// NewRouter creates a new swap router with the given options
func NewRouter(opts ...RouterOption) *Router {
	r := &Router{
		providers: make([]SwapProvider, 0),
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
func NewDefaultRouter() *Router {
	return NewRouter(
		WithProvider(NewTHORChainProvider(nil)),
		WithProvider(NewMayachainProvider(nil)),
		WithProvider(NewLiFiProvider("")),
		WithProvider(NewOneInchProvider("")),
		WithProvider(NewJupiterProvider("")),
		WithProvider(NewUniswapProvider()),
	)
}

// FindRoute finds the first available provider that can handle the swap route
func (r *Router) FindRoute(ctx context.Context, from, to Asset) (*RouteResult, error) {
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

		// For cross-chain swaps, check destination chain availability
		if from.Chain != to.Chain {
			available, err = provider.IsAvailable(ctx, to.Chain)
			if err != nil {
				lastErr = fmt.Errorf("provider %s availability check failed for destination: %w", provider.Name(), err)
				continue
			}
			if !available {
				lastErr = fmt.Errorf("provider %s is not available for destination chain %s", provider.Name(), to.Chain)
				continue
			}
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

// GetQuote gets a swap quote from the first available provider
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

		// For cross-chain swaps, check destination chain availability
		if req.From.Chain != req.To.Chain {
			available, err = provider.IsAvailable(ctx, req.To.Chain)
			if err != nil || !available {
				if err != nil {
					lastErr = err
				}
				continue
			}
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
func (r *Router) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	// Find the provider that issued the quote
	var provider SwapProvider
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

// BuildSwapBundle builds a complete swap bundle including approval tx (if needed) and swap tx.
// For ERC20 tokens on EVM chains, this always includes an approval transaction for the exact
// swap amount, bundled with the swap transaction. Both transactions are built with sequential
// nonces so they can be signed in a single keysign session.
//
// Security: This method enforces mandatory approval bundling - an approval tx can only exist
// as part of a swap bundle, preventing orphaned approvals that could waste gas.
func (r *Router) BuildSwapBundle(ctx context.Context, req SwapBundleRequest) (*SwapBundle, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	// Find the provider that issued the quote
	var provider SwapProvider
	for _, p := range r.providers {
		if p.Name() == req.Quote.Provider {
			provider = p
			break
		}
	}

	if provider == nil {
		return nil, fmt.Errorf("provider %s not found", req.Quote.Provider)
	}

	// Build the swap transaction first to get the transaction data
	swapResult, err := provider.BuildTx(ctx, SwapRequest{
		Quote:       req.Quote,
		Sender:      req.Sender,
		Destination: req.Destination,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to build swap tx: %w", err)
	}

	// Get chain ID for EVM chains
	chainID, err := GetEVMChainID(req.Quote.FromAsset.Chain)
	if err != nil {
		// Non-EVM chain - return just the swap tx without approval, include memo
		return &SwapBundle{
			ApprovalTx: nil,
			SwapTx: &TxData{
				To:    swapResult.ToAddress,
				Value: swapResult.Value,
				Data:  swapResult.TxData,
				Memo:  swapResult.Memo,
			},
			Provider:       req.Quote.Provider,
			ExpectedOutput: req.Quote.ExpectedOutput,
			Quote:          req.Quote,
		}, nil
	}

	// Get nonce - either from request or fetch from chain
	var nonce uint64
	if req.Nonce != nil {
		nonce = *req.Nonce
	} else {
		nonce, err = GetNonce(ctx, req.Quote.FromAsset.Chain, req.Sender)
		if err != nil {
			return nil, fmt.Errorf("failed to get nonce: %w", err)
		}
	}

	// Check if approval is needed
	needsApproval := req.Quote.NeedsApproval || IsApprovalRequired(req.Quote.FromAsset)

	var approvalTx *TxData
	swapNonce := nonce

	if needsApproval {
		// Determine the spender (router) address
		spender := req.Quote.ApprovalSpender
		if spender == "" {
			spender = req.Quote.Router
		}
		if spender == "" {
			return nil, fmt.Errorf("no router/spender address available for approval")
		}

		// Determine the approval amount
		approvalAmount := req.Quote.ApprovalAmount
		if approvalAmount == nil {
			approvalAmount = req.Quote.FromAmount
		}
		if approvalAmount == nil {
			return nil, fmt.Errorf("no approval amount available")
		}

		// Build approval transaction
		approvalTx, err = BuildApprovalTx(BuildApprovalInput{
			TokenAddress: req.Quote.FromAsset.Address,
			Spender:      spender,
			Amount:       approvalAmount,
			Nonce:        nonce,
			GasLimit:     DefaultApprovalGasLimit,
			ChainID:      chainID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to build approval tx: %w", err)
		}

		// Swap tx uses next nonce
		swapNonce = nonce + 1
	}

	// Build the swap TxData with correct nonce
	swapTx := &TxData{
		To:       swapResult.ToAddress,
		Value:    swapResult.Value,
		Data:     swapResult.TxData,
		Nonce:    swapNonce,
		GasLimit: 300000, // Default swap gas limit, should be estimated
		ChainID:  chainID,
	}

	return &SwapBundle{
		ApprovalTx:     approvalTx,
		SwapTx:         swapTx,
		Provider:       req.Quote.Provider,
		ExpectedOutput: req.Quote.ExpectedOutput,
		Quote:          req.Quote,
	}, nil
}

