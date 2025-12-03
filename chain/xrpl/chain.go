package xrpl

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/types"
	xrpgo "github.com/xyield/xrpl-go/binary-codec"
	"github.com/xyield/xrpl-go/model/transactions"
	xrptypes "github.com/xyield/xrpl-go/model/transactions/types"
)

// Chain implements the types.Chain interface for XRPL.
type Chain struct{}

// NewChain creates a new XRPL chain instance.
func NewChain() *Chain {
	return &Chain{}
}

// ID returns the unique identifier for the XRPL chain.
func (c *Chain) ID() string {
	return "xrpl"
}

// Name returns a human-readable name for the XRPL chain.
func (c *Chain) Name() string {
	return "XRP Ledger"
}

// Description returns a detailed description of the XRPL chain.
func (c *Chain) Description() string {
	return "XRP Ledger is a decentralized public blockchain for fast, low-cost payments."
}

// SupportedProtocols returns the list of protocol IDs supported by XRPL.
func (c *Chain) SupportedProtocols() []string {
	return []string{"xrp"}
}

// ParsedXRPLTransaction implements the types.DecodedTransaction interface for XRPL.
type ParsedXRPLTransaction struct {
	payment *transactions.Payment
}

// ChainIdentifier returns "xrpl".
func (p *ParsedXRPLTransaction) ChainIdentifier() string {
	return "xrpl"
}

// Hash returns the transaction hash (empty for unsigned transactions).
func (p *ParsedXRPLTransaction) Hash() string {
	return ""
}

// From returns the sender account address.
func (p *ParsedXRPLTransaction) From() string {
	return string(p.payment.Account)
}

// To returns the destination address.
func (p *ParsedXRPLTransaction) To() string {
	return string(p.payment.Destination)
}

// Value returns the amount being transferred.
func (p *ParsedXRPLTransaction) Value() *big.Int {
	if p.payment.Amount == nil {
		return nil
	}
	xrpAmount, ok := p.payment.Amount.(xrptypes.XRPCurrencyAmount)
	if !ok {
		return nil
	}
	return big.NewInt(int64(xrpAmount))
}

// Data returns nil as XRPL uses structured transactions.
func (p *ParsedXRPLTransaction) Data() []byte {
	return nil
}

// Nonce returns the sequence number.
func (p *ParsedXRPLTransaction) Nonce() uint64 {
	return uint64(p.payment.Sequence)
}

// GasPrice returns nil as XRPL uses a different fee model.
func (p *ParsedXRPLTransaction) GasPrice() *big.Int {
	return nil
}

// GasLimit returns 0 as XRPL uses a different fee model.
func (p *ParsedXRPLTransaction) GasLimit() uint64 {
	return 0
}

// GetPayment returns the underlying XRPL Payment transaction.
func (p *ParsedXRPLTransaction) GetPayment() *transactions.Payment {
	return p.payment
}

// GetMemo extracts memo data from the payment transaction.
func (p *ParsedXRPLTransaction) GetMemo() (string, error) {
	return ExtractMemoFromXRPPayment(p.payment)
}

// ExtractMemoFromXRPPayment extracts memo data from XRPL Payment transaction.
func ExtractMemoFromXRPPayment(payment *transactions.Payment) (string, error) {
	if len(payment.Memos) == 0 {
		return "", fmt.Errorf("no memo found in payment transaction")
	}

	// XRPL memos are typically hex-encoded, need to decode
	memo := payment.Memos[0]
	if memo.Memo.MemoData == "" {
		return "", fmt.Errorf("empty memo data")
	}

	// Decode hex to string (THORChain memos are text)
	memoBytes, err := hex.DecodeString(memo.Memo.MemoData)
	if err != nil {
		// If not hex, treat as plain string
		return memo.Memo.MemoData, nil
	}

	return string(memoBytes), nil
}

// ParseTransaction decodes a raw XRPL transaction from hex string.
func (c *Chain) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	txBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex: %w", err)
	}

	return c.ParseTransactionBytes(txBytes)
}

// ParseTransactionBytes decodes a raw XRPL transaction from bytes.
func (c *Chain) ParseTransactionBytes(txBytes []byte) (types.DecodedTransaction, error) {
	// Convert bytes to hex string for binary codec
	hexStr := hex.EncodeToString(txBytes)

	// Use XRPL binary codec to decode hex to JSON
	jsonData, err := xrpgo.Decode(hexStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode XRPL binary format: %w", err)
	}

	// Convert map to JSON bytes for unmarshaling
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal decoded JSON: %w", err)
	}

	// Unmarshal into Payment struct
	var payment transactions.Payment
	if err := json.Unmarshal(jsonBytes, &payment); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XRPL Payment transaction: %w", err)
	}

	// Validate required fields for Payment transactions
	if string(payment.Account) == "" {
		return nil, fmt.Errorf("account field is required")
	}
	if string(payment.Destination) == "" {
		return nil, fmt.Errorf("destination field is required")
	}
	if payment.Amount == nil {
		return nil, fmt.Errorf("amount field is required")
	}

	return &ParsedXRPLTransaction{payment: &payment}, nil
}

// GetProtocol returns a protocol handler for the given ID.
func (c *Chain) GetProtocol(id string) (types.Protocol, error) {
	if id == "xrp" {
		return NewXRP(), nil
	}
	return nil, fmt.Errorf("protocol %q not found on XRPL", id)
}

// ComputeTxHash computes the transaction hash from the proposed transaction and signatures.
func (c *Chain) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	// XRPL transaction hash computation would require proper signature assembly
	// For now, return an error as this needs chain-specific implementation
	return "", fmt.Errorf("ComputeTxHash not yet implemented for XRPL")
}

