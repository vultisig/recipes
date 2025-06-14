package ethereum

import (
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func TestUniswapV2Validator_ValidateTransaction(t *testing.T) {
	validator := NewUniswapV2Validator()

	tests := []struct {
		name          string
		functionName  string
		params        map[string]interface{}
		expectError   bool
		errorContains string
	}{
		{
			name:         "Valid swapExactETHForTokens",
			functionName: "swapExactETHForTokens",
			params: map[string]interface{}{
				"amountOutMin": big.NewInt(1000000000000000000), // 1 ETH worth
				"path": []interface{}{
					"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // WETH
					"0xA0b86a33E6BA3b1b3b3b3b3b3b3b3b3b3b3b3b3b", // Sample token
				},
				"to":       "0x742d35cc6671FbF82f0a8C2A9C8fBc9b8b8b8b8b",
				"deadline": big.NewInt(time.Now().Unix() + 1800), // 30 minutes from now
			},
			expectError: false,
		},
		{
			name:         "Invalid swapExactETHForTokens - zero amountOutMin",
			functionName: "swapExactETHForTokens",
			params: map[string]interface{}{
				"amountOutMin": big.NewInt(0), // Invalid: no slippage protection
				"path": []interface{}{
					"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					"0xA0b86a33E6BA3b1b3b3b3b3b3b3b3b3b3b3b3b3b",
				},
				"to":       "0x742d35cc6671FbF82f0a8C2A9C8fBc9b8b8b8b8b",
				"deadline": big.NewInt(time.Now().Unix() + 1800),
			},
			expectError:   true,
			errorContains: "amountOutMin must be positive",
		},
		{
			name:         "Invalid swapExactETHForTokens - zero address in path",
			functionName: "swapExactETHForTokens",
			params: map[string]interface{}{
				"amountOutMin": big.NewInt(1000000000000000000),
				"path": []interface{}{
					"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					"0x0000000000000000000000000000000000000000", // Invalid: zero address
				},
				"to":       "0x742d35cc6671FbF82f0a8C2A9C8fBc9b8b8b8b8b",
				"deadline": big.NewInt(time.Now().Unix() + 1800),
			},
			expectError:   true,
			errorContains: "zero address not allowed in swap path",
		},
		{
			name:         "Invalid swapExactETHForTokens - past deadline",
			functionName: "swapExactETHForTokens",
			params: map[string]interface{}{
				"amountOutMin": big.NewInt(1000000000000000000),
				"path": []interface{}{
					"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					"0xA0b86a33E6BA3b1b3b3b3b3b3b3b3b3b3b3b3b3b",
				},
				"to":       "0x742d35cc6671FbF82f0a8C2A9C8fBc9b8b8b8b8b",
				"deadline": big.NewInt(1600000000), // Past deadline
			},
			expectError:   true,
			errorContains: "deadline must be in the future",
		},
		{
			name:         "Invalid swapExactETHForTokens - zero address recipient",
			functionName: "swapExactETHForTokens",
			params: map[string]interface{}{
				"amountOutMin": big.NewInt(1000000000000000000),
				"path": []interface{}{
					"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					"0xA0b86a33E6BA3b1b3b3b3b3b3b3b3b3b3b3b3b3b",
				},
				"to":       "0x0000000000000000000000000000000000000000", // Invalid: zero address
				"deadline": big.NewInt(time.Now().Unix() + 1800),
			},
			expectError:   true,
			errorContains: "parameter to cannot be zero address",
		},
		{
			name:         "Valid addLiquidity",
			functionName: "addLiquidity",
			params: map[string]interface{}{
				"tokenA":         "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
				"tokenB":         "0xA0b86a33E6BA3b1b3b3b3b3b3b3b3b3b3b3b3b3b",
				"amountADesired": big.NewInt(1000000000000000000), // 1 ETH
				"amountBDesired": big.NewInt(1000000000000000000), // 1 token
				"amountAMin":     big.NewInt(950000000000000000),  // 95% of desired (good slippage)
				"amountBMin":     big.NewInt(950000000000000000),  // 95% of desired
				"to":             "0x742d35cc6671FbF82f0a8C2A9C8fBc9b8b8b8b8b",
				"deadline":       big.NewInt(time.Now().Unix() + 1800),
			},
			expectError: false,
		},
		{
			name:         "Invalid addLiquidity - same tokenA and tokenB",
			functionName: "addLiquidity",
			params: map[string]interface{}{
				"tokenA":         "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
				"tokenB":         "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // Same as tokenA
				"amountADesired": big.NewInt(1000000000000000000),
				"amountBDesired": big.NewInt(1000000000000000000),
				"amountAMin":     big.NewInt(950000000000000000),
				"amountBMin":     big.NewInt(950000000000000000),
				"to":             "0x742d35cc6671FbF82f0a8C2A9C8fBc9b8b8b8b8b",
				"deadline":       big.NewInt(time.Now().Unix() + 1800),
			},
			expectError:   true,
			errorContains: "tokenA and tokenB cannot be the same",
		},
		{
			name:         "Invalid addLiquidity - poor slippage protection",
			functionName: "addLiquidity",
			params: map[string]interface{}{
				"tokenA":         "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
				"tokenB":         "0xA0b86a33E6BA3b1b3b3b3b3b3b3b3b3b3b3b3b3b",
				"amountADesired": big.NewInt(1000000000000000000),
				"amountBDesired": big.NewInt(1000000000000000000),
				"amountAMin":     big.NewInt(100000000000000000), // Only 10% slippage protection (too low)
				"amountBMin":     big.NewInt(950000000000000000),
				"to":             "0x742d35cc6671FbF82f0a8C2A9C8fBc9b8b8b8b8b",
				"deadline":       big.NewInt(time.Now().Unix() + 1800),
			},
			expectError:   true,
			errorContains: "amountAMin too low compared to amountADesired",
		},
		{
			name:         "Valid swapETHForExactTokens (exact output)",
			functionName: "swapETHForExactTokens",
			params: map[string]interface{}{
				"amountOut":   big.NewInt(1000000000000000000), // Want exactly 1 token
				"amountInMax": big.NewInt(2000000000000000000), // Willing to spend up to 2 ETH
				"path": []interface{}{
					"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					"0xA0b86a33E6BA3b1b3b3b3b3b3b3b3b3b3b3b3b3b",
				},
				"to":       "0x742d35cc6671FbF82f0a8C2A9C8fBc9b8b8b8b8b",
				"deadline": big.NewInt(time.Now().Unix() + 1800),
			},
			expectError: false,
		},
		{
			name:         "Invalid swapETHForExactTokens - zero amountInMax",
			functionName: "swapETHForExactTokens",
			params: map[string]interface{}{
				"amountOut":   big.NewInt(1000000000000000000),
				"amountInMax": big.NewInt(0), // Invalid: no slippage protection
				"path": []interface{}{
					"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					"0xA0b86a33E6BA3b1b3b3b3b3b3b3b3b3b3b3b3b3b",
				},
				"to":       "0x742d35cc6671FbF82f0a8C2A9C8fBc9b8b8b8b8b",
				"deadline": big.NewInt(time.Now().Unix() + 1800),
			},
			expectError:   true,
			errorContains: "amountInMax must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateTransaction(tt.functionName, tt.params)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tt.errorContains != "" && !contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain '%s', but got: %s", tt.errorContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %s", err.Error())
				}
			}
		})
	}
}

func TestUniswapV2Validator_GetProtocolID(t *testing.T) {
	validator := NewUniswapV2Validator()
	if validator.GetProtocolID() != "uniswapv2_router" {
		t.Errorf("Expected protocol ID 'uniswapv2_router', got '%s'", validator.GetProtocolID())
	}
}

func TestUniswapV2Validator_ValidateAddresses(t *testing.T) {
	validator := NewUniswapV2Validator()

	tests := []struct {
		name          string
		params        map[string]interface{}
		expectError   bool
		errorContains string
	}{
		{
			name: "Valid string address",
			params: map[string]interface{}{
				"to": "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
			},
			expectError: false,
		},
		{
			name: "Valid common.Address",
			params: map[string]interface{}{
				"to": common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"),
			},
			expectError: false,
		},
		{
			name: "Zero address string (lowercase)",
			params: map[string]interface{}{
				"to": "0x0000000000000000000000000000000000000000",
			},
			expectError:   true,
			errorContains: "cannot be zero address",
		},
		{
			name: "Zero address string (uppercase)",
			params: map[string]interface{}{
				"to": "0X0000000000000000000000000000000000000000",
			},
			expectError:   true,
			errorContains: "cannot be zero address",
		},
		{
			name: "Zero address common.Address",
			params: map[string]interface{}{
				"to": common.Address{},
			},
			expectError:   true,
			errorContains: "cannot be zero address",
		},
		{
			name: "Invalid address string",
			params: map[string]interface{}{
				"to": "invalid-address",
			},
			expectError:   true,
			errorContains: "must be a valid Ethereum address",
		},
		{
			name: "Invalid address type",
			params: map[string]interface{}{
				"to": 12345,
			},
			expectError:   true,
			errorContains: "must be a string or common.Address",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.validateAddresses(tt.params)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tt.errorContains != "" && !contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain '%s', but got: %s", tt.errorContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %s", err.Error())
				}
			}
		})
	}
}

func TestUniswapV2Validator_ValidateSwapPath(t *testing.T) {
	validator := NewUniswapV2Validator()

	tests := []struct {
		name          string
		params        map[string]interface{}
		expectError   bool
		errorContains string
	}{
		{
			name: "Valid []interface{} path",
			params: map[string]interface{}{
				"path": []interface{}{
					"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
				},
				"amountOutMin": big.NewInt(1000000000000000000),
				"to":           "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
				"deadline":     big.NewInt(time.Now().Unix() + 1800),
			},
			expectError: false,
		},
		{
			name: "Valid []common.Address path",
			params: map[string]interface{}{
				"path": []common.Address{
					common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
					common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"),
				},
				"amountOutMin": big.NewInt(1000000000000000000),
				"to":           "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
				"deadline":     big.NewInt(time.Now().Unix() + 1800),
			},
			expectError: false,
		},
		{
			name: "Valid []string path",
			params: map[string]interface{}{
				"path": []string{
					"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
				},
				"amountOutMin": big.NewInt(1000000000000000000),
				"to":           "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
				"deadline":     big.NewInt(time.Now().Unix() + 1800),
			},
			expectError: false,
		},
		{
			name: "Invalid path type",
			params: map[string]interface{}{
				"path":         12345,
				"amountOutMin": big.NewInt(1000000000000000000),
				"to":           "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
				"deadline":     big.NewInt(time.Now().Unix() + 1800),
			},
			expectError:   true,
			errorContains: "invalid path type: expected []interface{}, []common.Address, or []string",
		},
		{
			name: "Path too short ([]string)",
			params: map[string]interface{}{
				"path": []string{
					"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
				},
				"amountOutMin": big.NewInt(1000000000000000000),
				"to":           "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
				"deadline":     big.NewInt(time.Now().Unix() + 1800),
			},
			expectError:   true,
			errorContains: "swap path must contain at least 2 addresses",
		},
		{
			name: "Path too long ([]string)",
			params: map[string]interface{}{
				"path": []string{
					"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
					"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
					"0xA0b86a33E6BA3b1b3b3b3b3b3b3b3b3b3b3b3b3b",
					"0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
					"0x742d35cc6671FbF82f0a8C2A9C8fBc9b8b8b8b8b",
				},
				"amountOutMin": big.NewInt(1000000000000000000),
				"to":           "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
				"deadline":     big.NewInt(time.Now().Unix() + 1800),
			},
			expectError:   true,
			errorContains: "swap path too long, maximum 4 tokens supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.validateSwapTransaction("swapExactETHForTokens", tt.params)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				if tt.errorContains != "" && !contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain '%s', but got: %s", tt.errorContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %s", err.Error())
				}
			}
		})
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
