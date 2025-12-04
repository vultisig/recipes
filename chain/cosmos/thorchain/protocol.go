package thorchain

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// RUNE implements the native THORChain RUNE protocol.
type RUNE struct {
	id          string
	name        string
	description string
	functions   []*types.Function
}

// NewRUNE creates a new RUNE protocol instance.
func NewRUNE() *RUNE {
	return &RUNE{
		id:          "rune",
		name:        "RUNE",
		description: "Native RUNE currency of the THORChain blockchain",
		functions: []*types.Function{
			{
				ID:          "transfer",
				Name:        "Transfer RUNE",
				Description: "Transfer RUNE to another address",
				Parameters: []*types.FunctionParam{
					{Name: "recipient", Type: "address", Description: "The THORChain address of the recipient"},
					{Name: "amount", Type: "decimal", Description: "The amount of RUNE to transfer"},
				},
			},
		},
	}
}

// ID returns the protocol identifier.
func (p *RUNE) ID() string {
	return p.id
}

// Name returns the protocol name.
func (p *RUNE) Name() string {
	return p.name
}

// ChainID returns the chain identifier.
func (p *RUNE) ChainID() string {
	return "thorchain"
}

// Description returns the protocol description.
func (p *RUNE) Description() string {
	return p.description
}

// Functions returns the available functions.
func (p *RUNE) Functions() []*types.Function {
	return p.functions
}

// GetFunction retrieves a function by ID.
func (p *RUNE) GetFunction(id string) (*types.Function, error) {
	for _, fn := range p.functions {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found in protocol %q", id, p.id)
}

// MatchFunctionCall matches a transaction against a policy function matcher.
func (p *RUNE) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	return false, nil, fmt.Errorf("THORChain function matching is handled by the engine")
}

// Swap implements the THORChain swap protocol.
type Swap struct {
	id          string
	name        string
	description string
	functions   []*types.Function
}

// NewSwap creates a new Swap protocol instance.
func NewSwap() *Swap {
	return &Swap{
		id:          "thorchain_swap",
		name:        "THORChain Swap",
		description: "Cross-chain swap protocol on THORChain",
		functions: []*types.Function{
			{
				ID:          "swap",
				Name:        "Swap",
				Description: "Perform a cross-chain swap via THORChain",
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
	return "thorchain"
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
	return false, nil, fmt.Errorf("THORChain function matching is handled by the engine")
}

