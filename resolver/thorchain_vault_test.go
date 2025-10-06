package resolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/vultisig/recipes/types"
)

func TestTHORChainVaultResolver_Supports(t *testing.T) {
	resolver := NewTHORChainVaultResolver()

	tests := []struct {
		name     string
		constant types.MagicConstant
		expected bool
	}{
		{
			name:     "supports THORCHAIN_VAULT",
			constant: types.MagicConstant_THORCHAIN_VAULT,
			expected: true,
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

func TestTHORChainVaultResolver_Resolve_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	resolver := NewTHORChainVaultResolver()

	tests := []struct {
		name    string
		chainID string
		assetID string
		wantErr bool
	}{
		{
			name:    "resolve BTC vault address",
			chainID: "bitcoin",
			assetID: "btc",
			wantErr: false,
		},
		{
			name:    "resolve ETH vault address",
			chainID: "ethereum",
			assetID: "eth",
			wantErr: false,
		},
		{
			name:    "resolve Base vault address",
			chainID: "base",
			assetID: "eth",
			wantErr: false,
		},
		{
			name:    "resolve XRP vault address",
			chainID: "ripple",
			assetID: "xrp",
			wantErr: false,
		},
		{
			name:    "resolve using uppercase chain",
			chainID: "bitcoin",
			assetID: "btc",
			wantErr: false,
		},
		{
			name:    "unsupported chain BSC should error",
			chainID: "bscchain",
			assetID: "bnb",
			wantErr: true,
		},
		{
			name:    "unsupported chain Arbitrum should error",
			chainID: "arbitrum",
			assetID: "eth",
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
				types.MagicConstant_THORCHAIN_VAULT,
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
				t.Errorf("Unexpected error for chainID %s: %v", tt.chainID, err)
				return
			}

			if address == "" {
				t.Errorf("Expected non-empty address for chainID %s", tt.chainID)
			}

			// THORChain vault addresses should not have memos
			if memo != "" {
				t.Errorf("Expected empty memo, got: %s", memo)
			}

			t.Logf("Successfully resolved %s vault address: %s", tt.chainID, address)
		})
	}
}

func TestTHORChainVaultResolver_UnsupportedConstant(t *testing.T) {
	resolver := NewTHORChainVaultResolver()

	_, _, err := resolver.Resolve(
		types.MagicConstant_VULTISIG_TREASURY,
		"bitcoin",
		"btc",
	)

	if err == nil {
		t.Error("Expected error for unsupported magic constant")
	}

	expectedError := "THORChainVaultResolver does not support type"
	if err != nil && !strings.HasPrefix(err.Error(), expectedError) {
		t.Errorf("Expected error message to start with '%s', got: %s", expectedError, err.Error())
	}
}

func TestTHORChainVaultResolver_UnsupportedChainError(t *testing.T) {
	resolver := NewTHORChainVaultResolver()

	// Test that unsupported chains return immediate error without API call
	_, _, err := resolver.Resolve(
		types.MagicConstant_THORCHAIN_VAULT,
		"polygon", // Not in our supported chainMap
		"matic",
	)

	if err == nil {
		t.Error("Expected error for unsupported chain polygon")
		return
	}

	expectedError := "chain Polygon not supported by ThorChain"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Expected error message to contain '%s', got: %s", expectedError, err.Error())
	}
}

func TestTHORChainVaultResolver_APIConsistency(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping API consistency test in short mode")
	}

	// Query THORChain API directly
	apiAddresses, err := queryTHORChainAPIDirect()
	if err != nil {
		t.Fatalf("Failed to query THORChain API directly: %v", err)
	}

	// Create our resolver
	resolver := NewTHORChainVaultResolver()

	// Test each supported chain
	supportedChains := []struct {
		chainID       string
		thorchainName string
	}{
		{"ethereum", "ETH"},
		{"bitcoin", "BTC"},
		{"base", "BASE"},
		{"ripple", "XRP"},
	}

	for _, chain := range supportedChains {
		t.Run(chain.chainID, func(t *testing.T) {
			// Get address from our resolver
			resolverAddress, _, err := resolver.Resolve(
				types.MagicConstant_THORCHAIN_VAULT,
				chain.chainID,
				"asset", // assetID doesn't matter for vault resolution
			)
			if err != nil {
				t.Fatalf("Resolver failed for %s: %v", chain.chainID, err)
			}

			// Find the corresponding address from direct API call
			var apiAddress string
			for _, addr := range apiAddresses {
				if strings.ToUpper(addr.Chain) == chain.thorchainName {
					// For EVM chains, expect router address if available, otherwise vault address
					if (chain.chainID == "ethereum" || chain.chainID == "base") && addr.Router != "" {
						apiAddress = addr.Router
					} else {
						apiAddress = addr.Address
					}
					break
				}
			}

			if apiAddress == "" {
				t.Fatalf("No API address found for THORChain chain %s", chain.thorchainName)
			}

			// Compare addresses
			if resolverAddress != apiAddress {
				t.Errorf("Address mismatch for %s:\n  Resolver: %s\n  API:      %s",
					chain.chainID, resolverAddress, apiAddress)
			} else {
				t.Logf("✓ %s addresses match: %s", chain.chainID, resolverAddress)
			}
		})
	}
}

// Query the THORChain API directly for comparison
func queryTHORChainAPIDirect() ([]InboundAddress, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get("https://thornode.ninerealms.com/thorchain/inbound_addresses")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("thornode inbound_addresses status %d: %s", resp.StatusCode, strings.TrimSpace(string(b)))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var addresses []InboundAddress
	if err := json.Unmarshal(body, &addresses); err != nil {
		return nil, err
	}

	return addresses, nil
}
