package swap

import (
	"context"
	"fmt"
	"math/big"
)

// ChainAdapter provides a chain-specific interface to the canonical swap router.
// This allows applications like app-recurring to use the router with their
// existing chain-specific architecture.
type ChainAdapter struct {
	router *Router
	chain  string
}

// NewChainAdapter creates a new adapter for a specific chain.
func NewChainAdapter(chain string) *ChainAdapter {
	return &ChainAdapter{
		router: NewDefaultRouter(),
		chain:  chain,
	}
}

// NewChainAdapterWithRouter creates a new adapter with a custom router.
func NewChainAdapterWithRouter(chain string, router *Router) *ChainAdapter {
	return &ChainAdapter{
		router: router,
		chain:  chain,
	}
}

// SwapInput contains input parameters for a swap.
// This matches the pattern used in app-recurring's Provider interface.
type SwapInput struct {
	FromToken    string   // Token address (empty for native)
	FromSymbol   string   // Token symbol
	FromDecimals int      // Token decimals
	FromAmount   *big.Int // Amount in base units
	FromAddress  string   // Sender address

	ToChain    string // Destination chain
	ToToken    string // Token address (empty for native)
	ToSymbol   string // Token symbol
	ToDecimals int    // Token decimals
	ToAddress  string // Destination address
}

// SwapOutput contains the result of preparing a swap.
type SwapOutput struct {
	// Provider used
	Provider string

	// Quote info
	ExpectedAmountOut *big.Int
	MinimumAmountOut  *big.Int

	// Transaction data
	TxTo    string   // Destination address (vault/router)
	TxValue *big.Int // Native token value
	TxData  []byte   // Calldata (for EVM)
	TxMemo  string   // Memo (for UTXO/Cosmos)

	// For EVM
	RouterAddress string
}

// MakeTx prepares a swap transaction using the canonical router.
// This method signature matches the Provider interface in app-recurring.
func (a *ChainAdapter) MakeTx(ctx context.Context, input SwapInput) (*big.Int, []byte, error) {
	from := Asset{
		Chain:    a.chain,
		Symbol:   input.FromSymbol,
		Address:  input.FromToken,
		Decimals: input.FromDecimals,
	}

	toChain := input.ToChain
	if toChain == "" {
		toChain = a.chain // Same-chain swap
	}

	to := Asset{
		Chain:    toChain,
		Symbol:   input.ToSymbol,
		Address:  input.ToToken,
		Decimals: input.ToDecimals,
	}

	// Get quote
	quote, err := a.router.GetQuote(ctx, QuoteRequest{
		From:        from,
		To:          to,
		Amount:      input.FromAmount,
		Sender:      input.FromAddress,
		Destination: input.ToAddress,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get quote: %w", err)
	}

	// Build transaction
	result, err := a.router.BuildTx(ctx, SwapRequest{
		Quote:       quote,
		Sender:      input.FromAddress,
		Destination: input.ToAddress,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build tx: %w", err)
	}

	return quote.ExpectedOutput, result.TxData, nil
}

// GetSwap prepares a swap with full output details.
func (a *ChainAdapter) GetSwap(ctx context.Context, input SwapInput) (*SwapOutput, error) {
	from := Asset{
		Chain:    a.chain,
		Symbol:   input.FromSymbol,
		Address:  input.FromToken,
		Decimals: input.FromDecimals,
	}

	toChain := input.ToChain
	if toChain == "" {
		toChain = a.chain
	}

	to := Asset{
		Chain:    toChain,
		Symbol:   input.ToSymbol,
		Address:  input.ToToken,
		Decimals: input.ToDecimals,
	}

	quote, err := a.router.GetQuote(ctx, QuoteRequest{
		From:        from,
		To:          to,
		Amount:      input.FromAmount,
		Sender:      input.FromAddress,
		Destination: input.ToAddress,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quote: %w", err)
	}

	result, err := a.router.BuildTx(ctx, SwapRequest{
		Quote:       quote,
		Sender:      input.FromAddress,
		Destination: input.ToAddress,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to build tx: %w", err)
	}

	return &SwapOutput{
		Provider:          quote.Provider,
		ExpectedAmountOut: quote.ExpectedOutput,
		MinimumAmountOut:  quote.MinimumOutput,
		TxTo:              result.ToAddress,
		TxValue:           result.Value,
		TxData:            result.TxData,
		TxMemo:            result.Memo,
		RouterAddress:     quote.Router,
	}, nil
}

// IsAvailable checks if swap functionality is available for this chain.
func (a *ChainAdapter) IsAvailable(ctx context.Context) (bool, error) {
	route, err := a.router.FindRoute(ctx, Asset{Chain: a.chain}, Asset{Chain: a.chain})
	if err != nil {
		return false, nil
	}
	return route.IsSupported, nil
}

// GetStatus returns the status of swap providers for this chain.
func (a *ChainAdapter) GetStatus(ctx context.Context) (*ProviderStatus, error) {
	// Try providers in priority order
	providers := []string{"THORChain", "Mayachain", "LiFi", "1inch", "Jupiter", "Uniswap"}

	for _, p := range providers {
		status, err := a.router.GetProviderStatus(ctx, p, a.chain)
		if err == nil && status.Available {
			return status, nil
		}
	}

	return &ProviderStatus{
		Chain:     a.chain,
		Available: false,
	}, nil
}

// MultiChainAdapter provides swap functionality across multiple chains.
type MultiChainAdapter struct {
	router   *Router
	adapters map[string]*ChainAdapter
}

// NewMultiChainAdapter creates a new multi-chain adapter.
func NewMultiChainAdapter() *MultiChainAdapter {
	router := NewDefaultRouter()
	return &MultiChainAdapter{
		router:   router,
		adapters: make(map[string]*ChainAdapter),
	}
}

// GetChainAdapter returns the adapter for a specific chain.
func (m *MultiChainAdapter) GetChainAdapter(chain string) *ChainAdapter {
	if adapter, ok := m.adapters[chain]; ok {
		return adapter
	}

	adapter := NewChainAdapterWithRouter(chain, m.router)
	m.adapters[chain] = adapter
	return adapter
}

// Swap performs a swap between any two assets.
func (m *MultiChainAdapter) Swap(ctx context.Context, input SwapInput, fromChain string) (*SwapOutput, error) {
	adapter := m.GetChainAdapter(fromChain)
	return adapter.GetSwap(ctx, input)
}

// ValidateRoute checks if a swap route is available.
func (m *MultiChainAdapter) ValidateRoute(ctx context.Context, fromChain, toChain string) error {
	from := Asset{Chain: fromChain}
	to := Asset{Chain: toChain}

	route, err := m.router.FindRoute(ctx, from, to)
	if err != nil {
		return err
	}

	if !route.IsSupported {
		return fmt.Errorf("route from %s to %s not supported", fromChain, toChain)
	}

	return nil
}

