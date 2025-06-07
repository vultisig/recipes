package ethereum

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vultisig/recipes/types"
)

// UniswapV2Validator provides Uniswap v2 specific validation logic
type UniswapV2Validator struct{}

// NewUniswapV2Validator creates a new Uniswap v2 validator
func NewUniswapV2Validator() *UniswapV2Validator {
	return &UniswapV2Validator{}
}

// GetProtocolID returns the protocol ID this validator is for
func (v *UniswapV2Validator) GetProtocolID() string {
	return "uniswapv2_router"
}

// CustomizeFunctions implements ProtocolValidator interface - customizes function descriptions
func (v *UniswapV2Validator) CustomizeFunctions(f *types.Function, abiFunc *ABIFunction) {
	// Enhance function descriptions
	switch abiFunc.Name {
	case "swapExactETHForTokens":
		f.Description = "Swap exact ETH for tokens with minimum output protection"
		v.addSwapValidations(f)
	case "swapExactTokensForETH":
		f.Description = "Swap exact tokens for ETH with minimum output protection"
		v.addSwapValidations(f)
	case "swapExactTokensForTokens":
		f.Description = "Swap exact tokens for other tokens with minimum output protection"
		v.addSwapValidations(f)
	case "addLiquidity":
		f.Description = "Add liquidity to a token pair with slippage protection"
		v.addLiquidityValidations(f)
	case "addLiquidityETH":
		f.Description = "Add liquidity to an ETH/token pair with slippage protection"
		v.addLiquidityValidations(f)
	case "removeLiquidity":
		f.Description = "Remove liquidity from a token pair with minimum output protection"
		v.addLiquidityValidations(f)
	case "removeLiquidityETH":
		f.Description = "Remove liquidity from an ETH/token pair with minimum output protection"
		v.addLiquidityValidations(f)
	}

	// Add deadline validation to all functions that have deadline parameter
	v.addDeadlineValidation(f)

	// Add amount validation to all functions with amount parameters
	v.addAmountValidations(f)

	// Add address validation to all functions with address parameters
	v.addAddressValidations(f)
}

// ValidateTransaction implements ProtocolValidator interface - validates Uniswap transactions
func (v *UniswapV2Validator) ValidateTransaction(functionName string, params map[string]interface{}) error {
	switch functionName {
	case "swapExactETHForTokens", "swapExactTokensForETH", "swapExactTokensForTokens":
		return v.validateSwapTransaction(params)
	case "addLiquidity", "addLiquidityETH":
		return v.validateAddLiquidityTransaction(params)
	case "removeLiquidity", "removeLiquidityETH":
		return v.validateRemoveLiquidityTransaction(params)
	}

	// Common validations for all functions
	if err := v.validateDeadline(params); err != nil {
		return err
	}

	if err := v.validateAmounts(params); err != nil {
		return err
	}

	if err := v.validateAddresses(params); err != nil {
		return err
	}

	return nil
}

// CreateUniswapV2Customizer returns a FunctionCustomizer for Uniswap v2 functions
// DEPRECATED: Use ProtocolValidator interface instead
func (v *UniswapV2Validator) CreateUniswapV2Customizer() FunctionCustomizer {
	return func(f *types.Function, abiFunc *ABIFunction) {
		v.CustomizeFunctions(f, abiFunc)
	}
}

// ValidateUniswapTransaction validates a Uniswap transaction against business rules
// DEPRECATED: Use ValidateTransaction instead
func (v *UniswapV2Validator) ValidateUniswapTransaction(functionName string, params map[string]interface{}) error {
	return v.ValidateTransaction(functionName, params)
}

// addSwapValidations adds swap-specific validation logic
func (v *UniswapV2Validator) addSwapValidations(f *types.Function) {
	for _, param := range f.Parameters {
		switch param.Name {
		case "amountOutMin":
			param.Description = "Minimum amount of output tokens (slippage protection). Should be calculated as: expectedOutput * (1 - slippageTolerance)"
		case "path":
			param.Description = "Array of token addresses representing the swap path. First address is input token, last is output token"
		case "to":
			param.Description = "Address that will receive the output tokens. Should be wallet address or approved contract"
		}
	}
}

// addLiquidityValidations adds liquidity-specific validation logic
func (v *UniswapV2Validator) addLiquidityValidations(f *types.Function) {
	for _, param := range f.Parameters {
		switch param.Name {
		case "amountAMin", "amountBMin", "amountTokenMin", "amountETHMin":
			param.Description = fmt.Sprintf("%s - Minimum amount for slippage protection (typically 95-99%% of desired)", param.Name)
		case "amountADesired", "amountBDesired", "amountTokenDesired":
			param.Description = fmt.Sprintf("%s - Desired amount (maximum you're willing to provide)", param.Name)
		case "liquidity":
			param.Description = "Amount of LP tokens to remove"
		}
	}
}

// addDeadlineValidation adds deadline parameter validation
func (v *UniswapV2Validator) addDeadlineValidation(f *types.Function) {
	for _, param := range f.Parameters {
		if param.Name == "deadline" {
			param.Description = "Unix timestamp deadline for transaction execution. Must be in the future (current time + reasonable buffer)"
		}
	}
}

// addAmountValidations adds amount parameter validation
func (v *UniswapV2Validator) addAmountValidations(f *types.Function) {
	for _, param := range f.Parameters {
		if strings.Contains(param.Name, "amount") || strings.Contains(param.Name, "value") {
			if param.Description == "" || !strings.Contains(param.Description, "Must be positive") {
				param.Description += " (Must be positive and within reasonable bounds)"
			}
		}
	}
}

// addAddressValidations adds address parameter validation
func (v *UniswapV2Validator) addAddressValidations(f *types.Function) {
	for _, param := range f.Parameters {
		if param.Type == "address" {
			switch param.Name {
			case "to":
				param.Description += " (Must be a valid Ethereum address, avoid zero address)"
			case "tokenA", "tokenB", "token":
				param.Description += " (Must be a valid ERC20 token contract address)"
			default:
				if strings.Contains(param.Name, "token") {
					param.Description += " (Must be a valid token contract address)"
				}
			}
		}
	}
}

// validateSwapTransaction validates swap-specific business rules
func (v *UniswapV2Validator) validateSwapTransaction(params map[string]interface{}) error {
	// Validate slippage protection
	if amountOutMin, ok := params["amountOutMin"]; ok {
		if amount, ok := amountOutMin.(*big.Int); ok {
			if amount.Cmp(big.NewInt(0)) <= 0 {
				return fmt.Errorf("amountOutMin must be positive for slippage protection")
			}
		}
	}

	// Validate swap path
	if path, ok := params["path"]; ok {
		if pathArray, ok := path.([]interface{}); ok {
			if len(pathArray) < 2 {
				return fmt.Errorf("swap path must contain at least 2 addresses (input and output tokens)")
			}
			if len(pathArray) > 4 {
				return fmt.Errorf("swap path too long, maximum 4 tokens supported for optimal gas usage")
			}

			// Validate each address in path
			for i, addr := range pathArray {
				if addrStr, ok := addr.(string); ok {
					if !common.IsHexAddress(addrStr) {
						return fmt.Errorf("invalid token address at path position %d: %s", i, addrStr)
					}
					if addrStr == "0x0000000000000000000000000000000000000000" {
						return fmt.Errorf("zero address not allowed in swap path at position %d", i)
					}
				}
			}
		}
	}

	return nil
}

// validateAddLiquidityTransaction validates add liquidity business rules
func (v *UniswapV2Validator) validateAddLiquidityTransaction(params map[string]interface{}) error {
	// Validate that minimum amounts are reasonable (should be 95-99% of desired)
	desiredA := v.getBigIntParam(params, "amountADesired")
	minA := v.getBigIntParam(params, "amountAMin")

	if desiredA != nil && minA != nil {
		// Calculate ratio: minA should be at least 90% of desiredA
		ratio := new(big.Int).Mul(minA, big.NewInt(100))
		ratio.Div(ratio, desiredA)

		if ratio.Cmp(big.NewInt(90)) < 0 {
			return fmt.Errorf("amountAMin too low compared to amountADesired (less than 90%%), increase slippage protection")
		}
	}

	return nil
}

// validateRemoveLiquidityTransaction validates remove liquidity business rules
func (v *UniswapV2Validator) validateRemoveLiquidityTransaction(params map[string]interface{}) error {
	// Validate liquidity amount
	if liquidity := v.getBigIntParam(params, "liquidity"); liquidity != nil {
		if liquidity.Cmp(big.NewInt(0)) <= 0 {
			return fmt.Errorf("liquidity amount must be positive")
		}
	}

	return nil
}

// validateDeadline validates deadline parameter
func (v *UniswapV2Validator) validateDeadline(params map[string]interface{}) error {
	if deadline := v.getBigIntParam(params, "deadline"); deadline != nil {
		currentTime := big.NewInt(time.Now().Unix())

		if deadline.Cmp(currentTime) <= 0 {
			return fmt.Errorf("deadline must be in the future (current: %s, deadline: %s)", currentTime.String(), deadline.String())
		}

		// Warn if deadline is too far in the future (more than 1 hour)
		oneHourFromNow := new(big.Int).Add(currentTime, big.NewInt(3600))
		if deadline.Cmp(oneHourFromNow) > 0 {
			return fmt.Errorf("deadline too far in future (more than 1 hour), this may be unsafe")
		}
	}

	return nil
}

// validateAmounts validates all amount parameters are positive
func (v *UniswapV2Validator) validateAmounts(params map[string]interface{}) error {
	for paramName, _ := range params {
		if strings.Contains(paramName, "amount") || strings.Contains(paramName, "value") {
			if amount := v.getBigIntParam(params, paramName); amount != nil {
				if amount.Cmp(big.NewInt(0)) <= 0 {
					return fmt.Errorf("parameter %s must be positive, got: %s", paramName, amount.String())
				}
			}
		}
	}

	return nil
}

// validateAddresses validates all address parameters
func (v *UniswapV2Validator) validateAddresses(params map[string]interface{}) error {
	for paramName, paramValue := range params {
		if strings.Contains(strings.ToLower(paramName), "to") || strings.Contains(strings.ToLower(paramName), "token") {
			if addrStr, ok := paramValue.(string); ok {
				if !common.IsHexAddress(addrStr) {
					return fmt.Errorf("parameter %s must be a valid Ethereum address, got: %s", paramName, addrStr)
				}
				if addrStr == "0x0000000000000000000000000000000000000000" {
					return fmt.Errorf("parameter %s cannot be zero address", paramName)
				}
			}
		}
	}

	return nil
}

// Helper function to safely get *big.Int from params
func (v *UniswapV2Validator) getBigIntParam(params map[string]interface{}, paramName string) *big.Int {
	if val, ok := params[paramName]; ok {
		switch v := val.(type) {
		case *big.Int:
			return v
		case string:
			if amount, ok := new(big.Int).SetString(v, 10); ok {
				return amount
			}
		}
	}
	return nil
}
