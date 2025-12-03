package thorchain

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

// Chain implements the types.Chain interface for THORChain.
type Chain struct {
	cdc codec.Codec
}

// NewChain creates a new THORChain chain instance.
func NewChain() *Chain {
	ir := codectypes.NewInterfaceRegistry()

	// Register crypto types (required for PubKey interfaces)
	cryptocodec.RegisterInterfaces(ir)

	// Register bank message types
	banktypes.RegisterInterfaces(ir)

	// Register the generated protobuf MsgDeposit for THORChain swaps
	ir.RegisterImplementations((*sdk.Msg)(nil), &types.MsgDeposit{})

	return &Chain{cdc: codec.NewProtoCodec(ir)}
}

// ID returns the unique identifier for the THORChain chain.
func (c *Chain) ID() string {
	return "thorchain"
}

// Name returns a human-readable name for the THORChain chain.
func (c *Chain) Name() string {
	return "THORChain"
}

// Description returns a detailed description of the THORChain chain.
func (c *Chain) Description() string {
	return "THORChain is a decentralized liquidity protocol for cross-chain swaps."
}

// SupportedProtocols returns the list of protocol IDs supported by THORChain.
func (c *Chain) SupportedProtocols() []string {
	return []string{"rune", "send", "thorchain_swap"}
}

// ParsedThorchainTransaction implements the types.DecodedTransaction interface for THORChain.
type ParsedThorchainTransaction struct {
	tx  *tx.Tx
	cdc codec.Codec
}

// ChainIdentifier returns "thorchain".
func (p *ParsedThorchainTransaction) ChainIdentifier() string {
	return "thorchain"
}

// Hash returns the transaction hash (empty for unsigned transactions).
func (p *ParsedThorchainTransaction) Hash() string {
	return ""
}

// From returns the sender address from the first message.
func (p *ParsedThorchainTransaction) From() string {
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
func (p *ParsedThorchainTransaction) To() string {
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

// Value returns nil as THORChain transactions handle value differently.
func (p *ParsedThorchainTransaction) Value() *big.Int {
	return nil
}

// Data returns nil as THORChain uses protobuf messages.
func (p *ParsedThorchainTransaction) Data() []byte {
	return nil
}

// Nonce returns 0 as THORChain uses sequence numbers differently.
func (p *ParsedThorchainTransaction) Nonce() uint64 {
	return 0
}

// GasPrice returns nil as THORChain uses a different fee model.
func (p *ParsedThorchainTransaction) GasPrice() *big.Int {
	return nil
}

// GasLimit returns 0 as THORChain uses a different fee model.
func (p *ParsedThorchainTransaction) GasLimit() uint64 {
	return 0
}

// GetTransaction returns the underlying Cosmos SDK transaction.
func (p *ParsedThorchainTransaction) GetTransaction() *tx.Tx {
	return p.tx
}

// GetMemo returns the transaction memo.
func (p *ParsedThorchainTransaction) GetMemo() string {
	if p.tx.Body != nil {
		return p.tx.Body.Memo
	}
	return ""
}

// ParseTransaction decodes a raw THORChain transaction from hex string.
func (c *Chain) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	txBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex: %w", err)
	}

	return c.ParseTransactionBytes(txBytes)
}

// ParseTransactionBytes decodes a raw THORChain transaction from bytes.
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

	return &ParsedThorchainTransaction{tx: &txData, cdc: c.cdc}, nil
}

// ComputeTxHash computes the transaction hash from the proposed transaction and signatures.
func (c *Chain) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	// THORChain transaction hash computation would require proper signature assembly
	// For now, return an error as this needs chain-specific implementation
	return "", fmt.Errorf("ComputeTxHash not yet implemented for THORChain")
}

// GetProtocol returns a protocol handler for the given ID.
func (c *Chain) GetProtocol(id string) (types.Protocol, error) {
	switch id {
	case "rune", "send":
		return NewRUNE(), nil
	case "thorchain_swap":
		return NewSwap(), nil
	}
	return nil, fmt.Errorf("protocol %q not found on THORChain", id)
}

