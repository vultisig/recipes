package solana

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/vultisig/mobile-tss-lib/tss"
)

type RPCClient interface {
	SendTransaction(ctx context.Context, tx *solana.Transaction) (solana.Signature, error)
}

type HTTPRPCClient struct {
	client *rpc.Client
}

type SDK struct {
	rpcClient RPCClient
}

func NewSDK(rpcClient RPCClient) *SDK {
	return &SDK{
		rpcClient: rpcClient,
	}
}

func NewHTTPRPCClient(endpoint string) *HTTPRPCClient {
	return &HTTPRPCClient{
		client: rpc.New(endpoint),
	}
}

func (c *HTTPRPCClient) SendTransaction(ctx context.Context, tx *solana.Transaction) (solana.Signature, error) {
	sig, err := c.client.SendTransactionWithOpts(
		ctx,
		tx,
		rpc.TransactionOpts{
			SkipPreflight:       false,
			PreflightCommitment: rpc.CommitmentFinalized,
		},
	)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("send transaction: %w", err)
	}

	return sig, nil
}

func (sdk *SDK) Sign(unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse) ([]byte, error) {
	if len(signatures) == 0 {
		return nil, fmt.Errorf("no signatures provided")
	}

	tx, err := solana.TransactionFromBytes(unsignedTxBytes)
	if err != nil {
		return nil, fmt.Errorf("decode transaction: %w", err)
	}

	if tx.Message.Header.NumRequiredSignatures == 0 {
		return nil, fmt.Errorf("transaction requires no signatures")
	}

	messageBytes, err := tx.Message.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("marshal message: %w", err)
	}

	derivedKey := sdk.deriveKeyFromMessage(messageBytes)

	sigResponse, exists := signatures[derivedKey]
	if !exists {
		return nil, fmt.Errorf("missing signature for derived key: %s", derivedKey)
	}

	rHex := cleanHex(sigResponse.R)
	rBytes, err := hex.DecodeString(rHex)
	if err != nil {
		return nil, fmt.Errorf("decode R: %w", err)
	}

	sHex := cleanHex(sigResponse.S)
	sBytes, err := hex.DecodeString(sHex)
	if err != nil {
		return nil, fmt.Errorf("decode S: %w", err)
	}

	if len(rBytes) != 32 {
		return nil, fmt.Errorf("r must be 32 bytes, got %d", len(rBytes))
	}
	if len(sBytes) != 32 {
		return nil, fmt.Errorf("s must be 32 bytes, got %d", len(sBytes))
	}

	ed25519Sig := make([]byte, 64)
	copy(ed25519Sig[0:32], rBytes)
	copy(ed25519Sig[32:64], sBytes)

	var sig solana.Signature
	copy(sig[:], ed25519Sig)

	tx.Signatures = []solana.Signature{sig}

	signedTxBytes, err := tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("marshal signed transaction: %w", err)
	}

	return signedTxBytes, nil
}

func (sdk *SDK) Broadcast(ctx context.Context, signedTxBytes []byte) error {
	tx, err := solana.TransactionFromBytes(signedTxBytes)
	if err != nil {
		return fmt.Errorf("decode signed transaction: %w", err)
	}

	_, err = sdk.rpcClient.SendTransaction(ctx, tx)
	if err != nil {
		return fmt.Errorf("send transaction: %w", err)
	}

	return nil
}

func (sdk *SDK) Send(ctx context.Context, unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse) error {
	signedTxBytes, err := sdk.Sign(unsignedTxBytes, signatures)
	if err != nil {
		return fmt.Errorf("sign transaction: %w", err)
	}

	err = sdk.Broadcast(ctx, signedTxBytes)
	if err != nil {
		return fmt.Errorf("broadcast transaction: %w", err)
	}

	return nil
}

func (sdk *SDK) deriveKeyFromMessage(messageBytes []byte) string {
	hash := sha256.Sum256(messageBytes)

	return base64.StdEncoding.EncodeToString(hash[:])
}

func cleanHex(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		return s[2:]
	}
	return s
}
