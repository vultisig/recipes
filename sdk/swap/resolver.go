package swap

import (
	"fmt"
	"sync"

	"github.com/vultisig/recipes/resolver"
	"github.com/vultisig/recipes/types"
)

var (
	registry     *resolver.MagicConstantRegistry
	registryOnce sync.Once
)

// getRegistry returns the shared magic constant registry
func getRegistry() *resolver.MagicConstantRegistry {
	registryOnce.Do(func() {
		registry = resolver.NewMagicConstantRegistry()
	})
	return registry
}

// RouterInfo contains resolved router information
type RouterInfo struct {
	Address string
	Chain   string
}

// ResolveLiFiRouter resolves the LiFi router address for a chain
func ResolveLiFiRouter(chain string) (*RouterInfo, error) {
	reg := getRegistry()
	r, err := reg.GetResolver(types.MagicConstant_LIFI_ROUTER)
	if err != nil {
		return nil, fmt.Errorf("no LiFi router resolver: %w", err)
	}

	address, _, err := r.Resolve(types.MagicConstant_LIFI_ROUTER, chain, "")
	if err != nil {
		return nil, fmt.Errorf("failed to resolve LiFi router for %s: %w", chain, err)
	}

	return &RouterInfo{
		Address: address,
		Chain:   chain,
	}, nil
}

// ResolveOneInchRouter resolves the 1inch router address for a chain
func ResolveOneInchRouter(chain string) (*RouterInfo, error) {
	reg := getRegistry()
	r, err := reg.GetResolver(types.MagicConstant_ONEINCH_ROUTER)
	if err != nil {
		return nil, fmt.Errorf("no 1inch router resolver: %w", err)
	}

	address, _, err := r.Resolve(types.MagicConstant_ONEINCH_ROUTER, chain, "")
	if err != nil {
		return nil, fmt.Errorf("failed to resolve 1inch router for %s: %w", chain, err)
	}

	return &RouterInfo{
		Address: address,
		Chain:   chain,
	}, nil
}

// ResolveUniswapRouter resolves the Uniswap Universal Router address for a chain
func ResolveUniswapRouter(chain string) (*RouterInfo, error) {
	reg := getRegistry()
	r, err := reg.GetResolver(types.MagicConstant_UNISWAP_UNIVERSAL_ROUTER)
	if err != nil {
		return nil, fmt.Errorf("no Uniswap router resolver: %w", err)
	}

	address, _, err := r.Resolve(types.MagicConstant_UNISWAP_UNIVERSAL_ROUTER, chain, "")
	if err != nil {
		return nil, fmt.Errorf("failed to resolve Uniswap router for %s: %w", chain, err)
	}

	return &RouterInfo{
		Address: address,
		Chain:   chain,
	}, nil
}

// ResolveTHORChainRouter resolves the THORChain router address for an EVM chain
func ResolveTHORChainRouter(chain string) (*RouterInfo, error) {
	reg := getRegistry()
	r, err := reg.GetResolver(types.MagicConstant_THORCHAIN_ROUTER)
	if err != nil {
		return nil, fmt.Errorf("no THORChain router resolver: %w", err)
	}

	address, _, err := r.Resolve(types.MagicConstant_THORCHAIN_ROUTER, chain, "")
	if err != nil {
		return nil, fmt.Errorf("failed to resolve THORChain router for %s: %w", chain, err)
	}

	return &RouterInfo{
		Address: address,
		Chain:   chain,
	}, nil
}

// ResolveMayaChainRouter resolves the Mayachain router address for an EVM chain
func ResolveMayaChainRouter(chain string) (*RouterInfo, error) {
	reg := getRegistry()
	r, err := reg.GetResolver(types.MagicConstant_MAYACHAIN_ROUTER)
	if err != nil {
		return nil, fmt.Errorf("no Mayachain router resolver: %w", err)
	}

	address, _, err := r.Resolve(types.MagicConstant_MAYACHAIN_ROUTER, chain, "")
	if err != nil {
		return nil, fmt.Errorf("failed to resolve Mayachain router for %s: %w", chain, err)
	}

	return &RouterInfo{
		Address: address,
		Chain:   chain,
	}, nil
}

// GetApprovalSpender returns the spender address that needs token approval for a provider
func GetApprovalSpender(provider, chain string) (string, error) {
	switch provider {
	case "THORChain":
		info, err := ResolveTHORChainRouter(chain)
		if err != nil {
			return "", err
		}
		return info.Address, nil
	case "Mayachain":
		info, err := ResolveMayaChainRouter(chain)
		if err != nil {
			return "", err
		}
		return info.Address, nil
	case "LiFi":
		info, err := ResolveLiFiRouter(chain)
		if err != nil {
			return "", err
		}
		return info.Address, nil
	case "1inch":
		info, err := ResolveOneInchRouter(chain)
		if err != nil {
			return "", err
		}
		return info.Address, nil
	case "Uniswap":
		info, err := ResolveUniswapRouter(chain)
		if err != nil {
			return "", err
		}
		return info.Address, nil
	default:
		return "", fmt.Errorf("unknown provider: %s", provider)
	}
}

