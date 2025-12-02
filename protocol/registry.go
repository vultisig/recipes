package protocol

import (
	"fmt"
	"sync"

	"github.com/vultisig/recipes/chain/utxo/bitcoin"
	"github.com/vultisig/recipes/chain/utxo/bitcoincash"
	"github.com/vultisig/recipes/chain/utxo/dogecoin"
	"github.com/vultisig/recipes/chain/utxo/litecoin"
	"github.com/vultisig/recipes/chain/utxo/zcash"
	"github.com/vultisig/recipes/types"
)

// Registry maintains a registry of available protocols
type Registry struct {
	protocols map[string]types.Protocol
	mu        sync.RWMutex
}

// NewRegistry creates a new protocol registry
func NewRegistry() *Registry {
	return &Registry{
		protocols: make(map[string]types.Protocol),
	}
}

// Register adds a protocol to the registry
func (r *Registry) Register(protocol types.Protocol) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := protocol.ID()
	if _, exists := r.protocols[id]; exists {
		return fmt.Errorf("protocol with ID %q already registered", id)
	}

	r.protocols[id] = protocol
	return nil
}

// Get retrieves a protocol from the registry by ID
func (r *Registry) Get(id string) (types.Protocol, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	protocol, exists := r.protocols[id]
	if !exists {
		return nil, fmt.Errorf("no protocol registered with ID %q", id)
	}

	return protocol, nil
}

// List returns all registered protocols
func (r *Registry) List() []types.Protocol {
	r.mu.RLock()
	defer r.mu.RUnlock()

	protocols := make([]types.Protocol, 0, len(r.protocols))
	for _, protocol := range r.protocols {
		protocols = append(protocols, protocol)
	}

	return protocols
}

// ListByChain returns all protocols for a specific chain
func (r *Registry) ListByChain(chainID string) []types.Protocol {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var chainProtocols []types.Protocol
	for _, protocol := range r.protocols {
		if protocol.ChainID() == chainID {
			chainProtocols = append(chainProtocols, protocol)
		}
	}

	return chainProtocols
}

// DefaultRegistry is the global protocol registry
var DefaultRegistry = NewRegistry()

// RegisterProtocol adds a protocol to the default registry
func RegisterProtocol(protocol types.Protocol) {
	err := DefaultRegistry.Register(protocol)
	if err != nil {
		// In production code, you might want to handle this differently
		panic(err)
	}
}

// GetProtocol retrieves a protocol from the default registry
func GetProtocol(id string) (types.Protocol, error) {
	return DefaultRegistry.Get(id)
}

// ListProtocols returns all registered protocols
func ListProtocols() []types.Protocol {
	return DefaultRegistry.List()
}

// ListProtocolsByChain returns all protocols for a specific chain
func ListProtocolsByChain(chainID string) []types.Protocol {
	return DefaultRegistry.ListByChain(chainID)
}

// init registers built-in protocols
func init() {
	// Register UTXO protocols
	RegisterProtocol(bitcoin.NewBTC())
	RegisterProtocol(bitcoincash.NewBCH())
	RegisterProtocol(dogecoin.NewDOGE())
	RegisterProtocol(litecoin.NewLTC())
	RegisterProtocol(zcash.NewZEC())
}
