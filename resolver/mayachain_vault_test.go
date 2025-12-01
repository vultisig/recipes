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

func TestMayaChainVaultResolver_Supports(t *testing.T) {
	resolver := NewMayaChainVaultResolver()

	tests := []struct {
		name     string
		constant types.MagicConstant
		expected bool
	}{
		{
			name:     "supports MAYACHAIN_VAULT",
			constant: types.MagicConstant_MAYACHAIN_VAULT,
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

func TestMayaChainVaultResolver_Resolve_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	resolver := NewMayaChainVaultResolver()

	tests := []struct {
		name    string
		chainID string
		assetID string
		wantErr bool
	}{
		{
			name:    "resolve ZEC vault address",
			chainID: "zcash",
			assetID: "zec",
			wantErr: false,
		},
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
			name:    "resolve DASH vault address",
			chainID: "dash",
			assetID: "dash",
			wantErr: false,
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
				types.MagicConstant_MAYACHAIN_VAULT,
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

			// MayaChain vault addresses should not have memos
			if memo != "" {
				t.Errorf("Expected empty memo, got: %s", memo)
			}

			t.Logf("Successfully resolved %s vault address: %s", tt.chainID, address)
		})
	}
}

func TestMayaChainVaultResolver_UnsupportedConstant(t *testing.T) {
	resolver := NewMayaChainVaultResolver()

	_, _, err := resolver.Resolve(
		types.MagicConstant_VULTISIG_TREASURY,
		"zcash",
		"zec",
	)

	if err == nil {
		t.Error("Expected error for unsupported magic constant")
	}

	expectedError := "MayaChainVaultResolver does not support type"
	if err != nil && !strings.HasPrefix(err.Error(), expectedError) {
		t.Errorf("Expected error message to start with '%s', got: %s", expectedError, err.Error())
	}
}

func TestMayaChainVaultResolver_UnsupportedChainError(t *testing.T) {
	resolver := NewMayaChainVaultResolver()

	// Test that unsupported chains return immediate error without API call
	_, _, err := resolver.Resolve(
		types.MagicConstant_MAYACHAIN_VAULT,
		"polygon", // Not in our supported chainMap
		"matic",
	)

	if err == nil {
		t.Error("Expected error for unsupported chain polygon")
		return
	}

	expectedError := "chain polygon not supported by MayaChain"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("Expected error message to contain '%s', got: %s", expectedError, err.Error())
	}
}

func TestMayaChainVaultResolver_APIConsistency(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping API consistency test in short mode")
	}

	// Query MayaChain API directly
	apiAddresses, err := queryMayaChainAPIDirect()
	if err != nil {
		t.Fatalf("Failed to query MayaChain API directly: %v", err)
	}

	// Create our resolver
	resolver := NewMayaChainVaultResolver()

	// Test each supported chain
	supportedChains := []struct {
		chainID       string
		mayachainName string
	}{
		{"zcash", "ZEC"},
		{"bitcoin", "BTC"},
		{"ethereum", "ETH"},
		{"dash", "DASH"},
	}

	for _, chain := range supportedChains {
		t.Run(chain.chainID, func(t *testing.T) {
			// Get address from our resolver
			resolverAddress, _, err := resolver.Resolve(
				types.MagicConstant_MAYACHAIN_VAULT,
				chain.chainID,
				"asset", // assetID doesn't matter for vault resolution
			)
			if err != nil {
				t.Fatalf("Resolver failed for %s: %v", chain.chainID, err)
			}

			// Find the corresponding address from direct API call
			var apiAddress string
			for _, addr := range apiAddresses {
				if strings.ToUpper(addr.Chain) == chain.mayachainName {
					// For chains with router, expect router address, otherwise vault address
					if addr.Router != "" {
						apiAddress = addr.Router
					} else {
						apiAddress = addr.Address
					}
					break
				}
			}

			if apiAddress == "" {
				t.Fatalf("No API address found for MayaChain chain %s", chain.mayachainName)
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

// Query the MayaChain API directly for comparison
func queryMayaChainAPIDirect() ([]MayaInboundAddress, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get("https://mayanode.mayachain.info/mayachain/inbound_addresses")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("mayanode inbound_addresses status %d: %s", resp.StatusCode, strings.TrimSpace(string(b)))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var addresses []MayaInboundAddress
	if err := json.Unmarshal(body, &addresses); err != nil {
		return nil, err
	}

	return addresses, nil
}

