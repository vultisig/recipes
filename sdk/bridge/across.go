package bridge

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	gethCommon "github.com/ethereum/go-ethereum/common"
)

const (
	acrossDefaultBaseURL = "https://across.to/api"
)

// ============================================================================
// Across Protocol Contract Addresses (SpokePool)
// ============================================================================
// Source: https://docs.across.to/reference/contract-addresses
// VERIFIED: These are the SpokePool contracts for each supported chain
// ============================================================================

var acrossSpokePoolAddresses = map[string]string{
	// Derived from Across `suggested-fees` API (spokePoolAddress) for each chain.
	"Ethereum":    "0x5c7BCd6E7De5423a257D81B442095A1a6ced35C5",
	"Optimism":    "0x6f26Bf09B1C792e3228e5467807a900A503c0281",
	"Arbitrum":    "0xe35e9842fceaCA96570B734083f4a58e8F7C5f2A",
	"Base":        "0x09aea4b2242abC8bb4BB78D537A67a245A7bEC64",
	"Polygon":     "0x9295ee1d8C5b022Be115A2AD3c30C72E34e7F096",
	"Linea":       "0x7E63A5f1a8F0B4d0934B2f2327DAED3F6bb2ee75",
	"Zksync":      "0xE0B015E54d54fc84a6cB9B666099c46adE9335FF",
	"Blast":       "0x2D509190Ed0172ba588407D4c2df918F955Cc6E1",
	"Hyperliquid": "0x35E63eA3eb0fb7A3bc543C71FB66412e1F6B0E04",
}

// Across chain IDs for API calls
var acrossChainIDs = map[string]int{
	"Ethereum":    1,
	"Arbitrum":    42161,
	"Optimism":    10,
	"Base":        8453,
	"Polygon":     137,
	"Linea":       59144,
	"Zksync":      324,
	"Blast":       81457,
	"Hyperliquid": 999,
}

// Across supported chains
var acrossSupportedChains = []string{
	"Ethereum",
	"Arbitrum",
	"Optimism",
	"Base",
	"Polygon",
	"Linea",
	"Zksync",
	"Blast",
	"Hyperliquid",
}

// AcrossProvider implements BridgeProvider for Across Protocol
type AcrossProvider struct {
	BaseProvider
	client  *http.Client
	baseURL string
}

const acrossSpokePoolABIJSON = `[
  {
    "type": "function",
    "name": "depositV3",
    "stateMutability": "payable",
    "inputs": [
      { "name": "depositor", "type": "address" },
      { "name": "recipient", "type": "address" },
      { "name": "inputToken", "type": "address" },
      { "name": "outputToken", "type": "address" },
      { "name": "inputAmount", "type": "uint256" },
      { "name": "outputAmount", "type": "uint256" },
      { "name": "destinationChainId", "type": "uint256" },
      { "name": "exclusiveRelayer", "type": "address" },
      { "name": "quoteTimestamp", "type": "uint32" },
      { "name": "fillDeadline", "type": "uint32" },
      { "name": "exclusivityParameter", "type": "uint32" },
      { "name": "message", "type": "bytes" }
    ],
    "outputs": []
  }
]`

var acrossSpokePoolABI = func() abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(acrossSpokePoolABIJSON))
	if err != nil {
		panic(err)
	}
	return parsed
}()

// NewAcrossProvider creates a new Across provider
func NewAcrossProvider() *AcrossProvider {
	return &AcrossProvider{
		BaseProvider: NewBaseProvider("Across", PriorityAcross, acrossSupportedChains),
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: acrossDefaultBaseURL,
	}
}

// SupportsRoute checks if Across can handle a bridge between two chains
func (p *AcrossProvider) SupportsRoute(from, to BridgeAsset) bool {
	// Must be different chains
	if from.Chain == to.Chain {
		return false
	}
	// Both chains must be supported
	return p.SupportsChain(from.Chain) && p.SupportsChain(to.Chain)
}

// IsAvailable checks if Across is available for a specific chain
func (p *AcrossProvider) IsAvailable(_ context.Context, chain string) (bool, error) {
	return p.SupportsChain(chain), nil
}

// GetStatus returns detailed availability status for a chain
func (p *AcrossProvider) GetStatus(_ context.Context, chain string) (*ProviderStatus, error) {
	bridgeAddr := acrossSpokePoolAddresses[chain]
	return &ProviderStatus{
		Chain:         chain,
		Available:     p.SupportsChain(chain),
		BridgeAddress: bridgeAddr,
	}, nil
}

// GetQuote gets a bridge quote from Across
func (p *AcrossProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	originChainID, ok := acrossChainIDs[req.From.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported from chain: %s", req.From.Chain)
	}
	destinationChainID, ok := acrossChainIDs[req.To.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported to chain: %s", req.To.Chain)
	}

	inputToken := req.From.Address
	if inputToken == "" {
		return nil, fmt.Errorf("Across requires explicit token contract address (native assets not supported yet)")
	}
	outputToken := req.To.Address
	if outputToken == "" {
		return nil, fmt.Errorf("Across requires explicit token contract address (native assets not supported yet)")
	}

	// Build suggested-fees URL (Across's quote endpoint)
	params := url.Values{}
	params.Set("originChainId", fmt.Sprintf("%d", originChainID))
	params.Set("destinationChainId", fmt.Sprintf("%d", destinationChainID))
	params.Set("inputToken", inputToken)
	params.Set("outputToken", outputToken)
	params.Set("amount", req.Amount.String())
	params.Set("recipient", req.Destination)

	quoteURL := fmt.Sprintf("%s/suggested-fees?%s", p.baseURL, params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, quoteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Accept", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call Across API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp acrossErrorResponse
		if json.Unmarshal(body, &errResp) == nil && errResp.Message != "" {
			return nil, fmt.Errorf("Across error: %s", errResp.Message)
		}
		return nil, fmt.Errorf("Across API error (status %d): %s", resp.StatusCode, string(body))
	}

	var quoteResp acrossQuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	totalFee, ok := new(big.Int).SetString(quoteResp.TotalRelayFee.Total, 10)
	if !ok {
		return nil, fmt.Errorf("invalid totalRelayFee.total: %s", quoteResp.TotalRelayFee.Total)
	}
	outputAmount, ok := new(big.Int).SetString(quoteResp.OutputAmount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid outputAmount: %s", quoteResp.OutputAmount)
	}

	// Get bridge contract address (SpokePool)
	router := quoteResp.SpokePoolAddress

	// CRITICAL: Verify the bridge address
	if err := p.verifyBridgeAddress(req.From.Chain, router); err != nil {
		return nil, fmt.Errorf("bridge address verification failed: %w", err)
	}

	return &Quote{
		Provider:       p.Name(),
		FromAsset:      req.From,
		ToAsset:        req.To,
		FromAmount:     req.Amount,
		ExpectedOutput: outputAmount,
		BridgeFee:      totalFee,
		EstimatedTime:  int64(quoteResp.EstimatedFillTimeSec),
		Router:         router,
	}, nil
}

// BuildTx builds an unsigned transaction for the bridge
func (p *AcrossProvider) BuildTx(ctx context.Context, req BridgeRequest) (*BridgeResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	originChainID, ok := acrossChainIDs[req.Quote.FromAsset.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported from chain: %s", req.Quote.FromAsset.Chain)
	}
	destinationChainID, ok := acrossChainIDs[req.Quote.ToAsset.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported to chain: %s", req.Quote.ToAsset.Chain)
	}

	inputToken := req.Quote.FromAsset.Address
	if inputToken == "" {
		return nil, fmt.Errorf("Across requires explicit token contract address (native assets not supported yet)")
	}
	outputToken := req.Quote.ToAsset.Address
	if outputToken == "" {
		return nil, fmt.Errorf("Across requires explicit token contract address (native assets not supported yet)")
	}

	// Fetch fresh suggested-fees to get deposit parameters + SpokePool address
	params := url.Values{}
	params.Set("originChainId", fmt.Sprintf("%d", originChainID))
	params.Set("destinationChainId", fmt.Sprintf("%d", destinationChainID))
	params.Set("inputToken", inputToken)
	params.Set("outputToken", outputToken)
	params.Set("amount", req.Quote.FromAmount.String())
	params.Set("recipient", req.Destination)
	suggestedFeesURL := fmt.Sprintf("%s/suggested-fees?%s", p.baseURL, params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, suggestedFeesURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Accept", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call Across API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Across API error (status %d): %s", resp.StatusCode, string(body))
	}

	var feeResp acrossQuoteResponse
	if err := json.Unmarshal(body, &feeResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	toAddress := feeResp.SpokePoolAddress
	if err := p.verifyBridgeAddress(req.Quote.FromAsset.Chain, toAddress); err != nil {
		return nil, fmt.Errorf("SECURITY: bridge address verification failed: %w", err)
	}

	outputAmount, ok := new(big.Int).SetString(feeResp.OutputAmount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid outputAmount: %s", feeResp.OutputAmount)
	}

	quoteTimestampU64, err := strconv.ParseUint(feeResp.Timestamp, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid timestamp: %w", err)
	}
	fillDeadlineU64, err := strconv.ParseUint(feeResp.FillDeadline, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid fillDeadline: %w", err)
	}
	exclusivityDeadlineU64 := uint64(0)
	if feeResp.ExclusivityDeadline != 0 {
		exclusivityDeadlineU64 = uint64(feeResp.ExclusivityDeadline)
	}

	calldata, err := acrossSpokePoolABI.Pack(
		"depositV3",
		gethCommon.HexToAddress(req.Sender),
		gethCommon.HexToAddress(req.Destination),
		gethCommon.HexToAddress(inputToken),
		gethCommon.HexToAddress(outputToken),
		req.Quote.FromAmount,
		outputAmount,
		big.NewInt(int64(destinationChainID)),
		gethCommon.HexToAddress(feeResp.ExclusiveRelayer),
		uint32(quoteTimestampU64),
		uint32(fillDeadlineU64),
		uint32(exclusivityDeadlineU64),
		[]byte{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build deposit calldata: %w", err)
	}

	// Get approval spender address - use SpokePool
	approvalSpender := toAddress

	// Token bridges need approval (not native token)
	needsApproval := req.Quote.FromAsset.Address != ""

	return &BridgeResult{
		Provider:        p.Name(),
		TxData:          calldata,
		Value:           big.NewInt(0),
		ToAddress:       toAddress,
		ExpectedOut:     req.Quote.ExpectedOutput,
		NeedsApproval:   needsApproval,
		ApprovalAddress: approvalSpender,
		ApprovalAmount:  req.Quote.FromAmount,
	}, nil
}

// verifyBridgeAddress verifies that the bridge address is a known Across SpokePool
// This is CRITICAL for security - sending to wrong address = lost funds
func (p *AcrossProvider) verifyBridgeAddress(chain, address string) error {
	knownAddress, ok := acrossSpokePoolAddresses[chain]
	if !ok {
		return fmt.Errorf("no known Across SpokePool address for chain %s", chain)
	}

	// Case-insensitive comparison for EVM addresses
	if !strings.EqualFold(address, knownAddress) {
		return fmt.Errorf(
			"address mismatch for chain %s: got %s, expected %s",
			chain, address, knownAddress,
		)
	}

	return nil
}

// Across API response types
type acrossQuoteResponse struct {
	EstimatedFillTimeSec int            `json:"estimatedFillTimeSec"`
	SpokePoolAddress     string         `json:"spokePoolAddress"`
	DestinationSpokePool string         `json:"destinationSpokePoolAddress"`
	Timestamp            string         `json:"timestamp"`
	FillDeadline         string         `json:"fillDeadline"`
	ExclusiveRelayer     string         `json:"exclusiveRelayer"`
	ExclusivityDeadline  int            `json:"exclusivityDeadline"`
	OutputAmount         string         `json:"outputAmount"`
	TotalRelayFee        acrossRelayFee `json:"totalRelayFee"`
}

type acrossRelayFee struct {
	Total string `json:"total"`
	Pct   string `json:"pct"`
}

type acrossErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
