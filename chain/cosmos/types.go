package cosmos

import (
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
)

// MessageType represents the type of Cosmos message
type MessageType int

const (
	MessageTypeUnknown MessageType = iota - 1
	MessageTypeSend
	MessageTypeDeposit
)

// String returns string representation of MessageType
func (mt MessageType) String() string {
	switch mt {
	case MessageTypeSend:
		return "MsgSend"
	case MessageTypeDeposit:
		return "MsgDeposit"
	default:
		return "Unknown"
	}
}

// TypeUrl constants for message identification
const (
	// MsgSend TypeUrls - standard Cosmos bank module
	TypeUrlCosmosMsgSend = "/cosmos.bank.v1beta1.MsgSend"

	// Custom chain-specific MsgSend/MsgDeposit TypeUrls
	TypeUrlCustomMsgSend    = "/types.MsgSend"
	TypeUrlCustomMsgDeposit = "/types.MsgDeposit"
)

// MessageTypeRegistry maps TypeUrls to MessageTypes for a chain
type MessageTypeRegistry struct {
	typeUrls map[string]MessageType
}

// NewMessageTypeRegistry creates a new registry with the given TypeUrl mappings
func NewMessageTypeRegistry(mappings map[string]MessageType) *MessageTypeRegistry {
	return &MessageTypeRegistry{typeUrls: mappings}
}

// GetMessageType returns the MessageType for a given TypeUrl
func (r *MessageTypeRegistry) GetMessageType(typeUrl string) (MessageType, error) {
	if mt, exists := r.typeUrls[typeUrl]; exists {
		return mt, nil
	}
	return MessageTypeUnknown, fmt.Errorf("unsupported TypeUrl: %s", typeUrl)
}

// GetSupportedTypeUrls returns all supported TypeUrls
func (r *MessageTypeRegistry) GetSupportedTypeUrls() []string {
	urls := make([]string, 0, len(r.typeUrls))
	for url := range r.typeUrls {
		urls = append(urls, url)
	}
	return urls
}

// ParsedCosmosTransaction implements the types.DecodedTransaction interface for Cosmos chains.
type ParsedCosmosTransaction struct {
	Tx           *tx.Tx
	Cdc          codec.Codec
	ChainID      string
	Bech32Prefix string
	// FromExtractor is an optional custom function to extract the From address.
	// If nil, the default extraction logic is used.
	FromExtractor func(*tx.Tx, codec.Codec, string) string
}

// ChainIdentifier returns the chain identifier.
func (p *ParsedCosmosTransaction) ChainIdentifier() string {
	return p.ChainID
}

// Hash returns the transaction hash (empty for unsigned transactions).
func (p *ParsedCosmosTransaction) Hash() string {
	return ""
}

// Value returns nil as Cosmos transactions handle value differently.
func (p *ParsedCosmosTransaction) Value() *big.Int {
	return nil
}

// Data returns nil as Cosmos uses protobuf messages.
func (p *ParsedCosmosTransaction) Data() []byte {
	return nil
}

// Nonce returns 0 as Cosmos uses sequence numbers differently.
func (p *ParsedCosmosTransaction) Nonce() uint64 {
	return 0
}

// GasPrice returns nil as Cosmos uses a different fee model.
func (p *ParsedCosmosTransaction) GasPrice() *big.Int {
	return nil
}

// GasLimit returns 0 as Cosmos uses a different fee model.
func (p *ParsedCosmosTransaction) GasLimit() uint64 {
	return 0
}

// GetTransaction returns the underlying Cosmos SDK transaction.
func (p *ParsedCosmosTransaction) GetTransaction() *tx.Tx {
	return p.Tx
}

// GetMemo returns the transaction memo.
func (p *ParsedCosmosTransaction) GetMemo() string {
	if p.Tx != nil && p.Tx.Body != nil {
		return p.Tx.Body.Memo
	}
	return ""
}

