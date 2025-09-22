// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package joe_factory

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

// ILBFactoryLBPairInformation is an auto generated low-level Go binding around an user-defined struct.
type ILBFactoryLBPairInformation struct {
	BinStep           uint16
	LBPair            common.Address
	CreatedByOwner    bool
	IgnoredForRouting bool
}

// JoeFactoryMetaData contains all meta data concerning the JoeFactory contract.
var JoeFactoryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"flashLoanFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LBFactory__AddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__BinStepHasNoPreset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__BinStepTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFees\",\"type\":\"uint256\"}],\"name\":\"LBFactory__FlashLoanFeeAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LBFactory__IdenticalAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBFactory__ImplementationNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__LBPairAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__LBPairDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBFactory__LBPairIgnoredIsAlreadyInTheSameState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__LBPairNotCreated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"LBPairImplementation\",\"type\":\"address\"}],\"name\":\"LBFactory__LBPairSafetyCheckFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__PresetIsLockedForUsers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBFactory__PresetOpenStateIsAlreadyInTheSameState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"LBFactory__QuoteAssetAlreadyWhitelisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"LBFactory__QuoteAssetNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"LBFactory__SameFeeRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"flashLoanFee\",\"type\":\"uint256\"}],\"name\":\"LBFactory__SameFlashLoanFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"LBPairImplementation\",\"type\":\"address\"}],\"name\":\"LBFactory__SameImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PairParametersHelper__InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__NoPendingOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__NotPendingOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__PendingOwnerAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeCast__Exceeds16Bits\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"y\",\"type\":\"int256\"}],\"name\":\"Uint128x128Math__PowUnderflow\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFlashLoanFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFlashLoanFee\",\"type\":\"uint256\"}],\"name\":\"FlashLoanFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"LBPairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ignored\",\"type\":\"bool\"}],\"name\":\"LBPairIgnoredStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLBPairImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"LBPairImplementation\",\"type\":\"address\"}],\"name\":\"LBPairImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"PendingOwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"PresetOpenStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"PresetRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filterPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decayPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reductionFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableFeeControl\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint256\"}],\"name\":\"PresetSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"QuoteAssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"QuoteAssetRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"addQuoteAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"becomeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"activeId\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"}],\"name\":\"createLBPair\",\"outputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"forceDecay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllBinSteps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"binStepWithPreset\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"}],\"name\":\"getAllLBPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"createdByOwner\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ignoredForRouting\",\"type\":\"bool\"}],\"internalType\":\"structILBFactory.LBPairInformation[]\",\"name\":\"lbPairsAvailable\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFlashLoanFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"flashLoanFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getLBPairAtIndex\",\"outputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"lbPair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLBPairImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lbPairImplementation\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"getLBPairInformation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"createdByOwner\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ignoredForRouting\",\"type\":\"bool\"}],\"internalType\":\"structILBFactory.LBPairInformation\",\"name\":\"lbPairInformation\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxFlashLoanFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinBinStep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minBinStep\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfLBPairs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lbPairNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfQuoteAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numberOfQuoteAssets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpenBinSteps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"openBinStep\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"getPreset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filterPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decayPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reductionFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableFeeControl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getQuoteAssetAtIndex\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isQuoteAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isQuote\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"}],\"name\":\"removePreset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"removeQuoteAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"filterPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"decayPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reductionFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"variableFeeControl\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"protocolShare\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint24\"}],\"name\":\"setFeesParametersOnPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"flashLoanFee\",\"type\":\"uint256\"}],\"name\":\"setFlashLoanFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"ignored\",\"type\":\"bool\"}],\"name\":\"setLBPairIgnored\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLBPairImplementation\",\"type\":\"address\"}],\"name\":\"setLBPairImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner_\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"filterPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"decayPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reductionFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"variableFeeControl\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"protocolShare\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"setPreset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"setPresetOpenState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "JoeFactory",
}

// JoeFactory is an auto generated Go binding around an Ethereum contract.
type JoeFactory struct {
	abi abi.ABI
}

// NewJoeFactory creates a new instance of JoeFactory.
func NewJoeFactory() *JoeFactory {
	parsed, err := JoeFactoryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &JoeFactory{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *JoeFactory) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address feeRecipient, uint256 flashLoanFee) returns()
func (joeFactory *JoeFactory) PackConstructor(feeRecipient common.Address, flashLoanFee *big.Int) []byte {
	enc, err := joeFactory.abi.Pack("", feeRecipient, flashLoanFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAddQuoteAsset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a440923.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addQuoteAsset(address quoteAsset) returns()
func (joeFactory *JoeFactory) PackAddQuoteAsset(quoteAsset common.Address) []byte {
	enc, err := joeFactory.abi.Pack("addQuoteAsset", quoteAsset)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddQuoteAsset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a440923.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addQuoteAsset(address quoteAsset) returns()
func (joeFactory *JoeFactory) TryPackAddQuoteAsset(quoteAsset common.Address) ([]byte, error) {
	return joeFactory.abi.Pack("addQuoteAsset", quoteAsset)
}

// PackBecomeOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf9dca989.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function becomeOwner() returns()
func (joeFactory *JoeFactory) PackBecomeOwner() []byte {
	enc, err := joeFactory.abi.Pack("becomeOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBecomeOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf9dca989.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function becomeOwner() returns()
func (joeFactory *JoeFactory) TryPackBecomeOwner() ([]byte, error) {
	return joeFactory.abi.Pack("becomeOwner")
}

// PackCreateLBPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x659ac74b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (joeFactory *JoeFactory) PackCreateLBPair(tokenX common.Address, tokenY common.Address, activeId *big.Int, binStep uint16) []byte {
	enc, err := joeFactory.abi.Pack("createLBPair", tokenX, tokenY, activeId, binStep)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreateLBPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x659ac74b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (joeFactory *JoeFactory) TryPackCreateLBPair(tokenX common.Address, tokenY common.Address, activeId *big.Int, binStep uint16) ([]byte, error) {
	return joeFactory.abi.Pack("createLBPair", tokenX, tokenY, activeId, binStep)
}

// UnpackCreateLBPair is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x659ac74b.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (joeFactory *JoeFactory) UnpackCreateLBPair(data []byte) (common.Address, error) {
	out, err := joeFactory.abi.Unpack("createLBPair", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackForceDecay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c78a941.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function forceDecay(address pair) returns()
func (joeFactory *JoeFactory) PackForceDecay(pair common.Address) []byte {
	enc, err := joeFactory.abi.Pack("forceDecay", pair)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackForceDecay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c78a941.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function forceDecay(address pair) returns()
func (joeFactory *JoeFactory) TryPackForceDecay(pair common.Address) ([]byte, error) {
	return joeFactory.abi.Pack("forceDecay", pair)
}

// PackGetAllBinSteps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b35875c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAllBinSteps() view returns(uint256[] binStepWithPreset)
func (joeFactory *JoeFactory) PackGetAllBinSteps() []byte {
	enc, err := joeFactory.abi.Pack("getAllBinSteps")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAllBinSteps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b35875c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAllBinSteps() view returns(uint256[] binStepWithPreset)
func (joeFactory *JoeFactory) TryPackGetAllBinSteps() ([]byte, error) {
	return joeFactory.abi.Pack("getAllBinSteps")
}

// UnpackGetAllBinSteps is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5b35875c.
//
// Solidity: function getAllBinSteps() view returns(uint256[] binStepWithPreset)
func (joeFactory *JoeFactory) UnpackGetAllBinSteps(data []byte) ([]*big.Int, error) {
	out, err := joeFactory.abi.Unpack("getAllBinSteps", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackGetAllLBPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6622e0d7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAllLBPairs(address tokenX, address tokenY) view returns((uint16,address,bool,bool)[] lbPairsAvailable)
func (joeFactory *JoeFactory) PackGetAllLBPairs(tokenX common.Address, tokenY common.Address) []byte {
	enc, err := joeFactory.abi.Pack("getAllLBPairs", tokenX, tokenY)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAllLBPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6622e0d7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAllLBPairs(address tokenX, address tokenY) view returns((uint16,address,bool,bool)[] lbPairsAvailable)
func (joeFactory *JoeFactory) TryPackGetAllLBPairs(tokenX common.Address, tokenY common.Address) ([]byte, error) {
	return joeFactory.abi.Pack("getAllLBPairs", tokenX, tokenY)
}

// UnpackGetAllLBPairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6622e0d7.
//
// Solidity: function getAllLBPairs(address tokenX, address tokenY) view returns((uint16,address,bool,bool)[] lbPairsAvailable)
func (joeFactory *JoeFactory) UnpackGetAllLBPairs(data []byte) ([]ILBFactoryLBPairInformation, error) {
	out, err := joeFactory.abi.Unpack("getAllLBPairs", data)
	if err != nil {
		return *new([]ILBFactoryLBPairInformation), err
	}
	out0 := *abi.ConvertType(out[0], new([]ILBFactoryLBPairInformation)).(*[]ILBFactoryLBPairInformation)
	return out0, nil
}

// PackGetFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4ccb20c0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getFeeRecipient() view returns(address feeRecipient)
func (joeFactory *JoeFactory) PackGetFeeRecipient() []byte {
	enc, err := joeFactory.abi.Pack("getFeeRecipient")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4ccb20c0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getFeeRecipient() view returns(address feeRecipient)
func (joeFactory *JoeFactory) TryPackGetFeeRecipient() ([]byte, error) {
	return joeFactory.abi.Pack("getFeeRecipient")
}

// UnpackGetFeeRecipient is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address feeRecipient)
func (joeFactory *JoeFactory) UnpackGetFeeRecipient(data []byte) (common.Address, error) {
	out, err := joeFactory.abi.Unpack("getFeeRecipient", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetFlashLoanFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfd90c2be.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getFlashLoanFee() view returns(uint256 flashLoanFee)
func (joeFactory *JoeFactory) PackGetFlashLoanFee() []byte {
	enc, err := joeFactory.abi.Pack("getFlashLoanFee")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetFlashLoanFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfd90c2be.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getFlashLoanFee() view returns(uint256 flashLoanFee)
func (joeFactory *JoeFactory) TryPackGetFlashLoanFee() ([]byte, error) {
	return joeFactory.abi.Pack("getFlashLoanFee")
}

// UnpackGetFlashLoanFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfd90c2be.
//
// Solidity: function getFlashLoanFee() view returns(uint256 flashLoanFee)
func (joeFactory *JoeFactory) UnpackGetFlashLoanFee(data []byte) (*big.Int, error) {
	out, err := joeFactory.abi.Unpack("getFlashLoanFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetLBPairAtIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7daf5d66.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLBPairAtIndex(uint256 index) view returns(address lbPair)
func (joeFactory *JoeFactory) PackGetLBPairAtIndex(index *big.Int) []byte {
	enc, err := joeFactory.abi.Pack("getLBPairAtIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLBPairAtIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7daf5d66.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLBPairAtIndex(uint256 index) view returns(address lbPair)
func (joeFactory *JoeFactory) TryPackGetLBPairAtIndex(index *big.Int) ([]byte, error) {
	return joeFactory.abi.Pack("getLBPairAtIndex", index)
}

// UnpackGetLBPairAtIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7daf5d66.
//
// Solidity: function getLBPairAtIndex(uint256 index) view returns(address lbPair)
func (joeFactory *JoeFactory) UnpackGetLBPairAtIndex(data []byte) (common.Address, error) {
	out, err := joeFactory.abi.Unpack("getLBPairAtIndex", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetLBPairImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf371065.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLBPairImplementation() view returns(address lbPairImplementation)
func (joeFactory *JoeFactory) PackGetLBPairImplementation() []byte {
	enc, err := joeFactory.abi.Pack("getLBPairImplementation")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLBPairImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf371065.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLBPairImplementation() view returns(address lbPairImplementation)
func (joeFactory *JoeFactory) TryPackGetLBPairImplementation() ([]byte, error) {
	return joeFactory.abi.Pack("getLBPairImplementation")
}

// UnpackGetLBPairImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaf371065.
//
// Solidity: function getLBPairImplementation() view returns(address lbPairImplementation)
func (joeFactory *JoeFactory) UnpackGetLBPairImplementation(data []byte) (common.Address, error) {
	out, err := joeFactory.abi.Unpack("getLBPairImplementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetLBPairInformation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x704037bd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLBPairInformation(address tokenA, address tokenB, uint256 binStep) view returns((uint16,address,bool,bool) lbPairInformation)
func (joeFactory *JoeFactory) PackGetLBPairInformation(tokenA common.Address, tokenB common.Address, binStep *big.Int) []byte {
	enc, err := joeFactory.abi.Pack("getLBPairInformation", tokenA, tokenB, binStep)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLBPairInformation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x704037bd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLBPairInformation(address tokenA, address tokenB, uint256 binStep) view returns((uint16,address,bool,bool) lbPairInformation)
func (joeFactory *JoeFactory) TryPackGetLBPairInformation(tokenA common.Address, tokenB common.Address, binStep *big.Int) ([]byte, error) {
	return joeFactory.abi.Pack("getLBPairInformation", tokenA, tokenB, binStep)
}

// UnpackGetLBPairInformation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x704037bd.
//
// Solidity: function getLBPairInformation(address tokenA, address tokenB, uint256 binStep) view returns((uint16,address,bool,bool) lbPairInformation)
func (joeFactory *JoeFactory) UnpackGetLBPairInformation(data []byte) (ILBFactoryLBPairInformation, error) {
	out, err := joeFactory.abi.Unpack("getLBPairInformation", data)
	if err != nil {
		return *new(ILBFactoryLBPairInformation), err
	}
	out0 := *abi.ConvertType(out[0], new(ILBFactoryLBPairInformation)).(*ILBFactoryLBPairInformation)
	return out0, nil
}

// PackGetMaxFlashLoanFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ce9aa1c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMaxFlashLoanFee() pure returns(uint256 maxFee)
func (joeFactory *JoeFactory) PackGetMaxFlashLoanFee() []byte {
	enc, err := joeFactory.abi.Pack("getMaxFlashLoanFee")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMaxFlashLoanFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8ce9aa1c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMaxFlashLoanFee() pure returns(uint256 maxFee)
func (joeFactory *JoeFactory) TryPackGetMaxFlashLoanFee() ([]byte, error) {
	return joeFactory.abi.Pack("getMaxFlashLoanFee")
}

// UnpackGetMaxFlashLoanFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8ce9aa1c.
//
// Solidity: function getMaxFlashLoanFee() pure returns(uint256 maxFee)
func (joeFactory *JoeFactory) UnpackGetMaxFlashLoanFee(data []byte) (*big.Int, error) {
	out, err := joeFactory.abi.Unpack("getMaxFlashLoanFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMinBinStep is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x701ab8c1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMinBinStep() pure returns(uint256 minBinStep)
func (joeFactory *JoeFactory) PackGetMinBinStep() []byte {
	enc, err := joeFactory.abi.Pack("getMinBinStep")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMinBinStep is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x701ab8c1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMinBinStep() pure returns(uint256 minBinStep)
func (joeFactory *JoeFactory) TryPackGetMinBinStep() ([]byte, error) {
	return joeFactory.abi.Pack("getMinBinStep")
}

// UnpackGetMinBinStep is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x701ab8c1.
//
// Solidity: function getMinBinStep() pure returns(uint256 minBinStep)
func (joeFactory *JoeFactory) UnpackGetMinBinStep(data []byte) (*big.Int, error) {
	out, err := joeFactory.abi.Unpack("getMinBinStep", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetNumberOfLBPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4e937c3a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNumberOfLBPairs() view returns(uint256 lbPairNumber)
func (joeFactory *JoeFactory) PackGetNumberOfLBPairs() []byte {
	enc, err := joeFactory.abi.Pack("getNumberOfLBPairs")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNumberOfLBPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4e937c3a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNumberOfLBPairs() view returns(uint256 lbPairNumber)
func (joeFactory *JoeFactory) TryPackGetNumberOfLBPairs() ([]byte, error) {
	return joeFactory.abi.Pack("getNumberOfLBPairs")
}

// UnpackGetNumberOfLBPairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4e937c3a.
//
// Solidity: function getNumberOfLBPairs() view returns(uint256 lbPairNumber)
func (joeFactory *JoeFactory) UnpackGetNumberOfLBPairs(data []byte) (*big.Int, error) {
	out, err := joeFactory.abi.Unpack("getNumberOfLBPairs", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetNumberOfQuoteAssets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x80c5061e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNumberOfQuoteAssets() view returns(uint256 numberOfQuoteAssets)
func (joeFactory *JoeFactory) PackGetNumberOfQuoteAssets() []byte {
	enc, err := joeFactory.abi.Pack("getNumberOfQuoteAssets")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNumberOfQuoteAssets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x80c5061e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNumberOfQuoteAssets() view returns(uint256 numberOfQuoteAssets)
func (joeFactory *JoeFactory) TryPackGetNumberOfQuoteAssets() ([]byte, error) {
	return joeFactory.abi.Pack("getNumberOfQuoteAssets")
}

// UnpackGetNumberOfQuoteAssets is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x80c5061e.
//
// Solidity: function getNumberOfQuoteAssets() view returns(uint256 numberOfQuoteAssets)
func (joeFactory *JoeFactory) UnpackGetNumberOfQuoteAssets(data []byte) (*big.Int, error) {
	out, err := joeFactory.abi.Unpack("getNumberOfQuoteAssets", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetOpenBinSteps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0282c9c1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getOpenBinSteps() view returns(uint256[] openBinStep)
func (joeFactory *JoeFactory) PackGetOpenBinSteps() []byte {
	enc, err := joeFactory.abi.Pack("getOpenBinSteps")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetOpenBinSteps is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0282c9c1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getOpenBinSteps() view returns(uint256[] openBinStep)
func (joeFactory *JoeFactory) TryPackGetOpenBinSteps() ([]byte, error) {
	return joeFactory.abi.Pack("getOpenBinSteps")
}

// UnpackGetOpenBinSteps is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0282c9c1.
//
// Solidity: function getOpenBinSteps() view returns(uint256[] openBinStep)
func (joeFactory *JoeFactory) UnpackGetOpenBinSteps(data []byte) ([]*big.Int, error) {
	out, err := joeFactory.abi.Unpack("getOpenBinSteps", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackGetPreset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaabc4b3c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPreset(uint256 binStep) view returns(uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulator, bool isOpen)
func (joeFactory *JoeFactory) PackGetPreset(binStep *big.Int) []byte {
	enc, err := joeFactory.abi.Pack("getPreset", binStep)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPreset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaabc4b3c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPreset(uint256 binStep) view returns(uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulator, bool isOpen)
func (joeFactory *JoeFactory) TryPackGetPreset(binStep *big.Int) ([]byte, error) {
	return joeFactory.abi.Pack("getPreset", binStep)
}

// GetPresetOutput serves as a container for the return parameters of contract
// method GetPreset.
type GetPresetOutput struct {
	BaseFactor               *big.Int
	FilterPeriod             *big.Int
	DecayPeriod              *big.Int
	ReductionFactor          *big.Int
	VariableFeeControl       *big.Int
	ProtocolShare            *big.Int
	MaxVolatilityAccumulator *big.Int
	IsOpen                   bool
}

// UnpackGetPreset is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaabc4b3c.
//
// Solidity: function getPreset(uint256 binStep) view returns(uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulator, bool isOpen)
func (joeFactory *JoeFactory) UnpackGetPreset(data []byte) (GetPresetOutput, error) {
	out, err := joeFactory.abi.Unpack("getPreset", data)
	outstruct := new(GetPresetOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.BaseFactor = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.FilterPeriod = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.DecayPeriod = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	outstruct.ReductionFactor = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.VariableFeeControl = abi.ConvertType(out[4], new(big.Int)).(*big.Int)
	outstruct.ProtocolShare = abi.ConvertType(out[5], new(big.Int)).(*big.Int)
	outstruct.MaxVolatilityAccumulator = abi.ConvertType(out[6], new(big.Int)).(*big.Int)
	outstruct.IsOpen = *abi.ConvertType(out[7], new(bool)).(*bool)
	return *outstruct, nil
}

// PackGetQuoteAssetAtIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0752092b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getQuoteAssetAtIndex(uint256 index) view returns(address asset)
func (joeFactory *JoeFactory) PackGetQuoteAssetAtIndex(index *big.Int) []byte {
	enc, err := joeFactory.abi.Pack("getQuoteAssetAtIndex", index)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetQuoteAssetAtIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0752092b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getQuoteAssetAtIndex(uint256 index) view returns(address asset)
func (joeFactory *JoeFactory) TryPackGetQuoteAssetAtIndex(index *big.Int) ([]byte, error) {
	return joeFactory.abi.Pack("getQuoteAssetAtIndex", index)
}

// UnpackGetQuoteAssetAtIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0752092b.
//
// Solidity: function getQuoteAssetAtIndex(uint256 index) view returns(address asset)
func (joeFactory *JoeFactory) UnpackGetQuoteAssetAtIndex(data []byte) (common.Address, error) {
	out, err := joeFactory.abi.Unpack("getQuoteAssetAtIndex", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsQuoteAsset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x27721842.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isQuoteAsset(address token) view returns(bool isQuote)
func (joeFactory *JoeFactory) PackIsQuoteAsset(token common.Address) []byte {
	enc, err := joeFactory.abi.Pack("isQuoteAsset", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsQuoteAsset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x27721842.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isQuoteAsset(address token) view returns(bool isQuote)
func (joeFactory *JoeFactory) TryPackIsQuoteAsset(token common.Address) ([]byte, error) {
	return joeFactory.abi.Pack("isQuoteAsset", token)
}

// UnpackIsQuoteAsset is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x27721842.
//
// Solidity: function isQuoteAsset(address token) view returns(bool isQuote)
func (joeFactory *JoeFactory) UnpackIsQuoteAsset(data []byte) (bool, error) {
	out, err := joeFactory.abi.Unpack("isQuoteAsset", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (joeFactory *JoeFactory) PackOwner() []byte {
	enc, err := joeFactory.abi.Pack("owner")
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
func (joeFactory *JoeFactory) TryPackOwner() ([]byte, error) {
	return joeFactory.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (joeFactory *JoeFactory) UnpackOwner(data []byte) (common.Address, error) {
	out, err := joeFactory.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pendingOwner() view returns(address)
func (joeFactory *JoeFactory) PackPendingOwner() []byte {
	enc, err := joeFactory.abi.Pack("pendingOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe30c3978.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pendingOwner() view returns(address)
func (joeFactory *JoeFactory) TryPackPendingOwner() ([]byte, error) {
	return joeFactory.abi.Pack("pendingOwner")
}

// UnpackPendingOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (joeFactory *JoeFactory) UnpackPendingOwner(data []byte) (common.Address, error) {
	out, err := joeFactory.abi.Unpack("pendingOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRemovePreset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe203a31f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removePreset(uint16 binStep) returns()
func (joeFactory *JoeFactory) PackRemovePreset(binStep uint16) []byte {
	enc, err := joeFactory.abi.Pack("removePreset", binStep)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemovePreset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe203a31f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removePreset(uint16 binStep) returns()
func (joeFactory *JoeFactory) TryPackRemovePreset(binStep uint16) ([]byte, error) {
	return joeFactory.abi.Pack("removePreset", binStep)
}

// PackRemoveQuoteAsset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xddbfd941.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeQuoteAsset(address quoteAsset) returns()
func (joeFactory *JoeFactory) PackRemoveQuoteAsset(quoteAsset common.Address) []byte {
	enc, err := joeFactory.abi.Pack("removeQuoteAsset", quoteAsset)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveQuoteAsset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xddbfd941.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeQuoteAsset(address quoteAsset) returns()
func (joeFactory *JoeFactory) TryPackRemoveQuoteAsset(quoteAsset common.Address) ([]byte, error) {
	return joeFactory.abi.Pack("removeQuoteAsset", quoteAsset)
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (joeFactory *JoeFactory) PackRenounceOwnership() []byte {
	enc, err := joeFactory.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function renounceOwnership() returns()
func (joeFactory *JoeFactory) TryPackRenounceOwnership() ([]byte, error) {
	return joeFactory.abi.Pack("renounceOwnership")
}

// PackRevokePendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67ab8a4e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function revokePendingOwner() returns()
func (joeFactory *JoeFactory) PackRevokePendingOwner() []byte {
	enc, err := joeFactory.abi.Pack("revokePendingOwner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRevokePendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67ab8a4e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function revokePendingOwner() returns()
func (joeFactory *JoeFactory) TryPackRevokePendingOwner() ([]byte, error) {
	return joeFactory.abi.Pack("revokePendingOwner")
}

// PackSetFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe74b981b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFeeRecipient(address feeRecipient) returns()
func (joeFactory *JoeFactory) PackSetFeeRecipient(feeRecipient common.Address) []byte {
	enc, err := joeFactory.abi.Pack("setFeeRecipient", feeRecipient)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFeeRecipient is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe74b981b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFeeRecipient(address feeRecipient) returns()
func (joeFactory *JoeFactory) TryPackSetFeeRecipient(feeRecipient common.Address) ([]byte, error) {
	return joeFactory.abi.Pack("setFeeRecipient", feeRecipient)
}

// PackSetFeesParametersOnPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x093ff769.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFeesParametersOnPair(address tokenX, address tokenY, uint16 binStep, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator) returns()
func (joeFactory *JoeFactory) PackSetFeesParametersOnPair(tokenX common.Address, tokenY common.Address, binStep uint16, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int) []byte {
	enc, err := joeFactory.abi.Pack("setFeesParametersOnPair", tokenX, tokenY, binStep, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFeesParametersOnPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x093ff769.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFeesParametersOnPair(address tokenX, address tokenY, uint16 binStep, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator) returns()
func (joeFactory *JoeFactory) TryPackSetFeesParametersOnPair(tokenX common.Address, tokenY common.Address, binStep uint16, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int) ([]byte, error) {
	return joeFactory.abi.Pack("setFeesParametersOnPair", tokenX, tokenY, binStep, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator)
}

// PackSetFlashLoanFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe92d0d5d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFlashLoanFee(uint256 flashLoanFee) returns()
func (joeFactory *JoeFactory) PackSetFlashLoanFee(flashLoanFee *big.Int) []byte {
	enc, err := joeFactory.abi.Pack("setFlashLoanFee", flashLoanFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFlashLoanFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe92d0d5d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFlashLoanFee(uint256 flashLoanFee) returns()
func (joeFactory *JoeFactory) TryPackSetFlashLoanFee(flashLoanFee *big.Int) ([]byte, error) {
	return joeFactory.abi.Pack("setFlashLoanFee", flashLoanFee)
}

// PackSetLBPairIgnored is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x69d56ea3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setLBPairIgnored(address tokenX, address tokenY, uint16 binStep, bool ignored) returns()
func (joeFactory *JoeFactory) PackSetLBPairIgnored(tokenX common.Address, tokenY common.Address, binStep uint16, ignored bool) []byte {
	enc, err := joeFactory.abi.Pack("setLBPairIgnored", tokenX, tokenY, binStep, ignored)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetLBPairIgnored is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x69d56ea3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setLBPairIgnored(address tokenX, address tokenY, uint16 binStep, bool ignored) returns()
func (joeFactory *JoeFactory) TryPackSetLBPairIgnored(tokenX common.Address, tokenY common.Address, binStep uint16, ignored bool) ([]byte, error) {
	return joeFactory.abi.Pack("setLBPairIgnored", tokenX, tokenY, binStep, ignored)
}

// PackSetLBPairImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb0384781.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setLBPairImplementation(address newLBPairImplementation) returns()
func (joeFactory *JoeFactory) PackSetLBPairImplementation(newLBPairImplementation common.Address) []byte {
	enc, err := joeFactory.abi.Pack("setLBPairImplementation", newLBPairImplementation)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetLBPairImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb0384781.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setLBPairImplementation(address newLBPairImplementation) returns()
func (joeFactory *JoeFactory) TryPackSetLBPairImplementation(newLBPairImplementation common.Address) ([]byte, error) {
	return joeFactory.abi.Pack("setLBPairImplementation", newLBPairImplementation)
}

// PackSetPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc42069ec.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setPendingOwner(address pendingOwner_) returns()
func (joeFactory *JoeFactory) PackSetPendingOwner(pendingOwner common.Address) []byte {
	enc, err := joeFactory.abi.Pack("setPendingOwner", pendingOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetPendingOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc42069ec.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setPendingOwner(address pendingOwner_) returns()
func (joeFactory *JoeFactory) TryPackSetPendingOwner(pendingOwner common.Address) ([]byte, error) {
	return joeFactory.abi.Pack("setPendingOwner", pendingOwner)
}

// PackSetPreset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x379ee803.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setPreset(uint16 binStep, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator, bool isOpen) returns()
func (joeFactory *JoeFactory) PackSetPreset(binStep uint16, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int, isOpen bool) []byte {
	enc, err := joeFactory.abi.Pack("setPreset", binStep, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator, isOpen)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetPreset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x379ee803.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setPreset(uint16 binStep, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator, bool isOpen) returns()
func (joeFactory *JoeFactory) TryPackSetPreset(binStep uint16, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int, isOpen bool) ([]byte, error) {
	return joeFactory.abi.Pack("setPreset", binStep, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator, isOpen)
}

// PackSetPresetOpenState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4cd161d3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setPresetOpenState(uint16 binStep, bool isOpen) returns()
func (joeFactory *JoeFactory) PackSetPresetOpenState(binStep uint16, isOpen bool) []byte {
	enc, err := joeFactory.abi.Pack("setPresetOpenState", binStep, isOpen)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetPresetOpenState is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4cd161d3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setPresetOpenState(uint16 binStep, bool isOpen) returns()
func (joeFactory *JoeFactory) TryPackSetPresetOpenState(binStep uint16, isOpen bool) ([]byte, error) {
	return joeFactory.abi.Pack("setPresetOpenState", binStep, isOpen)
}

// JoeFactoryFeeRecipientSet represents a FeeRecipientSet event raised by the JoeFactory contract.
type JoeFactoryFeeRecipientSet struct {
	OldRecipient common.Address
	NewRecipient common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const JoeFactoryFeeRecipientSetEventName = "FeeRecipientSet"

// ContractEventName returns the user-defined event name.
func (JoeFactoryFeeRecipientSet) ContractEventName() string {
	return JoeFactoryFeeRecipientSetEventName
}

// UnpackFeeRecipientSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeRecipientSet(address oldRecipient, address newRecipient)
func (joeFactory *JoeFactory) UnpackFeeRecipientSetEvent(log *types.Log) (*JoeFactoryFeeRecipientSet, error) {
	event := "FeeRecipientSet"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryFeeRecipientSet)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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

// JoeFactoryFlashLoanFeeSet represents a FlashLoanFeeSet event raised by the JoeFactory contract.
type JoeFactoryFlashLoanFeeSet struct {
	OldFlashLoanFee *big.Int
	NewFlashLoanFee *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const JoeFactoryFlashLoanFeeSetEventName = "FlashLoanFeeSet"

// ContractEventName returns the user-defined event name.
func (JoeFactoryFlashLoanFeeSet) ContractEventName() string {
	return JoeFactoryFlashLoanFeeSetEventName
}

// UnpackFlashLoanFeeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FlashLoanFeeSet(uint256 oldFlashLoanFee, uint256 newFlashLoanFee)
func (joeFactory *JoeFactory) UnpackFlashLoanFeeSetEvent(log *types.Log) (*JoeFactoryFlashLoanFeeSet, error) {
	event := "FlashLoanFeeSet"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryFlashLoanFeeSet)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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

// JoeFactoryLBPairCreated represents a LBPairCreated event raised by the JoeFactory contract.
type JoeFactoryLBPairCreated struct {
	TokenX  common.Address
	TokenY  common.Address
	BinStep *big.Int
	LBPair  common.Address
	Pid     *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const JoeFactoryLBPairCreatedEventName = "LBPairCreated"

// ContractEventName returns the user-defined event name.
func (JoeFactoryLBPairCreated) ContractEventName() string {
	return JoeFactoryLBPairCreatedEventName
}

// UnpackLBPairCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LBPairCreated(address indexed tokenX, address indexed tokenY, uint256 indexed binStep, address LBPair, uint256 pid)
func (joeFactory *JoeFactory) UnpackLBPairCreatedEvent(log *types.Log) (*JoeFactoryLBPairCreated, error) {
	event := "LBPairCreated"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryLBPairCreated)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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

// JoeFactoryLBPairIgnoredStateChanged represents a LBPairIgnoredStateChanged event raised by the JoeFactory contract.
type JoeFactoryLBPairIgnoredStateChanged struct {
	LBPair  common.Address
	Ignored bool
	Raw     *types.Log // Blockchain specific contextual infos
}

const JoeFactoryLBPairIgnoredStateChangedEventName = "LBPairIgnoredStateChanged"

// ContractEventName returns the user-defined event name.
func (JoeFactoryLBPairIgnoredStateChanged) ContractEventName() string {
	return JoeFactoryLBPairIgnoredStateChangedEventName
}

// UnpackLBPairIgnoredStateChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LBPairIgnoredStateChanged(address indexed LBPair, bool ignored)
func (joeFactory *JoeFactory) UnpackLBPairIgnoredStateChangedEvent(log *types.Log) (*JoeFactoryLBPairIgnoredStateChanged, error) {
	event := "LBPairIgnoredStateChanged"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryLBPairIgnoredStateChanged)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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

// JoeFactoryLBPairImplementationSet represents a LBPairImplementationSet event raised by the JoeFactory contract.
type JoeFactoryLBPairImplementationSet struct {
	OldLBPairImplementation common.Address
	LBPairImplementation    common.Address
	Raw                     *types.Log // Blockchain specific contextual infos
}

const JoeFactoryLBPairImplementationSetEventName = "LBPairImplementationSet"

// ContractEventName returns the user-defined event name.
func (JoeFactoryLBPairImplementationSet) ContractEventName() string {
	return JoeFactoryLBPairImplementationSetEventName
}

// UnpackLBPairImplementationSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event LBPairImplementationSet(address oldLBPairImplementation, address LBPairImplementation)
func (joeFactory *JoeFactory) UnpackLBPairImplementationSetEvent(log *types.Log) (*JoeFactoryLBPairImplementationSet, error) {
	event := "LBPairImplementationSet"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryLBPairImplementationSet)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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

// JoeFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the JoeFactory contract.
type JoeFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const JoeFactoryOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (JoeFactoryOwnershipTransferred) ContractEventName() string {
	return JoeFactoryOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (joeFactory *JoeFactory) UnpackOwnershipTransferredEvent(log *types.Log) (*JoeFactoryOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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

// JoeFactoryPendingOwnerSet represents a PendingOwnerSet event raised by the JoeFactory contract.
type JoeFactoryPendingOwnerSet struct {
	PendingOwner common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const JoeFactoryPendingOwnerSetEventName = "PendingOwnerSet"

// ContractEventName returns the user-defined event name.
func (JoeFactoryPendingOwnerSet) ContractEventName() string {
	return JoeFactoryPendingOwnerSetEventName
}

// UnpackPendingOwnerSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PendingOwnerSet(address indexed pendingOwner)
func (joeFactory *JoeFactory) UnpackPendingOwnerSetEvent(log *types.Log) (*JoeFactoryPendingOwnerSet, error) {
	event := "PendingOwnerSet"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryPendingOwnerSet)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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

// JoeFactoryPresetOpenStateChanged represents a PresetOpenStateChanged event raised by the JoeFactory contract.
type JoeFactoryPresetOpenStateChanged struct {
	BinStep *big.Int
	IsOpen  bool
	Raw     *types.Log // Blockchain specific contextual infos
}

const JoeFactoryPresetOpenStateChangedEventName = "PresetOpenStateChanged"

// ContractEventName returns the user-defined event name.
func (JoeFactoryPresetOpenStateChanged) ContractEventName() string {
	return JoeFactoryPresetOpenStateChangedEventName
}

// UnpackPresetOpenStateChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PresetOpenStateChanged(uint256 indexed binStep, bool indexed isOpen)
func (joeFactory *JoeFactory) UnpackPresetOpenStateChangedEvent(log *types.Log) (*JoeFactoryPresetOpenStateChanged, error) {
	event := "PresetOpenStateChanged"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryPresetOpenStateChanged)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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

// JoeFactoryPresetRemoved represents a PresetRemoved event raised by the JoeFactory contract.
type JoeFactoryPresetRemoved struct {
	BinStep *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const JoeFactoryPresetRemovedEventName = "PresetRemoved"

// ContractEventName returns the user-defined event name.
func (JoeFactoryPresetRemoved) ContractEventName() string {
	return JoeFactoryPresetRemovedEventName
}

// UnpackPresetRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PresetRemoved(uint256 indexed binStep)
func (joeFactory *JoeFactory) UnpackPresetRemovedEvent(log *types.Log) (*JoeFactoryPresetRemoved, error) {
	event := "PresetRemoved"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryPresetRemoved)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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

// JoeFactoryPresetSet represents a PresetSet event raised by the JoeFactory contract.
type JoeFactoryPresetSet struct {
	BinStep                  *big.Int
	BaseFactor               *big.Int
	FilterPeriod             *big.Int
	DecayPeriod              *big.Int
	ReductionFactor          *big.Int
	VariableFeeControl       *big.Int
	ProtocolShare            *big.Int
	MaxVolatilityAccumulator *big.Int
	Raw                      *types.Log // Blockchain specific contextual infos
}

const JoeFactoryPresetSetEventName = "PresetSet"

// ContractEventName returns the user-defined event name.
func (JoeFactoryPresetSet) ContractEventName() string {
	return JoeFactoryPresetSetEventName
}

// UnpackPresetSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PresetSet(uint256 indexed binStep, uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulator)
func (joeFactory *JoeFactory) UnpackPresetSetEvent(log *types.Log) (*JoeFactoryPresetSet, error) {
	event := "PresetSet"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryPresetSet)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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

// JoeFactoryQuoteAssetAdded represents a QuoteAssetAdded event raised by the JoeFactory contract.
type JoeFactoryQuoteAssetAdded struct {
	QuoteAsset common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const JoeFactoryQuoteAssetAddedEventName = "QuoteAssetAdded"

// ContractEventName returns the user-defined event name.
func (JoeFactoryQuoteAssetAdded) ContractEventName() string {
	return JoeFactoryQuoteAssetAddedEventName
}

// UnpackQuoteAssetAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuoteAssetAdded(address indexed quoteAsset)
func (joeFactory *JoeFactory) UnpackQuoteAssetAddedEvent(log *types.Log) (*JoeFactoryQuoteAssetAdded, error) {
	event := "QuoteAssetAdded"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryQuoteAssetAdded)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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

// JoeFactoryQuoteAssetRemoved represents a QuoteAssetRemoved event raised by the JoeFactory contract.
type JoeFactoryQuoteAssetRemoved struct {
	QuoteAsset common.Address
	Raw        *types.Log // Blockchain specific contextual infos
}

const JoeFactoryQuoteAssetRemovedEventName = "QuoteAssetRemoved"

// ContractEventName returns the user-defined event name.
func (JoeFactoryQuoteAssetRemoved) ContractEventName() string {
	return JoeFactoryQuoteAssetRemovedEventName
}

// UnpackQuoteAssetRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event QuoteAssetRemoved(address indexed quoteAsset)
func (joeFactory *JoeFactory) UnpackQuoteAssetRemovedEvent(log *types.Log) (*JoeFactoryQuoteAssetRemoved, error) {
	event := "QuoteAssetRemoved"
	if log.Topics[0] != joeFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(JoeFactoryQuoteAssetRemoved)
	if len(log.Data) > 0 {
		if err := joeFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range joeFactory.abi.Events[event].Inputs {
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
func (joeFactory *JoeFactory) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryAddressZero"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryBinStepHasNoPreset"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryBinStepHasNoPresetError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryBinStepTooLow"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryBinStepTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryFlashLoanFeeAboveMax"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryFlashLoanFeeAboveMaxError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryIdenticalAddresses"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryIdenticalAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryImplementationNotSet"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryImplementationNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryLBPairAlreadyExists"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryLBPairAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryLBPairDoesNotExist"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryLBPairDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryLBPairIgnoredIsAlreadyInTheSameState"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryLBPairIgnoredIsAlreadyInTheSameStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryLBPairNotCreated"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryLBPairNotCreatedError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryLBPairSafetyCheckFailed"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryLBPairSafetyCheckFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryPresetIsLockedForUsers"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryPresetIsLockedForUsersError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryPresetOpenStateIsAlreadyInTheSameState"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryPresetOpenStateIsAlreadyInTheSameStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryQuoteAssetAlreadyWhitelisted"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryQuoteAssetAlreadyWhitelistedError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactoryQuoteAssetNotWhitelisted"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactoryQuoteAssetNotWhitelistedError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactorySameFeeRecipient"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactorySameFeeRecipientError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactorySameFlashLoanFee"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactorySameFlashLoanFeeError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["LBFactorySameImplementation"].ID.Bytes()[:4]) {
		return joeFactory.UnpackLBFactorySameImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["PairParametersHelperInvalidParameter"].ID.Bytes()[:4]) {
		return joeFactory.UnpackPairParametersHelperInvalidParameterError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["PendingOwnableAddressZero"].ID.Bytes()[:4]) {
		return joeFactory.UnpackPendingOwnableAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["PendingOwnableNoPendingOwner"].ID.Bytes()[:4]) {
		return joeFactory.UnpackPendingOwnableNoPendingOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["PendingOwnableNotOwner"].ID.Bytes()[:4]) {
		return joeFactory.UnpackPendingOwnableNotOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["PendingOwnableNotPendingOwner"].ID.Bytes()[:4]) {
		return joeFactory.UnpackPendingOwnableNotPendingOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["PendingOwnablePendingOwnerAlreadySet"].ID.Bytes()[:4]) {
		return joeFactory.UnpackPendingOwnablePendingOwnerAlreadySetError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["SafeCastExceeds16Bits"].ID.Bytes()[:4]) {
		return joeFactory.UnpackSafeCastExceeds16BitsError(raw[4:])
	}
	if bytes.Equal(raw[:4], joeFactory.abi.Errors["Uint128x128MathPowUnderflow"].ID.Bytes()[:4]) {
		return joeFactory.UnpackUint128x128MathPowUnderflowError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// JoeFactoryLBFactoryAddressZero represents a LBFactory__AddressZero error raised by the JoeFactory contract.
type JoeFactoryLBFactoryAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__AddressZero()
func JoeFactoryLBFactoryAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x95cf3ee426a9a8fb0d88a6a89d772d47913ea3716bcd4ee9b8e55a4e1623d3a3")
}

// UnpackLBFactoryAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__AddressZero()
func (joeFactory *JoeFactory) UnpackLBFactoryAddressZeroError(raw []byte) (*JoeFactoryLBFactoryAddressZero, error) {
	out := new(JoeFactoryLBFactoryAddressZero)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryBinStepHasNoPreset represents a LBFactory__BinStepHasNoPreset error raised by the JoeFactory contract.
type JoeFactoryLBFactoryBinStepHasNoPreset struct {
	BinStep *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__BinStepHasNoPreset(uint256 binStep)
func JoeFactoryLBFactoryBinStepHasNoPresetErrorID() common.Hash {
	return common.HexToHash("0xfb22c17ea6871cce960eeb0f571efea1ee47141ab54cce694d172ff2a555414a")
}

// UnpackLBFactoryBinStepHasNoPresetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__BinStepHasNoPreset(uint256 binStep)
func (joeFactory *JoeFactory) UnpackLBFactoryBinStepHasNoPresetError(raw []byte) (*JoeFactoryLBFactoryBinStepHasNoPreset, error) {
	out := new(JoeFactoryLBFactoryBinStepHasNoPreset)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryBinStepHasNoPreset", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryBinStepTooLow represents a LBFactory__BinStepTooLow error raised by the JoeFactory contract.
type JoeFactoryLBFactoryBinStepTooLow struct {
	BinStep *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__BinStepTooLow(uint256 binStep)
func JoeFactoryLBFactoryBinStepTooLowErrorID() common.Hash {
	return common.HexToHash("0x4f958e71e0afd0bc4d2b7428ca60875ddced91e100f35580f0b3142453c8f905")
}

// UnpackLBFactoryBinStepTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__BinStepTooLow(uint256 binStep)
func (joeFactory *JoeFactory) UnpackLBFactoryBinStepTooLowError(raw []byte) (*JoeFactoryLBFactoryBinStepTooLow, error) {
	out := new(JoeFactoryLBFactoryBinStepTooLow)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryBinStepTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryFlashLoanFeeAboveMax represents a LBFactory__FlashLoanFeeAboveMax error raised by the JoeFactory contract.
type JoeFactoryLBFactoryFlashLoanFeeAboveMax struct {
	Fees    *big.Int
	MaxFees *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__FlashLoanFeeAboveMax(uint256 fees, uint256 maxFees)
func JoeFactoryLBFactoryFlashLoanFeeAboveMaxErrorID() common.Hash {
	return common.HexToHash("0x5e8988c1fe2fcc1d77c82675e5a641411a5d072f369593ed8886b952d8bc677d")
}

// UnpackLBFactoryFlashLoanFeeAboveMaxError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__FlashLoanFeeAboveMax(uint256 fees, uint256 maxFees)
func (joeFactory *JoeFactory) UnpackLBFactoryFlashLoanFeeAboveMaxError(raw []byte) (*JoeFactoryLBFactoryFlashLoanFeeAboveMax, error) {
	out := new(JoeFactoryLBFactoryFlashLoanFeeAboveMax)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryFlashLoanFeeAboveMax", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryIdenticalAddresses represents a LBFactory__IdenticalAddresses error raised by the JoeFactory contract.
type JoeFactoryLBFactoryIdenticalAddresses struct {
	Token common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__IdenticalAddresses(address token)
func JoeFactoryLBFactoryIdenticalAddressesErrorID() common.Hash {
	return common.HexToHash("0x2f9b18530bc9f80a3f3be00112d250c0ffeb27f86c0bb6ee167f6f409d8ed66c")
}

// UnpackLBFactoryIdenticalAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__IdenticalAddresses(address token)
func (joeFactory *JoeFactory) UnpackLBFactoryIdenticalAddressesError(raw []byte) (*JoeFactoryLBFactoryIdenticalAddresses, error) {
	out := new(JoeFactoryLBFactoryIdenticalAddresses)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryIdenticalAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryImplementationNotSet represents a LBFactory__ImplementationNotSet error raised by the JoeFactory contract.
type JoeFactoryLBFactoryImplementationNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__ImplementationNotSet()
func JoeFactoryLBFactoryImplementationNotSetErrorID() common.Hash {
	return common.HexToHash("0xa2d3f3e4f4a6cdfa14f29f61c3cf0fe24b26f8d3abf9199ee8e1ecf1231afccc")
}

// UnpackLBFactoryImplementationNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__ImplementationNotSet()
func (joeFactory *JoeFactory) UnpackLBFactoryImplementationNotSetError(raw []byte) (*JoeFactoryLBFactoryImplementationNotSet, error) {
	out := new(JoeFactoryLBFactoryImplementationNotSet)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryImplementationNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryLBPairAlreadyExists represents a LBFactory__LBPairAlreadyExists error raised by the JoeFactory contract.
type JoeFactoryLBFactoryLBPairAlreadyExists struct {
	TokenX  common.Address
	TokenY  common.Address
	BinStep *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__LBPairAlreadyExists(address tokenX, address tokenY, uint256 _binStep)
func JoeFactoryLBFactoryLBPairAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0xcb27a4351d737a0ff50c3dbad28939649bd2c4f1ee4f19f125b36e87053fbdfa")
}

// UnpackLBFactoryLBPairAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__LBPairAlreadyExists(address tokenX, address tokenY, uint256 _binStep)
func (joeFactory *JoeFactory) UnpackLBFactoryLBPairAlreadyExistsError(raw []byte) (*JoeFactoryLBFactoryLBPairAlreadyExists, error) {
	out := new(JoeFactoryLBFactoryLBPairAlreadyExists)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryLBPairAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryLBPairDoesNotExist represents a LBFactory__LBPairDoesNotExist error raised by the JoeFactory contract.
type JoeFactoryLBFactoryLBPairDoesNotExist struct {
	TokenX  common.Address
	TokenY  common.Address
	BinStep *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__LBPairDoesNotExist(address tokenX, address tokenY, uint256 binStep)
func JoeFactoryLBFactoryLBPairDoesNotExistErrorID() common.Hash {
	return common.HexToHash("0x40aa464491ae651910569d649a0c4a1290d738ed935d8a3ab954ec3cab830807")
}

// UnpackLBFactoryLBPairDoesNotExistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__LBPairDoesNotExist(address tokenX, address tokenY, uint256 binStep)
func (joeFactory *JoeFactory) UnpackLBFactoryLBPairDoesNotExistError(raw []byte) (*JoeFactoryLBFactoryLBPairDoesNotExist, error) {
	out := new(JoeFactoryLBFactoryLBPairDoesNotExist)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryLBPairDoesNotExist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryLBPairIgnoredIsAlreadyInTheSameState represents a LBFactory__LBPairIgnoredIsAlreadyInTheSameState error raised by the JoeFactory contract.
type JoeFactoryLBFactoryLBPairIgnoredIsAlreadyInTheSameState struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__LBPairIgnoredIsAlreadyInTheSameState()
func JoeFactoryLBFactoryLBPairIgnoredIsAlreadyInTheSameStateErrorID() common.Hash {
	return common.HexToHash("0x00ddcccaec6f4a336ff38228d67d86dc754c4eacf872f5c1404ac2d078b8e1d5")
}

// UnpackLBFactoryLBPairIgnoredIsAlreadyInTheSameStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__LBPairIgnoredIsAlreadyInTheSameState()
func (joeFactory *JoeFactory) UnpackLBFactoryLBPairIgnoredIsAlreadyInTheSameStateError(raw []byte) (*JoeFactoryLBFactoryLBPairIgnoredIsAlreadyInTheSameState, error) {
	out := new(JoeFactoryLBFactoryLBPairIgnoredIsAlreadyInTheSameState)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryLBPairIgnoredIsAlreadyInTheSameState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryLBPairNotCreated represents a LBFactory__LBPairNotCreated error raised by the JoeFactory contract.
type JoeFactoryLBFactoryLBPairNotCreated struct {
	TokenX  common.Address
	TokenY  common.Address
	BinStep *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__LBPairNotCreated(address tokenX, address tokenY, uint256 binStep)
func JoeFactoryLBFactoryLBPairNotCreatedErrorID() common.Hash {
	return common.HexToHash("0xb65ee953d97bb6da6189040df0137e54fd24406e395285586e73d50b34fe580d")
}

// UnpackLBFactoryLBPairNotCreatedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__LBPairNotCreated(address tokenX, address tokenY, uint256 binStep)
func (joeFactory *JoeFactory) UnpackLBFactoryLBPairNotCreatedError(raw []byte) (*JoeFactoryLBFactoryLBPairNotCreated, error) {
	out := new(JoeFactoryLBFactoryLBPairNotCreated)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryLBPairNotCreated", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryLBPairSafetyCheckFailed represents a LBFactory__LBPairSafetyCheckFailed error raised by the JoeFactory contract.
type JoeFactoryLBFactoryLBPairSafetyCheckFailed struct {
	LBPairImplementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__LBPairSafetyCheckFailed(address LBPairImplementation)
func JoeFactoryLBFactoryLBPairSafetyCheckFailedErrorID() common.Hash {
	return common.HexToHash("0x147ce15eacd45a885d783af593b0b0d669448df3300a4acc578c4fc91c263d73")
}

// UnpackLBFactoryLBPairSafetyCheckFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__LBPairSafetyCheckFailed(address LBPairImplementation)
func (joeFactory *JoeFactory) UnpackLBFactoryLBPairSafetyCheckFailedError(raw []byte) (*JoeFactoryLBFactoryLBPairSafetyCheckFailed, error) {
	out := new(JoeFactoryLBFactoryLBPairSafetyCheckFailed)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryLBPairSafetyCheckFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryPresetIsLockedForUsers represents a LBFactory__PresetIsLockedForUsers error raised by the JoeFactory contract.
type JoeFactoryLBFactoryPresetIsLockedForUsers struct {
	User    common.Address
	BinStep *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__PresetIsLockedForUsers(address user, uint256 binStep)
func JoeFactoryLBFactoryPresetIsLockedForUsersErrorID() common.Hash {
	return common.HexToHash("0x09f85fcee71c154193f7230cb2b326aa974d4cf2325ceb0875c61cd21749d4f4")
}

// UnpackLBFactoryPresetIsLockedForUsersError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__PresetIsLockedForUsers(address user, uint256 binStep)
func (joeFactory *JoeFactory) UnpackLBFactoryPresetIsLockedForUsersError(raw []byte) (*JoeFactoryLBFactoryPresetIsLockedForUsers, error) {
	out := new(JoeFactoryLBFactoryPresetIsLockedForUsers)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryPresetIsLockedForUsers", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryPresetOpenStateIsAlreadyInTheSameState represents a LBFactory__PresetOpenStateIsAlreadyInTheSameState error raised by the JoeFactory contract.
type JoeFactoryLBFactoryPresetOpenStateIsAlreadyInTheSameState struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__PresetOpenStateIsAlreadyInTheSameState()
func JoeFactoryLBFactoryPresetOpenStateIsAlreadyInTheSameStateErrorID() common.Hash {
	return common.HexToHash("0x237c71b69c5b9673f7547276aac502f625f280a365695460d21edaa00c4dfd1b")
}

// UnpackLBFactoryPresetOpenStateIsAlreadyInTheSameStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__PresetOpenStateIsAlreadyInTheSameState()
func (joeFactory *JoeFactory) UnpackLBFactoryPresetOpenStateIsAlreadyInTheSameStateError(raw []byte) (*JoeFactoryLBFactoryPresetOpenStateIsAlreadyInTheSameState, error) {
	out := new(JoeFactoryLBFactoryPresetOpenStateIsAlreadyInTheSameState)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryPresetOpenStateIsAlreadyInTheSameState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryQuoteAssetAlreadyWhitelisted represents a LBFactory__QuoteAssetAlreadyWhitelisted error raised by the JoeFactory contract.
type JoeFactoryLBFactoryQuoteAssetAlreadyWhitelisted struct {
	QuoteAsset common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__QuoteAssetAlreadyWhitelisted(address quoteAsset)
func JoeFactoryLBFactoryQuoteAssetAlreadyWhitelistedErrorID() common.Hash {
	return common.HexToHash("0x03ce0ad98523e8eab371b68d4430b3df4f3fae5327677afb9ab0e9062958f628")
}

// UnpackLBFactoryQuoteAssetAlreadyWhitelistedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__QuoteAssetAlreadyWhitelisted(address quoteAsset)
func (joeFactory *JoeFactory) UnpackLBFactoryQuoteAssetAlreadyWhitelistedError(raw []byte) (*JoeFactoryLBFactoryQuoteAssetAlreadyWhitelisted, error) {
	out := new(JoeFactoryLBFactoryQuoteAssetAlreadyWhitelisted)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryQuoteAssetAlreadyWhitelisted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactoryQuoteAssetNotWhitelisted represents a LBFactory__QuoteAssetNotWhitelisted error raised by the JoeFactory contract.
type JoeFactoryLBFactoryQuoteAssetNotWhitelisted struct {
	QuoteAsset common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__QuoteAssetNotWhitelisted(address quoteAsset)
func JoeFactoryLBFactoryQuoteAssetNotWhitelistedErrorID() common.Hash {
	return common.HexToHash("0x8e888ef33a60481591c3ba7f8a9f3ca09aa65e94b6ae9eed665bfea2123a6daa")
}

// UnpackLBFactoryQuoteAssetNotWhitelistedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__QuoteAssetNotWhitelisted(address quoteAsset)
func (joeFactory *JoeFactory) UnpackLBFactoryQuoteAssetNotWhitelistedError(raw []byte) (*JoeFactoryLBFactoryQuoteAssetNotWhitelisted, error) {
	out := new(JoeFactoryLBFactoryQuoteAssetNotWhitelisted)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactoryQuoteAssetNotWhitelisted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactorySameFeeRecipient represents a LBFactory__SameFeeRecipient error raised by the JoeFactory contract.
type JoeFactoryLBFactorySameFeeRecipient struct {
	FeeRecipient common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__SameFeeRecipient(address feeRecipient)
func JoeFactoryLBFactorySameFeeRecipientErrorID() common.Hash {
	return common.HexToHash("0x4fcea9716e68f71f8ef2f5b29b24024c44e61604394f29d56ad0b50ff368e205")
}

// UnpackLBFactorySameFeeRecipientError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__SameFeeRecipient(address feeRecipient)
func (joeFactory *JoeFactory) UnpackLBFactorySameFeeRecipientError(raw []byte) (*JoeFactoryLBFactorySameFeeRecipient, error) {
	out := new(JoeFactoryLBFactorySameFeeRecipient)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactorySameFeeRecipient", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactorySameFlashLoanFee represents a LBFactory__SameFlashLoanFee error raised by the JoeFactory contract.
type JoeFactoryLBFactorySameFlashLoanFee struct {
	FlashLoanFee *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__SameFlashLoanFee(uint256 flashLoanFee)
func JoeFactoryLBFactorySameFlashLoanFeeErrorID() common.Hash {
	return common.HexToHash("0x6ea8c7a46ec20271923901984d6cd982fc4d3559eddf7449a9edc17f951f34c9")
}

// UnpackLBFactorySameFlashLoanFeeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__SameFlashLoanFee(uint256 flashLoanFee)
func (joeFactory *JoeFactory) UnpackLBFactorySameFlashLoanFeeError(raw []byte) (*JoeFactoryLBFactorySameFlashLoanFee, error) {
	out := new(JoeFactoryLBFactorySameFlashLoanFee)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactorySameFlashLoanFee", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryLBFactorySameImplementation represents a LBFactory__SameImplementation error raised by the JoeFactory contract.
type JoeFactoryLBFactorySameImplementation struct {
	LBPairImplementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LBFactory__SameImplementation(address LBPairImplementation)
func JoeFactoryLBFactorySameImplementationErrorID() common.Hash {
	return common.HexToHash("0x6f69dca87faf6a8b6ba04573bbbe93887545d6171153e10584cbe3a9b3599ee4")
}

// UnpackLBFactorySameImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LBFactory__SameImplementation(address LBPairImplementation)
func (joeFactory *JoeFactory) UnpackLBFactorySameImplementationError(raw []byte) (*JoeFactoryLBFactorySameImplementation, error) {
	out := new(JoeFactoryLBFactorySameImplementation)
	if err := joeFactory.abi.UnpackIntoInterface(out, "LBFactorySameImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryPairParametersHelperInvalidParameter represents a PairParametersHelper__InvalidParameter error raised by the JoeFactory contract.
type JoeFactoryPairParametersHelperInvalidParameter struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PairParametersHelper__InvalidParameter()
func JoeFactoryPairParametersHelperInvalidParameterErrorID() common.Hash {
	return common.HexToHash("0x1c07203f6db94b4be38204d02e95f6f8d1532f93488dbd6c4c558c59340ae84c")
}

// UnpackPairParametersHelperInvalidParameterError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PairParametersHelper__InvalidParameter()
func (joeFactory *JoeFactory) UnpackPairParametersHelperInvalidParameterError(raw []byte) (*JoeFactoryPairParametersHelperInvalidParameter, error) {
	out := new(JoeFactoryPairParametersHelperInvalidParameter)
	if err := joeFactory.abi.UnpackIntoInterface(out, "PairParametersHelperInvalidParameter", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryPendingOwnableAddressZero represents a PendingOwnable__AddressZero error raised by the JoeFactory contract.
type JoeFactoryPendingOwnableAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PendingOwnable__AddressZero()
func JoeFactoryPendingOwnableAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x91f38515a80fb31d7d944ac161cc931bbada62b5516c273de86121c3000f77f7")
}

// UnpackPendingOwnableAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PendingOwnable__AddressZero()
func (joeFactory *JoeFactory) UnpackPendingOwnableAddressZeroError(raw []byte) (*JoeFactoryPendingOwnableAddressZero, error) {
	out := new(JoeFactoryPendingOwnableAddressZero)
	if err := joeFactory.abi.UnpackIntoInterface(out, "PendingOwnableAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryPendingOwnableNoPendingOwner represents a PendingOwnable__NoPendingOwner error raised by the JoeFactory contract.
type JoeFactoryPendingOwnableNoPendingOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PendingOwnable__NoPendingOwner()
func JoeFactoryPendingOwnableNoPendingOwnerErrorID() common.Hash {
	return common.HexToHash("0xecfad6bfa11d62ca0c0b5e6cd057ed384b108d60e6430f452462eda4b41f3679")
}

// UnpackPendingOwnableNoPendingOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PendingOwnable__NoPendingOwner()
func (joeFactory *JoeFactory) UnpackPendingOwnableNoPendingOwnerError(raw []byte) (*JoeFactoryPendingOwnableNoPendingOwner, error) {
	out := new(JoeFactoryPendingOwnableNoPendingOwner)
	if err := joeFactory.abi.UnpackIntoInterface(out, "PendingOwnableNoPendingOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryPendingOwnableNotOwner represents a PendingOwnable__NotOwner error raised by the JoeFactory contract.
type JoeFactoryPendingOwnableNotOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PendingOwnable__NotOwner()
func JoeFactoryPendingOwnableNotOwnerErrorID() common.Hash {
	return common.HexToHash("0x9f216c139c46ec4615f8f7d48426871204dd462e59ccb85402570783ada30f48")
}

// UnpackPendingOwnableNotOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PendingOwnable__NotOwner()
func (joeFactory *JoeFactory) UnpackPendingOwnableNotOwnerError(raw []byte) (*JoeFactoryPendingOwnableNotOwner, error) {
	out := new(JoeFactoryPendingOwnableNotOwner)
	if err := joeFactory.abi.UnpackIntoInterface(out, "PendingOwnableNotOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryPendingOwnableNotPendingOwner represents a PendingOwnable__NotPendingOwner error raised by the JoeFactory contract.
type JoeFactoryPendingOwnableNotPendingOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PendingOwnable__NotPendingOwner()
func JoeFactoryPendingOwnableNotPendingOwnerErrorID() common.Hash {
	return common.HexToHash("0x7304d012d604433a0dcb757adf69a0d01e9bed6a0936a4e79cd610ded2506b50")
}

// UnpackPendingOwnableNotPendingOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PendingOwnable__NotPendingOwner()
func (joeFactory *JoeFactory) UnpackPendingOwnableNotPendingOwnerError(raw []byte) (*JoeFactoryPendingOwnableNotPendingOwner, error) {
	out := new(JoeFactoryPendingOwnableNotPendingOwner)
	if err := joeFactory.abi.UnpackIntoInterface(out, "PendingOwnableNotPendingOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryPendingOwnablePendingOwnerAlreadySet represents a PendingOwnable__PendingOwnerAlreadySet error raised by the JoeFactory contract.
type JoeFactoryPendingOwnablePendingOwnerAlreadySet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PendingOwnable__PendingOwnerAlreadySet()
func JoeFactoryPendingOwnablePendingOwnerAlreadySetErrorID() common.Hash {
	return common.HexToHash("0x716b1fbf473dbdcc204315446c79a2203620729fa97231c5158a9c76ec6bb1ce")
}

// UnpackPendingOwnablePendingOwnerAlreadySetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PendingOwnable__PendingOwnerAlreadySet()
func (joeFactory *JoeFactory) UnpackPendingOwnablePendingOwnerAlreadySetError(raw []byte) (*JoeFactoryPendingOwnablePendingOwnerAlreadySet, error) {
	out := new(JoeFactoryPendingOwnablePendingOwnerAlreadySet)
	if err := joeFactory.abi.UnpackIntoInterface(out, "PendingOwnablePendingOwnerAlreadySet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactorySafeCastExceeds16Bits represents a SafeCast__Exceeds16Bits error raised by the JoeFactory contract.
type JoeFactorySafeCastExceeds16Bits struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCast__Exceeds16Bits()
func JoeFactorySafeCastExceeds16BitsErrorID() common.Hash {
	return common.HexToHash("0x64ae406d4033e0f06c19644ac75469557b82e3f7a8234d830c3f98b4a9aa5c70")
}

// UnpackSafeCastExceeds16BitsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCast__Exceeds16Bits()
func (joeFactory *JoeFactory) UnpackSafeCastExceeds16BitsError(raw []byte) (*JoeFactorySafeCastExceeds16Bits, error) {
	out := new(JoeFactorySafeCastExceeds16Bits)
	if err := joeFactory.abi.UnpackIntoInterface(out, "SafeCastExceeds16Bits", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// JoeFactoryUint128x128MathPowUnderflow represents a Uint128x128Math__PowUnderflow error raised by the JoeFactory contract.
type JoeFactoryUint128x128MathPowUnderflow struct {
	X *big.Int
	Y *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Uint128x128Math__PowUnderflow(uint256 x, int256 y)
func JoeFactoryUint128x128MathPowUnderflowErrorID() common.Hash {
	return common.HexToHash("0x3b74b31a47085961f6ba62bb07eda2a0a84b8784e0043a4818302c92815aed2a")
}

// UnpackUint128x128MathPowUnderflowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Uint128x128Math__PowUnderflow(uint256 x, int256 y)
func (joeFactory *JoeFactory) UnpackUint128x128MathPowUnderflowError(raw []byte) (*JoeFactoryUint128x128MathPowUnderflow, error) {
	out := new(JoeFactoryUint128x128MathPowUnderflow)
	if err := joeFactory.abi.UnpackIntoInterface(out, "Uint128x128MathPowUnderflow", raw); err != nil {
		return nil, err
	}
	return out, nil
}
