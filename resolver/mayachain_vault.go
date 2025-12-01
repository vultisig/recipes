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

type MayaChainVaultResolver struct {
	client  *http.Client
	baseURL string
}

// MayaInboundAddress represents a single inbound address from MayaChain API
type MayaInboundAddress struct {
	Chain   string `json:"chain"`
	PubKey  string `json:"pub_key"`
	Address string `json:"address"`
	Router  string `json:"router,omitempty"`
	Halted  bool   `json:"halted"`
	GasRate string `json:"gas_rate"`
}

// NewMayaChainVaultResolver creates a new resolver for MayaChain Asgard vault addresses
func NewMayaChainVaultResolver() Resolver {
	return &MayaChainVaultResolver{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://mayanode.mayachain.info",
	}
}

func (r *MayaChainVaultResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_MAYACHAIN_VAULT
}

// Resolve converts MAYACHAIN_VAULT magic constant to current Asgard vault address
func (r *MayaChainVaultResolver) Resolve(constant types.MagicConstant, chainID, assetID string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("MayaChainVaultResolver does not support type: %v", constant)
	}

	// Query MayaChain inbound addresses API
	inboundAddresses, err := r.getInboundAddresses()
	if err != nil {
		return "", "", fmt.Errorf("failed to get MayaChain inbound addresses: %w", err)
	}

	// Find the address for the requested chain
	address, err := r.findAddressForChain(inboundAddresses, chainID)
	if err != nil {
		return "", "", fmt.Errorf("failed to find address for chain %s: %w", chainID, err)
	}

	return address, "", nil
}

// getInboundAddresses queries the MayaChain API for current inbound addresses
func (r *MayaChainVaultResolver) getInboundAddresses() ([]MayaInboundAddress, error) {
	url := r.baseURL + "/mayachain/inbound_addresses"

	resp, err := r.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

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

// findAddressForChain finds the inbound address for a specific chain
func (r *MayaChainVaultResolver) findAddressForChain(addresses []MayaInboundAddress, chainID string) (string, error) {
	// Get MayaChain's symbol for this chain
	mayaSymbol, err := r.getMayaChainSymbol(chainID)
	if err != nil {
		return "", fmt.Errorf("chain %s not supported: %w", chainID, err)
	}

	// Find the address for the chain
	for _, addr := range addresses {
		if strings.ToUpper(addr.Chain) == mayaSymbol {
			if addr.Halted {
				return "", fmt.Errorf("inbound address for chain %s is currently halted", chainID)
			}

			// For chains with router (like EVM), use router address
			if addr.Router != "" {
				return addr.Router, nil
			}
			return addr.Address, nil
		}
	}

	return "", fmt.Errorf("no inbound address found for chain %s", chainID)
}

// getMayaChainSymbol maps chain IDs to MayaChain's expected symbols
func (r *MayaChainVaultResolver) getMayaChainSymbol(chainID string) (string, error) {
	// Normalize chainID to lowercase for comparison
	switch strings.ToLower(chainID) {
	case "zcash":
		return "ZEC", nil
	case "bitcoin":
		return "BTC", nil
	case "ethereum":
		return "ETH", nil
	case "dash":
		return "DASH", nil
	case "thorchain":
		return "THOR", nil
	case "kujira":
		return "KUJI", nil
	case "arbitrum":
		return "ARB", nil
	default:
		return "", fmt.Errorf("chain %s not supported by MayaChain", chainID)
	}
}

