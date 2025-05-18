package types

import (
	"fmt"
	"strings"
)

// ABIInput represents an input or output parameter for an ABI function.
type ABIInput struct {
	Name string
	Type string
	// Components []ABIInput // For tuple types, if needed later
	// Indexed bool      // For event parameters, if needed later
}

// ABIOutput represents an output parameter for an ABI function.
// Often similar to ABIInput, so can be aliased or kept separate for distinction.
type ABIOutput ABIInput

// ABIFunction represents a function definition within an ABI.
type ABIFunction struct {
	Name            string
	Inputs          []ABIInput
	Outputs         []ABIOutput
	StateMutability string // e.g., "view", "nonpayable", "payable"
	Type            string // e.g., "function", "constructor", "event"
	Constant        bool   // True if 'constant' or 'view'/'pure' stateMutability
}

// IsPayable checks if the function can receive Ether.
func (af *ABIFunction) IsPayable() bool {
	return af.StateMutability == "payable"
}

// ABI represents a parsed contract Application Binary Interface.
// It holds the definitions for functions and events.
type ABI struct {
	Functions []ABIFunction
	Events    []ABIFunction // Can reuse ABIFunction for events if structure is similar enough, or define ABIEvent
	// Constructor ABIFunction // Optional: if constructor needs to be specially handled
	RawJson string // Stores the raw JSON string of the ABI, useful for parsing with other libraries
}

// MapTypeToParamType converts an ABI type string to a simplified policy parameter type string.
// This is a placeholder and should be made more robust.
func MapTypeToParamType(abiType string) string {
	// TODO: Implement a more comprehensive mapping based on policy engine's type system
	switch abiType {
	case "address":
		return "address"
	case "bool":
		return "boolean"
	case "string":
		return "string"
	case "bytes":
		return "bytes"
	default:
		if strings.HasPrefix(abiType, "uint") || strings.HasPrefix(abiType, "int") {
			return "decimal" // Or more specific like "uint256", "int64"
		}
		if strings.HasSuffix(abiType, "[]") {
			return "array" // Generic array type
		}
		return "string" // Default or unknown type
	}
}

// GetInputDescription creates a basic description for an ABI input parameter.
// This is a placeholder.
func GetInputDescription(input ABIInput) string {
	return fmt.Sprintf("Parameter '%s' of type '%s'", input.Name, input.Type)
}
