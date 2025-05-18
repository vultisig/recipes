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

// Protocol defines the interface for a specific protocol on a chain (e.g., ETH, ERC20)
type Protocol interface {
	// ID returns the unique identifier for the protocol (e.g., "eth", "usdc")
	ID() string
	// Name returns a human-readable name for the protocol (e.g., "Ethereum", "USD Coin")
	Name() string
	// ChainID returns the identifier of the chain this protocol belongs to
	ChainID() string
	// Description returns a brief description of the protocol
	Description() string
	// Functions returns a list of supported functions by this protocol
	Functions() []*Function
	// GetFunction returns a specific function by its ID
	GetFunction(id string) (*Function, error)

	// MatchFunctionCall checks if a decoded transaction matches the criteria for a specific function call
	// defined in a policy. It returns true if it matches, along with extracted parameters relevant
	// to the function call, or an error if the matching process fails.
	MatchFunctionCall(decodedTx DecodedTransaction, policyMatcher *PolicyFunctionMatcher) (matches bool, extractedParams map[string]interface{}, err error)
}

// ResourcePath builds a fully qualified resource path for a function
func FullyQualifiedResourcePath(chainID, protocolID, functionID string) string {
	return fmt.Sprintf("%s.%s.%s", chainID, protocolID, functionID)
}
