package resolver

import (
	"testing"

	"github.com/vultisig/recipes/types"
)

func TestNewMagicResolver(t *testing.T) {
	resolver := NewMagicResolver()

	if resolver == nil {
		t.Fatal("NewMagicResolver() returned nil")
	}

	if resolver.treasuryConfig == nil {
		t.Fatal("treasuryConfig is nil")
	}

	// Check that some expected chains are configured
	expectedChains := []string{"ethereum", "bitcoin"}
	for _, chain := range expectedChains {
		if _, exists := resolver.treasuryConfig[chain]; !exists {
			t.Errorf("Expected chain %s not found in treasuryConfig", chain)
		}
	}
}

func TestMagicResolver_Resolve_VultisigTreasury_Success(t *testing.T) {
	resolver := NewMagicResolver()

	testCases := []struct {
		name     string
		chainID  string
		assetID  string
		expected string
	}{
		{
			name:     "ethereum_eth",
			chainID:  "ethereum",
			assetID:  "eth",
			expected: "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
		},
		{
			name:     "ethereum_usdc",
			chainID:  "ethereum",
			assetID:  "usdc",
			expected: "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
		},
		{
			name:     "ethereum_dai",
			chainID:  "ethereum",
			assetID:  "dai",
			expected: "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
		},
		{
			name:     "ethereum_weth",
			chainID:  "ethereum",
			assetID:  "weth",
			expected: "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
		},
		{
			name:     "bitcoin_btc",
			chainID:  "bitcoin",
			assetID:  "btc",
			expected: "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := resolver.Resolve(types.MagicConstant_MAGIC_CONSTANT_VULTISIG_TREASURY, tc.chainID, tc.assetID)

			if err != nil {
				t.Fatalf("Resolve() returned error: %v", err)
			}

			if result != tc.expected {
				t.Errorf("Resolve() = %v, expected %v", result, tc.expected)
			}
		})
	}
}

func TestMagicResolver_Resolve_FallbackToDefault(t *testing.T) {
	resolver := NewMagicResolver()

	testCases := []struct {
		name     string
		chainID  string
		assetID  string
		expected string
	}{
		{
			name:     "ethereum_unknown_asset_fallback",
			chainID:  "ethereum",
			assetID:  "unknown",
			expected: "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64", // Should fallback to default
		},
		{
			name:     "ethereum_uniswap_fallback",
			chainID:  "ethereum",
			assetID:  "uniswapv2",
			expected: "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64", // Should fallback to default
		},
		{
			name:     "bitcoin_unknown_asset_fallback",
			chainID:  "bitcoin",
			assetID:  "unknown",
			expected: "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4", // Should fallback to default
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := resolver.Resolve(types.MagicConstant_MAGIC_CONSTANT_VULTISIG_TREASURY, tc.chainID, tc.assetID)

			if err != nil {
				t.Fatalf("Expected fallback to default, but got error: %v", err)
			}

			if result != tc.expected {
				t.Errorf("Resolve() = %v, expected fallback to %v", result, tc.expected)
			}
		})
	}
}

func TestMagicResolver_Resolve_UnsupportedChain(t *testing.T) {
	resolver := NewMagicResolver()

	testCases := []struct {
		name    string
		chainID string
		assetID string
	}{
		{
			name:    "unknown_chain",
			chainID: "unknown",
			assetID: "eth",
		},
		{
			name:    "polygon_chain",
			chainID: "polygon",
			assetID: "matic",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := resolver.Resolve(types.MagicConstant_MAGIC_CONSTANT_VULTISIG_TREASURY, tc.chainID, tc.assetID)

			if err == nil {
				t.Fatalf("Expected error for unsupported chain %s, but got result: %v", tc.chainID, result)
			}

			expectedError := "no treasury address configured for chain " + tc.chainID
			if err.Error() != expectedError {
				t.Errorf("Expected error message %q, got %q", expectedError, err.Error())
			}

			if result != "" {
				t.Errorf("Expected empty result for error case, got %q", result)
			}
		})
	}
}

func TestMagicResolver_Resolve_UnsupportedAssetNoDefault(t *testing.T) {
	// Create a resolver with no default configured for testing
	resolver := &MagicResolver{
		treasuryConfig: map[string]map[string]string{
			"testchain": {
				"specific": "0xspecificaddress",
				// No "default" key
			},
		},
	}

	result, err := resolver.Resolve(types.MagicConstant_MAGIC_CONSTANT_VULTISIG_TREASURY, "testchain", "unknown")

	if err == nil {
		t.Fatalf("Expected error when no default exists, but got result: %v", result)
	}

	expectedError := "no treasury address configured for asset unknown on chain testchain"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, err.Error())
	}

	if result != "" {
		t.Errorf("Expected empty result for error case, got %q", result)
	}
}

func TestMagicResolver_Resolve_UnsupportedMagicConstant(t *testing.T) {
	resolver := NewMagicResolver()

	// Test with an unspecified magic constant
	result, err := resolver.Resolve(types.MagicConstant_MAGIC_CONSTANT_UNSPECIFIED, "ethereum", "eth")

	if err == nil {
		t.Fatalf("Expected error for unsupported magic constant, but got result: %v", result)
	}

	expectedError := "unsupported magic constant: MAGIC_CONSTANT_UNSPECIFIED"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, err.Error())
	}

	if result != "" {
		t.Errorf("Expected empty result for error case, got %q", result)
	}
}

func TestMagicResolver_resolveTreasury_Success(t *testing.T) {
	resolver := NewMagicResolver()

	// Test the private method through the public interface
	result, err := resolver.resolveTreasury("ethereum", "usdc")

	if err != nil {
		t.Fatalf("resolveTreasury() returned error: %v", err)
	}

	expected := "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64"
	if result != expected {
		t.Errorf("resolveTreasury() = %v, expected %v", result, expected)
	}
}

func TestMagicResolver_resolveTreasury_FallbackToDefault(t *testing.T) {
	resolver := NewMagicResolver()

	// Test fallback behavior
	result, err := resolver.resolveTreasury("ethereum", "nonexistent")

	if err != nil {
		t.Fatalf("Expected fallback to default, but got error: %v", err)
	}

	expected := "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64" // Should be default address
	if result != expected {
		t.Errorf("resolveTreasury() = %v, expected fallback to %v", result, expected)
	}
}

func TestMagicResolver_resolveTreasury_ChainNotFound(t *testing.T) {
	resolver := NewMagicResolver()

	result, err := resolver.resolveTreasury("nonexistent", "eth")

	if err == nil {
		t.Fatalf("Expected error for nonexistent chain, but got result: %v", result)
	}

	expectedError := "no treasury address configured for chain nonexistent"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, err.Error())
	}
}

func TestMagicResolver_resolveTreasury_AssetAndNoDefaultNotFound(t *testing.T) {
	// Create a resolver with no default for testing edge case
	resolver := &MagicResolver{
		treasuryConfig: map[string]map[string]string{
			"testchain": {
				"btc": "bc1qtest",
				// No "default" key
			},
		},
	}

	result, err := resolver.resolveTreasury("testchain", "nonexistent")

	if err == nil {
		t.Fatalf("Expected error for nonexistent asset with no default, but got result: %v", result)
	}

	expectedError := "no treasury address configured for asset nonexistent on chain testchain"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, err.Error())
	}
}
