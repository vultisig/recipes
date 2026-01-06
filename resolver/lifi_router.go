package resolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/vultisig/recipes/types"
)

// LiFi contract addresses per chain
// These are the official LiFi Diamond contracts
// Source: https://docs.li.fi/list-of-all-lifi-contract-addresses
var lifiRouters = map[string]string{
	"Ethereum":  "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Arbitrum":  "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Optimism":  "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Polygon":   "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"BSC":       "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Avalanche": "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Base":      "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Fantom":    "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
	"Gnosis":    "0x1231DEB6f5749EF6cE6943a275A1D3E7486F4EaE",
}

type LiFiRouterResolver struct {
	client  *http.Client
	baseURL string
}

func NewLiFiRouterResolver() Resolver {
	return &LiFiRouterResolver{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://li.quest/v1",
	}
}

func (r *LiFiRouterResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_LIFI_ROUTER
}

func (r *LiFiRouterResolver) Resolve(constant types.MagicConstant, chainID, _ string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("LiFiRouterResolver does not support type: %v", constant)
	}

	// First check our known addresses
	if router, ok := lifiRouters[chainID]; ok {
		return router, "", nil
	}

	// Fallback: fetch from LiFi chains endpoint
	router, err := r.fetchRouterFromAPI(chainID)
	if err != nil {
		return "", "", fmt.Errorf("failed to get LiFi router for chain %s: %w", chainID, err)
	}

	return router, "", nil
}

func (r *LiFiRouterResolver) fetchRouterFromAPI(chainID string) (string, error) {
	url := r.baseURL + "/chains"

	resp, err := r.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var chainsResp lifiChainsResponse
	if err := json.Unmarshal(body, &chainsResp); err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %w", err)
	}

	for _, chain := range chainsResp.Chains {
		if chain.Name == chainID || chain.Key == chainID {
			if chain.DiamondAddress != "" {
				return chain.DiamondAddress, nil
			}
		}
	}

	return "", fmt.Errorf("no LiFi router found for chain %s", chainID)
}

type lifiChainsResponse struct {
	Chains []lifiChain `json:"chains"`
}

type lifiChain struct {
	Key            string `json:"key"`
	Name           string `json:"name"`
	ChainID        int64  `json:"id"`
	DiamondAddress string `json:"diamondAddress"`
}

