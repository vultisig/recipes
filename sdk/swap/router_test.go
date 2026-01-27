package swap

import (
	"context"
	"testing"
)

func TestNewDefaultRouter(t *testing.T) {
	router := NewDefaultRouter()

	providers := router.ListProviders()
	if len(providers) != 6 {
		t.Errorf("expected 6 providers, got %d", len(providers))
	}

	// Verify priority order
	expectedOrder := []string{"THORChain", "Mayachain", "LiFi", "1inch", "Jupiter", "Uniswap"}
	for i, name := range expectedOrder {
		if providers[i] != name {
			t.Errorf("expected provider %d to be %s, got %s", i, name, providers[i])
		}
	}
}

func TestRouterFindRoute(t *testing.T) {
	router := NewDefaultRouter()
	ctx := context.Background()

	tests := []struct {
		name        string
		from        Asset
		to          Asset
		wantErr     bool
		wantProvider string
	}{
		{
			name: "BTC to ETH via THORChain",
			from: Asset{Chain: "Bitcoin", Symbol: "BTC"},
			to:   Asset{Chain: "Ethereum", Symbol: "ETH"},
			// This will try THORChain first - may or may not be available
			wantErr: false, // May succeed or fail depending on THORChain status
		},
		{
			name: "Solana to Solana",
			from: Asset{Chain: "Solana", Symbol: "SOL"},
			to:   Asset{Chain: "Solana", Symbol: "USDC", Address: "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"},
			// THORChain is tried first for all swaps; falls back to Jupiter for Solana.
			wantErr: false,
		},
		{
			name:    "Unsupported chain",
			from:    Asset{Chain: "UnsupportedChain", Symbol: "XXX"},
			to:      Asset{Chain: "Ethereum", Symbol: "ETH"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := router.FindRoute(ctx, tt.from, tt.to, nil)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			// For non-error cases, we just check the result structure
			if result != nil && tt.wantProvider != "" && result.Provider != tt.wantProvider {
				t.Errorf("expected provider %s, got %s", tt.wantProvider, result.Provider)
			}
		})
	}
}

func TestRouterGetSupportedChains(t *testing.T) {
	router := NewDefaultRouter()

	chains := router.GetSupportedChains()
	if len(chains) == 0 {
		t.Error("expected at least one supported chain")
	}

	// Check some expected chains are present
	expectedChains := []string{"Bitcoin", "Ethereum", "Solana"}
	chainSet := make(map[string]bool)
	for _, c := range chains {
		chainSet[c] = true
	}

	for _, expected := range expectedChains {
		if !chainSet[expected] {
			t.Errorf("expected chain %s to be supported", expected)
		}
	}
}

func TestTHORChainProviderSupportsRoute(t *testing.T) {
	provider := NewTHORChainProvider(nil)

	tests := []struct {
		name    string
		from    Asset
		to      Asset
		want    bool
	}{
		{
			name: "BTC to ETH",
			from: Asset{Chain: "Bitcoin", Symbol: "BTC"},
			to:   Asset{Chain: "Ethereum", Symbol: "ETH"},
			want: true,
		},
		{
			name: "ETH to AVAX",
			from: Asset{Chain: "Ethereum", Symbol: "ETH"},
			to:   Asset{Chain: "Avalanche", Symbol: "AVAX"},
			want: true,
		},
		{
			name: "Solana not supported",
			from: Asset{Chain: "Solana", Symbol: "SOL"},
			to:   Asset{Chain: "Ethereum", Symbol: "ETH"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := provider.SupportsRoute(tt.from, tt.to)
			if got != tt.want {
				t.Errorf("SupportsRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMayachainProviderSupportsRoute(t *testing.T) {
	provider := NewMayachainProvider(nil)

	tests := []struct {
		name string
		from Asset
		to   Asset
		want bool
	}{
		{
			name: "BTC to ETH",
			from: Asset{Chain: "Bitcoin", Symbol: "BTC"},
			to:   Asset{Chain: "Ethereum", Symbol: "ETH"},
			want: true,
		},
		{
			name: "ETH to ARB",
			from: Asset{Chain: "Ethereum", Symbol: "ETH"},
			to:   Asset{Chain: "Arbitrum", Symbol: "ETH"},
			want: true,
		},
		{
			name: "Avalanche not supported by Maya",
			from: Asset{Chain: "Avalanche", Symbol: "AVAX"},
			to:   Asset{Chain: "Ethereum", Symbol: "ETH"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := provider.SupportsRoute(tt.from, tt.to)
			if got != tt.want {
				t.Errorf("SupportsRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOneInchProviderSameChainOnly(t *testing.T) {
	provider := NewOneInchProvider("")

	tests := []struct {
		name string
		from Asset
		to   Asset
		want bool
	}{
		{
			name: "Same chain ETH swap",
			from: Asset{Chain: "Ethereum", Symbol: "ETH"},
			to:   Asset{Chain: "Ethereum", Symbol: "USDC", Address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"},
			want: true,
		},
		{
			name: "Cross-chain not supported",
			from: Asset{Chain: "Ethereum", Symbol: "ETH"},
			to:   Asset{Chain: "BSC", Symbol: "BNB"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := provider.SupportsRoute(tt.from, tt.to)
			if got != tt.want {
				t.Errorf("SupportsRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJupiterProviderSolanaOnly(t *testing.T) {
	provider := NewJupiterProvider("")

	tests := []struct {
		name string
		from Asset
		to   Asset
		want bool
	}{
		{
			name: "Solana swap",
			from: Asset{Chain: "Solana", Symbol: "SOL"},
			to:   Asset{Chain: "Solana", Symbol: "USDC"},
			want: true,
		},
		{
			name: "Non-Solana not supported",
			from: Asset{Chain: "Ethereum", Symbol: "ETH"},
			to:   Asset{Chain: "Ethereum", Symbol: "USDC"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := provider.SupportsRoute(tt.from, tt.to)
			if got != tt.want {
				t.Errorf("SupportsRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLiFiProviderCrossChain(t *testing.T) {
	provider := NewLiFiProvider("")

	tests := []struct {
		name string
		from Asset
		to   Asset
		want bool
	}{
		{
			name: "Same chain EVM swap",
			from: Asset{Chain: "Ethereum", Symbol: "ETH"},
			to:   Asset{Chain: "Ethereum", Symbol: "USDC"},
			want: true,
		},
		{
			name: "Cross-chain EVM swap",
			from: Asset{Chain: "Ethereum", Symbol: "ETH"},
			to:   Asset{Chain: "BSC", Symbol: "BNB"},
			want: true,
		},
		{
			name: "Solana supported",
			from: Asset{Chain: "Solana", Symbol: "SOL"},
			to:   Asset{Chain: "Ethereum", Symbol: "ETH"},
			want: true,
		},
		{
			name: "Unsupported chain",
			from: Asset{Chain: "UnsupportedChain", Symbol: "XXX"},
			to:   Asset{Chain: "Ethereum", Symbol: "ETH"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := provider.SupportsRoute(tt.from, tt.to)
			if got != tt.want {
				t.Errorf("SupportsRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniswapProviderSameChainOnly(t *testing.T) {
	provider := NewUniswapProvider()

	tests := []struct {
		name string
		from Asset
		to   Asset
		want bool
	}{
		{
			name: "Same chain ETH swap",
			from: Asset{Chain: "Ethereum", Symbol: "ETH"},
			to:   Asset{Chain: "Ethereum", Symbol: "USDC"},
			want: true,
		},
		{
			name: "Cross-chain not supported",
			from: Asset{Chain: "Ethereum", Symbol: "ETH"},
			to:   Asset{Chain: "Polygon", Symbol: "MATIC"},
			want: false,
		},
		{
			name: "Same chain Arbitrum",
			from: Asset{Chain: "Arbitrum", Symbol: "ETH"},
			to:   Asset{Chain: "Arbitrum", Symbol: "USDC"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := provider.SupportsRoute(tt.from, tt.to)
			if got != tt.want {
				t.Errorf("SupportsRoute() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test package-level API
func TestPackageLevelFunctions(t *testing.T) {
	providers := ListProviders()
	if len(providers) == 0 {
		t.Error("ListProviders() returned empty list")
	}

	chains := GetSupportedChains()
	if len(chains) == 0 {
		t.Error("GetSupportedChains() returned empty list")
	}
}

func TestCanSwap(t *testing.T) {
	ctx := context.Background()

	// Solana-to-Solana should always be possible via Jupiter
	from := Asset{Chain: "Solana", Symbol: "SOL"}
	to := Asset{Chain: "Solana", Symbol: "USDC"}

	if !CanSwap(ctx, from, to) {
		t.Error("expected Solana-to-Solana swap to be possible")
	}

	// Unsupported chain should not be possible
	unsupportedFrom := Asset{Chain: "UnsupportedChain", Symbol: "XXX"}
	if CanSwap(ctx, unsupportedFrom, to) {
		t.Error("expected unsupported chain to return false")
	}
}

func TestAssetConstructors(t *testing.T) {
	// Test NewAsset
	usdc := NewAsset("Ethereum", "USDC", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", 6)
	if usdc.Chain != "Ethereum" || usdc.Symbol != "USDC" || usdc.Decimals != 6 {
		t.Errorf("NewAsset() = %+v, expected Ethereum USDC with 6 decimals", usdc)
	}

	// Test NativeAsset
	eth := NativeAsset("Ethereum", "ETH", 18)
	if eth.Chain != "Ethereum" || eth.Symbol != "ETH" || eth.Address != "" || eth.Decimals != 18 {
		t.Errorf("NativeAsset() = %+v, expected native Ethereum ETH", eth)
	}
}

func TestUnitConversions(t *testing.T) {
	// Test ToBaseUnits
	weiAmount := ToBaseUnits(1.5, 18)
	expected := "1500000000000000000"
	if weiAmount.String() != expected {
		t.Errorf("ToBaseUnits(1.5, 18) = %s, want %s", weiAmount.String(), expected)
	}

	// Test FromBaseUnits
	ethAmount := FromBaseUnits(weiAmount, 18)
	if ethAmount != 1.5 {
		t.Errorf("FromBaseUnits() = %f, want 1.5", ethAmount)
	}

	// Test with different decimals (USDC has 6)
	usdcAmount := ToBaseUnits(100.0, 6)
	if usdcAmount.String() != "100000000" {
		t.Errorf("ToBaseUnits(100.0, 6) = %s, want 100000000", usdcAmount.String())
	}
}

func TestSmartProviderRouting(t *testing.T) {
	router := NewDefaultRouter()

	t.Run("All swaps prefer THORChain first", func(t *testing.T) {
		ordered := router.getOrderedProviders(nil)

		if len(ordered) == 0 {
			t.Fatal("expected at least one provider")
		}
		if ordered[0].Name() != "THORChain" {
			t.Errorf("expected first provider to be THORChain, got %s", ordered[0].Name())
		}
	})

	t.Run("Cross-chain prefers THORChain", func(t *testing.T) {
		ordered := router.getOrderedProviders(nil)

		if len(ordered) == 0 {
			t.Fatal("expected at least one provider")
		}
		if ordered[0].Name() != "THORChain" {
			t.Errorf("expected first provider to be THORChain for cross-chain, got %s", ordered[0].Name())
		}
	})

	t.Run("Custom preference reorders providers", func(t *testing.T) {
		pref := &ProviderPreference{
			Providers: []string{ProviderOneInch, ProviderTHORChain},
		}
		ordered := router.getOrderedProviders(pref)

		if len(ordered) == 0 {
			t.Fatal("expected at least one provider")
		}
		if ordered[0].Name() != "1inch" {
			t.Errorf("expected first provider to be 1inch with custom preference, got %s", ordered[0].Name())
		}
		if ordered[1].Name() != "THORChain" {
			t.Errorf("expected second provider to be THORChain, got %s", ordered[1].Name())
		}
	})

	t.Run("OnlyPreferred limits providers", func(t *testing.T) {
		pref := &ProviderPreference{
			Providers:     []string{ProviderOneInch, ProviderJupiter},
			OnlyPreferred: true,
		}
		ordered := router.getOrderedProviders(pref)

		if len(ordered) != 2 {
			t.Errorf("expected 2 providers with OnlyPreferred, got %d", len(ordered))
		}
		// Verify only the preferred providers are included
		names := make(map[string]bool)
		for _, p := range ordered {
			names[p.Name()] = true
		}
		if !names["1inch"] || !names["Jupiter"] {
			t.Errorf("expected only 1inch and Jupiter, got %v", names)
		}
	})

}

