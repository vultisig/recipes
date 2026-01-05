package gaia

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// ATOM implements the native Cosmos ATOM protocol.
type ATOM struct {
	id          string
	name        string
	description string
	functions   []*types.Function
}

// NewATOM creates a new ATOM protocol instance.
func NewATOM() *ATOM {
	return &ATOM{
		id:          "atom",
		name:        "ATOM",
		description: "Native ATOM currency of the Cosmos Hub blockchain",
		functions: []*types.Function{
			{
				ID:          "transfer",
				Name:        "Transfer ATOM",
				Description: "Transfer ATOM to another address",
				Parameters: []*types.FunctionParam{
					{Name: "recipient", Type: "address", Description: "The Cosmos address of the recipient"},
					{Name: "amount", Type: "decimal", Description: "The amount of ATOM to transfer (in uatom)"},
				},
			},
		},
	}
}

// ID returns the protocol identifier.
func (p *ATOM) ID() string {
	return p.id
}

// Name returns the protocol name.
func (p *ATOM) Name() string {
	return p.name
}

// ChainID returns the chain identifier.
func (p *ATOM) ChainID() string {
	return "cosmos"
}

// Description returns the protocol description.
func (p *ATOM) Description() string {
	return p.description
}

// Functions returns the available functions.
func (p *ATOM) Functions() []*types.Function {
	return p.functions
}

// GetFunction retrieves a function by ID.
func (p *ATOM) GetFunction(id string) (*types.Function, error) {
	for _, fn := range p.functions {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found in protocol %q", id, p.id)
}

// MatchFunctionCall matches a transaction against a policy function matcher.
func (p *ATOM) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	return false, nil, fmt.Errorf("cosmos function matching is handled by the engine")
}

