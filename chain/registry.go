package chain

import (
	"fmt"
	"sync"

	"github.com/vultisig/recipes/chain/cosmos/gaia"
	"github.com/vultisig/recipes/chain/cosmos/maya"
	"github.com/vultisig/recipes/chain/cosmos/thorchain"
	"github.com/vultisig/recipes/chain/evm"
	"github.com/vultisig/recipes/chain/solana"
	"github.com/vultisig/recipes/chain/tron"
	"github.com/vultisig/recipes/chain/utxo/bitcoin"
	"github.com/vultisig/recipes/chain/utxo/bitcoincash"
	"github.com/vultisig/recipes/chain/utxo/dash"
	"github.com/vultisig/recipes/chain/utxo/dogecoin"
	"github.com/vultisig/recipes/chain/utxo/litecoin"
	"github.com/vultisig/recipes/chain/utxo/zcash"
	"github.com/vultisig/recipes/chain/xrpl"
	"github.com/vultisig/recipes/types"
)

// Registry maintains a registry of available blockchain chains
type Registry struct {
	chains map[string]types.Chain
	mu     sync.RWMutex
}

// NewRegistry creates a new chain registry
func NewRegistry() *Registry {
	return &Registry{
		chains: make(map[string]types.Chain),
	}
}

// Register adds a chain to the registry
func (r *Registry) Register(chain types.Chain) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.chains[chain.ID()] = chain
	return nil
}

// Get retrieves a chain from the registry by ID
func (r *Registry) Get(id string) (types.Chain, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	chain, exists := r.chains[id]
	if !exists {
		return nil, fmt.Errorf("no chain registered with ID %q", id)
	}

	return chain, nil
}

// List returns all registered chains
func (r *Registry) List() []types.Chain {
	r.mu.RLock()
	defer r.mu.RUnlock()

	chains := make([]types.Chain, 0, len(r.chains))
	for _, chain := range r.chains {
		chains = append(chains, chain)
	}

	return chains
}

// DefaultRegistry is the global chain registry
var DefaultRegistry = NewRegistry()

// RegisterChain adds a chain to the default registry
func RegisterChain(chain types.Chain) {
	err := DefaultRegistry.Register(chain)
	if err != nil {
		// In production code, you might want to handle this differently
		panic(err)
	}
}

// GetChain retrieves a chain from the default registry
func GetChain(id string) (types.Chain, error) {
	return DefaultRegistry.Get(id)
}

// ExtractTxBytes extracts transaction bytes using the appropriate chain handler.
func ExtractTxBytes(chainID string, txData string) ([]byte, error) {
	c, err := GetChain(chainID)
	if err != nil {
		return nil, fmt.Errorf("unsupported chain %q: %w", chainID, err)
	}
	return c.ExtractTxBytes(txData)
}

// init registers built-in chains
func init() {
	// Register UTXO chains
	RegisterChain(bitcoin.NewChain())
	RegisterChain(bitcoincash.NewChain())
	RegisterChain(dash.NewChain())
	RegisterChain(dogecoin.NewChain())
	RegisterChain(litecoin.NewChain())
	RegisterChain(zcash.NewChain())

	// Register all EVM chains using the generic implementation
	for _, config := range evm.AllEVMChainConfigs() {
		RegisterChain(evm.NewChainFromConfig(config))
	}

	// Register other chains
	RegisterChain(thorchain.NewChain())
	RegisterChain(xrpl.NewChain())
	RegisterChain(gaia.NewChain())
	RegisterChain(maya.NewChain())
	RegisterChain(tron.NewChain())
	RegisterChain(solana.NewChain())
}
