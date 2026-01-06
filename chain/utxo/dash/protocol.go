package dash

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/vultisig/recipes/types"
)

// DASH implements the Protocol interface for the Dash protocol
type DASH struct{}

// ID returns the unique identifier for the DASH protocol
func (d *DASH) ID() string {
	return "dash"
}

// Name returns a human-readable name for the DASH protocol
func (d *DASH) Name() string {
	return "Dash"
}

// ChainID returns the ID of the chain this protocol belongs to
func (d *DASH) ChainID() string {
	return "dash"
}

// Description returns a detailed description of the DASH protocol
func (d *DASH) Description() string {
	return "The native cryptocurrency of the Dash blockchain, used for transactions and value transfer with InstantSend capabilities."
}

// Functions returns a list of available functions for this protocol
func (d *DASH) Functions() []*types.Function {
	return []*types.Function{
		{
			ID:          "transfer",
			Name:        "Transfer DASH",
			Description: "Transfer Dash to another address",
			Parameters: []*types.FunctionParam{
				{
					Name:        "recipient",
					Type:        "address",
					Description: "The Dash address of the recipient",
				},
				{
					Name:        "amount",
					Type:        "decimal",
					Description: "The amount of Dash to transfer",
				},
			},
		},
	}
}

// GetFunction retrieves a specific function by ID
func (d *DASH) GetFunction(id string) (*types.Function, error) {
	for _, fn := range d.Functions() {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found for protocol DASH", id)
}

func (d *DASH) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	// Check if this is a Dash transaction
	if decodedTx.ChainIdentifier() != "dash" {
		return false, nil, fmt.Errorf("expected Dash transaction, got %s", decodedTx.ChainIdentifier())
	}

	// Only support transfer function
	if policyMatcher.FunctionID != "transfer" {
		return false, nil, nil
	}

	// Extract parameters from the transaction
	params := make(map[string]interface{})
	params["recipient"] = decodedTx.To()
	params["amount"] = decodedTx.Value() // Amount as *big.Int in duffs

	// Also store a string representation for display
	displayParams := make(map[string]interface{})
	displayParams["recipient"] = decodedTx.To()
	displayParams["amount"] = decodedTx.Value().String() // Amount in duffs as string

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

// NewDASH creates a new DASH protocol instance
func NewDASH() types.Protocol {
	return &DASH{}
}


