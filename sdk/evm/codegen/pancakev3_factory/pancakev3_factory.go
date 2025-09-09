// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancakev3_factory

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

// Pancakev3FactoryMetaData contains all meta data concerning the Pancakev3Factory contract.
var Pancakev3FactoryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolDeployer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"FeeAmountEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"FeeAmountExtraInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lmPoolDeployer\",\"type\":\"address\"}],\"name\":\"SetLmPoolDeployer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"WhiteListAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount0Requested\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1Requested\",\"type\":\"uint128\"}],\"name\":\"collectProtocol\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount0\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount1\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"enableFeeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"feeAmountTickSpacingExtraInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lmPoolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"whitelistRequested\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setFeeAmountExtraInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeProtocol0\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeProtocol1\",\"type\":\"uint32\"}],\"name\":\"setFeeProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lmPool\",\"type\":\"address\"}],\"name\":\"setLmPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lmPoolDeployer\",\"type\":\"address\"}],\"name\":\"setLmPoolDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"setWhiteListAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Pancakev3Factory",
}

// Pancakev3Factory is an auto generated Go binding around an Ethereum contract.
type Pancakev3Factory struct {
	abi abi.ABI
}

// NewPancakev3Factory creates a new instance of Pancakev3Factory.
func NewPancakev3Factory() *Pancakev3Factory {
	parsed, err := Pancakev3FactoryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Pancakev3Factory{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Pancakev3Factory) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _poolDeployer) returns()
func (pancakev3Factory *Pancakev3Factory) PackConstructor(_poolDeployer common.Address) []byte {
	enc, err := pancakev3Factory.abi.Pack("", _poolDeployer)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCollectProtocol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43db87da.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (pancakev3Factory *Pancakev3Factory) PackCollectProtocol(pool common.Address, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) []byte {
	enc, err := pancakev3Factory.abi.Pack("collectProtocol", pool, recipient, amount0Requested, amount1Requested)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCollectProtocol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43db87da.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (pancakev3Factory *Pancakev3Factory) TryPackCollectProtocol(pool common.Address, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) ([]byte, error) {
	return pancakev3Factory.abi.Pack("collectProtocol", pool, recipient, amount0Requested, amount1Requested)
}

// CollectProtocolOutput serves as a container for the return parameters of contract
// method CollectProtocol.
type CollectProtocolOutput struct {
	Amount0 *big.Int
	Amount1 *big.Int
}

// UnpackCollectProtocol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x43db87da.
//
// Solidity: function collectProtocol(address pool, address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (pancakev3Factory *Pancakev3Factory) UnpackCollectProtocol(data []byte) (CollectProtocolOutput, error) {
	out, err := pancakev3Factory.abi.Unpack("collectProtocol", data)
	outstruct := new(CollectProtocolOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Amount0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Amount1 = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackCreatePool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1671295.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (pancakev3Factory *Pancakev3Factory) PackCreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) []byte {
	enc, err := pancakev3Factory.abi.Pack("createPool", tokenA, tokenB, fee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreatePool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1671295.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (pancakev3Factory *Pancakev3Factory) TryPackCreatePool(tokenA common.Address, tokenB common.Address, fee *big.Int) ([]byte, error) {
	return pancakev3Factory.abi.Pack("createPool", tokenA, tokenB, fee)
}

// UnpackCreatePool is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1671295.
//
// Solidity: function createPool(address tokenA, address tokenB, uint24 fee) returns(address pool)
func (pancakev3Factory *Pancakev3Factory) UnpackCreatePool(data []byte) (common.Address, error) {
	out, err := pancakev3Factory.abi.Unpack("createPool", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackEnableFeeAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a7c195f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (pancakev3Factory *Pancakev3Factory) PackEnableFeeAmount(fee *big.Int, tickSpacing *big.Int) []byte {
	enc, err := pancakev3Factory.abi.Pack("enableFeeAmount", fee, tickSpacing)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEnableFeeAmount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a7c195f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function enableFeeAmount(uint24 fee, int24 tickSpacing) returns()
func (pancakev3Factory *Pancakev3Factory) TryPackEnableFeeAmount(fee *big.Int, tickSpacing *big.Int) ([]byte, error) {
	return pancakev3Factory.abi.Pack("enableFeeAmount", fee, tickSpacing)
}

// PackFeeAmountTickSpacing is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22afcccb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (pancakev3Factory *Pancakev3Factory) PackFeeAmountTickSpacing(arg0 *big.Int) []byte {
	enc, err := pancakev3Factory.abi.Pack("feeAmountTickSpacing", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFeeAmountTickSpacing is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22afcccb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (pancakev3Factory *Pancakev3Factory) TryPackFeeAmountTickSpacing(arg0 *big.Int) ([]byte, error) {
	return pancakev3Factory.abi.Pack("feeAmountTickSpacing", arg0)
}

// UnpackFeeAmountTickSpacing is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x22afcccb.
//
// Solidity: function feeAmountTickSpacing(uint24 ) view returns(int24)
func (pancakev3Factory *Pancakev3Factory) UnpackFeeAmountTickSpacing(data []byte) (*big.Int, error) {
	out, err := pancakev3Factory.abi.Unpack("feeAmountTickSpacing", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackFeeAmountTickSpacingExtraInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x88e8006d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (pancakev3Factory *Pancakev3Factory) PackFeeAmountTickSpacingExtraInfo(arg0 *big.Int) []byte {
	enc, err := pancakev3Factory.abi.Pack("feeAmountTickSpacingExtraInfo", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFeeAmountTickSpacingExtraInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x88e8006d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (pancakev3Factory *Pancakev3Factory) TryPackFeeAmountTickSpacingExtraInfo(arg0 *big.Int) ([]byte, error) {
	return pancakev3Factory.abi.Pack("feeAmountTickSpacingExtraInfo", arg0)
}

// FeeAmountTickSpacingExtraInfoOutput serves as a container for the return parameters of contract
// method FeeAmountTickSpacingExtraInfo.
type FeeAmountTickSpacingExtraInfoOutput struct {
	WhitelistRequested bool
	Enabled            bool
}

// UnpackFeeAmountTickSpacingExtraInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x88e8006d.
//
// Solidity: function feeAmountTickSpacingExtraInfo(uint24 ) view returns(bool whitelistRequested, bool enabled)
func (pancakev3Factory *Pancakev3Factory) UnpackFeeAmountTickSpacingExtraInfo(data []byte) (FeeAmountTickSpacingExtraInfoOutput, error) {
	out, err := pancakev3Factory.abi.Unpack("feeAmountTickSpacingExtraInfo", data)
	outstruct := new(FeeAmountTickSpacingExtraInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.WhitelistRequested = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	return *outstruct, nil
}

// PackGetPool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1698ee82.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (pancakev3Factory *Pancakev3Factory) PackGetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) []byte {
	enc, err := pancakev3Factory.abi.Pack("getPool", arg0, arg1, arg2)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1698ee82.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (pancakev3Factory *Pancakev3Factory) TryPackGetPool(arg0 common.Address, arg1 common.Address, arg2 *big.Int) ([]byte, error) {
	return pancakev3Factory.abi.Pack("getPool", arg0, arg1, arg2)
}

// UnpackGetPool is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1698ee82.
//
// Solidity: function getPool(address , address , uint24 ) view returns(address)
func (pancakev3Factory *Pancakev3Factory) UnpackGetPool(data []byte) (common.Address, error) {
	out, err := pancakev3Factory.abi.Unpack("getPool", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackLmPoolDeployer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e492ac8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (pancakev3Factory *Pancakev3Factory) PackLmPoolDeployer() []byte {
	enc, err := pancakev3Factory.abi.Pack("lmPoolDeployer")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLmPoolDeployer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e492ac8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (pancakev3Factory *Pancakev3Factory) TryPackLmPoolDeployer() ([]byte, error) {
	return pancakev3Factory.abi.Pack("lmPoolDeployer")
}

// UnpackLmPoolDeployer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5e492ac8.
//
// Solidity: function lmPoolDeployer() view returns(address)
func (pancakev3Factory *Pancakev3Factory) UnpackLmPoolDeployer(data []byte) (common.Address, error) {
	out, err := pancakev3Factory.abi.Unpack("lmPoolDeployer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (pancakev3Factory *Pancakev3Factory) PackOwner() []byte {
	enc, err := pancakev3Factory.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (pancakev3Factory *Pancakev3Factory) TryPackOwner() ([]byte, error) {
	return pancakev3Factory.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (pancakev3Factory *Pancakev3Factory) UnpackOwner(data []byte) (common.Address, error) {
	out, err := pancakev3Factory.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPoolDeployer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3119049a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function poolDeployer() view returns(address)
func (pancakev3Factory *Pancakev3Factory) PackPoolDeployer() []byte {
	enc, err := pancakev3Factory.abi.Pack("poolDeployer")
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
func (pancakev3Factory *Pancakev3Factory) TryPackPoolDeployer() ([]byte, error) {
	return pancakev3Factory.abi.Pack("poolDeployer")
}

// UnpackPoolDeployer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (pancakev3Factory *Pancakev3Factory) UnpackPoolDeployer(data []byte) (common.Address, error) {
	out, err := pancakev3Factory.abi.Unpack("poolDeployer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetFeeAmountExtraInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ff38e80.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFeeAmountExtraInfo(uint24 fee, bool whitelistRequested, bool enabled) returns()
func (pancakev3Factory *Pancakev3Factory) PackSetFeeAmountExtraInfo(fee *big.Int, whitelistRequested bool, enabled bool) []byte {
	enc, err := pancakev3Factory.abi.Pack("setFeeAmountExtraInfo", fee, whitelistRequested, enabled)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFeeAmountExtraInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ff38e80.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFeeAmountExtraInfo(uint24 fee, bool whitelistRequested, bool enabled) returns()
func (pancakev3Factory *Pancakev3Factory) TryPackSetFeeAmountExtraInfo(fee *big.Int, whitelistRequested bool, enabled bool) ([]byte, error) {
	return pancakev3Factory.abi.Pack("setFeeAmountExtraInfo", fee, whitelistRequested, enabled)
}

// PackSetFeeProtocol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7e8435e6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFeeProtocol(address pool, uint32 feeProtocol0, uint32 feeProtocol1) returns()
func (pancakev3Factory *Pancakev3Factory) PackSetFeeProtocol(pool common.Address, feeProtocol0 uint32, feeProtocol1 uint32) []byte {
	enc, err := pancakev3Factory.abi.Pack("setFeeProtocol", pool, feeProtocol0, feeProtocol1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFeeProtocol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7e8435e6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFeeProtocol(address pool, uint32 feeProtocol0, uint32 feeProtocol1) returns()
func (pancakev3Factory *Pancakev3Factory) TryPackSetFeeProtocol(pool common.Address, feeProtocol0 uint32, feeProtocol1 uint32) ([]byte, error) {
	return pancakev3Factory.abi.Pack("setFeeProtocol", pool, feeProtocol0, feeProtocol1)
}

// PackSetLmPool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11ff5e8d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setLmPool(address pool, address lmPool) returns()
func (pancakev3Factory *Pancakev3Factory) PackSetLmPool(pool common.Address, lmPool common.Address) []byte {
	enc, err := pancakev3Factory.abi.Pack("setLmPool", pool, lmPool)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetLmPool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11ff5e8d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setLmPool(address pool, address lmPool) returns()
func (pancakev3Factory *Pancakev3Factory) TryPackSetLmPool(pool common.Address, lmPool common.Address) ([]byte, error) {
	return pancakev3Factory.abi.Pack("setLmPool", pool, lmPool)
}

// PackSetLmPoolDeployer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x80d6a792.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setLmPoolDeployer(address _lmPoolDeployer) returns()
func (pancakev3Factory *Pancakev3Factory) PackSetLmPoolDeployer(lmPoolDeployer common.Address) []byte {
	enc, err := pancakev3Factory.abi.Pack("setLmPoolDeployer", lmPoolDeployer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetLmPoolDeployer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x80d6a792.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setLmPoolDeployer(address _lmPoolDeployer) returns()
func (pancakev3Factory *Pancakev3Factory) TryPackSetLmPoolDeployer(lmPoolDeployer common.Address) ([]byte, error) {
	return pancakev3Factory.abi.Pack("setLmPoolDeployer", lmPoolDeployer)
}

// PackSetOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13af4035.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setOwner(address _owner) returns()
func (pancakev3Factory *Pancakev3Factory) PackSetOwner(owner common.Address) []byte {
	enc, err := pancakev3Factory.abi.Pack("setOwner", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13af4035.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setOwner(address _owner) returns()
func (pancakev3Factory *Pancakev3Factory) TryPackSetOwner(owner common.Address) ([]byte, error) {
	return pancakev3Factory.abi.Pack("setOwner", owner)
}

// PackSetWhiteListAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe4a86a99.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setWhiteListAddress(address user, bool verified) returns()
func (pancakev3Factory *Pancakev3Factory) PackSetWhiteListAddress(user common.Address, verified bool) []byte {
	enc, err := pancakev3Factory.abi.Pack("setWhiteListAddress", user, verified)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetWhiteListAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe4a86a99.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setWhiteListAddress(address user, bool verified) returns()
func (pancakev3Factory *Pancakev3Factory) TryPackSetWhiteListAddress(user common.Address, verified bool) ([]byte, error) {
	return pancakev3Factory.abi.Pack("setWhiteListAddress", user, verified)
}

// Pancakev3FactoryFeeAmountEnabled represents a FeeAmountEnabled event raised by the Pancakev3Factory contract.
type Pancakev3FactoryFeeAmountEnabled struct {
	Fee         *big.Int
	TickSpacing *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const Pancakev3FactoryFeeAmountEnabledEventName = "FeeAmountEnabled"

// ContractEventName returns the user-defined event name.
func (Pancakev3FactoryFeeAmountEnabled) ContractEventName() string {
	return Pancakev3FactoryFeeAmountEnabledEventName
}

// UnpackFeeAmountEnabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeAmountEnabled(uint24 indexed fee, int24 indexed tickSpacing)
func (pancakev3Factory *Pancakev3Factory) UnpackFeeAmountEnabledEvent(log *types.Log) (*Pancakev3FactoryFeeAmountEnabled, error) {
	event := "FeeAmountEnabled"
	if log.Topics[0] != pancakev3Factory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Pancakev3FactoryFeeAmountEnabled)
	if len(log.Data) > 0 {
		if err := pancakev3Factory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pancakev3Factory.abi.Events[event].Inputs {
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

// Pancakev3FactoryFeeAmountExtraInfoUpdated represents a FeeAmountExtraInfoUpdated event raised by the Pancakev3Factory contract.
type Pancakev3FactoryFeeAmountExtraInfoUpdated struct {
	Fee                *big.Int
	WhitelistRequested bool
	Enabled            bool
	Raw                *types.Log // Blockchain specific contextual infos
}

const Pancakev3FactoryFeeAmountExtraInfoUpdatedEventName = "FeeAmountExtraInfoUpdated"

// ContractEventName returns the user-defined event name.
func (Pancakev3FactoryFeeAmountExtraInfoUpdated) ContractEventName() string {
	return Pancakev3FactoryFeeAmountExtraInfoUpdatedEventName
}

// UnpackFeeAmountExtraInfoUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeAmountExtraInfoUpdated(uint24 indexed fee, bool whitelistRequested, bool enabled)
func (pancakev3Factory *Pancakev3Factory) UnpackFeeAmountExtraInfoUpdatedEvent(log *types.Log) (*Pancakev3FactoryFeeAmountExtraInfoUpdated, error) {
	event := "FeeAmountExtraInfoUpdated"
	if log.Topics[0] != pancakev3Factory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Pancakev3FactoryFeeAmountExtraInfoUpdated)
	if len(log.Data) > 0 {
		if err := pancakev3Factory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pancakev3Factory.abi.Events[event].Inputs {
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

// Pancakev3FactoryOwnerChanged represents a OwnerChanged event raised by the Pancakev3Factory contract.
type Pancakev3FactoryOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const Pancakev3FactoryOwnerChangedEventName = "OwnerChanged"

// ContractEventName returns the user-defined event name.
func (Pancakev3FactoryOwnerChanged) ContractEventName() string {
	return Pancakev3FactoryOwnerChangedEventName
}

// UnpackOwnerChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (pancakev3Factory *Pancakev3Factory) UnpackOwnerChangedEvent(log *types.Log) (*Pancakev3FactoryOwnerChanged, error) {
	event := "OwnerChanged"
	if log.Topics[0] != pancakev3Factory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Pancakev3FactoryOwnerChanged)
	if len(log.Data) > 0 {
		if err := pancakev3Factory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pancakev3Factory.abi.Events[event].Inputs {
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

// Pancakev3FactoryPoolCreated represents a PoolCreated event raised by the Pancakev3Factory contract.
type Pancakev3FactoryPoolCreated struct {
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Pool        common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const Pancakev3FactoryPoolCreatedEventName = "PoolCreated"

// ContractEventName returns the user-defined event name.
func (Pancakev3FactoryPoolCreated) ContractEventName() string {
	return Pancakev3FactoryPoolCreatedEventName
}

// UnpackPoolCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, uint24 indexed fee, int24 tickSpacing, address pool)
func (pancakev3Factory *Pancakev3Factory) UnpackPoolCreatedEvent(log *types.Log) (*Pancakev3FactoryPoolCreated, error) {
	event := "PoolCreated"
	if log.Topics[0] != pancakev3Factory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Pancakev3FactoryPoolCreated)
	if len(log.Data) > 0 {
		if err := pancakev3Factory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pancakev3Factory.abi.Events[event].Inputs {
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

// Pancakev3FactorySetLmPoolDeployer represents a SetLmPoolDeployer event raised by the Pancakev3Factory contract.
type Pancakev3FactorySetLmPoolDeployer struct {
	LmPoolDeployer common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const Pancakev3FactorySetLmPoolDeployerEventName = "SetLmPoolDeployer"

// ContractEventName returns the user-defined event name.
func (Pancakev3FactorySetLmPoolDeployer) ContractEventName() string {
	return Pancakev3FactorySetLmPoolDeployerEventName
}

// UnpackSetLmPoolDeployerEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SetLmPoolDeployer(address indexed lmPoolDeployer)
func (pancakev3Factory *Pancakev3Factory) UnpackSetLmPoolDeployerEvent(log *types.Log) (*Pancakev3FactorySetLmPoolDeployer, error) {
	event := "SetLmPoolDeployer"
	if log.Topics[0] != pancakev3Factory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Pancakev3FactorySetLmPoolDeployer)
	if len(log.Data) > 0 {
		if err := pancakev3Factory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pancakev3Factory.abi.Events[event].Inputs {
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

// Pancakev3FactoryWhiteListAdded represents a WhiteListAdded event raised by the Pancakev3Factory contract.
type Pancakev3FactoryWhiteListAdded struct {
	User     common.Address
	Verified bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const Pancakev3FactoryWhiteListAddedEventName = "WhiteListAdded"

// ContractEventName returns the user-defined event name.
func (Pancakev3FactoryWhiteListAdded) ContractEventName() string {
	return Pancakev3FactoryWhiteListAddedEventName
}

// UnpackWhiteListAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WhiteListAdded(address indexed user, bool verified)
func (pancakev3Factory *Pancakev3Factory) UnpackWhiteListAddedEvent(log *types.Log) (*Pancakev3FactoryWhiteListAdded, error) {
	event := "WhiteListAdded"
	if log.Topics[0] != pancakev3Factory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Pancakev3FactoryWhiteListAdded)
	if len(log.Data) > 0 {
		if err := pancakev3Factory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range pancakev3Factory.abi.Events[event].Inputs {
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
