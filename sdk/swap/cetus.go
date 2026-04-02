package swap

import (
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
	cetusDefaultAPIURL   = "https://api-sui.cetus.zone"
	cetusDefaultSlippage = 50 // 0.5% in bps
	suiNativeCoinType    = "0x2::sui::SUI"
)

// Cetus only supports SUI
var cetusSupportedChains = []string{
	"Sui",
}

// CetusProvider implements SwapProvider for Cetus DEX aggregator on SUI.
type CetusProvider struct {
	BaseProvider
	client *http.Client
	apiURL string
}

// NewCetusProvider creates a new Cetus provider.
func NewCetusProvider(apiURL string) *CetusProvider {
	if apiURL == "" {
		apiURL = cetusDefaultAPIURL
	}
	return &CetusProvider{
		BaseProvider: NewBaseProvider("Cetus", PriorityCetus, cetusSupportedChains),
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		apiURL: apiURL,
	}
}

// SupportsRoute checks if Cetus can handle a swap — SUI-to-SUI only.
func (p *CetusProvider) SupportsRoute(from, to Asset) bool {
	return from.Chain == "Sui" && to.Chain == "Sui"
}

// IsAvailable checks if Cetus is available for a chain.
func (p *CetusProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	return chain == "Sui", nil
}

// GetStatus returns availability status.
func (p *CetusProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	return &ProviderStatus{
		Chain:     chain,
		Available: chain == "Sui",
	}, nil
}

// GetQuote gets a swap quote from Cetus aggregator.
func (p *CetusProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	if req.From.Chain != "Sui" || req.To.Chain != "Sui" {
		return nil, fmt.Errorf("Cetus only supports Sui swaps")
	}

	fromCoinType := req.From.Address
	if fromCoinType == "" {
		fromCoinType = suiNativeCoinType
	}
	toCoinType := req.To.Address
	if toCoinType == "" {
		toCoinType = suiNativeCoinType
	}

	// Build route query (v param required by Cetus API)
	params := url.Values{}
	params.Set("from", fromCoinType)
	params.Set("target", toCoinType)
	params.Set("amount", req.Amount.String())
	params.Set("by_amount_in", "true")
	params.Set("v", "1010405") // SDK version required by API

	routeURL := fmt.Sprintf("%s/router_v3/find_routes?%s", p.apiURL, params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, routeURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Accept", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call Cetus API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Cetus API error (status %d): %s", resp.StatusCode, string(body))
	}

	var routeResp cetusRouteResponse
	if err := json.Unmarshal(body, &routeResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if routeResp.Code != 200 || routeResp.Data == nil {
		msg := routeResp.Msg
		if msg == "" {
			msg = "no route found"
		}
		return nil, fmt.Errorf("Cetus routing failed: %s", msg)
	}

	outAmount := new(big.Int).SetUint64(routeResp.Data.AmountOut)
	if outAmount.Sign() == 0 {
		return nil, fmt.Errorf("Cetus returned zero output for this swap")
	}

	if routeResp.Data.InsufficientLiquidity {
		return nil, fmt.Errorf("insufficient liquidity for this swap pair")
	}

	providerData, err := json.Marshal(routeResp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal route data: %w", err)
	}

	return &Quote{
		Provider:       p.Name(),
		FromAsset:      req.From,
		ToAsset:        req.To,
		FromAmount:     req.Amount,
		ExpectedOutput: outAmount,
		ProviderData:   providerData,
	}, nil
}

// BuildTx builds an unsigned SUI transaction for the Cetus swap.
// For now, returns the route data — the MCP layer handles SUI TX construction.
func (p *CetusProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	// The route data from GetQuote contains the swap instructions
	// The MCP/app layer will construct the actual SUI transaction
	return &SwapResult{
		Provider:    p.Name(),
		TxData:      req.Quote.ProviderData,
		ExpectedOut: req.Quote.ExpectedOutput,
	}, nil
}

// --- Cetus API types ---

type cetusRouteResponse struct {
	Code int             `json:"code"`
	Msg  string          `json:"msg"`
	Data *cetusRouteData `json:"data"`
}

type cetusRouteData struct {
	AmountOut             uint64  `json:"amount_out"`
	AmountIn              uint64  `json:"amount_in"`
	InsufficientLiquidity bool    `json:"insufficient_liquidity"`
	RequestID             string  `json:"request_id"`
}
