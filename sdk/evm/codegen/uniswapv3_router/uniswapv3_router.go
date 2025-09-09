// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv3_router

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// IApproveAndCallIncreaseLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IApproveAndCallIncreaseLiquidityParams struct {
	Token0     common.Address
	Token1     common.Address
	TokenId    *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
}

// IApproveAndCallMintParams is an auto generated low-level Go binding around an user-defined struct.
type IApproveAndCallMintParams struct {
	Token0     common.Address
	Token1     common.Address
	Fee        *big.Int
	TickLower  *big.Int
	TickUpper  *big.Int
	Amount0Min *big.Int
	Amount1Min *big.Int
	Recipient  common.Address
}

// IV3SwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// IV3SwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// IV3SwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// IV3SwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type IV3SwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// Uniswapv3RouterMetaData contains all meta data concerning the Uniswapv3Router contract.
var Uniswapv3RouterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factoryV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factoryV3\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_positionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveMax\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveMaxMinusOne\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveZeroThenMax\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"approveZeroThenMaxMinusOne\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callPositionManager\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"paths\",\"type\":\"bytes[]\"},{\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"},{\"internalType\":\"uint24\",\"name\":\"maximumTickDivergence\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"secondsAgo\",\"type\":\"uint32\"}],\"name\":\"checkOracleSlippage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint24\",\"name\":\"maximumTickDivergence\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"secondsAgo\",\"type\":\"uint32\"}],\"name\":\"checkOracleSlippage\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structIV3SwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIV3SwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structIV3SwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIV3SwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getApprovalType\",\"outputs\":[{\"internalType\":\"enumIApproveAndCall.ApprovalType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"}],\"internalType\":\"structIApproveAndCall.IncreaseLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"increaseLiquidity\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIApproveAndCall.MintParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"previousBlockhash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"positionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"wrapETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "Uniswapv3Router",
}

// Uniswapv3Router is an auto generated Go binding around an Ethereum contract.
type Uniswapv3Router struct {
	abi abi.ABI
}

// NewUniswapv3Router creates a new instance of Uniswapv3Router.
func NewUniswapv3Router() *Uniswapv3Router {
	parsed, err := Uniswapv3RouterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Uniswapv3Router{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Uniswapv3Router) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _factoryV2, address factoryV3, address _positionManager, address _WETH9) returns()
func (uniswapv3Router *Uniswapv3Router) PackConstructor(_factoryV2 common.Address, factoryV3 common.Address, _positionManager common.Address, _WETH9 common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("", _factoryV2, factoryV3, _positionManager, _WETH9)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWETH9 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4aa4a4fc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function WETH9() view returns(address)
func (uniswapv3Router *Uniswapv3Router) PackWETH9() []byte {
	enc, err := uniswapv3Router.abi.Pack("WETH9")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWETH9 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4aa4a4fc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function WETH9() view returns(address)
func (uniswapv3Router *Uniswapv3Router) TryPackWETH9() ([]byte, error) {
	return uniswapv3Router.abi.Pack("WETH9")
}

// UnpackWETH9 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (uniswapv3Router *Uniswapv3Router) UnpackWETH9(data []byte) (common.Address, error) {
	out, err := uniswapv3Router.abi.Unpack("WETH9", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackApproveMax is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x571ac8b0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approveMax(address token) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackApproveMax(token common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("approveMax", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApproveMax is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x571ac8b0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approveMax(address token) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackApproveMax(token common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("approveMax", token)
}

// PackApproveMaxMinusOne is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcab372ce.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackApproveMaxMinusOne(token common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("approveMaxMinusOne", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApproveMaxMinusOne is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcab372ce.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approveMaxMinusOne(address token) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackApproveMaxMinusOne(token common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("approveMaxMinusOne", token)
}

// PackApproveZeroThenMax is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x639d71a9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackApproveZeroThenMax(token common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("approveZeroThenMax", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApproveZeroThenMax is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x639d71a9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approveZeroThenMax(address token) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackApproveZeroThenMax(token common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("approveZeroThenMax", token)
}

// PackApproveZeroThenMaxMinusOne is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab3fdd50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackApproveZeroThenMaxMinusOne(token common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("approveZeroThenMaxMinusOne", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApproveZeroThenMaxMinusOne is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab3fdd50.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approveZeroThenMaxMinusOne(address token) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackApproveZeroThenMaxMinusOne(token common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("approveZeroThenMaxMinusOne", token)
}

// PackCallPositionManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb3a2af13.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (uniswapv3Router *Uniswapv3Router) PackCallPositionManager(data []byte) []byte {
	enc, err := uniswapv3Router.abi.Pack("callPositionManager", data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCallPositionManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb3a2af13.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (uniswapv3Router *Uniswapv3Router) TryPackCallPositionManager(data []byte) ([]byte, error) {
	return uniswapv3Router.abi.Pack("callPositionManager", data)
}

// UnpackCallPositionManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb3a2af13.
//
// Solidity: function callPositionManager(bytes data) payable returns(bytes result)
func (uniswapv3Router *Uniswapv3Router) UnpackCallPositionManager(data []byte) ([]byte, error) {
	out, err := uniswapv3Router.abi.Unpack("callPositionManager", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackCheckOracleSlippage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xefdeed8e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (uniswapv3Router *Uniswapv3Router) PackCheckOracleSlippage(paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) []byte {
	enc, err := uniswapv3Router.abi.Pack("checkOracleSlippage", paths, amounts, maximumTickDivergence, secondsAgo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCheckOracleSlippage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xefdeed8e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function checkOracleSlippage(bytes[] paths, uint128[] amounts, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (uniswapv3Router *Uniswapv3Router) TryPackCheckOracleSlippage(paths [][]byte, amounts []*big.Int, maximumTickDivergence *big.Int, secondsAgo uint32) ([]byte, error) {
	return uniswapv3Router.abi.Pack("checkOracleSlippage", paths, amounts, maximumTickDivergence, secondsAgo)
}

// PackCheckOracleSlippage0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf25801a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (uniswapv3Router *Uniswapv3Router) PackCheckOracleSlippage0(path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) []byte {
	enc, err := uniswapv3Router.abi.Pack("checkOracleSlippage0", path, maximumTickDivergence, secondsAgo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCheckOracleSlippage0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf25801a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function checkOracleSlippage(bytes path, uint24 maximumTickDivergence, uint32 secondsAgo) view returns()
func (uniswapv3Router *Uniswapv3Router) TryPackCheckOracleSlippage0(path []byte, maximumTickDivergence *big.Int, secondsAgo uint32) ([]byte, error) {
	return uniswapv3Router.abi.Pack("checkOracleSlippage0", path, maximumTickDivergence, secondsAgo)
}

// PackExactInput is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb858183f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (uniswapv3Router *Uniswapv3Router) PackExactInput(params IV3SwapRouterExactInputParams) []byte {
	enc, err := uniswapv3Router.abi.Pack("exactInput", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExactInput is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb858183f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (uniswapv3Router *Uniswapv3Router) TryPackExactInput(params IV3SwapRouterExactInputParams) ([]byte, error) {
	return uniswapv3Router.abi.Pack("exactInput", params)
}

// UnpackExactInput is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb858183f.
//
// Solidity: function exactInput((bytes,address,uint256,uint256) params) payable returns(uint256 amountOut)
func (uniswapv3Router *Uniswapv3Router) UnpackExactInput(data []byte) (*big.Int, error) {
	out, err := uniswapv3Router.abi.Unpack("exactInput", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackExactInputSingle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x04e45aaf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (uniswapv3Router *Uniswapv3Router) PackExactInputSingle(params IV3SwapRouterExactInputSingleParams) []byte {
	enc, err := uniswapv3Router.abi.Pack("exactInputSingle", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExactInputSingle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x04e45aaf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (uniswapv3Router *Uniswapv3Router) TryPackExactInputSingle(params IV3SwapRouterExactInputSingleParams) ([]byte, error) {
	return uniswapv3Router.abi.Pack("exactInputSingle", params)
}

// UnpackExactInputSingle is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x04e45aaf.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (uniswapv3Router *Uniswapv3Router) UnpackExactInputSingle(data []byte) (*big.Int, error) {
	out, err := uniswapv3Router.abi.Unpack("exactInputSingle", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackExactOutput is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09b81346.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (uniswapv3Router *Uniswapv3Router) PackExactOutput(params IV3SwapRouterExactOutputParams) []byte {
	enc, err := uniswapv3Router.abi.Pack("exactOutput", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExactOutput is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09b81346.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (uniswapv3Router *Uniswapv3Router) TryPackExactOutput(params IV3SwapRouterExactOutputParams) ([]byte, error) {
	return uniswapv3Router.abi.Pack("exactOutput", params)
}

// UnpackExactOutput is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x09b81346.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256) params) payable returns(uint256 amountIn)
func (uniswapv3Router *Uniswapv3Router) UnpackExactOutput(data []byte) (*big.Int, error) {
	out, err := uniswapv3Router.abi.Unpack("exactOutput", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackExactOutputSingle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5023b4df.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (uniswapv3Router *Uniswapv3Router) PackExactOutputSingle(params IV3SwapRouterExactOutputSingleParams) []byte {
	enc, err := uniswapv3Router.abi.Pack("exactOutputSingle", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExactOutputSingle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5023b4df.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (uniswapv3Router *Uniswapv3Router) TryPackExactOutputSingle(params IV3SwapRouterExactOutputSingleParams) ([]byte, error) {
	return uniswapv3Router.abi.Pack("exactOutputSingle", params)
}

// UnpackExactOutputSingle is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5023b4df.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (uniswapv3Router *Uniswapv3Router) UnpackExactOutputSingle(data []byte) (*big.Int, error) {
	out, err := uniswapv3Router.abi.Unpack("exactOutputSingle", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc45a0155.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function factory() view returns(address)
func (uniswapv3Router *Uniswapv3Router) PackFactory() []byte {
	enc, err := uniswapv3Router.abi.Pack("factory")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc45a0155.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function factory() view returns(address)
func (uniswapv3Router *Uniswapv3Router) TryPackFactory() ([]byte, error) {
	return uniswapv3Router.abi.Pack("factory")
}

// UnpackFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (uniswapv3Router *Uniswapv3Router) UnpackFactory(data []byte) (common.Address, error) {
	out, err := uniswapv3Router.abi.Unpack("factory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFactoryV2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x68e0d4e1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function factoryV2() view returns(address)
func (uniswapv3Router *Uniswapv3Router) PackFactoryV2() []byte {
	enc, err := uniswapv3Router.abi.Pack("factoryV2")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFactoryV2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x68e0d4e1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function factoryV2() view returns(address)
func (uniswapv3Router *Uniswapv3Router) TryPackFactoryV2() ([]byte, error) {
	return uniswapv3Router.abi.Pack("factoryV2")
}

// UnpackFactoryV2 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x68e0d4e1.
//
// Solidity: function factoryV2() view returns(address)
func (uniswapv3Router *Uniswapv3Router) UnpackFactoryV2(data []byte) (common.Address, error) {
	out, err := uniswapv3Router.abi.Unpack("factoryV2", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetApprovalType is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdee00f35.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (uniswapv3Router *Uniswapv3Router) PackGetApprovalType(token common.Address, amount *big.Int) []byte {
	enc, err := uniswapv3Router.abi.Pack("getApprovalType", token, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetApprovalType is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdee00f35.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (uniswapv3Router *Uniswapv3Router) TryPackGetApprovalType(token common.Address, amount *big.Int) ([]byte, error) {
	return uniswapv3Router.abi.Pack("getApprovalType", token, amount)
}

// UnpackGetApprovalType is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdee00f35.
//
// Solidity: function getApprovalType(address token, uint256 amount) returns(uint8)
func (uniswapv3Router *Uniswapv3Router) UnpackGetApprovalType(data []byte) (uint8, error) {
	out, err := uniswapv3Router.abi.Unpack("getApprovalType", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackIncreaseLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf100b205.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (uniswapv3Router *Uniswapv3Router) PackIncreaseLiquidity(params IApproveAndCallIncreaseLiquidityParams) []byte {
	enc, err := uniswapv3Router.abi.Pack("increaseLiquidity", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIncreaseLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf100b205.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (uniswapv3Router *Uniswapv3Router) TryPackIncreaseLiquidity(params IApproveAndCallIncreaseLiquidityParams) ([]byte, error) {
	return uniswapv3Router.abi.Pack("increaseLiquidity", params)
}

// UnpackIncreaseLiquidity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf100b205.
//
// Solidity: function increaseLiquidity((address,address,uint256,uint256,uint256) params) payable returns(bytes result)
func (uniswapv3Router *Uniswapv3Router) UnpackIncreaseLiquidity(data []byte) ([]byte, error) {
	out, err := uniswapv3Router.abi.Unpack("increaseLiquidity", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11ed56c9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (uniswapv3Router *Uniswapv3Router) PackMint(params IApproveAndCallMintParams) []byte {
	enc, err := uniswapv3Router.abi.Pack("mint", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMint is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11ed56c9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (uniswapv3Router *Uniswapv3Router) TryPackMint(params IApproveAndCallMintParams) ([]byte, error) {
	return uniswapv3Router.abi.Pack("mint", params)
}

// UnpackMint is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x11ed56c9.
//
// Solidity: function mint((address,address,uint24,int24,int24,uint256,uint256,address) params) payable returns(bytes result)
func (uniswapv3Router *Uniswapv3Router) UnpackMint(data []byte) ([]byte, error) {
	out, err := uniswapv3Router.abi.Unpack("mint", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackMulticall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f0464d1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (uniswapv3Router *Uniswapv3Router) PackMulticall(previousBlockhash [32]byte, data [][]byte) []byte {
	enc, err := uniswapv3Router.abi.Pack("multicall", previousBlockhash, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMulticall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f0464d1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (uniswapv3Router *Uniswapv3Router) TryPackMulticall(previousBlockhash [32]byte, data [][]byte) ([]byte, error) {
	return uniswapv3Router.abi.Pack("multicall", previousBlockhash, data)
}

// UnpackMulticall is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1f0464d1.
//
// Solidity: function multicall(bytes32 previousBlockhash, bytes[] data) payable returns(bytes[])
func (uniswapv3Router *Uniswapv3Router) UnpackMulticall(data []byte) ([][]byte, error) {
	out, err := uniswapv3Router.abi.Unpack("multicall", data)
	if err != nil {
		return *new([][]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	return out0, nil
}

// PackMulticall0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ae401dc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (uniswapv3Router *Uniswapv3Router) PackMulticall0(deadline *big.Int, data [][]byte) []byte {
	enc, err := uniswapv3Router.abi.Pack("multicall0", deadline, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMulticall0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ae401dc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (uniswapv3Router *Uniswapv3Router) TryPackMulticall0(deadline *big.Int, data [][]byte) ([]byte, error) {
	return uniswapv3Router.abi.Pack("multicall0", deadline, data)
}

// UnpackMulticall0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5ae401dc.
//
// Solidity: function multicall(uint256 deadline, bytes[] data) payable returns(bytes[])
func (uniswapv3Router *Uniswapv3Router) UnpackMulticall0(data []byte) ([][]byte, error) {
	out, err := uniswapv3Router.abi.Unpack("multicall0", data)
	if err != nil {
		return *new([][]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	return out0, nil
}

// PackMulticall1 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac9650d8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (uniswapv3Router *Uniswapv3Router) PackMulticall1(data [][]byte) []byte {
	enc, err := uniswapv3Router.abi.Pack("multicall1", data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMulticall1 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac9650d8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (uniswapv3Router *Uniswapv3Router) TryPackMulticall1(data [][]byte) ([]byte, error) {
	return uniswapv3Router.abi.Pack("multicall1", data)
}

// UnpackMulticall1 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (uniswapv3Router *Uniswapv3Router) UnpackMulticall1(data []byte) ([][]byte, error) {
	out, err := uniswapv3Router.abi.Unpack("multicall1", data)
	if err != nil {
		return *new([][]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	return out0, nil
}

// PackPositionManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x791b98bc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function positionManager() view returns(address)
func (uniswapv3Router *Uniswapv3Router) PackPositionManager() []byte {
	enc, err := uniswapv3Router.abi.Pack("positionManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPositionManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x791b98bc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function positionManager() view returns(address)
func (uniswapv3Router *Uniswapv3Router) TryPackPositionManager() ([]byte, error) {
	return uniswapv3Router.abi.Pack("positionManager")
}

// UnpackPositionManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x791b98bc.
//
// Solidity: function positionManager() view returns(address)
func (uniswapv3Router *Uniswapv3Router) UnpackPositionManager(data []byte) (common.Address, error) {
	out, err := uniswapv3Router.abi.Unpack("positionManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPull is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2d5d56b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackPull(token common.Address, value *big.Int) []byte {
	enc, err := uniswapv3Router.abi.Pack("pull", token, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPull is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2d5d56b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pull(address token, uint256 value) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackPull(token common.Address, value *big.Int) ([]byte, error) {
	return uniswapv3Router.abi.Pack("pull", token, value)
}

// PackRefundETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12210e8a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function refundETH() payable returns()
func (uniswapv3Router *Uniswapv3Router) PackRefundETH() []byte {
	enc, err := uniswapv3Router.abi.Pack("refundETH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRefundETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12210e8a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function refundETH() payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackRefundETH() ([]byte, error) {
	return uniswapv3Router.abi.Pack("refundETH")
}

// PackSelfPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3995c67.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackSelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := uniswapv3Router.abi.Pack("selfPermit", token, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSelfPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3995c67.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackSelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return uniswapv3Router.abi.Pack("selfPermit", token, value, deadline, v, r, s)
}

// PackSelfPermitAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4659a494.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackSelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := uniswapv3Router.abi.Pack("selfPermitAllowed", token, nonce, expiry, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSelfPermitAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4659a494.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackSelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return uniswapv3Router.abi.Pack("selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// PackSelfPermitAllowedIfNecessary is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4a78f0c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackSelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := uniswapv3Router.abi.Pack("selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSelfPermitAllowedIfNecessary is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4a78f0c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackSelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return uniswapv3Router.abi.Pack("selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// PackSelfPermitIfNecessary is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc2e3140a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackSelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := uniswapv3Router.abi.Pack("selfPermitIfNecessary", token, value, deadline, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSelfPermitIfNecessary is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc2e3140a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackSelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return uniswapv3Router.abi.Pack("selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// PackSwapExactTokensForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x472b43f3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (uniswapv3Router *Uniswapv3Router) PackSwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("swapExactTokensForTokens", amountIn, amountOutMin, path, to)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactTokensForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x472b43f3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (uniswapv3Router *Uniswapv3Router) TryPackSwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("swapExactTokensForTokens", amountIn, amountOutMin, path, to)
}

// UnpackSwapExactTokensForTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x472b43f3.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to) payable returns(uint256 amountOut)
func (uniswapv3Router *Uniswapv3Router) UnpackSwapExactTokensForTokens(data []byte) (*big.Int, error) {
	out, err := uniswapv3Router.abi.Unpack("swapExactTokensForTokens", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSwapTokensForExactTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42712a67.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (uniswapv3Router *Uniswapv3Router) PackSwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("swapTokensForExactTokens", amountOut, amountInMax, path, to)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapTokensForExactTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42712a67.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (uniswapv3Router *Uniswapv3Router) TryPackSwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("swapTokensForExactTokens", amountOut, amountInMax, path, to)
}

// UnpackSwapTokensForExactTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x42712a67.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to) payable returns(uint256 amountIn)
func (uniswapv3Router *Uniswapv3Router) UnpackSwapTokensForExactTokens(data []byte) (*big.Int, error) {
	out, err := uniswapv3Router.abi.Unpack("swapTokensForExactTokens", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSweepToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdf2ab5bb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackSweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("sweepToken", token, amountMinimum, recipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSweepToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdf2ab5bb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackSweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("sweepToken", token, amountMinimum, recipient)
}

// PackSweepToken0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe90a182f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackSweepToken0(token common.Address, amountMinimum *big.Int) []byte {
	enc, err := uniswapv3Router.abi.Pack("sweepToken0", token, amountMinimum)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSweepToken0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe90a182f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackSweepToken0(token common.Address, amountMinimum *big.Int) ([]byte, error) {
	return uniswapv3Router.abi.Pack("sweepToken0", token, amountMinimum)
}

// PackSweepTokenWithFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3068c554.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackSweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("sweepTokenWithFee", token, amountMinimum, feeBips, feeRecipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSweepTokenWithFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3068c554.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackSweepTokenWithFee(token common.Address, amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("sweepTokenWithFee", token, amountMinimum, feeBips, feeRecipient)
}

// PackSweepTokenWithFee0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe0e189a0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackSweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("sweepTokenWithFee0", token, amountMinimum, recipient, feeBips, feeRecipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSweepTokenWithFee0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe0e189a0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackSweepTokenWithFee0(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("sweepTokenWithFee0", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// PackUniswapV3SwapCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa461e33.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (uniswapv3Router *Uniswapv3Router) PackUniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) []byte {
	enc, err := uniswapv3Router.abi.Pack("uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUniswapV3SwapCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa461e33.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (uniswapv3Router *Uniswapv3Router) TryPackUniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) ([]byte, error) {
	return uniswapv3Router.abi.Pack("uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// PackUnwrapWETH9 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49404b7c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackUnwrapWETH9(amountMinimum *big.Int, recipient common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("unwrapWETH9", amountMinimum, recipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnwrapWETH9 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49404b7c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackUnwrapWETH9(amountMinimum *big.Int, recipient common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("unwrapWETH9", amountMinimum, recipient)
}

// PackUnwrapWETH90 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49616997.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackUnwrapWETH90(amountMinimum *big.Int) []byte {
	enc, err := uniswapv3Router.abi.Pack("unwrapWETH90", amountMinimum)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnwrapWETH90 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49616997.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackUnwrapWETH90(amountMinimum *big.Int) ([]byte, error) {
	return uniswapv3Router.abi.Pack("unwrapWETH90", amountMinimum)
}

// PackUnwrapWETH9WithFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b2c0a37.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackUnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnwrapWETH9WithFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b2c0a37.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackUnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// PackUnwrapWETH9WithFee0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd4ef38de.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackUnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) []byte {
	enc, err := uniswapv3Router.abi.Pack("unwrapWETH9WithFee0", amountMinimum, feeBips, feeRecipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnwrapWETH9WithFee0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd4ef38de.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, uint256 feeBips, address feeRecipient) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackUnwrapWETH9WithFee0(amountMinimum *big.Int, feeBips *big.Int, feeRecipient common.Address) ([]byte, error) {
	return uniswapv3Router.abi.Pack("unwrapWETH9WithFee0", amountMinimum, feeBips, feeRecipient)
}

// PackWrapETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c58db4f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (uniswapv3Router *Uniswapv3Router) PackWrapETH(value *big.Int) []byte {
	enc, err := uniswapv3Router.abi.Pack("wrapETH", value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWrapETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c58db4f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function wrapETH(uint256 value) payable returns()
func (uniswapv3Router *Uniswapv3Router) TryPackWrapETH(value *big.Int) ([]byte, error) {
	return uniswapv3Router.abi.Pack("wrapETH", value)
}
