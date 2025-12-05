package bitcoincash

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/vultisig/recipes/types"
)

// BCH implements the Protocol interface for the Bitcoin Cash protocol
type BCH struct{}

// ID returns the unique identifier for the BCH protocol
func (b *BCH) ID() string {
	return "bch"
}

// Name returns a human-readable name for the BCH protocol
func (b *BCH) Name() string {
	return "Bitcoin Cash"
}

// ChainID returns the ID of the chain this protocol belongs to
func (b *BCH) ChainID() string {
	return "bitcoincash"
}

// Description returns a detailed description of the BCH protocol
func (b *BCH) Description() string {
	return "The native cryptocurrency of the Bitcoin Cash blockchain, used for transactions and value transfer."
}

// Functions returns a list of available functions for this protocol
func (b *BCH) Functions() []*types.Function {
	return []*types.Function{
		{
			ID:          "transfer",
			Name:        "Transfer BCH",
			Description: "Transfer Bitcoin Cash to another address",
			Parameters: []*types.FunctionParam{
				{
					Name:        "recipient",
					Type:        "address",
					Description: "The Bitcoin Cash address of the recipient",
				},
				{
					Name:        "amount",
					Type:        "decimal",
					Description: "The amount of Bitcoin Cash to transfer",
				},
			},
		},
	}
}

// GetFunction retrieves a specific function by ID
func (b *BCH) GetFunction(id string) (*types.Function, error) {
	for _, fn := range b.Functions() {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found for protocol BCH", id)
}

func (b *BCH) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	// Check if this is a Bitcoin Cash transaction
	if decodedTx.ChainIdentifier() != "bitcoincash" {
		return false, nil, fmt.Errorf("expected Bitcoin Cash transaction, got %s", decodedTx.ChainIdentifier())
	}

	// Only support transfer function
	if policyMatcher.FunctionID != "transfer" {
		return false, nil, nil
	}

	// Extract parameters from the transaction
	params := make(map[string]interface{})
	params["recipient"] = decodedTx.To()
	params["amount"] = decodedTx.Value() // Amount as *big.Int in satoshis

	// Also store a string representation for display
	displayParams := make(map[string]interface{})
	displayParams["recipient"] = decodedTx.To()
	displayParams["amount"] = decodedTx.Value().String() // Amount in satoshis as string

	// Check constraints
	for _, pc := range policyMatcher.Constraints {
		if pc == nil {
			continue
		}

		paramName := pc.GetParameterName()
		constraint := pc.GetConstraint()

		if constraint == nil {
			return false, nil, fmt.Errorf("nil constraint found for parameter %q", paramName)
		}

		paramValue, exists := params[paramName]
		if !exists {
			if constraint.GetRequired() {
				return false, nil, fmt.Errorf("required parameter %s not found", paramName)
			}
			continue
		}

		// Validate based on constraint type
		switch constraint.GetType() {
		case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
			valStr := ""
			switch v := paramValue.(type) {
			case string:
				valStr = v
			case *big.Int:
				valStr = v.String()
			default:
				return false, nil, fmt.Errorf("parameter %q has unsupported type %T", paramName, paramValue)
			}

			if !strings.EqualFold(valStr, constraint.GetFixedValue()) {
				return false, nil, nil // Constraint not met
			}

		case types.ConstraintType_CONSTRAINT_TYPE_MAX:
			var amount *big.Int
			switch v := paramValue.(type) {
			case *big.Int:
				amount = v
			case string:
				var ok bool
				amount, ok = new(big.Int).SetString(v, 10)
				if !ok {
					return false, nil, fmt.Errorf("parameter %q value %q is not a valid number", paramName, v)
				}
			default:
				return false, nil, fmt.Errorf("parameter %q has unsupported type %T for MAX constraint", paramName, paramValue)
			}

			maxValue, ok := new(big.Int).SetString(constraint.GetMaxValue(), 10)
			if !ok {
				return false, nil, fmt.Errorf("constraint max_value %q is not a valid number", constraint.GetMaxValue())
			}

			if amount.Cmp(maxValue) > 0 {
				return false, nil, nil // Amount exceeds maximum
			}

		default:
			return false, nil, fmt.Errorf("unsupported constraint type %v for parameter %q", constraint.GetType(), paramName)
		}
	}

	return true, displayParams, nil
}

// NewBCH creates a new BCH protocol instance
func NewBCH() types.Protocol {
	return &BCH{}
}

