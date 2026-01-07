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
	thorChainTRON = "TRX"
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
	"Tron":        thorChainTRON,
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

	// Try all endpoints with fallback
	var lastErr error
	for _, endpoint := range p.endpoints {
		quote, err := p.fetchQuote(ctx, endpoint, params, req)
		if err == nil {
			return quote, nil
		}
		lastErr = err
	}

	return nil, fmt.Errorf("all THORChain endpoints failed: %w", lastErr)
}

// fetchQuote fetches a quote from a specific THORChain endpoint
func (p *THORChainProvider) fetchQuote(ctx context.Context, endpoint string, params url.Values, req QuoteRequest) (*Quote, error) {
	quoteURL := fmt.Sprintf("%s/thorchain/quote/swap?%s", endpoint, params.Encode())

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

	// Check if this is an EVM chain that requires router calldata
	if IsEVMChain(req.Quote.FromAsset.Chain) && req.Quote.Router != "" {
		// Build router calldata for EVM chains
		txData, value, err := p.encodeRouterDeposit(req)
		if err != nil {
			return nil, fmt.Errorf("failed to encode router deposit: %w", err)
		}
		result.TxData = txData
		result.Value = value
		result.ToAddress = req.Quote.Router

		// For EVM token swaps, set approval info
		if req.Quote.FromAsset.Address != "" {
			result.NeedsApproval = true
			result.ApprovalAddress = req.Quote.Router
			result.ApprovalAmount = req.Quote.FromAmount
		}
	} else {
		// For UTXO/Cosmos chains, just set the value to send
		result.Value = req.Quote.FromAmount
	}

	return result, nil
}

// THORChain router function selectors
// depositWithExpiry(address,address,uint256,string,uint256): 0x44bc937b
var thorChainDepositWithExpirySig = []byte{0x44, 0xbc, 0x93, 0x7b}

// encodeRouterDeposit encodes a THORChain router depositWithExpiry call
// depositWithExpiry(address vault, address asset, uint256 amount, string memo, uint256 expiration)
func (p *THORChainProvider) encodeRouterDeposit(req SwapRequest) ([]byte, *big.Int, error) {
	vault := req.Quote.InboundAddress
	asset := req.Quote.FromAsset.Address
	amount := req.Quote.FromAmount
	memo := req.Quote.Memo
	expiry := req.Quote.Expiry

	// For native ETH, asset is the zero address
	if asset == "" {
		asset = "0x0000000000000000000000000000000000000000"
	}

	// Default expiry to 15 minutes from now if not set
	if expiry == 0 {
		expiry = time.Now().Unix() + 900
	}

	// Encode calldata
	calldata, err := encodeDepositWithExpiry(vault, asset, amount, memo, big.NewInt(expiry))
	if err != nil {
		return nil, nil, err
	}

	// Value is the amount for native token, 0 for ERC20
	value := big.NewInt(0)
	if req.Quote.FromAsset.Address == "" {
		value = amount
	}

	return calldata, value, nil
}

// encodeDepositWithExpiry encodes the THORChain router depositWithExpiry function call
// Function signature: depositWithExpiry(address,address,uint256,string,uint256)
func encodeDepositWithExpiry(vault, asset string, amount *big.Int, memo string, expiry *big.Int) ([]byte, error) {
	// Validate addresses
	vaultBytes, err := addressToBytes(vault)
	if err != nil {
		return nil, fmt.Errorf("invalid vault address: %w", err)
	}
	assetBytes, err := addressToBytes(asset)
	if err != nil {
		return nil, fmt.Errorf("invalid asset address: %w", err)
	}

	// Calculate memo byte length (padded to 32 bytes)
	memoBytes := []byte(memo)
	memoPaddedLen := ((len(memoBytes) + 31) / 32) * 32

	// Calculate total calldata size:
	// 4 (selector) + 32*5 (5 params) + 32 (memo offset points to) + memoPaddedLen
	// Parameters: vault(32) + asset(32) + amount(32) + memo_offset(32) + expiry(32)
	// Dynamic data: memo_length(32) + memo_data(memoPaddedLen)
	totalLen := 4 + 32*5 + 32 + memoPaddedLen

	calldata := make([]byte, totalLen)

	// Function selector
	copy(calldata[0:4], thorChainDepositWithExpirySig)
	offset := 4

	// Vault address (32 bytes, left-padded)
	copy(calldata[offset+12:offset+32], vaultBytes)
	offset += 32

	// Asset address (32 bytes, left-padded)
	copy(calldata[offset+12:offset+32], assetBytes)
	offset += 32

	// Amount (32 bytes, left-padded)
	amountBytes := amount.Bytes()
	copy(calldata[offset+32-len(amountBytes):offset+32], amountBytes)
	offset += 32

	// Memo offset (points to dynamic data section: 5*32 = 160)
	memoOffset := big.NewInt(160)
	memoOffsetBytes := memoOffset.Bytes()
	copy(calldata[offset+32-len(memoOffsetBytes):offset+32], memoOffsetBytes)
	offset += 32

	// Expiry (32 bytes, left-padded)
	expiryBytes := expiry.Bytes()
	copy(calldata[offset+32-len(expiryBytes):offset+32], expiryBytes)
	offset += 32

	// Dynamic data: memo length
	memoLenBytes := big.NewInt(int64(len(memoBytes))).Bytes()
	copy(calldata[offset+32-len(memoLenBytes):offset+32], memoLenBytes)
	offset += 32

	// Dynamic data: memo content (right-padded)
	copy(calldata[offset:offset+len(memoBytes)], memoBytes)

	return calldata, nil
}

// addressToBytes converts a hex address string to 20 bytes
func addressToBytes(addr string) ([]byte, error) {
	addr = strings.TrimPrefix(strings.ToLower(addr), "0x")
	if len(addr) != 40 {
		return nil, fmt.Errorf("invalid address length: %d", len(addr))
	}
	return hex.DecodeString(addr)
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
	// Use address as-is to preserve EIP-55 checksums for EVM chains
	return fmt.Sprintf("%s.%s-%s", network, asset.Symbol, asset.Address), nil
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

