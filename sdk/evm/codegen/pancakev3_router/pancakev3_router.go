// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakev3_router

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

// ISwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// ISwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// ISwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// Pancakev3RouterMetaData contains all meta data concerning the Pancakev3Router contract.
var Pancakev3RouterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"pancakeV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "Pancakev3Router",
}

// Pancakev3Router is an auto generated Go binding around an Ethereum contract.
type Pancakev3Router struct {
	abi abi.ABI
}

// NewPancakev3Router creates a new instance of Pancakev3Router.
func NewPancakev3Router() *Pancakev3Router {
	parsed, err := Pancakev3RouterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Pancakev3Router{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Pancakev3Router) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _deployer, address _factory, address _WETH9) returns()
func (pancakev3Router *Pancakev3Router) PackConstructor(_deployer common.Address, _factory common.Address, _WETH9 common.Address) []byte {
	enc, err := pancakev3Router.abi.Pack("", _deployer, _factory, _WETH9)
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
func (pancakev3Router *Pancakev3Router) PackWETH9() []byte {
	enc, err := pancakev3Router.abi.Pack("WETH9")
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
func (pancakev3Router *Pancakev3Router) TryPackWETH9() ([]byte, error) {
	return pancakev3Router.abi.Pack("WETH9")
}

// UnpackWETH9 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (pancakev3Router *Pancakev3Router) UnpackWETH9(data []byte) (common.Address, error) {
	out, err := pancakev3Router.abi.Unpack("WETH9", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDeployer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5f39488.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deployer() view returns(address)
func (pancakev3Router *Pancakev3Router) PackDeployer() []byte {
	enc, err := pancakev3Router.abi.Pack("deployer")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeployer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5f39488.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deployer() view returns(address)
func (pancakev3Router *Pancakev3Router) TryPackDeployer() ([]byte, error) {
	return pancakev3Router.abi.Pack("deployer")
}

// UnpackDeployer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (pancakev3Router *Pancakev3Router) UnpackDeployer(data []byte) (common.Address, error) {
	out, err := pancakev3Router.abi.Unpack("deployer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackExactInput is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc04b8d59.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (pancakev3Router *Pancakev3Router) PackExactInput(params ISwapRouterExactInputParams) []byte {
	enc, err := pancakev3Router.abi.Pack("exactInput", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExactInput is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc04b8d59.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (pancakev3Router *Pancakev3Router) TryPackExactInput(params ISwapRouterExactInputParams) ([]byte, error) {
	return pancakev3Router.abi.Pack("exactInput", params)
}

// UnpackExactInput is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (pancakev3Router *Pancakev3Router) UnpackExactInput(data []byte) (*big.Int, error) {
	out, err := pancakev3Router.abi.Unpack("exactInput", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackExactInputSingle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x414bf389.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (pancakev3Router *Pancakev3Router) PackExactInputSingle(params ISwapRouterExactInputSingleParams) []byte {
	enc, err := pancakev3Router.abi.Pack("exactInputSingle", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExactInputSingle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x414bf389.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (pancakev3Router *Pancakev3Router) TryPackExactInputSingle(params ISwapRouterExactInputSingleParams) ([]byte, error) {
	return pancakev3Router.abi.Pack("exactInputSingle", params)
}

// UnpackExactInputSingle is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (pancakev3Router *Pancakev3Router) UnpackExactInputSingle(data []byte) (*big.Int, error) {
	out, err := pancakev3Router.abi.Unpack("exactInputSingle", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackExactOutput is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf28c0498.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (pancakev3Router *Pancakev3Router) PackExactOutput(params ISwapRouterExactOutputParams) []byte {
	enc, err := pancakev3Router.abi.Pack("exactOutput", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExactOutput is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf28c0498.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (pancakev3Router *Pancakev3Router) TryPackExactOutput(params ISwapRouterExactOutputParams) ([]byte, error) {
	return pancakev3Router.abi.Pack("exactOutput", params)
}

// UnpackExactOutput is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (pancakev3Router *Pancakev3Router) UnpackExactOutput(data []byte) (*big.Int, error) {
	out, err := pancakev3Router.abi.Unpack("exactOutput", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackExactOutputSingle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb3e2198.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (pancakev3Router *Pancakev3Router) PackExactOutputSingle(params ISwapRouterExactOutputSingleParams) []byte {
	enc, err := pancakev3Router.abi.Pack("exactOutputSingle", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExactOutputSingle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb3e2198.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (pancakev3Router *Pancakev3Router) TryPackExactOutputSingle(params ISwapRouterExactOutputSingleParams) ([]byte, error) {
	return pancakev3Router.abi.Pack("exactOutputSingle", params)
}

// UnpackExactOutputSingle is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (pancakev3Router *Pancakev3Router) UnpackExactOutputSingle(data []byte) (*big.Int, error) {
	out, err := pancakev3Router.abi.Unpack("exactOutputSingle", data)
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
func (pancakev3Router *Pancakev3Router) PackFactory() []byte {
	enc, err := pancakev3Router.abi.Pack("factory")
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
func (pancakev3Router *Pancakev3Router) TryPackFactory() ([]byte, error) {
	return pancakev3Router.abi.Pack("factory")
}

// UnpackFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (pancakev3Router *Pancakev3Router) UnpackFactory(data []byte) (common.Address, error) {
	out, err := pancakev3Router.abi.Unpack("factory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackMulticall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac9650d8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (pancakev3Router *Pancakev3Router) PackMulticall(data [][]byte) []byte {
	enc, err := pancakev3Router.abi.Pack("multicall", data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMulticall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac9650d8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (pancakev3Router *Pancakev3Router) TryPackMulticall(data [][]byte) ([]byte, error) {
	return pancakev3Router.abi.Pack("multicall", data)
}

// UnpackMulticall is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (pancakev3Router *Pancakev3Router) UnpackMulticall(data []byte) ([][]byte, error) {
	out, err := pancakev3Router.abi.Unpack("multicall", data)
	if err != nil {
		return *new([][]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	return out0, nil
}

// PackPancakeV3SwapCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23a69e75.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (pancakev3Router *Pancakev3Router) PackPancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) []byte {
	enc, err := pancakev3Router.abi.Pack("pancakeV3SwapCallback", amount0Delta, amount1Delta, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPancakeV3SwapCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23a69e75.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pancakeV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (pancakev3Router *Pancakev3Router) TryPackPancakeV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) ([]byte, error) {
	return pancakev3Router.abi.Pack("pancakeV3SwapCallback", amount0Delta, amount1Delta, data)
}

// PackRefundETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12210e8a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function refundETH() payable returns()
func (pancakev3Router *Pancakev3Router) PackRefundETH() []byte {
	enc, err := pancakev3Router.abi.Pack("refundETH")
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
func (pancakev3Router *Pancakev3Router) TryPackRefundETH() ([]byte, error) {
	return pancakev3Router.abi.Pack("refundETH")
}

// PackSelfPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3995c67.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (pancakev3Router *Pancakev3Router) PackSelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := pancakev3Router.abi.Pack("selfPermit", token, value, deadline, v, r, s)
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
func (pancakev3Router *Pancakev3Router) TryPackSelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return pancakev3Router.abi.Pack("selfPermit", token, value, deadline, v, r, s)
}

// PackSelfPermitAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4659a494.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (pancakev3Router *Pancakev3Router) PackSelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := pancakev3Router.abi.Pack("selfPermitAllowed", token, nonce, expiry, v, r, s)
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
func (pancakev3Router *Pancakev3Router) TryPackSelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return pancakev3Router.abi.Pack("selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// PackSelfPermitAllowedIfNecessary is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4a78f0c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (pancakev3Router *Pancakev3Router) PackSelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := pancakev3Router.abi.Pack("selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
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
func (pancakev3Router *Pancakev3Router) TryPackSelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return pancakev3Router.abi.Pack("selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// PackSelfPermitIfNecessary is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc2e3140a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (pancakev3Router *Pancakev3Router) PackSelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := pancakev3Router.abi.Pack("selfPermitIfNecessary", token, value, deadline, v, r, s)
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
func (pancakev3Router *Pancakev3Router) TryPackSelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return pancakev3Router.abi.Pack("selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// PackSweepToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdf2ab5bb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (pancakev3Router *Pancakev3Router) PackSweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) []byte {
	enc, err := pancakev3Router.abi.Pack("sweepToken", token, amountMinimum, recipient)
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
func (pancakev3Router *Pancakev3Router) TryPackSweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) ([]byte, error) {
	return pancakev3Router.abi.Pack("sweepToken", token, amountMinimum, recipient)
}

// PackSweepTokenWithFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe0e189a0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (pancakev3Router *Pancakev3Router) PackSweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) []byte {
	enc, err := pancakev3Router.abi.Pack("sweepTokenWithFee", token, amountMinimum, recipient, feeBips, feeRecipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSweepTokenWithFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe0e189a0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (pancakev3Router *Pancakev3Router) TryPackSweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) ([]byte, error) {
	return pancakev3Router.abi.Pack("sweepTokenWithFee", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// PackUnwrapWETH9 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49404b7c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (pancakev3Router *Pancakev3Router) PackUnwrapWETH9(amountMinimum *big.Int, recipient common.Address) []byte {
	enc, err := pancakev3Router.abi.Pack("unwrapWETH9", amountMinimum, recipient)
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
func (pancakev3Router *Pancakev3Router) TryPackUnwrapWETH9(amountMinimum *big.Int, recipient common.Address) ([]byte, error) {
	return pancakev3Router.abi.Pack("unwrapWETH9", amountMinimum, recipient)
}

// PackUnwrapWETH9WithFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b2c0a37.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (pancakev3Router *Pancakev3Router) PackUnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) []byte {
	enc, err := pancakev3Router.abi.Pack("unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
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
func (pancakev3Router *Pancakev3Router) TryPackUnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) ([]byte, error) {
	return pancakev3Router.abi.Pack("unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
}
