package resolver

import (
	"testing"

	"github.com/vultisig/recipes/types"
)

func TestTreasuryResolverSupports(t *testing.T) {
	resolver := NewDefaultTreasuryResolver()

	// Test that it supports the treasury magic constant
	if !resolver.Supports(types.MagicConstant_VULTISIG_TREASURY) {
		t.Error("TreasuryResolver should support VULTISIG_TREASURY")
	}

	// Test that it doesn't support other magic constants
	if resolver.Supports(types.MagicConstant_UNSPECIFIED) {
		t.Error("TreasuryResolver should not support UNSPECIFIED")
	}
}

func TestTreasuryResolverResolve(t *testing.T) {
	resolver := NewDefaultTreasuryResolver()

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
			expected: "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
			wantErr:  false,
		},
		{
			name:     "ethereum usdc",
			chainID:  "ethereum",
			assetID:  "usdc",
			expected: "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
			wantErr:  false,
		},
		{
			name:     "bitcoin btc",
			chainID:  "bitcoin",
			assetID:  "btc",
			expected: "bc1qelza2cr7w92dmzgkmhdn0a4vcqpe9rfpfknr6a",
			wantErr:  false,
		},
		{
			name:     "ethereum unknown asset (should use default)",
			chainID:  "ethereum",
			assetID:  "unknown",
			expected: "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
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
			expected: "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := resolver.Resolve(types.MagicConstant_VULTISIG_TREASURY, tt.chainID, tt.assetID)
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
	resolver := NewDefaultTreasuryResolver()

	// Test with non-treasury magic constant
	_, _, err := resolver.Resolve(types.MagicConstant_UNSPECIFIED, "ethereum", "eth")
	if err == nil {
		t.Error("TreasuryResolver.Resolve() should return error for non-treasury magic constant")
	}
}

func TestTreasuryResolverConfigStructure(t *testing.T) {
	resolver := NewDefaultTreasuryResolver().(*TreasuryResolver)

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
