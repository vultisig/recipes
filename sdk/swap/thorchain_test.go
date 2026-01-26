package swap

import (
	"math/big"
	"testing"
)

func TestEncodeRouterDeposit_NativeETH(t *testing.T) {
	// Create a THORChain provider
	provider := NewTHORChainProvider(nil)

	// Create a mock quote for native ETH swap
	quote := &Quote{
		Provider: "THORChain",
		FromAsset: Asset{
			Chain:    "Ethereum",
			Symbol:   "ETH",
			Address:  "", // Native ETH - empty address
			Decimals: 18,
		},
		ToAsset: Asset{
			Chain:    "Base",
			Symbol:   "ETH",
			Address:  "",
			Decimals: 18,
		},
		FromAmount:     big.NewInt(2000000000000000), // 0.002 ETH
		ExpectedOutput: big.NewInt(1900000000000000),
		InboundAddress: "0x9318E3C2456E7B1eC557Fc3706D5aA5B6A452512",
		Router:         "0xD37BbE5744D730a1d98d8DC97c42F0Ca46aD7146",
		Memo:           "=:BASE.ETH:0xd27FcacfC7fA6a3B7674Bf3D38f863C89Cc6F6Ea:149818/3/0",
		Expiry:         1706100000,
	}

	req := SwapRequest{
		Quote:       quote,
		Sender:      "0xd27FcacfC7fA6a3B7674Bf3D38f863C89Cc6F6Ea",
		Destination: "0xd27FcacfC7fA6a3B7674Bf3D38f863C89Cc6F6Ea",
	}

	// Call encodeRouterDeposit
	calldata, value, err := provider.encodeRouterDeposit(req)
	if err != nil {
		t.Fatalf("encodeRouterDeposit failed: %v", err)
	}

	t.Logf("Calldata length: %d", len(calldata))
	t.Logf("Calldata hex: %x", calldata)
	t.Logf("Value: %s", value.String())

	// For native ETH:
	// - Value should be the swap amount (2000000000000000)
	// - Calldata amount parameter should be 0
	if value.Cmp(big.NewInt(2000000000000000)) != 0 {
		t.Errorf("Expected value=2000000000000000, got %s", value.String())
	}

	// Check that the amount in calldata is 0
	// The amount is at bytes 68-100 (third parameter, 0-indexed)
	// Function selector: 4 bytes
	// vault address: 32 bytes (offset 4-36)
	// asset address: 32 bytes (offset 36-68)
	// amount: 32 bytes (offset 68-100)
	if len(calldata) < 100 {
		t.Fatalf("Calldata too short: %d bytes", len(calldata))
	}

	amountBytes := calldata[68:100]
	amountInCalldata := new(big.Int).SetBytes(amountBytes)

	t.Logf("Amount in calldata: %s", amountInCalldata.String())

	if amountInCalldata.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected amount in calldata to be 0 for native ETH, got %s", amountInCalldata.String())
	}
}

func TestEncodeRouterDeposit_ERC20Token(t *testing.T) {
	// Create a THORChain provider
	provider := NewTHORChainProvider(nil)

	// Create a mock quote for USDC (ERC20) swap
	quote := &Quote{
		Provider: "THORChain",
		FromAsset: Asset{
			Chain:    "Ethereum",
			Symbol:   "USDC",
			Address:  "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // USDC address
			Decimals: 6,
		},
		ToAsset: Asset{
			Chain:    "Base",
			Symbol:   "ETH",
			Address:  "",
			Decimals: 18,
		},
		FromAmount:     big.NewInt(100000000), // 100 USDC
		ExpectedOutput: big.NewInt(50000000000000000),
		InboundAddress: "0x9318E3C2456E7B1eC557Fc3706D5aA5B6A452512",
		Router:         "0xD37BbE5744D730a1d98d8DC97c42F0Ca46aD7146",
		Memo:           "=:BASE.ETH:0xd27FcacfC7fA6a3B7674Bf3D38f863C89Cc6F6Ea:149818/3/0",
		Expiry:         1706100000,
	}

	req := SwapRequest{
		Quote:       quote,
		Sender:      "0xd27FcacfC7fA6a3B7674Bf3D38f863C89Cc6F6Ea",
		Destination: "0xd27FcacfC7fA6a3B7674Bf3D38f863C89Cc6F6Ea",
	}

	// Call encodeRouterDeposit
	calldata, value, err := provider.encodeRouterDeposit(req)
	if err != nil {
		t.Fatalf("encodeRouterDeposit failed: %v", err)
	}

	t.Logf("Calldata length: %d", len(calldata))
	t.Logf("Value: %s", value.String())

	// For ERC20:
	// - Value should be 0 (no ETH sent)
	// - Calldata amount parameter should be the token amount
	if value.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Expected value=0 for ERC20, got %s", value.String())
	}

	// Check that the amount in calldata is the token amount
	amountBytes := calldata[68:100]
	amountInCalldata := new(big.Int).SetBytes(amountBytes)

	t.Logf("Amount in calldata: %s", amountInCalldata.String())

	if amountInCalldata.Cmp(big.NewInt(100000000)) != 0 {
		t.Errorf("Expected amount in calldata to be 100000000 for ERC20, got %s", amountInCalldata.String())
	}
}
