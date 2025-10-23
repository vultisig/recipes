package ethereum

import (
	"fmt"
	
	"github.com/vultisig/recipes/types"
)

type AngleMerklDistributor struct {}

func NewAngleMerklDistributor() *AngleMerklDistributor {
	return &AngleMerklDistributor{}
}

func (a *AngleMerklDistributor) GetProtocolID() string {
	return "angle_merkl_distributor"
}

func (a *AngleMerklDistributor) CustomizeFunctions(f *types.Function, abiFunc *ABIFunction) {
	switch abiFunc.Name {
	case "claim":
		f.Description = "Claim tokens from the merkl distributor"
		a.addClaimValidations(f)
	}

}

func (a *AngleMerklDistributor) ValidateTransaction(functionName string, params map[string]interface{}) error {
	switch functionName {
	case "claim":
		return a.validateClaimTransaction(params)
	}

	return nil
}

func (a *AngleMerklDistributor) addClaimValidations(f *types.Function) {
	for _, param := range f.Parameters {
		switch param.Name {
		case "users":
			param.Description = "Array of addresses to claim tokens for"
		case "tokens":
			param.Description = "Array of token addresses to claim"
		case "amounts":
			param.Description = "Array of amounts to claim"
		case "proofs":
			param.Description = "Array of proofs to claim"
		}
	}
}

func (a *AngleMerklDistributor) validateClaimTransaction(params map[string]interface{}) error {
	// Get users
	_, ok := params["users"]
	if !ok {
		return fmt.Errorf("users are required")
	}

	return nil
}