package ethereum

import (
	"encoding/json"
	"fmt"
)

// ABIInput represents an input parameter to a function
type ABIInput struct {
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	Indexed      bool       `json:"indexed,omitempty"`
	Components   []ABIInput `json:"components,omitempty"`
	InternalType string     `json:"internalType,omitempty"`
}

// ABIOutput represents an output parameter from a function
type ABIOutput struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Components   []ABIOutput `json:"components,omitempty"`
	InternalType string      `json:"internalType,omitempty"`
}

// ABIFunction represents a function in an ABI
type ABIFunction struct {
	Name            string      `json:"name"`
	Type            string      `json:"type"`
	Inputs          []ABIInput  `json:"inputs"`
	Outputs         []ABIOutput `json:"outputs,omitempty"`
	StateMutability string      `json:"stateMutability,omitempty"`
	Constant        bool        `json:"constant,omitempty"`
	Payable         bool        `json:"payable,omitempty"`
}

// ABI represents an Ethereum ABI
type ABI struct {
	Functions []ABIFunction
	RawJson   string
}

// ParseABI parses an ABI JSON into an ABI struct
func ParseABI(abiJSON []byte) (*ABI, error) {
	var abiFunctions []ABIFunction
	if err := json.Unmarshal(abiJSON, &abiFunctions); err != nil {
		return nil, fmt.Errorf("error unmarshaling ABI: %w", err)
	}

	// Filter to only include functions (not events, constructors, etc.)
	var functions []ABIFunction
	for _, fn := range abiFunctions {
		if fn.Type == "function" {
			functions = append(functions, fn)
		}
	}

	return &ABI{
		Functions: functions,
		RawJson:   string(abiJSON),
	}, nil
}

// GetFunction returns a function by name
func (a *ABI) GetFunction(name string) (*ABIFunction, bool) {
	for _, fn := range a.Functions {
		if fn.Name == name {
			return &fn, true
		}
	}
	return nil, false
}

// MapTypeToParamType maps Ethereum types to policy parameter types
func MapTypeToParamType(ethType string) string {
	switch {
	case ethType == "address":
		return "address"
	case ethType == "bool":
		return "boolean"
	case ethType == "string":
		return "string"
	case ethType == "bytes":
		return "bytes"
	case len(ethType) >= 5 && ethType[:5] == "bytes":
		return "bytes" // bytes1, bytes32, etc.
	case ethType == "uint256" || ethType == "int256" ||
		len(ethType) >= 4 && (ethType[:4] == "uint" || ethType[:3] == "int"):
		return "number"
	default:
		// For complex types (arrays, tuples, etc.), return "complex"
		return "complex"
	}
}

// GetInputDescription generates a human-readable description for an input
func GetInputDescription(input ABIInput) string {
	typeName := input.Type
	if input.InternalType != "" {
		typeName = input.InternalType
	}

	return fmt.Sprintf("%s parameter of type %s", input.Name, typeName)
}

// IsPayable returns whether the function is payable
func (f *ABIFunction) IsPayable() bool {
	return f.Payable || f.StateMutability == "payable"
}

// IsView returns whether the function is a view function
func (f *ABIFunction) IsView() bool {
	return f.Constant || f.StateMutability == "view" || f.StateMutability == "pure"
}
