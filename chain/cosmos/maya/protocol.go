package maya

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// CACAO implements the native MAYAChain CACAO protocol.
type CACAO struct {
	id          string
	name        string
	description string
	functions   []*types.Function
}

// NewCACAO creates a new CACAO protocol instance.
func NewCACAO() *CACAO {
	return &CACAO{
		id:          "cacao",
		name:        "CACAO",
		description: "Native CACAO currency of the MAYAChain blockchain",
		functions: []*types.Function{
			{
				ID:          "transfer",
				Name:        "Transfer CACAO",
				Description: "Transfer CACAO to another address",
				Parameters: []*types.FunctionParam{
					{Name: "recipient", Type: "address", Description: "The MAYAChain address of the recipient"},
					{Name: "amount", Type: "decimal", Description: "The amount of CACAO to transfer"},
				},
			},
		},
	}
}

// ID returns the protocol identifier.
func (p *CACAO) ID() string {
	return p.id
}

// Name returns the protocol name.
func (p *CACAO) Name() string {
	return p.name
}

// ChainID returns the chain identifier.
func (p *CACAO) ChainID() string {
	return "mayachain"
}

// Description returns the protocol description.
func (p *CACAO) Description() string {
	return p.description
}

// Functions returns the available functions.
func (p *CACAO) Functions() []*types.Function {
	return p.functions
}

// GetFunction retrieves a function by ID.
func (p *CACAO) GetFunction(id string) (*types.Function, error) {
	for _, fn := range p.functions {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found in protocol %q", id, p.id)
}

// MatchFunctionCall matches a transaction against a policy function matcher.
func (p *CACAO) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	return false, nil, fmt.Errorf("mayachain function matching is handled by the engine")
}

// Swap implements the MAYAChain swap protocol.
type Swap struct {
	id          string
	name        string
	description string
	functions   []*types.Function
}

// NewSwap creates a new Swap protocol instance.
func NewSwap() *Swap {
	return &Swap{
		id:          "mayachain_swap",
		name:        "MAYAChain Swap",
		description: "Cross-chain swap protocol on MAYAChain",
		functions: []*types.Function{
			{
				ID:          "swap",
				Name:        "Swap",
				Description: "Perform a cross-chain swap via MAYAChain",
				Parameters: []*types.FunctionParam{
					{Name: "from_asset", Type: "string", Description: "The source asset"},
					{Name: "amount", Type: "decimal", Description: "The amount to swap"},
					{Name: "memo", Type: "string", Description: "The swap memo containing destination details"},
				},
			},
		},
	}
}

// ID returns the protocol identifier.
func (p *Swap) ID() string {
	return p.id
}

// Name returns the protocol name.
func (p *Swap) Name() string {
	return p.name
}

// ChainID returns the chain identifier.
func (p *Swap) ChainID() string {
	return "mayachain"
}

// Description returns the protocol description.
func (p *Swap) Description() string {
	return p.description
}

// Functions returns the available functions.
func (p *Swap) Functions() []*types.Function {
	return p.functions
}

// GetFunction retrieves a function by ID.
func (p *Swap) GetFunction(id string) (*types.Function, error) {
	for _, fn := range p.functions {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found in protocol %q", id, p.id)
}

// MatchFunctionCall matches a transaction against a policy function matcher.
func (p *Swap) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	return false, nil, fmt.Errorf("mayachain function matching is handled by the engine")
}

