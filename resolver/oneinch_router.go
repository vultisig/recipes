package resolver

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// 1inch Aggregation Router V6 addresses per chain
// Source: https://docs.1inch.io/docs/aggregation-protocol/smart-contract-addresses/
var oneInchRouters = map[string]string{
	"Ethereum":  "0x111111125421cA6dc452d289314280a0f8842A65",
	"BSC":       "0x111111125421cA6dc452d289314280a0f8842A65",
	"Polygon":   "0x111111125421cA6dc452d289314280a0f8842A65",
	"Avalanche": "0x111111125421cA6dc452d289314280a0f8842A65",
	"Arbitrum":  "0x111111125421cA6dc452d289314280a0f8842A65",
	"Optimism":  "0x111111125421cA6dc452d289314280a0f8842A65",
	"Base":      "0x111111125421cA6dc452d289314280a0f8842A65",
	"Fantom":    "0x111111125421cA6dc452d289314280a0f8842A65",
	"Gnosis":    "0x111111125421cA6dc452d289314280a0f8842A65",
}

type OneInchRouterResolver struct{}

func NewOneInchRouterResolver() Resolver {
	return &OneInchRouterResolver{}
}

func (r *OneInchRouterResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_ONEINCH_ROUTER
}

func (r *OneInchRouterResolver) Resolve(constant types.MagicConstant, chainID, _ string) (string, string, error) {
	if !r.Supports(constant) {
		return "", "", fmt.Errorf("OneInchRouterResolver does not support type: %v", constant)
	}

	router, ok := oneInchRouters[chainID]
	if !ok {
		return "", "", fmt.Errorf("1inch not supported on chain %s", chainID)
	}

	return router, "", nil
}

