// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aerodrome_factory

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

// AerodromeFactoryMetaData contains all meta data concerning the AerodromeFactory contract.
var AerodromeFactoryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FeeInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFeeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVoter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFee\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SetCustomFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeManager\",\"type\":\"address\"}],\"name\":\"SetFeeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"SetPauseState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"SetPauser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"SetVoter\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_FEE_INDICATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"customFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_stable\",\"type\":\"bool\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setCustomFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeManager\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setPauseState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"setPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"setVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stableFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"volatileFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "AerodromeFactory",
}

// AerodromeFactory is an auto generated Go binding around an Ethereum contract.
type AerodromeFactory struct {
	abi abi.ABI
}

// NewAerodromeFactory creates a new instance of AerodromeFactory.
func NewAerodromeFactory() *AerodromeFactory {
	parsed, err := AerodromeFactoryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AerodromeFactory{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AerodromeFactory) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _implementation) returns()
func (aerodromeFactory *AerodromeFactory) PackConstructor(_implementation common.Address) []byte {
	enc, err := aerodromeFactory.abi.Pack("", _implementation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMAXFEE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc063e1a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) PackMAXFEE() []byte {
	enc, err := aerodromeFactory.abi.Pack("MAX_FEE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMAXFEE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc063e1a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) TryPackMAXFEE() ([]byte, error) {
	return aerodromeFactory.abi.Pack("MAX_FEE")
}

// UnpackMAXFEE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) UnpackMAXFEE(data []byte) (*big.Int, error) {
	out, err := aerodromeFactory.abi.Unpack("MAX_FEE", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackZEROFEEINDICATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38c55d46.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ZERO_FEE_INDICATOR() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) PackZEROFEEINDICATOR() []byte {
	enc, err := aerodromeFactory.abi.Pack("ZERO_FEE_INDICATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackZEROFEEINDICATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38c55d46.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ZERO_FEE_INDICATOR() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) TryPackZEROFEEINDICATOR() ([]byte, error) {
	return aerodromeFactory.abi.Pack("ZERO_FEE_INDICATOR")
}

// UnpackZEROFEEINDICATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x38c55d46.
//
// Solidity: function ZERO_FEE_INDICATOR() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) UnpackZEROFEEINDICATOR(data []byte) (*big.Int, error) {
	out, err := aerodromeFactory.abi.Unpack("ZERO_FEE_INDICATOR", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackAllPools is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41d1de97.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (aerodromeFactory *AerodromeFactory) PackAllPools(arg0 *big.Int) []byte {
	enc, err := aerodromeFactory.abi.Pack("allPools", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllPools is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41d1de97.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (aerodromeFactory *AerodromeFactory) TryPackAllPools(arg0 *big.Int) ([]byte, error) {
	return aerodromeFactory.abi.Pack("allPools", arg0)
}

// UnpackAllPools is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (aerodromeFactory *AerodromeFactory) UnpackAllPools(data []byte) (common.Address, error) {
	out, err := aerodromeFactory.abi.Unpack("allPools", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackAllPoolsLength is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xefde4e64.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) PackAllPoolsLength() []byte {
	enc, err := aerodromeFactory.abi.Pack("allPoolsLength")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllPoolsLength is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xefde4e64.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) TryPackAllPoolsLength() ([]byte, error) {
	return aerodromeFactory.abi.Pack("allPoolsLength")
}

// UnpackAllPoolsLength is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) UnpackAllPoolsLength(data []byte) (*big.Int, error) {
	out, err := aerodromeFactory.abi.Unpack("allPoolsLength", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCreatePool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36bf95a0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createPool(address tokenA, address tokenB, bool stable) returns(address pool)
func (aerodromeFactory *AerodromeFactory) PackCreatePool(tokenA common.Address, tokenB common.Address, stable bool) []byte {
	enc, err := aerodromeFactory.abi.Pack("createPool", tokenA, tokenB, stable)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreatePool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36bf95a0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createPool(address tokenA, address tokenB, bool stable) returns(address pool)
func (aerodromeFactory *AerodromeFactory) TryPackCreatePool(tokenA common.Address, tokenB common.Address, stable bool) ([]byte, error) {
	return aerodromeFactory.abi.Pack("createPool", tokenA, tokenB, stable)
}

// UnpackCreatePool is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x36bf95a0.
//
// Solidity: function createPool(address tokenA, address tokenB, bool stable) returns(address pool)
func (aerodromeFactory *AerodromeFactory) UnpackCreatePool(data []byte) (common.Address, error) {
	out, err := aerodromeFactory.abi.Unpack("createPool", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCreatePool0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1671295.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (aerodromeFactory *AerodromeFactory) PackCreatePool0(tokenA common.Address, tokenB common.Address, fee *big.Int) []byte {
	enc, err := aerodromeFactory.abi.Pack("createPool0", tokenA, tokenB, fee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreatePool0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1671295.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (aerodromeFactory *AerodromeFactory) TryPackCreatePool0(tokenA common.Address, tokenB common.Address, fee *big.Int) ([]byte, error) {
	return aerodromeFactory.abi.Pack("createPool0", tokenA, tokenB, fee)
}

// UnpackCreatePool0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (aerodromeFactory *AerodromeFactory) UnpackCreatePool0(data []byte) (common.Address, error) {
	out, err := aerodromeFactory.abi.Unpack("createPool0", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCustomFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d419abc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function customFee(address ) view returns(uint256)
func (aerodromeFactory *AerodromeFactory) PackCustomFee(arg0 common.Address) []byte {
	enc, err := aerodromeFactory.abi.Pack("customFee", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCustomFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d419abc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function customFee(address ) view returns(uint256)
func (aerodromeFactory *AerodromeFactory) TryPackCustomFee(arg0 common.Address) ([]byte, error) {
	return aerodromeFactory.abi.Pack("customFee", arg0)
}

// UnpackCustomFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4d419abc.
//
// Solidity: function customFee(address ) view returns(uint256)
func (aerodromeFactory *AerodromeFactory) UnpackCustomFee(data []byte) (*big.Int, error) {
	out, err := aerodromeFactory.abi.Unpack("customFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackFeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd0fb0203.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function feeManager() view returns(address)
func (aerodromeFactory *AerodromeFactory) PackFeeManager() []byte {
	enc, err := aerodromeFactory.abi.Pack("feeManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd0fb0203.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function feeManager() view returns(address)
func (aerodromeFactory *AerodromeFactory) TryPackFeeManager() ([]byte, error) {
	return aerodromeFactory.abi.Pack("feeManager")
}

// UnpackFeeManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (aerodromeFactory *AerodromeFactory) UnpackFeeManager(data []byte) (common.Address, error) {
	out, err := aerodromeFactory.abi.Unpack("feeManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcc56b2c5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getFee(address pool, bool _stable) view returns(uint256)
func (aerodromeFactory *AerodromeFactory) PackGetFee(pool common.Address, stable bool) []byte {
	enc, err := aerodromeFactory.abi.Pack("getFee", pool, stable)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcc56b2c5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getFee(address pool, bool _stable) view returns(uint256)
func (aerodromeFactory *AerodromeFactory) TryPackGetFee(pool common.Address, stable bool) ([]byte, error) {
	return aerodromeFactory.abi.Pack("getFee", pool, stable)
}

// UnpackGetFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc56b2c5.
//
// Solidity: function getFee(address pool, bool _stable) view returns(uint256)
func (aerodromeFactory *AerodromeFactory) UnpackGetFee(data []byte) (*big.Int, error) {
	out, err := aerodromeFactory.abi.Unpack("getFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetPool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1698ee82.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address)
func (aerodromeFactory *AerodromeFactory) PackGetPool(tokenA common.Address, tokenB common.Address, fee *big.Int) []byte {
	enc, err := aerodromeFactory.abi.Pack("getPool", tokenA, tokenB, fee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1698ee82.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address)
func (aerodromeFactory *AerodromeFactory) TryPackGetPool(tokenA common.Address, tokenB common.Address, fee *big.Int) ([]byte, error) {
	return aerodromeFactory.abi.Pack("getPool", tokenA, tokenB, fee)
}

// UnpackGetPool is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1698ee82.
//
// Solidity: function getPool(address tokenA, address tokenB, uint24 fee) view returns(address)
func (aerodromeFactory *AerodromeFactory) UnpackGetPool(data []byte) (common.Address, error) {
	out, err := aerodromeFactory.abi.Unpack("getPool", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetPool0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79bc57d5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address)
func (aerodromeFactory *AerodromeFactory) PackGetPool0(tokenA common.Address, tokenB common.Address, stable bool) []byte {
	enc, err := aerodromeFactory.abi.Pack("getPool0", tokenA, tokenB, stable)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPool0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79bc57d5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address)
func (aerodromeFactory *AerodromeFactory) TryPackGetPool0(tokenA common.Address, tokenB common.Address, stable bool) ([]byte, error) {
	return aerodromeFactory.abi.Pack("getPool0", tokenA, tokenB, stable)
}

// UnpackGetPool0 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x79bc57d5.
//
// Solidity: function getPool(address tokenA, address tokenB, bool stable) view returns(address)
func (aerodromeFactory *AerodromeFactory) UnpackGetPool0(data []byte) (common.Address, error) {
	out, err := aerodromeFactory.abi.Unpack("getPool0", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c60da1b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function implementation() view returns(address)
func (aerodromeFactory *AerodromeFactory) PackImplementation() []byte {
	enc, err := aerodromeFactory.abi.Pack("implementation")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c60da1b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function implementation() view returns(address)
func (aerodromeFactory *AerodromeFactory) TryPackImplementation() ([]byte, error) {
	return aerodromeFactory.abi.Pack("implementation")
}

// UnpackImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (aerodromeFactory *AerodromeFactory) UnpackImplementation(data []byte) (common.Address, error) {
	out, err := aerodromeFactory.abi.Unpack("implementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb187bd26.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isPaused() view returns(bool)
func (aerodromeFactory *AerodromeFactory) PackIsPaused() []byte {
	enc, err := aerodromeFactory.abi.Pack("isPaused")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb187bd26.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isPaused() view returns(bool)
func (aerodromeFactory *AerodromeFactory) TryPackIsPaused() ([]byte, error) {
	return aerodromeFactory.abi.Pack("isPaused")
}

// UnpackIsPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (aerodromeFactory *AerodromeFactory) UnpackIsPaused(data []byte) (bool, error) {
	out, err := aerodromeFactory.abi.Unpack("isPaused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsPool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b16ebb7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isPool(address pool) view returns(bool)
func (aerodromeFactory *AerodromeFactory) PackIsPool(pool common.Address) []byte {
	enc, err := aerodromeFactory.abi.Pack("isPool", pool)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsPool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b16ebb7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isPool(address pool) view returns(bool)
func (aerodromeFactory *AerodromeFactory) TryPackIsPool(pool common.Address) ([]byte, error) {
	return aerodromeFactory.abi.Pack("isPool", pool)
}

// UnpackIsPool is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b16ebb7.
//
// Solidity: function isPool(address pool) view returns(bool)
func (aerodromeFactory *AerodromeFactory) UnpackIsPool(data []byte) (bool, error) {
	out, err := aerodromeFactory.abi.Unpack("isPool", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fd0506d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pauser() view returns(address)
func (aerodromeFactory *AerodromeFactory) PackPauser() []byte {
	enc, err := aerodromeFactory.abi.Pack("pauser")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fd0506d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pauser() view returns(address)
func (aerodromeFactory *AerodromeFactory) TryPackPauser() ([]byte, error) {
	return aerodromeFactory.abi.Pack("pauser")
}

// UnpackPauser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (aerodromeFactory *AerodromeFactory) UnpackPauser(data []byte) (common.Address, error) {
	out, err := aerodromeFactory.abi.Unpack("pauser", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetCustomFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd49466a8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setCustomFee(address pool, uint256 fee) returns()
func (aerodromeFactory *AerodromeFactory) PackSetCustomFee(pool common.Address, fee *big.Int) []byte {
	enc, err := aerodromeFactory.abi.Pack("setCustomFee", pool, fee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetCustomFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd49466a8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setCustomFee(address pool, uint256 fee) returns()
func (aerodromeFactory *AerodromeFactory) TryPackSetCustomFee(pool common.Address, fee *big.Int) ([]byte, error) {
	return aerodromeFactory.abi.Pack("setCustomFee", pool, fee)
}

// PackSetFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1f76b44.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFee(bool _stable, uint256 _fee) returns()
func (aerodromeFactory *AerodromeFactory) PackSetFee(stable bool, fee *big.Int) []byte {
	enc, err := aerodromeFactory.abi.Pack("setFee", stable, fee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe1f76b44.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFee(bool _stable, uint256 _fee) returns()
func (aerodromeFactory *AerodromeFactory) TryPackSetFee(stable bool, fee *big.Int) ([]byte, error) {
	return aerodromeFactory.abi.Pack("setFee", stable, fee)
}

// PackSetFeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x472d35b9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (aerodromeFactory *AerodromeFactory) PackSetFeeManager(feeManager common.Address) []byte {
	enc, err := aerodromeFactory.abi.Pack("setFeeManager", feeManager)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x472d35b9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFeeManager(address _feeManager) returns()
func (aerodromeFactory *AerodromeFactory) TryPackSetFeeManager(feeManager common.Address) ([]byte, error) {
	return aerodromeFactory.abi.Pack("setFeeManager", feeManager)
}

// PackSetPauseState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcdb88ad1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setPauseState(bool _state) returns()
func (aerodromeFactory *AerodromeFactory) PackSetPauseState(state bool) []byte {
	enc, err := aerodromeFactory.abi.Pack("setPauseState", state)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetPauseState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcdb88ad1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setPauseState(bool _state) returns()
func (aerodromeFactory *AerodromeFactory) TryPackSetPauseState(state bool) ([]byte, error) {
	return aerodromeFactory.abi.Pack("setPauseState", state)
}

// PackSetPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d88af4a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setPauser(address _pauser) returns()
func (aerodromeFactory *AerodromeFactory) PackSetPauser(pauser common.Address) []byte {
	enc, err := aerodromeFactory.abi.Pack("setPauser", pauser)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d88af4a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setPauser(address _pauser) returns()
func (aerodromeFactory *AerodromeFactory) TryPackSetPauser(pauser common.Address) ([]byte, error) {
	return aerodromeFactory.abi.Pack("setPauser", pauser)
}

// PackSetVoter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4bc2a657.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setVoter(address _voter) returns()
func (aerodromeFactory *AerodromeFactory) PackSetVoter(voter common.Address) []byte {
	enc, err := aerodromeFactory.abi.Pack("setVoter", voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetVoter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4bc2a657.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setVoter(address _voter) returns()
func (aerodromeFactory *AerodromeFactory) TryPackSetVoter(voter common.Address) ([]byte, error) {
	return aerodromeFactory.abi.Pack("setVoter", voter)
}

// PackStableFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40bbd775.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function stableFee() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) PackStableFee() []byte {
	enc, err := aerodromeFactory.abi.Pack("stableFee")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackStableFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40bbd775.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function stableFee() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) TryPackStableFee() ([]byte, error) {
	return aerodromeFactory.abi.Pack("stableFee")
}

// UnpackStableFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x40bbd775.
//
// Solidity: function stableFee() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) UnpackStableFee(data []byte) (*big.Int, error) {
	out, err := aerodromeFactory.abi.Unpack("stableFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackVolatileFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5084ed03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function volatileFee() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) PackVolatileFee() []byte {
	enc, err := aerodromeFactory.abi.Pack("volatileFee")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVolatileFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5084ed03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function volatileFee() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) TryPackVolatileFee() ([]byte, error) {
	return aerodromeFactory.abi.Pack("volatileFee")
}

// UnpackVolatileFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5084ed03.
//
// Solidity: function volatileFee() view returns(uint256)
func (aerodromeFactory *AerodromeFactory) UnpackVolatileFee(data []byte) (*big.Int, error) {
	out, err := aerodromeFactory.abi.Unpack("volatileFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackVoter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46c96aac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function voter() view returns(address)
func (aerodromeFactory *AerodromeFactory) PackVoter() []byte {
	enc, err := aerodromeFactory.abi.Pack("voter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVoter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46c96aac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function voter() view returns(address)
func (aerodromeFactory *AerodromeFactory) TryPackVoter() ([]byte, error) {
	return aerodromeFactory.abi.Pack("voter")
}

// UnpackVoter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (aerodromeFactory *AerodromeFactory) UnpackVoter(data []byte) (common.Address, error) {
	out, err := aerodromeFactory.abi.Unpack("voter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// AerodromeFactoryPoolCreated represents a PoolCreated event raised by the AerodromeFactory contract.
type AerodromeFactoryPoolCreated struct {
	Token0 common.Address
	Token1 common.Address
	Stable bool
	Pool   common.Address
	Arg4   *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const AerodromeFactoryPoolCreatedEventName = "PoolCreated"

// ContractEventName returns the user-defined event name.
func (AerodromeFactoryPoolCreated) ContractEventName() string {
	return AerodromeFactoryPoolCreatedEventName
}

// UnpackPoolCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, bool indexed stable, address pool, uint256 arg4)
func (aerodromeFactory *AerodromeFactory) UnpackPoolCreatedEvent(log *types.Log) (*AerodromeFactoryPoolCreated, error) {
	event := "PoolCreated"
	if log.Topics[0] != aerodromeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AerodromeFactoryPoolCreated)
	if len(log.Data) > 0 {
		if err := aerodromeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range aerodromeFactory.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// AerodromeFactorySetCustomFee represents a SetCustomFee event raised by the AerodromeFactory contract.
type AerodromeFactorySetCustomFee struct {
	Pool common.Address
	Fee  *big.Int
	Raw  *types.Log // Blockchain specific contextual infos
}

const AerodromeFactorySetCustomFeeEventName = "SetCustomFee"

// ContractEventName returns the user-defined event name.
func (AerodromeFactorySetCustomFee) ContractEventName() string {
	return AerodromeFactorySetCustomFeeEventName
}

// UnpackSetCustomFeeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SetCustomFee(address indexed pool, uint256 fee)
func (aerodromeFactory *AerodromeFactory) UnpackSetCustomFeeEvent(log *types.Log) (*AerodromeFactorySetCustomFee, error) {
	event := "SetCustomFee"
	if log.Topics[0] != aerodromeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AerodromeFactorySetCustomFee)
	if len(log.Data) > 0 {
		if err := aerodromeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range aerodromeFactory.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// AerodromeFactorySetFeeManager represents a SetFeeManager event raised by the AerodromeFactory contract.
type AerodromeFactorySetFeeManager struct {
	FeeManager common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const AerodromeFactorySetFeeManagerEventName = "SetFeeManager"

// ContractEventName returns the user-defined event name.
func (AerodromeFactorySetFeeManager) ContractEventName() string {
	return AerodromeFactorySetFeeManagerEventName
}

// UnpackSetFeeManagerEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SetFeeManager(address feeManager)
func (aerodromeFactory *AerodromeFactory) UnpackSetFeeManagerEvent(log *types.Log) (*AerodromeFactorySetFeeManager, error) {
	event := "SetFeeManager"
	if log.Topics[0] != aerodromeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AerodromeFactorySetFeeManager)
	if len(log.Data) > 0 {
		if err := aerodromeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range aerodromeFactory.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// AerodromeFactorySetPauseState represents a SetPauseState event raised by the AerodromeFactory contract.
type AerodromeFactorySetPauseState struct {
	State bool
	Raw   *types.Log // Blockchain specific contextual infos
}

const AerodromeFactorySetPauseStateEventName = "SetPauseState"

// ContractEventName returns the user-defined event name.
func (AerodromeFactorySetPauseState) ContractEventName() string {
	return AerodromeFactorySetPauseStateEventName
}

// UnpackSetPauseStateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SetPauseState(bool state)
func (aerodromeFactory *AerodromeFactory) UnpackSetPauseStateEvent(log *types.Log) (*AerodromeFactorySetPauseState, error) {
	event := "SetPauseState"
	if log.Topics[0] != aerodromeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AerodromeFactorySetPauseState)
	if len(log.Data) > 0 {
		if err := aerodromeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range aerodromeFactory.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// AerodromeFactorySetPauser represents a SetPauser event raised by the AerodromeFactory contract.
type AerodromeFactorySetPauser struct {
	Pauser common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const AerodromeFactorySetPauserEventName = "SetPauser"

// ContractEventName returns the user-defined event name.
func (AerodromeFactorySetPauser) ContractEventName() string {
	return AerodromeFactorySetPauserEventName
}

// UnpackSetPauserEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SetPauser(address pauser)
func (aerodromeFactory *AerodromeFactory) UnpackSetPauserEvent(log *types.Log) (*AerodromeFactorySetPauser, error) {
	event := "SetPauser"
	if log.Topics[0] != aerodromeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AerodromeFactorySetPauser)
	if len(log.Data) > 0 {
		if err := aerodromeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range aerodromeFactory.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// AerodromeFactorySetVoter represents a SetVoter event raised by the AerodromeFactory contract.
type AerodromeFactorySetVoter struct {
	Voter common.Address
	Raw   *types.Log // Blockchain specific contextual infos
}

const AerodromeFactorySetVoterEventName = "SetVoter"

// ContractEventName returns the user-defined event name.
func (AerodromeFactorySetVoter) ContractEventName() string {
	return AerodromeFactorySetVoterEventName
}

// UnpackSetVoterEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SetVoter(address voter)
func (aerodromeFactory *AerodromeFactory) UnpackSetVoterEvent(log *types.Log) (*AerodromeFactorySetVoter, error) {
	event := "SetVoter"
	if log.Topics[0] != aerodromeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(AerodromeFactorySetVoter)
	if len(log.Data) > 0 {
		if err := aerodromeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range aerodromeFactory.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (aerodromeFactory *AerodromeFactory) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], aerodromeFactory.abi.Errors["FeeInvalid"].ID.Bytes()[:4]) {
		return aerodromeFactory.UnpackFeeInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeFactory.abi.Errors["FeeTooHigh"].ID.Bytes()[:4]) {
		return aerodromeFactory.UnpackFeeTooHighError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeFactory.abi.Errors["InvalidPool"].ID.Bytes()[:4]) {
		return aerodromeFactory.UnpackInvalidPoolError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeFactory.abi.Errors["NotFeeManager"].ID.Bytes()[:4]) {
		return aerodromeFactory.UnpackNotFeeManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeFactory.abi.Errors["NotPauser"].ID.Bytes()[:4]) {
		return aerodromeFactory.UnpackNotPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeFactory.abi.Errors["NotVoter"].ID.Bytes()[:4]) {
		return aerodromeFactory.UnpackNotVoterError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeFactory.abi.Errors["PoolAlreadyExists"].ID.Bytes()[:4]) {
		return aerodromeFactory.UnpackPoolAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeFactory.abi.Errors["SameAddress"].ID.Bytes()[:4]) {
		return aerodromeFactory.UnpackSameAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeFactory.abi.Errors["ZeroAddress"].ID.Bytes()[:4]) {
		return aerodromeFactory.UnpackZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeFactory.abi.Errors["ZeroFee"].ID.Bytes()[:4]) {
		return aerodromeFactory.UnpackZeroFeeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// AerodromeFactoryFeeInvalid represents a FeeInvalid error raised by the AerodromeFactory contract.
type AerodromeFactoryFeeInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeInvalid()
func AerodromeFactoryFeeInvalidErrorID() common.Hash {
	return common.HexToHash("0x52dadcf9294ce72d94f3d13037a25e2ac6d04ed9976e67cfb21ac508d3c417bc")
}

// UnpackFeeInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeInvalid()
func (aerodromeFactory *AerodromeFactory) UnpackFeeInvalidError(raw []byte) (*AerodromeFactoryFeeInvalid, error) {
	out := new(AerodromeFactoryFeeInvalid)
	if err := aerodromeFactory.abi.UnpackIntoInterface(out, "FeeInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeFactoryFeeTooHigh represents a FeeTooHigh error raised by the AerodromeFactory contract.
type AerodromeFactoryFeeTooHigh struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeTooHigh()
func AerodromeFactoryFeeTooHighErrorID() common.Hash {
	return common.HexToHash("0xcd4e6167a0147beade9e7daca0e52cd42e992cd9c3dc1dd3ce8a2b6956f53601")
}

// UnpackFeeTooHighError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeTooHigh()
func (aerodromeFactory *AerodromeFactory) UnpackFeeTooHighError(raw []byte) (*AerodromeFactoryFeeTooHigh, error) {
	out := new(AerodromeFactoryFeeTooHigh)
	if err := aerodromeFactory.abi.UnpackIntoInterface(out, "FeeTooHigh", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeFactoryInvalidPool represents a InvalidPool error raised by the AerodromeFactory contract.
type AerodromeFactoryInvalidPool struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPool()
func AerodromeFactoryInvalidPoolErrorID() common.Hash {
	return common.HexToHash("0x2083cd4046029c86451c2e139a91b950c0f6ebc0dea58aa338e3029a5f151d99")
}

// UnpackInvalidPoolError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPool()
func (aerodromeFactory *AerodromeFactory) UnpackInvalidPoolError(raw []byte) (*AerodromeFactoryInvalidPool, error) {
	out := new(AerodromeFactoryInvalidPool)
	if err := aerodromeFactory.abi.UnpackIntoInterface(out, "InvalidPool", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeFactoryNotFeeManager represents a NotFeeManager error raised by the AerodromeFactory contract.
type AerodromeFactoryNotFeeManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotFeeManager()
func AerodromeFactoryNotFeeManagerErrorID() common.Hash {
	return common.HexToHash("0xf5d267ebbcd4925f9b76695d51986b620233a685bd12e9e1fac0bef372a72273")
}

// UnpackNotFeeManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotFeeManager()
func (aerodromeFactory *AerodromeFactory) UnpackNotFeeManagerError(raw []byte) (*AerodromeFactoryNotFeeManager, error) {
	out := new(AerodromeFactoryNotFeeManager)
	if err := aerodromeFactory.abi.UnpackIntoInterface(out, "NotFeeManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeFactoryNotPauser represents a NotPauser error raised by the AerodromeFactory contract.
type AerodromeFactoryNotPauser struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotPauser()
func AerodromeFactoryNotPauserErrorID() common.Hash {
	return common.HexToHash("0x492f6781990f089ef1c9cc82418ca3bcc261fbf6004c4e65cbb104af186dd9c0")
}

// UnpackNotPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotPauser()
func (aerodromeFactory *AerodromeFactory) UnpackNotPauserError(raw []byte) (*AerodromeFactoryNotPauser, error) {
	out := new(AerodromeFactoryNotPauser)
	if err := aerodromeFactory.abi.UnpackIntoInterface(out, "NotPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeFactoryNotVoter represents a NotVoter error raised by the AerodromeFactory contract.
type AerodromeFactoryNotVoter struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotVoter()
func AerodromeFactoryNotVoterErrorID() common.Hash {
	return common.HexToHash("0xc18384c1c33bcb495abe640b7d5186662eb8df3b8d2dca64b2cc26383f2fa337")
}

// UnpackNotVoterError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotVoter()
func (aerodromeFactory *AerodromeFactory) UnpackNotVoterError(raw []byte) (*AerodromeFactoryNotVoter, error) {
	out := new(AerodromeFactoryNotVoter)
	if err := aerodromeFactory.abi.UnpackIntoInterface(out, "NotVoter", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeFactoryPoolAlreadyExists represents a PoolAlreadyExists error raised by the AerodromeFactory contract.
type AerodromeFactoryPoolAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PoolAlreadyExists()
func AerodromeFactoryPoolAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x03119322497e04818eb069b67d4f395df8f307626d617ae15435b24b5f3bd20a")
}

// UnpackPoolAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PoolAlreadyExists()
func (aerodromeFactory *AerodromeFactory) UnpackPoolAlreadyExistsError(raw []byte) (*AerodromeFactoryPoolAlreadyExists, error) {
	out := new(AerodromeFactoryPoolAlreadyExists)
	if err := aerodromeFactory.abi.UnpackIntoInterface(out, "PoolAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeFactorySameAddress represents a SameAddress error raised by the AerodromeFactory contract.
type AerodromeFactorySameAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SameAddress()
func AerodromeFactorySameAddressErrorID() common.Hash {
	return common.HexToHash("0x367558c3b8d26eef2bf229cfbd1d28748768736e5b36de06552b152117e1be8e")
}

// UnpackSameAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SameAddress()
func (aerodromeFactory *AerodromeFactory) UnpackSameAddressError(raw []byte) (*AerodromeFactorySameAddress, error) {
	out := new(AerodromeFactorySameAddress)
	if err := aerodromeFactory.abi.UnpackIntoInterface(out, "SameAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeFactoryZeroAddress represents a ZeroAddress error raised by the AerodromeFactory contract.
type AerodromeFactoryZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroAddress()
func AerodromeFactoryZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xd92e233df2717d4a40030e20904abd27b68fcbeede117eaaccbbdac9618c8c73")
}

// UnpackZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroAddress()
func (aerodromeFactory *AerodromeFactory) UnpackZeroAddressError(raw []byte) (*AerodromeFactoryZeroAddress, error) {
	out := new(AerodromeFactoryZeroAddress)
	if err := aerodromeFactory.abi.UnpackIntoInterface(out, "ZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeFactoryZeroFee represents a ZeroFee error raised by the AerodromeFactory contract.
type AerodromeFactoryZeroFee struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroFee()
func AerodromeFactoryZeroFeeErrorID() common.Hash {
	return common.HexToHash("0xaf13986d418e4b80e36f0b9d3502f4029fc089b272f0666af24393def6b9f097")
}

// UnpackZeroFeeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroFee()
func (aerodromeFactory *AerodromeFactory) UnpackZeroFeeError(raw []byte) (*AerodromeFactoryZeroFee, error) {
	out := new(AerodromeFactoryZeroFee)
	if err := aerodromeFactory.abi.UnpackIntoInterface(out, "ZeroFee", raw); err != nil {
		return nil, err
	}
	return out, nil
}
