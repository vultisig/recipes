package base

import (
	"github.com/vultisig/recipes/chain/evm/ethereum"
	"github.com/vultisig/recipes/types"
)

// NewBaseETH creates a new Base ETH protocol using the shared native EVM implementation
func NewBaseETH() types.Protocol {
	return ethereum.NewNativeEVMProtocol(
		"base",
		"eth",
		"Base ETH",
		"Native Ether on the Base L2 blockchain",
	)
}
