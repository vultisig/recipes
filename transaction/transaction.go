package transaction

import (
	"encoding/json"
	"fmt"
	"time"
)

// ParamValue represents a parameter value with its type
type ParamValue struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// Transaction represents a cryptocurrency transaction to be evaluated against policies
type Transaction struct {
	// ID is a unique identifier for the transaction
	ID string `json:"id"`

	// ResourcePath is the fully qualified resource path (chain.protocol.function)
	ResourcePath string `json:"resource_path"`

	// ChainID is the identifier for the blockchain
	ChainID string `json:"chain_id"`

	// ProtocolID is the identifier for the protocol
	ProtocolID string `json:"protocol_id"`

	// FunctionID is the identifier for the function
	FunctionID string `json:"function_id"`

	// Parameters contains the transaction parameters
	Parameters map[string]ParamValue `json:"parameters"`

	// Value is the native asset value of the transaction (if applicable)
	Value string `json:"value,omitempty"`

	// Timestamp is when the transaction was created
	Timestamp time.Time `json:"timestamp"`
}

// NewTransaction creates a new transaction with the given resource path and parameters
func NewTransaction(chainID, protocolID, functionID string, params map[string]ParamValue) (*Transaction, error) {
	resourcePath := fmt.Sprintf("%s.%s.%s", chainID, protocolID, functionID)

	return &Transaction{
		ResourcePath: resourcePath,
		ChainID:      chainID,
		ProtocolID:   protocolID,
		FunctionID:   functionID,
		Parameters:   params,
		Timestamp:    time.Now(),
	}, nil
}

// GetParam returns a parameter value by name
func (t *Transaction) GetParam(name string) (ParamValue, bool) {
	val, ok := t.Parameters[name]
	return val, ok
}

// String returns a string representation of the transaction
func (t *Transaction) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error marshaling transaction: %v", err)
	}
	return string(b)
}

// Example of creating a BTC transfer transaction
func NewBTCTransfer(recipient string, amount string) (*Transaction, error) {
	params := map[string]ParamValue{
		"recipient": {
			Type:  "address",
			Value: recipient,
		},
		"amount": {
			Type:  "decimal",
			Value: amount,
		},
	}

	return NewTransaction("bitcoin", "btc", "transfer", params)
}
