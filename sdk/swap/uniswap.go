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
	uniswapDefaultBaseURL = "https://api.uniswap.org/v2"
	uniswapNativeAddress  = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
)

// Uniswap chain IDs
var uniswapChainIDs = map[string]int{
	"Ethereum": 1,
	"Polygon":  137,
	"Arbitrum": 42161,
	"Optimism": 10,
	"Base":     8453,
}

// Uniswap supported EVM chains
var uniswapSupportedChains = []string{
	"Ethereum",
	"Polygon",
	"Arbitrum",
	"Optimism",
	"Base",
}

// UniswapProvider implements SwapProvider for Uniswap
type UniswapProvider struct {
	BaseProvider
	client  *http.Client
	baseURL string
}

// NewUniswapProvider creates a new Uniswap provider
func NewUniswapProvider() *UniswapProvider {
	return &UniswapProvider{
		BaseProvider: NewBaseProvider("Uniswap", PriorityUniswap, uniswapSupportedChains),
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		baseURL: uniswapDefaultBaseURL,
	}
}

// SupportsRoute checks if Uniswap can handle a swap between two assets
// Uniswap only supports same-chain swaps
func (p *UniswapProvider) SupportsRoute(from, to Asset) bool {
	return from.Chain == to.Chain && p.SupportsChain(from.Chain)
}

// IsAvailable checks if Uniswap is available for a specific chain
func (p *UniswapProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	return p.SupportsChain(chain), nil
}

// GetStatus returns detailed availability status for a chain
func (p *UniswapProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	return &ProviderStatus{
		Chain:     chain,
		Available: p.SupportsChain(chain),
	}, nil
}

// GetQuote gets a swap quote from Uniswap
func (p *UniswapProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	chainID, ok := uniswapChainIDs[req.From.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported chain: %s", req.From.Chain)
	}

	// Format token addresses
	tokenIn := req.From.Address
	if tokenIn == "" {
		tokenIn = uniswapNativeAddress
	}
	tokenOut := req.To.Address
	if tokenOut == "" {
		tokenOut = uniswapNativeAddress
	}

	// Build quote URL
	params := url.Values{}
	params.Set("tokenIn", strings.ToLower(tokenIn))
	params.Set("tokenOut", strings.ToLower(tokenOut))
	params.Set("amount", req.Amount.String())
	params.Set("type", "EXACT_INPUT")
	params.Set("chainId", fmt.Sprintf("%d", chainID))
	params.Set("recipient", req.Destination)
	params.Set("enableUniversalRouter", "true")

	quoteURL := fmt.Sprintf("%s/quote?%s", p.baseURL, params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, quoteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Accept", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call Uniswap API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Uniswap API error (status %d): %s", resp.StatusCode, string(body))
	}

	var quoteResp uniswapQuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	quoteAmount, ok := new(big.Int).SetString(quoteResp.Quote, 10)
	if !ok {
		return nil, fmt.Errorf("invalid quote amount: %s", quoteResp.Quote)
	}

	// Get router address for approval
	routerAddress := quoteResp.MethodParameters.To
	if routerAddress == "" {
		if routerInfo, err := ResolveUniswapRouter(req.From.Chain); err == nil && routerInfo.Address != "" {
			routerAddress = routerInfo.Address
		} else {
			routerAddress = "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD" // Fallback Universal Router
		}
	}

	// Check if approval is needed (ERC20 token)
	needsApproval := IsApprovalRequired(req.From)

	return &Quote{
		Provider:        p.Name(),
		FromAsset:       req.From,
		ToAsset:         req.To,
		FromAmount:      req.Amount,
		ExpectedOutput:  quoteAmount,
		Router:          routerAddress,
		NeedsApproval:   needsApproval,
		ApprovalSpender: routerAddress,
		ApprovalAmount:  req.Amount,
	}, nil
}

// BuildTx builds an unsigned transaction for the swap
func (p *UniswapProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	chainID, ok := uniswapChainIDs[req.Quote.FromAsset.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported chain: %s", req.Quote.FromAsset.Chain)
	}

	tokenIn := req.Quote.FromAsset.Address
	if tokenIn == "" {
		tokenIn = uniswapNativeAddress
	}
	tokenOut := req.Quote.ToAsset.Address
	if tokenOut == "" {
		tokenOut = uniswapNativeAddress
	}

	// Build quote URL with swap request
	params := url.Values{}
	params.Set("tokenIn", strings.ToLower(tokenIn))
	params.Set("tokenOut", strings.ToLower(tokenOut))
	params.Set("amount", req.Quote.FromAmount.String())
	params.Set("type", "EXACT_INPUT")
	params.Set("chainId", fmt.Sprintf("%d", chainID))
	params.Set("recipient", req.Destination)
	params.Set("enableUniversalRouter", "true")

	quoteURL := fmt.Sprintf("%s/quote?%s", p.baseURL, params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, quoteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Accept", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call Uniswap API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Uniswap API error (status %d): %s", resp.StatusCode, string(body))
	}

	var quoteResp uniswapQuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	value := big.NewInt(0)
	if strings.EqualFold(tokenIn, uniswapNativeAddress) {
		value = req.Quote.FromAmount
	}

	// Get the router address from resolver
	toAddress := quoteResp.MethodParameters.To
	if toAddress == "" {
		if routerInfo, err := ResolveUniswapRouter(req.Quote.FromAsset.Chain); err == nil {
			toAddress = routerInfo.Address
		} else {
			toAddress = "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD" // Fallback Universal Router
		}
	}

	// Check if approval is needed using consistent logic
	needsApproval := IsApprovalRequired(req.Quote.FromAsset)
	approvalSpender := toAddress

	// Decode hex-encoded calldata
	txData, err := decodeUniswapHexData(quoteResp.MethodParameters.Calldata)
	if err != nil {
		return nil, fmt.Errorf("failed to decode calldata: %w", err)
	}

	return &SwapResult{
		Provider:        p.Name(),
		TxData:          txData,
		Value:           value,
		ToAddress:       toAddress,
		ExpectedOut:     req.Quote.ExpectedOutput,
		NeedsApproval:   needsApproval,
		ApprovalAddress: approvalSpender,
		ApprovalAmount:  req.Quote.FromAmount,
	}, nil
}

// decodeUniswapHexData decodes a hex-encoded string (with optional 0x prefix) to bytes
func decodeUniswapHexData(hexStr string) ([]byte, error) {
	hexStr = strings.TrimPrefix(hexStr, "0x")
	if hexStr == "" {
		return nil, nil
	}
	return hex.DecodeString(hexStr)
}

// Uniswap API types
type uniswapQuoteResponse struct {
	Quote            string                  `json:"quote"`
	QuoteGasAdjusted string                  `json:"quoteGasAdjusted"`
	GasUseEstimate   string                  `json:"gasUseEstimate"`
	MethodParameters uniswapMethodParameters `json:"methodParameters"`
}

type uniswapMethodParameters struct {
	Calldata string `json:"calldata"`
	Value    string `json:"value"`
	To       string `json:"to"`
}
