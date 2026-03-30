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

	// Preference specifies which providers to use and in what order.
	// If nil, uses default provider order.
	Preference *ProviderPreference
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
		Preference:  params.Preference,
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

	route, err := s.router.FindRoute(ctx, from, to, params.Preference)
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

	route, err := s.router.FindRoute(ctx, from, to, params.Preference)
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

// SwapTxBundle contains both approval and swap transactions bundled together.
// For ERC20 swaps on EVM chains, both transactions are included and must be
// signed in a single keysign session with sequential nonces.
type SwapTxBundle struct {
	// Provider that will execute the swap
	Provider string

	// Quote information
	ExpectedOutput *big.Int
	MinimumOutput  *big.Int

	// ApprovalTx is the ERC20 approval transaction (nil if not needed).
	// When present, this transaction uses nonce N.
	ApprovalTx *TxData

	// SwapTx is the main swap transaction (always present).
	// Uses nonce N+1 if ApprovalTx is present, otherwise nonce N.
	SwapTx *TxData

	// Memo for non-EVM chains (UTXO/Cosmos)
	Memo string

	// Original quote for reference
	Quote *Quote

	// NeedsApproval indicates if an approval transaction is required
	NeedsApproval bool
}

// GetSwapTxBundle returns a complete transaction bundle ready for signing.
// For ERC20 tokens on EVM chains, this includes both the approval transaction
// and the swap transaction, bundled together with sequential nonces.
//
// This is the recommended entry point for applications that need to handle
// ERC20 approvals. The returned bundle should be processed as follows:
//
//  1. If bundle.ApprovalTx != nil, create a keysign message for it
//  2. Create a keysign message for bundle.SwapTx
//  3. Sign both messages in a single keysign session
//  4. Broadcast the approval tx first, wait for confirmation
//  5. Broadcast the swap tx
//
// Example:
//
//	bundle, err := service.GetSwapTxBundle(ctx, params)
//	if err != nil {
//	    return err
//	}
//
//	var messages []KeysignMessage
//	if bundle.ApprovalTx != nil {
//	    messages = append(messages, buildKeysignMessage(bundle.ApprovalTx))
//	}
//	messages = append(messages, buildKeysignMessage(bundle.SwapTx))
//
//	signatures := keysign(messages)
//	// ... broadcast in order
func (s *Service) GetSwapTxBundle(ctx context.Context, params SwapParams) (*SwapTxBundle, error) {
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
		Preference:  params.Preference,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get quote: %w", err)
	}

	// Build the swap bundle with approval if needed
	bundle, err := s.router.BuildSwapBundle(ctx, SwapBundleRequest{
		Quote:       quote,
		Sender:      params.Sender,
		Destination: params.Destination,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to build swap bundle: %w", err)
	}

	return &SwapTxBundle{
		Provider:       bundle.Provider,
		ExpectedOutput: bundle.ExpectedOutput,
		MinimumOutput:  quote.MinimumOutput,
		ApprovalTx:     bundle.ApprovalTx,
		SwapTx:         bundle.SwapTx,
		Memo:           quote.Memo,
		Quote:          quote,
		NeedsApproval:  bundle.ApprovalTx != nil,
	}, nil
}

// GetQuote returns a swap quote without building transactions.
// Use this when you need to show the user expected amounts before
// they confirm the swap.
func (s *Service) GetQuote(ctx context.Context, params SwapParams) (*Quote, error) {
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

	return s.router.GetQuote(ctx, QuoteRequest{
		From:        from,
		To:          to,
		Amount:      params.Amount,
		Sender:      params.Sender,
		Destination: params.Destination,
		Preference:  params.Preference,
	})
}

// RequiresApproval checks if a swap will require an ERC20 approval.
// This can be called before GetSwapTxBundle to inform the user.
func (s *Service) RequiresApproval(params SwapParams) bool {
	from := Asset{
		Chain:   params.FromChain,
		Symbol:  params.FromSymbol,
		Address: params.FromAddress,
	}
	return IsApprovalRequired(from)
}

