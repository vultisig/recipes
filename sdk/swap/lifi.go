package swap

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
	lifiDefaultBaseURL   = "https://li.quest/v1"
	lifiIntegratorName   = "vultisig"
)

// LiFi chain IDs
// Note: LiFi supports Fantom but we exclude it for consistency with 1inch removal
var lifiChainIDs = map[string]int{
	"Ethereum":  1,
	"BSC":       56,
	"Polygon":   137,
	"Avalanche": 43114,
	"Arbitrum":  42161,
	"Optimism":  10,
	"Base":      8453,
	"Gnosis":    100,
	"Solana":    1151111081099710, // LiFi Solana chain ID
}

// LiFi supported chains (EVM + Solana for cross-chain)
var lifiSupportedChains = []string{
	"Ethereum",
	"BSC",
	"Polygon",
	"Avalanche",
	"Arbitrum",
	"Optimism",
	"Base",
	"Gnosis",
	"Solana",
}

// LiFiProvider implements SwapProvider for LiFi aggregator
type LiFiProvider struct {
	BaseProvider
	client  *http.Client
	baseURL string
	apiKey  string
}

// NewLiFiProvider creates a new LiFi provider
func NewLiFiProvider(apiKey string) *LiFiProvider {
	return &LiFiProvider{
		BaseProvider: NewBaseProvider("LiFi", PriorityLiFi, lifiSupportedChains),
		client: &http.Client{
			Timeout: 30 * time.Second, // LiFi can be slow for complex routes
		},
		baseURL: lifiDefaultBaseURL,
		apiKey:  apiKey,
	}
}

// SupportsRoute checks if LiFi can handle a swap between two assets
func (p *LiFiProvider) SupportsRoute(from, to Asset) bool {
	return p.SupportsChain(from.Chain) && p.SupportsChain(to.Chain)
}

// IsAvailable checks if LiFi is available for a specific chain
// LiFi is generally always available if the chain is supported
func (p *LiFiProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	return p.SupportsChain(chain), nil
}

// GetStatus returns detailed availability status for a chain
func (p *LiFiProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	return &ProviderStatus{
		Chain:     chain,
		Available: p.SupportsChain(chain),
	}, nil
}

// GetQuote gets a swap quote from LiFi
func (p *LiFiProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	fromChainID, ok := lifiChainIDs[req.From.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported from chain: %s", req.From.Chain)
	}
	toChainID, ok := lifiChainIDs[req.To.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported to chain: %s", req.To.Chain)
	}

	// Format token addresses (use ticker for native tokens)
	fromToken := req.From.Address
	if fromToken == "" {
		fromToken = req.From.Symbol
	}
	toToken := req.To.Address
	if toToken == "" {
		toToken = req.To.Symbol
	}

	// Build quote URL
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
	defer resp.Body.Close()

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

	// Determine router address for approval
	routerAddress := ""
	if quoteResp.TransactionRequest != nil {
		routerAddress = quoteResp.TransactionRequest.To
	}

	// Try to get the canonical router from resolver
	if routerInfo, err := ResolveLiFiRouter(req.From.Chain); err == nil && routerInfo.Address != "" {
		routerAddress = routerInfo.Address
	}

	// Check if approval is needed (ERC20 token on EVM chain)
	needsApproval := IsApprovalRequired(req.From)

	quote := &Quote{
		Provider:        p.Name(),
		FromAsset:       req.From,
		ToAsset:         req.To,
		FromAmount:      req.Amount,
		ExpectedOutput:  toAmount,
		Router:          routerAddress,
		NeedsApproval:   needsApproval,
		ApprovalSpender: routerAddress,
		ApprovalAmount:  req.Amount,
	}

	return quote, nil
}

// BuildTx builds an unsigned transaction for the swap
func (p *LiFiProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
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
		fromToken = req.Quote.FromAsset.Symbol
	}
	toToken := req.Quote.ToAsset.Address
	if toToken == "" {
		toToken = req.Quote.ToAsset.Symbol
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
	defer resp.Body.Close()

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

	value := big.NewInt(0)
	if quoteResp.TransactionRequest.Value != "" {
		var ok bool
		value, ok = new(big.Int).SetString(quoteResp.TransactionRequest.Value, 0)
		if !ok {
			return nil, fmt.Errorf("invalid transaction value: %s", quoteResp.TransactionRequest.Value)
		}
	}

	toAmount, ok := new(big.Int).SetString(quoteResp.Estimate.ToAmount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid toAmount: %s", quoteResp.Estimate.ToAmount)
	}

	// Get the approval spender address from resolver
	approvalSpender := quoteResp.TransactionRequest.To
	if routerInfo, err := ResolveLiFiRouter(req.Quote.FromAsset.Chain); err == nil {
		approvalSpender = routerInfo.Address
	}

	// Check if approval is needed using consistent logic
	needsApproval := IsApprovalRequired(req.Quote.FromAsset)

	// Decode hex-encoded transaction data
	txData, err := decodeLiFiHexData(quoteResp.TransactionRequest.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tx data: %w", err)
	}

	return &SwapResult{
		Provider:        p.Name(),
		TxData:          txData,
		Value:           value,
		ToAddress:       quoteResp.TransactionRequest.To,
		ExpectedOut:     toAmount,
		NeedsApproval:   needsApproval,
		ApprovalAddress: approvalSpender,
		ApprovalAmount:  req.Quote.FromAmount,
	}, nil
}

// decodeLiFiHexData decodes a hex-encoded string (with optional 0x prefix) to bytes
func decodeLiFiHexData(hexStr string) ([]byte, error) {
	hexStr = strings.TrimPrefix(hexStr, "0x")
	if hexStr == "" {
		return nil, nil
	}
	return hex.DecodeString(hexStr)
}

// LiFi API types
type lifiQuoteResponse struct {
	Estimate           lifiEstimate           `json:"estimate"`
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

