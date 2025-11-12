package thorchain

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
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

// secp256k1 curve parameters for low-S enforcement
var secpN, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
var secpHalfN = new(big.Int).Rsh(new(big.Int).Set(secpN), 1)

// NewSDK creates a new THORChain SDK instance
func NewSDK(rpcClient RPCClient) *SDK {
	return &SDK{
		rpcClient: rpcClient,
		codec:     MakeCodec(),
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
		codec:  MakeCodec(),
	}, nil
}

// MakeCodec creates a standardized codec for THORChain transactions.
// This codec is used across SDK, engine, and tests to ensure consistency.
// External plugins can use this to maintain compatibility with the recipes engine.
func MakeCodec() codec.Codec {
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

// Sign applies TSS signatures to an unsigned THORChain transaction.
//
// REQUIREMENTS: The unsigned transaction MUST have properly formatted AuthInfo with:
// - SignerInfos[0].PublicKey set to a valid secp256k1 public key (33 bytes compressed)
// - SignerInfos[0].ModeInfo set to SIGN_MODE_DIRECT
// - SignerInfos[0].Sequence matching the sequence used in MessageHash()
//
// If these requirements are not met, the signed transaction will be rejected by the network.
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

	// Enforce low-S to prevent transaction malleability
	sLow, err := normalizeLowS(sBytes)
	if err != nil {
		return nil, fmt.Errorf("low-S normalization failed: %w", err)
	}
	sBytes = sLow

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

	// Create the sign document following Cosmos SDK SIGN_MODE_DIRECT
	bodyBytes := mustMarshalTxBody(sdk.codec, unsignedTx.Body)
	authInfoBytes := mustMarshalAuthInfo(sdk.codec, unsignedTx.AuthInfo)

	signDoc := &tx.SignDoc{
		BodyBytes:     bodyBytes,
		AuthInfoBytes: authInfoBytes,
		ChainId:       "thorchain-1", // THORChain mainnet chain ID
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

// cleanHex removes 0x prefix from hex strings
func cleanHex(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		return s[2:]
	}
	return s
}

// normalizeLowS ensures the S value is in the lower half of the curve order to prevent malleability
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
