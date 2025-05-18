package bitcoin

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// BTC implements the Protocol interface for the Bitcoin protocol
type BTC struct{}

// ID returns the unique identifier for the BTC protocol
func (b *BTC) ID() string {
	return "btc"
}

// Name returns a human-readable name for the BTC protocol
func (b *BTC) Name() string {
	return "Bitcoin"
}

// ChainID returns the ID of the chain this protocol belongs to
func (b *BTC) ChainID() string {
	return "bitcoin"
}

// Description returns a detailed description of the BTC protocol
func (b *BTC) Description() string {
	return "The native cryptocurrency of the Bitcoin blockchain, used for transactions and value transfer."
}

// Functions returns a list of available functions for this protocol
func (b *BTC) Functions() []*types.Function {
	return []*types.Function{
		{
			ID:          "transfer",
			Name:        "Transfer BTC",
			Description: "Transfer Bitcoin to another address",
			Parameters: []*types.FunctionParam{
				{
					Name:        "recipient",
					Type:        "address",
					Description: "The Bitcoin address of the recipient",
				},
				{
					Name:        "amount",
					Type:        "decimal",
					Description: "The amount of Bitcoin to transfer",
				},
			},
		},
	}
}

// GetFunction retrieves a specific function by ID
func (b *BTC) GetFunction(id string) (*types.Function, error) {
	for _, fn := range b.Functions() {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found for protocol BTC", id)
}

func (b *BTC) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	return false, nil, fmt.Errorf("function matching not supported on Bitcoin")
}

// NewBTC creates a new BTC protocol instance
func NewBTC() types.Protocol {
	return &BTC{}
}
