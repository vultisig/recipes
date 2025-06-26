// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv2_router

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

// Uniswapv2RouterMetaData contains all meta data concerning the Uniswapv2Router contract.
var Uniswapv2RouterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "Uniswapv2Router",
}

// Uniswapv2Router is an auto generated Go binding around an Ethereum contract.
type Uniswapv2Router struct {
	abi abi.ABI
}

// NewUniswapv2Router creates a new instance of Uniswapv2Router.
func NewUniswapv2Router() *Uniswapv2Router {
	parsed, err := Uniswapv2RouterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Uniswapv2Router{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Uniswapv2Router) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _factory, address _WETH) returns()
func (uniswapv2Router *Uniswapv2Router) PackConstructor(_factory common.Address, _WETH common.Address) []byte {
	enc, err := uniswapv2Router.abi.Pack("", _factory, _WETH)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (uniswapv2Router *Uniswapv2Router) PackWETH() []byte {
	enc, err := uniswapv2Router.abi.Pack("WETH")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackWETH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (uniswapv2Router *Uniswapv2Router) UnpackWETH(data []byte) (common.Address, error) {
	out, err := uniswapv2Router.abi.Unpack("WETH", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackAddLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (uniswapv2Router *Uniswapv2Router) PackAddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// AddLiquidityOutput serves as a container for the return parameters of contract
// method AddLiquidity.
type AddLiquidityOutput struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}

// UnpackAddLiquidity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (uniswapv2Router *Uniswapv2Router) UnpackAddLiquidity(data []byte) (AddLiquidityOutput, error) {
	out, err := uniswapv2Router.abi.Unpack("addLiquidity", data)
	outstruct := new(AddLiquidityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Liquidity = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackAddLiquidityETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (uniswapv2Router *Uniswapv2Router) PackAddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// AddLiquidityETHOutput serves as a container for the return parameters of contract
// method AddLiquidityETH.
type AddLiquidityETHOutput struct {
	AmountToken *big.Int
	AmountETH   *big.Int
	Liquidity   *big.Int
}

// UnpackAddLiquidityETH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf305d719.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (uniswapv2Router *Uniswapv2Router) UnpackAddLiquidityETH(data []byte) (AddLiquidityETHOutput, error) {
	out, err := uniswapv2Router.abi.Unpack("addLiquidityETH", data)
	outstruct := new(AddLiquidityETHOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountToken = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountETH = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Liquidity = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (uniswapv2Router *Uniswapv2Router) PackFactory() []byte {
	enc, err := uniswapv2Router.abi.Pack("factory")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (uniswapv2Router *Uniswapv2Router) UnpackFactory(data []byte) (common.Address, error) {
	out, err := uniswapv2Router.abi.Unpack("factory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackGetAmountIn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (uniswapv2Router *Uniswapv2Router) PackGetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("getAmountIn", amountOut, reserveIn, reserveOut)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAmountIn is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (uniswapv2Router *Uniswapv2Router) UnpackGetAmountIn(data []byte) (*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("getAmountIn", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetAmountOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (uniswapv2Router *Uniswapv2Router) PackGetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("getAmountOut", amountIn, reserveIn, reserveOut)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAmountOut is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (uniswapv2Router *Uniswapv2Router) UnpackGetAmountOut(data []byte) (*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("getAmountOut", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackGetAmountsIn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) PackGetAmountsIn(amountOut *big.Int, path []common.Address) []byte {
	enc, err := uniswapv2Router.abi.Pack("getAmountsIn", amountOut, path)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAmountsIn is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) UnpackGetAmountsIn(data []byte) ([]*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("getAmountsIn", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackGetAmountsOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) PackGetAmountsOut(amountIn *big.Int, path []common.Address) []byte {
	enc, err := uniswapv2Router.abi.Pack("getAmountsOut", amountIn, path)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetAmountsOut is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) UnpackGetAmountsOut(data []byte) ([]*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("getAmountsOut", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackQuote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (uniswapv2Router *Uniswapv2Router) PackQuote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("quote", amountA, reserveA, reserveB)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackQuote is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (uniswapv2Router *Uniswapv2Router) UnpackQuote(data []byte) (*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("quote", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackRemoveLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (uniswapv2Router *Uniswapv2Router) PackRemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// RemoveLiquidityOutput serves as a container for the return parameters of contract
// method RemoveLiquidity.
type RemoveLiquidityOutput struct {
	AmountA *big.Int
	AmountB *big.Int
}

// UnpackRemoveLiquidity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (uniswapv2Router *Uniswapv2Router) UnpackRemoveLiquidity(data []byte) (RemoveLiquidityOutput, error) {
	out, err := uniswapv2Router.abi.Unpack("removeLiquidity", data)
	outstruct := new(RemoveLiquidityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackRemoveLiquidityETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (uniswapv2Router *Uniswapv2Router) PackRemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// RemoveLiquidityETHOutput serves as a container for the return parameters of contract
// method RemoveLiquidityETH.
type RemoveLiquidityETHOutput struct {
	AmountToken *big.Int
	AmountETH   *big.Int
}

// UnpackRemoveLiquidityETH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x02751cec.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (uniswapv2Router *Uniswapv2Router) UnpackRemoveLiquidityETH(data []byte) (RemoveLiquidityETHOutput, error) {
	out, err := uniswapv2Router.abi.Unpack("removeLiquidityETH", data)
	outstruct := new(RemoveLiquidityETHOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountToken = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountETH = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackRemoveLiquidityETHSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (uniswapv2Router *Uniswapv2Router) PackRemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("removeLiquidityETHSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRemoveLiquidityETHSupportingFeeOnTransferTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (uniswapv2Router *Uniswapv2Router) UnpackRemoveLiquidityETHSupportingFeeOnTransferTokens(data []byte) (*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("removeLiquidityETHSupportingFeeOnTransferTokens", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackRemoveLiquidityETHWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (uniswapv2Router *Uniswapv2Router) PackRemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := uniswapv2Router.abi.Pack("removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// RemoveLiquidityETHWithPermitOutput serves as a container for the return parameters of contract
// method RemoveLiquidityETHWithPermit.
type RemoveLiquidityETHWithPermitOutput struct {
	AmountToken *big.Int
	AmountETH   *big.Int
}

// UnpackRemoveLiquidityETHWithPermit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xded9382a.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (uniswapv2Router *Uniswapv2Router) UnpackRemoveLiquidityETHWithPermit(data []byte) (RemoveLiquidityETHWithPermitOutput, error) {
	out, err := uniswapv2Router.abi.Unpack("removeLiquidityETHWithPermit", data)
	outstruct := new(RemoveLiquidityETHWithPermitOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountToken = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountETH = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (uniswapv2Router *Uniswapv2Router) PackRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := uniswapv2Router.abi.Pack("removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (uniswapv2Router *Uniswapv2Router) UnpackRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(data []byte) (*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackRemoveLiquidityWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (uniswapv2Router *Uniswapv2Router) PackRemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := uniswapv2Router.abi.Pack("removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// RemoveLiquidityWithPermitOutput serves as a container for the return parameters of contract
// method RemoveLiquidityWithPermit.
type RemoveLiquidityWithPermitOutput struct {
	AmountA *big.Int
	AmountB *big.Int
}

// UnpackRemoveLiquidityWithPermit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (uniswapv2Router *Uniswapv2Router) UnpackRemoveLiquidityWithPermit(data []byte) (RemoveLiquidityWithPermitOutput, error) {
	out, err := uniswapv2Router.abi.Unpack("removeLiquidityWithPermit", data)
	outstruct := new(RemoveLiquidityWithPermitOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackSwapETHForExactTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) PackSwapETHForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("swapETHForExactTokens", amountOut, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSwapETHForExactTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) UnpackSwapETHForExactTokens(data []byte) ([]*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("swapETHForExactTokens", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackSwapExactETHForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) PackSwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("swapExactETHForTokens", amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSwapExactETHForTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) UnpackSwapExactETHForTokens(data []byte) ([]*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("swapExactETHForTokens", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackSwapExactETHForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb6f9de95.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (uniswapv2Router *Uniswapv2Router) PackSwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSwapExactTokensForETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) PackSwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSwapExactTokensForETH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) UnpackSwapExactTokensForETH(data []byte) ([]*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("swapExactTokensForETH", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackSwapExactTokensForETHSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x791ac947.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (uniswapv2Router *Uniswapv2Router) PackSwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSwapExactTokensForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) PackSwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSwapExactTokensForTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) UnpackSwapExactTokensForTokens(data []byte) ([]*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("swapExactTokensForTokens", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackSwapExactTokensForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (uniswapv2Router *Uniswapv2Router) PackSwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSwapTokensForExactETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) PackSwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("swapTokensForExactETH", amountOut, amountInMax, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSwapTokensForExactETH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) UnpackSwapTokensForExactETH(data []byte) ([]*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("swapTokensForExactETH", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}

// PackSwapTokensForExactTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) PackSwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := uniswapv2Router.abi.Pack("swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSwapTokensForExactTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (uniswapv2Router *Uniswapv2Router) UnpackSwapTokensForExactTokens(data []byte) ([]*big.Int, error) {
	out, err := uniswapv2Router.abi.Unpack("swapTokensForExactTokens", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, err
}
