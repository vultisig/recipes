package resolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/vultisig/recipes/types"
)

type MayaChainRouterResolver struct {
	client  *http.Client
	baseURL string
}

// NewMayaChainRouterResolver creates a new resolver for MayaChain router addresses
func NewMayaChainRouterResolver() Resolver {
	return &MayaChainRouterResolver{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://mayanode.mayachain.info",
	}
}

func (r *MayaChainRouterResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_MAYACHAIN_ROUTER
}

// Resolve converts MAYACHAIN_ROUTER magic constant to current router address
func (r *MayaChainRouterResolver) Resolve(constant types.MagicConstant, chainID, _ string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("MayaChainRouterResolver does not support type: %v", constant)
	}

	// Get MayaChain's symbol for this chain
	mayaSymbol, err := r.getMayaChainSymbol(chainID)
	if err != nil {
		return "", "", fmt.Errorf("chain %s not supported: %w", chainID, err)
	}

	// For non-EVM chains, router is not applicable
	if !r.isEVMChain(chainID) {
		return "", "", fmt.Errorf("MAYACHAIN_ROUTER is only available for EVM chains, %s is not an EVM chain", chainID)
	}

	// Query MayaChain inbound addresses API
	inboundAddresses, err := r.getInboundAddresses()
	if err != nil {
		return "", "", fmt.Errorf("failed to get MayaChain inbound addresses: %w", err)
	}

	// Find the router for the requested chain
	router, err := r.findRouterForChain(inboundAddresses, mayaSymbol, chainID)
	if err != nil {
		return "", "", fmt.Errorf("failed to find router for chain %s: %w", chainID, err)
	}

	return router, "", nil
}

// getInboundAddresses queries the MayaChain API for current inbound addresses
func (r *MayaChainRouterResolver) getInboundAddresses() ([]MayaInboundAddress, error) {
	url := r.baseURL + "/mayachain/inbound_addresses"

	resp, err := r.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var inboundAddresses []MayaInboundAddress
	if err := json.Unmarshal(body, &inboundAddresses); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return inboundAddresses, nil
}

// findRouterForChain finds the router address for a specific chain
func (r *MayaChainRouterResolver) findRouterForChain(addresses []MayaInboundAddress, mayaSymbol, chainID string) (string, error) {
	for _, addr := range addresses {
		if strings.ToUpper(addr.Chain) == mayaSymbol {
			if addr.Halted {
				return "", fmt.Errorf("inbound address for chain %s is currently halted", chainID)
			}

			if addr.Router == "" {
				return "", fmt.Errorf("no router address available for chain %s", chainID)
			}

			return addr.Router, nil
		}
	}

	return "", fmt.Errorf("no inbound address found for chain %s", chainID)
}

// getMayaChainSymbol maps chain IDs to MayaChain's expected symbols
// MayaChain router is used for:
// - Arbitrum: all swaps to/from ARB
// - Ethereum: swaps to Maya-only chains (ZEC, DASH) that THORChain doesn't support
func (r *MayaChainRouterResolver) getMayaChainSymbol(chainID string) (string, error) {
	switch strings.ToLower(chainID) {
	case "arbitrum":
		return "ARB", nil
	case "ethereum":
		return "ETH", nil
	default:
		return "", fmt.Errorf("chain %s not supported by MayaChain router", chainID)
	}
}

// isEVMChain returns true if the chain is an EVM-compatible chain supported by MayaChain
func (r *MayaChainRouterResolver) isEVMChain(chainID string) bool {
	switch strings.ToLower(chainID) {
	case "arbitrum", "ethereum":
		return true
	default:
		return false
	}
}

