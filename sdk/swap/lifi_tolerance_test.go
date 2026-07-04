package swap

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// Regression: LiFiProvider.GetQuote must forward a positive caller tolerance as
// the `slippage` query (a decimal fraction), and omit it otherwise so LiFi's
// own conservative default (~0.5%) stays in effect. The tolerance is carried
// into the Quote so BuildTx re-quotes with the same slippage.
func TestLiFiGetQuoteForwardsToleranceBps(t *testing.T) {
	var captured url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.URL.Query()
		_, _ = w.Write([]byte(`{"estimate":{"toAmount":"990000"},"transactionRequest":{"to":"0x1111111254EEB25477B68fb85Ed929f73A960582"}}`))
	}))
	defer srv.Close()

	p := NewLiFiProvider("")
	p.baseURL = srv.URL
	base := QuoteRequest{
		From:        Asset{Chain: "Ethereum", Symbol: "ETH", Decimals: 18},
		To:          Asset{Chain: "Ethereum", Symbol: "USDC", Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", Decimals: 6},
		Amount:      big.NewInt(1_000_000_000_000_000_000),
		Sender:      "0x1234567890123456789012345678901234567890",
		Destination: "0x1234567890123456789012345678901234567890",
	}

	t.Run("positive tolerance -> slippage fraction", func(t *testing.T) {
		captured = nil
		tol := 300
		req := base
		req.ToleranceBps = &tol
		q, err := p.GetQuote(context.Background(), req)
		if err != nil {
			t.Fatalf("GetQuote: %v", err)
		}
		if got := captured.Get("slippage"); got != "0.03" {
			t.Errorf("slippage = %q, want 0.03 (300bps)", got)
		}
		if q.ToleranceBps == nil || *q.ToleranceBps != 300 {
			t.Errorf("Quote.ToleranceBps = %v, want 300", q.ToleranceBps)
		}
	})

	t.Run("omitted tolerance -> no slippage param (LiFi default)", func(t *testing.T) {
		captured = nil
		if _, err := p.GetQuote(context.Background(), base); err != nil {
			t.Fatalf("GetQuote: %v", err)
		}
		if captured.Has("slippage") {
			t.Errorf("slippage param should be omitted when no tolerance given, got %q", captured.Get("slippage"))
		}
	})
}
