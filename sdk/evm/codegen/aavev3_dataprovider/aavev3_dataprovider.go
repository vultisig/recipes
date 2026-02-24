// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aavev3_dataprovider

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

// Aavev3DataproviderMetaData contains all meta data concerning the Aavev3Dataprovider contract.
var Aavev3DataproviderMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveConfigurationData\",\"outputs\":[{\"name\":\"decimals\",\"type\":\"uint256\"},{\"name\":\"ltv\",\"type\":\"uint256\"},{\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"name\":\"liquidationBonus\",\"type\":\"uint256\"},{\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"name\":\"isActive\",\"type\":\"bool\"},{\"name\":\"isFrozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "Aavev3Dataprovider",
}

// Aavev3Dataprovider is an auto generated Go binding around an Ethereum contract.
type Aavev3Dataprovider struct {
	abi abi.ABI
}

// NewAavev3Dataprovider creates a new instance of Aavev3Dataprovider.
func NewAavev3Dataprovider() *Aavev3Dataprovider {
	parsed, err := Aavev3DataproviderMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Aavev3Dataprovider{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Aavev3Dataprovider) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetReserveConfigurationData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e150141.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (aavev3Dataprovider *Aavev3Dataprovider) PackGetReserveConfigurationData(asset common.Address) []byte {
	enc, err := aavev3Dataprovider.abi.Pack("getReserveConfigurationData", asset)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetReserveConfigurationData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e150141.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (aavev3Dataprovider *Aavev3Dataprovider) TryPackGetReserveConfigurationData(asset common.Address) ([]byte, error) {
	return aavev3Dataprovider.abi.Pack("getReserveConfigurationData", asset)
}

// GetReserveConfigurationDataOutput serves as a container for the return parameters of contract
// method GetReserveConfigurationData.
type GetReserveConfigurationDataOutput struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}

// UnpackGetReserveConfigurationData is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (aavev3Dataprovider *Aavev3Dataprovider) UnpackGetReserveConfigurationData(data []byte) (GetReserveConfigurationDataOutput, error) {
	out, err := aavev3Dataprovider.abi.Unpack("getReserveConfigurationData", data)
	outstruct := new(GetReserveConfigurationDataOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Decimals = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Ltv = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.LiquidationThreshold = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	outstruct.LiquidationBonus = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.ReserveFactor = abi.ConvertType(out[4], new(big.Int)).(*big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.BorrowingEnabled = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.StableBorrowRateEnabled = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.IsActive = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.IsFrozen = *abi.ConvertType(out[9], new(bool)).(*bool)
	return *outstruct, nil
}
