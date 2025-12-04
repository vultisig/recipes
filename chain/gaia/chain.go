package gaia

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/vultisig/mobile-tss-lib/tss"
	"github.com/vultisig/recipes/types"
)

// Chain implements the types.Chain interface for Cosmos (GAIA).
type Chain struct {
	cdc codec.Codec
}

// NewChain creates a new Cosmos/GAIA chain instance.
func NewChain() *Chain {
	ir := codectypes.NewInterfaceRegistry()

	// Register crypto types (required for PubKey interfaces)
	cryptocodec.RegisterInterfaces(ir)

	// Register bank message types
	banktypes.RegisterInterfaces(ir)

	return &Chain{cdc: codec.NewProtoCodec(ir)}
}

// ID returns the unique identifier for the Cosmos chain.
func (c *Chain) ID() string {
	return "cosmos"
}

// Name returns a human-readable name for the Cosmos chain.
func (c *Chain) Name() string {
	return "Cosmos"
}

// Description returns a detailed description of the Cosmos chain.
func (c *Chain) Description() string {
	return "Cosmos (GAIA) is the hub of the Cosmos network, enabling cross-chain communication via IBC."
}

// SupportedProtocols returns the list of protocol IDs supported by Cosmos.
func (c *Chain) SupportedProtocols() []string {
	return []string{"atom", "send"}
}

// ParsedGaiaTransaction implements the types.DecodedTransaction interface for Cosmos.
type ParsedGaiaTransaction struct {
	tx  *tx.Tx
	cdc codec.Codec
}

// ChainIdentifier returns "cosmos".
func (p *ParsedGaiaTransaction) ChainIdentifier() string {
	return "cosmos"
}

// Hash returns the transaction hash (empty for unsigned transactions).
func (p *ParsedGaiaTransaction) Hash() string {
	return ""
}

// From returns the sender address from the first message.
func (p *ParsedGaiaTransaction) From() string {
	if p.tx.Body == nil || len(p.tx.Body.Messages) == 0 {
		return ""
	}
	// Try to extract from MsgSend
	msg := p.tx.Body.Messages[0]
	var sdkMsg sdk.Msg
	if err := p.cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return ""
	}
	if msgSend, ok := sdkMsg.(*banktypes.MsgSend); ok {
		return msgSend.FromAddress
	}
	return ""
}

// To returns the recipient address from the first message.
func (p *ParsedGaiaTransaction) To() string {
	if p.tx.Body == nil || len(p.tx.Body.Messages) == 0 {
		return ""
	}
	// Try to extract from MsgSend
	msg := p.tx.Body.Messages[0]
	var sdkMsg sdk.Msg
	if err := p.cdc.UnpackAny(msg, &sdkMsg); err != nil {
		return ""
	}
	if msgSend, ok := sdkMsg.(*banktypes.MsgSend); ok {
		return msgSend.ToAddress
	}
	return ""
}

// Value returns nil as Cosmos transactions handle value differently.
func (p *ParsedGaiaTransaction) Value() *big.Int {
	return nil
}

// Data returns nil as Cosmos uses protobuf messages.
func (p *ParsedGaiaTransaction) Data() []byte {
	return nil
}

// Nonce returns 0 as Cosmos uses sequence numbers differently.
func (p *ParsedGaiaTransaction) Nonce() uint64 {
	return 0
}

// GasPrice returns nil as Cosmos uses a different fee model.
func (p *ParsedGaiaTransaction) GasPrice() *big.Int {
	return nil
}

// GasLimit returns 0 as Cosmos uses a different fee model.
func (p *ParsedGaiaTransaction) GasLimit() uint64 {
	return 0
}

// GetTransaction returns the underlying Cosmos SDK transaction.
func (p *ParsedGaiaTransaction) GetTransaction() *tx.Tx {
	return p.tx
}

// GetMemo returns the transaction memo.
func (p *ParsedGaiaTransaction) GetMemo() string {
	if p.tx.Body != nil {
		return p.tx.Body.Memo
	}
	return ""
}

// ParseTransaction decodes a raw Cosmos transaction from hex string.
func (c *Chain) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	txBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex: %w", err)
	}

	return c.ParseTransactionBytes(txBytes)
}

// ParseTransactionBytes decodes a raw Cosmos transaction from bytes.
func (c *Chain) ParseTransactionBytes(txBytes []byte) (types.DecodedTransaction, error) {
	const maxTxBytes = 32 * 1024 // 32 KB
	if len(txBytes) > maxTxBytes {
		return nil, fmt.Errorf("transaction too large: %d bytes (max %d)", len(txBytes), maxTxBytes)
	}

	if len(txBytes) == 0 {
		return nil, fmt.Errorf("empty transaction data")
	}

	var txData tx.Tx
	if err := c.cdc.Unmarshal(txBytes, &txData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal protobuf transaction: %w", err)
	}

	return &ParsedGaiaTransaction{tx: &txData, cdc: c.cdc}, nil
}

// ComputeTxHash computes the transaction hash from the proposed transaction and signatures.
func (c *Chain) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	// Cosmos transaction hash computation would require proper signature assembly
	return "", fmt.Errorf("computeTxHash not yet implemented for Cosmos")
}

// GetProtocol returns a protocol handler for the given ID.
func (c *Chain) GetProtocol(id string) (types.Protocol, error) {
	switch id {
	case "atom", "send":
		return NewATOM(), nil
	}
	return nil, fmt.Errorf("protocol %q not found on Cosmos", id)
}

