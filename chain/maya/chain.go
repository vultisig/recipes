package maya

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

// Chain implements the types.Chain interface for MAYAChain.
type Chain struct {
	cdc codec.Codec
}

// NewChain creates a new MAYAChain chain instance.
func NewChain() *Chain {
	ir := codectypes.NewInterfaceRegistry()

	// Register crypto types (required for PubKey interfaces)
	cryptocodec.RegisterInterfaces(ir)

	// Register bank message types
	banktypes.RegisterInterfaces(ir)

	// Register the generated protobuf MsgDeposit for MAYAChain swaps
	ir.RegisterImplementations((*sdk.Msg)(nil), &types.MsgDeposit{})

	return &Chain{cdc: codec.NewProtoCodec(ir)}
}

// ID returns the unique identifier for the MAYAChain chain.
func (c *Chain) ID() string {
	return "mayachain"
}

// Name returns a human-readable name for the MAYAChain chain.
func (c *Chain) Name() string {
	return "MAYAChain"
}

// Description returns a detailed description of the MAYAChain chain.
func (c *Chain) Description() string {
	return "MAYAChain is a decentralized cross-chain liquidity protocol forked from THORChain."
}

// SupportedProtocols returns the list of protocol IDs supported by MAYAChain.
func (c *Chain) SupportedProtocols() []string {
	return []string{"cacao", "send", "mayachain_swap"}
}

// ParsedMayaTransaction implements the types.DecodedTransaction interface for MAYAChain.
type ParsedMayaTransaction struct {
	tx  *tx.Tx
	cdc codec.Codec
}

// ChainIdentifier returns "mayachain".
func (p *ParsedMayaTransaction) ChainIdentifier() string {
	return "mayachain"
}

// Hash returns the transaction hash (empty for unsigned transactions).
func (p *ParsedMayaTransaction) Hash() string {
	return ""
}

// From returns the sender address from the first message.
func (p *ParsedMayaTransaction) From() string {
	if p.tx.Body == nil || len(p.tx.Body.Messages) == 0 {
		return ""
	}
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
func (p *ParsedMayaTransaction) To() string {
	if p.tx.Body == nil || len(p.tx.Body.Messages) == 0 {
		return ""
	}
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

// Value returns nil as MAYAChain transactions handle value differently.
func (p *ParsedMayaTransaction) Value() *big.Int {
	return nil
}

// Data returns nil as MAYAChain uses protobuf messages.
func (p *ParsedMayaTransaction) Data() []byte {
	return nil
}

// Nonce returns 0 as MAYAChain uses sequence numbers differently.
func (p *ParsedMayaTransaction) Nonce() uint64 {
	return 0
}

// GasPrice returns nil as MAYAChain uses a different fee model.
func (p *ParsedMayaTransaction) GasPrice() *big.Int {
	return nil
}

// GasLimit returns 0 as MAYAChain uses a different fee model.
func (p *ParsedMayaTransaction) GasLimit() uint64 {
	return 0
}

// GetTransaction returns the underlying Cosmos SDK transaction.
func (p *ParsedMayaTransaction) GetTransaction() *tx.Tx {
	return p.tx
}

// GetMemo returns the transaction memo.
func (p *ParsedMayaTransaction) GetMemo() string {
	if p.tx.Body != nil {
		return p.tx.Body.Memo
	}
	return ""
}

// ParseTransaction decodes a raw MAYAChain transaction from hex string.
func (c *Chain) ParseTransaction(txHex string) (types.DecodedTransaction, error) {
	txBytes, err := hex.DecodeString(strings.TrimPrefix(txHex, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex: %w", err)
	}

	return c.ParseTransactionBytes(txBytes)
}

// ParseTransactionBytes decodes a raw MAYAChain transaction from bytes.
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

	return &ParsedMayaTransaction{tx: &txData, cdc: c.cdc}, nil
}

// ComputeTxHash computes the transaction hash from the proposed transaction and signatures.
func (c *Chain) ComputeTxHash(proposedTx []byte, sigs []tss.KeysignResponse) (string, error) {
	return "", fmt.Errorf("computeTxHash not yet implemented for MAYAChain")
}

// GetProtocol returns a protocol handler for the given ID.
func (c *Chain) GetProtocol(id string) (types.Protocol, error) {
	switch id {
	case "cacao", "send":
		return NewCACAO(), nil
	case "mayachain_swap":
		return NewSwap(), nil
	}
	return nil, fmt.Errorf("protocol %q not found on MAYAChain", id)
}

