package swap

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// FUND-SAFETY regression: MayachainProvider.GetQuote must forward tolerance_bps
// into the Mayanode /quote/swap request so the returned memo carries a LIM
// (minimum-output) field. A missing tolerance_bps yields a market-order memo
// ("=:c:maya1...", no limit) that executes with ZERO slippage protection —
// full MEV/sandwich exposure on the whole amount — regardless of the caller's
// requested tolerance. This mirrors THORChainProvider.GetQuote, which was
// already forwarding it.
func TestMayachainGetQuoteForwardsToleranceBps(t *testing.T) {
	var captured url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.URL.Query()
		// We only assert on the OUTGOING query; the request is captured before
		// this returns, so a stub error response is enough (no need to build a
		// full valid quote body).
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":"stub"}`))
	}))
	defer srv.Close()

	p := NewMayachainProvider([]string{srv.URL})
	base := QuoteRequest{
		From:        Asset{Chain: "Bitcoin", Symbol: "BTC", Decimals: 8},
		To:          Asset{Chain: "MayaChain", Symbol: "CACAO", Decimals: 10},
		Amount:      big.NewInt(100_000_000),
		Destination: "maya1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq4n8gyj",
	}

	t.Run("explicit tolerance is forwarded", func(t *testing.T) {
		captured = nil
		tol := 150
		req := base
		req.ToleranceBps = &tol
		_, _ = p.GetQuote(context.Background(), req)
		if captured == nil {
			t.Fatal("request never reached the server (GetQuote errored before fetchQuote)")
		}
		if got := captured.Get("tolerance_bps"); got != "150" {
			t.Errorf("tolerance_bps = %q, want 150", got)
		}
	})

	t.Run("omitted tolerance defaults to 2500 (parity with THORChain)", func(t *testing.T) {
		captured = nil
		_, _ = p.GetQuote(context.Background(), base)
		if captured == nil {
			t.Fatal("request never reached the server")
		}
		if got := captured.Get("tolerance_bps"); got != "2500" {
			t.Errorf("default tolerance_bps = %q, want 2500", got)
		}
	})
}
