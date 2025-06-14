package thorchain

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/vultisig/recipes/types"
)

// Thorchain constants for basic transfers
const (
	// RUNE specifications
	RuneDecimals  = 8
	RuneBaseDenom = "rune"

	// TCY specifications
	TCYDecimals  = 8
	TCYBaseDenom = "tcy"
)

// BaseProtocol provides common functionality for Thorchain protocols
type BaseProtocol struct {
	id          string
	name        string
	description string
	functions   []*types.Function
}

// ID returns the protocol identifier
func (p *BaseProtocol) ID() string {
	return p.id
}

// Name returns the protocol name
func (p *BaseProtocol) Name() string {
	return p.name
}

// ChainID returns the chain identifier
func (p *BaseProtocol) ChainID() string {
	return "thorchain"
}

// Description returns the protocol description
func (p *BaseProtocol) Description() string {
	return p.description
}

// Functions returns the available functions
func (p *BaseProtocol) Functions() []*types.Function {
	return p.functions
}

// GetFunction retrieves a function by ID
func (p *BaseProtocol) GetFunction(id string) (*types.Function, error) {
	for _, fn := range p.functions {
		if fn.ID == id {
			return fn, nil
		}
	}
	return nil, fmt.Errorf("function %q not found in protocol %q", id, p.id)
}

// RUNE implements the native RUNE protocol
type RUNE struct {
	BaseProtocol
}

// NewRUNE creates a new RUNE protocol instance
func NewRUNE() types.Protocol {
	return &RUNE{
		BaseProtocol: BaseProtocol{
			id:          "rune",
			name:        "RUNE",
			description: "Native RUNE token of the Thorchain blockchain",
			functions: []*types.Function{
				{
					ID:          "transfer",
					Name:        "Transfer RUNE",
					Description: "Transfer RUNE tokens to another address",
					Parameters: []*types.FunctionParam{
						{
							Name:        "recipient",
							Type:        "address",
							Description: "The Thorchain recipient address (thor... format)",
						},
						{
							Name:        "amount",
							Type:        "decimal",
							Description: "The amount of RUNE to transfer (8 decimals)",
						},
						{
							Name:        "memo",
							Type:        "string",
							Description: "Optional memo field (max 80 characters)",
						},
					},
				},
			},
		},
	}
}

// TCY implements the TCY token protocol
type TCY struct {
	BaseProtocol
}

// NewTCY creates a new TCY protocol instance
func NewTCY() types.Protocol {
	return &TCY{
		BaseProtocol: BaseProtocol{
			id:          "tcy",
			name:        "TCY",
			description: "TCY token on the Thorchain blockchain",
			functions: []*types.Function{
				{
					ID:          "transfer",
					Name:        "Transfer TCY",
					Description: "Transfer TCY tokens to another address",
					Parameters: []*types.FunctionParam{
						{
							Name:        "recipient",
							Type:        "address",
							Description: "The Thorchain recipient address (thor... format)",
						},
						{
							Name:        "amount",
							Type:        "decimal",
							Description: "The amount of TCY to transfer (8 decimals)",
						},
						{
							Name:        "memo",
							Type:        "string",
							Description: "Optional memo field",
						},
					},
				},
			},
		},
	}
}

// MatchFunctionCall for RUNE protocol
func (r *RUNE) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	if decodedTx.ChainIdentifier() != "thorchain" {
		return false, nil, fmt.Errorf("expected Thorchain transaction, got %s", decodedTx.ChainIdentifier())
	}

	if policyMatcher.FunctionID != "transfer" {
		return false, nil, nil // This protocol only supports transfer operations
	}

	// Add sanity checks to ensure transaction has valid data
	recipient := decodedTx.To()
	if recipient == "" {
		return false, nil, nil // No recipient, not a valid transfer
	}

	amount := decodedTx.Value()
	if amount == nil || amount.Cmp(big.NewInt(0)) <= 0 {
		return false, nil, nil // No amount or zero amount, not a valid transfer
	}

	params := map[string]interface{}{
		"recipient": strings.ToLower(recipient),
		"amount":    amount,
	}

	// Add memo if present
	if len(decodedTx.Data()) > 0 {
		params["memo"] = string(decodedTx.Data())
	}

	// Basic constraint evaluation
	constraintsMet, err := r.evaluateBasicConstraints(params, policyMatcher.Constraints)
	if err != nil {
		return false, nil, fmt.Errorf("error evaluating constraints: %w", err)
	}

	if !constraintsMet {
		return false, nil, nil
	}

	return true, params, nil
}

// MatchFunctionCall for TCY protocol
func (t *TCY) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	if decodedTx.ChainIdentifier() != "thorchain" {
		return false, nil, fmt.Errorf("expected Thorchain transaction, got %s", decodedTx.ChainIdentifier())
	}

	if policyMatcher.FunctionID != "transfer" {
		return false, nil, nil // This protocol only supports transfer operations
	}

	// Add sanity checks to ensure transaction has valid data
	recipient := decodedTx.To()
	if recipient == "" {
		return false, nil, nil // No recipient, not a valid transfer
	}

	amount := decodedTx.Value()
	if amount == nil || amount.Cmp(big.NewInt(0)) <= 0 {
		return false, nil, nil // No amount or zero amount, not a valid transfer
	}

	params := map[string]interface{}{
		"recipient": strings.ToLower(recipient),
		"amount":    amount,
	}

	// Add memo if present
	if len(decodedTx.Data()) > 0 {
		params["memo"] = string(decodedTx.Data())
	}

	// Basic constraint evaluation
	constraintsMet, err := t.evaluateBasicConstraints(params, policyMatcher.Constraints)
	if err != nil {
		return false, nil, fmt.Errorf("error evaluating constraints: %w", err)
	}

	if !constraintsMet {
		return false, nil, nil
	}

	return true, params, nil
}

// evaluateBasicConstraints provides simplified constraint evaluation
func (r *RUNE) evaluateBasicConstraints(params map[string]interface{}, constraints []*types.ParameterConstraint) (bool, error) {
	return evaluateBasicConstraintsHelper(params, constraints)
}

// evaluateBasicConstraints for TCY - uses shared helper
func (t *TCY) evaluateBasicConstraints(params map[string]interface{}, constraints []*types.ParameterConstraint) (bool, error) {
	return evaluateBasicConstraintsHelper(params, constraints)
}

// evaluateBasicConstraintsHelper provides shared constraint evaluation logic for both RUNE and TCY
func evaluateBasicConstraintsHelper(params map[string]interface{}, constraints []*types.ParameterConstraint) (bool, error) {
	for _, pc := range constraints {
		if pc == nil {
			continue
		}

		paramName := pc.GetParameterName()
		constraint := pc.GetConstraint()

		if constraint == nil {
			return false, fmt.Errorf("nil constraint found for parameter %q", paramName)
		}

		paramValue, exists := params[paramName]
		if !exists {
			if constraint.GetRequired() {
				return false, fmt.Errorf("required parameter %s not found", paramName)
			}
			continue
		}

		// Handle basic constraint types (focus on most common ones)
		switch constraint.GetType() {
		case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
			valStr := fmt.Sprintf("%v", paramValue)
			if !strings.EqualFold(valStr, constraint.GetFixedValue()) {
				return false, nil
			}

		case types.ConstraintType_CONSTRAINT_TYPE_MAX:
			if amount, ok := paramValue.(*big.Int); ok {
				maxValue, ok := new(big.Int).SetString(constraint.GetMaxValue(), 10)
				if !ok {
					return false, fmt.Errorf("invalid max_value: %s", constraint.GetMaxValue())
				}
				if amount.Cmp(maxValue) > 0 {
					return false, nil
				}
			}

		default:
			// Return false for unsupported constraint types to prevent bypassing restrictions
			return false, fmt.Errorf("unsupported constraint type: %v for parameter %s", constraint.GetType(), paramName)
		}
	}

	return true, nil
}
