package ethereum

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// BaseProtocol provides common functionality for Ethereum protocols
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
	return "ethereum"
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
	return nil, fmt.Errorf("function %q not found", id)
}

// ETH implements the native Ethereum protocol
type ETH struct {
	BaseProtocol
}

// NewETH creates a new ETH protocol
func NewETH() types.Protocol {
	functions := []*types.Function{
		{
			ID:          "transfer",
			Name:        "Transfer ETH",
			Description: "Transfer Ether to another address",
			Parameters: []*types.FunctionParam{
				{
					Name:        "recipient",
					Type:        "address",
					Description: "The Ethereum address of the recipient",
				},
				{
					Name:        "amount",
					Type:        "decimal",
					Description: "The amount of Ether to transfer",
				},
			},
		},
	}

	return &ETH{
		BaseProtocol: BaseProtocol{
			id:          "eth",
			name:        "Ethereum",
			description: "The native cryptocurrency of the Ethereum blockchain",
			functions:   functions,
		},
	}
}

// ABIProtocol implements a protocol defined by an ABI
type ABIProtocol struct {
	BaseProtocol
	abi *ABI
}

// FunctionCustomizer is a function that customizes a generated function
type FunctionCustomizer func(f *types.Function, abiFunc *ABIFunction)

// NewABIProtocolWithCustomization creates a new protocol from an ABI with customization
func NewABIProtocolWithCustomization(id string, name string, description string, abi *ABI, customizer FunctionCustomizer) types.Protocol {
	functions := make([]*types.Function, 0, len(abi.Functions))

	// Convert ABI functions to policy functions
	for _, abiFunc := range abi.Functions {
		// Skip internal functions that start with an underscore
		if len(abiFunc.Name) > 0 && abiFunc.Name[0] == '_' {
			continue
		}

		// Skip non-externally callable functions
		if abiFunc.StateMutability == "internal" || abiFunc.StateMutability == "private" {
			continue
		}

		// Create function parameters
		params := make([]*types.FunctionParam, 0, len(abiFunc.Inputs))
		for _, input := range abiFunc.Inputs {
			paramType := MapTypeToParamType(input.Type)
			paramName := input.Name
			if paramName == "" {
				paramName = fmt.Sprintf("param%d", len(params))
			}

			params = append(params, &types.FunctionParam{
				Name:        paramName,
				Type:        paramType,
				Description: GetInputDescription(input),
			})
		}

		// Add value parameter for payable functions
		if abiFunc.IsPayable() {
			params = append(params, &types.FunctionParam{
				Name:        "value",
				Type:        "decimal",
				Description: "The amount of ETH to send with the transaction",
			})
		}

		// Create the function
		function := &types.Function{
			ID:          abiFunc.Name,
			Name:        fmt.Sprintf("%s.%s", id, abiFunc.Name),
			Description: fmt.Sprintf("Call the %s function on %s", abiFunc.Name, name),
			Parameters:  params,
		}

		// Apply customization if provided
		if customizer != nil {
			customizer(function, &abiFunc)
		}

		functions = append(functions, function)
	}

	return &ABIProtocol{
		BaseProtocol: BaseProtocol{
			id:          id,
			name:        name,
			description: description,
			functions:   functions,
		},
		abi: abi,
	}
}

// NewABIProtocol creates a new protocol from an ABI
func NewABIProtocol(id string, name string, description string, abi *ABI) types.Protocol {
	return NewABIProtocolWithCustomization(id, name, description, abi, nil)
}
