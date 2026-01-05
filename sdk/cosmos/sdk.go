package cosmos

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

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

	chainCosmos "github.com/vultisig/recipes/chain/cosmos"
)

// SDK represents the Cosmos SDK for transaction signing and broadcasting
type SDK struct {
	rpcClient RPCClient
	cdc       codec.Codec
}

// NewSDK creates a new Cosmos SDK instance
func NewSDK(rpcClient RPCClient) *SDK {
	ir := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(ir)
	banktypes.RegisterInterfaces(ir)

	return &SDK{
		rpcClient: rpcClient,
		cdc:       codec.NewProtoCodec(ir),
	}
}

// Sign applies TSS signatures to an unsigned Cosmos transaction
func (sdk *SDK) Sign(unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse, pubKey []byte) ([]byte, error) {
	if len(signatures) == 0 {
		return nil, fmt.Errorf("no signatures provided")
	}
	if len(pubKey) != 33 {
		return nil, fmt.Errorf("pubkey must be 33 bytes (compressed), got %d", len(pubKey))
	}

	// Get the first (and typically only) signature
	var sig tss.KeysignResponse
	for _, v := range signatures {
		sig = v
		break
	}

	// Decode R and S from hex strings
	rHex := chainCosmos.CleanHex(sig.R)
	rBytes, err := hex.DecodeString(rHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode R: %w", err)
	}
	sHex := chainCosmos.CleanHex(sig.S)
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
	sLow, err := chainCosmos.NormalizeLowS(sBytes)
	if err != nil {
		return nil, fmt.Errorf("low-S normalization failed: %w", err)
	}

	// Unmarshal the unsigned transaction
	var unsignedTx tx.Tx
	if err := sdk.cdc.Unmarshal(unsignedTxBytes, &unsignedTx); err != nil {
		return nil, fmt.Errorf("failed to unmarshal unsigned transaction: %w", err)
	}

	// Create the secp256k1 public key
	cosmosPubKey := &secp256k1.PubKey{Key: pubKey}

	// Create signature bytes (r || s format for secp256k1)
	sigBytes := make([]byte, 64)
	copy(sigBytes[:32], rBytes)
	copy(sigBytes[32:], sLow)

	// Create SignerInfo with the public key
	pubKeyAny, err := codectypes.NewAnyWithValue(cosmosPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubkey any: %w", err)
	}

	// Extract sequence from existing SignerInfos if available
	var sequence uint64
	if unsignedTx.AuthInfo != nil && len(unsignedTx.AuthInfo.SignerInfos) > 0 {
		sequence = unsignedTx.AuthInfo.SignerInfos[0].Sequence
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
		Sequence: sequence,
	}

	// Preserve existing AuthInfo or create new one
	if unsignedTx.AuthInfo == nil {
		defaultCoins, err := sdk.sdkCoins(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create default coins: %w", err)
		}
		unsignedTx.AuthInfo = &tx.AuthInfo{
			Fee: &tx.Fee{
				Amount:   defaultCoins,
				GasLimit: 200000,
			},
		}
	}
	unsignedTx.AuthInfo.SignerInfos = []*tx.SignerInfo{signerInfo}

	// Set the signatures
	unsignedTx.Signatures = [][]byte{sigBytes}

	// Marshal the signed transaction
	signedTxBytes, err := sdk.cdc.Marshal(&unsignedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal signed transaction: %w", err)
	}

	return signedTxBytes, nil
}

// sdkCoins converts coins or returns empty coins
func (s *SDK) sdkCoins(coins []struct{ Denom, Amount string }) (cosmostypes.Coins, error) {
	if len(coins) == 0 {
		return cosmostypes.NewCoins(), nil
	}
	result := make(cosmostypes.Coins, 0, len(coins))
	for _, c := range coins {
		amount, ok := math.NewIntFromString(c.Amount)
		if !ok {
			return nil, fmt.Errorf("invalid amount %q for denom %s", c.Amount, c.Denom)
		}
		result = append(result, cosmostypes.NewCoin(c.Denom, amount))
	}
	return result, nil
}

// Broadcast submits a signed transaction to the Cosmos network
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

// ComputeTxHash computes the transaction hash for a signed Cosmos transaction
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


