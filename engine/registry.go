package engine

import (
	"fmt"

	"github.com/vultisig/recipes/engine/btc"
	"github.com/vultisig/recipes/engine/evm"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

// ChainEngine defines the interface that all chain-specific engines must implement
type ChainEngine interface {
	Evaluate(rule *types.Rule, txBytes []byte) error
	Supports(chain common.Chain) bool
}

// ChainEngineRegistry manages chain-specific engines
type ChainEngineRegistry struct {
	engines []ChainEngine
}

// NewChainEngineRegistry creates a new engine registry with all engines registered
func NewChainEngineRegistry() *ChainEngineRegistry {
	registry := &ChainEngineRegistry{
		engines: make([]ChainEngine, 0),
	}

	// Register non-EVM engines (EVM engines are created on-demand in GetEngine)

	// Register BTC engine
	registry.Register(&btc.Btc{})

	return registry
}

// Register adds an engine to the registry
func (r *ChainEngineRegistry) Register(engine ChainEngine) {
	r.engines = append(r.engines, engine)
}

// GetEngine returns the appropriate engine for the given chain.
// EVM engines are created on-demand to avoid unnecessary ABI loading.
func (r *ChainEngineRegistry) GetEngine(chain common.Chain) (ChainEngine, error) {
	// Handle EVM chains with on-demand instantiation
	if chain.IsEvm() {
		nativeSymbol, err := chain.NativeSymbol()
		if err != nil {
			return nil, fmt.Errorf("failed to get native symbol for %s: %w", chain.String(), err)
		}

		evmEngine, err := evm.NewEvm(nativeSymbol)
		if err != nil {
			return nil, fmt.Errorf("failed to create EVM engine for %s: %w", chain.String(), err)
		}

		return evmEngine, nil
	}

	// For non-EVM chains, use the pre-registered engines
	for _, engine := range r.engines {
		if engine.Supports(chain) {
			return engine, nil
		}
	}

	return nil, fmt.Errorf("no engine found for chain: %v", chain)
}
