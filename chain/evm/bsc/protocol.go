package bsc

import (
	"github.com/vultisig/recipes/chain/evm/ethereum"
	"github.com/vultisig/recipes/types"
)

// NewBNB creates a new BSC BNB protocol using the shared native EVM implementation
func NewBNB() types.Protocol {
	return ethereum.NewNativeEVMProtocol(
		"bsc",
		"bnb",
		"BSC BNB",
		"Native BNB on the BNB Smart Chain",
	)
}
