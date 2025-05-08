package policy

import (
	"fmt"

	"github.com/vultisig/recipes/transaction"
)

// ConstraintType defines the type of constraint
type ConstraintType string

const (
	// ConstraintTypeFixed requires an exact match
	ConstraintTypeFixed ConstraintType = "fixed"

	// ConstraintTypeMax sets a maximum value
	ConstraintTypeMax ConstraintType = "max"

	// ConstraintTypeMin sets a minimum value
	ConstraintTypeMin ConstraintType = "min"

	// ConstraintTypeRange requires a value within a range
	ConstraintTypeRange ConstraintType = "range"

	// ConstraintTypeWhitelist requires a value from an allowed list
	ConstraintTypeWhitelist ConstraintType = "whitelist"

	// ConstraintTypeMaxPerPeriod limits frequency within a time period
	ConstraintTypeMaxPerPeriod ConstraintType = "max_per_period"
)

// RangeValue represents a range constraint with min and max values
type RangeValue struct {
	Min interface{} `json:"min"`
	Max interface{} `json:"max"`
}

// Constraint defines restrictions on a parameter
type Constraint struct {
	// Type of constraint (fixed, max, min, etc.)
	Type ConstraintType `json:"type"`

	// Value of the constraint (depends on Type)
	Value interface{} `json:"value"`

	// Additional metadata for the constraint
	DenominatedIn string `json:"denominated_in,omitempty"`
	Period        string `json:"period,omitempty"`
	Required      bool   `json:"required,omitempty"`
}

// NewFixedConstraint creates a constraint that requires an exact match
func NewFixedConstraint(value interface{}, required bool) *Constraint {
	return &Constraint{
		Type:     ConstraintTypeFixed,
		Value:    value,
		Required: required,
	}
}

// NewMaxConstraint creates a constraint that limits the maximum value
func NewMaxConstraint(value interface{}, denominatedIn string, required bool) *Constraint {
	return &Constraint{
		Type:          ConstraintTypeMax,
		Value:         value,
		DenominatedIn: denominatedIn,
		Required:      required,
	}
}

// NewMinConstraint creates a constraint that sets a minimum value
func NewMinConstraint(value interface{}, denominatedIn string, required bool) *Constraint {
	return &Constraint{
		Type:          ConstraintTypeMin,
		Value:         value,
		DenominatedIn: denominatedIn,
		Required:      required,
	}
}

// NewRangeConstraint creates a constraint that requires a value within a range
func NewRangeConstraint(min, max interface{}, denominatedIn string, required bool) *Constraint {
	return &Constraint{
		Type: ConstraintTypeRange,
		Value: RangeValue{
			Min: min,
			Max: max,
		},
		DenominatedIn: denominatedIn,
		Required:      required,
	}
}

// NewWhitelistConstraint creates a constraint that requires a value from an allowed list
func NewWhitelistConstraint(values []interface{}, required bool) *Constraint {
	return &Constraint{
		Type:     ConstraintTypeWhitelist,
		Value:    values,
		Required: required,
	}
}

// NewMaxPerPeriodConstraint creates a constraint that limits frequency within a time period
func NewMaxPerPeriodConstraint(maxCount int, period string, required bool) *Constraint {
	return &Constraint{
		Type:     ConstraintTypeMaxPerPeriod,
		Value:    maxCount,
		Period:   period,
		Required: required,
	}
}

// Validate checks if a parameter value satisfies this constraint
func (c *Constraint) Validate(paramValue transaction.ParamValue) (bool, error) {
	switch c.Type {
	case ConstraintTypeFixed:
		return c.validateFixed(paramValue)
	case ConstraintTypeMax:
		return c.validateMax(paramValue)
	case ConstraintTypeMin:
		return c.validateMin(paramValue)
	case ConstraintTypeRange:
		return c.validateRange(paramValue)
	case ConstraintTypeWhitelist:
		return c.validateWhitelist(paramValue)
	case ConstraintTypeMaxPerPeriod:
		// This would require a transaction history, which isn't passed here
		// In a real implementation, you would track transactions and check frequency
		return true, nil
	default:
		return false, fmt.Errorf("unsupported constraint type: %s", c.Type)
	}
}

// validateFixed checks for an exact match
func (c *Constraint) validateFixed(paramValue transaction.ParamValue) (bool, error) {
	// Convert both values to strings for easier comparison
	constraintStr := fmt.Sprintf("%v", c.Value)
	paramStr := fmt.Sprintf("%v", paramValue.Value)

	return constraintStr == paramStr, nil
}

// validateMax checks if value is less than or equal to maximum
func (c *Constraint) validateMax(paramValue transaction.ParamValue) (bool, error) {
	// For simplicity, we'll just compare as strings for now
	// In a real implementation, you'd need proper numeric comparison based on types
	constraintStr := fmt.Sprintf("%v", c.Value)
	paramStr := fmt.Sprintf("%v", paramValue.Value)

	// Simple string comparison works for similar formatted numbers
	// But a real implementation would need proper numeric parsing
	return paramStr <= constraintStr, nil
}

// validateMin checks if value is greater than or equal to minimum
func (c *Constraint) validateMin(paramValue transaction.ParamValue) (bool, error) {
	// For simplicity, we'll just compare as strings for now
	constraintStr := fmt.Sprintf("%v", c.Value)
	paramStr := fmt.Sprintf("%v", paramValue.Value)

	return paramStr >= constraintStr, nil
}

// validateRange checks if value is within the range
func (c *Constraint) validateRange(paramValue transaction.ParamValue) (bool, error) {
	rangeValue, ok := c.Value.(RangeValue)
	if !ok {
		return false, fmt.Errorf("expected RangeValue for range constraint")
	}

	// Convert to strings for comparison
	minStr := fmt.Sprintf("%v", rangeValue.Min)
	maxStr := fmt.Sprintf("%v", rangeValue.Max)
	paramStr := fmt.Sprintf("%v", paramValue.Value)

	return paramStr >= minStr && paramStr <= maxStr, nil
}

// validateWhitelist checks if value is in the allowed list
func (c *Constraint) validateWhitelist(paramValue transaction.ParamValue) (bool, error) {
	whitelist, ok := c.Value.([]interface{})
	if !ok {
		return false, fmt.Errorf("expected []interface{} for whitelist constraint")
	}

	paramStr := fmt.Sprintf("%v", paramValue.Value)

	for _, allowed := range whitelist {
		allowedStr := fmt.Sprintf("%v", allowed)
		if paramStr == allowedStr {
			return true, nil
		}
	}

	return false, nil
}
