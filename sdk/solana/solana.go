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
	sdk "github.com/vultisig/recipes/sdk"
)

type RPCClient interface {
	SendTransaction(ctx context.Context, tx *solana.Transaction) (solana.Signature, error)
	GetAccountInfo(ctx context.Context, account solana.PublicKey) (*rpc.GetAccountInfoResult, error)
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

func (c *HTTPRPCClient) GetAccountInfo(ctx context.Context, account solana.PublicKey) (*rpc.GetAccountInfoResult, error) {
	return c.client.GetAccountInfo(ctx, account)
}

func (sdk *SDK) Sign(unsignedTxBytes []byte, signatures map[string]tss.KeysignResponse) ([]byte, error) {
	if len(signatures) != 1 {
		return nil, fmt.Errorf("must be 1 signature, got %d", len(signatures))
	}

	tx, err := solana.TransactionFromBytes(unsignedTxBytes)
	if err != nil {
		return nil, fmt.Errorf("decode transaction: %w", err)
	}

	if tx.Message.Header.NumRequiredSignatures == 0 {
		return nil, fmt.Errorf("unexpected no signatures")
	}
	if tx.Message.Header.NumRequiredSignatures > 1 {
		return nil, fmt.Errorf("multi-signature transactions are not supported")
	}

	var sigResponse tss.KeysignResponse
	for _, v := range signatures {
		sigResponse = v
		break
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

func (sdk *SDK) MessageHash(unsignedTx []byte) ([]byte, error) {
	tx, err := solana.TransactionFromBytes(unsignedTx)
	if err != nil {
		return nil, fmt.Errorf("decode transaction: %w", err)
	}

	messageBytes, err := tx.Message.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("marshal message: %w", err)
	}

	return messageBytes, nil
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

// GetTokenProgram queries the mint account to determine which token program owns it.
// Returns the token program public key string (TokenProgramID for legacy SPL tokens or
// Token2022ProgramID for Token-2022 tokens).
// If mint is empty or SOL (native), returns empty string (no token program needed).
func (sdk *SDK) GetTokenProgram(ctx context.Context, mint string) (string, error) {
	// Native SOL doesn't have a token program
	if mint == "" || mint == "SOL" || mint == "So11111111111111111111111111111111111111112" {
		return "", nil
	}

	mintPubkey, err := solana.PublicKeyFromBase58(mint)
	if err != nil {
		return "", fmt.Errorf("invalid mint address: %w", err)
	}

	accountInfo, err := sdk.rpcClient.GetAccountInfo(ctx, mintPubkey)
	if err != nil {
		return "", fmt.Errorf("failed to get mint account info: %w", err)
	}

	if accountInfo.Value == nil {
		return "", fmt.Errorf("mint account not found: %s", mint)
	}

	owner := accountInfo.Value.Owner
	if owner == solana.TokenProgramID {
		return solana.TokenProgramID.String(), nil
	}
	if owner == solana.Token2022ProgramID {
		return solana.Token2022ProgramID.String(), nil
	}

	return "", fmt.Errorf("mint account is not owned by a token program: %s", owner)
}

// DeriveSigningHashes derives the signing hash from unsigned transaction bytes.
// For Solana, this extracts the message bytes which are signed directly.
// Returns a single DerivedHash since Solana transactions have one signature per signer.
func (s *SDK) DeriveSigningHashes(txBytes []byte, _ sdk.DeriveOptions) ([]sdk.DerivedHash, error) {
	messageBytes, err := s.MessageHash(txBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to derive message hash: %w", err)
	}

	// For Solana, the message bytes are signed directly (Ed25519)
	// The Hash field is sha256 of the message for lookup purposes
	hash := sha256.Sum256(messageBytes)

	return []sdk.DerivedHash{
		{
			Message: messageBytes,
			Hash:    hash[:],
		},
	}, nil
}
