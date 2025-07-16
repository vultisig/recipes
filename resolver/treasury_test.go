package resolver

import (
	"testing"

	"github.com/vultisig/recipes/types"
)

func TestTreasuryResolverSupports(t *testing.T) {
	resolver := NewTreasuryResolver()

	// Test that it supports the treasury magic constant
	if !resolver.Supports(types.MagicConstant_MAGIC_CONSTANT_VULTISIG_TREASURY) {
		t.Error("TreasuryResolver should support MAGIC_CONSTANT_VULTISIG_TREASURY")
	}

	// Test that it doesn't support other magic constants
	if resolver.Supports(types.MagicConstant_MAGIC_CONSTANT_UNSPECIFIED) {
		t.Error("TreasuryResolver should not support MAGIC_CONSTANT_UNSPECIFIED")
	}
}

func TestTreasuryResolverResolve(t *testing.T) {
	resolver := NewTreasuryResolver()

	tests := []struct {
		name     string
		chainID  string
		assetID  string
		expected string
		wantErr  bool
	}{
		{
			name:     "ethereum eth",
			chainID:  "ethereum",
			assetID:  "eth",
			expected: "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			wantErr:  false,
		},
		{
			name:     "ethereum usdc",
			chainID:  "ethereum",
			assetID:  "usdc",
			expected: "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			wantErr:  false,
		},
		{
			name:     "bitcoin btc",
			chainID:  "bitcoin",
			assetID:  "btc",
			expected: "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
			wantErr:  false,
		},
		{
			name:     "ethereum unknown asset (should use default)",
			chainID:  "ethereum",
			assetID:  "unknown",
			expected: "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			wantErr:  false,
		},
		{
			name:     "unsupported chain",
			chainID:  "unsupported",
			assetID:  "eth",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "chain without default",
			chainID:  "ethereum",
			assetID:  "unknown",
			expected: "0x742C4B65cc6cd34b45b3b99d50e3677b1e4b9b6e64",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolver.Resolve(types.MagicConstant_MAGIC_CONSTANT_VULTISIG_TREASURY, tt.chainID, tt.assetID)
			if (err != nil) != tt.wantErr {
				t.Errorf("TreasuryResolver.Resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("TreasuryResolver.Resolve() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTreasuryResolverResolveNonTreasuryConstant(t *testing.T) {
	resolver := NewTreasuryResolver()

	// Test with non-treasury magic constant
	_, err := resolver.Resolve(types.MagicConstant_MAGIC_CONSTANT_UNSPECIFIED, "ethereum", "eth")
	if err == nil {
		t.Error("TreasuryResolver.Resolve() should return error for non-treasury magic constant")
	}
}

func TestTreasuryResolverConfigStructure(t *testing.T) {
	resolver := NewTreasuryResolver().(*TreasuryResolver)

	// Test that the config has expected chains
	expectedChains := []string{"ethereum", "bitcoin", "arbitrum"}
	for _, chain := range expectedChains {
		if _, exists := resolver.treasuryConfig[chain]; !exists {
			t.Errorf("TreasuryResolver config missing chain: %s", chain)
		}
	}

	// Test that ethereum chain has expected assets
	ethereumConfig := resolver.treasuryConfig["ethereum"]
	expectedAssets := []string{"eth", "usdc", "dai", "weth", "default"}
	for _, asset := range expectedAssets {
		if _, exists := ethereumConfig[asset]; !exists {
			t.Errorf("TreasuryResolver ethereum config missing asset: %s", asset)
		}
	}
}
