package aavev3

import (
	"math/big"
	"testing"
)

func TestRayToAPY(t *testing.T) {
	tests := []struct {
		name string
		ray  *big.Int
		want float64
	}{
		{"zero", big.NewInt(0), 0},
		{"3.5%", new(big.Int).Mul(big.NewInt(35), new(big.Int).Exp(big.NewInt(10), big.NewInt(24), nil)), 3.5},
		{"100% = 1e27", new(big.Int).Exp(big.NewInt(10), big.NewInt(27), nil), 100.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RayToAPY(tt.ray)
			if got != tt.want {
				t.Errorf("RayToAPY = %f, want %f", got, tt.want)
			}
		})
	}
}

func TestParseAmount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		decimals int
		want     *big.Int
		wantErr  bool
	}{
		{"integer", "100", 6, big.NewInt(100_000_000), false},
		{"decimal", "0.5", 6, big.NewInt(500_000), false},
		{"fractional with more precision", "100.123456", 6, big.NewInt(100_123_456), false},
		{"max", "max", 18, MaxUint256, false},
		{"MAX case insensitive", "MAX", 6, MaxUint256, false},
		{"18 decimals", "1.5", 18, new(big.Int).Mul(big.NewInt(15), new(big.Int).Exp(big.NewInt(10), big.NewInt(17), nil)), false},
		{"truncates excess precision", "1.1234567", 6, big.NewInt(1_123_456), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAmount(tt.input, tt.decimals)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.Cmp(tt.want) != 0 {
				t.Errorf("ParseAmount(%q, %d) = %s, want %s", tt.input, tt.decimals, got, tt.want)
			}
		})
	}
}
