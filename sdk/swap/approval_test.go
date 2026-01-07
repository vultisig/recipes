package swap

import (
	"math/big"
	"testing"
)

func TestEncodeApproveCalldata(t *testing.T) {
	tests := []struct {
		name    string
		spender string
		amount  *big.Int
		wantErr bool
	}{
		{
			name:    "valid approval - small amount",
			spender: "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD",
			amount:  big.NewInt(1000000), // 1 USDC
			wantErr: false,
		},
		{
			name:    "valid approval - large amount",
			spender: "0x1111111254EEB25477B68fb85Ed929f73A960582",
			amount:  new(big.Int).Mul(big.NewInt(1000000), big.NewInt(1e18)), // 1M tokens
			wantErr: false,
		},
		{
			name:    "valid approval - with 0x prefix",
			spender: "0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F",
			amount:  big.NewInt(1e18),
			wantErr: false,
		},
		{
			name:    "valid approval - without 0x prefix",
			spender: "d9e1cE17f2641f24aE83637ab66a2cca9C378B9F",
			amount:  big.NewInt(1e18),
			wantErr: false,
		},
		{
			name:    "invalid spender - too short",
			spender: "0x1234",
			amount:  big.NewInt(1000),
			wantErr: true,
		},
		{
			name:    "invalid spender - too long",
			spender: "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD00",
			amount:  big.NewInt(1000),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calldata, err := EncodeApproveCalldata(tt.spender, tt.amount)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Verify calldata length
			if len(calldata) != 68 {
				t.Errorf("expected calldata length 68, got %d", len(calldata))
			}

			// Verify function selector (0x095ea7b3)
			expectedSelector := []byte{0x09, 0x5e, 0xa7, 0xb3}
			for i := 0; i < 4; i++ {
				if calldata[i] != expectedSelector[i] {
					t.Errorf("invalid function selector byte %d: expected %x, got %x", i, expectedSelector[i], calldata[i])
				}
			}
		})
	}
}

func TestDecodeApproveCalldata(t *testing.T) {
	// Test roundtrip encoding/decoding
	originalSpender := "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"
	originalAmount := big.NewInt(1000000000000000000) // 1e18

	calldata, err := EncodeApproveCalldata(originalSpender, originalAmount)
	if err != nil {
		t.Fatalf("failed to encode: %v", err)
	}

	spender, amount, err := DecodeApproveCalldata(calldata)
	if err != nil {
		t.Fatalf("failed to decode: %v", err)
	}

	// Compare spender (normalize to lowercase)
	if spender != "0x3fc91a3afd70395cd496c647d5a6cc9d4b2b7fad" {
		t.Errorf("spender mismatch: expected %s, got %s", "0x3fc91a3afd70395cd496c647d5a6cc9d4b2b7fad", spender)
	}

	// Compare amount
	if amount.Cmp(originalAmount) != 0 {
		t.Errorf("amount mismatch: expected %s, got %s", originalAmount.String(), amount.String())
	}
}

func TestDecodeApproveCalldata_Invalid(t *testing.T) {
	tests := []struct {
		name     string
		calldata []byte
	}{
		{
			name:     "too short",
			calldata: make([]byte, 67),
		},
		{
			name:     "too long",
			calldata: make([]byte, 69),
		},
		{
			name:     "wrong selector",
			calldata: append([]byte{0x00, 0x00, 0x00, 0x00}, make([]byte, 64)...),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := DecodeApproveCalldata(tt.calldata)
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	}
}

func TestBuildApprovalTx(t *testing.T) {
	tests := []struct {
		name    string
		input   BuildApprovalInput
		wantErr bool
	}{
		{
			name: "valid approval tx",
			input: BuildApprovalInput{
				TokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // USDC
				Spender:      "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD", // Uniswap Router
				Amount:       big.NewInt(1000000),                          // 1 USDC
				Nonce:        5,
				GasLimit:     60000,
				ChainID:      big.NewInt(1), // Ethereum mainnet
			},
			wantErr: false,
		},
		{
			name: "missing token address",
			input: BuildApprovalInput{
				Spender: "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD",
				Amount:  big.NewInt(1000000),
				ChainID: big.NewInt(1),
			},
			wantErr: true,
		},
		{
			name: "missing spender",
			input: BuildApprovalInput{
				TokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Amount:       big.NewInt(1000000),
				ChainID:      big.NewInt(1),
			},
			wantErr: true,
		},
		{
			name: "zero amount",
			input: BuildApprovalInput{
				TokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Spender:      "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD",
				Amount:       big.NewInt(0),
				ChainID:      big.NewInt(1),
			},
			wantErr: true,
		},
		{
			name: "missing chain ID",
			input: BuildApprovalInput{
				TokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Spender:      "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD",
				Amount:       big.NewInt(1000000),
			},
			wantErr: true,
		},
		{
			name: "default gas limit applied",
			input: BuildApprovalInput{
				TokenAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
				Spender:      "0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD",
				Amount:       big.NewInt(1000000),
				ChainID:      big.NewInt(1),
				// GasLimit not set - should use default
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, err := BuildApprovalTx(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Verify transaction data
			if tx.To != tt.input.TokenAddress {
				t.Errorf("To address mismatch: expected %s, got %s", tt.input.TokenAddress, tx.To)
			}

			if tx.Value.Sign() != 0 {
				t.Errorf("expected zero value, got %s", tx.Value.String())
			}

			if len(tx.Data) != 68 {
				t.Errorf("expected calldata length 68, got %d", len(tx.Data))
			}

			if tx.Nonce != tt.input.Nonce {
				t.Errorf("nonce mismatch: expected %d, got %d", tt.input.Nonce, tx.Nonce)
			}

			expectedGas := tt.input.GasLimit
			if expectedGas == 0 {
				expectedGas = DefaultApprovalGasLimit
			}
			if tx.GasLimit != expectedGas {
				t.Errorf("gas limit mismatch: expected %d, got %d", expectedGas, tx.GasLimit)
			}

			if tx.ChainID.Cmp(tt.input.ChainID) != 0 {
				t.Errorf("chain ID mismatch: expected %s, got %s", tt.input.ChainID.String(), tx.ChainID.String())
			}
		})
	}
}

func TestIsApprovalRequired(t *testing.T) {
	tests := []struct {
		name     string
		asset    Asset
		expected bool
	}{
		{
			name: "Native ETH - no approval",
			asset: Asset{
				Chain:   "Ethereum",
				Symbol:  "ETH",
				Address: "",
			},
			expected: false,
		},
		{
			name: "ERC20 USDC on Ethereum - needs approval",
			asset: Asset{
				Chain:   "Ethereum",
				Symbol:  "USDC",
				Address: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			},
			expected: true,
		},
		{
			name: "Native BTC - no approval",
			asset: Asset{
				Chain:   "Bitcoin",
				Symbol:  "BTC",
				Address: "",
			},
			expected: false,
		},
		{
			name: "Native SOL - no approval (non-EVM)",
			asset: Asset{
				Chain:   "Solana",
				Symbol:  "SOL",
				Address: "",
			},
			expected: false,
		},
		{
			name: "SPL token on Solana - no approval (non-EVM)",
			asset: Asset{
				Chain:   "Solana",
				Symbol:  "USDC",
				Address: "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			},
			expected: false,
		},
		{
			name: "ERC20 on Arbitrum - needs approval",
			asset: Asset{
				Chain:   "Arbitrum",
				Symbol:  "USDC",
				Address: "0xaf88d065e77c8cC2239327C5EDb3A432268e5831",
			},
			expected: true,
		},
		{
			name: "Native BNB on BSC - no approval",
			asset: Asset{
				Chain:   "BSC",
				Symbol:  "BNB",
				Address: "",
			},
			expected: false,
		},
		{
			name: "BEP20 on BSC - needs approval",
			asset: Asset{
				Chain:   "BSC",
				Symbol:  "USDT",
				Address: "0x55d398326f99059fF775485246999027B3197955",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsApprovalRequired(tt.asset)
			if result != tt.expected {
				t.Errorf("IsApprovalRequired(%v) = %v, want %v", tt.asset, result, tt.expected)
			}
		})
	}
}

func TestIsEVMChain(t *testing.T) {
	evmChains := []string{"Ethereum", "BSC", "Polygon", "Arbitrum", "Avalanche", "Base", "Optimism"}
	nonEVMChains := []string{"Bitcoin", "Solana", "THORChain", "Cosmos", "XRP", "Litecoin"}

	for _, chain := range evmChains {
		if !IsEVMChain(chain) {
			t.Errorf("expected %s to be an EVM chain", chain)
		}
	}

	for _, chain := range nonEVMChains {
		if IsEVMChain(chain) {
			t.Errorf("expected %s to NOT be an EVM chain", chain)
		}
	}
}

