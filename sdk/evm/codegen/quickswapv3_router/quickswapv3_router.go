// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quickswapv3_router

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
	TokenIn          common.Address
	TokenOut         common.Address
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
	LimitSqrtPrice   *big.Int
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
	TokenIn         common.Address
	TokenOut        common.Address
	Fee             *big.Int
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
	LimitSqrtPrice  *big.Int
}

// Quickswapv3RouterMetaData contains all meta data concerning the Quickswapv3Router contract.
var Quickswapv3RouterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WNativeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolDeployer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WNativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"algebraSwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingleSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"limitSqrtPrice\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundNativeToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWNativeToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWNativeTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "Quickswapv3Router",
}

// Quickswapv3Router is an auto generated Go binding around an Ethereum contract.
type Quickswapv3Router struct {
	abi abi.ABI
}

// NewQuickswapv3Router creates a new instance of Quickswapv3Router.
func NewQuickswapv3Router() *Quickswapv3Router {
	parsed, err := Quickswapv3RouterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Quickswapv3Router{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Quickswapv3Router) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _factory, address _WNativeToken, address _poolDeployer) returns()
func (quickswapv3Router *Quickswapv3Router) PackConstructor(_factory common.Address, _WNativeToken common.Address, _poolDeployer common.Address) []byte {
	enc, err := quickswapv3Router.abi.Pack("", _factory, _WNativeToken, _poolDeployer)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWNativeToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8af3ac85.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function WNativeToken() view returns(address)
func (quickswapv3Router *Quickswapv3Router) PackWNativeToken() []byte {
	enc, err := quickswapv3Router.abi.Pack("WNativeToken")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWNativeToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8af3ac85.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function WNativeToken() view returns(address)
func (quickswapv3Router *Quickswapv3Router) TryPackWNativeToken() ([]byte, error) {
	return quickswapv3Router.abi.Pack("WNativeToken")
}

// UnpackWNativeToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8af3ac85.
//
// Solidity: function WNativeToken() view returns(address)
func (quickswapv3Router *Quickswapv3Router) UnpackWNativeToken(data []byte) (common.Address, error) {
	out, err := quickswapv3Router.abi.Unpack("WNativeToken", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackAlgebraSwapCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c8958f6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (quickswapv3Router *Quickswapv3Router) PackAlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) []byte {
	enc, err := quickswapv3Router.abi.Pack("algebraSwapCallback", amount0Delta, amount1Delta, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAlgebraSwapCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c8958f6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (quickswapv3Router *Quickswapv3Router) TryPackAlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) ([]byte, error) {
	return quickswapv3Router.abi.Pack("algebraSwapCallback", amount0Delta, amount1Delta, data)
}

// PackExactInput is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc04b8d59.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (quickswapv3Router *Quickswapv3Router) PackExactInput(params ISwapRouterExactInputParams) []byte {
	enc, err := quickswapv3Router.abi.Pack("exactInput", params)
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
func (quickswapv3Router *Quickswapv3Router) TryPackExactInput(params ISwapRouterExactInputParams) ([]byte, error) {
	return quickswapv3Router.abi.Pack("exactInput", params)
}

// UnpackExactInput is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (quickswapv3Router *Quickswapv3Router) UnpackExactInput(data []byte) (*big.Int, error) {
	out, err := quickswapv3Router.abi.Unpack("exactInput", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackExactInputSingle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc651188.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (quickswapv3Router *Quickswapv3Router) PackExactInputSingle(params ISwapRouterExactInputSingleParams) []byte {
	enc, err := quickswapv3Router.abi.Pack("exactInputSingle", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExactInputSingle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc651188.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (quickswapv3Router *Quickswapv3Router) TryPackExactInputSingle(params ISwapRouterExactInputSingleParams) ([]byte, error) {
	return quickswapv3Router.abi.Pack("exactInputSingle", params)
}

// UnpackExactInputSingle is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc651188.
//
// Solidity: function exactInputSingle((address,address,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (quickswapv3Router *Quickswapv3Router) UnpackExactInputSingle(data []byte) (*big.Int, error) {
	out, err := quickswapv3Router.abi.Unpack("exactInputSingle", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackExactInputSingleSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb87d2524.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (quickswapv3Router *Quickswapv3Router) PackExactInputSingleSupportingFeeOnTransferTokens(params ISwapRouterExactInputSingleParams) []byte {
	enc, err := quickswapv3Router.abi.Pack("exactInputSingleSupportingFeeOnTransferTokens", params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExactInputSingleSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb87d2524.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (quickswapv3Router *Quickswapv3Router) TryPackExactInputSingleSupportingFeeOnTransferTokens(params ISwapRouterExactInputSingleParams) ([]byte, error) {
	return quickswapv3Router.abi.Pack("exactInputSingleSupportingFeeOnTransferTokens", params)
}

// UnpackExactInputSingleSupportingFeeOnTransferTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb87d2524.
//
// Solidity: function exactInputSingleSupportingFeeOnTransferTokens((address,address,address,uint256,uint256,uint256,uint160) params) returns(uint256 amountOut)
func (quickswapv3Router *Quickswapv3Router) UnpackExactInputSingleSupportingFeeOnTransferTokens(data []byte) (*big.Int, error) {
	out, err := quickswapv3Router.abi.Unpack("exactInputSingleSupportingFeeOnTransferTokens", data)
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
func (quickswapv3Router *Quickswapv3Router) PackExactOutput(params ISwapRouterExactOutputParams) []byte {
	enc, err := quickswapv3Router.abi.Pack("exactOutput", params)
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
func (quickswapv3Router *Quickswapv3Router) TryPackExactOutput(params ISwapRouterExactOutputParams) ([]byte, error) {
	return quickswapv3Router.abi.Pack("exactOutput", params)
}

// UnpackExactOutput is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (quickswapv3Router *Quickswapv3Router) UnpackExactOutput(data []byte) (*big.Int, error) {
	out, err := quickswapv3Router.abi.Unpack("exactOutput", data)
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
func (quickswapv3Router *Quickswapv3Router) PackExactOutputSingle(params ISwapRouterExactOutputSingleParams) []byte {
	enc, err := quickswapv3Router.abi.Pack("exactOutputSingle", params)
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
func (quickswapv3Router *Quickswapv3Router) TryPackExactOutputSingle(params ISwapRouterExactOutputSingleParams) ([]byte, error) {
	return quickswapv3Router.abi.Pack("exactOutputSingle", params)
}

// UnpackExactOutputSingle is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (quickswapv3Router *Quickswapv3Router) UnpackExactOutputSingle(data []byte) (*big.Int, error) {
	out, err := quickswapv3Router.abi.Unpack("exactOutputSingle", data)
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
func (quickswapv3Router *Quickswapv3Router) PackFactory() []byte {
	enc, err := quickswapv3Router.abi.Pack("factory")
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
func (quickswapv3Router *Quickswapv3Router) TryPackFactory() ([]byte, error) {
	return quickswapv3Router.abi.Pack("factory")
}

// UnpackFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (quickswapv3Router *Quickswapv3Router) UnpackFactory(data []byte) (common.Address, error) {
	out, err := quickswapv3Router.abi.Unpack("factory", data)
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
func (quickswapv3Router *Quickswapv3Router) PackMulticall(data [][]byte) []byte {
	enc, err := quickswapv3Router.abi.Pack("multicall", data)
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
func (quickswapv3Router *Quickswapv3Router) TryPackMulticall(data [][]byte) ([]byte, error) {
	return quickswapv3Router.abi.Pack("multicall", data)
}

// UnpackMulticall is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (quickswapv3Router *Quickswapv3Router) UnpackMulticall(data []byte) ([][]byte, error) {
	out, err := quickswapv3Router.abi.Unpack("multicall", data)
	if err != nil {
		return *new([][]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	return out0, nil
}

// PackPoolDeployer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3119049a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function poolDeployer() view returns(address)
func (quickswapv3Router *Quickswapv3Router) PackPoolDeployer() []byte {
	enc, err := quickswapv3Router.abi.Pack("poolDeployer")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPoolDeployer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3119049a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function poolDeployer() view returns(address)
func (quickswapv3Router *Quickswapv3Router) TryPackPoolDeployer() ([]byte, error) {
	return quickswapv3Router.abi.Pack("poolDeployer")
}

// UnpackPoolDeployer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (quickswapv3Router *Quickswapv3Router) UnpackPoolDeployer(data []byte) (common.Address, error) {
	out, err := quickswapv3Router.abi.Unpack("poolDeployer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRefundNativeToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41865270.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function refundNativeToken() payable returns()
func (quickswapv3Router *Quickswapv3Router) PackRefundNativeToken() []byte {
	enc, err := quickswapv3Router.abi.Pack("refundNativeToken")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRefundNativeToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41865270.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function refundNativeToken() payable returns()
func (quickswapv3Router *Quickswapv3Router) TryPackRefundNativeToken() ([]byte, error) {
	return quickswapv3Router.abi.Pack("refundNativeToken")
}

// PackSelfPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3995c67.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (quickswapv3Router *Quickswapv3Router) PackSelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := quickswapv3Router.abi.Pack("selfPermit", token, value, deadline, v, r, s)
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
func (quickswapv3Router *Quickswapv3Router) TryPackSelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return quickswapv3Router.abi.Pack("selfPermit", token, value, deadline, v, r, s)
}

// PackSelfPermitAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4659a494.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (quickswapv3Router *Quickswapv3Router) PackSelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := quickswapv3Router.abi.Pack("selfPermitAllowed", token, nonce, expiry, v, r, s)
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
func (quickswapv3Router *Quickswapv3Router) TryPackSelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return quickswapv3Router.abi.Pack("selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// PackSelfPermitAllowedIfNecessary is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4a78f0c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (quickswapv3Router *Quickswapv3Router) PackSelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := quickswapv3Router.abi.Pack("selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
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
func (quickswapv3Router *Quickswapv3Router) TryPackSelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return quickswapv3Router.abi.Pack("selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// PackSelfPermitIfNecessary is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc2e3140a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (quickswapv3Router *Quickswapv3Router) PackSelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := quickswapv3Router.abi.Pack("selfPermitIfNecessary", token, value, deadline, v, r, s)
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
func (quickswapv3Router *Quickswapv3Router) TryPackSelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return quickswapv3Router.abi.Pack("selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// PackSweepToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdf2ab5bb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (quickswapv3Router *Quickswapv3Router) PackSweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) []byte {
	enc, err := quickswapv3Router.abi.Pack("sweepToken", token, amountMinimum, recipient)
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
func (quickswapv3Router *Quickswapv3Router) TryPackSweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) ([]byte, error) {
	return quickswapv3Router.abi.Pack("sweepToken", token, amountMinimum, recipient)
}

// PackSweepTokenWithFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe0e189a0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (quickswapv3Router *Quickswapv3Router) PackSweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) []byte {
	enc, err := quickswapv3Router.abi.Pack("sweepTokenWithFee", token, amountMinimum, recipient, feeBips, feeRecipient)
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
func (quickswapv3Router *Quickswapv3Router) TryPackSweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) ([]byte, error) {
	return quickswapv3Router.abi.Pack("sweepTokenWithFee", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// PackUnwrapWNativeToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x69bc35b2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unwrapWNativeToken(uint256 amountMinimum, address recipient) payable returns()
func (quickswapv3Router *Quickswapv3Router) PackUnwrapWNativeToken(amountMinimum *big.Int, recipient common.Address) []byte {
	enc, err := quickswapv3Router.abi.Pack("unwrapWNativeToken", amountMinimum, recipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnwrapWNativeToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x69bc35b2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unwrapWNativeToken(uint256 amountMinimum, address recipient) payable returns()
func (quickswapv3Router *Quickswapv3Router) TryPackUnwrapWNativeToken(amountMinimum *big.Int, recipient common.Address) ([]byte, error) {
	return quickswapv3Router.abi.Pack("unwrapWNativeToken", amountMinimum, recipient)
}

// PackUnwrapWNativeTokenWithFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc60696ec.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unwrapWNativeTokenWithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (quickswapv3Router *Quickswapv3Router) PackUnwrapWNativeTokenWithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) []byte {
	enc, err := quickswapv3Router.abi.Pack("unwrapWNativeTokenWithFee", amountMinimum, recipient, feeBips, feeRecipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnwrapWNativeTokenWithFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc60696ec.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unwrapWNativeTokenWithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (quickswapv3Router *Quickswapv3Router) TryPackUnwrapWNativeTokenWithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) ([]byte, error) {
	return quickswapv3Router.abi.Pack("unwrapWNativeTokenWithFee", amountMinimum, recipient, feeBips, feeRecipient)
}
