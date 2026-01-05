package resolver

import (
	"strings"
	"testing"

	"github.com/vultisig/recipes/types"
)

func TestTHORChainRouterResolver_Supports(t *testing.T) {
	resolver := NewTHORChainRouterResolver()

	tests := []struct {
		name     string
		constant types.MagicConstant
		expected bool
	}{
		{
			name:     "supports THORCHAIN_ROUTER",
			constant: types.MagicConstant_THORCHAIN_ROUTER,
			expected: true,
		},
		{
			name:     "does not support THORCHAIN_VAULT",
			constant: types.MagicConstant_THORCHAIN_VAULT,
			expected: false,
		},
		{
			name:     "does not support VULTISIG_TREASURY",
			constant: types.MagicConstant_VULTISIG_TREASURY,
			expected: false,
		},
		{
			name:     "does not support UNSPECIFIED",
			constant: types.MagicConstant_UNSPECIFIED,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := resolver.Supports(tt.constant)
			if result != tt.expected {
				t.Errorf("Supports(%v) = %v, expected %v", tt.constant, result, tt.expected)
			}
		})
	}
}

func TestTHORChainRouterResolver_Resolve_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	resolver := NewTHORChainRouterResolver()

	tests := []struct {
		name    string
		chainID string
		assetID string
		wantErr bool
	}{
		{
			name:    "resolve ETH router address",
			chainID: "ethereum",
			assetID: "eth",
			wantErr: false,
		},
		{
			name:    "resolve AVAX router address",
			chainID: "avalanche",
			assetID: "avax",
			wantErr: false,
		},
		{
			name:    "resolve BSC router address",
			chainID: "bsc",
			assetID: "bnb",
			wantErr: false,
		},
		{
			name:    "resolve Base router address",
			chainID: "base",
			assetID: "eth",
			wantErr: false,
		},
		{
			name:    "error on non-EVM chain BTC",
			chainID: "bitcoin",
			assetID: "btc",
			wantErr: true,
		},
		{
			name:    "error on non-EVM chain XRP",
			chainID: "ripple",
			assetID: "xrp",
			wantErr: true,
		},
		{
			name:    "unsupported chain should error",
			chainID: "unsupported_chain",
			assetID: "token",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			address, memo, err := resolver.Resolve(
				types.MagicConstant_THORCHAIN_ROUTER,
				tt.chainID,
				tt.assetID,
			)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error for chainID %s, but got none", tt.chainID)
				}
				return
			}

			if err != nil {
				if strings.Contains(err.Error(), "is currently halted") {
					t.Skipf("Chain %s is currently halted on THORChain, skipping test", tt.chainID)
					return
				}
				t.Errorf("Unexpected error for chainID %s: %v", tt.chainID, err)
				return
			}

			if address == "" {
				t.Errorf("Expected non-empty router address for chainID %s", tt.chainID)
			}

			if memo != "" {
				t.Errorf("Expected empty memo, got: %s", memo)
			}

			t.Logf("Successfully resolved %s router address: %s", tt.chainID, address)
		})
	}
}

func TestTHORChainRouterResolver_UnsupportedConstant(t *testing.T) {
	resolver := NewTHORChainRouterResolver()

	_, _, err := resolver.Resolve(
		types.MagicConstant_THORCHAIN_VAULT,
		"ethereum",
		"eth",
	)

	if err == nil {
		t.Error("Expected error for unsupported magic constant")
	}

	expectedError := "THORChainRouterResolver does not support type"
	if err != nil && !strings.HasPrefix(err.Error(), expectedError) {
		t.Errorf("Expected error message to start with '%s', got: %s", expectedError, err.Error())
	}
}

func TestTHORChainRouterResolver_NonEvmChainError(t *testing.T) {
	resolver := NewTHORChainRouterResolver()

	nonEvmChains := []string{"bitcoin", "ripple", "litecoin", "dogecoin", "bitcoincash"}

	for _, chainID := range nonEvmChains {
		t.Run(chainID, func(t *testing.T) {
			_, _, err := resolver.Resolve(
				types.MagicConstant_THORCHAIN_ROUTER,
				chainID,
				"asset",
			)

			if err == nil {
				t.Errorf("Expected error for non-EVM chain %s", chainID)
				return
			}

			expectedError := "is only available for EVM chains"
			if !strings.Contains(err.Error(), expectedError) {
				t.Errorf("Expected error message to contain '%s', got: %s", expectedError, err.Error())
			}
		})
	}
}

func TestTHORChainRouterResolver_APIConsistency(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping API consistency test in short mode")
	}

	apiAddresses, err := queryTHORChainAPIDirect()
	if err != nil {
		t.Fatalf("Failed to query THORChain API directly: %v", err)
	}

	resolver := NewTHORChainRouterResolver()

	supportedEvmChains := []struct {
		chainID       string
		thorchainName string
	}{
		{"ethereum", "ETH"},
		{"avalanche", "AVAX"},
		{"bsc", "BSC"},
		{"base", "BASE"},
	}

	for _, chain := range supportedEvmChains {
		t.Run(chain.chainID, func(t *testing.T) {
			resolverAddress, _, err := resolver.Resolve(
				types.MagicConstant_THORCHAIN_ROUTER,
				chain.chainID,
				"asset",
			)
			if err != nil {
				if strings.Contains(err.Error(), "is currently halted") {
					t.Skipf("Chain %s is currently halted on THORChain, skipping test", chain.chainID)
					return
				}
				t.Fatalf("Resolver failed for %s: %v", chain.chainID, err)
			}

			var apiRouterAddress string
			for _, addr := range apiAddresses {
				if strings.ToUpper(addr.Chain) == chain.thorchainName {
					apiRouterAddress = addr.Router
					break
				}
			}

			if apiRouterAddress == "" {
				t.Fatalf("No router address found in API response for %s", chain.chainID)
			}

			if resolverAddress != apiRouterAddress {
				t.Errorf("Router address mismatch for %s: resolver=%s, api=%s",
					chain.chainID, resolverAddress, apiRouterAddress)
			}

			t.Logf("✓ %s router address matches: %s", chain.chainID, resolverAddress)
		})
	}
}
