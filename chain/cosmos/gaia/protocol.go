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

// StakingRedelegate implements the Cosmos Hub redelegate protocol.
type StakingRedelegate struct {
	id          string
	name        string
	description string
	functions   []*types.Function
}

// NewStakingRedelegate creates a new StakingRedelegate protocol instance.
func NewStakingRedelegate() *StakingRedelegate {
	return &StakingRedelegate{
		id:          "staking_redelegate",
		name:        "Cosmos Staking Redelegate",
		description: "Redelegate staked ATOM from one validator to another on the Cosmos Hub",
		functions: []*types.Function{
			{
				ID:          "redelegate",
				Name:        "Redelegate ATOM",
				Description: "Move staked ATOM from a source validator to a destination validator",
				Parameters: []*types.FunctionParam{
					// delegator_address carries the account (cosmos1...) HRP.
					{Name: "delegator_address", Type: "address", Description: "The Cosmos account address of the delegator (cosmos1... prefix)"},
					// validator_src_address and validator_dst_address MUST carry the validator
					// operator (cosmosvaloper1...) HRP — not the delegator HRP. The engine
					// enforces this at extraction time; rules that constrain these fields will
					// reject any cosmos1... address.
					{Name: "validator_src_address", Type: "address", Description: "The source validator operator address (cosmosvaloper1... prefix)"},
					{Name: "validator_dst_address", Type: "address", Description: "The destination validator operator address (cosmosvaloper1... prefix)"},
					{Name: "amount", Type: "decimal", Description: "The amount of ATOM to redelegate (in uatom)"},
					{Name: "denom", Type: "string", Description: "The coin denomination (e.g. uatom)"},
				},
			},
		},
	}
}

// ID returns the protocol identifier.
func (p *StakingRedelegate) ID() string { return p.id }

// Name returns the protocol name.
func (p *StakingRedelegate) Name() string { return p.name }

// ChainID returns the chain identifier.
func (p *StakingRedelegate) ChainID() string { return "cosmos" }

// Description returns the protocol description.
func (p *StakingRedelegate) Description() string { return p.description }

// Functions returns the available functions.
func (p *StakingRedelegate) Functions() []*types.Function { return p.functions }

// GetFunction retrieves a function by ID.
func (p *StakingRedelegate) GetFunction(id string) (*types.Function, error) {
	for _, fn := range p.functions {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found in protocol %q", id, p.id)
}

// MatchFunctionCall matches a transaction against a policy function matcher.
func (p *StakingRedelegate) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	return false, nil, fmt.Errorf("cosmos function matching is handled by the engine")
}

// StakingWithdrawRewards implements the Cosmos Hub withdraw delegator rewards protocol.
type StakingWithdrawRewards struct {
	id          string
	name        string
	description string
	functions   []*types.Function
}

// NewStakingWithdrawRewards creates a new StakingWithdrawRewards protocol instance.
func NewStakingWithdrawRewards() *StakingWithdrawRewards {
	return &StakingWithdrawRewards{
		id:          "staking_withdraw_rewards",
		name:        "Cosmos Staking Withdraw Rewards",
		description: "Withdraw staking rewards accrued from a validator on the Cosmos Hub",
		functions: []*types.Function{
			{
				ID:          "withdraw_rewards",
				Name:        "Withdraw Staking Rewards",
				Description: "Claim accumulated ATOM staking rewards from a validator",
				Parameters: []*types.FunctionParam{
					// delegator_address carries the account (cosmos1...) HRP.
					{Name: "delegator_address", Type: "address", Description: "The Cosmos account address of the delegator (cosmos1... prefix)"},
					// validator_address MUST carry the validator operator (cosmosvaloper1...)
					// HRP — not the delegator HRP. The engine enforces this at extraction
					// time; rules constraining this field will reject any cosmos1... address.
					{Name: "validator_address", Type: "address", Description: "The validator operator address to claim rewards from (cosmosvaloper1... prefix)"},
				},
			},
		},
	}
}

// ID returns the protocol identifier.
func (p *StakingWithdrawRewards) ID() string { return p.id }

// Name returns the protocol name.
func (p *StakingWithdrawRewards) Name() string { return p.name }

// ChainID returns the chain identifier.
func (p *StakingWithdrawRewards) ChainID() string { return "cosmos" }

// Description returns the protocol description.
func (p *StakingWithdrawRewards) Description() string { return p.description }

// Functions returns the available functions.
func (p *StakingWithdrawRewards) Functions() []*types.Function { return p.functions }

// GetFunction retrieves a function by ID.
func (p *StakingWithdrawRewards) GetFunction(id string) (*types.Function, error) {
	for _, fn := range p.functions {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found in protocol %q", id, p.id)
}

// MatchFunctionCall matches a transaction against a policy function matcher.
func (p *StakingWithdrawRewards) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	return false, nil, fmt.Errorf("cosmos function matching is handled by the engine")
}
