package swap

import (
	"math/big"
	"testing"
)

// TestMayaCACAODecimalScaling exercises the two decimal-conversion paths that
// are unique to CACAO on MayaChain.
//
// MayaChain empirically uses 10 decimal places for CACAO (verified 2026-06-01
// against /mayachain/quote/swap: 0.01 BTC → ~5776 CACAO at 10-dec, which
// matches the pool ratio; treating the same value as 8-dec would give ~100x
// the correct output).
//
// Non-CACAO assets (BTC, ETH, ZEC, DASH, …) still use the 8-decimal
// "thor-amount" convention, same as THORChain.
func TestMayaCACAODecimalScaling(t *testing.T) {
	t.Run("isMayaCACAO/true", func(t *testing.T) {
		a := Asset{Chain: "MayaChain", Symbol: "CACAO", Address: ""}
		if !isMayaCACAO(a) {
			t.Fatal("expected isMayaCACAO(MayaChain/CACAO native) = true")
		}
	})

	t.Run("isMayaCACAO/false_token", func(t *testing.T) {
		a := Asset{Chain: "MayaChain", Symbol: "MAYA", Address: "maya1something"}
		if isMayaCACAO(a) {
			t.Fatal("expected isMayaCACAO for non-native asset = false")
		}
	})

	t.Run("isMayaCACAO/false_other_chain", func(t *testing.T) {
		a := Asset{Chain: "Ethereum", Symbol: "ETH", Address: ""}
		if isMayaCACAO(a) {
			t.Fatal("expected isMayaCACAO for non-MayaChain asset = false")
		}
	})

	// For a CACAO input, the amount passed to the API should equal the original
	// base-unit amount (no scaling, since CACAO is already in 10-dec and Maya
	// wants 10-dec for CACAO input).
	t.Run("CACAO_input_no_scaling", func(t *testing.T) {
		// 1000 CACAO in 10-dec native base units
		cacao1000 := new(big.Int).Mul(big.NewInt(1000), new(big.Int).Exp(big.NewInt(10), big.NewInt(10), nil))

		from := Asset{Chain: "MayaChain", Symbol: "CACAO", Decimals: 10}
		var mayaAmount *big.Int
		if isMayaCACAO(from) {
			mayaAmount = new(big.Int).Set(cacao1000)
		} else {
			mayaAmount = toThorChainAmount(cacao1000, 10)
		}
		// Must equal the input — no division by 100 here.
		if mayaAmount.Cmp(cacao1000) != 0 {
			t.Fatalf("CACAO input scaling: got %s, want %s (100x error if 8-dec path taken)",
				mayaAmount, cacao1000)
		}
	})

	// For a non-CACAO input (e.g. ETH with 18 decimals), the amount IS
	// scaled to 8-dec via toThorChainAmount — unchanged from pre-fix behaviour.
	t.Run("ETH_input_scales_to_8dec", func(t *testing.T) {
		// 1 ETH in 18-dec base units
		eth1 := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
		from := Asset{Chain: "Ethereum", Symbol: "ETH", Decimals: 18}
		var mayaAmount *big.Int
		if isMayaCACAO(from) {
			mayaAmount = new(big.Int).Set(eth1)
		} else {
			mayaAmount = toThorChainAmount(eth1, 18)
		}
		// 1 ETH in 18-dec → 1e8 in 8-dec
		want := new(big.Int).Exp(big.NewInt(10), big.NewInt(8), nil)
		if mayaAmount.Cmp(want) != 0 {
			t.Fatalf("ETH input scaling: got %s, want %s", mayaAmount, want)
		}
	})

	// For a CACAO output, the raw API value (already in 10-dec) must pass
	// through unchanged. The previous bug multiplied this by 100 (treating 10-dec
	// as 8-dec and scaling to 10-dec), yielding a 100x inflated estimate.
	t.Run("CACAO_output_no_scaling", func(t *testing.T) {
		// Suppose Maya returns 57763399808087 for 0.01 BTC → CACAO.
		// At 10-dec this is ~5776 CACAO, which matches the pool price.
		apiOut := big.NewInt(57763399808087)
		to := Asset{Chain: "MayaChain", Symbol: "CACAO", Decimals: 10}
		var expectedOutput *big.Int
		if isMayaCACAO(to) {
			expectedOutput = apiOut
		} else {
			expectedOutput = fromThorChainAmount(apiOut, to.Decimals)
		}
		// Must equal the raw API value — no multiplication by 100.
		if expectedOutput.Cmp(apiOut) != 0 {
			t.Fatalf("CACAO output scaling: got %s, want %s (100x error if 8-dec path taken)",
				expectedOutput, apiOut)
		}
		// Sanity: reading the value at 10 dec should give ~5776.33 CACAO,
		// matching the live pool price for 0.01 BTC.
		cacoaHuman := new(big.Float).SetInt(expectedOutput)
		cacoaHuman.Quo(cacoaHuman, new(big.Float).SetFloat64(1e10))
		f, _ := cacoaHuman.Float64()
		if f < 1000 || f > 50000 {
			t.Fatalf("CACAO output sanity: got %.2f CACAO for 0.01 BTC, expected ~5776 (range 1000–50000)", f)
		}
	})

	// For a non-CACAO output (e.g. ETH), fromThorChainAmount scales from 8-dec
	// to native decimals — unchanged from pre-fix behaviour.
	t.Run("ETH_output_scales_from_8dec", func(t *testing.T) {
		// Maya returns 1e8 in 8-dec for 1 ETH
		apiOut := new(big.Int).Exp(big.NewInt(10), big.NewInt(8), nil)
		to := Asset{Chain: "Ethereum", Symbol: "ETH", Decimals: 18}
		var expectedOutput *big.Int
		if isMayaCACAO(to) {
			expectedOutput = apiOut
		} else {
			expectedOutput = fromThorChainAmount(apiOut, to.Decimals)
		}
		// 1e8 at 8-dec → 1e18 at 18-dec
		want := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
		if expectedOutput.Cmp(want) != 0 {
			t.Fatalf("ETH output scaling: got %s, want %s", expectedOutput, want)
		}
	})
}
