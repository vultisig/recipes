package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vultisig/vultisig-go/common"
)

func TestChainEngineRegistry(t *testing.T) {
	registry, _ := NewChainEngineRegistry()

	tests := []struct {
		name        string
		chain       common.Chain
		shouldFind  bool
		description string
	}{
		{
			name:        "EVM chains",
			chain:       common.Ethereum,
			shouldFind:  true,
			description: "Ethereum should be supported by EVM engine",
		},
		{
			name:        "EVM chains - BSC",
			chain:       common.BscChain,
			shouldFind:  true,
			description: "BSC should be supported by EVM engine",
		},
		{
			name:        "Bitcoin",
			chain:       common.Bitcoin,
			shouldFind:  true,
			description: "Bitcoin should be supported by BTC engine",
		},
		{
			name:        "XRP",
			chain:       common.XRP,
			shouldFind:  true,
			description: "XRP should be supported by XRPL engine",
		},
		{
			name:        "Solana",
			chain:       common.Solana,
			shouldFind:  true,
			description: "Solana should be supported by Solana engine",
		},
		{
			name:        "Thorchain",
			chain:       common.THORChain,
			shouldFind:  true,
			description: "THORChain should be supported by Thorchain engine",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine, err := registry.GetEngine(tt.chain)

			if tt.shouldFind {
				if err != nil {
					t.Errorf("Expected to find engine for %s, but got error: %v", tt.chain.String(), err)
				}
				if engine == nil {
					t.Errorf("Expected to find engine for %s, but got nil", tt.chain.String())
				}

				// Verify the engine actually supports this chain
				if !engine.Supports(tt.chain) {
					t.Errorf("Engine claims to not support %s, but was returned for it", tt.chain.String())
				}
			} else {
				if err == nil {
					t.Errorf("Expected no engine for %s, but found one", tt.chain.String())
				}
				if engine != nil {
					t.Errorf("Expected nil engine for %s, but got: %T", tt.chain.String(), engine)
				}
			}
		})
	}
}

func TestChainEngineInterface(t *testing.T) {
	registry, err := NewChainEngineRegistry()
	require.NoError(t, err)

	// Test that all registered engines implement the interface correctly
	supportedChains := []common.Chain{
		common.Ethereum, common.BscChain, common.Arbitrum, // EVM chains
		common.Bitcoin, // BTC chains
		common.XRP, // XRPL chains
		common.Solana,  // Solana
		common.THORChain, // Thorchain
	}

	for _, chain := range supportedChains {
		t.Run(chain.String(), func(t *testing.T) {
			engine, err := registry.GetEngine(chain)
			if err != nil {
				t.Fatalf("Failed to get engine for %s: %v", chain.String(), err)
			}

			// Test that the engine correctly reports support
			if !engine.Supports(chain) {
				t.Errorf("Engine does not support %s but was returned for it", chain.String())
			}
		})
	}
}
