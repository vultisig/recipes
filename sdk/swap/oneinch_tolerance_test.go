package swap

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func intPtr(i int) *int { return &i }

// Unit-test the bps->percent conversion + fallbacks in isolation.
func TestOneInchSlippageParam(t *testing.T) {
	cases := []struct {
		name string
		in   *int
		want string
	}{
		{"nil -> 1% default", nil, "1"},
		{"zero -> 1% default", intPtr(0), "1"},
		{"negative -> 1% default", intPtr(-5), "1"},
		{"250bps -> 2.5%", intPtr(250), "2.5"},
		{"50bps -> 0.5%", intPtr(50), "0.5"},
		{"100bps -> 1%", intPtr(100), "1"},
		{"5000bps -> 50%", intPtr(5000), "50"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := oneInchSlippageParam(c.in); got != c.want {
				t.Errorf("oneInchSlippageParam(%v) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}

// Wiring: GetQuote must send the caller's tolerance as the `slippage` query
// param (in percent), not the hardcoded 1% default, and carry it into the
// returned Quote so BuildTx signs with the same slippage.
func TestOneInchGetQuoteForwardsToleranceBps(t *testing.T) {
	var captured url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.URL.Query()
		_, _ = w.Write([]byte(`{"dstAmount":"990000","tx":{"to":"0x1111111254EEB25477B68fb85Ed929f73A960582","data":"0x","value":"0"}}`))
	}))
	defer srv.Close()

	p := NewOneInchProvider("")
	p.baseURL = srv.URL
	tol := 250
	q, err := p.GetQuote(context.Background(), QuoteRequest{
		From:         Asset{Chain: "Ethereum", Symbol: "ETH", Decimals: 18},
		To:           Asset{Chain: "Ethereum", Symbol: "USDC", Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", Decimals: 6},
		Amount:       big.NewInt(1_000_000_000_000_000_000),
		Sender:       "0x1234567890123456789012345678901234567890",
		Destination:  "0x1234567890123456789012345678901234567890",
		ToleranceBps: &tol,
	})
	if err != nil {
		t.Fatalf("GetQuote: %v", err)
	}
	if got := captured.Get("slippage"); got != "2.5" {
		t.Errorf("outgoing slippage = %q, want 2.5", got)
	}
	if q.ToleranceBps == nil || *q.ToleranceBps != 250 {
		t.Errorf("Quote.ToleranceBps = %v, want 250 (must be carried for BuildTx)", q.ToleranceBps)
	}
}
