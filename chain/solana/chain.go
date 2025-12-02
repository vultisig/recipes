package solana

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/types"
)

// Chain implements the types.Chain interface for Solana.
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

// ParsedSolanaTransaction implements the types.DecodedTransaction interface for Solana.
type ParsedSolanaTransaction struct {
	tx *solana.Transaction
}

// ChainIdentifier returns "solana".
func (p *ParsedSolanaTransaction) ChainIdentifier() string {
	return "solana"
}

// Hash returns the transaction signature (first signature).
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

// To returns empty string as Solana transactions don't have a single recipient.
func (p *ParsedSolanaTransaction) To() string {
	return ""
}

// Value returns nil as Solana transactions handle value differently.
func (p *ParsedSolanaTransaction) Value() *big.Int {
	return nil
}

// Data returns the raw transaction data.
func (p *ParsedSolanaTransaction) Data() []byte {
	return nil
}

// Nonce returns 0 as Solana uses a different mechanism (recent blockhash).
func (p *ParsedSolanaTransaction) Nonce() uint64 {
	return 0
}

// GasPrice returns nil as Solana uses compute units and priority fees.
func (p *ParsedSolanaTransaction) GasPrice() *big.Int {
	return nil
}

// GasLimit returns 0 as Solana uses compute units.
func (p *ParsedSolanaTransaction) GasLimit() uint64 {
	return 0
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
func (c *Chain) ParseTransactionBytes(txBytes []byte) (types.DecodedTransaction, error) {
	tx, err := solana.TransactionFromDecoder(bin.NewBorshDecoder(txBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to decode Solana transaction: %w", err)
	}

	return &ParsedSolanaTransaction{tx: tx}, nil
}

// ComputeTxHash computes the transaction hash from the proposed transaction and signatures.
func (c *Chain) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	// Solana transaction hash computation would require proper signature assembly
	// For now, return an error as this needs chain-specific implementation
	return "", fmt.Errorf("ComputeTxHash not yet implemented for Solana")
}

// GetProtocol returns a protocol handler for the given ID.
func (c *Chain) GetProtocol(id string) (types.Protocol, error) {
	if id == "sol" {
		return NewSOL(), nil
	}
	return nil, fmt.Errorf("protocol %q not found on Solana", id)
}

