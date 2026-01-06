package swap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"time"
)

const (
	jupiterDefaultAPIURL   = "https://api.jup.ag"
	jupiterDefaultSlippage = 100 // 1% in bps
	solanaNativeMint       = "So11111111111111111111111111111111111111112"
)

// Jupiter only supports Solana
var jupiterSupportedChains = []string{
	"Solana",
}

// JupiterProvider implements SwapProvider for Jupiter aggregator
type JupiterProvider struct {
	BaseProvider
	client *http.Client
	apiURL string
}

// NewJupiterProvider creates a new Jupiter provider
func NewJupiterProvider(apiURL string) *JupiterProvider {
	if apiURL == "" {
		apiURL = jupiterDefaultAPIURL
	}
	return &JupiterProvider{
		BaseProvider: NewBaseProvider("Jupiter", PriorityJupiter, jupiterSupportedChains),
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		apiURL: apiURL,
	}
}

// SupportsRoute checks if Jupiter can handle a swap between two assets
// Jupiter only supports Solana-to-Solana swaps
func (p *JupiterProvider) SupportsRoute(from, to Asset) bool {
	return from.Chain == "Solana" && to.Chain == "Solana"
}

// IsAvailable checks if Jupiter is available for a specific chain
func (p *JupiterProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	return chain == "Solana", nil
}

// GetStatus returns detailed availability status for a chain
func (p *JupiterProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	return &ProviderStatus{
		Chain:     chain,
		Available: chain == "Solana",
	}, nil
}

// GetQuote gets a swap quote from Jupiter
func (p *JupiterProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	if req.From.Chain != "Solana" || req.To.Chain != "Solana" {
		return nil, fmt.Errorf("Jupiter only supports Solana swaps")
	}

	// Format mint addresses
	inputMint := req.From.Address
	if inputMint == "" {
		inputMint = solanaNativeMint
	}
	outputMint := req.To.Address
	if outputMint == "" {
		outputMint = solanaNativeMint
	}

	// Build quote URL
	params := url.Values{}
	params.Set("swapMode", "ExactIn")
	params.Set("inputMint", inputMint)
	params.Set("outputMint", outputMint)
	params.Set("amount", req.Amount.String())
	params.Set("slippageBps", fmt.Sprintf("%d", jupiterDefaultSlippage))

	quoteURL := fmt.Sprintf("%s/swap/v1/quote?%s", p.apiURL, params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, quoteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Accept", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call Jupiter API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Jupiter API error (status %d): %s", resp.StatusCode, string(body))
	}

	var quoteResp jupiterQuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	outAmount, ok := new(big.Int).SetString(quoteResp.OutAmount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid outAmount: %s", quoteResp.OutAmount)
	}

	// Parse minimum output (otherAmountThreshold) for slippage protection
	var minOutput *big.Int
	if quoteResp.OtherAmountThreshold != "" {
		minOutput, _ = new(big.Int).SetString(quoteResp.OtherAmountThreshold, 10)
	}

	// Store the quote response for use in BuildTx (avoids re-fetching)
	providerData, err := json.Marshal(quoteResp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal quote data: %w", err)
	}

	return &Quote{
		Provider:       p.Name(),
		FromAsset:      req.From,
		ToAsset:        req.To,
		FromAmount:     req.Amount,
		ExpectedOutput: outAmount,
		MinimumOutput:  minOutput,
		ProviderData:   providerData,
	}, nil
}

// BuildTx builds an unsigned transaction for the swap
func (p *JupiterProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	// Use the stored quote data from GetQuote (avoids re-fetching)
	if len(req.Quote.ProviderData) == 0 {
		return nil, fmt.Errorf("quote missing provider data - was this quote obtained from GetQuote?")
	}

	var quoteResponse jupiterQuoteResponse
	if err := json.Unmarshal(req.Quote.ProviderData, &quoteResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stored quote: %w", err)
	}

	// Get swap instructions using the stored quote
	swapReqBody := jupiterSwapRequest{
		UserPublicKey:           req.Sender,
		QuoteResponse:           quoteResponse,
		WrapAndUnwrapSol:        true,
		UseSharedAccounts:       true,
		AsLegacyTransaction:     false,
		DynamicComputeUnitLimit: true,
	}

	swapJSON, err := json.Marshal(swapReqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal swap request: %w", err)
	}

	swapURL := fmt.Sprintf("%s/swap/v1/swap-instructions", p.apiURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, swapURL, bytes.NewReader(swapJSON))
	if err != nil {
		return nil, fmt.Errorf("failed to create swap request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get swap instructions: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Jupiter swap error (status %d): %s", resp.StatusCode, string(body))
	}

	return &SwapResult{
		Provider:    p.Name(),
		TxData:      body, // Raw instruction data for Solana tx building
		ExpectedOut: req.Quote.ExpectedOutput,
	}, nil
}

// Jupiter API types
type jupiterQuoteResponse struct {
	InputMint            string              `json:"inputMint"`
	InAmount             string              `json:"inAmount"`
	OutputMint           string              `json:"outputMint"`
	OutAmount            string              `json:"outAmount"`
	OtherAmountThreshold string              `json:"otherAmountThreshold"`
	SwapMode             string              `json:"swapMode"`
	SlippageBps          int                 `json:"slippageBps"`
	PriceImpactPct       string              `json:"priceImpactPct"`
	RoutePlan            []jupiterRoutePlan  `json:"routePlan"`
}

type jupiterRoutePlan struct {
	SwapInfo jupiterSwapInfo `json:"swapInfo"`
	Percent  int             `json:"percent"`
}

type jupiterSwapInfo struct {
	AmmKey     string `json:"ammKey"`
	Label      string `json:"label"`
	InputMint  string `json:"inputMint"`
	OutputMint string `json:"outputMint"`
	InAmount   string `json:"inAmount"`
	OutAmount  string `json:"outAmount"`
	FeeAmount  string `json:"feeAmount"`
	FeeMint    string `json:"feeMint"`
}

type jupiterSwapRequest struct {
	UserPublicKey           string               `json:"userPublicKey"`
	QuoteResponse           jupiterQuoteResponse `json:"quoteResponse"`
	WrapAndUnwrapSol        bool                 `json:"wrapAndUnwrapSol,omitempty"`
	UseSharedAccounts       bool                 `json:"useSharedAccounts,omitempty"`
	AsLegacyTransaction     bool                 `json:"asLegacyTransaction,omitempty"`
	DynamicComputeUnitLimit bool                 `json:"dynamicComputeUnitLimit,omitempty"`
}

