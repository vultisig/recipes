package thorchain

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/vultisig/mobile-tss-lib/tss"
)

// RPCClient interface for THORChain CometBFT RPC calls
type RPCClient interface {
	BroadcastTxSync(ctx context.Context, tx []byte) (*coretypes.ResultBroadcastTx, error)
}

// CometBFTRPCClient implements RPCClient using CometBFT HTTP client
type CometBFTRPCClient struct {
	client *http.HTTP
	codec  codec.Codec
}

// SDK represents the THORChain SDK for transaction signing and broadcasting
type SDK struct {
	rpcClient RPCClient
	codec     codec.Codec
}

// THORChain mainnet endpoints
var MainnetEndpoints = []string{
	"https://rpc.thorchain.info",
	"https://thornode.ninerealms.com",
	"https://rpc-thorchain.nodeist.net",
}

// THORChain testnet endpoints
var TestnetEndpoints = []string{
	"https://testnet.rpc.thorchain.info",
	"https://testnet.thornode.thorchain.info",
}

// NewSDK creates a new THORChain SDK instance
func NewSDK(rpcClient RPCClient) *SDK {
	return &SDK{
		rpcClient: rpcClient,
		codec:     makeCodec(),
	}
}

// NewCometBFTRPCClient creates a new CometBFT RPC client
func NewCometBFTRPCClient(endpoint string) (*CometBFTRPCClient, error) {
	client, err := http.New(endpoint, "/websocket")
	if err != nil {
		return nil, fmt.Errorf("failed to create CometBFT client: %w", err)
	}

	return &CometBFTRPCClient{
		client: client,
		codec:  makeCodec(),
	}, nil
}

// makeCodec creates a codec for THORChain transactions
func makeCodec() codec.Codec {
	registry := types.NewInterfaceRegistry()
	return codec.NewProtoCodec(registry)
}

// BroadcastTxSync submits a signed transaction to the THORChain network using CometBFT client
func (c *CometBFTRPCClient) BroadcastTxSync(ctx context.Context, txBytes []byte) (*coretypes.ResultBroadcastTx, error) {
	result, err := c.client.BroadcastTxSync(ctx, txBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to broadcast transaction: %w", err)
	}

	// Check if the transaction was accepted
	if result.Code != 0 {
		return result, fmt.Errorf("transaction rejected: code=%d, log=%s, codespace=%s",
			result.Code, result.Log, result.Codespace)
	}

	return result, nil
}

// Sign applies TSS signatures to an unsigned THORChain transaction
func (sdk *SDK) Sign(unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse) ([]byte, error) {
	if len(signatures) == 0 {
		return nil, fmt.Errorf("no signatures provided")
	}

	// THORChain typically uses single signatures
	if len(signatures) != 1 {
		return nil, fmt.Errorf("expected 1 signature, got %d", len(signatures))
	}

	// Parse the unsigned transaction using Cosmos SDK
	var unsignedTx tx.Tx
	if err := sdk.codec.Unmarshal(unsignedTxBytes, &unsignedTx); err != nil {
		return nil, fmt.Errorf("failed to unmarshal unsigned transaction: %w", err)
	}

	// Get the signature
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

	// Convert to DER format for Cosmos SDK
	derSig, err := sdk.toDER(rBytes, sBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to DER: %w", err)
	}

	// Create the signed transaction with proper Cosmos SDK structure
	signedTx := tx.Tx{
		Body:       unsignedTx.Body,
		AuthInfo:   unsignedTx.AuthInfo,
		Signatures: [][]byte{derSig},
	}

	// Marshal the signed transaction back to bytes
	signedTxBytes, err := sdk.codec.Marshal(&signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal signed transaction: %w", err)
	}

	return signedTxBytes, nil
}

// Broadcast submits a signed transaction to the THORChain network
func (sdk *SDK) Broadcast(ctx context.Context, signedTxBytes []byte) error {
	if sdk.rpcClient == nil {
		return fmt.Errorf("rpc client not configured")
	}

	_, err := sdk.rpcClient.BroadcastTxSync(ctx, signedTxBytes)
	return err
}

// Send is a convenience method that signs and broadcasts the transaction
func (sdk *SDK) Send(ctx context.Context, unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse) error {
	// Sign the transaction
	signedTxBytes, err := sdk.Sign(unsignedTxBytes, signatures)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Broadcast the signed transaction
	return sdk.Broadcast(ctx, signedTxBytes)
}

// MessageHash calculates the message hash for signing
func (sdk *SDK) MessageHash(unsignedTxBytes []byte) ([]byte, error) {
	// THORChain uses standard Cosmos SDK transaction signing
	hash := sha256.Sum256(unsignedTxBytes)
	return hash[:], nil
}

// deriveKeyFromMessage derives a key from a message hash for TSS lookup
func (sdk *SDK) deriveKeyFromMessage(messageHash []byte) string {
	hash := sha256.Sum256(messageHash)
	return base64.StdEncoding.EncodeToString(hash[:])
}

// toDER converts R,S to DER format for Cosmos SDK signatures
func (sdk *SDK) toDER(r, s []byte) ([]byte, error) {
	// Remove leading zeros
	r = sdk.trimLeftZeros(r)
	s = sdk.trimLeftZeros(s)

	// Add 0x00 prefix if high bit is set
	if len(r) > 0 && (r[0]&0x80) != 0 {
		r = append([]byte{0x00}, r...)
	}
	if len(s) > 0 && (s[0]&0x80) != 0 {
		s = append([]byte{0x00}, s...)
	}

	// Build DER sequence: SEQUENCE { INTEGER r, INTEGER s }
	intR := append([]byte{0x02, byte(len(r))}, r...)
	intS := append([]byte{0x02, byte(len(s))}, s...)
	content := append(intR, intS...)
	
	if len(content) > 255 {
		return nil, fmt.Errorf("DER signature too long")
	}
	
	seq := []byte{0x30, byte(len(content))}
	return append(seq, content...), nil
}

// trimLeftZeros removes leading zeros from byte slice, keeping at least one byte
func (sdk *SDK) trimLeftZeros(b []byte) []byte {
	i := 0
	for i < len(b) && b[i] == 0x00 {
		i++
	}
	if i == len(b) {
		return []byte{0x00} // Keep at least one zero
	}
	return b[i:]
}

// cleanHex removes 0x prefix from hex strings
func cleanHex(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		return s[2:]
	}
	return s
}