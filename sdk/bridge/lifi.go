package bridge

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	lifiDefaultBaseURL = "https://li.quest/v1"
	lifiIntegratorName = "vultisig"
)

// LiFi Diamond contract addresses per chain
// Source: https://docs.li.fi/list-of-all-lifi-contract-addresses
// VERIFIED: These addresses are the same across all EVM chains (Diamond proxy pattern)
var lifiDiamondAddresses = map[string]string{
	"Ethereum":  "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Arbitrum":  "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Optimism":  "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Polygon":   "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"BSC":       "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Avalanche": "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Base":      "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Fantom":    "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Gnosis":    "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Blast":     "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Zksync":    "0x341e94069f53234fE6DabeF707aD424830525715", // zkSync has different address
}

// LiFi chain IDs for API calls
var lifiChainIDs = map[string]int{
	"Ethereum":  1,
	"BSC":       56,
	"Polygon":   137,
	"Avalanche": 43114,
	"Arbitrum":  42161,
	"Optimism":  10,
	"Base":      8453,
	"Fantom":    250,
	"Gnosis":    100,
	"Blast":     81457,
	"Zksync":    324,
}

// LiFi supported chains for bridging
var lifiSupportedChains = []string{
	"Ethereum",
	"BSC",
	"Polygon",
	"Avalanche",
	"Arbitrum",
	"Optimism",
	"Base",
	"Fantom",
	"Gnosis",
	"Blast",
	"Zksync",
}

// LiFiProvider implements BridgeProvider for LiFi bridge aggregator
type LiFiProvider struct {
	BaseProvider
	client  *http.Client
	baseURL string
	apiKey  string
}

// NewLiFiProvider creates a new LiFi bridge provider
func NewLiFiProvider(apiKey string) *LiFiProvider {
	return &LiFiProvider{
		BaseProvider: NewBaseProvider("LiFi", PriorityLiFi, lifiSupportedChains),
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: lifiDefaultBaseURL,
		apiKey:  apiKey,
	}
}

// SupportsRoute checks if LiFi can handle a bridge between two chains
func (p *LiFiProvider) SupportsRoute(from, to BridgeAsset) bool {
	// Must be different chains
	if from.Chain == to.Chain {
		return false
	}
	// Both chains must be supported
	return p.SupportsChain(from.Chain) && p.SupportsChain(to.Chain)
}

// IsAvailable checks if LiFi is available for a specific chain
func (p *LiFiProvider) IsAvailable(_ context.Context, chain string) (bool, error) {
	return p.SupportsChain(chain), nil
}

// GetStatus returns detailed availability status for a chain
func (p *LiFiProvider) GetStatus(_ context.Context, chain string) (*ProviderStatus, error) {
	bridgeAddr := lifiDiamondAddresses[chain]
	return &ProviderStatus{
		Chain:         chain,
		Available:     p.SupportsChain(chain),
		BridgeAddress: bridgeAddr,
	}, nil
}

// GetQuote gets a bridge quote from LiFi
func (p *LiFiProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	fromChainID, ok := lifiChainIDs[req.From.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported from chain: %s", req.From.Chain)
	}
	toChainID, ok := lifiChainIDs[req.To.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported to chain: %s", req.To.Chain)
	}

	// Format token addresses
	fromToken := req.From.Address
	if fromToken == "" {
		fromToken = "0x0000000000000000000000000000000000000000" // Native token
	}
	toToken := req.To.Address
	if toToken == "" {
		toToken = "0x0000000000000000000000000000000000000000" // Native token
	}

	// Build quote URL with bridge-specific parameters
	params := url.Values{}
	params.Set("fromChain", fmt.Sprintf("%d", fromChainID))
	params.Set("toChain", fmt.Sprintf("%d", toChainID))
	params.Set("fromToken", fromToken)
	params.Set("toToken", toToken)
	params.Set("fromAmount", req.Amount.String())
	params.Set("fromAddress", req.Sender)
	params.Set("toAddress", req.Destination)
	params.Set("integrator", lifiIntegratorName)

	quoteURL := fmt.Sprintf("%s/quote?%s", p.baseURL, params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, quoteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Accept", "application/json")
	if p.apiKey != "" {
		httpReq.Header.Set("x-lifi-api-key", p.apiKey)
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call LiFi API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp lifiErrorResponse
		if json.Unmarshal(body, &errResp) == nil && errResp.Message != "" {
			return nil, fmt.Errorf("LiFi error: %s", errResp.Message)
		}
		return nil, fmt.Errorf("LiFi API error (status %d): %s", resp.StatusCode, string(body))
	}

	var quoteResp lifiQuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	toAmount, ok := new(big.Int).SetString(quoteResp.Estimate.ToAmount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid toAmount: %s", quoteResp.Estimate.ToAmount)
	}

	// Get bridge router address
	router := lifiDiamondAddresses[req.From.Chain]
	if quoteResp.TransactionRequest != nil && quoteResp.TransactionRequest.To != "" {
		router = quoteResp.TransactionRequest.To
	}

	// CRITICAL: Verify the router address matches known LiFi addresses
	if err := p.verifyBridgeAddress(req.From.Chain, router); err != nil {
		return nil, fmt.Errorf("bridge address verification failed: %w", err)
	}

	return &Quote{
		Provider:       p.Name(),
		FromAsset:      req.From,
		ToAsset:        req.To,
		FromAmount:     req.Amount,
		ExpectedOutput: toAmount,
		Router:         router,
	}, nil
}

// BuildTx builds an unsigned transaction for the bridge
func (p *LiFiProvider) BuildTx(ctx context.Context, req BridgeRequest) (*BridgeResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	fromChainID, ok := lifiChainIDs[req.Quote.FromAsset.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported from chain: %s", req.Quote.FromAsset.Chain)
	}
	toChainID, ok := lifiChainIDs[req.Quote.ToAsset.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported to chain: %s", req.Quote.ToAsset.Chain)
	}

	fromToken := req.Quote.FromAsset.Address
	if fromToken == "" {
		fromToken = "0x0000000000000000000000000000000000000000"
	}
	toToken := req.Quote.ToAsset.Address
	if toToken == "" {
		toToken = "0x0000000000000000000000000000000000000000"
	}

	// Re-fetch quote with transaction data
	params := url.Values{}
	params.Set("fromChain", fmt.Sprintf("%d", fromChainID))
	params.Set("toChain", fmt.Sprintf("%d", toChainID))
	params.Set("fromToken", fromToken)
	params.Set("toToken", toToken)
	params.Set("fromAmount", req.Quote.FromAmount.String())
	params.Set("fromAddress", req.Sender)
	params.Set("toAddress", req.Destination)
	params.Set("integrator", lifiIntegratorName)

	quoteURL := fmt.Sprintf("%s/quote?%s", p.baseURL, params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, quoteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Accept", "application/json")
	if p.apiKey != "" {
		httpReq.Header.Set("x-lifi-api-key", p.apiKey)
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call LiFi API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("LiFi API error (status %d): %s", resp.StatusCode, string(body))
	}

	var quoteResp lifiQuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if quoteResp.TransactionRequest == nil {
		return nil, fmt.Errorf("no transaction request in LiFi response")
	}

	// CRITICAL: Triple-verify the bridge address before building transaction
	toAddress := quoteResp.TransactionRequest.To
	if err := p.verifyBridgeAddress(req.Quote.FromAsset.Chain, toAddress); err != nil {
		return nil, fmt.Errorf("SECURITY: bridge address verification failed: %w", err)
	}

	value := big.NewInt(0)
	if quoteResp.TransactionRequest.Value != "" {
		var ok bool
		value, ok = new(big.Int).SetString(quoteResp.TransactionRequest.Value, 0)
		if !ok {
			return nil, fmt.Errorf("failed to parse transaction value: %s", quoteResp.TransactionRequest.Value)
		}
	}

	toAmount, ok := new(big.Int).SetString(quoteResp.Estimate.ToAmount, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse toAmount: %s", quoteResp.Estimate.ToAmount)
	}

	// Decode hex-encoded transaction data from API response
	txDataHex := strings.TrimPrefix(quoteResp.TransactionRequest.Data, "0x")
	txData, err := hex.DecodeString(txDataHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode transaction data: %w", err)
	}

	// Get the approval spender address - use LiFi Diamond
	approvalSpender := lifiDiamondAddresses[req.Quote.FromAsset.Chain]

	// Token bridges need approval (not native token)
	needsApproval := req.Quote.FromAsset.Address != ""

	return &BridgeResult{
		Provider:        p.Name(),
		TxData:          txData,
		Value:           value,
		ToAddress:       toAddress,
		ExpectedOut:     toAmount,
		NeedsApproval:   needsApproval,
		ApprovalAddress: approvalSpender,
		ApprovalAmount:  req.Quote.FromAmount,
	}, nil
}

// verifyBridgeAddress verifies that the bridge address is a known LiFi address
// This is CRITICAL for security - sending to wrong address = lost funds
func (p *LiFiProvider) verifyBridgeAddress(chain, address string) error {
	knownAddress, ok := lifiDiamondAddresses[chain]
	if !ok {
		return fmt.Errorf("no known LiFi address for chain %s", chain)
	}

	// Case-insensitive comparison for addresses
	if !strings.EqualFold(address, knownAddress) {
		return fmt.Errorf(
			"address mismatch for chain %s: got %s, expected %s",
			chain, address, knownAddress,
		)
	}

	return nil
}

// LiFi API response types
type lifiQuoteResponse struct {
	Estimate           lifiEstimate            `json:"estimate"`
	TransactionRequest *lifiTransactionRequest `json:"transactionRequest,omitempty"`
}

type lifiEstimate struct {
	FromAmount string `json:"fromAmount"`
	ToAmount   string `json:"toAmount"`
}

type lifiTransactionRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Data     string `json:"data"`
	Value    string `json:"value"`
	GasLimit string `json:"gasLimit,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
}

type lifiErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
