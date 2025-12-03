package engine

import (
	"fmt"

	"github.com/vultisig/recipes/engine/evm"
	"github.com/vultisig/recipes/engine/solana"
	"github.com/vultisig/recipes/engine/thorchain"
	"github.com/vultisig/recipes/engine/utxo/bitcoin"
	"github.com/vultisig/recipes/engine/utxo/bitcoincash"
	"github.com/vultisig/recipes/engine/utxo/dogecoin"
	"github.com/vultisig/recipes/engine/utxo/litecoin"
	"github.com/vultisig/recipes/engine/utxo/zcash"
	"github.com/vultisig/recipes/engine/xrpl"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

// ChainEngine defines the interface that all chain-specific engines must implement
type ChainEngine interface {
	Evaluate(rule *types.Rule, txBytes []byte) error
	Supports(chain common.Chain) bool
}

// SupportedEVMChains is the list of all EVM chains that have engines registered.
// To add a new EVM chain, simply add it to this list and ensure it has a valid
// NativeSymbol() implementation in the vultisig-go/common package.
var SupportedEVMChains = []common.Chain{
	common.Ethereum, common.BscChain, common.Arbitrum, common.Avalanche,
	common.Base, common.Blast, common.CronosChain, common.Optimism,
	common.Polygon, common.Zksync,
}

// ChainEngineRegistry manages chain-specific engines
type ChainEngineRegistry struct {
	engines []ChainEngine
}

// NewChainEngineRegistry creates a new engine registry with all engines registered
func NewChainEngineRegistry() (*ChainEngineRegistry, error) {
	registry := &ChainEngineRegistry{
		engines: make([]ChainEngine, 0),
	}

	// Register EVM engines - one for each supported EVM chain
	for _, chain := range SupportedEVMChains {
		nativeSymbol, err := chain.NativeSymbol()
		if err != nil {
			return nil, fmt.Errorf("failed to get native symbol for %s: %w", chain.String(), err)
		}

		evmEngine, err := evm.NewEvm(nativeSymbol)
		if err != nil {
			return nil, fmt.Errorf("failed to create evm engine for %s: %w", chain.String(), err)
		}

		registry.Register(evmEngine)
	}

	// Register Bitcoin engine
	registry.Register(bitcoin.NewBtc())

	// Register Bitcoin Cash engine
	registry.Register(bitcoincash.NewBitcoinCash())

	// Register Dogecoin engine
	registry.Register(dogecoin.NewDogecoin())

	// Register Litecoin engine
	registry.Register(litecoin.NewLitecoin())

	// Register XRPL engine
	registry.Register(xrpl.NewXRPL())

	// Register Thorchain engine
	registry.Register(thorchain.NewThorchain())

	// Register Zcash engine
	registry.Register(zcash.NewZcash())

	// Register Solana engine
	solEng, err := solana.NewSolana()
	if err != nil {
		return nil, fmt.Errorf("failed to create solana engine: %s", err)
	}
	registry.Register(solEng)

	return registry, nil
}

// Register adds an engine to the registry
func (r *ChainEngineRegistry) Register(engine ChainEngine) {
	r.engines = append(r.engines, engine)
}

// GetEngine returns the first engine that supports the given chain
func (r *ChainEngineRegistry) GetEngine(chain common.Chain) (ChainEngine, error) {
	for _, engine := range r.engines {
		if engine.Supports(chain) {
			return engine, nil
		}
	}
	return nil, fmt.Errorf("no engine found for chain: %v", chain)
}
