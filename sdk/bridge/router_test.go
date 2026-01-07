package bridge

import (
	"context"
	"math/big"
	"strings"
	"testing"
)

// TestBridgeAddressVerification verifies all bridge addresses are correctly formatted
// and match the expected official addresses
func TestBridgeAddressVerification(t *testing.T) {
	// Test LiFi Diamond addresses
	t.Run("LiFi Diamond Addresses", func(t *testing.T) {
		expectedLiFi := "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE"
		chainsWithSameAddress := []string{
			"Ethereum", "Arbitrum", "Optimism", "Polygon",
			"BSC", "Avalanche", "Base", "Fantom", "Gnosis", "Blast",
		}

		for _, chain := range chainsWithSameAddress {
			addr, ok := lifiDiamondAddresses[chain]
			if !ok {
				t.Errorf("Missing LiFi address for chain %s", chain)
				continue
			}
			if !strings.EqualFold(addr, expectedLiFi) {
				t.Errorf("LiFi address mismatch for %s: got %s, want %s", chain, addr, expectedLiFi)
			}
		}

		// zkSync has a different address
		zkSyncAddr := lifiDiamondAddresses["Zksync"]
		if strings.EqualFold(zkSyncAddr, expectedLiFi) {
			t.Error("zkSync should have a different LiFi address")
		}
		if !strings.HasPrefix(zkSyncAddr, "0x") || len(zkSyncAddr) != 42 {
			t.Errorf("Invalid zkSync LiFi address format: %s", zkSyncAddr)
		}
	})

	// Test L1 Bridge addresses (on Ethereum)
	t.Run("L1 Bridge Addresses", func(t *testing.T) {
		testCases := []struct {
			chain    string
			expected string
		}{
			{"Arbitrum", "0x72Ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef"},
			{"Optimism", "0x40E0C049f4671846E9Cff93AAEd88f2B48E527bB"},
			{"Base", "0x3154Cf16ccdb4C6d922629664174b904d80F2C35"},
		}

		for _, tc := range testCases {
			config, ok := l1BridgeAddresses[tc.chain]
			if !ok {
				t.Errorf("Missing L1 bridge config for %s", tc.chain)
				continue
			}
			if !strings.EqualFold(config.GatewayRouter, tc.expected) {
				t.Errorf("L1 bridge address mismatch for %s: got %s, want %s",
					tc.chain, config.GatewayRouter, tc.expected)
			}
		}
	})

	// Test L2 Bridge addresses
	t.Run("L2 Bridge Addresses", func(t *testing.T) {
		testCases := []struct {
			chain    string
			expected string
		}{
			{"Arbitrum", "0x5288c571Fd7aD117beA99bF60FE0846C4E84F933"},
			{"Optimism", "0x4200000000000000000000000000000000000010"},
			{"Base", "0x4200000000000000000000000000000000000010"},
		}

		for _, tc := range testCases {
			config, ok := l2BridgeAddresses[tc.chain]
			if !ok {
				t.Errorf("Missing L2 bridge config for %s", tc.chain)
				continue
			}
			if !strings.EqualFold(config.GatewayRouter, tc.expected) {
				t.Errorf("L2 bridge address mismatch for %s: got %s, want %s",
					tc.chain, config.GatewayRouter, tc.expected)
			}
		}
	})
}

// TestRouterProviderOrder verifies providers are sorted by priority
func TestRouterProviderOrder(t *testing.T) {
	router := NewDefaultRouter()
	providers := router.ListProviders()

	if len(providers) != 4 {
		t.Errorf("Expected 4 providers, got %d", len(providers))
	}

	// Expected order by priority:
	// 1. NativeL2 (priority 1)
	// 2. LiFi (priority 2)
	// 3. Across (priority 3)
	// 4. deBridge (priority 4)
	expectedOrder := []string{"NativeL2", "LiFi", "Across", "deBridge"}
	for i, expected := range expectedOrder {
		if i < len(providers) && providers[i] != expected {
			t.Errorf("Provider at index %d should be %s, got %s", i, expected, providers[i])
		}
	}
}

// TestLiFiProviderSupportsRoute verifies route support logic
func TestLiFiProviderSupportsRoute(t *testing.T) {
	provider := NewLiFiProvider("")

	testCases := []struct {
		name     string
		from     BridgeAsset
		to       BridgeAsset
		expected bool
	}{
		{
			name:     "ETH to Arbitrum ETH",
			from:     BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			to:       BridgeAsset{Chain: "Arbitrum", Symbol: "ETH"},
			expected: true,
		},
		{
			name:     "Same chain not supported",
			from:     BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			to:       BridgeAsset{Chain: "Ethereum", Symbol: "USDC"},
			expected: false,
		},
		{
			name:     "Unsupported source chain",
			from:     BridgeAsset{Chain: "Bitcoin", Symbol: "BTC"},
			to:       BridgeAsset{Chain: "Ethereum", Symbol: "WBTC"},
			expected: false,
		},
		{
			name:     "Cross EVM supported",
			from:     BridgeAsset{Chain: "Arbitrum", Symbol: "USDC"},
			to:       BridgeAsset{Chain: "Base", Symbol: "USDC"},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := provider.SupportsRoute(tc.from, tc.to)
			if result != tc.expected {
				t.Errorf("SupportsRoute(%v, %v) = %v, want %v",
					tc.from, tc.to, result, tc.expected)
			}
		})
	}
}

// TestNativeL2ProviderSupportsRoute verifies native bridge route support
func TestNativeL2ProviderSupportsRoute(t *testing.T) {
	provider := NewNativeL2Provider()

	testCases := []struct {
		name     string
		from     BridgeAsset
		to       BridgeAsset
		expected bool
	}{
		{
			name:     "ETH L1 to Arbitrum L2",
			from:     BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			to:       BridgeAsset{Chain: "Arbitrum", Symbol: "ETH"},
			expected: true,
		},
		{
			name:     "Arbitrum L2 to ETH L1",
			from:     BridgeAsset{Chain: "Arbitrum", Symbol: "ETH"},
			to:       BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			expected: true,
		},
		{
			name:     "L2 to L2 not supported by native bridge",
			from:     BridgeAsset{Chain: "Arbitrum", Symbol: "ETH"},
			to:       BridgeAsset{Chain: "Base", Symbol: "ETH"},
			expected: false,
		},
		{
			name:     "Same chain not supported",
			from:     BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			to:       BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := provider.SupportsRoute(tc.from, tc.to)
			if result != tc.expected {
				t.Errorf("SupportsRoute(%v, %v) = %v, want %v",
					tc.from, tc.to, result, tc.expected)
			}
		})
	}
}

// TestDeBridgeProviderSupportsRoute verifies deBridge route support
func TestDeBridgeProviderSupportsRoute(t *testing.T) {
	provider := NewDeBridgeProvider()

	testCases := []struct {
		name     string
		from     BridgeAsset
		to       BridgeAsset
		expected bool
	}{
		{
			name:     "ETH to Arbitrum",
			from:     BridgeAsset{Chain: "Ethereum", Symbol: "USDC", Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"},
			to:       BridgeAsset{Chain: "Arbitrum", Symbol: "USDC", Address: "0xaf88d065e77c8cC2239327C5EDb3A432268e5831"},
			expected: true,
		},
		{
			name:     "Base to Ethereum",
			from:     BridgeAsset{Chain: "Base", Symbol: "USDC", Address: "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"},
			to:       BridgeAsset{Chain: "Ethereum", Symbol: "USDC", Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"},
			expected: true,
		},
		{
			name:     "Unsupported chain fails",
			from:     BridgeAsset{Chain: "Bitcoin", Symbol: "BTC"},
			to:       BridgeAsset{Chain: "Ethereum", Symbol: "BTC"},
			expected: false,
		},
		{
			name:     "Same chain not supported",
			from:     BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			to:       BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := provider.SupportsRoute(tc.from, tc.to)
			if result != tc.expected {
				t.Errorf("SupportsRoute(%v, %v) = %v, want %v",
					tc.from, tc.to, result, tc.expected)
			}
		})
	}
}

// TestAcrossProviderSupportsRoute verifies Across route support
func TestAcrossProviderSupportsRoute(t *testing.T) {
	provider := NewAcrossProvider()

	testCases := []struct {
		name     string
		from     BridgeAsset
		to       BridgeAsset
		expected bool
	}{
		{
			name:     "ETH to Hyperliquid",
			from:     BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			to:       BridgeAsset{Chain: "Hyperliquid", Symbol: "ETH"},
			expected: true,
		},
		{
			name:     "Arbitrum to Base",
			from:     BridgeAsset{Chain: "Arbitrum", Symbol: "ETH"},
			to:       BridgeAsset{Chain: "Base", Symbol: "ETH"},
			expected: true,
		},
		{
			name:     "Same chain not supported",
			from:     BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			to:       BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			expected: false,
		},
		{
			name:     "Solana not supported by Across",
			from:     BridgeAsset{Chain: "Ethereum", Symbol: "USDC"},
			to:       BridgeAsset{Chain: "Solana", Symbol: "USDC"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := provider.SupportsRoute(tc.from, tc.to)
			if result != tc.expected {
				t.Errorf("SupportsRoute(%v, %v) = %v, want %v",
					tc.from, tc.to, result, tc.expected)
			}
		})
	}
}

// TestDeBridgeAddressVerification verifies deBridge address verification
func TestDeBridgeAddressVerification(t *testing.T) {
	provider := NewDeBridgeProvider()

	t.Run("Valid address passes", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Ethereum", "0xeF4fB24aD0916217251F553c0596F8Edc630EB66")
		if err != nil {
			t.Errorf("Expected valid address to pass verification: %v", err)
		}
	})

	t.Run("Case insensitive", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Ethereum", "0xef4fb24ad0916217251f553c0596f8edc630eb66")
		if err != nil {
			t.Errorf("Address verification should be case insensitive: %v", err)
		}
	})

	t.Run("Invalid address fails", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Ethereum", "0xdeadbeef")
		if err == nil {
			t.Error("Expected invalid address to fail verification")
		}
	})
}

// TestAcrossAddressVerification verifies Across address verification
func TestAcrossAddressVerification(t *testing.T) {
	provider := NewAcrossProvider()

	t.Run("Valid Ethereum SpokePool", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Ethereum", "0x5c7BCd6E7De5423a257D81B442095A1a6ced35C5")
		if err != nil {
			t.Errorf("Expected valid address to pass verification: %v", err)
		}
	})

	t.Run("Hyperliquid SpokePool", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Hyperliquid", "0x35E63eA3eb0fb7A3bc543C71FB66412e1F6B0E04")
		if err != nil {
			t.Errorf("Expected Hyperliquid address to pass verification: %v", err)
		}
	})

	t.Run("Invalid address fails", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Ethereum", "0xdeadbeef")
		if err == nil {
			t.Error("Expected invalid address to fail verification")
		}
	})
}

// TestBridgeAddressFormat verifies all addresses are properly formatted
func TestBridgeAddressFormat(t *testing.T) {
	validateEVMAddress := func(t *testing.T, name, addr string) {
		t.Helper()
		if !strings.HasPrefix(addr, "0x") {
			t.Errorf("%s: address should start with 0x: %s", name, addr)
		}
		if len(addr) != 42 {
			t.Errorf("%s: address should be 42 chars (including 0x): %s has %d", name, addr, len(addr))
		}
	}

	// Check LiFi addresses
	for chain, addr := range lifiDiamondAddresses {
		validateEVMAddress(t, "LiFi "+chain, addr)
	}

	// Check L1 bridge addresses
	for chain, config := range l1BridgeAddresses {
		validateEVMAddress(t, "L1 "+chain, config.GatewayRouter)
	}

	// Check L2 bridge addresses
	for chain, config := range l2BridgeAddresses {
		validateEVMAddress(t, "L2 "+chain, config.GatewayRouter)
	}

	// Check deBridge addresses
	for chain, addr := range debridgeDlnSourceAddresses {
		validateEVMAddress(t, "deBridge "+chain, addr)
	}

	// Check Across SpokePool addresses
	for chain, addr := range acrossSpokePoolAddresses {
		validateEVMAddress(t, "Across "+chain, addr)
	}
}

// TestLiFiAddressVerification verifies address verification function
func TestLiFiAddressVerification(t *testing.T) {
	provider := NewLiFiProvider("")

	t.Run("Valid address passes", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Ethereum", "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE")
		if err != nil {
			t.Errorf("Expected valid address to pass verification: %v", err)
		}
	})

	t.Run("Case insensitive", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Ethereum", "0x1231deb6f5749ef6ce6943a275a1d3e7486f4eae")
		if err != nil {
			t.Errorf("Address verification should be case insensitive: %v", err)
		}
	})

	t.Run("Invalid address fails", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Ethereum", "0xdeadbeef")
		if err == nil {
			t.Error("Expected invalid address to fail verification")
		}
	})

	t.Run("Unknown chain fails", func(t *testing.T) {
		err := provider.verifyBridgeAddress("UnknownChain", "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE")
		if err == nil {
			t.Error("Expected unknown chain to fail verification")
		}
	})
}

// TestNativeL2AddressVerification verifies native bridge address verification
func TestNativeL2AddressVerification(t *testing.T) {
	provider := NewNativeL2Provider()

	// Use known addresses from the maps
	arbL1Addr := l1BridgeAddresses["Arbitrum"].GatewayRouter
	arbL2Addr := l2BridgeAddresses["Arbitrum"].GatewayRouter

	t.Run("L1 to L2 valid", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Ethereum", "Arbitrum", arbL1Addr)
		if err != nil {
			t.Errorf("Expected valid L1 address to pass: %v", err)
		}
	})

	t.Run("L2 to L1 valid", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Arbitrum", "Ethereum", arbL2Addr)
		if err != nil {
			t.Errorf("Expected valid L2 address to pass: %v", err)
		}
	})

	t.Run("Invalid address fails", func(t *testing.T) {
		err := provider.verifyBridgeAddress("Ethereum", "Arbitrum", "0xdeadbeef")
		if err == nil {
			t.Error("Expected invalid address to fail")
		}
	})
}

// TestGetQuoteWithMockedAPI would test quote functionality with mocked responses
// Skipped in short tests as it would require HTTP mocking
func TestGetQuoteWithMockedAPI(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping API test in short mode")
	}
	// In a real implementation, we'd mock the HTTP client here
}

// TestRouterFindRoute tests the route finding logic
func TestRouterFindRoute(t *testing.T) {
	router := NewDefaultRouter()
	ctx := context.Background()

	t.Run("Cross chain route found", func(t *testing.T) {
		from := BridgeAsset{Chain: "Ethereum", Symbol: "ETH"}
		to := BridgeAsset{Chain: "Arbitrum", Symbol: "ETH"}

		result, err := router.FindRoute(ctx, from, to)
		if err != nil {
			t.Errorf("FindRoute failed: %v", err)
		}
		if result == nil {
			t.Error("Expected route result")
		}
		if !result.IsSupported {
			t.Error("Expected route to be supported")
		}
		// NativeL2 has highest priority for ETH L1->L2 routes
		if result.Provider != "NativeL2" {
			t.Errorf("Expected NativeL2 provider for ETH L1->L2 route, got %s", result.Provider)
		}
	})

	t.Run("Same chain route fails", func(t *testing.T) {
		from := BridgeAsset{Chain: "Ethereum", Symbol: "ETH"}
		to := BridgeAsset{Chain: "Ethereum", Symbol: "ETH"}

		_, err := router.FindRoute(ctx, from, to)
		if err == nil {
			t.Error("Expected same chain route to fail")
		}
	})
}

// TestNativeBridgeQuote tests native bridge quote generation
func TestNativeBridgeQuote(t *testing.T) {
	provider := NewNativeL2Provider()
	ctx := context.Background()

	// Use known address from the map
	arbL1Addr := l1BridgeAddresses["Arbitrum"].GatewayRouter

	t.Run("L1 to L2 quote", func(t *testing.T) {
		req := QuoteRequest{
			From:        BridgeAsset{Chain: "Ethereum", Symbol: "ETH"},
			To:          BridgeAsset{Chain: "Arbitrum", Symbol: "ETH"},
			Amount:      big.NewInt(1e18),
			Sender:      "0x1234567890123456789012345678901234567890",
			Destination: "0x1234567890123456789012345678901234567890",
		}

		quote, err := provider.GetQuote(ctx, req)
		if err != nil {
			t.Errorf("GetQuote failed: %v", err)
		}
		if quote == nil {
			t.Fatal("Expected quote")
		}

		// Native bridges are 1:1
		if quote.ExpectedOutput.Cmp(req.Amount) != 0 {
			t.Errorf("Native bridge should be 1:1, got %s want %s",
				quote.ExpectedOutput.String(), req.Amount.String())
		}

		// Verify bridge address
		if !strings.EqualFold(quote.Router, arbL1Addr) {
			t.Errorf("Wrong router address: got %s, want %s",
				quote.Router, arbL1Addr)
		}
	})
}
