// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package joe_router

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

// JoeRouterMetaData contains all meta data concerning the JoeRouter contract.
var JoeRouterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WAVAX\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WAVAX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityAVAXSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityAVAXWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAVAXForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAXSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "JoeRouter",
}

// JoeRouter is an auto generated Go binding around an Ethereum contract.
type JoeRouter struct {
	abi abi.ABI
}

// NewJoeRouter creates a new instance of JoeRouter.
func NewJoeRouter() *JoeRouter {
	parsed, err := JoeRouterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &JoeRouter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *JoeRouter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _factory, address _WAVAX) returns()
func (joeRouter *JoeRouter) PackConstructor(_factory common.Address, _WAVAX common.Address) []byte {
	enc, err := joeRouter.abi.Pack("", _factory, _WAVAX)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWAVAX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x73b295c2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function WAVAX() view returns(address)
func (joeRouter *JoeRouter) PackWAVAX() []byte {
	enc, err := joeRouter.abi.Pack("WAVAX")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWAVAX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x73b295c2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function WAVAX() view returns(address)
func (joeRouter *JoeRouter) TryPackWAVAX() ([]byte, error) {
	return joeRouter.abi.Pack("WAVAX")
}

// UnpackWAVAX is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x73b295c2.
//
// Solidity: function WAVAX() view returns(address)
func (joeRouter *JoeRouter) UnpackWAVAX(data []byte) (common.Address, error) {
	out, err := joeRouter.abi.Unpack("WAVAX", data)
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
func (joeRouter *JoeRouter) PackAddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
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
func (joeRouter *JoeRouter) TryPackAddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
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
func (joeRouter *JoeRouter) UnpackAddLiquidity(data []byte) (AddLiquidityOutput, error) {
	out, err := joeRouter.abi.Unpack("addLiquidity", data)
	outstruct := new(AddLiquidityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Liquidity = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackAddLiquidityAVAX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf91b3f72.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (joeRouter *JoeRouter) PackAddLiquidityAVAX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("addLiquidityAVAX", token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddLiquidityAVAX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf91b3f72.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (joeRouter *JoeRouter) TryPackAddLiquidityAVAX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("addLiquidityAVAX", token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
}

// AddLiquidityAVAXOutput serves as a container for the return parameters of contract
// method AddLiquidityAVAX.
type AddLiquidityAVAXOutput struct {
	AmountToken *big.Int
	AmountAVAX  *big.Int
	Liquidity   *big.Int
}

// UnpackAddLiquidityAVAX is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (joeRouter *JoeRouter) UnpackAddLiquidityAVAX(data []byte) (AddLiquidityAVAXOutput, error) {
	out, err := joeRouter.abi.Unpack("addLiquidityAVAX", data)
	outstruct := new(AddLiquidityAVAXOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountToken = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountAVAX = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Liquidity = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc45a0155.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function factory() view returns(address)
func (joeRouter *JoeRouter) PackFactory() []byte {
	enc, err := joeRouter.abi.Pack("factory")
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
func (joeRouter *JoeRouter) TryPackFactory() ([]byte, error) {
	return joeRouter.abi.Pack("factory")
}

// UnpackFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (joeRouter *JoeRouter) UnpackFactory(data []byte) (common.Address, error) {
	out, err := joeRouter.abi.Unpack("factory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAmountIn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85f8c259.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (joeRouter *JoeRouter) PackGetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("getAmountIn", amountOut, reserveIn, reserveOut)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAmountIn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85f8c259.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (joeRouter *JoeRouter) TryPackGetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("getAmountIn", amountOut, reserveIn, reserveOut)
}

// UnpackGetAmountIn is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (joeRouter *JoeRouter) UnpackGetAmountIn(data []byte) (*big.Int, error) {
	out, err := joeRouter.abi.Unpack("getAmountIn", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetAmountOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x054d50d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (joeRouter *JoeRouter) PackGetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("getAmountOut", amountIn, reserveIn, reserveOut)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAmountOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x054d50d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (joeRouter *JoeRouter) TryPackGetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("getAmountOut", amountIn, reserveIn, reserveOut)
}

// UnpackGetAmountOut is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (joeRouter *JoeRouter) UnpackGetAmountOut(data []byte) (*big.Int, error) {
	out, err := joeRouter.abi.Unpack("getAmountOut", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetAmountsIn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f00ca74.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (joeRouter *JoeRouter) PackGetAmountsIn(amountOut *big.Int, path []common.Address) []byte {
	enc, err := joeRouter.abi.Pack("getAmountsIn", amountOut, path)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAmountsIn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f00ca74.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (joeRouter *JoeRouter) TryPackGetAmountsIn(amountOut *big.Int, path []common.Address) ([]byte, error) {
	return joeRouter.abi.Pack("getAmountsIn", amountOut, path)
}

// UnpackGetAmountsIn is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (joeRouter *JoeRouter) UnpackGetAmountsIn(data []byte) ([]*big.Int, error) {
	out, err := joeRouter.abi.Unpack("getAmountsIn", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackGetAmountsOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd06ca61f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (joeRouter *JoeRouter) PackGetAmountsOut(amountIn *big.Int, path []common.Address) []byte {
	enc, err := joeRouter.abi.Pack("getAmountsOut", amountIn, path)
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
func (joeRouter *JoeRouter) TryPackGetAmountsOut(amountIn *big.Int, path []common.Address) ([]byte, error) {
	return joeRouter.abi.Pack("getAmountsOut", amountIn, path)
}

// UnpackGetAmountsOut is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (joeRouter *JoeRouter) UnpackGetAmountsOut(data []byte) ([]*big.Int, error) {
	out, err := joeRouter.abi.Unpack("getAmountsOut", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackQuote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad615dec.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (joeRouter *JoeRouter) PackQuote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("quote", amountA, reserveA, reserveB)
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
func (joeRouter *JoeRouter) TryPackQuote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("quote", amountA, reserveA, reserveB)
}

// UnpackQuote is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (joeRouter *JoeRouter) UnpackQuote(data []byte) (*big.Int, error) {
	out, err := joeRouter.abi.Unpack("quote", data)
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
func (joeRouter *JoeRouter) PackRemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
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
func (joeRouter *JoeRouter) TryPackRemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
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
func (joeRouter *JoeRouter) UnpackRemoveLiquidity(data []byte) (RemoveLiquidityOutput, error) {
	out, err := joeRouter.abi.Unpack("removeLiquidity", data)
	outstruct := new(RemoveLiquidityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackRemoveLiquidityAVAX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x33c6b725.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (joeRouter *JoeRouter) PackRemoveLiquidityAVAX(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("removeLiquidityAVAX", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidityAVAX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x33c6b725.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (joeRouter *JoeRouter) TryPackRemoveLiquidityAVAX(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("removeLiquidityAVAX", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXOutput serves as a container for the return parameters of contract
// method RemoveLiquidityAVAX.
type RemoveLiquidityAVAXOutput struct {
	AmountToken *big.Int
	AmountAVAX  *big.Int
}

// UnpackRemoveLiquidityAVAX is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x33c6b725.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (joeRouter *JoeRouter) UnpackRemoveLiquidityAVAX(data []byte) (RemoveLiquidityAVAXOutput, error) {
	out, err := joeRouter.abi.Unpack("removeLiquidityAVAX", data)
	outstruct := new(RemoveLiquidityAVAXOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountToken = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountAVAX = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackRemoveLiquidityAVAXSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x73bc79cf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (joeRouter *JoeRouter) PackRemoveLiquidityAVAXSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("removeLiquidityAVAXSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidityAVAXSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x73bc79cf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (joeRouter *JoeRouter) TryPackRemoveLiquidityAVAXSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("removeLiquidityAVAXSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// UnpackRemoveLiquidityAVAXSupportingFeeOnTransferTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x73bc79cf.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (joeRouter *JoeRouter) UnpackRemoveLiquidityAVAXSupportingFeeOnTransferTokens(data []byte) (*big.Int, error) {
	out, err := joeRouter.abi.Unpack("removeLiquidityAVAXSupportingFeeOnTransferTokens", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRemoveLiquidityAVAXWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c407024.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (joeRouter *JoeRouter) PackRemoveLiquidityAVAXWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := joeRouter.abi.Pack("removeLiquidityAVAXWithPermit", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidityAVAXWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c407024.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (joeRouter *JoeRouter) TryPackRemoveLiquidityAVAXWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return joeRouter.abi.Pack("removeLiquidityAVAXWithPermit", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermitOutput serves as a container for the return parameters of contract
// method RemoveLiquidityAVAXWithPermit.
type RemoveLiquidityAVAXWithPermitOutput struct {
	AmountToken *big.Int
	AmountAVAX  *big.Int
}

// UnpackRemoveLiquidityAVAXWithPermit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2c407024.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (joeRouter *JoeRouter) UnpackRemoveLiquidityAVAXWithPermit(data []byte) (RemoveLiquidityAVAXWithPermitOutput, error) {
	out, err := joeRouter.abi.Unpack("removeLiquidityAVAXWithPermit", data)
	outstruct := new(RemoveLiquidityAVAXWithPermitOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountToken = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountAVAX = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackRemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fc27226.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (joeRouter *JoeRouter) PackRemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := joeRouter.abi.Pack("removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fc27226.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (joeRouter *JoeRouter) TryPackRemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return joeRouter.abi.Pack("removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// UnpackRemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9fc27226.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (joeRouter *JoeRouter) UnpackRemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(data []byte) (*big.Int, error) {
	out, err := joeRouter.abi.Unpack("removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens", data)
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
func (joeRouter *JoeRouter) PackRemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) []byte {
	enc, err := joeRouter.abi.Pack("removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
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
func (joeRouter *JoeRouter) TryPackRemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) ([]byte, error) {
	return joeRouter.abi.Pack("removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
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
func (joeRouter *JoeRouter) UnpackRemoveLiquidityWithPermit(data []byte) (RemoveLiquidityWithPermitOutput, error) {
	out, err := joeRouter.abi.Unpack("removeLiquidityWithPermit", data)
	outstruct := new(RemoveLiquidityWithPermitOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackSwapAVAXForExactTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a657e67.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (joeRouter *JoeRouter) PackSwapAVAXForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("swapAVAXForExactTokens", amountOut, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapAVAXForExactTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a657e67.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (joeRouter *JoeRouter) TryPackSwapAVAXForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("swapAVAXForExactTokens", amountOut, path, to, deadline)
}

// UnpackSwapAVAXForExactTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (joeRouter *JoeRouter) UnpackSwapAVAXForExactTokens(data []byte) ([]*big.Int, error) {
	out, err := joeRouter.abi.Unpack("swapAVAXForExactTokens", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackSwapExactAVAXForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa2a1623d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (joeRouter *JoeRouter) PackSwapExactAVAXForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("swapExactAVAXForTokens", amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactAVAXForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa2a1623d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (joeRouter *JoeRouter) TryPackSwapExactAVAXForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("swapExactAVAXForTokens", amountOutMin, path, to, deadline)
}

// UnpackSwapExactAVAXForTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (joeRouter *JoeRouter) UnpackSwapExactAVAXForTokens(data []byte) ([]*big.Int, error) {
	out, err := joeRouter.abi.Unpack("swapExactAVAXForTokens", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackSwapExactAVAXForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc57559dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (joeRouter *JoeRouter) PackSwapExactAVAXForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("swapExactAVAXForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactAVAXForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc57559dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (joeRouter *JoeRouter) TryPackSwapExactAVAXForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("swapExactAVAXForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// PackSwapExactTokensForAVAX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x676528d1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) PackSwapExactTokensForAVAX(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("swapExactTokensForAVAX", amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactTokensForAVAX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x676528d1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) TryPackSwapExactTokensForAVAX(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("swapExactTokensForAVAX", amountIn, amountOutMin, path, to, deadline)
}

// UnpackSwapExactTokensForAVAX is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) UnpackSwapExactTokensForAVAX(data []byte) ([]*big.Int, error) {
	out, err := joeRouter.abi.Unpack("swapExactTokensForAVAX", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackSwapExactTokensForAVAXSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x762b1562.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (joeRouter *JoeRouter) PackSwapExactTokensForAVAXSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("swapExactTokensForAVAXSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactTokensForAVAXSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x762b1562.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (joeRouter *JoeRouter) TryPackSwapExactTokensForAVAXSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("swapExactTokensForAVAXSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// PackSwapExactTokensForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38ed1739.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) PackSwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactTokensForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38ed1739.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) TryPackSwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// UnpackSwapExactTokensForTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) UnpackSwapExactTokensForTokens(data []byte) ([]*big.Int, error) {
	out, err := joeRouter.abi.Unpack("swapExactTokensForTokens", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackSwapExactTokensForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c11d795.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (joeRouter *JoeRouter) PackSwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactTokensForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c11d795.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (joeRouter *JoeRouter) TryPackSwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// PackSwapTokensForExactAVAX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a42416a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) PackSwapTokensForExactAVAX(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("swapTokensForExactAVAX", amountOut, amountInMax, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapTokensForExactAVAX is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a42416a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) TryPackSwapTokensForExactAVAX(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("swapTokensForExactAVAX", amountOut, amountInMax, path, to, deadline)
}

// UnpackSwapTokensForExactAVAX is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) UnpackSwapTokensForExactAVAX(data []byte) ([]*big.Int, error) {
	out, err := joeRouter.abi.Unpack("swapTokensForExactAVAX", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackSwapTokensForExactTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8803dbee.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) PackSwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) []byte {
	enc, err := joeRouter.abi.Pack("swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapTokensForExactTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8803dbee.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) TryPackSwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) ([]byte, error) {
	return joeRouter.abi.Pack("swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// UnpackSwapTokensForExactTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (joeRouter *JoeRouter) UnpackSwapTokensForExactTokens(data []byte) ([]*big.Int, error) {
	out, err := joeRouter.abi.Unpack("swapTokensForExactTokens", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}
