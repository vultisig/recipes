package resolver

import (
	"testing"

	"github.com/vultisig/recipes/types"
)

func TestMagicConstantRegistry(t *testing.T) {
	// Create a new registry
	registry := NewMagicConstantRegistry()

	// Test that the registry is initialized with at least one resolver
	if len(registry.resolvers) == 0 {
		t.Error("expected registry to have at least one resolver")
	}

	// Test GetResolver with supported magic constant
	resolver, err := registry.GetResolver(types.MagicConstant_VULTISIG_TREASURY)
	if err != nil {
		t.Errorf("GetResolver() error = %v, want nil", err)
	}
	if resolver == nil {
		t.Error("GetResolver() returned nil resolver")
	}

	// Test GetResolver with unsupported magic constant
	// Using a different magic constant that treasury resolver doesn't support
	_, err = registry.GetResolver(types.MagicConstant_UNSPECIFIED)
	if err == nil {
		t.Error("GetResolver() should return error for unsupported magic constant")
	}
}

func TestMagicConstantRegistryRegister(t *testing.T) {
	registry := NewMagicConstantRegistry()
	initialCount := len(registry.resolvers)

	// Create a mock resolver for testing
	mockResolver := &mockResolver{}
	registry.Register(mockResolver)

	// Verify the resolver was added
	if len(registry.resolvers) != initialCount+1 {
		t.Errorf("Register() did not add resolver, expected %d resolvers, got %d", initialCount+1, len(registry.resolvers))
	}

	// Test that the mock resolver can be retrieved
	resolver, err := registry.GetResolver(types.MagicConstant_UNSPECIFIED)
	if err != nil {
		t.Errorf("GetResolver() error = %v, want nil", err)
	}
	if resolver != mockResolver {
		t.Error("GetResolver() returned wrong resolver")
	}
}

// mockResolver is a simple mock implementation of the Resolver interface for testing
type mockResolver struct{}

func (m *mockResolver) Supports(constant types.MagicConstant) bool {
	return constant == types.MagicConstant_UNSPECIFIED
}

func (m *mockResolver) Resolve(constant types.MagicConstant, chainID, assetID string) (string, string, error) {
	return "mock-address", "", nil
}
