package solana

import (
	"encoding/hex"
	"fmt"
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

// ParsedSolanaTransaction wraps a decoded Solana transaction.
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

// ParseTransaction decodes a raw Solana transaction from hex string.
func (c *Chain) ParseTransaction(txHex string) (*ParsedSolanaTransaction, error) {
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
		// Parse R and S from hex
		rBytes, err := hex.DecodeString(sig.R)
		if err != nil {
			return "", fmt.Errorf("failed to decode R for signature %d: %w", i, err)
		}
		sBytes, err := hex.DecodeString(sig.S)
		if err != nil {
			return "", fmt.Errorf("failed to decode S for signature %d: %w", i, err)
		}

		// Solana uses Ed25519 signatures which are 64 bytes (R || S)
		// Pad R and S to 32 bytes each if needed
		rPadded := make([]byte, 32)
		sPadded := make([]byte, 32)
		copy(rPadded[32-len(rBytes):], rBytes)
		copy(sPadded[32-len(sBytes):], sBytes)

		// Concatenate R || S to form the 64-byte signature
		var solSig solana.Signature
		copy(solSig[:32], rPadded)
		copy(solSig[32:], sPadded)

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

