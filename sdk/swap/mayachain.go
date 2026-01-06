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

// Mayachain network identifiers
const (
	mayaChainBTC  = "BTC"
	mayaChainETH  = "ETH"
	mayaChainARB  = "ARB"
	mayaChainZEC  = "ZEC"
	mayaChainDASH = "DASH"
	mayaChainTHOR = "THOR"
	mayaChainMAYA = "MAYA"
	mayaChainKUJI = "KUJI"
)

// Mayachain supported chains mapped to their network identifiers
var mayaChainNetworks = map[string]string{
	"Bitcoin":   mayaChainBTC,
	"Ethereum":  mayaChainETH,
	"Arbitrum":  mayaChainARB,
	"Zcash":     mayaChainZEC,
	"Dash":      mayaChainDASH,
	"THORChain": mayaChainTHOR,
	"MayaChain": mayaChainMAYA,
	"Kujira":    mayaChainKUJI,
}

// Default Mayachain endpoints
var defaultMayaChainEndpoints = []string{
	"https://mayanode.mayachain.info",
	"https://maya-api.polkachu.com",
}

// MayachainProvider implements SwapProvider for Mayachain
type MayachainProvider struct {
	BaseProvider
	client    *http.Client
	endpoints []string
}

// NewMayachainProvider creates a new Mayachain provider
func NewMayachainProvider(endpoints []string) *MayachainProvider {
	if len(endpoints) == 0 {
		endpoints = defaultMayaChainEndpoints
	}

	chains := make([]string, 0, len(mayaChainNetworks))
	for chain := range mayaChainNetworks {
		chains = append(chains, chain)
	}

	return &MayachainProvider{
		BaseProvider: NewBaseProvider("Mayachain", PriorityMayachain, chains),
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		endpoints: endpoints,
	}
}

// SupportsRoute checks if Mayachain can handle a swap between two assets
func (p *MayachainProvider) SupportsRoute(from, to Asset) bool {
	_, fromOk := mayaChainNetworks[from.Chain]
	_, toOk := mayaChainNetworks[to.Chain]
	return fromOk && toOk
}

// IsAvailable checks if Mayachain is available for a specific chain
func (p *MayachainProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	status, err := p.GetStatus(ctx, chain)
	if err != nil {
		return false, err
	}
	return status.Available, nil
}

// GetStatus returns detailed availability status for a chain
func (p *MayachainProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	network, ok := mayaChainNetworks[chain]
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

// GetQuote gets a swap quote from Mayachain
func (p *MayachainProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
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

	quoteURL := fmt.Sprintf("%s/mayachain/quote/swap?%s", p.endpoints[0], params.Encode())

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
		var errResp mayaChainErrorResponse
		if json.Unmarshal(body, &errResp) == nil && errResp.Error != "" {
			return nil, fmt.Errorf("Mayachain quote error: %s", errResp.Error)
		}
		return nil, fmt.Errorf("Mayachain quote failed with status %d: %s", resp.StatusCode, string(body))
	}

	var quoteResp mayaChainQuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, fmt.Errorf("failed to parse quote response: %w", err)
	}

	expectedOutput, ok := new(big.Int).SetString(quoteResp.ExpectedAmountOut, 10)
	if !ok {
		return nil, fmt.Errorf("invalid expected_amount_out: %s", quoteResp.ExpectedAmountOut)
	}

	return &Quote{
		Provider:        p.Name(),
		FromAsset:       req.From,
		ToAsset:         req.To,
		FromAmount:      req.Amount,
		ExpectedOutput:  expectedOutput,
		Memo:            quoteResp.Memo,
		InboundAddress:  quoteResp.InboundAddress,
		Router:          quoteResp.Router,
		Expiry:          quoteResp.Expiry,
		StreamingSwap:   quoteResp.MaxStreamingQuantity > 0,
		StreamingBlocks: quoteResp.StreamingSwapBlocks,
	}, nil
}

// BuildTx builds an unsigned transaction for the swap
func (p *MayachainProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	return &SwapResult{
		Provider:    p.Name(),
		ToAddress:   req.Quote.InboundAddress,
		Memo:        req.Quote.Memo,
		ExpectedOut: req.Quote.ExpectedOutput,
	}, nil
}

// formatAsset formats an asset for Mayachain API
func (p *MayachainProvider) formatAsset(asset Asset) (string, error) {
	network, ok := mayaChainNetworks[asset.Chain]
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

// getInboundAddresses fetches inbound addresses from Mayachain
func (p *MayachainProvider) getInboundAddresses(ctx context.Context) ([]mayaChainInboundAddress, error) {
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

func (p *MayachainProvider) fetchInboundAddresses(ctx context.Context, endpoint string) ([]mayaChainInboundAddress, error) {
	reqURL := endpoint + "/mayachain/inbound_addresses"

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

	var addresses []mayaChainInboundAddress
	if err := json.Unmarshal(body, &addresses); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return addresses, nil
}

// Mayachain API response types
type mayaChainInboundAddress struct {
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

type mayaChainQuoteResponse struct {
	InboundAddress       string `json:"inbound_address"`
	Router               string `json:"router"`
	Expiry               int64  `json:"expiry"`
	Memo                 string `json:"memo"`
	ExpectedAmountOut    string `json:"expected_amount_out"`
	DustThreshold        string `json:"dust_threshold"`
	MaxStreamingQuantity int64  `json:"max_streaming_quantity"`
	StreamingSwapBlocks  int64  `json:"streaming_swap_blocks"`
}

type mayaChainErrorResponse struct {
	Error string `json:"error"`
}

