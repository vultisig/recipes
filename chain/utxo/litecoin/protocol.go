package litecoin

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/vultisig/recipes/types"
)

// LTC implements the Protocol interface for the Litecoin protocol
type LTC struct{}

// ID returns the unique identifier for the LTC protocol
func (l *LTC) ID() string {
	return "ltc"
}

// Name returns a human-readable name for the LTC protocol
func (l *LTC) Name() string {
	return "Litecoin"
}

// ChainID returns the ID of the chain this protocol belongs to
func (l *LTC) ChainID() string {
	return "litecoin"
}

// Description returns a detailed description of the LTC protocol
func (l *LTC) Description() string {
	return "The native cryptocurrency of the Litecoin blockchain, used for transactions and value transfer."
}

// Functions returns a list of available functions for this protocol
func (l *LTC) Functions() []*types.Function {
	return []*types.Function{
		{
			ID:          "transfer",
			Name:        "Transfer LTC",
			Description: "Transfer Litecoin to another address",
			Parameters: []*types.FunctionParam{
				{
					Name:        "recipient",
					Type:        "address",
					Description: "The Litecoin address of the recipient",
				},
				{
					Name:        "amount",
					Type:        "decimal",
					Description: "The amount of Litecoin to transfer",
				},
			},
		},
	}
}

// GetFunction retrieves a specific function by ID
func (l *LTC) GetFunction(id string) (*types.Function, error) {
	for _, fn := range l.Functions() {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found for protocol LTC", id)
}

func (l *LTC) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	// Check if this is a Litecoin transaction
	if decodedTx.ChainIdentifier() != "litecoin" {
		return false, nil, fmt.Errorf("expected Litecoin transaction, got %s", decodedTx.ChainIdentifier())
	}

	// Only support transfer function
	if policyMatcher.FunctionID != "transfer" {
		return false, nil, nil
	}

	// Extract parameters from the transaction
	params := make(map[string]interface{})
	params["recipient"] = decodedTx.To()
	params["amount"] = decodedTx.Value() // Amount as *big.Int in litoshi

	// Also store a string representation for display
	displayParams := make(map[string]interface{})
	displayParams["recipient"] = decodedTx.To()
	displayParams["amount"] = decodedTx.Value().String() // Amount in litoshi as string

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

// NewLTC creates a new LTC protocol instance
func NewLTC() types.Protocol {
	return &LTC{}
}

