package types

// Chain represents a blockchain network
type Chain interface {
	// ID returns the unique identifier for the chain
	ID() string

	// Name returns a human-readable name for the chain
	Name() string

	// SupportedProtocols returns the list of protocol IDs supported by this chain
	SupportedProtocols() []string
}
