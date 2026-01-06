package bridge

import (
	"context"
	"math/big"
)

// BridgeAsset represents a token/coin on a specific chain for bridging
type BridgeAsset struct {
	Chain    string // Chain identifier (e.g., "Ethereum", "Arbitrum")
	Symbol   string // Token symbol (e.g., "ETH", "USDC")
	Address  string // Contract address for tokens, empty for native coins
	Decimals int    // Token decimals
}

// Provider priority constants - lower number = higher priority
const (
	PriorityNative   = 1 // Native L2 bridges, cheapest + most canonical for L1 <-> L2
	PriorityLiFi     = 2 // LiFi aggregates multiple bridges, best rates for general cross-chain
	PriorityAcross   = 3 // Across Protocol, fast cross-chain, supports Hyperliquid
	PriorityDeBridge = 4 // deBridge (DLN), cross-chain liquidity network (EVM-only in this implementation)
)

// BridgeProvider defines the interface for bridge providers
type BridgeProvider interface {
	// Name returns the provider's name
	Name() string

	// Priority returns the provider's priority (lower = higher priority)
	Priority() int

	// SupportedChains returns the list of chains this provider supports
	SupportedChains() []string

	// SupportsRoute checks if the provider can handle a bridge between two chains
	SupportsRoute(from, to BridgeAsset) bool

	// IsAvailable checks if the provider is available for a specific chain
	IsAvailable(ctx context.Context, chain string) (bool, error)

	// GetStatus returns detailed availability status for a chain
	GetStatus(ctx context.Context, chain string) (*ProviderStatus, error)

	// GetQuote gets a bridge quote from the provider
	GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error)

	// BuildTx builds an unsigned transaction for the bridge
	BuildTx(ctx context.Context, req BridgeRequest) (*BridgeResult, error)
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

// QuoteRequest contains parameters for getting a bridge quote
type QuoteRequest struct {
	From        BridgeAsset // Source asset
	To          BridgeAsset // Destination asset (same token, different chain)
	Amount      *big.Int    // Amount in smallest unit (e.g., wei)
	Sender      string      // Sender address on source chain
	Destination string      // Recipient address on destination chain
}

// Quote represents a bridge quote from a provider
type Quote struct {
	Provider       string      // Provider name
	FromAsset      BridgeAsset // Source asset
	ToAsset        BridgeAsset // Destination asset
	FromAmount     *big.Int    // Input amount
	ExpectedOutput *big.Int    // Expected output amount
	BridgeFee      *big.Int    // Bridge fee
	EstimatedTime  int64       // Estimated time in seconds
	Router         string      // Router/bridge contract address
	Expiry         int64       // Quote expiry timestamp
}

// BridgeRequest contains all parameters needed to build a bridge transaction
type BridgeRequest struct {
	Quote       *Quote
	Sender      string // Sender address
	Destination string // Destination address on target chain
}

// BridgeResult contains the result of building a bridge transaction
type BridgeResult struct {
	Provider    string   // Provider that was used
	TxData      []byte   // Unsigned transaction data
	Value       *big.Int // Native token value to send (for EVM)
	ToAddress   string   // Transaction destination address (bridge contract)
	ExpectedOut *big.Int // Expected output amount

	// Approval info (for ERC20 bridges)
	NeedsApproval   bool     // True if token approval is needed
	ApprovalAddress string   // Spender address for approval
	ApprovalAmount  *big.Int // Amount to approve (nil = unlimited)
}

// ProviderStatus represents the availability status of a provider for a chain
type ProviderStatus struct {
	Chain         string
	Available     bool
	BridgeAddress string // Bridge contract address for this chain
}
