package chain

// Bitcoin implements the Chain interface for the Bitcoin blockchain
type Bitcoin struct{}

// ID returns the unique identifier for the Bitcoin chain
func (b *Bitcoin) ID() string {
	return "bitcoin"
}

// Name returns a human-readable name for the Bitcoin chain
func (b *Bitcoin) Name() string {
	return "Bitcoin"
}

// SupportedProtocols returns the list of protocol IDs supported by the Bitcoin chain
func (b *Bitcoin) SupportedProtocols() []string {
	return []string{"btc"}
}

// NewBitcoin creates a new Bitcoin chain instance
func NewBitcoin() Chain {
	return &Bitcoin{}
}
