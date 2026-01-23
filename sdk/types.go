package sdk

// DeriveOptions contains options for deriving signing hashes.
type DeriveOptions struct {
	// SignBytes is required for Cosmos chains where signBytes cannot be derived from txBytes.
	// For other chains, this field is ignored.
	SignBytes []byte
}

// DerivedHash represents a derived signing hash and its lookup key.
type DerivedHash struct {
	// Message is the actual bytes to sign (will be base64-encoded for KeysignMessage.Message)
	Message []byte
	// Hash is the lookup key (will be base64-encoded for KeysignMessage.Hash)
	Hash []byte
}
