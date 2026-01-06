package swap

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// THORChain network identifiers
const (
	thorChainBTC  = "BTC"
	thorChainLTC  = "LTC"
	thorChainDOGE = "DOGE"
	thorChainBCH  = "BCH"
	thorChainETH  = "ETH"
	thorChainBSC  = "BSC"
	thorChainBASE = "BASE"
	thorChainAVAX = "AVAX"
	thorChainXRP  = "XRP"
	thorChainGAIA = "GAIA"
	thorChainTHOR = "THOR"
)

// THORChain supported chains mapped to their network identifiers
var thorChainNetworks = map[string]string{
	"Bitcoin":     thorChainBTC,
	"Litecoin":    thorChainLTC,
	"Dogecoin":    thorChainDOGE,
	"BitcoinCash": thorChainBCH,
	"Ethereum":    thorChainETH,
	"BSC":         thorChainBSC,
	"Base":        thorChainBASE,
	"Avalanche":   thorChainAVAX,
	"XRP":         thorChainXRP,
	"Cosmos":      thorChainGAIA,
	"THORChain":   thorChainTHOR,
}

// Default THORChain endpoints
var defaultTHORChainEndpoints = []string{
	"https://thornode.ninerealms.com",
	"https://thornode.thorchain.info",
}

// THORChainProvider implements SwapProvider for THORChain
type THORChainProvider struct {
	BaseProvider
	client    *http.Client
	endpoints []string
}

// NewTHORChainProvider creates a new THORChain provider
func NewTHORChainProvider(endpoints []string) *THORChainProvider {
	if len(endpoints) == 0 {
		endpoints = defaultTHORChainEndpoints
	}

	chains := make([]string, 0, len(thorChainNetworks))
	for chain := range thorChainNetworks {
		chains = append(chains, chain)
	}

	return &THORChainProvider{
		BaseProvider: NewBaseProvider("THORChain", PriorityTHORChain, chains),
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		endpoints: endpoints,
	}
}

// SupportsRoute checks if THORChain can handle a swap between two assets
func (p *THORChainProvider) SupportsRoute(from, to Asset) bool {
	_, fromOk := thorChainNetworks[from.Chain]
	_, toOk := thorChainNetworks[to.Chain]
	return fromOk && toOk
}

// IsAvailable checks if THORChain is available for a specific chain
func (p *THORChainProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	status, err := p.GetStatus(ctx, chain)
	if err != nil {
		return false, err
	}
	return status.Available, nil
}

// GetStatus returns detailed availability status for a chain
func (p *THORChainProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	network, ok := thorChainNetworks[chain]
	if !ok {
		return &ProviderStatus{
			Chain:     chain,
			Available: false,
		}, nil
	}

	addresses, err := p.getInboundAddresses(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get inbound addresses: %w", err)
	}

	for _, addr := range addresses {
		if strings.EqualFold(addr.Chain, network) {
			available := !addr.Halted && !addr.GlobalTradingPaused && !addr.ChainTradingPaused
			return &ProviderStatus{
				Chain:               chain,
				Available:           available,
				Halted:              addr.Halted,
				GlobalTradingPaused: addr.GlobalTradingPaused,
				ChainTradingPaused:  addr.ChainTradingPaused,
				Router:              addr.Router,
				InboundAddress:      addr.Address,
			}, nil
		}
	}

	return &ProviderStatus{
		Chain:     chain,
		Available: false,
	}, nil
}

// GetQuote gets a swap quote from THORChain
func (p *THORChainProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	fromAsset, err := p.formatAsset(req.From)
	if err != nil {
		return nil, fmt.Errorf("invalid from asset: %w", err)
	}

	toAsset, err := p.formatAsset(req.To)
	if err != nil {
		return nil, fmt.Errorf("invalid to asset: %w", err)
	}

	params := url.Values{}
	params.Set("from_asset", fromAsset)
	params.Set("to_asset", toAsset)
	params.Set("amount", req.Amount.String())
	params.Set("destination", req.Destination)
	params.Set("streaming_interval", "3")
	params.Set("streaming_quantity", "0")
	params.Set("tolerance_bps", "2500")

	quoteURL := fmt.Sprintf("%s/thorchain/quote/swap?%s", p.endpoints[0], params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, quoteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get quote: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp thorChainErrorResponse
		if json.Unmarshal(body, &errResp) == nil && errResp.Error != "" {
			return nil, fmt.Errorf("THORChain quote error: %s", errResp.Error)
		}
		return nil, fmt.Errorf("THORChain quote failed with status %d: %s", resp.StatusCode, string(body))
	}

	var quoteResp thorChainQuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, fmt.Errorf("failed to parse quote response: %w", err)
	}

	expectedOutput, ok := new(big.Int).SetString(quoteResp.ExpectedAmountOut, 10)
	if !ok {
		return nil, fmt.Errorf("invalid expected_amount_out: %s", quoteResp.ExpectedAmountOut)
	}

	// Determine router address - try resolver first, fallback to quote response
	routerAddress := quoteResp.Router
	if routerInfo, err := ResolveTHORChainRouter(req.From.Chain); err == nil && routerInfo.Address != "" {
		routerAddress = routerInfo.Address
	}

	// Check if approval is needed (ERC20 token on EVM chain)
	needsApproval := IsApprovalRequired(req.From)

	return &Quote{
		Provider:        p.Name(),
		FromAsset:       req.From,
		ToAsset:         req.To,
		FromAmount:      req.Amount,
		ExpectedOutput:  expectedOutput,
		Memo:            quoteResp.Memo,
		InboundAddress:  quoteResp.InboundAddress,
		Router:          routerAddress,
		Expiry:          quoteResp.Expiry,
		StreamingSwap:   quoteResp.MaxStreamingQuantity > 0,
		StreamingBlocks: quoteResp.StreamingSwapBlocks,
		NeedsApproval:   needsApproval,
		ApprovalSpender: routerAddress,
		ApprovalAmount:  req.Amount,
	}, nil
}

// BuildTx builds an unsigned transaction for the swap
func (p *THORChainProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	result := &SwapResult{
		Provider:    p.Name(),
		ToAddress:   req.Quote.InboundAddress,
		Memo:        req.Quote.Memo,
		ExpectedOut: req.Quote.ExpectedOutput,
	}

	// For EVM token swaps, set approval info using the router
	if req.Quote.FromAsset.Address != "" && req.Quote.Router != "" {
		result.NeedsApproval = true
		result.ApprovalAddress = req.Quote.Router
		result.ApprovalAmount = req.Quote.FromAmount
	}

	return result, nil
}

// formatAsset formats an asset for THORChain API
func (p *THORChainProvider) formatAsset(asset Asset) (string, error) {
	network, ok := thorChainNetworks[asset.Chain]
	if !ok {
		return "", fmt.Errorf("unsupported chain: %s", asset.Chain)
	}

	// Native asset format: CHAIN.SYMBOL (e.g., BTC.BTC, ETH.ETH)
	if asset.Address == "" {
		return fmt.Sprintf("%s.%s", network, asset.Symbol), nil
	}

	// Token format: CHAIN.SYMBOL-ADDRESS (e.g., ETH.USDC-0x...)
	return fmt.Sprintf("%s.%s-%s", network, asset.Symbol, strings.ToUpper(asset.Address)), nil
}

// getInboundAddresses fetches inbound addresses from THORChain
func (p *THORChainProvider) getInboundAddresses(ctx context.Context) ([]thorChainInboundAddress, error) {
	var lastErr error

	for _, endpoint := range p.endpoints {
		addresses, err := p.fetchInboundAddresses(ctx, endpoint)
		if err != nil {
			lastErr = err
			continue
		}
		return addresses, nil
	}

	return nil, fmt.Errorf("all endpoints failed, last error: %w", lastErr)
}

func (p *THORChainProvider) fetchInboundAddresses(ctx context.Context, endpoint string) ([]thorChainInboundAddress, error) {
	reqURL := endpoint + "/thorchain/inbound_addresses"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var addresses []thorChainInboundAddress
	if err := json.Unmarshal(body, &addresses); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return addresses, nil
}

// THORChain API response types
type thorChainInboundAddress struct {
	Chain               string `json:"chain"`
	PubKey              string `json:"pub_key"`
	Address             string `json:"address"`
	Router              string `json:"router"`
	Halted              bool   `json:"halted"`
	GlobalTradingPaused bool   `json:"global_trading_paused"`
	ChainTradingPaused  bool   `json:"chain_trading_paused"`
	GasRate             string `json:"gas_rate"`
	DustThreshold       string `json:"dust_threshold"`
}

type thorChainQuoteResponse struct {
	InboundAddress       string `json:"inbound_address"`
	Router               string `json:"router"`
	Expiry               int64  `json:"expiry"`
	Memo                 string `json:"memo"`
	ExpectedAmountOut    string `json:"expected_amount_out"`
	DustThreshold        string `json:"dust_threshold"`
	MaxStreamingQuantity int64  `json:"max_streaming_quantity"`
	StreamingSwapBlocks  int64  `json:"streaming_swap_blocks"`
}

type thorChainErrorResponse struct {
	Error string `json:"error"`
}

