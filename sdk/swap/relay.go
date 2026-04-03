package swap

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/gagliardetto/solana-go"
	addresslookuptable "github.com/gagliardetto/solana-go/programs/address-lookup-table"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/vultisig/recipes/resolver"
)

const (
	relayDefaultBaseURL  = "https://api.relay.link"
	relayReferrer        = "vultisig"
	relayAffiliateBps    = "50" // 0.5%
)

// Relay chain IDs
var relayChainIDs = map[string]int{
	"Ethereum":  1,
	"BSC":       56,
	"Polygon":   137,
	"Avalanche": 43114,
	"Arbitrum":  42161,
	"Optimism":  10,
	"Base":      8453,
	"Mantle":    5000,
	"Zksync":    324,
	"Sei":       1329,
	"Solana":    792703809,
}

// relaySupportedAllChains lists all Relay-supported chains.
var relaySupportedAllChains = func() []string {
	chains := make([]string, 0, len(relayChainIDs))
	for chain := range relayChainIDs {
		chains = append(chains, chain)
	}
	return chains
}()

// RelayProvider implements SwapProvider for Relay.link
type RelayProvider struct {
	BaseProvider
	client  *http.Client
	baseURL string
	solRPC  *rpc.Client
}

// NewRelayProvider creates a new Relay provider
func NewRelayProvider(solRPC *rpc.Client) *RelayProvider {
	return &RelayProvider{
		BaseProvider: NewBaseProvider("Relay", PriorityRelay, relaySupportedAllChains),
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		baseURL: relayDefaultBaseURL,
		solRPC:  solRPC,
	}
}

// SupportsRoute checks if Relay can handle a swap between two assets
func (p *RelayProvider) SupportsRoute(from, to Asset) bool {
	return p.SupportsChain(from.Chain) && p.SupportsChain(to.Chain)
}

// IsAvailable checks if Relay is available for a specific chain
func (p *RelayProvider) IsAvailable(ctx context.Context, chain string) (bool, error) {
	return p.SupportsChain(chain), nil
}

// GetStatus returns detailed availability status for a chain
func (p *RelayProvider) GetStatus(ctx context.Context, chain string) (*ProviderStatus, error) {
	return &ProviderStatus{
		Chain:     chain,
		Available: p.SupportsChain(chain),
	}, nil
}

// GetQuote gets a swap quote from Relay
func (p *RelayProvider) GetQuote(ctx context.Context, req QuoteRequest) (*Quote, error) {
	originChainID, ok := relayChainIDs[req.From.Chain]
	if !ok {
		return nil, fmt.Errorf("relay: unsupported origin chain %q", req.From.Chain)
	}
	destChainID, ok := relayChainIDs[req.To.Chain]
	if !ok {
		return nil, fmt.Errorf("relay: unsupported destination chain %q", req.To.Chain)
	}

	originCurrency := req.From.Address
	if originCurrency == "" {
		originCurrency = relayNativeAddress(req.From.Chain)
	}
	destCurrency := req.To.Address
	if destCurrency == "" {
		destCurrency = relayNativeAddress(req.To.Chain)
	}

	recipient := req.Destination
	if recipient == "" {
		recipient = req.Sender
	}

	quoteReq := relayQuoteRequest{
		User:                req.Sender,
		OriginChainID:       originChainID,
		DestinationChainID:  destChainID,
		OriginCurrency:      originCurrency,
		DestinationCurrency: destCurrency,
		Amount:              req.Amount.String(),
		TradeType:           "EXACT_INPUT",
		Recipient:           recipient,
		Referrer:            relayReferrer,
		AppFees: []relayAppFee{{
			Recipient: resolver.DefaultEVMTreasuryAddress,
			Fee:       relayAffiliateBps,
		}},
	}

	// Reduce Solana transaction size by using shared accounts.
	if req.From.Chain == "Solana" {
		useShared := true
		quoteReq.UseSharedAccounts = &useShared
	}

	quoteResp, err := p.postQuote(ctx, quoteReq)
	if err != nil {
		return nil, err
	}

	// Parse expected output amount
	expectedOutput, ok := new(big.Int).SetString(quoteResp.Details.CurrencyOut.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("relay: invalid output amount %q", quoteResp.Details.CurrencyOut.Amount)
	}

	// Separate approval and swap transaction steps.
	var approvalStep *relayStepData
	var txStep *relayStepData
	for i := range quoteResp.Steps {
		step := &quoteResp.Steps[i]
		if len(step.Items) == 0 {
			continue
		}
		if step.ID == "approve" {
			approvalStep = &step.Items[0].Data
		} else if step.Kind == "transaction" && txStep == nil {
			txStep = &step.Items[0].Data
		}
	}

	// Spender is the swap TX destination (router), not the approval TX target (token contract).
	needsApproval := approvalStep != nil
	var approvalSpender string
	if needsApproval && txStep != nil {
		approvalSpender = txStep.To
	}

	// Cache quote response in ProviderData for BuildTx.
	providerData, err := json.Marshal(quoteResp)
	if err != nil {
		return nil, fmt.Errorf("relay: marshal provider data: %w", err)
	}

	// Router address for approval spender resolution.
	var router string
	if txStep != nil {
		router = txStep.To
	}

	quote := &Quote{
		Provider:        p.Name(),
		FromAsset:       req.From,
		ToAsset:         req.To,
		FromAmount:      req.Amount,
		ExpectedOutput:  expectedOutput,
		MinimumOutput:   expectedOutput, // Relay is solver-based: no AMM slippage
		Router:          router,
		NeedsApproval:   needsApproval,
		ApprovalSpender: approvalSpender,
		ApprovalAmount:  req.Amount,
		ProviderData:    providerData,
	}

	return quote, nil
}

// BuildTx builds an unsigned transaction for the swap
func (p *RelayProvider) BuildTx(ctx context.Context, req SwapRequest) (*SwapResult, error) {
	if req.Quote == nil {
		return nil, fmt.Errorf("quote is required")
	}

	// Restore cached quote response.
	var quoteResp relayQuoteResponse
	if len(req.Quote.ProviderData) > 0 {
		if err := json.Unmarshal(req.Quote.ProviderData, &quoteResp); err != nil {
			return nil, fmt.Errorf("relay: unmarshal provider data: %w", err)
		}
	} else {
		// Should not happen — GetQuote always sets ProviderData.
		return nil, fmt.Errorf("relay: missing provider data, call GetQuote first")
	}

	// Find the swap transaction step.
	var approvalStep *relayStepData
	var txStep *relayStepData
	for i := range quoteResp.Steps {
		step := &quoteResp.Steps[i]
		if len(step.Items) == 0 {
			continue
		}
		if step.ID == "approve" {
			approvalStep = &step.Items[0].Data
		} else if step.Kind == "transaction" && txStep == nil {
			txStep = &step.Items[0].Data
		}
	}
	if txStep == nil {
		return nil, fmt.Errorf("relay: no transaction step in quote response")
	}

	fromChain := req.Quote.FromAsset.Chain

	// Build the result based on chain type.
	if fromChain == "Solana" {
		txData, err := p.buildSolanaTransaction(ctx, txStep, req.Sender)
		if err != nil {
			return nil, fmt.Errorf("relay: build solana tx: %w", err)
		}
		return &SwapResult{
			Provider:    p.Name(),
			TxData:      txData,
			Value:       big.NewInt(0),
			ToAddress:   "", // Solana: program addresses are in the instructions
			ExpectedOut: req.Quote.ExpectedOutput,
		}, nil
	}

	// EVM chains
	value := big.NewInt(0)
	if txStep.Value != "" {
		var ok bool
		value, ok = new(big.Int).SetString(txStep.Value, 0)
		if !ok {
			return nil, fmt.Errorf("relay: invalid tx value %q", txStep.Value)
		}
	}

	// Decode EVM calldata
	txData, err := hex.DecodeString(trimHexPrefix(txStep.Data))
	if err != nil {
		return nil, fmt.Errorf("relay: decode tx data: %w", err)
	}

	result := &SwapResult{
		Provider:    p.Name(),
		TxData:      txData,
		Value:       value,
		ToAddress:   txStep.To,
		ExpectedOut: req.Quote.ExpectedOutput,
	}

	// Populate approval info from the cached response.
	if approvalStep != nil {
		result.NeedsApproval = true
		result.ApprovalAddress = approvalStep.To

		if req.Quote.FromAmount == nil {
			return nil, fmt.Errorf("relay: approval required but FromAmount is nil")
		}
		result.ApprovalAmount = req.Quote.FromAmount
	}

	return result, nil
}

// buildSolanaTransaction assembles a VersionedTransaction from Relay response instructions.
func (p *RelayProvider) buildSolanaTransaction(ctx context.Context, txData *relayStepData, senderAddr string) ([]byte, error) {
	if p.solRPC == nil {
		return nil, fmt.Errorf("solana RPC client required for Solana swaps")
	}

	feePayer, err := solana.PublicKeyFromBase58(senderAddr)
	if err != nil {
		return nil, fmt.Errorf("parse sender address: %w", err)
	}

	// Convert Relay instructions to solana-go instructions.
	instructions := make([]solana.Instruction, 0, len(txData.Instructions))
	for i, inst := range txData.Instructions {
		programID, err := solana.PublicKeyFromBase58(inst.ProgramID)
		if err != nil {
			return nil, fmt.Errorf("parse program id for instruction %d: %w", i, err)
		}

		accounts := make([]*solana.AccountMeta, 0, len(inst.Keys))
		for _, key := range inst.Keys {
			pk, err := solana.PublicKeyFromBase58(key.Pubkey)
			if err != nil {
				return nil, fmt.Errorf("parse account key %s: %w", key.Pubkey, err)
			}
			accounts = append(accounts, solana.NewAccountMeta(pk, key.IsWritable, key.IsSigner))
		}

		// Relay returns instruction data as hex-encoded string (typically no 0x prefix).
		data, err := hex.DecodeString(trimHexPrefix(inst.Data))
		if err != nil {
			return nil, fmt.Errorf("decode instruction data %d: %w", i, err)
		}

		instructions = append(instructions, solana.NewInstruction(programID, accounts, data))
	}

	// Fetch recent blockhash.
	block, err := p.solRPC.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, fmt.Errorf("get latest blockhash: %w", err)
	}

	opts := []solana.TransactionOption{
		solana.TransactionPayer(feePayer),
	}

	// Resolve address lookup tables if present.
	if len(txData.AddressLookupTableAddresses) > 0 {
		altMap, err := resolveRelayALTs(ctx, p.solRPC, txData.AddressLookupTableAddresses)
		if err != nil {
			return nil, fmt.Errorf("resolve address lookup tables: %w", err)
		}
		opts = append(opts, solana.TransactionAddressTables(altMap))
	}

	tx, err := solana.NewTransaction(instructions, block.Value.Blockhash, opts...)
	if err != nil {
		return nil, fmt.Errorf("create solana transaction: %w", err)
	}

	txBytes, err := tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("marshal solana transaction: %w", err)
	}

	return txBytes, nil
}

// resolveRelayALTs fetches address lookup table accounts from the Solana chain.
func resolveRelayALTs(ctx context.Context, solRPC *rpc.Client, altAddresses []string) (map[solana.PublicKey]solana.PublicKeySlice, error) {
	result := make(map[solana.PublicKey]solana.PublicKeySlice, len(altAddresses))

	for _, addr := range altAddresses {
		pk, err := solana.PublicKeyFromBase58(addr)
		if err != nil {
			return nil, fmt.Errorf("parse ALT address %s: %w", addr, err)
		}

		state, err := addresslookuptable.GetAddressLookupTableStateWithOpts(ctx, solRPC, pk, &rpc.GetAccountInfoOpts{
			Commitment: rpc.CommitmentFinalized,
		})
		if err != nil {
			return nil, fmt.Errorf("fetch ALT %s: %w", addr, err)
		}

		result[pk] = state.Addresses
	}

	return result, nil
}

// postQuote calls POST /quote on the Relay API.
func (p *RelayProvider) postQuote(ctx context.Context, req relayQuoteRequest) (*relayQuoteResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("relay: marshal quote request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/quote", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("relay: create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("relay: http request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("relay: read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("relay: api error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var quoteResp relayQuoteResponse
	if err := json.Unmarshal(respBody, &quoteResp); err != nil {
		return nil, fmt.Errorf("relay: parse quote response: %w", err)
	}

	return &quoteResp, nil
}

// relayNativeAddress returns the Relay native currency address for a chain.
func relayNativeAddress(chain string) string {
	if chain == "Solana" {
		return "11111111111111111111111111111111"
	}
	return "0x0000000000000000000000000000000000000000"
}

// trimHexPrefix removes a "0x" prefix from hex data if present.
func trimHexPrefix(s string) string {
	if len(s) >= 2 && s[:2] == "0x" {
		return s[2:]
	}
	return s
}

// --- Relay API types ---

type relayQuoteRequest struct {
	User                string        `json:"user"`
	OriginChainID       int           `json:"originChainId"`
	DestinationChainID  int           `json:"destinationChainId"`
	OriginCurrency      string        `json:"originCurrency"`
	DestinationCurrency string        `json:"destinationCurrency"`
	Amount              string        `json:"amount"`
	TradeType           string        `json:"tradeType"`
	Recipient           string        `json:"recipient"`
	Referrer            string        `json:"referrer"`
	AppFees             []relayAppFee `json:"appFees,omitempty"`

	// UseSharedAccounts prevents certain ATA creation instructions in Solana routing,
	// reducing transaction size.
	UseSharedAccounts *bool `json:"useSharedAccounts,omitempty"`
}

type relayAppFee struct {
	Recipient string `json:"recipient"`
	Fee       string `json:"fee"` // basis points
}

type relayQuoteResponse struct {
	Steps   []relayStep   `json:"steps"`
	Details relayDetails  `json:"details"`
}

type relayStep struct {
	ID    string          `json:"id"`
	Kind  string          `json:"kind"`
	Items []relayStepItem `json:"items"`
}

type relayStepItem struct {
	Data relayStepData `json:"data"`
}

// relayStepData contains the raw transaction fields from Relay.
type relayStepData struct {
	// EVM fields
	From  string `json:"from,omitempty"`
	To    string `json:"to,omitempty"`
	Data  string `json:"data,omitempty"`
	Value string `json:"value,omitempty"`

	// Solana fields
	Instructions                []relaySolInstruction `json:"instructions,omitempty"`
	AddressLookupTableAddresses []string              `json:"addressLookupTableAddresses,omitempty"`
}

type relaySolInstruction struct {
	Keys      []relaySolAccountKey `json:"keys"`
	ProgramID string               `json:"programId"`
	Data      string               `json:"data"`
}

type relaySolAccountKey struct {
	Pubkey     string `json:"pubkey"`
	IsSigner   bool   `json:"isSigner"`
	IsWritable bool   `json:"isWritable"`
}

type relayDetails struct {
	CurrencyOut  relayCurrencyOut `json:"currencyOut"`
	Rate         string           `json:"rate"`
	TimeEstimate int              `json:"timeEstimate"`
}

type relayCurrencyOut struct {
	Amount          string            `json:"amount"`
	AmountFormatted string            `json:"amountFormatted"`
	Currency        relayCurrencyMeta `json:"currency"`
}

type relayCurrencyMeta struct {
	Symbol string `json:"symbol"`
}
