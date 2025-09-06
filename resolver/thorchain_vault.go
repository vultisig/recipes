package resolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

type THORChainVaultResolver struct {
	client  *http.Client
	baseURL string
}

// InboundAddress represents a single inbound address from THORChain API
type InboundAddress struct {
	Chain   string `json:"chain"`
	PubKey  string `json:"pub_key"`
	Address string `json:"address"`
	Router  string `json:"router,omitempty"`
	Halted  bool   `json:"halted"`
	GasRate string `json:"gas_rate"`
}

// NewTHORChainVaultResolver creates a new resolver for THORChain Asgard vault addresses
func NewTHORChainVaultResolver() Resolver {
	return &THORChainVaultResolver{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://thornode.ninerealms.com",
	}
}

func (r *THORChainVaultResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_THORCHAIN_VAULT
}

// Resolve converts THORCHAIN_VAULT magic constant to current Asgard vault address
func (r *THORChainVaultResolver) Resolve(constant types.MagicConstant, chainID, assetID string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("THORChainVaultResolver does not support type: %v", constant)
	}

	// Query THORChain inbound addresses API
	inboundAddresses, err := r.getInboundAddresses()
	if err != nil {
		return "", "", fmt.Errorf("failed to get THORChain inbound addresses: %w", err)
	}

	// Find the address for the requested chain
	address, err := r.findAddressForChain(inboundAddresses, chainID)
	if err != nil {
		return "", "", fmt.Errorf("failed to find address for chain %s: %w", chainID, err)
	}

	return address, "", nil
}

// getInboundAddresses queries the THORChain API for current inbound addresses
func (r *THORChainVaultResolver) getInboundAddresses() ([]InboundAddress, error) {
	url := r.baseURL + "/thorchain/inbound_addresses"

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

	var inboundAddresses []InboundAddress
	if err := json.Unmarshal(body, &inboundAddresses); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return inboundAddresses, nil
}

// findAddressForChain finds the inbound address for a specific chain
func (r *THORChainVaultResolver) findAddressForChain(addresses []InboundAddress, chainID string) (string, error) {
	// Convert chainID string to Chain enum
	chain, err := common.FromString(chainID)
	if err != nil {
		return "", fmt.Errorf("unsupported chain: %s", err)
	}

	// Get ThorChain's native symbol for this chain
	thorchainSymbol, err := r.getThorChainSymbol(chain)
	if err != nil {
		return "", fmt.Errorf("chain %s not supported: %w", chainID, err)
	}

	// Find the address for the chain
	for _, addr := range addresses {
		if strings.ToUpper(addr.Chain) == thorchainSymbol {
			if addr.Halted {
				return "", fmt.Errorf("inbound address for chain %s is currently halted", chainID)
			}
			return addr.Address, nil
		}
	}

	return "", fmt.Errorf("no inbound address found for chain %s", chainID)
}

// getThorChainSymbol maps our Chain enum to ThorChain's expected symbols
func (r *THORChainVaultResolver) getThorChainSymbol(chain common.Chain) (string, error) {
	switch chain {
	case common.Bitcoin:
		return "BTC", nil
	case common.Ethereum:
		return "ETH", nil
	case common.Avalanche:
		return "AVAX", nil
	case common.BscChain:
		return "BSC", nil
	case common.Base:
		return "BASE", nil
	case common.Litecoin:
		return "LTC", nil
	case common.Dogecoin:
		return "DOGE", nil
	case common.BitcoinCash:
		return "BCH", nil
	default:
		return "", fmt.Errorf("chain %s not supported by ThorChain", chain.String())
	}
}
