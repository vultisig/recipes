package protocol

import (
	"fmt"
	"sync"

	"github.com/vultisig/recipes/bitcoin"
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

// buildKey creates a unique key from chainID and protocolID
func (r *Registry) buildKey(chainID, protocolID string) string {
	return fmt.Sprintf("%s.%s", chainID, protocolID)
}

// Register adds a protocol to the registry
func (r *Registry) Register(protocol types.Protocol) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := r.buildKey(protocol.ChainID(), protocol.ID())
	if _, exists := r.protocols[key]; exists {
		return fmt.Errorf("protocol with Key %q already registered", key)
	}

	r.protocols[key] = protocol
	return nil
}

// Get retrieves a protocol from the registry by ID
func (r *Registry) Get(chainID, protocolID string) (types.Protocol, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := r.buildKey(chainID, protocolID)
	protocol, exists := r.protocols[key]
	if !exists {
		return nil, fmt.Errorf("no protocol registered with Key %q", key)
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
	// Register BTC protocol
	RegisterProtocol(bitcoin.NewBTC())
}
