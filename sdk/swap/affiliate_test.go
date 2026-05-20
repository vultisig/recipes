package swap

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// captureQuoteParams starts a test HTTP server that captures the URL query params
// from the first incoming request and returns a minimal valid JSON response so
// that the provider's JSON decoder doesn't error out.
func captureQuoteParams(t *testing.T) (server *httptest.Server, params func() url.Values) {
	t.Helper()
	captured := make(chan url.Values, 1)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case captured <- r.URL.Query():
		default:
		}
		// Minimal response - enough for the JSON decoder to not fail on the outer
		// envelope even if the provider returns an empty quote.  The test only
		// cares about the URL params, not the returned Quote struct.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Return an empty JSON object - providers should handle this gracefully
		// (they'll return an error, but we already have the params by then).
		_, _ = w.Write([]byte(`{}`))
	}))
	return srv, func() url.Values {
		select {
		case v := <-captured:
			return v
		default:
			return url.Values{}
		}
	}
}

func baseQuoteRequest() QuoteRequest {
	return QuoteRequest{
		From: Asset{
			Chain:    "Bitcoin",
			Symbol:   "BTC",
			Decimals: 8,
		},
		To: Asset{
			Chain:    "Ethereum",
			Symbol:   "ETH",
			Decimals: 18,
		},
		Amount:      big.NewInt(100000), // 0.001 BTC in satoshis
		Destination: "0xdestination",
		Sender:      "bc1qsender",
	}
}

func intPtr(v int) *int { return &v }

// ---------- THORChain affiliate tests ----------

func TestTHORChainAffiliate_NilBps_NoParams(t *testing.T) {
	srv, getParams := captureQuoteParams(t)
	defer srv.Close()

	provider := NewTHORChainProvider([]string{srv.URL})
	req := baseQuoteRequest()
	req.AffiliateBps = nil // not set at all
	req.AffiliateAddress = "thor1vultisig"

	// error is expected (empty response), but we only care about URL params
	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("affiliate") != "" {
		t.Errorf("expected no 'affiliate' param, got %q", params.Get("affiliate"))
	}
	if params.Get("affiliate_bps") != "" {
		t.Errorf("expected no 'affiliate_bps' param, got %q", params.Get("affiliate_bps"))
	}
}

func TestTHORChainAffiliate_ZeroBps_NoParams(t *testing.T) {
	srv, getParams := captureQuoteParams(t)
	defer srv.Close()

	provider := NewTHORChainProvider([]string{srv.URL})
	req := baseQuoteRequest()
	req.AffiliateBps = intPtr(0) // explicit zero
	req.AffiliateAddress = "thor1vultisig"

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("affiliate") != "" {
		t.Errorf("expected no 'affiliate' param for bps=0, got %q", params.Get("affiliate"))
	}
	if params.Get("affiliate_bps") != "" {
		t.Errorf("expected no 'affiliate_bps' param for bps=0, got %q", params.Get("affiliate_bps"))
	}
}

func TestTHORChainAffiliate_PositiveBps_EmptyAddress_NoParams(t *testing.T) {
	srv, getParams := captureQuoteParams(t)
	defer srv.Close()

	provider := NewTHORChainProvider([]string{srv.URL})
	req := baseQuoteRequest()
	req.AffiliateBps = intPtr(50)
	req.AffiliateAddress = "" // empty - should skip silently

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("affiliate") != "" {
		t.Errorf("expected no 'affiliate' param when address empty, got %q", params.Get("affiliate"))
	}
	if params.Get("affiliate_bps") != "" {
		t.Errorf("expected no 'affiliate_bps' param when address empty, got %q", params.Get("affiliate_bps"))
	}
}

func TestTHORChainAffiliate_PositiveBps_NonEmptyAddress_BothParams(t *testing.T) {
	srv, getParams := captureQuoteParams(t)
	defer srv.Close()

	provider := NewTHORChainProvider([]string{srv.URL})
	req := baseQuoteRequest()
	req.AffiliateBps = intPtr(50)
	req.AffiliateAddress = "thor1vultisig"

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if got := params.Get("affiliate"); got != "thor1vultisig" {
		t.Errorf("expected affiliate=thor1vultisig, got %q", got)
	}
	if got := params.Get("affiliate_bps"); got != "50" {
		t.Errorf("expected affiliate_bps=50, got %q", got)
	}
}

// ---------- Mayachain affiliate tests ----------

func TestMayachainAffiliate_NilBps_NoParams(t *testing.T) {
	srv, getParams := captureQuoteParams(t)
	defer srv.Close()

	provider := NewMayachainProvider([]string{srv.URL})
	req := baseQuoteRequest()
	req.From.Chain = "Bitcoin" // Mayachain supports BTC
	req.AffiliateBps = nil
	req.AffiliateAddress = "maya1vultisig"

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("affiliate") != "" {
		t.Errorf("expected no 'affiliate' param, got %q", params.Get("affiliate"))
	}
	if params.Get("affiliate_bps") != "" {
		t.Errorf("expected no 'affiliate_bps' param, got %q", params.Get("affiliate_bps"))
	}
}

func TestMayachainAffiliate_ZeroBps_NoParams(t *testing.T) {
	srv, getParams := captureQuoteParams(t)
	defer srv.Close()

	provider := NewMayachainProvider([]string{srv.URL})
	req := baseQuoteRequest()
	req.From.Chain = "Bitcoin"
	req.AffiliateBps = intPtr(0)
	req.AffiliateAddress = "maya1vultisig"

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("affiliate") != "" {
		t.Errorf("expected no 'affiliate' param for bps=0, got %q", params.Get("affiliate"))
	}
	if params.Get("affiliate_bps") != "" {
		t.Errorf("expected no 'affiliate_bps' param for bps=0, got %q", params.Get("affiliate_bps"))
	}
}

func TestMayachainAffiliate_PositiveBps_EmptyAddress_NoParams(t *testing.T) {
	srv, getParams := captureQuoteParams(t)
	defer srv.Close()

	provider := NewMayachainProvider([]string{srv.URL})
	req := baseQuoteRequest()
	req.From.Chain = "Bitcoin"
	req.AffiliateBps = intPtr(50)
	req.AffiliateAddress = ""

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("affiliate") != "" {
		t.Errorf("expected no 'affiliate' param when address empty, got %q", params.Get("affiliate"))
	}
	if params.Get("affiliate_bps") != "" {
		t.Errorf("expected no 'affiliate_bps' param when address empty, got %q", params.Get("affiliate_bps"))
	}
}

func TestMayachainAffiliate_PositiveBps_NonEmptyAddress_BothParams(t *testing.T) {
	srv, getParams := captureQuoteParams(t)
	defer srv.Close()

	provider := NewMayachainProvider([]string{srv.URL})
	req := baseQuoteRequest()
	req.From.Chain = "Bitcoin"
	req.AffiliateBps = intPtr(50)
	req.AffiliateAddress = "maya1vultisig"

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if got := params.Get("affiliate"); got != "maya1vultisig" {
		t.Errorf("expected affiliate=maya1vultisig, got %q", got)
	}
	if got := params.Get("affiliate_bps"); got != "50" {
		t.Errorf("expected affiliate_bps=50, got %q", got)
	}
}
