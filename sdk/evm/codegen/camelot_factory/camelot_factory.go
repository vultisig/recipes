// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package camelot_factory

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

// CamelotFactoryMetaData contains all meta data concerning the CamelotFactory contract.
var CamelotFactoryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolDeployer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newDefaultCommunityFee\",\"type\":\"uint8\"}],\"name\":\"DefaultCommunityFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFarmingAddress\",\"type\":\"address\"}],\"name\":\"FarmingAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"name\":\"FeeConfiguration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"Owner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"Pool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVaultAddress\",\"type\":\"address\"}],\"name\":\"VaultAddress\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseFeeConfiguration\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultCommunityFee\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"farmingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolByPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDeployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"alpha1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"alpha2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"beta1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"beta2\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"gamma1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gamma2\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"volumeBeta\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"volumeGamma\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFee\",\"type\":\"uint16\"}],\"name\":\"setBaseFeeConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newDefaultCommunityFee\",\"type\":\"uint8\"}],\"name\":\"setDefaultCommunityFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_farmingAddress\",\"type\":\"address\"}],\"name\":\"setFarmingAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"name\":\"setVaultAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "CamelotFactory",
}

// CamelotFactory is an auto generated Go binding around an Ethereum contract.
type CamelotFactory struct {
	abi abi.ABI
}

// NewCamelotFactory creates a new instance of CamelotFactory.
func NewCamelotFactory() *CamelotFactory {
	parsed, err := CamelotFactoryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &CamelotFactory{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *CamelotFactory) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _poolDeployer, address _vaultAddress) returns()
func (camelotFactory *CamelotFactory) PackConstructor(_poolDeployer common.Address, _vaultAddress common.Address) []byte {
	enc, err := camelotFactory.abi.Pack("", _poolDeployer, _vaultAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBaseFeeConfiguration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9832853a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (camelotFactory *CamelotFactory) PackBaseFeeConfiguration() []byte {
	enc, err := camelotFactory.abi.Pack("baseFeeConfiguration")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBaseFeeConfiguration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9832853a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (camelotFactory *CamelotFactory) TryPackBaseFeeConfiguration() ([]byte, error) {
	return camelotFactory.abi.Pack("baseFeeConfiguration")
}

// BaseFeeConfigurationOutput serves as a container for the return parameters of contract
// method BaseFeeConfiguration.
type BaseFeeConfigurationOutput struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
}

// UnpackBaseFeeConfiguration is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9832853a.
//
// Solidity: function baseFeeConfiguration() view returns(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (camelotFactory *CamelotFactory) UnpackBaseFeeConfiguration(data []byte) (BaseFeeConfigurationOutput, error) {
	out, err := camelotFactory.abi.Unpack("baseFeeConfiguration", data)
	outstruct := new(BaseFeeConfigurationOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Alpha1 = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Alpha2 = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.Beta1 = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Beta2 = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Gamma1 = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Gamma2 = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.VolumeBeta = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.VolumeGamma = *abi.ConvertType(out[7], new(uint16)).(*uint16)
	outstruct.BaseFee = *abi.ConvertType(out[8], new(uint16)).(*uint16)
	return *outstruct, nil
}

// PackCreatePool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe3433615.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (camelotFactory *CamelotFactory) PackCreatePool(tokenA common.Address, tokenB common.Address) []byte {
	enc, err := camelotFactory.abi.Pack("createPool", tokenA, tokenB)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreatePool is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe3433615.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (camelotFactory *CamelotFactory) TryPackCreatePool(tokenA common.Address, tokenB common.Address) ([]byte, error) {
	return camelotFactory.abi.Pack("createPool", tokenA, tokenB)
}

// UnpackCreatePool is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe3433615.
//
// Solidity: function createPool(address tokenA, address tokenB) returns(address pool)
func (camelotFactory *CamelotFactory) UnpackCreatePool(data []byte) (common.Address, error) {
	out, err := camelotFactory.abi.Unpack("createPool", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDefaultCommunityFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f8a39dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function defaultCommunityFee() view returns(uint8)
func (camelotFactory *CamelotFactory) PackDefaultCommunityFee() []byte {
	enc, err := camelotFactory.abi.Pack("defaultCommunityFee")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDefaultCommunityFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f8a39dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function defaultCommunityFee() view returns(uint8)
func (camelotFactory *CamelotFactory) TryPackDefaultCommunityFee() ([]byte, error) {
	return camelotFactory.abi.Pack("defaultCommunityFee")
}

// UnpackDefaultCommunityFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f8a39dd.
//
// Solidity: function defaultCommunityFee() view returns(uint8)
func (camelotFactory *CamelotFactory) UnpackDefaultCommunityFee(data []byte) (uint8, error) {
	out, err := camelotFactory.abi.Unpack("defaultCommunityFee", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackFarmingAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a2ade58.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function farmingAddress() view returns(address)
func (camelotFactory *CamelotFactory) PackFarmingAddress() []byte {
	enc, err := camelotFactory.abi.Pack("farmingAddress")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFarmingAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8a2ade58.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function farmingAddress() view returns(address)
func (camelotFactory *CamelotFactory) TryPackFarmingAddress() ([]byte, error) {
	return camelotFactory.abi.Pack("farmingAddress")
}

// UnpackFarmingAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8a2ade58.
//
// Solidity: function farmingAddress() view returns(address)
func (camelotFactory *CamelotFactory) UnpackFarmingAddress(data []byte) (common.Address, error) {
	out, err := camelotFactory.abi.Unpack("farmingAddress", data)
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
func (camelotFactory *CamelotFactory) PackOwner() []byte {
	enc, err := camelotFactory.abi.Pack("owner")
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
func (camelotFactory *CamelotFactory) TryPackOwner() ([]byte, error) {
	return camelotFactory.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (camelotFactory *CamelotFactory) UnpackOwner(data []byte) (common.Address, error) {
	out, err := camelotFactory.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPoolByPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd9a641e1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (camelotFactory *CamelotFactory) PackPoolByPair(arg0 common.Address, arg1 common.Address) []byte {
	enc, err := camelotFactory.abi.Pack("poolByPair", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPoolByPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd9a641e1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (camelotFactory *CamelotFactory) TryPackPoolByPair(arg0 common.Address, arg1 common.Address) ([]byte, error) {
	return camelotFactory.abi.Pack("poolByPair", arg0, arg1)
}

// UnpackPoolByPair is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd9a641e1.
//
// Solidity: function poolByPair(address , address ) view returns(address)
func (camelotFactory *CamelotFactory) UnpackPoolByPair(data []byte) (common.Address, error) {
	out, err := camelotFactory.abi.Unpack("poolByPair", data)
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
func (camelotFactory *CamelotFactory) PackPoolDeployer() []byte {
	enc, err := camelotFactory.abi.Pack("poolDeployer")
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
func (camelotFactory *CamelotFactory) TryPackPoolDeployer() ([]byte, error) {
	return camelotFactory.abi.Pack("poolDeployer")
}

// UnpackPoolDeployer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3119049a.
//
// Solidity: function poolDeployer() view returns(address)
func (camelotFactory *CamelotFactory) UnpackPoolDeployer(data []byte) (common.Address, error) {
	out, err := camelotFactory.abi.Unpack("poolDeployer", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetBaseFeeConfiguration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d6d7e93.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (camelotFactory *CamelotFactory) PackSetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) []byte {
	enc, err := camelotFactory.abi.Pack("setBaseFeeConfiguration", alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetBaseFeeConfiguration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d6d7e93.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setBaseFeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee) returns()
func (camelotFactory *CamelotFactory) TryPackSetBaseFeeConfiguration(alpha1 uint16, alpha2 uint16, beta1 uint32, beta2 uint32, gamma1 uint16, gamma2 uint16, volumeBeta uint32, volumeGamma uint16, baseFee uint16) ([]byte, error) {
	return camelotFactory.abi.Pack("setBaseFeeConfiguration", alpha1, alpha2, beta1, beta2, gamma1, gamma2, volumeBeta, volumeGamma, baseFee)
}

// PackSetDefaultCommunityFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x371e3521.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setDefaultCommunityFee(uint8 newDefaultCommunityFee) returns()
func (camelotFactory *CamelotFactory) PackSetDefaultCommunityFee(newDefaultCommunityFee uint8) []byte {
	enc, err := camelotFactory.abi.Pack("setDefaultCommunityFee", newDefaultCommunityFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetDefaultCommunityFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x371e3521.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setDefaultCommunityFee(uint8 newDefaultCommunityFee) returns()
func (camelotFactory *CamelotFactory) TryPackSetDefaultCommunityFee(newDefaultCommunityFee uint8) ([]byte, error) {
	return camelotFactory.abi.Pack("setDefaultCommunityFee", newDefaultCommunityFee)
}

// PackSetFarmingAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb001f618.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (camelotFactory *CamelotFactory) PackSetFarmingAddress(farmingAddress common.Address) []byte {
	enc, err := camelotFactory.abi.Pack("setFarmingAddress", farmingAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFarmingAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb001f618.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFarmingAddress(address _farmingAddress) returns()
func (camelotFactory *CamelotFactory) TryPackSetFarmingAddress(farmingAddress common.Address) ([]byte, error) {
	return camelotFactory.abi.Pack("setFarmingAddress", farmingAddress)
}

// PackSetOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13af4035.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setOwner(address _owner) returns()
func (camelotFactory *CamelotFactory) PackSetOwner(owner common.Address) []byte {
	enc, err := camelotFactory.abi.Pack("setOwner", owner)
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
func (camelotFactory *CamelotFactory) TryPackSetOwner(owner common.Address) ([]byte, error) {
	return camelotFactory.abi.Pack("setOwner", owner)
}

// PackSetVaultAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85535cc5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (camelotFactory *CamelotFactory) PackSetVaultAddress(vaultAddress common.Address) []byte {
	enc, err := camelotFactory.abi.Pack("setVaultAddress", vaultAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetVaultAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85535cc5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setVaultAddress(address _vaultAddress) returns()
func (camelotFactory *CamelotFactory) TryPackSetVaultAddress(vaultAddress common.Address) ([]byte, error) {
	return camelotFactory.abi.Pack("setVaultAddress", vaultAddress)
}

// PackVaultAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x430bf08a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function vaultAddress() view returns(address)
func (camelotFactory *CamelotFactory) PackVaultAddress() []byte {
	enc, err := camelotFactory.abi.Pack("vaultAddress")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVaultAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x430bf08a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function vaultAddress() view returns(address)
func (camelotFactory *CamelotFactory) TryPackVaultAddress() ([]byte, error) {
	return camelotFactory.abi.Pack("vaultAddress")
}

// UnpackVaultAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (camelotFactory *CamelotFactory) UnpackVaultAddress(data []byte) (common.Address, error) {
	out, err := camelotFactory.abi.Unpack("vaultAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// CamelotFactoryDefaultCommunityFee represents a DefaultCommunityFee event raised by the CamelotFactory contract.
type CamelotFactoryDefaultCommunityFee struct {
	NewDefaultCommunityFee uint8
	Raw                    *types.Log // Blockchain specific contextual infos
}

const CamelotFactoryDefaultCommunityFeeEventName = "DefaultCommunityFee"

// ContractEventName returns the user-defined event name.
func (CamelotFactoryDefaultCommunityFee) ContractEventName() string {
	return CamelotFactoryDefaultCommunityFeeEventName
}

// UnpackDefaultCommunityFeeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultCommunityFee(uint8 newDefaultCommunityFee)
func (camelotFactory *CamelotFactory) UnpackDefaultCommunityFeeEvent(log *types.Log) (*CamelotFactoryDefaultCommunityFee, error) {
	event := "DefaultCommunityFee"
	if log.Topics[0] != camelotFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CamelotFactoryDefaultCommunityFee)
	if len(log.Data) > 0 {
		if err := camelotFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range camelotFactory.abi.Events[event].Inputs {
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

// CamelotFactoryFarmingAddress represents a FarmingAddress event raised by the CamelotFactory contract.
type CamelotFactoryFarmingAddress struct {
	NewFarmingAddress common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const CamelotFactoryFarmingAddressEventName = "FarmingAddress"

// ContractEventName returns the user-defined event name.
func (CamelotFactoryFarmingAddress) ContractEventName() string {
	return CamelotFactoryFarmingAddressEventName
}

// UnpackFarmingAddressEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FarmingAddress(address indexed newFarmingAddress)
func (camelotFactory *CamelotFactory) UnpackFarmingAddressEvent(log *types.Log) (*CamelotFactoryFarmingAddress, error) {
	event := "FarmingAddress"
	if log.Topics[0] != camelotFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CamelotFactoryFarmingAddress)
	if len(log.Data) > 0 {
		if err := camelotFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range camelotFactory.abi.Events[event].Inputs {
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

// CamelotFactoryFeeConfiguration represents a FeeConfiguration event raised by the CamelotFactory contract.
type CamelotFactoryFeeConfiguration struct {
	Alpha1      uint16
	Alpha2      uint16
	Beta1       uint32
	Beta2       uint32
	Gamma1      uint16
	Gamma2      uint16
	VolumeBeta  uint32
	VolumeGamma uint16
	BaseFee     uint16
	Raw         *types.Log // Blockchain specific contextual infos
}

const CamelotFactoryFeeConfigurationEventName = "FeeConfiguration"

// ContractEventName returns the user-defined event name.
func (CamelotFactoryFeeConfiguration) ContractEventName() string {
	return CamelotFactoryFeeConfigurationEventName
}

// UnpackFeeConfigurationEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeConfiguration(uint16 alpha1, uint16 alpha2, uint32 beta1, uint32 beta2, uint16 gamma1, uint16 gamma2, uint32 volumeBeta, uint16 volumeGamma, uint16 baseFee)
func (camelotFactory *CamelotFactory) UnpackFeeConfigurationEvent(log *types.Log) (*CamelotFactoryFeeConfiguration, error) {
	event := "FeeConfiguration"
	if log.Topics[0] != camelotFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CamelotFactoryFeeConfiguration)
	if len(log.Data) > 0 {
		if err := camelotFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range camelotFactory.abi.Events[event].Inputs {
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

// CamelotFactoryOwner represents a Owner event raised by the CamelotFactory contract.
type CamelotFactoryOwner struct {
	NewOwner common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const CamelotFactoryOwnerEventName = "Owner"

// ContractEventName returns the user-defined event name.
func (CamelotFactoryOwner) ContractEventName() string {
	return CamelotFactoryOwnerEventName
}

// UnpackOwnerEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Owner(address indexed newOwner)
func (camelotFactory *CamelotFactory) UnpackOwnerEvent(log *types.Log) (*CamelotFactoryOwner, error) {
	event := "Owner"
	if log.Topics[0] != camelotFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CamelotFactoryOwner)
	if len(log.Data) > 0 {
		if err := camelotFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range camelotFactory.abi.Events[event].Inputs {
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

// CamelotFactoryPool represents a Pool event raised by the CamelotFactory contract.
type CamelotFactoryPool struct {
	Token0 common.Address
	Token1 common.Address
	Pool   common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const CamelotFactoryPoolEventName = "Pool"

// ContractEventName returns the user-defined event name.
func (CamelotFactoryPool) ContractEventName() string {
	return CamelotFactoryPoolEventName
}

// UnpackPoolEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Pool(address indexed token0, address indexed token1, address pool)
func (camelotFactory *CamelotFactory) UnpackPoolEvent(log *types.Log) (*CamelotFactoryPool, error) {
	event := "Pool"
	if log.Topics[0] != camelotFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CamelotFactoryPool)
	if len(log.Data) > 0 {
		if err := camelotFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range camelotFactory.abi.Events[event].Inputs {
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

// CamelotFactoryVaultAddress represents a VaultAddress event raised by the CamelotFactory contract.
type CamelotFactoryVaultAddress struct {
	NewVaultAddress common.Address
	Raw             *types.Log // Blockchain specific contextual infos
}

const CamelotFactoryVaultAddressEventName = "VaultAddress"

// ContractEventName returns the user-defined event name.
func (CamelotFactoryVaultAddress) ContractEventName() string {
	return CamelotFactoryVaultAddressEventName
}

// UnpackVaultAddressEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VaultAddress(address indexed newVaultAddress)
func (camelotFactory *CamelotFactory) UnpackVaultAddressEvent(log *types.Log) (*CamelotFactoryVaultAddress, error) {
	event := "VaultAddress"
	if log.Topics[0] != camelotFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(CamelotFactoryVaultAddress)
	if len(log.Data) > 0 {
		if err := camelotFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range camelotFactory.abi.Events[event].Inputs {
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
