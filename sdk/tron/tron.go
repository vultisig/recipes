package tron

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/vultisig/mobile-tss-lib/tss"
)

// RPCClient interface for TRON JSON-RPC calls
type RPCClient interface {
	BroadcastTransaction(ctx context.Context, signedTx *SignedTransaction) (*BroadcastResponse, error)
}

// HTTPRPCClient implements RPCClient using HTTP
type HTTPRPCClient struct {
	endpoints []string
	client    *http.Client
}

// SDK represents the TRON SDK for transaction signing and broadcasting
type SDK struct {
	rpcClient RPCClient
}

// SignedTransaction represents a signed TRON transaction
type SignedTransaction struct {
	TxID       string   `json:"txID"`
	RawData    *RawData `json:"raw_data,omitempty"`
	RawDataHex string   `json:"raw_data_hex"`
	Signature  []string `json:"signature"`
}

// RawData represents the raw data of a TRON transaction
type RawData struct {
	Contract      []Contract `json:"contract"`
	RefBlockBytes string     `json:"ref_block_bytes"`
	RefBlockHash  string     `json:"ref_block_hash"`
	Expiration    int64      `json:"expiration"`
	Timestamp     int64      `json:"timestamp"`
	FeeLimit      int64      `json:"fee_limit,omitempty"`
	Data          string     `json:"data,omitempty"`
}

// Contract represents a contract in a TRON transaction
type Contract struct {
	Parameter Parameter `json:"parameter"`
	Type      string    `json:"type"`
}

// Parameter represents the parameter of a contract
type Parameter struct {
	Value   Value  `json:"value"`
	TypeUrl string `json:"type_url"`
}

// Value represents the value of a contract parameter
type Value struct {
	Amount       int64  `json:"amount,omitempty"`
	OwnerAddress string `json:"owner_address"`
	ToAddress    string `json:"to_address,omitempty"`
}

// BroadcastResponse represents the response from broadcasting a transaction
type BroadcastResponse struct {
	Result  bool   `json:"result"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	TxID    string `json:"txid,omitempty"`
}

// TRON mainnet endpoints
var MainnetEndpoints = []string{
	"https://api.trongrid.io",
	"https://api.tronstack.io",
}

// TRON testnet (Shasta) endpoints
var TestnetEndpoints = []string{
	"https://api.shasta.trongrid.io",
}

var secpN, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
var secpHalfN = new(big.Int).Rsh(new(big.Int).Set(secpN), 1)

// NewSDK creates a new TRON SDK instance
func NewSDK(rpcClient RPCClient) *SDK {
	return &SDK{
		rpcClient: rpcClient,
	}
}

// NewHTTPRPCClient creates a new HTTP RPC client with the given endpoints
func NewHTTPRPCClient(endpoints []string) *HTTPRPCClient {
	return &HTTPRPCClient{
		endpoints: endpoints,
		client:    &http.Client{Timeout: 30 * time.Second},
	}
}

// BroadcastTransaction broadcasts a signed transaction to the TRON network
func (c *HTTPRPCClient) BroadcastTransaction(ctx context.Context, signedTx *SignedTransaction) (*BroadcastResponse, error) {
	requestBody, err := json.Marshal(signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal transaction: %w", err)
	}

	var lastErr error
	for _, endpoint := range c.endpoints {
		url := endpoint + "/wallet/broadcasttransaction"

		req, reqErr := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(requestBody))
		if reqErr != nil {
			lastErr = fmt.Errorf("failed to create request for %s: %w", endpoint, reqErr)
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("failed to send request to %s: %w", endpoint, err)
			continue
		}

		body, readErr := io.ReadAll(resp.Body)
		_ = resp.Body.Close()

		if readErr != nil {
			lastErr = fmt.Errorf("failed to read response from %s: %w", endpoint, readErr)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("HTTP error from %s: %d, body: %s", endpoint, resp.StatusCode, string(body))
			continue
		}

		var broadcastResp BroadcastResponse
		if err := json.Unmarshal(body, &broadcastResp); err != nil {
			lastErr = fmt.Errorf("failed to parse response from %s: %w", endpoint, err)
			continue
		}

		if !broadcastResp.Result {
			lastErr = fmt.Errorf("broadcast failed at %s: code=%s, message=%s",
				endpoint, broadcastResp.Code, broadcastResp.Message)
			continue
		}

		return &broadcastResp, nil
	}

	return nil, fmt.Errorf("all endpoints failed, last error: %w", lastErr)
}

// Sign applies TSS signatures to an unsigned TRON transaction
func (sdk *SDK) Sign(unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse, pubKey []byte) ([]byte, error) {
	if len(signatures) == 0 {
		return nil, fmt.Errorf("no signatures provided")
	}
	if len(pubKey) != 33 && len(pubKey) != 65 {
		return nil, fmt.Errorf("pubkey must be 33 or 65 bytes, got %d", len(pubKey))
	}

	var sig tss.KeysignResponse
	for _, v := range signatures {
		sig = v
		break
	}

	rHex := cleanHex(sig.R)
	rBytes, err := hex.DecodeString(rHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode R: %w", err)
	}
	sHex := cleanHex(sig.S)
	sBytes, err := hex.DecodeString(sHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode S: %w", err)
	}

	if len(rBytes) != 32 {
		return nil, fmt.Errorf("r must be 32 bytes, got %d", len(rBytes))
	}
	if len(sBytes) != 32 {
		return nil, fmt.Errorf("s must be 32 bytes, got %d", len(sBytes))
	}

	// Normalize S to low-S form
	sLow, err := normalizeLowS(sBytes)
	if err != nil {
		return nil, fmt.Errorf("low-S normalization failed: %w", err)
	}

	// Compute transaction ID (SHA256 of raw_data)
	txID := sha256.Sum256(unsignedTxBytes)
	txIDHex := hex.EncodeToString(txID[:])

	// TRON signature format: r (32 bytes) || s (32 bytes) || v (1 byte)
	// Recovery ID (v) needs to be computed based on the signature
	sigBytes := make([]byte, 65)
	copy(sigBytes[:32], rBytes)
	copy(sigBytes[32:64], sLow)
	
	// Parse recovery ID from hex string
	recoveryID := byte(0)
	if sig.RecoveryID != "" {
		recIDHex := cleanHex(sig.RecoveryID)
		if len(recIDHex) > 0 {
			recIDBytes, err := hex.DecodeString(recIDHex)
			if err != nil {
				return nil, fmt.Errorf("failed to decode recovery ID: %w", err)
			}
			if len(recIDBytes) > 0 {
				recoveryID = recIDBytes[0]
			}
		}
	}
	sigBytes[64] = recoveryID + 27 // Add 27 to recovery ID

	sigHex := hex.EncodeToString(sigBytes)

	// Create signed transaction JSON
	// Note: TRON nodes accept either raw_data (JSON object) or raw_data_hex (serialized protobuf).
	// Using raw_data_hex is simpler and sufficient for broadcast - the node will deserialize it.
	signedTx := &SignedTransaction{
		TxID:       txIDHex,
		RawDataHex: hex.EncodeToString(unsignedTxBytes),
		Signature:  []string{sigHex},
	}

	signedTxBytes, err := json.Marshal(signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal signed transaction: %w", err)
	}

	return signedTxBytes, nil
}

// Broadcast submits a signed transaction to the TRON network
func (sdk *SDK) Broadcast(ctx context.Context, signedTxBytes []byte) (*BroadcastResponse, error) {
	if sdk.rpcClient == nil {
		return nil, fmt.Errorf("rpc client not configured")
	}

	var signedTx SignedTransaction
	if err := json.Unmarshal(signedTxBytes, &signedTx); err != nil {
		return nil, fmt.Errorf("failed to unmarshal signed transaction: %w", err)
	}

	return sdk.rpcClient.BroadcastTransaction(ctx, &signedTx)
}

// Send is a convenience method that signs and broadcasts the transaction
func (sdk *SDK) Send(ctx context.Context, unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse, pubKey []byte) (*BroadcastResponse, error) {
	signedTxBytes, err := sdk.Sign(unsignedTxBytes, signatures, pubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	return sdk.Broadcast(ctx, signedTxBytes)
}

// ComputeTxHash computes the transaction hash for a TRON transaction
func (sdk *SDK) ComputeTxHash(rawDataBytes []byte) string {
	hash := sha256.Sum256(rawDataBytes)
	return hex.EncodeToString(hash[:])
}

func normalizeLowS(s []byte) ([]byte, error) {
	if len(s) == 0 {
		return nil, fmt.Errorf("empty s")
	}
	var sb big.Int
	sb.SetBytes(s)
	if sb.Sign() <= 0 || sb.Cmp(secpN) >= 0 {
		return nil, fmt.Errorf("s not in [1, N-1]")
	}
	if sb.Cmp(secpHalfN) > 0 {
		sb.Sub(secpN, &sb)
	}
	out := sb.Bytes()
	if len(out) < 32 {
		pad := make([]byte, 32-len(out))
		out = append(pad, out...)
	}
	return out, nil
}

func cleanHex(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		return s[2:]
	}
	return s
}

