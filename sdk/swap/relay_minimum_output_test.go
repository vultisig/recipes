package swap

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
)

// FUND-SAFETY regression: RelayProvider.GetQuote must report the API's real
// currencyOut.minimumAmount as Quote.MinimumOutput, not copy expectedOutput.
// Relay quotes are NOT zero-slippage (live: isFixedRate:false, ~1% gap between
// amount and minimumAmount), so copying expectedOutput hands callers a false
// zero-slippage floor.
func TestRelayGetQuoteUsesMinimumAmount(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Minimal quote: distinct amount vs minimumAmount, no steps.
		_, _ = w.Write([]byte(`{"steps":[],"details":{"currencyOut":{"amount":"1000000","minimumAmount":"990000","currency":{"symbol":"USDC"}}}}`))
	}))
	defer srv.Close()

	p := NewRelayProvider(nil)
	p.baseURL = srv.URL

	q, err := p.GetQuote(context.Background(), QuoteRequest{
		From:        Asset{Chain: "Ethereum", Symbol: "ETH", Decimals: 18},
		To:          Asset{Chain: "Ethereum", Symbol: "USDC", Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", Decimals: 6},
		Amount:      big.NewInt(1_000_000_000_000_000_000),
		Destination: "0x1234567890123456789012345678901234567890",
		Sender:      "0x1234567890123456789012345678901234567890",
	})
	if err != nil {
		t.Fatalf("GetQuote: %v", err)
	}
	if q.ExpectedOutput.String() != "1000000" {
		t.Errorf("ExpectedOutput = %s, want 1000000", q.ExpectedOutput)
	}
	if q.MinimumOutput.String() != "990000" {
		t.Errorf("MinimumOutput = %s, want 990000 (currencyOut.minimumAmount, NOT expectedOutput)", q.MinimumOutput)
	}

	// When the API omits minimumAmount, fall back to expectedOutput (no crash).
	srv2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"steps":[],"details":{"currencyOut":{"amount":"1000000","currency":{"symbol":"USDC"}}}}`))
	}))
	defer srv2.Close()
	p.baseURL = srv2.URL
	q2, err := p.GetQuote(context.Background(), QuoteRequest{
		From:        Asset{Chain: "Ethereum", Symbol: "ETH", Decimals: 18},
		To:          Asset{Chain: "Ethereum", Symbol: "USDC", Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", Decimals: 6},
		Amount:      big.NewInt(1_000_000_000_000_000_000),
		Destination: "0x1234567890123456789012345678901234567890",
		Sender:      "0x1234567890123456789012345678901234567890",
	})
	if err != nil {
		t.Fatalf("GetQuote (no minimumAmount): %v", err)
	}
	if q2.MinimumOutput.String() != "1000000" {
		t.Errorf("fallback MinimumOutput = %s, want 1000000 (expectedOutput)", q2.MinimumOutput)
	}
}
