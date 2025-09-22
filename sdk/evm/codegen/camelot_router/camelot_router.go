// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package camelot_router

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

// CamelotRouterMetaData contains all meta data concerning the CamelotRouter contract.
var CamelotRouterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token2\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "CamelotRouter",
}

// CamelotRouter is an auto generated Go binding around an Ethereum contract.
type CamelotRouter struct {
	abi abi.ABI
}

// NewCamelotRouter creates a new instance of CamelotRouter.
func NewCamelotRouter() *CamelotRouter {
	parsed, err := CamelotRouterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &CamelotRouter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *CamelotRouter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _factory, address _WETH) returns()
func (camelotRouter *CamelotRouter) PackConstructor(_factory common.Address, _WETH common.Address) []byte {
	enc, err := camelotRouter.abi.Pack("", _factory, _WETH)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad5c4648.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function WETH() view returns(address)
func (camelotRouter *CamelotRouter) PackWETH() []byte {
	enc, err := camelotRouter.abi.Pack("WETH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad5c4648.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function WETH() view returns(address)
func (camelotRouter *CamelotRouter) TryPackWETH() ([]byte, error) {
	return camelotRouter.abi.Pack("WETH")
}

// UnpackWETH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (camelotRouter *CamelotRouter) UnpackWETH(data []byte) (common.Address, error) {
	out, err := camelotRouter.abi.Unpack("WETH", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackAddLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe8e33700.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (camelotRouter *CamelotRouter) PackAddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := camelotRouter.abi.Pack("addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe8e33700.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (camelotRouter *CamelotRouter) TryPackAddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return camelotRouter.abi.Pack("addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
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
func (camelotRouter *CamelotRouter) UnpackAddLiquidity(data []byte) (AddLiquidityOutput, error) {
	out, err := camelotRouter.abi.Unpack("addLiquidity", data)
	outstruct := new(AddLiquidityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Liquidity = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackAddLiquidityETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf305d719.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (camelotRouter *CamelotRouter) PackAddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := camelotRouter.abi.Pack("addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddLiquidityETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf305d719.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (camelotRouter *CamelotRouter) TryPackAddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return camelotRouter.abi.Pack("addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
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
func (camelotRouter *CamelotRouter) UnpackAddLiquidityETH(data []byte) (AddLiquidityETHOutput, error) {
	out, err := camelotRouter.abi.Unpack("addLiquidityETH", data)
	outstruct := new(AddLiquidityETHOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountToken = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountETH = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Liquidity = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc45a0155.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function factory() view returns(address)
func (camelotRouter *CamelotRouter) PackFactory() []byte {
	enc, err := camelotRouter.abi.Pack("factory")
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
func (camelotRouter *CamelotRouter) TryPackFactory() ([]byte, error) {
	return camelotRouter.abi.Pack("factory")
}

// UnpackFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (camelotRouter *CamelotRouter) UnpackFactory(data []byte) (common.Address, error) {
	out, err := camelotRouter.abi.Unpack("factory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAmountsOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd06ca61f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (camelotRouter *CamelotRouter) PackGetAmountsOut(amountIn *big.Int, path []common.Address) []byte {
	enc, err := camelotRouter.abi.Pack("getAmountsOut", amountIn, path)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAmountsOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd06ca61f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (camelotRouter *CamelotRouter) TryPackGetAmountsOut(amountIn *big.Int, path []common.Address) ([]byte, error) {
	return camelotRouter.abi.Pack("getAmountsOut", amountIn, path)
}

// UnpackGetAmountsOut is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (camelotRouter *CamelotRouter) UnpackGetAmountsOut(data []byte) ([]*big.Int, error) {
	out, err := camelotRouter.abi.Unpack("getAmountsOut", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackGetPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe6a43905.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPair(address token1, address token2) view returns(address)
func (camelotRouter *CamelotRouter) PackGetPair(token1 common.Address, token2 common.Address) []byte {
	enc, err := camelotRouter.abi.Pack("getPair", token1, token2)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe6a43905.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPair(address token1, address token2) view returns(address)
func (camelotRouter *CamelotRouter) TryPackGetPair(token1 common.Address, token2 common.Address) ([]byte, error) {
	return camelotRouter.abi.Pack("getPair", token1, token2)
}

// UnpackGetPair is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe6a43905.
//
// Solidity: function getPair(address token1, address token2) view returns(address)
func (camelotRouter *CamelotRouter) UnpackGetPair(data []byte) (common.Address, error) {
	out, err := camelotRouter.abi.Unpack("getPair", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackQuote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad615dec.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (camelotRouter *CamelotRouter) PackQuote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) []byte {
	enc, err := camelotRouter.abi.Pack("quote", amountA, reserveA, reserveB)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad615dec.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (camelotRouter *CamelotRouter) TryPackQuote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) ([]byte, error) {
	return camelotRouter.abi.Pack("quote", amountA, reserveA, reserveB)
}

// UnpackQuote is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (camelotRouter *CamelotRouter) UnpackQuote(data []byte) (*big.Int, error) {
	out, err := camelotRouter.abi.Unpack("quote", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRemoveLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbaa2abde.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (camelotRouter *CamelotRouter) PackRemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := camelotRouter.abi.Pack("removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbaa2abde.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (camelotRouter *CamelotRouter) TryPackRemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return camelotRouter.abi.Pack("removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
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
func (camelotRouter *CamelotRouter) UnpackRemoveLiquidity(data []byte) (RemoveLiquidityOutput, error) {
	out, err := camelotRouter.abi.Unpack("removeLiquidity", data)
	outstruct := new(RemoveLiquidityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackRemoveLiquidityETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02751cec.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (camelotRouter *CamelotRouter) PackRemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := camelotRouter.abi.Pack("removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidityETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02751cec.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (camelotRouter *CamelotRouter) TryPackRemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return camelotRouter.abi.Pack("removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
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
func (camelotRouter *CamelotRouter) UnpackRemoveLiquidityETH(data []byte) (RemoveLiquidityETHOutput, error) {
	out, err := camelotRouter.abi.Unpack("removeLiquidityETH", data)
	outstruct := new(RemoveLiquidityETHOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountToken = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountETH = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackRemoveLiquidityETHSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf2979eb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (camelotRouter *CamelotRouter) PackRemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := camelotRouter.abi.Pack("removeLiquidityETHSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidityETHSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf2979eb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (camelotRouter *CamelotRouter) TryPackRemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return camelotRouter.abi.Pack("removeLiquidityETHSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// UnpackRemoveLiquidityETHSupportingFeeOnTransferTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaf2979eb.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (camelotRouter *CamelotRouter) UnpackRemoveLiquidityETHSupportingFeeOnTransferTokens(data []byte) (*big.Int, error) {
	out, err := camelotRouter.abi.Unpack("removeLiquidityETHSupportingFeeOnTransferTokens", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRemoveLiquidityETHWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xded9382a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (camelotRouter *CamelotRouter) PackRemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := camelotRouter.abi.Pack("removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidityETHWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xded9382a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidityETHWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountETH)
func (camelotRouter *CamelotRouter) TryPackRemoveLiquidityETHWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return camelotRouter.abi.Pack("removeLiquidityETHWithPermit", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
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
func (camelotRouter *CamelotRouter) UnpackRemoveLiquidityETHWithPermit(data []byte) (RemoveLiquidityETHWithPermitOutput, error) {
	out, err := camelotRouter.abi.Unpack("removeLiquidityETHWithPermit", data)
	outstruct := new(RemoveLiquidityETHWithPermitOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountToken = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountETH = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b0d5984.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (camelotRouter *CamelotRouter) PackRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := camelotRouter.abi.Pack("removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b0d5984.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (camelotRouter *CamelotRouter) TryPackRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return camelotRouter.abi.Pack("removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline, approveMax, v, r, s)
}

// UnpackRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b0d5984.
//
// Solidity: function removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountETH)
func (camelotRouter *CamelotRouter) UnpackRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(data []byte) (*big.Int, error) {
	out, err := camelotRouter.abi.Unpack("removeLiquidityETHWithPermitSupportingFeeOnTransferTokens", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRemoveLiquidityWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2195995c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (camelotRouter *CamelotRouter) PackRemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := camelotRouter.abi.Pack("removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidityWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2195995c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (camelotRouter *CamelotRouter) TryPackRemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return camelotRouter.abi.Pack("removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
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
func (camelotRouter *CamelotRouter) UnpackRemoveLiquidityWithPermit(data []byte) (RemoveLiquidityWithPermitOutput, error) {
	out, err := camelotRouter.abi.Unpack("removeLiquidityWithPermit", data)
	outstruct := new(RemoveLiquidityWithPermitOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackSwapExactETHForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb4822be3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, address referrer, uint256 deadline) payable returns()
func (camelotRouter *CamelotRouter) PackSwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, referrer common.Address, deadline *big.Int) []byte {
	enc, err := camelotRouter.abi.Pack("swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, referrer, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactETHForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb4822be3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, address referrer, uint256 deadline) payable returns()
func (camelotRouter *CamelotRouter) TryPackSwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, referrer common.Address, deadline *big.Int) ([]byte, error) {
	return camelotRouter.abi.Pack("swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, referrer, deadline)
}

// PackSwapExactTokensForETHSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52aa4c22.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, address referrer, uint256 deadline) returns()
func (camelotRouter *CamelotRouter) PackSwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, referrer common.Address, deadline *big.Int) []byte {
	enc, err := camelotRouter.abi.Pack("swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, referrer, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactTokensForETHSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52aa4c22.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, address referrer, uint256 deadline) returns()
func (camelotRouter *CamelotRouter) TryPackSwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, referrer common.Address, deadline *big.Int) ([]byte, error) {
	return camelotRouter.abi.Pack("swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, referrer, deadline)
}

// PackSwapExactTokensForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac3893ba.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, address referrer, uint256 deadline) returns()
func (camelotRouter *CamelotRouter) PackSwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, referrer common.Address, deadline *big.Int) []byte {
	enc, err := camelotRouter.abi.Pack("swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, referrer, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactTokensForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac3893ba.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, address referrer, uint256 deadline) returns()
func (camelotRouter *CamelotRouter) TryPackSwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, referrer common.Address, deadline *big.Int) ([]byte, error) {
	return camelotRouter.abi.Pack("swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, referrer, deadline)
}
