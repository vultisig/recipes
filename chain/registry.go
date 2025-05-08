package chain

import (
	"fmt"
	"sync"

	"github.com/vultisig/recipes/bitcoin"
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

	id := chain.ID()
	if _, exists := r.chains[id]; exists {
		return fmt.Errorf("chain with ID %q already registered", id)
	}

	r.chains[id] = chain
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

// init registers built-in chains
func init() {
	// Register Bitcoin chain
	RegisterChain(bitcoin.NewChain())
}
