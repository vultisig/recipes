package types

import (
	"fmt"
)

// FunctionParam represents a parameter for a protocol function
type FunctionParam struct {
	// Name of the parameter
	Name string

	// Type of the parameter (e.g., "address", "amount", etc.)
	Type string

	// Description provides details about the parameter's purpose
	Description string
}

// Function represents an action that can be performed on a protocol
type Function struct {
	// ID is the unique identifier for the function
	ID string

	// Name is a human-readable name for the function
	Name string

	// Description provides details about what the function does
	Description string

	// Parameters is a list of parameters that the function accepts
	Parameters []*FunctionParam
}

// Protocol represents a cryptocurrency protocol within a chain
type Protocol interface {
	// ID returns the unique identifier for the protocol
	ID() string

	// Name returns a human-readable name for the protocol
	Name() string

	// ChainID returns the ID of the chain this protocol belongs to
	ChainID() string

	// Description returns a detailed description of the protocol
	Description() string

	// Functions returns a list of available functions for this protocol
	Functions() []*Function

	// GetFunction retrieves a specific function by ID
	GetFunction(id string) (*Function, error)
}

// ResourcePath builds a fully qualified resource path for a function
func ResourcePath(chainID, protocolID, functionID string) string {
	return fmt.Sprintf("%s.%s.%s", chainID, protocolID, functionID)
}
