package resolver

// Resolver defines the interface for magic constant resolution
type Resolver interface {
	// Resolve converts a magic constant to an actual address + memo
	Resolve(chainID, assetID string) (string, string, error)
}
