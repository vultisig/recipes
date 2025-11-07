package thorchain

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cometbft/cometbft/rpc/client/http"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/vultisig/mobile-tss-lib/tss"
	vtypes "github.com/vultisig/recipes/types"
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
	// Register crypto types (required for PubKey interfaces)
	cryptocodec.RegisterInterfaces(registry)
	// Register bank module message types
	banktypes.RegisterInterfaces(registry)
	// Register THORChain MsgDeposit for swaps
	registry.RegisterImplementations((*sdk.Msg)(nil), &vtypes.MsgDeposit{})
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

	// For Cosmos SDK SIGN_MODE_DIRECT, use raw 64-byte signature: R(32) || S(32)
	rawSig := append(rBytes, sBytes...)

	// Create the signed transaction with proper Cosmos SDK structure
	signedTx := tx.Tx{
		Body:       unsignedTx.Body,
		AuthInfo:   unsignedTx.AuthInfo,
		Signatures: [][]byte{rawSig},
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

// MessageHash calculates the sign document hash for Cosmos SDK signing
func (sdk *SDK) MessageHash(unsignedTxBytes []byte, accountNumber uint64, sequence uint64) ([]byte, error) {
	// Parse the unsigned transaction
	var unsignedTx tx.Tx
	if err := sdk.codec.Unmarshal(unsignedTxBytes, &unsignedTx); err != nil {
		return nil, fmt.Errorf("failed to unmarshal unsigned transaction: %w", err)
	}

	// Validate AuthInfo exists
	if unsignedTx.AuthInfo == nil {
		return nil, fmt.Errorf("transaction missing AuthInfo")
	}

	// Extract signing info from the transaction
	if len(unsignedTx.AuthInfo.SignerInfos) == 0 {
		return nil, fmt.Errorf("transaction missing signer info")
	}

	// Validate and update sequence in AuthInfo to match the expected sequence
	signerInfo := unsignedTx.AuthInfo.SignerInfos[0]
	if signerInfo.Sequence != sequence {
		signerInfo.Sequence = sequence
	}

	chainId := getChainIdFromTx(&unsignedTx)

	// Create the sign document following Cosmos SDK SIGN_MODE_DIRECT
	bodyBytes := mustMarshalTxBody(sdk.codec, unsignedTx.Body)
	authInfoBytes := mustMarshalAuthInfo(sdk.codec, unsignedTx.AuthInfo)

	signDoc := &tx.SignDoc{
		BodyBytes:     bodyBytes,
		AuthInfoBytes: authInfoBytes,
		ChainId:       chainId,
		AccountNumber: accountNumber,
	}

	// Marshal and hash the sign document
	signDocBytes, err := sdk.codec.Marshal(signDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal sign document: %w", err)
	}

	hash := sha256.Sum256(signDocBytes)
	return hash[:], nil
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

// Helper functions for sign document creation

// mustMarshalTxBody marshals transaction body, panics on error (should not happen with valid tx)
func mustMarshalTxBody(cdc codec.Codec, body *tx.TxBody) []byte {
	if body == nil {
		panic("nil transaction body")
	}
	bytes, err := cdc.Marshal(body)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal tx body: %v", err))
	}
	return bytes
}

// mustMarshalAuthInfo marshals auth info, panics on error (should not happen with valid tx)
func mustMarshalAuthInfo(cdc codec.Codec, authInfo *tx.AuthInfo) []byte {
	if authInfo == nil {
		panic("nil auth info")
	}
	bytes, err := cdc.Marshal(authInfo)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal auth info: %v", err))
	}
	return bytes
}

// getChainIdFromTx extracts chain ID from transaction (should be embedded in the unsigned tx)
func getChainIdFromTx(_ *tx.Tx) string {
	// For THORChain, the chain ID should be "thorchain-1" for mainnet
	// In a real implementation, this would be extracted from the transaction context
	// or provided as a parameter during transaction building
	return "thorchain-1"
}


// parseUint64 safely parses a string to uint64
func parseUint64(s string) (uint64, error) {
	result := uint64(0)
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, fmt.Errorf("invalid number: %s", s)
		}
		result = result*10 + uint64(char-'0')
	}
	return result, nil
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
