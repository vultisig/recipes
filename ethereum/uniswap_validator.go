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
	case "swapETHForExactTokens":
		f.Description = "Swap ETH for exact amount of tokens with maximum input protection"
		v.addSwapValidations(f)
	case "swapTokensForExactETH":
		f.Description = "Swap tokens for exact ETH with maximum input protection"
		v.addSwapValidations(f)
	case "swapTokensForExactTokens":
		f.Description = "Swap tokens for exact amount of other tokens with maximum input protection"
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
	case "removeLiquidityETHSupportingFeeOnTransferTokens":
		f.Description = "Remove liquidity from ETH/token pair supporting fee-on-transfer tokens"
		v.addLiquidityValidations(f)
	case "removeLiquidityWithPermit":
		f.Description = "Remove liquidity using permit signature for gas-less approval"
		v.addLiquidityValidations(f)
	case "removeLiquidityETHWithPermit":
		f.Description = "Remove ETH/token liquidity using permit signature"
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
	case "swapExactETHForTokens", "swapExactTokensForETH", "swapExactTokensForTokens",
		"swapETHForExactTokens", "swapTokensForExactETH", "swapTokensForExactTokens",
		"swapExactETHForTokensSupportingFeeOnTransferTokens",
		"swapExactTokensForETHSupportingFeeOnTransferTokens",
		"swapExactTokensForTokensSupportingFeeOnTransferTokens":
		return v.validateSwapTransaction(functionName, params)
	case "addLiquidity", "addLiquidityETH":
		return v.validateAddLiquidityTransaction(params)
	case "removeLiquidity", "removeLiquidityETH", "removeLiquidityETHSupportingFeeOnTransferTokens",
		"removeLiquidityWithPermit", "removeLiquidityETHWithPermit",
		"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens":
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
		case "amountInMax":
			param.Description = "Maximum amount of input tokens willing to spend (slippage protection for exact output swaps)"
		case "amountOut":
			param.Description = "Exact amount of output tokens desired"
		case "amountIn":
			param.Description = "Exact amount of input tokens to swap"
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
		case "tokenA", "tokenB", "token":
			param.Description = fmt.Sprintf("%s - Token contract address for liquidity pair", param.Name)
		case "liquidity":
			param.Description = "Amount of LP tokens to remove"
		case "approveMax":
			param.Description = "Whether to approve maximum uint256 amount for permit"
		case "v", "r", "s":
			param.Description = fmt.Sprintf("%s component of permit signature for gas-less approval", param.Name)
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
func (v *UniswapV2Validator) validateSwapTransaction(functionName string, params map[string]interface{}) error {
	// Validate slippage protection based on swap type
	switch functionName {
	case "swapExactETHForTokens", "swapExactTokensForETH", "swapExactTokensForTokens",
		"swapExactETHForTokensSupportingFeeOnTransferTokens",
		"swapExactTokensForETHSupportingFeeOnTransferTokens",
		"swapExactTokensForTokensSupportingFeeOnTransferTokens":
		// For exact input swaps, validate amountOutMin
		if amountOutMin, ok := params["amountOutMin"]; ok {
			if amount, ok := amountOutMin.(*big.Int); ok {
				if amount.Cmp(big.NewInt(0)) <= 0 {
					return fmt.Errorf("amountOutMin must be positive for slippage protection")
				}
			}
		}
	case "swapETHForExactTokens", "swapTokensForExactETH", "swapTokensForExactTokens":
		// For exact output swaps, validate amountInMax and amountOut
		if amountInMax, ok := params["amountInMax"]; ok {
			if amount, ok := amountInMax.(*big.Int); ok {
				if amount.Cmp(big.NewInt(0)) <= 0 {
					return fmt.Errorf("amountInMax must be positive for slippage protection")
				}
			}
		}
		if amountOut, ok := params["amountOut"]; ok {
			if amount, ok := amountOut.(*big.Int); ok {
				if amount.Cmp(big.NewInt(0)) <= 0 {
					return fmt.Errorf("amountOut must be positive")
				}
			}
		}
	}

	// Validate swap path
	if path, ok := params["path"]; ok {
		var pathAddresses []string
		var pathLength int

		// Handle both []interface{} and []common.Address types
		if pathArray, ok := path.([]interface{}); ok {
			pathLength = len(pathArray)
			if pathLength < 2 {
				return fmt.Errorf("swap path must contain at least 2 addresses (input and output tokens)")
			}
			if pathLength > 4 {
				return fmt.Errorf("swap path too long, maximum 4 tokens supported for optimal gas usage")
			}

			// Convert []interface{} to []string
			pathAddresses = make([]string, pathLength)
			for i, addr := range pathArray {
				if addrStr, ok := addr.(string); ok {
					pathAddresses[i] = strings.ToLower(addrStr)
				} else {
					return fmt.Errorf("invalid address type at path position %d: expected string, got %T", i, addr)
				}
			}
		} else if pathArray, ok := path.([]common.Address); ok {
			pathLength = len(pathArray)
			if pathLength < 2 {
				return fmt.Errorf("swap path must contain at least 2 addresses (input and output tokens)")
			}
			if pathLength > 4 {
				return fmt.Errorf("swap path too long, maximum 4 tokens supported for optimal gas usage")
			}

			// Convert []common.Address to []string
			pathAddresses = make([]string, pathLength)
			for i, addr := range pathArray {
				pathAddresses[i] = strings.ToLower(addr.Hex())
			}
		} else {
			return fmt.Errorf("invalid path type: expected []interface{} or []common.Address, got %T", path)
		}

		// Validate each address in path
		for i, addrStr := range pathAddresses {
			if !common.IsHexAddress(addrStr) {
				return fmt.Errorf("invalid token address at path position %d: %s", i, addrStr)
			}
			if addrStr == "0x0000000000000000000000000000000000000000" {
				return fmt.Errorf("zero address not allowed in swap path at position %d", i)
			}
		}
	}

	// Common validations for swap transactions
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

// validateAddLiquidityTransaction validates add liquidity business rules
func (v *UniswapV2Validator) validateAddLiquidityTransaction(params map[string]interface{}) error {
	// Validate that minimum amounts are reasonable (should be 95-99% of desired)
	desiredA := v.getBigIntParam(params, "amountADesired")
	minA := v.getBigIntParam(params, "amountAMin")
	desiredToken := v.getBigIntParam(params, "amountTokenDesired")
	minToken := v.getBigIntParam(params, "amountTokenMin")

	// Check ratio for token A (in addLiquidity)
	if desiredA != nil && minA != nil {
		if err := v.validateLiquidityRatio(desiredA, minA, "amountADesired", "amountAMin"); err != nil {
			return err
		}
	}

	// Check ratio for token (in addLiquidityETH)
	if desiredToken != nil && minToken != nil {
		if err := v.validateLiquidityRatio(desiredToken, minToken, "amountTokenDesired", "amountTokenMin"); err != nil {
			return err
		}
	}

	// Validate token addresses
	if tokenA, ok := params["tokenA"]; ok {
		if tokenB, ok := params["tokenB"]; ok {
			if tokenA == tokenB {
				return fmt.Errorf("tokenA and tokenB cannot be the same")
			}
		}
	}

	// Common validations for liquidity transactions
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

// validateLiquidityRatio validates that minimum amount is reasonable compared to desired
func (v *UniswapV2Validator) validateLiquidityRatio(desired, minimum *big.Int, desiredName, minName string) error {
	if desired.Cmp(big.NewInt(0)) <= 0 {
		return fmt.Errorf("%s must be positive", desiredName)
	}
	if minimum.Cmp(big.NewInt(0)) <= 0 {
		return fmt.Errorf("%s must be positive", minName)
	}

	// Calculate ratio: minimum should be at least 90% of desired
	ratio := new(big.Int).Mul(minimum, big.NewInt(100))
	ratio.Div(ratio, desired)

	if ratio.Cmp(big.NewInt(90)) < 0 {
		return fmt.Errorf("%s too low compared to %s (less than 90%%), increase slippage protection", minName, desiredName)
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

	// Common validations for remove liquidity transactions
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
	// Check specific amount parameter names from Uniswap v2 ABI
	amountParams := []string{
		"amountIn", "amountOut", "amountOutMin", "amountInMax",
		"amountADesired", "amountBDesired", "amountTokenDesired",
		"amountAMin", "amountBMin", "amountTokenMin", "amountETHMin",
		"liquidity",
	}

	for _, paramName := range amountParams {
		if amount := v.getBigIntParam(params, paramName); amount != nil {
			if amount.Cmp(big.NewInt(0)) <= 0 {
				return fmt.Errorf("parameter %s must be positive, got: %s", paramName, amount.String())
			}

			// Additional check for extremely large values that could cause overflow
			maxValue := new(big.Int).Exp(big.NewInt(2), big.NewInt(255), nil) // 2^255
			if amount.Cmp(maxValue) >= 0 {
				return fmt.Errorf("parameter %s value too large, may cause overflow: %s", paramName, amount.String())
			}
		}
	}

	return nil
}

// validateAddresses validates all address parameters
func (v *UniswapV2Validator) validateAddresses(params map[string]interface{}) error {
	// Check specific address parameter names from Uniswap v2 ABI
	addressParams := []string{"to", "tokenA", "tokenB", "token"}

	for _, paramName := range addressParams {
		if paramValue, ok := params[paramName]; ok {
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
