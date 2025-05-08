package policy

import (
	"strings"
)

// ResourcePath represents a parsed resource path
type ResourcePath struct {
	ChainID    string
	ProtocolID string
	FunctionID string
	Full       string
}

// ParseResourcePath parses a resource path into its components
func ParseResourcePath(path string) *ResourcePath {
	parts := strings.Split(path, ".")

	// Default values
	rp := &ResourcePath{
		Full: path,
	}

	// Extract components based on available parts
	if len(parts) >= 1 {
		rp.ChainID = parts[0]
	}

	if len(parts) >= 2 {
		rp.ProtocolID = parts[1]
	}

	if len(parts) >= 3 {
		rp.FunctionID = parts[2]
	}

	return rp
}

// IsWildcard checks if the resource path contains a wildcard
func (rp *ResourcePath) IsWildcard() bool {
	return strings.Contains(rp.Full, "*")
}

// Matches checks if this resource path matches a target path
func (rp *ResourcePath) Matches(targetPath string) bool {
	// If resource is a wildcard, check if target starts with the prefix
	if rp.IsWildcard() {
		prefix := strings.TrimSuffix(rp.Full, "*")
		return strings.HasPrefix(targetPath, prefix)
	}

	// Otherwise, exact match required
	return rp.Full == targetPath
}

// String returns the full resource path
func (rp *ResourcePath) String() string {
	return rp.Full
}
