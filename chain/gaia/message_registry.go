package gaia

import "fmt"

// MessageType represents the type of Cosmos message
type MessageType int

const (
	MessageTypeSend MessageType = iota
)

// TypeUrl constants for message identification
const (
	// MsgSend TypeUrl - standard Cosmos bank module
	TypeUrlCosmosMsgSend = "/cosmos.bank.v1beta1.MsgSend"
)

// typeUrlRegistry maps TypeUrls to MessageTypes
var typeUrlRegistry = map[string]MessageType{
	TypeUrlCosmosMsgSend: MessageTypeSend,
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
	default:
		return "Unknown"
	}
}

