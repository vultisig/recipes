package maya

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/vultisig/mobile-tss-lib/tss"
)

// RPCClient interface for MAYAChain JSON-RPC calls
type RPCClient interface {
	BroadcastTx(ctx context.Context, txBytes []byte) (*BroadcastTxResponse, error)
}

// HTTPRPCClient implements RPCClient using HTTP
type HTTPRPCClient struct {
	endpoints []string
	client    *http.Client
}

// SDK represents the MAYAChain SDK for transaction signing and broadcasting
type SDK struct {
	rpcClient RPCClient
	cdc       codec.Codec
}

// BroadcastTxRequest is the request payload for broadcasting transactions
type BroadcastTxRequest struct {
	TxBytes []byte `json:"tx_bytes"`
	Mode    string `json:"mode"`
}

// BroadcastTxResponse is the response from broadcasting a transaction
type BroadcastTxResponse struct {
	TxResponse *TxResponse `json:"tx_response"`
}

// TxResponse contains the result of a transaction broadcast
type TxResponse struct {
	Code      uint32 `json:"code"`
	Codespace string `json:"codespace"`
	Data      string `json:"data"`
	RawLog    string `json:"raw_log"`
	TxHash    string `json:"txhash"`
}

// MAYAChain mainnet endpoints (REST API)
var MainnetEndpoints = []string{
	"https://mayanode.mayachain.info",
	"https://maya-api.polkachu.com",
}

// MAYAChain stagenet endpoints
var StagenetEndpoints = []string{
	"https://stagenet.mayanode.mayachain.info",
}

var secpN, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
var secpHalfN = new(big.Int).Rsh(new(big.Int).Set(secpN), 1)

// NewSDK creates a new MAYAChain SDK instance
func NewSDK(rpcClient RPCClient) *SDK {
	ir := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(ir)
	banktypes.RegisterInterfaces(ir)

	return &SDK{
		rpcClient: rpcClient,
		cdc:       codec.NewProtoCodec(ir),
	}
}

// NewHTTPRPCClient creates a new HTTP RPC client with the given endpoints
func NewHTTPRPCClient(endpoints []string) *HTTPRPCClient {
	return &HTTPRPCClient{
		endpoints: endpoints,
		client:    &http.Client{Timeout: 30 * time.Second},
	}
}

// BroadcastTx broadcasts a signed transaction to the MAYAChain network
func (c *HTTPRPCClient) BroadcastTx(ctx context.Context, txBytes []byte) (*BroadcastTxResponse, error) {
	txBase64 := base64.StdEncoding.EncodeToString(txBytes)

	requestBody := map[string]interface{}{
		"tx_bytes": txBase64,
		"mode":     "BROADCAST_MODE_SYNC",
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	var lastErr error
	for _, endpoint := range c.endpoints {
		url := endpoint + "/cosmos/tx/v1beta1/txs"

		req, reqErr := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(requestJSON))
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

		var broadcastResp BroadcastTxResponse
		if err := json.Unmarshal(body, &broadcastResp); err != nil {
			lastErr = fmt.Errorf("failed to parse response from %s: %w", endpoint, err)
			continue
		}

		if broadcastResp.TxResponse != nil && broadcastResp.TxResponse.Code != 0 {
			lastErr = fmt.Errorf("broadcast failed at %s: code=%d, log=%s",
				endpoint, broadcastResp.TxResponse.Code, broadcastResp.TxResponse.RawLog)
			continue
		}

		return &broadcastResp, nil
	}

	return nil, fmt.Errorf("all endpoints failed, last error: %w", lastErr)
}

// Sign applies TSS signatures to an unsigned MAYAChain transaction
func (sdk *SDK) Sign(unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse, pubKey []byte) ([]byte, error) {
	if len(signatures) == 0 {
		return nil, fmt.Errorf("no signatures provided")
	}
	if len(pubKey) != 33 {
		return nil, fmt.Errorf("pubkey must be 33 bytes (compressed), got %d", len(pubKey))
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

	sLow, err := normalizeLowS(sBytes)
	if err != nil {
		return nil, fmt.Errorf("low-S normalization failed: %w", err)
	}

	var unsignedTx tx.Tx
	if err := sdk.cdc.Unmarshal(unsignedTxBytes, &unsignedTx); err != nil {
		return nil, fmt.Errorf("failed to unmarshal unsigned transaction: %w", err)
	}

	cosmosPubKey := &secp256k1.PubKey{Key: pubKey}

	sigBytes := make([]byte, 64)
	copy(sigBytes[:32], rBytes)
	copy(sigBytes[32:], sLow)

	pubKeyAny, err := codectypes.NewAnyWithValue(cosmosPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubkey any: %w", err)
	}

	signerInfo := &tx.SignerInfo{
		PublicKey: pubKeyAny,
		ModeInfo: &tx.ModeInfo{
			Sum: &tx.ModeInfo_Single_{
				Single: &tx.ModeInfo_Single{
					Mode: signing.SignMode_SIGN_MODE_DIRECT,
				},
			},
		},
		Sequence: 0,
	}

	if unsignedTx.AuthInfo == nil {
		unsignedTx.AuthInfo = &tx.AuthInfo{
			Fee: &tx.Fee{
				Amount:   sdk.sdkCoins(nil),
				GasLimit: 200000,
			},
		}
	}
	unsignedTx.AuthInfo.SignerInfos = []*tx.SignerInfo{signerInfo}
	unsignedTx.Signatures = [][]byte{sigBytes}

	signedTxBytes, err := sdk.cdc.Marshal(&unsignedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal signed transaction: %w", err)
	}

	return signedTxBytes, nil
}

func (s *SDK) sdkCoins(coins []struct{ Denom, Amount string }) cosmostypes.Coins {
	if len(coins) == 0 {
		return cosmostypes.NewCoins()
	}
	result := make(cosmostypes.Coins, 0, len(coins))
	for _, c := range coins {
		amount, ok := math.NewIntFromString(c.Amount)
		if !ok {
			continue
		}
		result = append(result, cosmostypes.NewCoin(c.Denom, amount))
	}
	return result
}

// Broadcast submits a signed transaction to the MAYAChain network
func (s *SDK) Broadcast(ctx context.Context, signedTxBytes []byte) (*BroadcastTxResponse, error) {
	if s.rpcClient == nil {
		return nil, fmt.Errorf("rpc client not configured")
	}

	return s.rpcClient.BroadcastTx(ctx, signedTxBytes)
}

// Send is a convenience method that signs and broadcasts the transaction
func (s *SDK) Send(ctx context.Context, unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse, pubKey []byte) (*BroadcastTxResponse, error) {
	signedTxBytes, err := s.Sign(unsignedTxBytes, signatures, pubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	return s.Broadcast(ctx, signedTxBytes)
}

// ComputeTxHash computes the transaction hash for a signed MAYAChain transaction
func (s *SDK) ComputeTxHash(signedTxBytes []byte) string {
	hash := sha256.Sum256(signedTxBytes)
	return strings.ToUpper(hex.EncodeToString(hash[:]))
}

// GetPubKeyFromBytes creates a Cosmos public key from compressed bytes
func GetPubKeyFromBytes(pubKeyBytes []byte) (cryptotypes.PubKey, error) {
	if len(pubKeyBytes) != 33 {
		return nil, fmt.Errorf("invalid pubkey length: expected 33, got %d", len(pubKeyBytes))
	}
	return &secp256k1.PubKey{Key: pubKeyBytes}, nil
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

