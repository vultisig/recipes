package solana

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/types"
)

// Chain represents the Solana blockchain.
type Chain struct{}

// NewChain creates a new Solana chain instance.
func NewChain() *Chain {
	return &Chain{}
}

// ID returns the unique identifier for the Solana chain.
func (c *Chain) ID() string {
	return "solana"
}

// Name returns a human-readable name for the Solana chain.
func (c *Chain) Name() string {
	return "Solana"
}

// Description returns a detailed description of the Solana chain.
func (c *Chain) Description() string {
	return "Solana is a high-performance blockchain supporting smart contracts and decentralized applications."
}

// SupportedProtocols returns the list of protocol IDs supported by Solana.
func (c *Chain) SupportedProtocols() []string {
	return []string{"sol"}
}

// ParsedSolanaTransaction wraps a decoded Solana transaction and implements types.DecodedTransaction.
type ParsedSolanaTransaction struct {
	tx *solana.Transaction
}

// GetTransaction returns the underlying Solana transaction.
func (p *ParsedSolanaTransaction) GetTransaction() *solana.Transaction {
	return p.tx
}

// GetInstructions returns the transaction instructions.
func (p *ParsedSolanaTransaction) GetInstructions() []solana.CompiledInstruction {
	return p.tx.Message.Instructions
}

// GetAccountKeys returns the account keys involved in the transaction.
func (p *ParsedSolanaTransaction) GetAccountKeys() []solana.PublicKey {
	return p.tx.Message.AccountKeys
}

// ChainIdentifier returns "solana".
func (p *ParsedSolanaTransaction) ChainIdentifier() string { return "solana" }

// Hash returns the transaction hash (first signature if present).
func (p *ParsedSolanaTransaction) Hash() string {
	if len(p.tx.Signatures) > 0 {
		return p.tx.Signatures[0].String()
	}
	return ""
}

// From returns the fee payer address.
func (p *ParsedSolanaTransaction) From() string {
	if len(p.tx.Message.AccountKeys) > 0 {
		return p.tx.Message.AccountKeys[0].String()
	}
	return ""
}

// To returns the recipient address (empty for Solana as transactions can have multiple recipients).
func (p *ParsedSolanaTransaction) To() string { return "" }

// Value returns zero (Solana uses lamports in instructions, not a single value field).
func (p *ParsedSolanaTransaction) Value() *big.Int { return big.NewInt(0) }

// Data returns empty bytes (Solana data is in instructions, not a single data field).
func (p *ParsedSolanaTransaction) Data() []byte { return nil }

// Nonce returns zero (Solana uses recent blockhash instead of nonces).
func (p *ParsedSolanaTransaction) Nonce() uint64 { return 0 }

// GasPrice returns zero (Solana uses compute units, not gas price).
func (p *ParsedSolanaTransaction) GasPrice() *big.Int { return big.NewInt(0) }

// GasLimit returns zero (Solana uses compute unit limits, not gas limit).
func (p *ParsedSolanaTransaction) GasLimit() uint64 { return 0 }

// ParseTransaction decodes a raw Solana transaction from hex string.
func (c *Chain) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	txBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex: %w", err)
	}

	tx, err := solana.TransactionFromDecoder(bin.NewBorshDecoder(txBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to decode Solana transaction: %w", err)
	}

	return &ParsedSolanaTransaction{tx: tx}, nil
}

// ParseTransactionBytes decodes a raw Solana transaction from bytes.
func (c *Chain) ParseTransactionBytes(txBytes []byte) (*ParsedSolanaTransaction, error) {
	tx, err := solana.TransactionFromDecoder(bin.NewBorshDecoder(txBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to decode Solana transaction: %w", err)
	}

	return &ParsedSolanaTransaction{tx: tx}, nil
}

// ComputeTxHash computes the transaction hash from the proposed transaction and signatures.
// For Solana, the transaction ID is the first signature (fee payer's signature).
func (c *Chain) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	if len(sigs) == 0 {
		return "", fmt.Errorf("at least one signature is required")
	}

	// Decode the transaction
	tx, err := solana.TransactionFromDecoder(bin.NewBorshDecoder(proposedTx))
	if err != nil {
		return "", fmt.Errorf("failed to decode transaction: %w", err)
	}

	// The number of signatures should match the number of required signers
	if len(sigs) != len(tx.Signatures) {
		return "", fmt.Errorf("signature count (%d) does not match required signers (%d)", len(sigs), len(tx.Signatures))
	}

	// Assemble and add signatures
	for i, sig := range sigs {
		// Parse R and S from hex (strip 0x prefix if present)
		rBytes, err := hex.DecodeString(cleanHex(sig.R))
		if err != nil {
			return "", fmt.Errorf("failed to decode R for signature %d: %w", i, err)
		}
		sBytes, err := hex.DecodeString(cleanHex(sig.S))
		if err != nil {
			return "", fmt.Errorf("failed to decode S for signature %d: %w", i, err)
		}

		// Validate R and S are exactly 32 bytes for Ed25519
		if len(rBytes) != 32 {
			return "", fmt.Errorf("r must be 32 bytes for signature %d, got %d", i, len(rBytes))
		}
		if len(sBytes) != 32 {
			return "", fmt.Errorf("s must be 32 bytes for signature %d, got %d", i, len(sBytes))
		}

		// Solana uses Ed25519 signatures which are 64 bytes (R || S)
		var solSig solana.Signature
		copy(solSig[:32], rBytes)
		copy(solSig[32:], sBytes)

		tx.Signatures[i] = solSig
	}

	// The transaction ID is the first signature (fee payer's signature)
	return tx.Signatures[0].String(), nil
}

// GetProtocol returns a protocol handler for the given ID.
func (c *Chain) GetProtocol(id string) (types.Protocol, error) {
	if id == "sol" {
		return NewSOL(), nil
	}
	return nil, fmt.Errorf("protocol %q not found on Solana", id)
}

func (c *Chain) ExtractTxBytes(txData string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(txData)
}

// cleanHex strips 0x/0X prefix and whitespace from a hex string.
func cleanHex(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		return s[2:]
	}
	return s
}

