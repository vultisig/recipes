// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package odosrouterv2

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

// OdosRouterV2inputTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type OdosRouterV2inputTokenInfo struct {
	TokenAddress common.Address
	AmountIn     *big.Int
	Receiver     common.Address
}

// OdosRouterV2outputTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type OdosRouterV2outputTokenInfo struct {
	TokenAddress  common.Address
	RelativeValue *big.Int
	Receiver      common.Address
}

// OdosRouterV2permit2Info is an auto generated low-level Go binding around an user-defined struct.
type OdosRouterV2permit2Info struct {
	ContractAddress common.Address
	Nonce           *big.Int
	Deadline        *big.Int
	Signature       []byte
}

// OdosRouterV2swapTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type OdosRouterV2swapTokenInfo struct {
	InputToken     common.Address
	InputAmount    *big.Int
	InputReceiver  common.Address
	OutputToken    common.Address
	OutputQuote    *big.Int
	OutputMin      *big.Int
	OutputReceiver common.Address
}

// Odosrouterv2MetaData contains all meta data concerning the Odosrouterv2 contract.
var Odosrouterv2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"slippage\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokensIn\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokensOut\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"SwapMulti\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_DENOM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFERRAL_WITH_FEE_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addressList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"referralLookup\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"referralFee\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_referralCode\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_referralFee\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"registerReferralCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_swapMultiFee\",\"type\":\"uint256\"}],\"name\":\"setSwapMultiFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"outputQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputReceiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.swapTokenInfo\",\"name\":\"tokenInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"pathDefinition\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapCompact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.inputTokenInfo[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relativeValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.outputTokenInfo[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"valueOutMin\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pathDefinition\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"swapMulti\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapMultiCompact\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapMultiFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOdosRouterV2.permit2Info\",\"name\":\"permit2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.inputTokenInfo[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relativeValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.outputTokenInfo[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"valueOutMin\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pathDefinition\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"swapMultiPermit2\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOdosRouterV2.permit2Info\",\"name\":\"permit2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"inputReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"outputQuote\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"outputReceiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.swapTokenInfo\",\"name\":\"tokenInfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"pathDefinition\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referralCode\",\"type\":\"uint32\"}],\"name\":\"swapPermit2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.inputTokenInfo[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"relativeValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structOdosRouterV2.outputTokenInfo[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"valueOutMin\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pathDefinition\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"swapRouterFunds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"transferRouterFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"writeAddressList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "Odosrouterv2",
}

// Odosrouterv2 is an auto generated Go binding around an Ethereum contract.
type Odosrouterv2 struct {
	abi abi.ABI
}

// NewOdosrouterv2 creates a new instance of Odosrouterv2.
func NewOdosrouterv2() *Odosrouterv2 {
	parsed, err := Odosrouterv2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Odosrouterv2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Odosrouterv2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackFEEDENOM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4886c675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function FEE_DENOM() view returns(uint256)
func (odosrouterv2 *Odosrouterv2) PackFEEDENOM() []byte {
	enc, err := odosrouterv2.abi.Pack("FEE_DENOM")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFEEDENOM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4886c675.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function FEE_DENOM() view returns(uint256)
func (odosrouterv2 *Odosrouterv2) TryPackFEEDENOM() ([]byte, error) {
	return odosrouterv2.abi.Pack("FEE_DENOM")
}

// UnpackFEEDENOM is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4886c675.
//
// Solidity: function FEE_DENOM() view returns(uint256)
func (odosrouterv2 *Odosrouterv2) UnpackFEEDENOM(data []byte) (*big.Int, error) {
	out, err := odosrouterv2.abi.Unpack("FEE_DENOM", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackREFERRALWITHFEETHRESHOLD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6c082c13.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function REFERRAL_WITH_FEE_THRESHOLD() view returns(uint256)
func (odosrouterv2 *Odosrouterv2) PackREFERRALWITHFEETHRESHOLD() []byte {
	enc, err := odosrouterv2.abi.Pack("REFERRAL_WITH_FEE_THRESHOLD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackREFERRALWITHFEETHRESHOLD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6c082c13.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function REFERRAL_WITH_FEE_THRESHOLD() view returns(uint256)
func (odosrouterv2 *Odosrouterv2) TryPackREFERRALWITHFEETHRESHOLD() ([]byte, error) {
	return odosrouterv2.abi.Pack("REFERRAL_WITH_FEE_THRESHOLD")
}

// UnpackREFERRALWITHFEETHRESHOLD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6c082c13.
//
// Solidity: function REFERRAL_WITH_FEE_THRESHOLD() view returns(uint256)
func (odosrouterv2 *Odosrouterv2) UnpackREFERRALWITHFEETHRESHOLD(data []byte) (*big.Int, error) {
	out, err := odosrouterv2.abi.Unpack("REFERRAL_WITH_FEE_THRESHOLD", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackAddressList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb810fb43.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addressList(uint256 ) view returns(address)
func (odosrouterv2 *Odosrouterv2) PackAddressList(arg0 *big.Int) []byte {
	enc, err := odosrouterv2.abi.Pack("addressList", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddressList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb810fb43.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addressList(uint256 ) view returns(address)
func (odosrouterv2 *Odosrouterv2) TryPackAddressList(arg0 *big.Int) ([]byte, error) {
	return odosrouterv2.abi.Pack("addressList", arg0)
}

// UnpackAddressList is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb810fb43.
//
// Solidity: function addressList(uint256 ) view returns(address)
func (odosrouterv2 *Odosrouterv2) UnpackAddressList(data []byte) (common.Address, error) {
	out, err := odosrouterv2.abi.Unpack("addressList", data)
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
func (odosrouterv2 *Odosrouterv2) PackOwner() []byte {
	enc, err := odosrouterv2.abi.Pack("owner")
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
func (odosrouterv2 *Odosrouterv2) TryPackOwner() ([]byte, error) {
	return odosrouterv2.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (odosrouterv2 *Odosrouterv2) UnpackOwner(data []byte) (common.Address, error) {
	out, err := odosrouterv2.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackReferralLookup is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf827065e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function referralLookup(uint32 ) view returns(uint64 referralFee, address beneficiary, bool registered)
func (odosrouterv2 *Odosrouterv2) PackReferralLookup(arg0 uint32) []byte {
	enc, err := odosrouterv2.abi.Pack("referralLookup", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReferralLookup is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf827065e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function referralLookup(uint32 ) view returns(uint64 referralFee, address beneficiary, bool registered)
func (odosrouterv2 *Odosrouterv2) TryPackReferralLookup(arg0 uint32) ([]byte, error) {
	return odosrouterv2.abi.Pack("referralLookup", arg0)
}

// ReferralLookupOutput serves as a container for the return parameters of contract
// method ReferralLookup.
type ReferralLookupOutput struct {
	ReferralFee uint64
	Beneficiary common.Address
	Registered  bool
}

// UnpackReferralLookup is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf827065e.
//
// Solidity: function referralLookup(uint32 ) view returns(uint64 referralFee, address beneficiary, bool registered)
func (odosrouterv2 *Odosrouterv2) UnpackReferralLookup(data []byte) (ReferralLookupOutput, error) {
	out, err := odosrouterv2.abi.Unpack("referralLookup", data)
	outstruct := new(ReferralLookupOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ReferralFee = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Beneficiary = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Registered = *abi.ConvertType(out[2], new(bool)).(*bool)
	return *outstruct, nil
}

// PackRegisterReferralCode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe10895f9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerReferralCode(uint32 _referralCode, uint64 _referralFee, address _beneficiary) returns()
func (odosrouterv2 *Odosrouterv2) PackRegisterReferralCode(referralCode uint32, referralFee uint64, beneficiary common.Address) []byte {
	enc, err := odosrouterv2.abi.Pack("registerReferralCode", referralCode, referralFee, beneficiary)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterReferralCode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe10895f9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerReferralCode(uint32 _referralCode, uint64 _referralFee, address _beneficiary) returns()
func (odosrouterv2 *Odosrouterv2) TryPackRegisterReferralCode(referralCode uint32, referralFee uint64, beneficiary common.Address) ([]byte, error) {
	return odosrouterv2.abi.Pack("registerReferralCode", referralCode, referralFee, beneficiary)
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (odosrouterv2 *Odosrouterv2) PackRenounceOwnership() []byte {
	enc, err := odosrouterv2.abi.Pack("renounceOwnership")
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
func (odosrouterv2 *Odosrouterv2) TryPackRenounceOwnership() ([]byte, error) {
	return odosrouterv2.abi.Pack("renounceOwnership")
}

// PackSetSwapMultiFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9286b93d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSwapMultiFee(uint256 _swapMultiFee) returns()
func (odosrouterv2 *Odosrouterv2) PackSetSwapMultiFee(swapMultiFee *big.Int) []byte {
	enc, err := odosrouterv2.abi.Pack("setSwapMultiFee", swapMultiFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSwapMultiFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9286b93d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSwapMultiFee(uint256 _swapMultiFee) returns()
func (odosrouterv2 *Odosrouterv2) TryPackSetSwapMultiFee(swapMultiFee *big.Int) ([]byte, error) {
	return odosrouterv2.abi.Pack("setSwapMultiFee", swapMultiFee)
}

// PackSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3b635ce4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swap((address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256 amountOut)
func (odosrouterv2 *Odosrouterv2) PackSwap(tokenInfo OdosRouterV2swapTokenInfo, pathDefinition []byte, executor common.Address, referralCode uint32) []byte {
	enc, err := odosrouterv2.abi.Pack("swap", tokenInfo, pathDefinition, executor, referralCode)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3b635ce4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swap((address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256 amountOut)
func (odosrouterv2 *Odosrouterv2) TryPackSwap(tokenInfo OdosRouterV2swapTokenInfo, pathDefinition []byte, executor common.Address, referralCode uint32) ([]byte, error) {
	return odosrouterv2.abi.Pack("swap", tokenInfo, pathDefinition, executor, referralCode)
}

// UnpackSwap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3b635ce4.
//
// Solidity: function swap((address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256 amountOut)
func (odosrouterv2 *Odosrouterv2) UnpackSwap(data []byte) (*big.Int, error) {
	out, err := odosrouterv2.abi.Unpack("swap", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSwapCompact is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83bd37f9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapCompact() payable returns(uint256)
func (odosrouterv2 *Odosrouterv2) PackSwapCompact() []byte {
	enc, err := odosrouterv2.abi.Pack("swapCompact")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapCompact is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83bd37f9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapCompact() payable returns(uint256)
func (odosrouterv2 *Odosrouterv2) TryPackSwapCompact() ([]byte, error) {
	return odosrouterv2.abi.Pack("swapCompact")
}

// UnpackSwapCompact is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x83bd37f9.
//
// Solidity: function swapCompact() payable returns(uint256)
func (odosrouterv2 *Odosrouterv2) UnpackSwapCompact(data []byte) (*big.Int, error) {
	out, err := odosrouterv2.abi.Unpack("swapCompact", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSwapMulti is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7bf2d6d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapMulti((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) PackSwapMulti(inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address, referralCode uint32) []byte {
	enc, err := odosrouterv2.abi.Pack("swapMulti", inputs, outputs, valueOutMin, pathDefinition, executor, referralCode)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapMulti is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7bf2d6d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapMulti((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) TryPackSwapMulti(inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address, referralCode uint32) ([]byte, error) {
	return odosrouterv2.abi.Pack("swapMulti", inputs, outputs, valueOutMin, pathDefinition, executor, referralCode)
}

// UnpackSwapMulti is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7bf2d6d4.
//
// Solidity: function swapMulti((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) UnpackSwapMulti(data []byte) ([]*big.Int, error) {
	out, err := odosrouterv2.abi.Unpack("swapMulti", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackSwapMultiCompact is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84a7f3dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapMultiCompact() payable returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) PackSwapMultiCompact() []byte {
	enc, err := odosrouterv2.abi.Pack("swapMultiCompact")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapMultiCompact is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84a7f3dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapMultiCompact() payable returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) TryPackSwapMultiCompact() ([]byte, error) {
	return odosrouterv2.abi.Pack("swapMultiCompact")
}

// UnpackSwapMultiCompact is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84a7f3dd.
//
// Solidity: function swapMultiCompact() payable returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) UnpackSwapMultiCompact(data []byte) ([]*big.Int, error) {
	out, err := odosrouterv2.abi.Unpack("swapMultiCompact", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackSwapMultiFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe7d3fc60.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapMultiFee() view returns(uint256)
func (odosrouterv2 *Odosrouterv2) PackSwapMultiFee() []byte {
	enc, err := odosrouterv2.abi.Pack("swapMultiFee")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapMultiFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe7d3fc60.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapMultiFee() view returns(uint256)
func (odosrouterv2 *Odosrouterv2) TryPackSwapMultiFee() ([]byte, error) {
	return odosrouterv2.abi.Pack("swapMultiFee")
}

// UnpackSwapMultiFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe7d3fc60.
//
// Solidity: function swapMultiFee() view returns(uint256)
func (odosrouterv2 *Odosrouterv2) UnpackSwapMultiFee(data []byte) (*big.Int, error) {
	out, err := odosrouterv2.abi.Unpack("swapMultiFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSwapMultiPermit2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x080c25b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapMultiPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) PackSwapMultiPermit2(permit2 OdosRouterV2permit2Info, inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address, referralCode uint32) []byte {
	enc, err := odosrouterv2.abi.Pack("swapMultiPermit2", permit2, inputs, outputs, valueOutMin, pathDefinition, executor, referralCode)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapMultiPermit2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x080c25b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapMultiPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) TryPackSwapMultiPermit2(permit2 OdosRouterV2permit2Info, inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address, referralCode uint32) ([]byte, error) {
	return odosrouterv2.abi.Pack("swapMultiPermit2", permit2, inputs, outputs, valueOutMin, pathDefinition, executor, referralCode)
}

// UnpackSwapMultiPermit2 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x080c25b3.
//
// Solidity: function swapMultiPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor, uint32 referralCode) payable returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) UnpackSwapMultiPermit2(data []byte) ([]*big.Int, error) {
	out, err := odosrouterv2.abi.Unpack("swapMultiPermit2", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackSwapPermit2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87b621b5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) returns(uint256 amountOut)
func (odosrouterv2 *Odosrouterv2) PackSwapPermit2(permit2 OdosRouterV2permit2Info, tokenInfo OdosRouterV2swapTokenInfo, pathDefinition []byte, executor common.Address, referralCode uint32) []byte {
	enc, err := odosrouterv2.abi.Pack("swapPermit2", permit2, tokenInfo, pathDefinition, executor, referralCode)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapPermit2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87b621b5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) returns(uint256 amountOut)
func (odosrouterv2 *Odosrouterv2) TryPackSwapPermit2(permit2 OdosRouterV2permit2Info, tokenInfo OdosRouterV2swapTokenInfo, pathDefinition []byte, executor common.Address, referralCode uint32) ([]byte, error) {
	return odosrouterv2.abi.Pack("swapPermit2", permit2, tokenInfo, pathDefinition, executor, referralCode)
}

// UnpackSwapPermit2 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87b621b5.
//
// Solidity: function swapPermit2((address,uint256,uint256,bytes) permit2, (address,uint256,address,address,uint256,uint256,address) tokenInfo, bytes pathDefinition, address executor, uint32 referralCode) returns(uint256 amountOut)
func (odosrouterv2 *Odosrouterv2) UnpackSwapPermit2(data []byte) (*big.Int, error) {
	out, err := odosrouterv2.abi.Unpack("swapPermit2", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSwapRouterFunds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x28be42f4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapRouterFunds((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor) returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) PackSwapRouterFunds(inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address) []byte {
	enc, err := odosrouterv2.abi.Pack("swapRouterFunds", inputs, outputs, valueOutMin, pathDefinition, executor)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapRouterFunds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x28be42f4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapRouterFunds((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor) returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) TryPackSwapRouterFunds(inputs []OdosRouterV2inputTokenInfo, outputs []OdosRouterV2outputTokenInfo, valueOutMin *big.Int, pathDefinition []byte, executor common.Address) ([]byte, error) {
	return odosrouterv2.abi.Pack("swapRouterFunds", inputs, outputs, valueOutMin, pathDefinition, executor)
}

// UnpackSwapRouterFunds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x28be42f4.
//
// Solidity: function swapRouterFunds((address,uint256,address)[] inputs, (address,uint256,address)[] outputs, uint256 valueOutMin, bytes pathDefinition, address executor) returns(uint256[] amountsOut)
func (odosrouterv2 *Odosrouterv2) UnpackSwapRouterFunds(data []byte) ([]*big.Int, error) {
	out, err := odosrouterv2.abi.Unpack("swapRouterFunds", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (odosrouterv2 *Odosrouterv2) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := odosrouterv2.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (odosrouterv2 *Odosrouterv2) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return odosrouterv2.abi.Pack("transferOwnership", newOwner)
}

// PackTransferRouterFunds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x174da621.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferRouterFunds(address[] tokens, uint256[] amounts, address dest) returns()
func (odosrouterv2 *Odosrouterv2) PackTransferRouterFunds(tokens []common.Address, amounts []*big.Int, dest common.Address) []byte {
	enc, err := odosrouterv2.abi.Pack("transferRouterFunds", tokens, amounts, dest)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferRouterFunds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x174da621.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferRouterFunds(address[] tokens, uint256[] amounts, address dest) returns()
func (odosrouterv2 *Odosrouterv2) TryPackTransferRouterFunds(tokens []common.Address, amounts []*big.Int, dest common.Address) ([]byte, error) {
	return odosrouterv2.abi.Pack("transferRouterFunds", tokens, amounts, dest)
}

// PackWriteAddressList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3596f9a2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function writeAddressList(address[] addresses) returns()
func (odosrouterv2 *Odosrouterv2) PackWriteAddressList(addresses []common.Address) []byte {
	enc, err := odosrouterv2.abi.Pack("writeAddressList", addresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWriteAddressList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3596f9a2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function writeAddressList(address[] addresses) returns()
func (odosrouterv2 *Odosrouterv2) TryPackWriteAddressList(addresses []common.Address) ([]byte, error) {
	return odosrouterv2.abi.Pack("writeAddressList", addresses)
}

// Odosrouterv2OwnershipTransferred represents a OwnershipTransferred event raised by the Odosrouterv2 contract.
type Odosrouterv2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const Odosrouterv2OwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (Odosrouterv2OwnershipTransferred) ContractEventName() string {
	return Odosrouterv2OwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (odosrouterv2 *Odosrouterv2) UnpackOwnershipTransferredEvent(log *types.Log) (*Odosrouterv2OwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != odosrouterv2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Odosrouterv2OwnershipTransferred)
	if len(log.Data) > 0 {
		if err := odosrouterv2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range odosrouterv2.abi.Events[event].Inputs {
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

// Odosrouterv2Swap represents a Swap event raised by the Odosrouterv2 contract.
type Odosrouterv2Swap struct {
	Sender       common.Address
	InputAmount  *big.Int
	InputToken   common.Address
	AmountOut    *big.Int
	OutputToken  common.Address
	Slippage     *big.Int
	ReferralCode uint32
	Raw          *types.Log // Blockchain specific contextual infos
}

const Odosrouterv2SwapEventName = "Swap"

// ContractEventName returns the user-defined event name.
func (Odosrouterv2Swap) ContractEventName() string {
	return Odosrouterv2SwapEventName
}

// UnpackSwapEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Swap(address sender, uint256 inputAmount, address inputToken, uint256 amountOut, address outputToken, int256 slippage, uint32 referralCode)
func (odosrouterv2 *Odosrouterv2) UnpackSwapEvent(log *types.Log) (*Odosrouterv2Swap, error) {
	event := "Swap"
	if log.Topics[0] != odosrouterv2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Odosrouterv2Swap)
	if len(log.Data) > 0 {
		if err := odosrouterv2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range odosrouterv2.abi.Events[event].Inputs {
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

// Odosrouterv2SwapMulti represents a SwapMulti event raised by the Odosrouterv2 contract.
type Odosrouterv2SwapMulti struct {
	Sender       common.Address
	AmountsIn    []*big.Int
	TokensIn     []common.Address
	AmountsOut   []*big.Int
	TokensOut    []common.Address
	ReferralCode uint32
	Raw          *types.Log // Blockchain specific contextual infos
}

const Odosrouterv2SwapMultiEventName = "SwapMulti"

// ContractEventName returns the user-defined event name.
func (Odosrouterv2SwapMulti) ContractEventName() string {
	return Odosrouterv2SwapMultiEventName
}

// UnpackSwapMultiEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SwapMulti(address sender, uint256[] amountsIn, address[] tokensIn, uint256[] amountsOut, address[] tokensOut, uint32 referralCode)
func (odosrouterv2 *Odosrouterv2) UnpackSwapMultiEvent(log *types.Log) (*Odosrouterv2SwapMulti, error) {
	event := "SwapMulti"
	if log.Topics[0] != odosrouterv2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Odosrouterv2SwapMulti)
	if len(log.Data) > 0 {
		if err := odosrouterv2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range odosrouterv2.abi.Events[event].Inputs {
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
