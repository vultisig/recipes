package ethereum

import (
	"github.com/vultisig/recipes/types"
)

// ProtocolValidator defines the interface for protocol-specific validation logic
type ProtocolValidator interface {
	// ValidateTransaction validates a transaction against protocol-specific business rules
	ValidateTransaction(functionName string, params map[string]interface{}) error

	// CustomizeFunctions allows the validator to customize function descriptions and parameters
	CustomizeFunctions(f *types.Function, abiFunc *ABIFunction)

	// GetProtocolID returns the protocol ID this validator is for
	GetProtocolID() string
}

// ValidatorRegistry manages custom validators for different protocols
type ValidatorRegistry struct {
	validators map[string]ProtocolValidator
}

// NewValidatorRegistry creates a new validator registry
func NewValidatorRegistry() *ValidatorRegistry {
	return &ValidatorRegistry{
		validators: make(map[string]ProtocolValidator),
	}
}

// RegisterValidator registers a custom validator for a protocol
func (vr *ValidatorRegistry) RegisterValidator(protocolID string, validator ProtocolValidator) {
	vr.validators[protocolID] = validator
}

// GetValidator returns the validator for a protocol, if any
func (vr *ValidatorRegistry) GetValidator(protocolID string) (ProtocolValidator, bool) {
	validator, exists := vr.validators[protocolID]
	return validator, exists
}

// Global validator registry instance
var GlobalValidatorRegistry = NewValidatorRegistry()
