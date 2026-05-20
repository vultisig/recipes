package swap

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// captureSwapParams starts a test HTTP server that captures the URL query
// params from the first incoming request and returns a minimal JSON response.
// Callers only care about the URL params, not the returned struct.
func captureSwapParams(t *testing.T, body string) (server *httptest.Server, params func() url.Values) {
	t.Helper()
	captured := make(chan url.Values, 1)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case captured <- r.URL.Query():
		default:
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
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

func intPtr(v int) *int { return &v }

func evmQuoteRequest() QuoteRequest {
	return QuoteRequest{
		From: Asset{
			Chain:    "Ethereum",
			Symbol:   "ETH",
			Decimals: 18,
		},
		To: Asset{
			Chain:    "Ethereum",
			Symbol:   "USDC",
			Address:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			Decimals: 6,
		},
		Amount:      big.NewInt(1e18),
		Destination: "0xdestination",
		Sender:      "0xsender",
	}
}

// minimal 1inch /swap response so the JSON decoder won't fail on DstAmount
const oneInchMinimalResp = `{"dstAmount":"1000000","tx":{"from":"0xsender","to":"0xrouter","data":"0x","value":"0","gas":200000,"gasPrice":"1000000000"}}`

// minimal LiFi /quote response
const lifiMinimalResp = `{"estimate":{"fromAmount":"1000000000000000000","toAmount":"1000000"},"transactionRequest":{"from":"0xsender","to":"0xrouter","data":"0x","value":"0"}}`

// ---------- 1inch affiliate tests ----------

func TestOneInchAffiliate_NilBps_NoParams(t *testing.T) {
	srv, getParams := captureSwapParams(t, oneInchMinimalResp)
	defer srv.Close()

	provider := &OneInchProvider{
		BaseProvider: NewBaseProvider("1inch", PriorityOneInch, oneInchSupportedChains),
		client:       srv.Client(),
		baseURL:      srv.URL,
	}
	req := evmQuoteRequest()
	req.AffiliateBps = nil
	req.AffiliateAddress = "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9"

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("referrer") != "" {
		t.Errorf("expected no 'referrer' param, got %q", params.Get("referrer"))
	}
	if params.Get("fee") != "" {
		t.Errorf("expected no 'fee' param, got %q", params.Get("fee"))
	}
}

func TestOneInchAffiliate_ZeroBps_NoParams(t *testing.T) {
	srv, getParams := captureSwapParams(t, oneInchMinimalResp)
	defer srv.Close()

	provider := &OneInchProvider{
		BaseProvider: NewBaseProvider("1inch", PriorityOneInch, oneInchSupportedChains),
		client:       srv.Client(),
		baseURL:      srv.URL,
	}
	req := evmQuoteRequest()
	req.AffiliateBps = intPtr(0)
	req.AffiliateAddress = "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9"

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("referrer") != "" {
		t.Errorf("expected no 'referrer' param for bps=0, got %q", params.Get("referrer"))
	}
	if params.Get("fee") != "" {
		t.Errorf("expected no 'fee' param for bps=0, got %q", params.Get("fee"))
	}
}

func TestOneInchAffiliate_PositiveBps_EmptyAddress_NoParams(t *testing.T) {
	srv, getParams := captureSwapParams(t, oneInchMinimalResp)
	defer srv.Close()

	provider := &OneInchProvider{
		BaseProvider: NewBaseProvider("1inch", PriorityOneInch, oneInchSupportedChains),
		client:       srv.Client(),
		baseURL:      srv.URL,
	}
	req := evmQuoteRequest()
	req.AffiliateBps = intPtr(50)
	req.AffiliateAddress = "" // empty - should skip silently

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("referrer") != "" {
		t.Errorf("expected no 'referrer' param when address empty, got %q", params.Get("referrer"))
	}
	if params.Get("fee") != "" {
		t.Errorf("expected no 'fee' param when address empty, got %q", params.Get("fee"))
	}
}

func TestOneInchAffiliate_PositiveBps_NonEmptyAddress_BothParams(t *testing.T) {
	srv, getParams := captureSwapParams(t, oneInchMinimalResp)
	defer srv.Close()

	provider := &OneInchProvider{
		BaseProvider: NewBaseProvider("1inch", PriorityOneInch, oneInchSupportedChains),
		client:       srv.Client(),
		baseURL:      srv.URL,
	}
	req := evmQuoteRequest()
	req.AffiliateBps = intPtr(50)
	req.AffiliateAddress = "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9"

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if got := params.Get("referrer"); got != "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9" {
		t.Errorf("expected referrer=0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9, got %q", got)
	}
	// 50 bps / 100 = 0.5000
	if got := params.Get("fee"); got != "0.5000" {
		t.Errorf("expected fee=0.5000 for 50bps, got %q", got)
	}
}

// ---------- LiFi affiliate tests ----------

func TestLiFiAffiliate_NilBps_NoFeeParam(t *testing.T) {
	srv, getParams := captureSwapParams(t, lifiMinimalResp)
	defer srv.Close()

	provider := &LiFiProvider{
		BaseProvider: NewBaseProvider("LiFi", PriorityLiFi, lifiSupportedChains),
		client:       srv.Client(),
		baseURL:      srv.URL,
	}
	req := evmQuoteRequest()
	req.AffiliateBps = nil
	req.AffiliateAddress = "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9"

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("fee") != "" {
		t.Errorf("expected no 'fee' param, got %q", params.Get("fee"))
	}
	// integrator must always be present
	if got := params.Get("integrator"); got != lifiIntegratorName {
		t.Errorf("expected integrator=%q, got %q", lifiIntegratorName, got)
	}
}

func TestLiFiAffiliate_ZeroBps_NoFeeParam(t *testing.T) {
	srv, getParams := captureSwapParams(t, lifiMinimalResp)
	defer srv.Close()

	provider := &LiFiProvider{
		BaseProvider: NewBaseProvider("LiFi", PriorityLiFi, lifiSupportedChains),
		client:       srv.Client(),
		baseURL:      srv.URL,
	}
	req := evmQuoteRequest()
	req.AffiliateBps = intPtr(0)
	req.AffiliateAddress = "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9"

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("fee") != "" {
		t.Errorf("expected no 'fee' param for bps=0, got %q", params.Get("fee"))
	}
}

func TestLiFiAffiliate_PositiveBps_EmptyAddress_NoFeeParam(t *testing.T) {
	srv, getParams := captureSwapParams(t, lifiMinimalResp)
	defer srv.Close()

	provider := &LiFiProvider{
		BaseProvider: NewBaseProvider("LiFi", PriorityLiFi, lifiSupportedChains),
		client:       srv.Client(),
		baseURL:      srv.URL,
	}
	req := evmQuoteRequest()
	req.AffiliateBps = intPtr(50)
	req.AffiliateAddress = "" // empty - should skip silently

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	if params.Get("fee") != "" {
		t.Errorf("expected no 'fee' param when address empty, got %q", params.Get("fee"))
	}
}

func TestLiFiAffiliate_PositiveBps_NonEmptyAddress_FeeParam(t *testing.T) {
	srv, getParams := captureSwapParams(t, lifiMinimalResp)
	defer srv.Close()

	provider := &LiFiProvider{
		BaseProvider: NewBaseProvider("LiFi", PriorityLiFi, lifiSupportedChains),
		client:       srv.Client(),
		baseURL:      srv.URL,
	}
	req := evmQuoteRequest()
	req.AffiliateBps = intPtr(50)
	req.AffiliateAddress = "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9"

	_, _ = provider.GetQuote(context.Background(), req)

	params := getParams()
	// 50 bps / 10000 = 0.0050
	if got := params.Get("fee"); got != "0.0050" {
		t.Errorf("expected fee=0.0050 for 50bps, got %q", got)
	}
	if got := params.Get("integrator"); got != lifiIntegratorName {
		t.Errorf("expected integrator=%q, got %q", lifiIntegratorName, got)
	}
}
