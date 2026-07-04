package swap

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// Regression: JupiterProvider.GetQuote must forward the caller's requested
// slippage (QuoteRequest.ToleranceBps) as slippageBps rather than always
// sending the hardcoded 1% default, so a caller can tighten or loosen the
// swap's min-out (baked into otherAmountThreshold).
func TestJupiterGetQuoteForwardsToleranceBps(t *testing.T) {
	var captured url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.URL.Query()
		_, _ = w.Write([]byte(`{"outAmount":"1000000","otherAmountThreshold":"990000","routePlan":[]}`))
	}))
	defer srv.Close()

	p := NewJupiterProvider(srv.URL)
	base := QuoteRequest{
		From:        Asset{Chain: "Solana", Symbol: "SOL", Decimals: 9},
		To:          Asset{Chain: "Solana", Symbol: "USDC", Address: "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v", Decimals: 6},
		Amount:      big.NewInt(1_000_000_000),
		Destination: "So11111111111111111111111111111111111111112",
		Sender:      "So11111111111111111111111111111111111111112",
	}

	t.Run("explicit tolerance forwarded", func(t *testing.T) {
		captured = nil
		tol := 250
		req := base
		req.ToleranceBps = &tol
		if _, err := p.GetQuote(context.Background(), req); err != nil {
			t.Fatalf("GetQuote: %v", err)
		}
		if got := captured.Get("slippageBps"); got != "250" {
			t.Errorf("slippageBps = %q, want 250", got)
		}
	})

	t.Run("omitted tolerance uses 1% default", func(t *testing.T) {
		captured = nil
		if _, err := p.GetQuote(context.Background(), base); err != nil {
			t.Fatalf("GetQuote: %v", err)
		}
		if got := captured.Get("slippageBps"); got != "100" {
			t.Errorf("default slippageBps = %q, want 100", got)
		}
	})

	t.Run("zero tolerance falls back to default (not exact-out)", func(t *testing.T) {
		captured = nil
		zero := 0
		req := base
		req.ToleranceBps = &zero
		if _, err := p.GetQuote(context.Background(), req); err != nil {
			t.Fatalf("GetQuote: %v", err)
		}
		if got := captured.Get("slippageBps"); got != "100" {
			t.Errorf("zero-tolerance slippageBps = %q, want 100 (fallback)", got)
		}
	})
}
