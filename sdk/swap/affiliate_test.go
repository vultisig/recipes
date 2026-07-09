package swap

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// stationAffiliateAddress / stationAffiliateBps mirror the values Station
// (money.terra.station) uses on the TS stack (agent-backend-ts
// STATION_AFFILIATE_CONFIG.native): affiliateFeeAddress='stvs', 50 bps
// (baseAffiliateBps in vultisig-sdk/packages/core/chain/swap/affiliate/config.ts).
// Test-local constants — callers select the real values, this package only
// threads whatever it's given.
const (
	stationAffiliateAddress = "stvs"
	stationAffiliateBps     = 50
	defaultAffiliateAddress = "v0"
	defaultAffiliateBps     = 50
)

// newThorQuoteStub starts an httptest server that echoes the request's query
// params back as the memo, so tests can assert on exactly what the provider
// sent without depending on real THORChain/Maya endpoints.
func newThorQuoteStub(t *testing.T) (*httptest.Server, *url.Values) {
	t.Helper()
	var captured url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.URL.Query()
		resp := thorChainQuoteResponse{
			InboundAddress:    "thor1inboundvault",
			Router:            "",
			Expiry:            1234567890,
			Memo:              "=:ETH.ETH:0xdest:0",
			ExpectedAmountOut: "100000000",
		}
		if a := captured.Get("affiliate"); a != "" {
			resp.Memo = "=:ETH.ETH:0xdest:0:" + a + ":" + captured.Get("affiliate_bps")
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	t.Cleanup(srv.Close)
	return srv, &captured
}

func newMayaQuoteStub(t *testing.T) (*httptest.Server, *url.Values) {
	t.Helper()
	var captured url.Values
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.URL.Query()
		resp := mayaChainQuoteResponse{
			InboundAddress:    "maya1inboundvault",
			Router:            "",
			Expiry:            1234567890,
			Memo:              "=:ETH.ETH:0xdest:0",
			ExpectedAmountOut: "100000000",
		}
		if a := captured.Get("affiliate"); a != "" {
			resp.Memo = "=:ETH.ETH:0xdest:0:" + a + ":" + captured.Get("affiliate_bps")
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	t.Cleanup(srv.Close)
	return srv, &captured
}

func ethAsset() Asset   { return Asset{Chain: "Ethereum", Symbol: "ETH", Decimals: 18} }
func rune_() Asset      { return Asset{Chain: "THORChain", Symbol: "RUNE", Decimals: 8} }
func cacaoAsset() Asset { return Asset{Chain: "MayaChain", Symbol: "CACAO", Decimals: 10} }

// (a) Station-app THOR swap: quote carries the Station affiliate (stvs, 50bps).
func TestTHORChainGetQuote_StationAffiliate(t *testing.T) {
	srv, captured := newThorQuoteStub(t)
	provider := NewTHORChainProvider([]string{srv.URL})

	bps := stationAffiliateBps
	req := QuoteRequest{
		From:             rune_(),
		To:               ethAsset(),
		Amount:           big.NewInt(1_000_000_000),
		Destination:      "0xdest",
		AffiliateAddress: stationAffiliateAddress,
		AffiliateBps:     &bps,
	}

	quote, err := provider.GetQuote(t.Context(), req)
	if err != nil {
		t.Fatalf("GetQuote failed: %v", err)
	}

	if got := captured.Get("affiliate"); got != stationAffiliateAddress {
		t.Errorf("affiliate param: got %q, want %q", got, stationAffiliateAddress)
	}
	if got := captured.Get("affiliate_bps"); got != "50" {
		t.Errorf("affiliate_bps param: got %q, want %q", got, "50")
	}
	if quote.Memo == "" {
		t.Fatal("expected non-empty memo")
	}
}

// (b) non-station/vultisig default THOR swap: carries the default affiliate (v0, 50bps).
func TestTHORChainGetQuote_DefaultAffiliate(t *testing.T) {
	srv, captured := newThorQuoteStub(t)
	provider := NewTHORChainProvider([]string{srv.URL})

	bps := defaultAffiliateBps
	req := QuoteRequest{
		From:             ethAsset(),
		To:               rune_(),
		Amount:           big.NewInt(1_000_000_000_000_000_000),
		Destination:      "thor1dest",
		AffiliateAddress: defaultAffiliateAddress,
		AffiliateBps:     &bps,
	}

	_, err := provider.GetQuote(t.Context(), req)
	if err != nil {
		t.Fatalf("GetQuote failed: %v", err)
	}

	if got := captured.Get("affiliate"); got != defaultAffiliateAddress {
		t.Errorf("affiliate param: got %q, want %q", got, defaultAffiliateAddress)
	}
	if got := captured.Get("affiliate_bps"); got != "50" {
		t.Errorf("affiliate_bps param: got %q, want %q", got, "50")
	}
}

// Fail-closed: no affiliate fields set on the request → no affiliate params
// sent at all (bare memo), never a silent default.
func TestTHORChainGetQuote_NoAffiliateWhenUnset(t *testing.T) {
	srv, captured := newThorQuoteStub(t)
	provider := NewTHORChainProvider([]string{srv.URL})

	req := QuoteRequest{
		From:        rune_(),
		To:          ethAsset(),
		Amount:      big.NewInt(1_000_000_000),
		Destination: "0xdest",
	}

	_, err := provider.GetQuote(t.Context(), req)
	if err != nil {
		t.Fatalf("GetQuote failed: %v", err)
	}

	if captured.Has("affiliate") {
		t.Errorf("expected no affiliate param when unset, got %q", captured.Get("affiliate"))
	}
	if captured.Has("affiliate_bps") {
		t.Errorf("expected no affiliate_bps param when unset, got %q", captured.Get("affiliate_bps"))
	}
}

// Mayachain mirrors the THORChain behaviour — same wiring gap, same fix.
func TestMayachainGetQuote_StationAffiliate(t *testing.T) {
	srv, captured := newMayaQuoteStub(t)
	provider := NewMayachainProvider([]string{srv.URL})

	bps := stationAffiliateBps
	req := QuoteRequest{
		From:             cacaoAsset(),
		To:               ethAsset(),
		Amount:           big.NewInt(10_000_000_000_000),
		Destination:      "0xdest",
		AffiliateAddress: stationAffiliateAddress,
		AffiliateBps:     &bps,
	}

	_, err := provider.GetQuote(t.Context(), req)
	if err != nil {
		t.Fatalf("GetQuote failed: %v", err)
	}

	if got := captured.Get("affiliate"); got != stationAffiliateAddress {
		t.Errorf("affiliate param: got %q, want %q", got, stationAffiliateAddress)
	}
	if got := captured.Get("affiliate_bps"); got != "50" {
		t.Errorf("affiliate_bps param: got %q, want %q", got, "50")
	}
}

func TestMayachainGetQuote_NoAffiliateWhenUnset(t *testing.T) {
	srv, captured := newMayaQuoteStub(t)
	provider := NewMayachainProvider([]string{srv.URL})

	req := QuoteRequest{
		From:        cacaoAsset(),
		To:          ethAsset(),
		Amount:      big.NewInt(10_000_000_000_000),
		Destination: "0xdest",
	}

	_, err := provider.GetQuote(t.Context(), req)
	if err != nil {
		t.Fatalf("GetQuote failed: %v", err)
	}

	if captured.Has("affiliate") {
		t.Errorf("expected no affiliate param when unset, got %q", captured.Get("affiliate"))
	}
}

// (c) the memo format matches what THORChain expects:
// =:ASSET:ADDR:LIM:AFFILIATE:BPS
func TestSetAffiliateParams_MemoFormat(t *testing.T) {
	params := url.Values{}
	bps := stationAffiliateBps
	req := QuoteRequest{AffiliateAddress: stationAffiliateAddress, AffiliateBps: &bps}
	setAffiliateParams(params, req)

	// Simulate the THOR-style memo the inbound vault/thornode would build
	// once affiliate+affiliate_bps are present on the quote request.
	memo := "=:ETH.ETH:0xdest:149818:" + params.Get("affiliate") + ":" + params.Get("affiliate_bps")
	want := "=:ETH.ETH:0xdest:149818:stvs:50"
	if memo != want {
		t.Errorf("memo format: got %q, want %q", memo, want)
	}
}

// Explicit zero bps is treated the same as unset (no affiliate line) — a
// caller passing AffiliateBps=0 is signalling "no fee", not "0% but still
// attribute to me".
func TestSetAffiliateParams_ZeroBpsOmitted(t *testing.T) {
	params := url.Values{}
	zero := 0
	req := QuoteRequest{AffiliateAddress: stationAffiliateAddress, AffiliateBps: &zero}
	setAffiliateParams(params, req)

	if params.Has("affiliate") || params.Has("affiliate_bps") {
		t.Errorf("expected no affiliate params for zero bps, got affiliate=%q affiliate_bps=%q",
			params.Get("affiliate"), params.Get("affiliate_bps"))
	}
}

// Out-of-range bps (>10000 = >100%) is rejected client-side rather than
// sent to thornode/maya, mirroring the toleranceBps bound-check.
func TestSetAffiliateParams_OutOfRangeBpsOmitted(t *testing.T) {
	params := url.Values{}
	tooHigh := 10001
	req := QuoteRequest{AffiliateAddress: stationAffiliateAddress, AffiliateBps: &tooHigh}
	setAffiliateParams(params, req)

	if params.Has("affiliate") || params.Has("affiliate_bps") {
		t.Errorf("expected no affiliate params for out-of-range bps, got affiliate=%q affiliate_bps=%q",
			params.Get("affiliate"), params.Get("affiliate_bps"))
	}
}

// AffiliateAddress set but AffiliateBps nil (or vice versa) — fail closed,
// no partial affiliate line.
func TestSetAffiliateParams_PartialFieldsOmitted(t *testing.T) {
	params := url.Values{}
	req := QuoteRequest{AffiliateAddress: stationAffiliateAddress, AffiliateBps: nil}
	setAffiliateParams(params, req)
	if params.Has("affiliate") {
		t.Errorf("expected no affiliate param when bps is nil, got %q", params.Get("affiliate"))
	}

	params2 := url.Values{}
	bps := stationAffiliateBps
	req2 := QuoteRequest{AffiliateAddress: "", AffiliateBps: &bps}
	setAffiliateParams(params2, req2)
	if params2.Has("affiliate_bps") {
		t.Errorf("expected no affiliate_bps param when address is empty, got %q", params2.Get("affiliate_bps"))
	}
}
