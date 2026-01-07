package swap

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

// ERC20 approve function selector: keccak256("approve(address,uint256)")[:4]
// This is the standard ERC20 approve function signature
var erc20ApproveFunctionSelector = []byte{0x09, 0x5e, 0xa7, 0xb3}

// BuildApprovalInput contains parameters for building an ERC20 approval transaction
type BuildApprovalInput struct {
	TokenAddress string   // ERC20 token contract address
	Spender      string   // Address to approve (router contract)
	Amount       *big.Int // Exact amount to approve
	Nonce        uint64   // Transaction nonce
	GasLimit     uint64   // Gas limit (default: 60000)
	ChainID      *big.Int // EVM chain ID
}

// DefaultApprovalGasLimit is the default gas limit for ERC20 approvals
const DefaultApprovalGasLimit = uint64(60000)

// BuildApprovalTx creates an ERC20 approval transaction.
// The approval is always for the exact amount - never unlimited.
// This is a security measure to prevent unlimited approval attacks.
func BuildApprovalTx(input BuildApprovalInput) (*TxData, error) {
	if input.TokenAddress == "" {
		return nil, fmt.Errorf("token address is required")
	}
	if input.Spender == "" {
		return nil, fmt.Errorf("spender address is required")
	}
	if input.Amount == nil || input.Amount.Sign() <= 0 {
		return nil, fmt.Errorf("amount must be positive")
	}
	if input.ChainID == nil {
		return nil, fmt.Errorf("chain ID is required")
	}

	// Build the calldata
	calldata, err := EncodeApproveCalldata(input.Spender, input.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to encode approve calldata: %w", err)
	}

	gasLimit := input.GasLimit
	if gasLimit == 0 {
		gasLimit = DefaultApprovalGasLimit
	}

	return &TxData{
		To:       input.TokenAddress,
		Value:    big.NewInt(0), // ERC20 approve sends 0 ETH
		Data:     calldata,
		Nonce:    input.Nonce,
		GasLimit: gasLimit,
		ChainID:  input.ChainID,
	}, nil
}

// EncodeApproveCalldata encodes the ERC20 approve function call.
// Format: 4 bytes selector + 32 bytes address + 32 bytes amount
func EncodeApproveCalldata(spender string, amount *big.Int) ([]byte, error) {
	// Validate and normalize spender address
	spender = strings.TrimPrefix(strings.ToLower(spender), "0x")
	if len(spender) != 40 {
		return nil, fmt.Errorf("invalid spender address length: %d", len(spender))
	}

	spenderBytes, err := hex.DecodeString(spender)
	if err != nil {
		return nil, fmt.Errorf("invalid spender address: %w", err)
	}

	// Build calldata: selector (4) + spender (32) + amount (32) = 68 bytes
	calldata := make([]byte, 68)

	// Copy function selector
	copy(calldata[0:4], erc20ApproveFunctionSelector)

	// Copy spender address (right-padded to 32 bytes, address is 20 bytes)
	// Address goes in bytes 16-36 of the 32-byte slot (left-padded with zeros)
	copy(calldata[16:36], spenderBytes)

	// Copy amount (32 bytes, big-endian, left-padded with zeros)
	amountBytes := amount.Bytes()
	if len(amountBytes) > 32 {
		return nil, fmt.Errorf("amount too large")
	}
	copy(calldata[68-len(amountBytes):68], amountBytes)

	return calldata, nil
}

// DecodeApproveCalldata decodes ERC20 approve calldata into spender and amount.
// This can be used for verification purposes.
func DecodeApproveCalldata(calldata []byte) (spender string, amount *big.Int, err error) {
	if len(calldata) != 68 {
		return "", nil, fmt.Errorf("invalid calldata length: expected 68, got %d", len(calldata))
	}

	// Verify function selector
	if calldata[0] != erc20ApproveFunctionSelector[0] ||
		calldata[1] != erc20ApproveFunctionSelector[1] ||
		calldata[2] != erc20ApproveFunctionSelector[2] ||
		calldata[3] != erc20ApproveFunctionSelector[3] {
		return "", nil, fmt.Errorf("invalid function selector")
	}

	// Extract spender address (bytes 16-36 of first 32-byte slot after selector)
	spenderBytes := calldata[16:36]
	spender = "0x" + hex.EncodeToString(spenderBytes)

	// Extract amount (bytes 36-68)
	amount = new(big.Int).SetBytes(calldata[36:68])

	return spender, amount, nil
}

