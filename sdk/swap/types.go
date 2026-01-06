package swap

import (
	"context"
	"math/big"
)

// Asset represents a token/coin on a specific chain
type Asset struct {
	Chain    string // Chain identifier (e.g., "Bitcoin", "Ethereum")
	Symbol   string // Token symbol (e.g., "BTC", "ETH", "USDC")
	Address  string // Contract address for tokens, empty for native coins
	Decimals int    // Token decimals
}

// Provider priority constants - lower number = higher priority
const (
	PriorityTHORChain = 1
	PriorityMayachain = 2
	PriorityLiFi      = 3
	PriorityOneInch   = 4
	PriorityJupiter   = 5
	PriorityUniswap   = 6
)

// SwapProvider defines the interface for swap providers
type SwapProvider interface {
	// Name returns the provider's name
	Name() string

	// Priority returns the provider's priority (lower = higher priority)
	Priority() int

	// SupportedChains returns the list of chains this provider supports
	SupportedChains() []string

	// SupportsRoute checks if the provider can handle a swap between two assets
	SupportsRoute(from, to Asset) bool

	// IsAvailable checks if the provider is available for a specific chain
	// For THORChain/Maya this checks halted status
	// For other providers this may check API availability
	IsAvailable(ctx context.Context, chain string) (bool, error)

	// GetStatus returns detailed availability status for a chain
	GetStatus(ctx context.Context, chain string) (*ProviderStatus, error)

	// GetQuote gets a swap quote from the provider
	GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error)

	// BuildTx builds an unsigned transaction for the swap
	BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error)
}

// BaseProvider provides common functionality for providers
type BaseProvider struct {
	ProviderName     string
	ProviderPriority int
	ProviderChains   []string
}

// NewBaseProvider creates a new base provider
func NewBaseProvider(name string, priority int, chains []string) BaseProvider {
	return BaseProvider{
		ProviderName:     name,
		ProviderPriority: priority,
		ProviderChains:   chains,
	}
}

// Name returns the provider's name
func (p *BaseProvider) Name() string {
	return p.ProviderName
}

// Priority returns the provider's priority
func (p *BaseProvider) Priority() int {
	return p.ProviderPriority
}

// SupportedChains returns the list of supported chains
func (p *BaseProvider) SupportedChains() []string {
	return p.ProviderChains
}

// SupportsChain checks if a chain is supported
func (p *BaseProvider) SupportsChain(chain string) bool {
	for _, c := range p.ProviderChains {
		if c == chain {
			return true
		}
	}
	return false
}

// QuoteRequest contains parameters for getting a swap quote
type QuoteRequest struct {
	From        Asset
	To          Asset
	Amount      *big.Int // Amount in smallest unit (e.g., satoshis, wei)
	Destination string   // Destination address for the swap output
	Sender      string   // Sender address
}

// Quote represents a swap quote from a provider
type Quote struct {
	Provider        string   // Provider name
	FromAsset       Asset    // Source asset
	ToAsset         Asset    // Destination asset
	FromAmount      *big.Int // Input amount
	ExpectedOutput  *big.Int // Expected output amount
	MinimumOutput   *big.Int // Minimum output with slippage
	EstimatedFees   *big.Int // Estimated fees
	Memo            string   // Transaction memo (for THORChain/Maya)
	InboundAddress  string   // Inbound vault address (for THORChain/Maya)
	Router          string   // Router contract address (for EVM)
	Expiry          int64    // Quote expiry timestamp
	StreamingSwap   bool     // Whether this is a streaming swap
	StreamingBlocks int64    // Number of streaming blocks
}

// SwapRequest contains all parameters needed to build a swap transaction
type SwapRequest struct {
	Quote       *Quote
	Sender      string // Sender address
	Destination string // Destination address for swap output
}

// SwapResult contains the result of building a swap transaction
type SwapResult struct {
	Provider    string   // Provider that was used
	TxData      []byte   // Unsigned transaction data
	Value       *big.Int // Native token value to send (for EVM)
	ToAddress   string   // Transaction destination address
	Memo        string   // Transaction memo
	ExpectedOut *big.Int // Expected output amount

	// Approval info (for ERC20 swaps)
	NeedsApproval   bool     // True if token approval is needed
	ApprovalAddress string   // Spender address for approval
	ApprovalAmount  *big.Int // Amount to approve (nil = unlimited)
}

// RouteResult contains the result of finding a swap route
type RouteResult struct {
	Provider    string // Provider name that can handle the route
	IsSupported bool   // Whether the route is supported
	Quote       *Quote // Optional quote if requested
}

// ProviderStatus represents the availability status of a provider for a chain
type ProviderStatus struct {
	Chain               string
	Available           bool
	Halted              bool
	GlobalTradingPaused bool
	ChainTradingPaused  bool
	Router              string
	InboundAddress      string
}

