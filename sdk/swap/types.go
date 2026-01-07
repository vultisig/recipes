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

	// Approval info for ERC20 swaps on EVM chains
	NeedsApproval   bool     // True if this is ERC20 on EVM chain
	ApprovalSpender string   // Router contract address to approve
	ApprovalAmount  *big.Int // Exact amount to approve (always exact, never unlimited)

	// ProviderData stores provider-specific quote data for use in BuildTx.
	// This avoids re-fetching quotes and ensures consistency between quote and tx.
	ProviderData []byte
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

// EVMChains is the list of EVM-compatible chains that support ERC20 approvals
var EVMChains = []string{
	"Ethereum",
	"BSC",
	"Polygon",
	"Arbitrum",
	"Avalanche",
	"Base",
	"Optimism",
	"Blast",
	"CronosChain",
	"ZkSync",
}

// IsEVMChain checks if a chain is an EVM-compatible chain
func IsEVMChain(chain string) bool {
	for _, c := range EVMChains {
		if c == chain {
			return true
		}
	}
	return false
}

// IsApprovalRequired checks if an asset requires ERC20 approval for swapping.
// Returns true if the asset is an ERC20 token (has contract address) on an EVM chain.
func IsApprovalRequired(asset Asset) bool {
	return IsEVMChain(asset.Chain) && asset.Address != ""
}

// TxData represents transaction data ready for signing
type TxData struct {
	To       string   // Destination address
	Value    *big.Int // Native token value (wei for EVM, satoshis for UTXO, etc.)
	Data     []byte   // Transaction calldata (for EVM)
	Memo     string   // Transaction memo (for UTXO/Cosmos chains)
	Nonce    uint64   // Transaction nonce (for EVM)
	GasLimit uint64   // Gas limit (for EVM)
	ChainID  *big.Int // EVM chain ID (nil for non-EVM)
}

// SwapBundle contains the approval tx (if needed) and swap tx bundled together.
// For ERC20 swaps on EVM chains, both transactions are included and must be
// signed in a single keysign session with sequential nonces.
type SwapBundle struct {
	// ApprovalTx is the ERC20 approval transaction (nil if not needed).
	// When present, uses nonce N.
	ApprovalTx *TxData

	// SwapTx is the main swap transaction (always present).
	// Uses nonce N+1 if ApprovalTx is present, otherwise nonce N.
	SwapTx *TxData

	// Provider is the swap provider name
	Provider string

	// ExpectedOutput is the expected amount of destination token
	ExpectedOutput *big.Int

	// Quote is the original quote used to build this bundle
	Quote *Quote
}

// SwapBundleRequest contains parameters for building a swap bundle
type SwapBundleRequest struct {
	Quote       *Quote
	Sender      string   // Sender address
	Destination string   // Destination address for swap output
	Nonce       *uint64  // Optional: starting nonce (fetched if nil)
	GasPrice    *big.Int // Optional: gas price (fetched if nil)
}

