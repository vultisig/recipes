package tron

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// TRX implements the native TRON TRX protocol.
type TRX struct {
	id          string
	name        string
	description string
	functions   []*types.Function
}

// NewTRX creates a new TRX protocol instance.
func NewTRX() *TRX {
	return &TRX{
		id:          "trx",
		name:        "TRX",
		description: "Native TRX currency of the TRON blockchain",
		functions: []*types.Function{
			{
				ID:          "transfer",
				Name:        "Transfer TRX",
				Description: "Transfer TRX to another address",
				Parameters: []*types.FunctionParam{
					{Name: "recipient", Type: "address", Description: "The TRON address of the recipient"},
					{Name: "amount", Type: "decimal", Description: "The amount of TRX to transfer (in SUN, 1 TRX = 1,000,000 SUN)"},
				},
			},
		},
	}
}

// TRC20 implements the TRON TRC20 token protocol.
type TRC20 struct {
	id          string
	name        string
	description string
	functions   []*types.Function
}

// NewTRC20 creates a new TRC20 protocol instance.
func NewTRC20() *TRC20 {
	return &TRC20{
		id:          "trc20",
		name:        "TRC20",
		description: "TRC20 token transfers on the TRON blockchain",
		functions: []*types.Function{
			{
				ID:          "transfer",
				Name:        "Transfer TRC20",
				Description: "Transfer TRC20 tokens to another address",
				Parameters: []*types.FunctionParam{
					{Name: "recipient", Type: "address", Description: "The TRON address of the recipient"},
					{Name: "amount", Type: "decimal", Description: "The amount of tokens to transfer (token smallest unit)"},
					{Name: "from_asset", Type: "address", Description: "TRC20 contract address"},
					{Name: "memo", Type: "string", Description: "Optional memo for swap routing"},
				},
			},
		},
	}
}

// ID returns the protocol identifier.
func (p *TRC20) ID() string {
	return p.id
}

// Name returns the protocol name.
func (p *TRC20) Name() string {
	return p.name
}

// ChainID returns the chain identifier.
func (p *TRC20) ChainID() string {
	return "tron"
}

// Description returns the protocol description.
func (p *TRC20) Description() string {
	return p.description
}

// Functions returns the available functions.
func (p *TRC20) Functions() []*types.Function {
	return p.functions
}

// GetFunction retrieves a function by ID.
func (p *TRC20) GetFunction(id string) (*types.Function, error) {
	for _, fn := range p.functions {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found in protocol %q", id, p.id)
}

// MatchFunctionCall matches a transaction against a policy function matcher.
func (p *TRC20) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	return false, nil, fmt.Errorf("tron function matching is handled by the engine")
}

// ID returns the protocol identifier.
func (p *TRX) ID() string {
	return p.id
}

// Name returns the protocol name.
func (p *TRX) Name() string {
	return p.name
}

// ChainID returns the chain identifier.
func (p *TRX) ChainID() string {
	return "tron"
}

// Description returns the protocol description.
func (p *TRX) Description() string {
	return p.description
}

// Functions returns the available functions.
func (p *TRX) Functions() []*types.Function {
	return p.functions
}

// GetFunction retrieves a function by ID.
func (p *TRX) GetFunction(id string) (*types.Function, error) {
	for _, fn := range p.functions {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found in protocol %q", id, p.id)
}

// MatchFunctionCall matches a transaction against a policy function matcher.
func (p *TRX) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	return false, nil, fmt.Errorf("tron function matching is handled by the engine")
}
