package chain

import (
	"fmt"
	"sync"
)

// Registry maintains a registry of available blockchain chains
type Registry struct {
	chains map[string]Chain
	mu     sync.RWMutex
}

// NewRegistry creates a new chain registry
func NewRegistry() *Registry {
	return &Registry{
		chains: make(map[string]Chain),
	}
}

// Register adds a chain to the registry
func (r *Registry) Register(chain Chain) error {
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
func (r *Registry) Get(id string) (Chain, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	chain, exists := r.chains[id]
	if !exists {
		return nil, fmt.Errorf("no chain registered with ID %q", id)
	}

	return chain, nil
}

// List returns all registered chains
func (r *Registry) List() []Chain {
	r.mu.RLock()
	defer r.mu.RUnlock()

	chains := make([]Chain, 0, len(r.chains))
	for _, chain := range r.chains {
		chains = append(chains, chain)
	}

	return chains
}

// DefaultRegistry is the global chain registry
var DefaultRegistry = NewRegistry()

// RegisterChain adds a chain to the default registry
func RegisterChain(chain Chain) {
	err := DefaultRegistry.Register(chain)
	if err != nil {
		// In production code, you might want to handle this differently
		panic(err)
	}
}

// GetChain retrieves a chain from the default registry
func GetChain(id string) (Chain, error) {
	return DefaultRegistry.Get(id)
}

// init registers built-in chains
func init() {
	// Register Bitcoin chain
	RegisterChain(NewBitcoin())
}
