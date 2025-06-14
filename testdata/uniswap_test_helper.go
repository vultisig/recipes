package testdata

import (
	"math/big"
	"strings"
	"time"

	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

// Sample Uniswap V2 Router ABI (just the functions we need for testing)
const uniswapV2RouterABI = `[
	{
		"inputs": [
			{"internalType": "uint256", "name": "amountOutMin", "type": "uint256"},
			{"internalType": "address[]", "name": "path", "type": "address[]"},
			{"internalType": "address", "name": "to", "type": "address"},
			{"internalType": "uint256", "name": "deadline", "type": "uint256"}
		],
		"name": "swapExactETHForTokens",
		"outputs": [{"internalType": "uint256[]", "name": "amounts", "type": "uint256[]"}],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{"internalType": "uint256", "name": "amountIn", "type": "uint256"},
			{"internalType": "uint256", "name": "amountOutMin", "type": "uint256"},
			{"internalType": "address[]", "name": "path", "type": "address[]"},
			{"internalType": "address", "name": "to", "type": "address"},
			{"internalType": "uint256", "name": "deadline", "type": "uint256"}
		],
		"name": "swapExactTokensForETH",
		"outputs": [{"internalType": "uint256[]", "name": "amounts", "type": "uint256[]"}],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{"internalType": "address", "name": "tokenA", "type": "address"},
			{"internalType": "address", "name": "tokenB", "type": "address"},
			{"internalType": "uint256", "name": "amountADesired", "type": "uint256"},
			{"internalType": "uint256", "name": "amountBDesired", "type": "uint256"},
			{"internalType": "uint256", "name": "amountAMin", "type": "uint256"},
			{"internalType": "uint256", "name": "amountBMin", "type": "uint256"},
			{"internalType": "address", "name": "to", "type": "address"},
			{"internalType": "uint256", "name": "deadline", "type": "uint256"}
		],
		"name": "addLiquidity",
		"outputs": [
			{"internalType": "uint256", "name": "amountA", "type": "uint256"},
			{"internalType": "uint256", "name": "amountB", "type": "uint256"},
			{"internalType": "uint256", "name": "liquidity", "type": "uint256"}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{"internalType": "address", "name": "tokenA", "type": "address"},
			{"internalType": "address", "name": "tokenB", "type": "address"},
			{"internalType": "uint256", "name": "liquidity", "type": "uint256"},
			{"internalType": "uint256", "name": "amountAMin", "type": "uint256"},
			{"internalType": "uint256", "name": "amountBMin", "type": "uint256"},
			{"internalType": "address", "name": "to", "type": "address"},
			{"internalType": "uint256", "name": "deadline", "type": "uint256"}
		],
		"name": "removeLiquidity",
		"outputs": [
			{"internalType": "uint256", "name": "amountA", "type": "uint256"},
			{"internalType": "uint256", "name": "amountB", "type": "uint256"}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

// CreateSampleUniswapTransaction creates a sample swapExactETHForTokens transaction
func CreateSampleUniswapTransaction() (*types.Transaction, error) {
	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(uniswapV2RouterABI))
	if err != nil {
		return nil, err
	}

	// Prepare function arguments
	amountOutMin := big.NewInt(1000000000000000000) // 1 ETH worth of tokens minimum
	path := []common.Address{
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), // WETH
		common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"), // Sample token
	}
	to := common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045") // Whitelisted recipient
	// Keep the tx valid for ~30 minutes - suitable for plugin operations
	deadline := big.NewInt(time.Now().Unix() + 1800)

	// Encode the function call
	data, err := parsedABI.Pack("swapExactETHForTokens", amountOutMin, path, to, deadline)
	if err != nil {
		return nil, err
	}

	// Create transaction
	tx := types.NewTransaction(
		0, // nonce
		common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"), // Uniswap V2 Router
		big.NewInt(1000000000000000000),                                   // 1 ETH value
		300000,                                                            // gas limit
		big.NewInt(20000000000),                                           // gas price (20 gwei)
		data,
	)

	return tx, nil
}

// Sample transaction parameters for testing
var SampleUniswapParams = map[string]interface{}{
	"amountOutMin": big.NewInt(1000000000000000000), // 1 ETH worth
	"path": []interface{}{
		"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // WETH
		"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984", // Sample token
	},
	"to":       "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
	"deadline": big.NewInt(time.Now().Unix() + 1800), // Dynamic future deadline
}

// Invalid parameters that should fail validation
var InvalidUniswapParams = map[string]interface{}{
	"amountOutMin": big.NewInt(0), // Invalid: zero slippage protection
	"path": []interface{}{
		"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2", // WETH
		"0x0000000000000000000000000000000000000000", // Invalid: zero address
	},
	"to":       "0x0000000000000000000000000000000000000000", // Invalid: zero address
	"deadline": big.NewInt(1600000000),                       // Invalid: past deadline
}

// Build a dynamic-fee swapExactETHForTokens transaction and return hex string
func BuildSwapExactETHForTokensHex(recipient string) (string, error) {
	parsedABI, err := abi.JSON(strings.NewReader(uniswapV2RouterABI))
	if err != nil {
		return "", err
	}

	amountOutMin := big.NewInt(1000000000000000000)
	path := []common.Address{
		common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"),
	}
	to := common.HexToAddress(recipient)
	deadline := big.NewInt(time.Now().Unix() + 1800)

	data, err := parsedABI.Pack("swapExactETHForTokens", amountOutMin, path, to, deadline)
	if err != nil {
		return "", err
	}

	router := common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")

	raw, err := buildUnsignedDynamicTx(&router, big.NewInt(1000000000000000000), data)
	if err != nil {
		return "", err
	}

	return "0x" + hex.EncodeToString(raw), nil
}

// Helper wrappers for tests
func ValidSwapExactETHForTokensTxHex() string {
	hex, _ := BuildSwapExactETHForTokensHex("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
	return hex
}

func InvalidRecipientSwapExactETHForTokensTxHex() string {
	hex, _ := BuildSwapExactETHForTokensHex("0x000000000000000000000000000000000000dead")
	return hex
}

// shared addresses
var (
	wethAddr    = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	sampleAddr  = common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")
	routerAddr  = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	whitelistTo = common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")
)

func buildUnsignedDynamicTx(to *common.Address, value *big.Int, data []byte) ([]byte, error) {
	unsigned := struct {
		ChainID    *big.Int
		Nonce      uint64
		GasTipCap  *big.Int
		GasFeeCap  *big.Int
		Gas        uint64
		To         *common.Address `rlp:"nil"`
		Value      *big.Int
		Data       []byte
		AccessList types.AccessList
	}{
		ChainID:    big.NewInt(1),
		Nonce:      0,
		GasTipCap:  big.NewInt(2_000_000_000),  // 2 gwei
		GasFeeCap:  big.NewInt(20_000_000_000), // 20 gwei
		Gas:        300_000,
		To:         to,
		Value:      value,
		Data:       data,
		AccessList: nil,
	}
	payload, err := rlp.EncodeToBytes(unsigned)
	if err != nil {
		return nil, err
	}
	return append([]byte{types.DynamicFeeTxType}, payload...), nil
}

// Build tx hex helper generic
func buildTxHex(data []byte, ethValue *big.Int) (string, error) {
	raw, err := buildUnsignedDynamicTx(&routerAddr, ethValue, data)
	if err != nil {
		return "", err
	}
	return "0x" + hex.EncodeToString(raw), nil
}

// ------- swapExactTokensForETH ---------
func BuildSwapExactTokensForETHTxHex(recipient common.Address, amountIn *big.Int) (string, error) {
	parsedABI, err := abi.JSON(strings.NewReader(uniswapV2RouterABI))
	if err != nil {
		return "", err
	}
	amountOutMin := big.NewInt(1000000000000000) // 0.001 ETH
	path := []common.Address{sampleAddr, wethAddr}
	deadline := big.NewInt(time.Now().Unix() + 1800)
	data, err := parsedABI.Pack("swapExactTokensForETH", amountIn, amountOutMin, path, recipient, deadline)
	if err != nil {
		return "", err
	}
	return buildTxHex(data, big.NewInt(0))
}

func bigInt(dec string) *big.Int { v, _ := new(big.Int).SetString(dec, 10); return v }

func ValidSwapExactTokensForETHTxHex() string {
	hexStr, _ := BuildSwapExactTokensForETHTxHex(whitelistTo, bigInt("100000000000000000000"))
	return hexStr
}

func ExceedAmountSwapExactTokensForETHTxHex() string {
	hexStr, _ := BuildSwapExactTokensForETHTxHex(whitelistTo, bigInt("600000000000000000000"))
	return hexStr
}

// ------- addLiquidity ---------
func BuildAddLiquidityTxHex(recipient common.Address) (string, error) {
	parsedABI, err := abi.JSON(strings.NewReader(uniswapV2RouterABI))
	if err != nil {
		return "", err
	}
	amtDesired := big.NewInt(1_000_000_000_000_000_000) // 1 token/ETH
	amtMin := big.NewInt(900_000_000_000_000_000)       // 0.9
	deadline := big.NewInt(time.Now().Unix() + 1800)
	data, err := parsedABI.Pack("addLiquidity", wethAddr, sampleAddr, amtDesired, amtDesired, amtMin, amtMin, recipient, deadline)
	if err != nil {
		return "", err
	}
	return buildTxHex(data, big.NewInt(0))
}

func ValidAddLiquidityTxHex() string {
	hexStr, _ := BuildAddLiquidityTxHex(whitelistTo)
	return hexStr
}

func InvalidTokenAddLiquidityTxHex() string {
	wrongToken := common.HexToAddress("0x000000000000000000000000000000000000dead")
	parsedABI, err := abi.JSON(strings.NewReader(uniswapV2RouterABI))
	if err != nil {
		return "" // Return empty string if ABI parsing fails
	}
	amt := big.NewInt(1_000_000_000_000_000_000)
	deadline := big.NewInt(time.Now().Unix() + 1800)
	data, err := parsedABI.Pack("addLiquidity", wrongToken, sampleAddr, amt, amt, amt, amt, whitelistTo, deadline)
	if err != nil {
		return "" // Return empty string if packing fails
	}
	hexStr, err := buildTxHex(data, big.NewInt(0))
	if err != nil {
		return "" // Return empty string if tx building fails
	}
	return hexStr
}

// ------- removeLiquidity ---------
func BuildRemoveLiquidityTxHex(recipient common.Address) (string, error) {
	parsedABI, err := abi.JSON(strings.NewReader(uniswapV2RouterABI))
	if err != nil {
		return "", err
	}
	liquidity := big.NewInt(1_000_000_000_000_000_000)
	amtMin := big.NewInt(1)
	deadline := big.NewInt(time.Now().Unix() + 1800)
	data, err := parsedABI.Pack("removeLiquidity", wethAddr, sampleAddr, liquidity, amtMin, amtMin, recipient, deadline)
	if err != nil {
		return "", err
	}
	return buildTxHex(data, big.NewInt(0))
}

func ValidRemoveLiquidityTxHex() string {
	hexStr, _ := BuildRemoveLiquidityTxHex(whitelistTo)
	return hexStr
}

func InvalidRecipientRemoveLiquidityTxHex() string {
	bad := common.HexToAddress("0x000000000000000000000000000000000000beef")
	hexStr, _ := BuildRemoveLiquidityTxHex(bad)
	return hexStr
}
