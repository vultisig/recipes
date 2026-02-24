// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aavev3_pool

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

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Configuration               Struct0
	LiquidityIndex              *big.Int
	CurrentLiquidityRate        *big.Int
	VariableBorrowIndex         *big.Int
	CurrentVariableBorrowRate   *big.Int
	CurrentStableBorrowRate     *big.Int
	LastUpdateTimestamp         *big.Int
	Id                          uint16
	ATokenAddress               common.Address
	StableDebtTokenAddress      common.Address
	VariableDebtTokenAddress    common.Address
	InterestRateStrategyAddress common.Address
	AccruedToTreasury           *big.Int
	Unbacked                    *big.Int
	IsolationModeTotalDebt      *big.Int
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Data *big.Int
}

// Aavev3PoolMetaData contains all meta data concerning the Aavev3Pool contract.
var Aavev3PoolMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"supply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"name\":\"referralCode\",\"type\":\"uint16\"},{\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"components\":[{\"components\":[{\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"configuration\",\"type\":\"tuple\"},{\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"name\":\"currentStableBorrowRate\",\"type\":\"uint128\"},{\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"name\":\"id\",\"type\":\"uint16\"},{\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"name\":\"accruedToTreasury\",\"type\":\"uint128\"},{\"name\":\"unbacked\",\"type\":\"uint128\"},{\"name\":\"isolationModeTotalDebt\",\"type\":\"uint128\"}],\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"name\":\"totalCollateralBase\",\"type\":\"uint256\"},{\"name\":\"totalDebtBase\",\"type\":\"uint256\"},{\"name\":\"availableBorrowsBase\",\"type\":\"uint256\"},{\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"name\":\"ltv\",\"type\":\"uint256\"},{\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "Aavev3Pool",
}

// Aavev3Pool is an auto generated Go binding around an Ethereum contract.
type Aavev3Pool struct {
	abi abi.ABI
}

// NewAavev3Pool creates a new instance of Aavev3Pool.
func NewAavev3Pool() *Aavev3Pool {
	parsed, err := Aavev3PoolMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Aavev3Pool{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Aavev3Pool) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackBorrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa415bcad.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (aavev3Pool *Aavev3Pool) PackBorrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) []byte {
	enc, err := aavev3Pool.abi.Pack("borrow", asset, amount, interestRateMode, referralCode, onBehalfOf)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBorrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa415bcad.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (aavev3Pool *Aavev3Pool) TryPackBorrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) ([]byte, error) {
	return aavev3Pool.abi.Pack("borrow", asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// PackGetReserveData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35ea6a75.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128))
func (aavev3Pool *Aavev3Pool) PackGetReserveData(asset common.Address) []byte {
	enc, err := aavev3Pool.abi.Pack("getReserveData", asset)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetReserveData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35ea6a75.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128))
func (aavev3Pool *Aavev3Pool) TryPackGetReserveData(asset common.Address) ([]byte, error) {
	return aavev3Pool.abi.Pack("getReserveData", asset)
}

// UnpackGetReserveData is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128))
func (aavev3Pool *Aavev3Pool) UnpackGetReserveData(data []byte) (Struct1, error) {
	out, err := aavev3Pool.abi.Unpack("getReserveData", data)
	if err != nil {
		return *new(Struct1), err
	}
	out0 := *abi.ConvertType(out[0], new(Struct1)).(*Struct1)
	return out0, nil
}

// PackGetUserAccountData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf92857c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (aavev3Pool *Aavev3Pool) PackGetUserAccountData(user common.Address) []byte {
	enc, err := aavev3Pool.abi.Pack("getUserAccountData", user)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetUserAccountData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf92857c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (aavev3Pool *Aavev3Pool) TryPackGetUserAccountData(user common.Address) ([]byte, error) {
	return aavev3Pool.abi.Pack("getUserAccountData", user)
}

// GetUserAccountDataOutput serves as a container for the return parameters of contract
// method GetUserAccountData.
type GetUserAccountDataOutput struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}

// UnpackGetUserAccountData is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (aavev3Pool *Aavev3Pool) UnpackGetUserAccountData(data []byte) (GetUserAccountDataOutput, error) {
	out, err := aavev3Pool.abi.Unpack("getUserAccountData", data)
	outstruct := new(GetUserAccountDataOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.TotalCollateralBase = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.TotalDebtBase = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.AvailableBorrowsBase = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	outstruct.CurrentLiquidationThreshold = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.Ltv = abi.ConvertType(out[4], new(big.Int)).(*big.Int)
	outstruct.HealthFactor = abi.ConvertType(out[5], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackRepay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x573ade81.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (aavev3Pool *Aavev3Pool) PackRepay(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) []byte {
	enc, err := aavev3Pool.abi.Pack("repay", asset, amount, interestRateMode, onBehalfOf)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRepay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x573ade81.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (aavev3Pool *Aavev3Pool) TryPackRepay(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) ([]byte, error) {
	return aavev3Pool.abi.Pack("repay", asset, amount, interestRateMode, onBehalfOf)
}

// UnpackRepay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (aavev3Pool *Aavev3Pool) UnpackRepay(data []byte) (*big.Int, error) {
	out, err := aavev3Pool.abi.Unpack("repay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x617ba037.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (aavev3Pool *Aavev3Pool) PackSupply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) []byte {
	enc, err := aavev3Pool.abi.Pack("supply", asset, amount, onBehalfOf, referralCode)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x617ba037.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (aavev3Pool *Aavev3Pool) TryPackSupply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) ([]byte, error) {
	return aavev3Pool.abi.Pack("supply", asset, amount, onBehalfOf, referralCode)
}

// PackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x69328dec.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (aavev3Pool *Aavev3Pool) PackWithdraw(asset common.Address, amount *big.Int, to common.Address) []byte {
	enc, err := aavev3Pool.abi.Pack("withdraw", asset, amount, to)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x69328dec.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (aavev3Pool *Aavev3Pool) TryPackWithdraw(asset common.Address, amount *big.Int, to common.Address) ([]byte, error) {
	return aavev3Pool.abi.Pack("withdraw", asset, amount, to)
}

// UnpackWithdraw is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (aavev3Pool *Aavev3Pool) UnpackWithdraw(data []byte) (*big.Int, error) {
	out, err := aavev3Pool.abi.Unpack("withdraw", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}
