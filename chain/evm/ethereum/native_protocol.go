package ethereum

import (
	"fmt"
	"strings"

	"github.com/vultisig/recipes/types"
)

// NativeEVMProtocol implements a native currency protocol for any EVM chain.
// This provides shared MatchFunctionCall logic that all EVM chains can reuse.
type NativeEVMProtocol struct {
	BaseProtocol
	nativeSymbol string // e.g., "eth", "bnb"
}

// NewNativeEVMProtocol creates a native currency protocol for any EVM chain.
// Parameters:
//   - chainID: chain identifier (e.g., "ethereum", "arbitrum", "bsc")
//   - nativeSymbol: the native token symbol (e.g., "eth", "bnb")
//   - name: human readable name (e.g., "Arbitrum ETH", "BSC BNB")
//   - description: protocol description
func NewNativeEVMProtocol(chainID, nativeSymbol, name, description string) types.Protocol {
	return &NativeEVMProtocol{
		BaseProtocol: BaseProtocol{
			id:          strings.ToLower(nativeSymbol),
			name:        name,
			description: description,
			functions: []*types.Function{
				{
					ID:          "transfer",
					Name:        fmt.Sprintf("Transfer %s", strings.ToUpper(nativeSymbol)),
					Description: fmt.Sprintf("Transfer %s to another address", strings.ToUpper(nativeSymbol)),
					Parameters: []*types.FunctionParam{
						{Name: "recipient", Type: "address", Description: fmt.Sprintf("The %s address of the recipient", chainID)},
						{Name: "amount", Type: "decimal", Description: fmt.Sprintf("The amount of %s to transfer", strings.ToUpper(nativeSymbol))},
					},
				},
			},
		},
		nativeSymbol: strings.ToLower(nativeSymbol),
	}
}

// ChainID returns the chain identifier - override in embedding struct if needed
func (p *NativeEVMProtocol) ChainID() string {
	return p.BaseProtocol.ChainID()
}

// MatchFunctionCall implements the matching logic for native EVM transfers.
// This is shared across all EVM chains.
// Note: policyMatcher.ResourcePath must be non-nil for this function to work correctly.
func (p *NativeEVMProtocol) MatchFunctionCall(decodedTx types.DecodedTransaction, policyMatcher *types.PolicyFunctionMatcher) (bool, map[string]interface{}, error) {
	if policyMatcher.FunctionID != "transfer" {
		return false, nil, fmt.Errorf("%s protocol only supports 'transfer' function for matching, got %s", strings.ToUpper(p.nativeSymbol), policyMatcher.FunctionID)
	}

	// For native transfer, transaction data should be empty.
	if len(decodedTx.Data()) > 0 {
		return false, nil, nil // Not a simple native transfer if data is present
	}

	if policyMatcher.ResourcePath == nil {
		return false, nil, fmt.Errorf("policy matcher ResourcePath must be set for %s native transfer matching", strings.ToUpper(p.nativeSymbol))
	}

	extractedParams := map[string]interface{}{
		"recipient": strings.ToLower(decodedTx.To()),
		"amount":    decodedTx.Value(),
		"chainId":   policyMatcher.ResourcePath.ChainId,
		"asset":     p.nativeSymbol,
	}

	constraintsMet, err := evaluateParameterConstraints(extractedParams, policyMatcher.Constraints)
	if err != nil {
		return false, nil, fmt.Errorf("error evaluating constraints for %s transfer: %w", strings.ToUpper(p.nativeSymbol), err)
	}

	if !constraintsMet {
		return false, nil, nil
	}

	return true, extractedParams, nil
}

