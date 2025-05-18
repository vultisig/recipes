package ethereum

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vultisig/recipes/types"
)

// evaluateParameterConstraints evaluates policy constraints against extracted transaction parameters.
// This function assumes that types.ParameterConstraint is correctly defined and available from the types package (e.g., via types/policy.pb.go).
func evaluateParameterConstraints(params map[string]interface{}, policyParamConstraints []*types.ParameterConstraint) (bool, error) {
	for _, pc := range policyParamConstraints {
		if pc == nil { // Added defensive nil check for the ParameterConstraint itself
			// This case should ideally not happen if policyParamConstraints comes directly from a valid protobuf message.
			// Depending on policy, this could be a strict error or a skippable entry.
			// For now, let's log a warning and skip, or return an error if strictness is required.
			fmt.Println("Warning: encountered a nil ParameterConstraint entry, skipping.")
			continue
		}
		paramName := pc.ParameterName
		constraint := pc.Constraint

		if constraint == nil {
			// Or log a warning, depends on how strict the policy interpretation should be
			return false, fmt.Errorf("nil constraint found for parameter %q", paramName)
		}

		paramValue, ok := params[paramName]
		if !ok {
			if constraint.Required { // Check if the constraint proto has a 'Required' field
				return false, fmt.Errorf("required parameter %q not found in transaction for constraint check", paramName)
			}
			continue // Parameter not present, and not strictly required by this constraint, skip to next constraint
		}

		switch cVal := constraint.Value.(type) { // Value is the oneof field
		case *types.Constraint_FixedValue:
			valStr, isStr := paramValueToString(paramValue)
			if !isStr {
				return false, fmt.Errorf("parameter %q (type %T) could not be converted to string for FixedValue comparison", paramName, paramValue)
			}
			if !strings.EqualFold(valStr, cVal.FixedValue) {
				return false, nil // Constraint not met
			}
		case *types.Constraint_WhitelistValues:
			valStr, isStr := paramValueToString(paramValue)
			if !isStr {
				return false, fmt.Errorf("parameter %q (type %T) could not be converted to string for Whitelist comparison", paramName, paramValue)
			}
			found := false
			for _, allowedVal := range cVal.WhitelistValues.GetValues() {
				if strings.EqualFold(valStr, allowedVal) {
					found = true
					break
				}
			}
			if !found {
				return false, nil // Not in whitelist
			}
		// TODO: Implement other constraint types from types.ConstraintType and oneof cases:
		// MaxValue, MinValue, RangeValue, MaxPerPeriodValue
		// These will require parsing string values from constraints (e.g., cVal.MaxValue) into appropriate types (e.g., *big.Int)
		// and comparing them with paramValue. Ensure robust type handling and comparisons.

		// Example for MaxValue (assuming paramValue is *big.Int and MaxValue is string representation of a number)
		case *types.Constraint_MaxValue:
			actualValueBigInt, ok := paramValue.(*big.Int)
			if !ok {
				return false, fmt.Errorf("parameter %q expected to be *big.Int for MaxValue, got %T", paramName, paramValue)
			}
			maxValueBigInt, ok := new(big.Int).SetString(cVal.MaxValue, 10)
			if !ok {
				return false, fmt.Errorf("invalid MaxValue string %q in policy for parameter %q", cVal.MaxValue, paramName)
			}
			if actualValueBigInt.Cmp(maxValueBigInt) > 0 { // actual > max
				return false, nil // Constraint not met
			}

		default:
			return false, fmt.Errorf("unsupported constraint value type: %T for parameter %q", constraint.Value, paramName)
		}
	}
	return true, nil // All parameter constraints met
}

// paramValueToString converts various parameter types to a string for comparison.
// Handles common.Address and *big.Int specifically.
func paramValueToString(paramValue interface{}) (string, bool) {
	switch v := paramValue.(type) {
	case string:
		return v, true
	case common.Address:
		return v.Hex(), true
	case *common.Address:
		if v == nil {
			return "", false
		} // Or handle as empty string if appropriate
		return v.Hex(), true
	case *big.Int:
		if v == nil {
			return "", false
		} // Or handle as "0" or error
		return v.String(), true
	default:
		// For other types, attempt a generic string conversion. This might not always be suitable.
		// Consider adding more specific type handling if needed.
		strVal := fmt.Sprintf("%v", v)
		return strVal, true // Assuming conversion is generally possible, though might not be canonical
	}
}

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
	return nil, fmt.Errorf("function %q not found in protocol %q", id, p.id)
}

// ETH implements the native Ethereum protocol
type ETH struct {
	BaseProtocol
}

// NewETH creates a new ETH protocol
func NewETH() types.Protocol {
	return &ETH{
		BaseProtocol: BaseProtocol{
			id:          "eth",
			name:        "Ethereum",
			description: "Native Ether currency of the Ethereum blockchain",
			functions: []*types.Function{
				{
					ID:          "transfer",
					Name:        "Transfer ETH",
					Description: "Transfer Ether to another address",
					Parameters: []*types.FunctionParam{
						{Name: "recipient", Type: "address", Description: "The Ethereum address of the recipient"},
						{Name: "amount", Type: "decimal", Description: "The amount of Ether to transfer"},
					},
				},
			},
		},
	}
}

// MatchFunctionCall for ETH protocol (native Ether transfer)
func (p *ETH) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	if policyMatcher.FunctionID != "transfer" {
		return false, nil, fmt.Errorf("ETH protocol only supports 'transfer' function for matching, got %s", policyMatcher.FunctionID)
	}

	// For native ETH transfer, transaction data should be empty.
	if len(decodedTx.Data()) > 0 {
		return false, nil, nil // Not a simple ETH transfer if data is present
	}

	extractedParams := map[string]interface{}{
		"recipient": strings.ToLower(decodedTx.To()), // Normalize to lowercase for policy matching
		"amount":    decodedTx.Value(),
		// "from":      strings.ToLower(decodedTx.From()), // Can also include sender if policies need it
	}

	// The PolicyFunctionMatcher holds ParameterConstraints directly now (hypothetically)
	// Or, the policy structure needs to be passed more directly to evaluateParameterConstraints.
	// The PolicyFunctionMatcher should contain []*types.ParameterConstraint if this is the common structure.
	// Let's assume PolicyFunctionMatcher.Constraints is actually of type []*types.ParameterConstraint.
	// This requires changing types/matcher.go definition for PolicyFunctionMatcher.
	// For now, I'll assume policyMatcher.Constraints is what evaluateParameterConstraints expects.
	// If policyMatcher.Constraints is still []*types.Constraint, we need a different approach.
	// Given PolicyRule has repeated ParameterConstraint, it is likely that a PolicyFunctionMatcher
	// would aggregate these. Let's stick to the assumption that policyMatcher.Constraints is []*types.ParameterConstraint

	constraintsMet, err := evaluateParameterConstraints(extractedParams, policyMatcher.Constraints) // This line assumes policyMatcher.Constraints is []*types.ParameterConstraint
	if err != nil {
		return false, nil, fmt.Errorf("error evaluating constraints for ETH transfer: %w", err)
	}

	if !constraintsMet {
		return false, nil, nil // Constraints not met
	}

	return true, extractedParams, nil
}

// ABIProtocol implements a protocol defined by an ABI
type ABIProtocol struct {
	BaseProtocol
	abiParsed *abi.ABI // Store the parsed go-ethereum ABI object
}

// FunctionCustomizer is a function that customizes a generated function
type FunctionCustomizer func(f *types.Function, abiFunc *ABIFunction)

// NewABIProtocolWithCustomization creates a new protocol from an ABI with customization
func NewABIProtocolWithCustomization(id string, name string, description string, localDomainABI *ABI, customizer FunctionCustomizer) types.Protocol {
	parsedGoEthABI, err := abi.JSON(strings.NewReader(localDomainABI.RawJson)) // Assuming ABI struct has RawJson string
	if err != nil {
		// Handle error: couldn't parse ABI. Maybe panic or return an error-protocol type.
		// For now, let's log and proceed, but MatchFunctionCall will fail.
		fmt.Printf("Error parsing ABI for %s from RawJson: %v. Protocol will not be functional.\n", id, err)
	}

	functions := make([]*types.Function, 0, len(localDomainABI.Functions))
	for _, domainAbiFunc := range localDomainABI.Functions {
		if len(domainAbiFunc.Name) > 0 && domainAbiFunc.Name[0] == '_' {
			continue
		}
		if domainAbiFunc.StateMutability == "internal" || domainAbiFunc.StateMutability == "private" {
			continue
		}

		params := make([]*types.FunctionParam, 0, len(domainAbiFunc.Inputs))
		for _, domainAbiInput := range domainAbiFunc.Inputs {
			paramType := types.MapTypeToParamType(domainAbiInput.Type)
			paramName := domainAbiInput.Name
			if paramName == "" {
				paramName = fmt.Sprintf("param%d", len(params))
			}
			params = append(params, &types.FunctionParam{
				Name: paramName, Type: paramType, Description: GetInputDescription(domainAbiInput),
			})
		}
		if domainAbiFunc.IsPayable() {
			params = append(params, &types.FunctionParam{
				Name: "value", Type: "decimal", Description: "The amount of ETH to send with the transaction",
			})
		}
		function := &types.Function{
			ID: domainAbiFunc.Name, Name: fmt.Sprintf("%s.%s", id, domainAbiFunc.Name),
			Description: fmt.Sprintf("Call the %s function on %s", domainAbiFunc.Name, name),
			Parameters:  params,
		}
		if customizer != nil {
			customizer(function, &domainAbiFunc)
		}
		functions = append(functions, function)
	}

	return &ABIProtocol{
		BaseProtocol: BaseProtocol{id: id, name: name, description: description, functions: functions},
		abiParsed:    &parsedGoEthABI,
	}
}

// NewABIProtocol creates a new protocol from an ABI
func NewABIProtocol(id string, name string, description string, localDomainABI *ABI) types.Protocol {
	return NewABIProtocolWithCustomization(id, name, description, localDomainABI, nil)
}

// MatchFunctionCall for ABIProtocol (e.g., ERC20 transfer)
func (p *ABIProtocol) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	if p.abiParsed == nil || p.abiParsed.Methods == nil {
		return false, nil, fmt.Errorf("ABI for protocol %q not parsed or empty, cannot match function call", p.ID())
	}

	abiMethod, methodExists := p.abiParsed.Methods[policyMatcher.FunctionID]
	if !methodExists {
		return false, nil, nil // Policy requests a function not in this ABI protocol
	}

	txData := decodedTx.Data()
	if len(txData) < 4 { // Method ID is 4 bytes
		return false, nil, nil // Data too short to contain a method ID
	}

	// Check if the method ID in the transaction data matches the expected method ID
	if !bytes.Equal(txData[:4], abiMethod.ID) {
		return false, nil, nil // Method ID mismatch
	}

	// Unpack parameters from transaction data (excluding the 4-byte method ID)
	rawParamsData := txData[4:]
	unpackedParamsMap := make(map[string]interface{}) // For go-ethereum/accounts/abi UnpackIntoMap
	err := abiMethod.Inputs.UnpackIntoMap(unpackedParamsMap, rawParamsData)
	if err != nil {
		return false, nil, fmt.Errorf("failed to unpack ABI data for method %s: %w. Data: %s", policyMatcher.FunctionID, err, hex.EncodeToString(rawParamsData))
	}

	// Normalize parameter names (e.g. _to -> to) and prepare for constraint evaluation
	extractedParams := make(map[string]interface{}) // For policy constraints
	for _, input := range abiMethod.Inputs {
		paramName := input.Name
		// If param name from ABI starts with _, trim it for matching against policy constraints (e.g. _to -> to)
		if strings.HasPrefix(paramName, "_") {
			paramName = strings.TrimPrefix(paramName, "_")
		}

		// The unpackedParamsMap uses original names from ABI. We need to handle this. Input.Name is reliable.
		val, valOk := unpackedParamsMap[input.Name]
		if !valOk {
			// This case should ideally not happen if UnpackIntoMap succeeded and ABI is consistent.
			// However, if inputs have no names in ABI, UnpackIntoMap might behave differently or require indexed access.
			// For named inputs (common), this should be fine.
			return false, nil, fmt.Errorf("parameter %s (original: %s) not found in unpacked data, though UnpackIntoMap succeeded", paramName, input.Name)
		}

		// Ensure addresses are consistently cased for policy checks (e.g., lowercase hex)
		if addr, ok := val.(common.Address); ok {
			extractedParams[paramName] = strings.ToLower(addr.Hex())
		} else {
			extractedParams[paramName] = val
		}
	}

	// If the function is payable, the transaction value might be a constrained parameter
	if abiMethod.Payable {
		extractedParams["value"] = decodedTx.Value() // Policy can constrain "value"
	}

	// Again, this assumes policyMatcher.Constraints is []*types.ParameterConstraint
	constraintsMet, err := evaluateParameterConstraints(extractedParams, policyMatcher.Constraints)
	if err != nil {
		return false, nil, fmt.Errorf("error evaluating constraints for ABI method %s: %w", policyMatcher.FunctionID, err)
	}

	if !constraintsMet {
		return false, nil, nil // Constraints not met
	}

	return true, extractedParams, nil
}

// End of ABIProtocol MatchFunctionCall. No more code should follow.
