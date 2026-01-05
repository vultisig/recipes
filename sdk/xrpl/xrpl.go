package xrpl

import (
	"bytes"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	ecdsa2 "github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"github.com/vultisig/mobile-tss-lib/tss"
	xrpgo "github.com/xyield/xrpl-go/binary-codec"
)

// RPCClient interface for XRPL JSON-RPC calls
type RPCClient interface {
	SubmitTransaction(ctx context.Context, txBlob string) error
}

// HTTPRPCClient implements RPCClient using HTTP JSON-RPC
type HTTPRPCClient struct {
	endpoints []string
	client    *http.Client
}

// SDK represents the XRP SDK for transaction signing and broadcasting
type SDK struct {
	rpcClient RPCClient
}

// JSONRPC request structure
type JSONRPCRequest struct {
	Method string `json:"method"`
	Params []any  `json:"params"`
	ID     int    `json:"id"`
}

// JSONRPC response structure
type JSONRPCResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *JSONRPCError   `json:"error"`
	ID     int             `json:"id"`
}

type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Submit request parameters
type SubmitParams struct {
	TxBlob   string `json:"tx_blob"`
	FailHard bool   `json:"fail_hard"`
}

// SubmitResponse represents the basic response from transaction submission
// The SDK only cares about successful submission, detailed tracking is handled by plugins
type SubmitResponse struct {
	EngineResult        string `json:"engine_result"`
	EngineResultCode    int    `json:"engine_result_code"`
	EngineResultMessage string `json:"engine_result_message"`
}

// XRPL mainnet endpoints
var MainnetEndpoints = []string{
	"https://s1.ripple.com:51234",
	"https://s2.ripple.com:51234",
	"https://xrplcluster.com",
}

// XRPL testnet endpoints
var TestnetEndpoints = []string{
	"https://s.altnet.rippletest.net:51234",
	"https://testnet.xrpl-labs.com",
}

var stxPrefix = []byte{0x53, 0x54, 0x58, 0x00} // "STX\0"
var secpN, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
var secpHalfN = new(big.Int).Rsh(new(big.Int).Set(secpN), 1)

// NewSDK creates a new XRP SDK instance
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

// SubmitTransaction submits a signed transaction to the XRPL network
func (c *HTTPRPCClient) SubmitTransaction(ctx context.Context, txBlob string) error {
	// Clean hex string (remove 0x prefix if present)
	txBlob = strings.TrimPrefix(strings.ToUpper(txBlob), "0X")

	submitParams := SubmitParams{
		TxBlob:   txBlob,
		FailHard: false,
	}

	request := JSONRPCRequest{
		Method: "submit",
		Params: []any{submitParams},
		ID:     1,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	var lastErr error
	for _, endpoint := range c.endpoints {
		req, reqErr := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(requestBody))
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

		var jr JSONRPCResponse
		if err := json.Unmarshal(body, &jr); err != nil {
			lastErr = fmt.Errorf("failed to parse response from %s: %w", endpoint, err)
			continue
		}

		if jr.Error != nil {
			lastErr = fmt.Errorf("submit error from %s: %d: %s", endpoint, jr.Error.Code, jr.Error.Message)
			continue
		}

		var submit struct {
			Status              string `json:"status"`
			EngineResult        string `json:"engine_result"`
			EngineResultCode    int    `json:"engine_result_code"`
			EngineResultMessage string `json:"engine_result_message"`
		}
		if err := json.Unmarshal(jr.Result, &submit); err != nil {
			lastErr = fmt.Errorf("failed to parse submit result from %s: %w", endpoint, err)
			continue
		}

		if strings.ToUpper(submit.EngineResult) != "TESSUCCESS" /* || strings.ToLower(submit.Status) != "success" */ {
			lastErr = fmt.Errorf("submit failed at %s: %s (%d): %s",
				endpoint, submit.EngineResult, submit.EngineResultCode, submit.EngineResultMessage)
			continue
		}
		return nil
	}

	return fmt.Errorf("all endpoints failed, last error: %w", lastErr)
}

// Sign applies TSS signatures to an unsigned XRPL transaction
func (sdk *SDK) Sign(unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse, pubKey []byte) ([]byte, error) {
	if len(signatures) == 0 {
		return nil, fmt.Errorf("no signatures provided")
	}
	if len(pubKey) != 33 {
		return nil, fmt.Errorf("pubkey must be 33 bytes (compressed), got %d", len(pubKey))
	}

	// Get the first (and typically only) signature for single-signature transactions
	var sig tss.KeysignResponse
	for _, v := range signatures {
		sig = v
		break
	}

	// Decode R and S from hex strings
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

	sLow, err := normalizeLowS(sBytes)
	if err != nil {
		return nil, fmt.Errorf("low-S normalization failed: %w", err)
	}
	sBytes = sLow

	// Convert bytes to hex string for binary codec
	baseHex := hex.EncodeToString(unsignedTxBytes)

	// Decode the unsigned base using the codec
	m, err := xrpgo.Decode(strings.ToUpper(baseHex))
	if err != nil {
		return nil, fmt.Errorf("binary-codec decode (unsigned) failed: %w", err)
	}

	// Ensure base doesn't already contain TxnSignature
	if _, has := m["TxnSignature"]; has {
		return nil, fmt.Errorf("base tx already contains TxnSignature; pass an unsigned base")
	}

	// SigningPubKey should already be present from the build script
	// If it's not present, add it as a fallback
	if _, has := m["SigningPubKey"]; !has {
		pubHex := strings.ToUpper(hex.EncodeToString(pubKey))
		m["SigningPubKey"] = pubHex
	}

	// Re-encode with SigningPubKey (but without TxnSignature) to get canonical bytes for signing
	withPubHex, err := xrpgo.Encode(m)
	if err != nil {
		return nil, fmt.Errorf("encode(with SigningPubKey) failed: %w", err)
	}
	withPubBytes, err := hex.DecodeString(withPubHex)
	if err != nil {
		return nil, fmt.Errorf("decode withPubHex failed: %w", err)
	}

	// Compute XRPL single-sign digest over canonical bytes
	preimage := append(append([]byte{}, stxPrefix...), withPubBytes...)
	digest := sha512Half(preimage)

	// Verify (r,s) locally against that digest
	ok, err := sdk.verifyRS(digest, rBytes, sBytes, pubKey)
	if err != nil {
		return nil, fmt.Errorf("signature verification failed: %w", err)
	}
	if !ok {
		return nil, fmt.Errorf("invalid signature: (R,S) do not verify against canonical XRPL digest with provided pubkey")
	}

	// Set TxnSignature (DER) into the map and re-encode canonically for the final tx_blob
	der := sdk.derEncodeRS(rBytes, sBytes)
	m["TxnSignature"] = strings.ToUpper(hex.EncodeToString(der))

	finalHex, err := xrpgo.Encode(m)
	if err != nil {
		return nil, fmt.Errorf("encode(final with TxnSignature) failed: %w", err)
	}

	finalBytes, err := hex.DecodeString(finalHex)
	if err != nil {
		return nil, fmt.Errorf("decode final hex failed: %w", err)
	}

	return finalBytes, nil
}

// Broadcast submits a signed transaction to the XRPL network
func (sdk *SDK) Broadcast(ctx context.Context, signedTxBytes []byte) error {
	if sdk.rpcClient == nil {
		return fmt.Errorf("rpc client not configured")
	}

	txHex := hex.EncodeToString(signedTxBytes)
	return sdk.rpcClient.SubmitTransaction(ctx, txHex)
}

// Send is a convenience method that signs and broadcasts the transaction
func (sdk *SDK) Send(ctx context.Context, unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse, pubKey []byte) error {
	// Sign the transaction
	signedTxBytes, err := sdk.Sign(unsignedTxBytes, signatures, pubKey)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Broadcast the signed transaction
	return sdk.Broadcast(ctx, signedTxBytes)
}

// sha512Half computes SHA512 and returns first 32 bytes
func sha512Half(b []byte) []byte {
	h := sha512.Sum512(b)
	return h[:32]
}

// derEncodeRS encodes R,S as ASN.1 DER format
func (sdk *SDK) derEncodeRS(r, s []byte) []byte {
	r = sdk.trimLeftZeros(r)
	s = sdk.trimLeftZeros(s)

	if len(r) == 0 || (r[0]&0x80) != 0 {
		r = append([]byte{0x00}, r...)
	}
	if len(s) == 0 || (s[0]&0x80) != 0 {
		s = append([]byte{0x00}, s...)
	}

	intR := append([]byte{0x02, byte(len(r))}, r...)
	intS := append([]byte{0x02, byte(len(s))}, s...)
	content := append(intR, intS...)
	seq := []byte{0x30, byte(len(content))}
	return append(seq, content...)
}

// trimLeftZeros removes leading zeros from byte slice
func (sdk *SDK) trimLeftZeros(b []byte) []byte {
	i := 0
	for i < len(b) && b[i] == 0x00 {
		i++
	}
	return b[i:]
}

// verifyRS verifies (r,s) signature over digest with compressed pubkey
func (sdk *SDK) verifyRS(digest, r, s []byte, pub33 []byte) (bool, error) {
	pk, err := secp256k1.ParsePubKey(pub33)
	if err != nil {
		return false, fmt.Errorf("parse pubkey: %w", err)
	}
	var rN, sN secp256k1.ModNScalar
	if rN.SetByteSlice(r) {
		return false, fmt.Errorf("r overflow/not in field")
	}
	if sN.SetByteSlice(s) {
		return false, fmt.Errorf("s overflow/not in field")
	}
	sig := ecdsa2.NewSignature(&rN, &sN)
	return sig.Verify(digest, pk), nil
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
	// left-pad to 32 bytes
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
