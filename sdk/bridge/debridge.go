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
	debridgeDefaultBaseURL = "https://dln.debridge.finance/v1.0"
)

// ============================================================================
// deBridge Contract Addresses (DLN Protocol)
// ============================================================================
// Source: https://docs.debridge.finance/dln-the-debridge-liquidity-network/deployed-contracts
// VERIFIED: These are the DlnSource contracts for initiating cross-chain orders
// ============================================================================

var debridgeDlnSourceAddresses = map[string]string{
	// EVM Chains - DlnSource contracts.
	// NOTE: DLN API returns this as tx.allowanceTarget and create-tx.tx.to.
	"Ethereum":  "0xeF4fB24aD0916217251F553c0596F8Edc630EB66",
	"Optimism":  "0xeF4fB24aD0916217251F553c0596F8Edc630EB66",
	"BSC":       "0xeF4fB24aD0916217251F553c0596F8Edc630EB66",
	"Polygon":   "0xeF4fB24aD0916217251F553c0596F8Edc630EB66",
	"Base":      "0xeF4fB24aD0916217251F553c0596F8Edc630EB66",
	"Arbitrum":  "0xeF4fB24aD0916217251F553c0596F8Edc630EB66",
	"Avalanche": "0xeF4fB24aD0916217251F553c0596F8Edc630EB66",
	"Linea":     "0xeF4fB24aD0916217251F553c0596F8Edc630EB66",
}

// deBridge chain IDs for API calls
var debridgeChainIDs = map[string]int{
	// Verified by DLN API validation errors (dstChainId allowlist).
	"Ethereum":  1,
	"Optimism":  10,
	"BSC":       56,
	"Polygon":   137,
	"Base":      8453,
	"Arbitrum":  42161,
	"Avalanche": 43114,
	"Linea":     59144,
}

// deBridge supported chains
var debridgeSupportedChains = []string{
	// EVM-only: this provider currently builds EVM transactions.
	"Ethereum",
	"Optimism",
	"BSC",
	"Polygon",
	"Base",
	"Arbitrum",
	"Avalanche",
	"Linea",
}

// DeBridgeProvider implements BridgeProvider for deBridge (DLN protocol)
type DeBridgeProvider struct {
	BaseProvider
	client  *http.Client
	baseURL string
}

// NewDeBridgeProvider creates a new deBridge provider
func NewDeBridgeProvider() *DeBridgeProvider {
	return &DeBridgeProvider{
		BaseProvider: NewBaseProvider("deBridge", PriorityDeBridge, debridgeSupportedChains),
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: debridgeDefaultBaseURL,
	}
}

// SupportsRoute checks if deBridge can handle a bridge between two chains
func (p *DeBridgeProvider) SupportsRoute(from, to BridgeAsset) bool {
	// Must be different chains
	if from.Chain == to.Chain {
		return false
	}
	// Both chains must be supported
	return p.SupportsChain(from.Chain) && p.SupportsChain(to.Chain)
}

// IsAvailable checks if deBridge is available for a specific chain
func (p *DeBridgeProvider) IsAvailable(_ context.Context, chain string) (bool, error) {
	return p.SupportsChain(chain), nil
}

// GetStatus returns detailed availability status for a chain
func (p *DeBridgeProvider) GetStatus(_ context.Context, chain string) (*ProviderStatus, error) {
	bridgeAddr := debridgeDlnSourceAddresses[chain]
	return &ProviderStatus{
		Chain:         chain,
		Available:     p.SupportsChain(chain),
		BridgeAddress: bridgeAddr,
	}, nil
}

// GetQuote gets a bridge quote from deBridge
func (p *DeBridgeProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	srcChainID, ok := debridgeChainIDs[req.From.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported from chain: %s", req.From.Chain)
	}
	dstChainID, ok := debridgeChainIDs[req.To.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported to chain: %s", req.To.Chain)
	}

	// Format token addresses (use 0x0 for native tokens)
	srcToken := req.From.Address
	if srcToken == "" {
		srcToken = "0x0000000000000000000000000000000000000000"
	}
	dstToken := req.To.Address
	if dstToken == "" {
		dstToken = "0x0000000000000000000000000000000000000000"
	}

	// Build quote URL
	params := url.Values{}
	params.Set("srcChainId", fmt.Sprintf("%d", srcChainID))
	params.Set("srcChainTokenIn", srcToken)
	params.Set("srcChainTokenInAmount", req.Amount.String())
	params.Set("dstChainId", fmt.Sprintf("%d", dstChainID))
	params.Set("dstChainTokenOut", dstToken)
	params.Set("dstChainTokenOutRecipient", req.Destination)
	params.Set("prependOperatingExpenses", "true")

	quoteURL := fmt.Sprintf("%s/dln/order/quote?%s", p.baseURL, params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, quoteURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Accept", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call deBridge API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp debridgeErrorResponse
		if json.Unmarshal(body, &errResp) == nil && errResp.Error != "" {
			return nil, fmt.Errorf("deBridge error: %s", errResp.Error)
		}
		return nil, fmt.Errorf("deBridge API error (status %d): %s", resp.StatusCode, string(body))
	}

	var quoteResp debridgeQuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Parse output amount
	toAmount, ok := new(big.Int).SetString(quoteResp.Estimation.DstChainTokenOut.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid output amount: %s", quoteResp.Estimation.DstChainTokenOut.Amount)
	}

	// Get bridge contract address from API (allowance target / spender).
	router := strings.TrimSpace(quoteResp.Tx.AllowanceTarget)
	if router == "" {
		return nil, fmt.Errorf("deBridge response missing tx.allowanceTarget")
	}

	// CRITICAL: Verify the bridge address
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
func (p *DeBridgeProvider) BuildTx(ctx context.Context, req BridgeRequest) (*BridgeResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	srcChainID, ok := debridgeChainIDs[req.Quote.FromAsset.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported from chain: %s", req.Quote.FromAsset.Chain)
	}
	dstChainID, ok := debridgeChainIDs[req.Quote.ToAsset.Chain]
	if !ok {
		return nil, fmt.Errorf("unsupported to chain: %s", req.Quote.ToAsset.Chain)
	}

	srcToken := req.Quote.FromAsset.Address
	if srcToken == "" {
		srcToken = "0x0000000000000000000000000000000000000000"
	}
	dstToken := req.Quote.ToAsset.Address
	if dstToken == "" {
		dstToken = "0x0000000000000000000000000000000000000000"
	}

	// Create order request - deBridge uses a different endpoint for tx building
	params := url.Values{}
	params.Set("srcChainId", fmt.Sprintf("%d", srcChainID))
	params.Set("srcChainTokenIn", srcToken)
	params.Set("srcChainTokenInAmount", req.Quote.FromAmount.String())
	params.Set("dstChainId", fmt.Sprintf("%d", dstChainID))
	params.Set("dstChainTokenOut", dstToken)
	params.Set("dstChainTokenOutRecipient", req.Destination)
	params.Set("srcChainOrderAuthorityAddress", req.Sender)
	params.Set("dstChainOrderAuthorityAddress", req.Destination)
	params.Set("prependOperatingExpenses", "true")

	createURL := fmt.Sprintf("%s/dln/order/create-tx?%s", p.baseURL, params.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, createURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Accept", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call deBridge API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("deBridge API error (status %d): %s", resp.StatusCode, string(body))
	}

	var createResp debridgeCreateTxResponse
	if err := json.Unmarshal(body, &createResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if createResp.Tx == nil {
		return nil, fmt.Errorf("no transaction in deBridge response")
	}

	// CRITICAL: Triple-verify the bridge address before building transaction
	toAddress := createResp.Tx.To
	if err := p.verifyBridgeAddress(req.Quote.FromAsset.Chain, toAddress); err != nil {
		return nil, fmt.Errorf("SECURITY: bridge address verification failed: %w", err)
	}

	value := big.NewInt(0)
	if createResp.Tx.Value != "" {
		var ok bool
		value, ok = new(big.Int).SetString(createResp.Tx.Value, 0)
		if !ok {
			return nil, fmt.Errorf("failed to parse transaction value: %s", createResp.Tx.Value)
		}
	}

	// Decode hex-encoded transaction data from API response
	txDataHex := strings.TrimPrefix(createResp.Tx.Data, "0x")
	txData, err := hex.DecodeString(txDataHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode transaction data: %w", err)
	}

	// Get approval spender address - use DlnSource contract (tx.to)
	approvalSpender := toAddress

	// Token bridges need approval (not native token)
	needsApproval := req.Quote.FromAsset.Address != ""

	return &BridgeResult{
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

// verifyBridgeAddress verifies that the bridge address is a known deBridge address
// This is CRITICAL for security - sending to wrong address = lost funds
func (p *DeBridgeProvider) verifyBridgeAddress(chain, address string) error {
	knownAddress, ok := debridgeDlnSourceAddresses[chain]
	if !ok {
		return fmt.Errorf("no known deBridge address for chain %s", chain)
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

// deBridge API response types
type debridgeQuoteResponse struct {
	Estimation debridgeEstimation `json:"estimation"`
	OrderId    string             `json:"orderId"`
	Tx         debridgeQuoteTx    `json:"tx"`
}

type debridgeQuoteTx struct {
	AllowanceTarget string `json:"allowanceTarget"`
}
type debridgeEstimation struct {
	SrcChainTokenIn  debridgeTokenAmount `json:"srcChainTokenIn"`
	DstChainTokenOut debridgeTokenAmount `json:"dstChainTokenOut"`
}

type debridgeTokenAmount struct {
	Address  string `json:"address"`
	Amount   string `json:"amount"`
	Decimals int    `json:"decimals"`
	Symbol   string `json:"symbol"`
}

type debridgeCreateTxResponse struct {
	Tx *debridgeTransaction `json:"tx,omitempty"`
}

type debridgeTransaction struct {
	To       string `json:"to"`
	Data     string `json:"data"`
	Value    string `json:"value"`
	GasLimit string `json:"gasLimit,omitempty"`
}

type debridgeErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
