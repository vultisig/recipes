package swap

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"
)

const (
	oneInchAPIVersion      = "v6.0"
	oneInchDefaultSlippage = 1
	oneInchNativeToken     = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
	oneInchDefaultBaseURL  = "https://api.1inch.dev"
)

// 1inch chain IDs
// Note: Fantom removed - 1inch uses Router V4 on Fantom, not V6, and is being deprecated
var oneInchChainIDs = map[string]int{
	"Ethereum":  1,
	"BSC":       56,
	"Polygon":   137,
	"Avalanche": 43114,
	"Arbitrum":  42161,
	"Optimism":  10,
	"Base":      8453,
	"Gnosis":    100,
}

// 1inch supported EVM chains (same-chain swaps only)
var oneInchSupportedChains = []string{
	"Ethereum",
	"BSC",
	"Polygon",
	"Avalanche",
	"Arbitrum",
	"Optimism",
	"Base",
	"Gnosis",
}

// OneInchProvider implements SwapProvider for 1inch aggregator
type OneInchProvider struct {
	BaseProvider
	client  *http.Client
	baseURL string
	apiKey  string
}

// NewOneInchProvider creates a new 1inch provider
func NewOneInchProvider(apiKey string) *OneInchProvider {
	return &OneInchProvider{
		BaseProvider: NewBaseProvider("1inch", PriorityOneInch, oneInchSupportedChains),
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		baseURL: oneInchDefaultBaseURL,
		apiKey:  apiKey,
	}
}

// SupportsRoute checks if 1inch can handle a swap between two assets
// 1inch only supports same-chain swaps
func (p *OneInchProvider) SupportsRoute(from, to Asset) bool {
	return from.Chain == to.Chain && p.SupportsChain(from.Chain)
}

// IsAvailable checks if 1inch is available for a specific chain
func (p *OneInchProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	return p.SupportsChain(chain), nil
}

// GetStatus returns detailed availability status for a chain
func (p *OneInchProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	return &ProviderStatus{
		Chain:     chain,
		Available: p.SupportsChain(chain),
	}, nil
}

// GetQuote gets a swap quote from 1inch
func (p *OneInchProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	chainID, ok := oneInchChainIDs[req.From.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported chain: %s", req.From.Chain)
	}

	// Format token addresses
	srcToken := req.From.Address
	if srcToken == "" {
		srcToken = oneInchNativeToken
	}
	dstToken := req.To.Address
	if dstToken == "" {
		dstToken = oneInchNativeToken
	}

	// Build request URL
	endpoint := fmt.Sprintf("%s/swap/%s/%d/swap", p.baseURL, oneInchAPIVersion, chainID)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add query parameters
	q := httpReq.URL.Query()
	q.Set("src", srcToken)
	q.Set("dst", dstToken)
	q.Set("amount", req.Amount.String())
	q.Set("from", req.Sender)
	q.Set("slippage", fmt.Sprintf("%d", oneInchDefaultSlippage))
	q.Set("disableEstimate", "true")
	q.Set("allowPartialFill", "false")
	q.Set("compatibility", "true")
	httpReq.URL.RawQuery = q.Encode()

	// Add API key header if provided
	if p.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call 1inch API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("1inch API error (status %d): %s", resp.StatusCode, string(body))
	}

	var swapResp oneInchSwapResponse
	if err := json.Unmarshal(body, &swapResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	dstAmount, ok := new(big.Int).SetString(swapResp.DstAmount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid dstAmount: %s", swapResp.DstAmount)
	}

	// Determine router address for approval
	routerAddress := swapResp.Tx.To
	if routerInfo, err := ResolveOneInchRouter(req.From.Chain); err == nil && routerInfo.Address != "" {
		routerAddress = routerInfo.Address
	}

	// Check if approval is needed (ERC20 token)
	needsApproval := IsApprovalRequired(req.From)

	return &Quote{
		Provider:        p.Name(),
		FromAsset:       req.From,
		ToAsset:         req.To,
		FromAmount:      req.Amount,
		ExpectedOutput:  dstAmount,
		Router:          routerAddress,
		NeedsApproval:   needsApproval,
		ApprovalSpender: routerAddress,
		ApprovalAmount:  req.Amount,
	}, nil
}

// BuildTx builds an unsigned transaction for the swap
func (p *OneInchProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	// For 1inch, we need to re-fetch with full tx data
	chainID, ok := oneInchChainIDs[req.Quote.FromAsset.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported chain: %s", req.Quote.FromAsset.Chain)
	}

	srcToken := req.Quote.FromAsset.Address
	if srcToken == "" {
		srcToken = oneInchNativeToken
	}
	dstToken := req.Quote.ToAsset.Address
	if dstToken == "" {
		dstToken = oneInchNativeToken
	}

	endpoint := fmt.Sprintf("%s/swap/%s/%d/swap", p.baseURL, oneInchAPIVersion, chainID)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	q := httpReq.URL.Query()
	q.Set("src", srcToken)
	q.Set("dst", dstToken)
	q.Set("amount", req.Quote.FromAmount.String())
	q.Set("from", req.Sender)
	q.Set("receiver", req.Destination)
	q.Set("slippage", fmt.Sprintf("%d", oneInchDefaultSlippage))
	q.Set("disableEstimate", "true")
	q.Set("allowPartialFill", "false")
	q.Set("compatibility", "true")
	httpReq.URL.RawQuery = q.Encode()

	if p.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call 1inch API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("1inch API error (status %d): %s", resp.StatusCode, string(body))
	}

	var swapResp oneInchSwapResponse
	if err := json.Unmarshal(body, &swapResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	value, ok := new(big.Int).SetString(swapResp.Tx.Value, 10)
	if !ok {
		value = big.NewInt(0)
	}

	// Get the approval spender address from resolver
	approvalSpender := swapResp.Tx.To
	if routerInfo, err := ResolveOneInchRouter(req.Quote.FromAsset.Chain); err == nil {
		approvalSpender = routerInfo.Address
	}

	// Check if approval is needed using consistent logic
	needsApproval := IsApprovalRequired(req.Quote.FromAsset)

	// Decode hex-encoded transaction data
	txData, err := decodeHexData(swapResp.Tx.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tx data: %w", err)
	}

	return &SwapResult{
		Provider:        p.Name(),
		TxData:          txData,
		Value:           value,
		ToAddress:       swapResp.Tx.To,
		ExpectedOut:     req.Quote.ExpectedOutput,
		NeedsApproval:   needsApproval,
		ApprovalAddress: approvalSpender,
		ApprovalAmount:  req.Quote.FromAmount,
	}, nil
}

// decodeHexData decodes a hex-encoded string (with optional 0x prefix) to bytes
func decodeHexData(hexStr string) ([]byte, error) {
	hexStr = strings.TrimPrefix(hexStr, "0x")
	if hexStr == "" {
		return nil, nil
	}
	return hex.DecodeString(hexStr)
}

// 1inch API response types
type oneInchSwapResponse struct {
	DstAmount string         `json:"dstAmount"`
	Tx        oneInchTxData  `json:"tx"`
}

type oneInchTxData struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Data     string `json:"data"`
	Value    string `json:"value"`
	Gas      int64  `json:"gas"`
	GasPrice string `json:"gasPrice"`
}

