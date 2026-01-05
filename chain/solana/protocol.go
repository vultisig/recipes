package solana

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// SOL implements the native Solana protocol.
type SOL struct {
	id          string
	name        string
	description string
	functions   []*types.Function
}

// NewSOL creates a new SOL protocol instance.
func NewSOL() *SOL {
	return &SOL{
		id:          "sol",
		name:        "Solana",
		description: "Native SOL currency of the Solana blockchain",
		functions: []*types.Function{
			{
				ID:          "transfer",
				Name:        "Transfer SOL",
				Description: "Transfer SOL to another address",
				Parameters: []*types.FunctionParam{
					{Name: "recipient", Type: "address", Description: "The Solana address of the recipient"},
					{Name: "amount", Type: "decimal", Description: "The amount of SOL to transfer"},
				},
			},
		},
	}
}

// ID returns the protocol identifier.
func (p *SOL) ID() string {
	return p.id
}

// Name returns the protocol name.
func (p *SOL) Name() string {
	return p.name
}

// ChainID returns the chain identifier.
func (p *SOL) ChainID() string {
	return "solana"
}

// Description returns the protocol description.
func (p *SOL) Description() string {
	return p.description
}

// Functions returns the available functions.
func (p *SOL) Functions() []*types.Function {
	return p.functions
}

// GetFunction retrieves a function by ID.
func (p *SOL) GetFunction(id string) (*types.Function, error) {
	for _, fn := range p.functions {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found in protocol %q", id, p.id)
}

// MatchFunctionCall matches a transaction against a policy function matcher.
func (p *SOL) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	// Solana function matching is handled by the engine via IDL
	return false, nil, fmt.Errorf("solana function matching is handled by the engine")
}

