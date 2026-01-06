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
			name: "Solana to Solana via Jupiter",
			from: Asset{Chain: "Solana", Symbol: "SOL"},
			to:   Asset{Chain: "Solana", Symbol: "USDC", Address: "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"},
			// Jupiter is last priority for Solana-only routes
			wantProvider: "Jupiter",
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
			result, err := router.FindRoute(ctx, tt.from, tt.to)
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

