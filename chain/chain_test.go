package chain

import (
	"testing"
)

// TestBitcoinChain tests the Bitcoin chain implementation
func TestBitcoinChain(t *testing.T) {
	btc := NewBitcoin()

	// Test ID
	if got := btc.ID(); got != "bitcoin" {
		t.Errorf("Bitcoin.ID() = %v, want %v", got, "bitcoin")
	}

	// Test Name
	if got := btc.Name(); got != "Bitcoin" {
		t.Errorf("Bitcoin.Name() = %v, want %v", got, "Bitcoin")
	}

	// Test SupportedProtocols
	protocols := btc.SupportedProtocols()
	if len(protocols) != 1 {
		t.Errorf("Bitcoin.SupportedProtocols() length = %v, want %v", len(protocols), 1)
	}
	if protocols[0] != "btc" {
		t.Errorf("Bitcoin.SupportedProtocols()[0] = %v, want %v", protocols[0], "btc")
	}
}

// TestRegistry tests the chain registry functionality
func TestRegistry(t *testing.T) {
	// Create a new registry
	reg := NewRegistry()

	// Test registering a chain
	btc := NewBitcoin()
	err := reg.Register(btc)
	if err != nil {
		t.Errorf("Register() error = %v, want nil", err)
	}

	// Test registering duplicate chain
	err = reg.Register(btc)
	if err == nil {
		t.Error("Register() duplicate chain error = nil, want error")
	}

	// Test getting a chain
	chain, err := reg.Get("bitcoin")
	if err != nil {
		t.Errorf("Get() error = %v, want nil", err)
	}
	if chain.ID() != "bitcoin" {
		t.Errorf("Get() chain ID = %v, want %v", chain.ID(), "bitcoin")
	}

	// Test getting non-existent chain
	_, err = reg.Get("nonexistent")
	if err == nil {
		t.Error("Get() non-existent chain error = nil, want error")
	}

	// Test listing chains
	chains := reg.List()
	if len(chains) != 1 {
		t.Errorf("List() length = %v, want %v", len(chains), 1)
	}
	if chains[0].ID() != "bitcoin" {
		t.Errorf("List()[0].ID() = %v, want %v", chains[0].ID(), "bitcoin")
	}
}

// TestDefaultRegistry tests the default registry functionality
func TestDefaultRegistry(t *testing.T) {
	// Test getting a chain from default registry
	chain, err := GetChain("bitcoin")
	if err != nil {
		t.Errorf("GetChain() error = %v, want nil", err)
	}
	if chain.ID() != "bitcoin" {
		t.Errorf("GetChain() chain ID = %v, want %v", chain.ID(), "bitcoin")
	}

	// Test getting non-existent chain from default registry
	_, err = GetChain("nonexistent")
	if err == nil {
		t.Error("GetChain() non-existent chain error = nil, want error")
	}
}
