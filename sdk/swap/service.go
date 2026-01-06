package swap

import (
	"context"
	"fmt"
	"math/big"
)

// Service provides a high-level swap service for applications.
// This is the recommended way for apps like app-recurring and verifier
// to interact with the swap functionality.
type Service struct {
	router *Router
}

// NewService creates a new swap service with the default router.
func NewService() *Service {
	return &Service{
		router: NewDefaultRouter(),
	}
}

// NewServiceWithRouter creates a new swap service with a custom router.
func NewServiceWithRouter(router *Router) *Service {
	return &Service{
		router: router,
	}
}

// SwapParams contains all parameters needed to execute a swap.
type SwapParams struct {
	// Source asset
	FromChain    string
	FromSymbol   string
	FromAddress  string // Contract address (empty for native)
	FromDecimals int

	// Destination asset
	ToChain    string
	ToSymbol   string
	ToAddress  string // Contract address (empty for native)
	ToDecimals int

	// Swap parameters
	Amount      *big.Int // Amount in base units
	Sender      string   // Sender wallet address
	Destination string   // Destination wallet address
}

// SwapTx contains the transaction data ready for signing.
type SwapTx struct {
	// Provider that will execute the swap
	Provider string

	// Quote information
	ExpectedOutput *big.Int
	MinimumOutput  *big.Int

	// Transaction data
	ToAddress string   // Vault/router address to send to
	Value     *big.Int // Native token value (for EVM)
	Data      []byte   // Transaction data
	Memo      string   // Memo (for UTXO/Cosmos chains)

	// For EVM chains
	GasLimit uint64

	// Original quote for reference
	Quote *Quote
}

// GetSwapTx returns transaction data ready for signing.
// This is the main entry point for applications.
func (s *Service) GetSwapTx(ctx context.Context, params SwapParams) (*SwapTx, error) {
	from := Asset{
		Chain:    params.FromChain,
		Symbol:   params.FromSymbol,
		Address:  params.FromAddress,
		Decimals: params.FromDecimals,
	}

	to := Asset{
		Chain:    params.ToChain,
		Symbol:   params.ToSymbol,
		Address:  params.ToAddress,
		Decimals: params.ToDecimals,
	}

	// Get quote
	quote, err := s.router.GetQuote(ctx, QuoteRequest{
		From:        from,
		To:          to,
		Amount:      params.Amount,
		Sender:      params.Sender,
		Destination: params.Destination,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quote: %w", err)
	}

	// Build transaction
	result, err := s.router.BuildTx(ctx, SwapRequest{
		Quote:       quote,
		Sender:      params.Sender,
		Destination: params.Destination,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to build transaction: %w", err)
	}

	return &SwapTx{
		Provider:       quote.Provider,
		ExpectedOutput: quote.ExpectedOutput,
		MinimumOutput:  quote.MinimumOutput,
		ToAddress:      result.ToAddress,
		Value:          result.Value,
		Data:           result.TxData,
		Memo:           result.Memo,
		Quote:          quote,
	}, nil
}

// ValidateRoute checks if a swap route is available without getting a quote.
func (s *Service) ValidateRoute(ctx context.Context, params SwapParams) error {
	from := Asset{
		Chain:    params.FromChain,
		Symbol:   params.FromSymbol,
		Address:  params.FromAddress,
		Decimals: params.FromDecimals,
	}

	to := Asset{
		Chain:    params.ToChain,
		Symbol:   params.ToSymbol,
		Address:  params.ToAddress,
		Decimals: params.ToDecimals,
	}

	route, err := s.router.FindRoute(ctx, from, to)
	if err != nil {
		return fmt.Errorf("no route available: %w", err)
	}

	if !route.IsSupported {
		return fmt.Errorf("route not supported")
	}

	return nil
}

// GetProviderForRoute returns the provider that would handle a swap.
func (s *Service) GetProviderForRoute(ctx context.Context, params SwapParams) (string, error) {
	from := Asset{
		Chain:    params.FromChain,
		Symbol:   params.FromSymbol,
		Address:  params.FromAddress,
		Decimals: params.FromDecimals,
	}

	to := Asset{
		Chain:    params.ToChain,
		Symbol:   params.ToSymbol,
		Address:  params.ToAddress,
		Decimals: params.ToDecimals,
	}

	route, err := s.router.FindRoute(ctx, from, to)
	if err != nil {
		return "", err
	}

	return route.Provider, nil
}

// IsCrossChain returns true if the swap is between different chains.
func (s *Service) IsCrossChain(params SwapParams) bool {
	return params.FromChain != params.ToChain
}

// GetSupportedChains returns all chains supported by at least one provider.
func (s *Service) GetSupportedChains() []string {
	return s.router.GetSupportedChains()
}

// IsChainSupported checks if a chain is supported for swapping.
func (s *Service) IsChainSupported(chain string) bool {
	for _, c := range s.router.GetSupportedChains() {
		if c == chain {
			return true
		}
	}
	return false
}

// GetChainStatus returns the availability status of a chain for swapping.
func (s *Service) GetChainStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	// Try THORChain first (highest priority for cross-chain)
	status, err := s.router.GetProviderStatus(ctx, "THORChain", chain)
	if err == nil && status.Available {
		return status, nil
	}

	// Try Mayachain
	status, err = s.router.GetProviderStatus(ctx, "Mayachain", chain)
	if err == nil && status.Available {
		return status, nil
	}

	// For EVM/Solana chains, check LiFi
	status, err = s.router.GetProviderStatus(ctx, "LiFi", chain)
	if err == nil && status.Available {
		return status, nil
	}

	// Return last status even if not available
	if status != nil {
		return status, nil
	}

	return &ProviderStatus{
		Chain:     chain,
		Available: false,
	}, nil
}

