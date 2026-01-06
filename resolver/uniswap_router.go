package resolver

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// Uniswap Universal Router addresses per chain
// Source: https://docs.uniswap.org/contracts/v3/reference/deployments
var uniswapUniversalRouters = map[string]string{
	"Ethereum": "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD",
	"Polygon":  "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD",
	"Arbitrum": "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD",
	"Optimism": "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD",
	"Base":     "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD",
}

// Uniswap SwapRouter02 addresses (fallback for chains without Universal Router)
var uniswapSwapRouter02 = map[string]string{
	"Ethereum": "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45",
	"Polygon":  "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45",
	"Arbitrum": "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45",
	"Optimism": "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45",
}

type UniswapRouterResolver struct{}

func NewUniswapRouterResolver() Resolver {
	return &UniswapRouterResolver{}
}

func (r *UniswapRouterResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_UNISWAP_UNIVERSAL_ROUTER
}

func (r *UniswapRouterResolver) Resolve(constant types.MagicConstant, chainID, _ string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("UniswapRouterResolver does not support type: %v", constant)
	}

	// Try Universal Router first
	if router, ok := uniswapUniversalRouters[chainID]; ok {
		return router, "", nil
	}

	// Fallback to SwapRouter02
	if router, ok := uniswapSwapRouter02[chainID]; ok {
		return router, "", nil
	}

	return "", "", fmt.Errorf("Uniswap not supported on chain %s", chainID)
}

