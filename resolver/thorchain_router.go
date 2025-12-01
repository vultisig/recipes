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

type THORChainRouterResolver struct {
	client  *http.Client
	baseURL string
}

func NewTHORChainRouterResolver() Resolver {
	return &THORChainRouterResolver{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://thornode.ninerealms.com",
	}
}

func (r *THORChainRouterResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_THORCHAIN_ROUTER
}

func (r *THORChainRouterResolver) Resolve(constant types.MagicConstant, chainID, _ string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("THORChainRouterResolver does not support type: %v", constant)
	}

	chain, err := common.FromString(chainID)
	if err != nil {
		return "", "", fmt.Errorf("unsupported chain: %w", err)
	}

	if !chain.IsEvm() {
		return "", "", fmt.Errorf("THORCHAIN_ROUTER is only available for EVM chains, %s is not an EVM chain", chainID)
	}

	inboundAddresses, err := r.getInboundAddresses()
	if err != nil {
		return "", "", fmt.Errorf("failed to get THORChain inbound addresses: %w", err)
	}

	router, err := r.findRouterForChain(inboundAddresses, chainID)
	if err != nil {
		return "", "", fmt.Errorf("failed to find router for chain %s: %w", chainID, err)
	}

	return router, "", nil
}

func (r *THORChainRouterResolver) getInboundAddresses() ([]InboundAddress, error) {
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
	err = json.Unmarshal(body, &inboundAddresses)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return inboundAddresses, nil
}

func (r *THORChainRouterResolver) findRouterForChain(addresses []InboundAddress, chainID string) (string, error) {
	chain, err := common.FromString(chainID)
	if err != nil {
		return "", fmt.Errorf("unsupported chain: %w", err)
	}

	thorchainSymbol, err := getThorChainSymbol(chain)
	if err != nil {
		return "", fmt.Errorf("chain %s not supported: %w", chainID, err)
	}

	for _, addr := range addresses {
		if strings.ToUpper(addr.Chain) == thorchainSymbol {
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
