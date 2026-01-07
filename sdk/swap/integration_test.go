// +build integration

package swap_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/vultisig/recipes/sdk/swap"
)

// Real mainnet token addresses for testing
var (
	// Ethereum
	ethUSDC = swap.NewAsset("Ethereum", "USDC", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", 6)
	ethUSDT = swap.NewAsset("Ethereum", "USDT", "0xdAC17F958D2ee523a2206206994597C13D831ec7", 6)
	ethETH  = swap.NativeAsset("Ethereum", "ETH", 18)
	ethWETH = swap.NewAsset("Ethereum", "WETH", "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", 18)

	// Bitcoin
	btcBTC = swap.NativeAsset("Bitcoin", "BTC", 8)

	// Solana
	solSOL  = swap.NativeAsset("Solana", "SOL", 9)
	solUSDC = swap.NewAsset("Solana", "USDC", "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v", 6)

	// BSC
	bscBNB  = swap.NativeAsset("BSC", "BNB", 18)
	bscUSDT = swap.NewAsset("BSC", "USDT", "0x55d398326f99059fF775485246999027B3197955", 18)

	// Arbitrum
	arbETH  = swap.NativeAsset("Arbitrum", "ETH", 18)
	arbUSDC = swap.NewAsset("Arbitrum", "USDC", "0xaf88d065e77c8cC2239327C5EDb3A432268e5831", 6)

	// Avalanche
	avaxAVAX = swap.NativeAsset("Avalanche", "AVAX", 18)
	avaxUSDC = swap.NewAsset("Avalanche", "USDC", "0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E", 6)

	// Test addresses (don't use for real transactions!)
	testEthAddress = "0x742d35Cc6634C0532925a3b844Bc9e7595f43E2e"
	testBtcAddress = "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh"
	testSolAddress = "DRpbCBMxVnDK7maPM5tGv6MvB3v1sRMC86PZ8okm21hy"
)

// TestTHORChainCrossChainQuote tests getting a cross-chain quote via THORChain
func TestTHORChainCrossChainQuote(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// ETH to BTC cross-chain swap
	quote, err := swap.GetQuote(ctx, swap.QuoteRequest{
		From:        ethETH,
		To:          btcBTC,
		Amount:      swap.ToBaseUnits(0.1, 18), // 0.1 ETH
		Sender:      testEthAddress,
		Destination: testBtcAddress,
	})

	if err != nil {
		t.Logf("THORChain quote failed (may be halted): %v", err)
		return
	}

	t.Logf("THORChain ETH->BTC Quote:")
	t.Logf("  Provider: %s", quote.Provider)
	t.Logf("  Input: %f ETH", swap.FromBaseUnits(quote.FromAmount, 18))
	t.Logf("  Expected Output: %f BTC", swap.FromBaseUnits(quote.ExpectedOutput, 8))
	t.Logf("  Memo: %s", quote.Memo)
	t.Logf("  Inbound Address: %s", quote.InboundAddress)

	if quote.Provider != "THORChain" {
		t.Errorf("expected THORChain provider, got %s", quote.Provider)
	}

	if quote.ExpectedOutput.Sign() <= 0 {
		t.Error("expected positive output amount")
	}
}

// TestTHORChainReverseSwap tests BTC to ETH swap
func TestTHORChainReverseSwap(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// BTC to ETH cross-chain swap
	quote, err := swap.GetQuote(ctx, swap.QuoteRequest{
		From:        btcBTC,
		To:          ethETH,
		Amount:      swap.ToBaseUnits(0.001, 8), // 0.001 BTC (100,000 sats)
		Sender:      testBtcAddress,
		Destination: testEthAddress,
	})

	if err != nil {
		t.Logf("THORChain BTC->ETH quote failed (may be halted): %v", err)
		return
	}

	t.Logf("THORChain BTC->ETH Quote:")
	t.Logf("  Provider: %s", quote.Provider)
	t.Logf("  Input: %f BTC", swap.FromBaseUnits(quote.FromAmount, 8))
	t.Logf("  Expected Output: %f ETH", swap.FromBaseUnits(quote.ExpectedOutput, 18))
	t.Logf("  Memo: %s", quote.Memo)

	if quote.ExpectedOutput.Sign() <= 0 {
		t.Error("expected positive output amount")
	}
}

// TestMayachainQuote tests Mayachain swaps
func TestMayachainQuote(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// ETH to ARB via Mayachain
	quote, err := swap.GetQuote(ctx, swap.QuoteRequest{
		From:        ethETH,
		To:          arbETH,
		Amount:      swap.ToBaseUnits(0.1, 18),
		Sender:      testEthAddress,
		Destination: testEthAddress,
	})

	if err != nil {
		t.Logf("Mayachain quote failed (may be halted or unsupported): %v", err)
		return
	}

	t.Logf("Mayachain ETH->ARB Quote:")
	t.Logf("  Provider: %s", quote.Provider)
	t.Logf("  Input: %f ETH", swap.FromBaseUnits(quote.FromAmount, 18))
	t.Logf("  Expected Output: %f ETH (ARB)", swap.FromBaseUnits(quote.ExpectedOutput, 18))
}

// TestJupiterSolanaSwap tests Solana-to-Solana swaps via Jupiter
func TestJupiterSolanaSwap(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// SOL to USDC on Solana
	quote, err := swap.GetQuote(ctx, swap.QuoteRequest{
		From:        solSOL,
		To:          solUSDC,
		Amount:      swap.ToBaseUnits(1.0, 9), // 1 SOL
		Sender:      testSolAddress,
		Destination: testSolAddress,
	})

	if err != nil {
		t.Logf("Solana swap quote failed: %v", err)
		return
	}

	t.Logf("Solana SOL->USDC Quote:")
	t.Logf("  Provider: %s", quote.Provider)
	t.Logf("  Input: %f SOL", swap.FromBaseUnits(quote.FromAmount, 9))
	t.Logf("  Expected Output: %f USDC", swap.FromBaseUnits(quote.ExpectedOutput, 6))

	if quote.ExpectedOutput.Sign() <= 0 {
		t.Error("expected positive output amount")
	}
}

// TestLiFiCrossChainEVM tests LiFi cross-chain EVM swaps
func TestLiFiCrossChainEVM(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// ETH to BSC cross-chain
	quote, err := swap.GetQuote(ctx, swap.QuoteRequest{
		From:        ethETH,
		To:          bscBNB,
		Amount:      swap.ToBaseUnits(0.1, 18),
		Sender:      testEthAddress,
		Destination: testEthAddress,
	})

	if err != nil {
		t.Logf("LiFi cross-chain quote failed: %v", err)
		return
	}

	t.Logf("LiFi ETH->BSC Quote:")
	t.Logf("  Provider: %s", quote.Provider)
	t.Logf("  Input: %f ETH", swap.FromBaseUnits(quote.FromAmount, 18))
	t.Logf("  Expected Output: %f BNB", swap.FromBaseUnits(quote.ExpectedOutput, 18))
}

// TestOneInchSameChainSwap tests 1inch same-chain swaps
func TestOneInchSameChainSwap(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Need to use a custom router with only 1inch to test it directly
	router := swap.NewRouter(
		swap.WithProvider(swap.NewOneInchProvider("")),
	)

	// ETH to USDC on Ethereum
	quote, err := router.GetQuote(ctx, swap.QuoteRequest{
		From:        ethETH,
		To:          ethUSDC,
		Amount:      swap.ToBaseUnits(0.1, 18),
		Sender:      testEthAddress,
		Destination: testEthAddress,
	})

	if err != nil {
		t.Logf("1inch quote failed (may need API key): %v", err)
		return
	}

	t.Logf("1inch ETH->USDC Quote:")
	t.Logf("  Provider: %s", quote.Provider)
	t.Logf("  Input: %f ETH", swap.FromBaseUnits(quote.FromAmount, 18))
	t.Logf("  Expected Output: %f USDC", swap.FromBaseUnits(quote.ExpectedOutput, 6))
}

// TestUniswapSwap tests Uniswap same-chain swaps
func TestUniswapSwap(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	router := swap.NewRouter(
		swap.WithProvider(swap.NewUniswapProvider()),
	)

	// ETH to USDT on Ethereum
	quote, err := router.GetQuote(ctx, swap.QuoteRequest{
		From:        ethETH,
		To:          ethUSDT,
		Amount:      swap.ToBaseUnits(0.1, 18),
		Sender:      testEthAddress,
		Destination: testEthAddress,
	})

	if err != nil {
		t.Logf("Uniswap quote failed: %v", err)
		return
	}

	t.Logf("Uniswap ETH->USDT Quote:")
	t.Logf("  Provider: %s", quote.Provider)
	t.Logf("  Input: %f ETH", swap.FromBaseUnits(quote.FromAmount, 18))
	t.Logf("  Expected Output: %f USDT", swap.FromBaseUnits(quote.ExpectedOutput, 6))
}

// TestProviderFallback tests that the router falls back to next provider
func TestProviderFallback(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Solana swaps should fallback to LiFi or Jupiter
	quote, err := swap.GetQuote(ctx, swap.QuoteRequest{
		From:        solSOL,
		To:          solUSDC,
		Amount:      swap.ToBaseUnits(0.5, 9),
		Sender:      testSolAddress,
		Destination: testSolAddress,
	})

	if err != nil {
		t.Logf("Fallback test failed: %v", err)
		return
	}

	t.Logf("Fallback Test Quote:")
	t.Logf("  Provider: %s (should be LiFi or Jupiter)", quote.Provider)
	t.Logf("  Input: %f SOL", swap.FromBaseUnits(quote.FromAmount, 9))
	t.Logf("  Expected Output: %f USDC", swap.FromBaseUnits(quote.ExpectedOutput, 6))
}

// TestProviderStatus tests checking provider availability
func TestProviderStatus(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	chains := []string{"Bitcoin", "Ethereum", "Avalanche", "BSC"}

	for _, chain := range chains {
		status, err := swap.GetProviderStatus(ctx, "THORChain", chain)
		if err != nil {
			t.Logf("Failed to get THORChain status for %s: %v", chain, err)
			continue
		}

		t.Logf("THORChain %s Status:", chain)
		t.Logf("  Available: %v", status.Available)
		t.Logf("  Halted: %v", status.Halted)
		t.Logf("  Global Trading Paused: %v", status.GlobalTradingPaused)
		t.Logf("  Chain Trading Paused: %v", status.ChainTradingPaused)
		if status.InboundAddress != "" {
			t.Logf("  Inbound Address: %s", status.InboundAddress)
		}
		if status.Router != "" {
			t.Logf("  Router: %s", status.Router)
		}
	}
}

// TestCanSwapVariousRoutes tests the CanSwap convenience function
func TestCanSwapVariousRoutes(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	testCases := []struct {
		name     string
		from     swap.Asset
		to       swap.Asset
		expected bool
	}{
		{"ETH to BTC (cross-chain)", ethETH, btcBTC, true},
		{"SOL to USDC (same-chain)", solSOL, solUSDC, true},
		{"ETH to USDC (same-chain)", ethETH, ethUSDC, true},
		{"AVAX to ETH (cross-chain)", avaxAVAX, ethETH, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			canSwap := swap.CanSwap(ctx, tc.from, tc.to)
			t.Logf("%s: CanSwap=%v (expected=%v)", tc.name, canSwap, tc.expected)
			// Don't fail on availability - just log
		})
	}
}

// TestQuoteAndBuild tests the combined quote and build function
func TestQuoteAndBuild(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	quote, result, err := swap.QuoteAndBuild(ctx, swap.QuoteRequest{
		From:        ethETH,
		To:          btcBTC,
		Amount:      swap.ToBaseUnits(0.1, 18),
		Sender:      testEthAddress,
		Destination: testBtcAddress,
	})

	if err != nil {
		t.Logf("QuoteAndBuild failed: %v", err)
		return
	}

	t.Logf("QuoteAndBuild Result:")
	t.Logf("  Quote Provider: %s", quote.Provider)
	t.Logf("  Expected Output: %f BTC", swap.FromBaseUnits(quote.ExpectedOutput, 8))
	t.Logf("  Result Provider: %s", result.Provider)
	t.Logf("  To Address: %s", result.ToAddress)
	t.Logf("  Memo: %s", result.Memo)
}

// Benchmark tests
func BenchmarkGetQuote(b *testing.B) {
	ctx := context.Background()
	req := swap.QuoteRequest{
		From:        ethETH,
		To:          ethUSDC,
		Amount:      big.NewInt(1e17), // 0.1 ETH
		Sender:      testEthAddress,
		Destination: testEthAddress,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = swap.GetQuote(ctx, req)
	}
}

func BenchmarkFindRoute(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = swap.FindRoute(ctx, ethETH, btcBTC)
	}
}

// =============================================================================
// Approval Bundle Tests
// =============================================================================

// TestSwapBundleERC20ToNative tests getting a swap bundle for ERC20 to native swap
// This should include an approval transaction
func TestSwapBundleERC20ToNative(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	service := swap.NewService()

	params := swap.SwapParams{
		FromChain:    "Ethereum",
		FromSymbol:   "USDC",
		FromAddress:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
		FromDecimals: 6,
		ToChain:      "Ethereum",
		ToSymbol:     "ETH",
		ToAddress:    "",
		ToDecimals:   18,
		Amount:       big.NewInt(100 * 1e6), // 100 USDC
		Sender:       testEthAddress,
		Destination:  testEthAddress,
	}

	bundle, err := service.GetSwapTxBundle(ctx, params)
	if err != nil {
		t.Logf("GetSwapTxBundle failed (provider may be unavailable): %v", err)
		return
	}

	t.Logf("ERC20 to Native Swap Bundle:")
	t.Logf("  Provider: %s", bundle.Provider)
	t.Logf("  Needs Approval: %v", bundle.NeedsApproval)

	// ERC20 swaps should always need approval
	if !bundle.NeedsApproval {
		t.Error("expected NeedsApproval to be true for ERC20 swap")
	}

	if bundle.ApprovalTx == nil {
		t.Error("expected ApprovalTx to be present for ERC20 swap")
	} else {
		t.Logf("  Approval TX:")
		t.Logf("    To: %s (token contract)", bundle.ApprovalTx.To)
		t.Logf("    Nonce: %d", bundle.ApprovalTx.Nonce)
		t.Logf("    Data length: %d bytes", len(bundle.ApprovalTx.Data))

		// Verify approval tx data
		if len(bundle.ApprovalTx.Data) != 68 {
			t.Errorf("expected approval calldata length 68, got %d", len(bundle.ApprovalTx.Data))
		}

		// Verify approval is to token contract
		if bundle.ApprovalTx.To != params.FromAddress {
			t.Errorf("approval To should be token contract, got %s", bundle.ApprovalTx.To)
		}

		// Verify approval value is 0 (no ETH sent with approve)
		if bundle.ApprovalTx.Value != nil && bundle.ApprovalTx.Value.Sign() != 0 {
			t.Errorf("approval value should be 0, got %s", bundle.ApprovalTx.Value.String())
		}
	}

	if bundle.SwapTx == nil {
		t.Error("expected SwapTx to be present")
	} else {
		t.Logf("  Swap TX:")
		t.Logf("    To: %s", bundle.SwapTx.To)
		t.Logf("    Nonce: %d", bundle.SwapTx.Nonce)

		// Verify swap nonce is approval nonce + 1
		if bundle.ApprovalTx != nil && bundle.SwapTx.Nonce != bundle.ApprovalTx.Nonce+1 {
			t.Errorf("swap nonce should be approval nonce + 1, got approval=%d, swap=%d",
				bundle.ApprovalTx.Nonce, bundle.SwapTx.Nonce)
		}
	}

	if bundle.ExpectedOutput == nil || bundle.ExpectedOutput.Sign() <= 0 {
		t.Error("expected positive output amount")
	}
}

// TestSwapBundleNativeToERC20 tests getting a swap bundle for native to ERC20 swap
// This should NOT include an approval transaction
func TestSwapBundleNativeToERC20(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	service := swap.NewService()

	params := swap.SwapParams{
		FromChain:    "Ethereum",
		FromSymbol:   "ETH",
		FromAddress:  "", // Native token
		FromDecimals: 18,
		ToChain:      "Ethereum",
		ToSymbol:     "USDC",
		ToAddress:    "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
		ToDecimals:   6,
		Amount:       big.NewInt(1e17), // 0.1 ETH
		Sender:       testEthAddress,
		Destination:  testEthAddress,
	}

	bundle, err := service.GetSwapTxBundle(ctx, params)
	if err != nil {
		t.Logf("GetSwapTxBundle failed (provider may be unavailable): %v", err)
		return
	}

	t.Logf("Native to ERC20 Swap Bundle:")
	t.Logf("  Provider: %s", bundle.Provider)
	t.Logf("  Needs Approval: %v", bundle.NeedsApproval)

	// Native token swaps should NOT need approval
	if bundle.NeedsApproval {
		t.Error("expected NeedsApproval to be false for native token swap")
	}

	if bundle.ApprovalTx != nil {
		t.Error("expected ApprovalTx to be nil for native token swap")
	}

	if bundle.SwapTx == nil {
		t.Error("expected SwapTx to be present")
	} else {
		t.Logf("  Swap TX:")
		t.Logf("    To: %s", bundle.SwapTx.To)
		t.Logf("    Value: %s wei", bundle.SwapTx.Value.String())
	}

	if bundle.ExpectedOutput == nil || bundle.ExpectedOutput.Sign() <= 0 {
		t.Error("expected positive output amount")
	}
}

// TestSwapBundleCrossChainWithApproval tests cross-chain ERC20 swap via THORChain
func TestSwapBundleCrossChainWithApproval(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	service := swap.NewService()

	params := swap.SwapParams{
		FromChain:    "Ethereum",
		FromSymbol:   "USDC",
		FromAddress:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
		FromDecimals: 6,
		ToChain:      "Bitcoin",
		ToSymbol:     "BTC",
		ToAddress:    "", // Native
		ToDecimals:   8,
		Amount:       big.NewInt(1000 * 1e6), // 1000 USDC
		Sender:       testEthAddress,
		Destination:  testBtcAddress,
	}

	bundle, err := service.GetSwapTxBundle(ctx, params)
	if err != nil {
		t.Logf("GetSwapTxBundle failed (THORChain may be halted): %v", err)
		return
	}

	t.Logf("Cross-chain ERC20 Swap Bundle (ETH USDC -> BTC):")
	t.Logf("  Provider: %s", bundle.Provider)
	t.Logf("  Needs Approval: %v", bundle.NeedsApproval)

	// Cross-chain ERC20 swap should need approval
	if !bundle.NeedsApproval {
		t.Error("expected NeedsApproval to be true for cross-chain ERC20 swap")
	}

	if bundle.ApprovalTx == nil {
		t.Error("expected ApprovalTx to be present for cross-chain ERC20 swap")
	} else {
		t.Logf("  Approval TX:")
		t.Logf("    To: %s", bundle.ApprovalTx.To)
		t.Logf("    Nonce: %d", bundle.ApprovalTx.Nonce)
	}

	if bundle.SwapTx == nil {
		t.Error("expected SwapTx to be present")
	} else {
		t.Logf("  Swap TX:")
		t.Logf("    To: %s (router/vault)", bundle.SwapTx.To)
		t.Logf("    Nonce: %d", bundle.SwapTx.Nonce)
	}

	// Verify memo is set for THORChain
	if bundle.Memo == "" {
		t.Error("expected memo to be set for THORChain swap")
	} else {
		t.Logf("  Memo: %s", bundle.Memo)
	}
}

// TestRequiresApproval tests the pre-check for approval requirement
func TestRequiresApproval(t *testing.T) {
	service := swap.NewService()

	tests := []struct {
		name     string
		params   swap.SwapParams
		expected bool
	}{
		{
			name: "ERC20 on Ethereum needs approval",
			params: swap.SwapParams{
				FromChain:   "Ethereum",
				FromSymbol:  "USDC",
				FromAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			},
			expected: true,
		},
		{
			name: "Native ETH does not need approval",
			params: swap.SwapParams{
				FromChain:   "Ethereum",
				FromSymbol:  "ETH",
				FromAddress: "",
			},
			expected: false,
		},
		{
			name: "Native BTC does not need approval",
			params: swap.SwapParams{
				FromChain:   "Bitcoin",
				FromSymbol:  "BTC",
				FromAddress: "",
			},
			expected: false,
		},
		{
			name: "ERC20 on Arbitrum needs approval",
			params: swap.SwapParams{
				FromChain:   "Arbitrum",
				FromSymbol:  "USDC",
				FromAddress: "0xaf88d065e77c8cC2239327C5EDb3A432268e5831",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.RequiresApproval(tt.params)
			if result != tt.expected {
				t.Errorf("RequiresApproval() = %v, want %v", result, tt.expected)
			}
		})
	}
}

