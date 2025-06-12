package ethereum

import (
	"math/big"
	"testing"
	"time"
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

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || (len(s) > len(substr) && contains(s[1:], substr) || s[:len(substr)] == substr))
}
