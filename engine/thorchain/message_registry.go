package thorchain

import "fmt"

// MessageType represents the type of THORChain message
type MessageType int

const (
	MessageTypeSend MessageType = iota
	MessageTypeDeposit
)

// TypeUrl constants for message identification
const (
	// MsgSend TypeUrls - both standard Cosmos and THORChain custom formats
	TypeUrlCosmosMsgSend    = "/cosmos.bank.v1beta1.MsgSend"
	TypeUrlThorchainMsgSend = "/types.MsgSend"
	
	// MsgDeposit TypeUrls
	TypeUrlThorchainMsgDeposit = "/types.MsgDeposit"
)

// typeUrlRegistry maps TypeUrls to MessageTypes
var typeUrlRegistry = map[string]MessageType{
	TypeUrlCosmosMsgSend:       MessageTypeSend,
	TypeUrlThorchainMsgSend:    MessageTypeSend,
	TypeUrlThorchainMsgDeposit: MessageTypeDeposit,
}

// GetMessageTypeFromTypeUrl returns the MessageType for a given TypeUrl
func GetMessageTypeFromTypeUrl(typeUrl string) (MessageType, error) {
	if messageType, exists := typeUrlRegistry[typeUrl]; exists {
		return messageType, nil
	}
	return MessageType(0), fmt.Errorf("unsupported TypeUrl: %s", typeUrl)
}

// GetSupportedTypeUrls returns all supported TypeUrls
func GetSupportedTypeUrls() []string {
	urls := make([]string, 0, len(typeUrlRegistry))
	for url := range typeUrlRegistry {
		urls = append(urls, url)
	}
	return urls
}

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