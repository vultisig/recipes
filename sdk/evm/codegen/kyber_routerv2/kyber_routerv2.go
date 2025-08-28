// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kyber_routerv2

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

// MetaAggregationRouterV2SwapDescriptionV2 is an auto generated low-level Go binding around an user-defined struct.
type MetaAggregationRouterV2SwapDescriptionV2 struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceivers    []common.Address
	SrcAmounts      []*big.Int
	FeeReceivers    []common.Address
	FeeAmounts      []*big.Int
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
	Permit          []byte
}

// MetaAggregationRouterV2SwapExecutionParams is an auto generated low-level Go binding around an user-defined struct.
type MetaAggregationRouterV2SwapExecutionParams struct {
	CallTarget    common.Address
	ApproveTarget common.Address
	TargetData    []byte
	Desc          MetaAggregationRouterV2SwapDescriptionV2
	ClientData    []byte
}

// KyberRouterv2MetaData contains all meta data concerning the KyberRouterv2 contract.
var KyberRouterv2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"ClientData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"output\",\"type\":\"address\"}],\"name\":\"Exchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBps\",\"type\":\"bool\"}],\"name\":\"Fee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"callTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTarget\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"feeReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapDescriptionV2\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapExecutionParams\",\"name\":\"execution\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"callTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"approveTarget\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"feeReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapDescriptionV2\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapExecutionParams\",\"name\":\"execution\",\"type\":\"tuple\"}],\"name\":\"swapGeneric\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"srcAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"feeReceivers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"feeAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structMetaAggregationRouterV2.SwapDescriptionV2\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"executorData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientData\",\"type\":\"bytes\"}],\"name\":\"swapSimpleMode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addr\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"value\",\"type\":\"bool[]\"}],\"name\":\"updateWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "KyberRouterv2",
}

// KyberRouterv2 is an auto generated Go binding around an Ethereum contract.
type KyberRouterv2 struct {
	abi abi.ABI
}

// NewKyberRouterv2 creates a new instance of KyberRouterv2.
func NewKyberRouterv2() *KyberRouterv2 {
	parsed, err := KyberRouterv2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &KyberRouterv2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *KyberRouterv2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _WETH) returns()
func (kyberRouterv2 *KyberRouterv2) PackConstructor(_WETH common.Address) []byte {
	enc, err := kyberRouterv2.abi.Pack("", _WETH)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad5c4648.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function WETH() view returns(address)
func (kyberRouterv2 *KyberRouterv2) PackWETH() []byte {
	enc, err := kyberRouterv2.abi.Pack("WETH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad5c4648.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function WETH() view returns(address)
func (kyberRouterv2 *KyberRouterv2) TryPackWETH() ([]byte, error) {
	return kyberRouterv2.abi.Pack("WETH")
}

// UnpackWETH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (kyberRouterv2 *KyberRouterv2) UnpackWETH(data []byte) (common.Address, error) {
	out, err := kyberRouterv2.abi.Unpack("WETH", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsWhitelist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc683630d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (kyberRouterv2 *KyberRouterv2) PackIsWhitelist(arg0 common.Address) []byte {
	enc, err := kyberRouterv2.abi.Pack("isWhitelist", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsWhitelist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc683630d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (kyberRouterv2 *KyberRouterv2) TryPackIsWhitelist(arg0 common.Address) ([]byte, error) {
	return kyberRouterv2.abi.Pack("isWhitelist", arg0)
}

// UnpackIsWhitelist is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc683630d.
//
// Solidity: function isWhitelist(address ) view returns(bool)
func (kyberRouterv2 *KyberRouterv2) UnpackIsWhitelist(data []byte) (bool, error) {
	out, err := kyberRouterv2.abi.Unpack("isWhitelist", data)
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
func (kyberRouterv2 *KyberRouterv2) PackOwner() []byte {
	enc, err := kyberRouterv2.abi.Pack("owner")
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
func (kyberRouterv2 *KyberRouterv2) TryPackOwner() ([]byte, error) {
	return kyberRouterv2.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (kyberRouterv2 *KyberRouterv2) UnpackOwner(data []byte) (common.Address, error) {
	out, err := kyberRouterv2.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (kyberRouterv2 *KyberRouterv2) PackRenounceOwnership() []byte {
	enc, err := kyberRouterv2.abi.Pack("renounceOwnership")
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
func (kyberRouterv2 *KyberRouterv2) TryPackRenounceOwnership() ([]byte, error) {
	return kyberRouterv2.abi.Pack("renounceOwnership")
}

// PackRescueFunds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78e3214f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (kyberRouterv2 *KyberRouterv2) PackRescueFunds(token common.Address, amount *big.Int) []byte {
	enc, err := kyberRouterv2.abi.Pack("rescueFunds", token, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRescueFunds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78e3214f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (kyberRouterv2 *KyberRouterv2) TryPackRescueFunds(token common.Address, amount *big.Int) ([]byte, error) {
	return kyberRouterv2.abi.Pack("rescueFunds", token, amount)
}

// PackSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe21fd0e9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swap((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (kyberRouterv2 *KyberRouterv2) PackSwap(execution MetaAggregationRouterV2SwapExecutionParams) []byte {
	enc, err := kyberRouterv2.abi.Pack("swap", execution)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe21fd0e9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swap((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (kyberRouterv2 *KyberRouterv2) TryPackSwap(execution MetaAggregationRouterV2SwapExecutionParams) ([]byte, error) {
	return kyberRouterv2.abi.Pack("swap", execution)
}

// SwapOutput serves as a container for the return parameters of contract
// method Swap.
type SwapOutput struct {
	ReturnAmount *big.Int
	GasUsed      *big.Int
}

// UnpackSwap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe21fd0e9.
//
// Solidity: function swap((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (kyberRouterv2 *KyberRouterv2) UnpackSwap(data []byte) (SwapOutput, error) {
	out, err := kyberRouterv2.abi.Unpack("swap", data)
	outstruct := new(SwapOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ReturnAmount = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.GasUsed = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackSwapGeneric is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x59e50fed.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapGeneric((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (kyberRouterv2 *KyberRouterv2) PackSwapGeneric(execution MetaAggregationRouterV2SwapExecutionParams) []byte {
	enc, err := kyberRouterv2.abi.Pack("swapGeneric", execution)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapGeneric is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x59e50fed.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapGeneric((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (kyberRouterv2 *KyberRouterv2) TryPackSwapGeneric(execution MetaAggregationRouterV2SwapExecutionParams) ([]byte, error) {
	return kyberRouterv2.abi.Pack("swapGeneric", execution)
}

// SwapGenericOutput serves as a container for the return parameters of contract
// method SwapGeneric.
type SwapGenericOutput struct {
	ReturnAmount *big.Int
	GasUsed      *big.Int
}

// UnpackSwapGeneric is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x59e50fed.
//
// Solidity: function swapGeneric((address,address,bytes,(address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes),bytes) execution) payable returns(uint256 returnAmount, uint256 gasUsed)
func (kyberRouterv2 *KyberRouterv2) UnpackSwapGeneric(data []byte) (SwapGenericOutput, error) {
	out, err := kyberRouterv2.abi.Unpack("swapGeneric", data)
	outstruct := new(SwapGenericOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ReturnAmount = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.GasUsed = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackSwapSimpleMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8af033fb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (kyberRouterv2 *KyberRouterv2) PackSwapSimpleMode(caller common.Address, desc MetaAggregationRouterV2SwapDescriptionV2, executorData []byte, clientData []byte) []byte {
	enc, err := kyberRouterv2.abi.Pack("swapSimpleMode", caller, desc, executorData, clientData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapSimpleMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8af033fb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (kyberRouterv2 *KyberRouterv2) TryPackSwapSimpleMode(caller common.Address, desc MetaAggregationRouterV2SwapDescriptionV2, executorData []byte, clientData []byte) ([]byte, error) {
	return kyberRouterv2.abi.Pack("swapSimpleMode", caller, desc, executorData, clientData)
}

// SwapSimpleModeOutput serves as a container for the return parameters of contract
// method SwapSimpleMode.
type SwapSimpleModeOutput struct {
	ReturnAmount *big.Int
	GasUsed      *big.Int
}

// UnpackSwapSimpleMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8af033fb.
//
// Solidity: function swapSimpleMode(address caller, (address,address,address[],uint256[],address[],uint256[],address,uint256,uint256,uint256,bytes) desc, bytes executorData, bytes clientData) returns(uint256 returnAmount, uint256 gasUsed)
func (kyberRouterv2 *KyberRouterv2) UnpackSwapSimpleMode(data []byte) (SwapSimpleModeOutput, error) {
	out, err := kyberRouterv2.abi.Unpack("swapSimpleMode", data)
	outstruct := new(SwapSimpleModeOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ReturnAmount = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.GasUsed = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (kyberRouterv2 *KyberRouterv2) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := kyberRouterv2.abi.Pack("transferOwnership", newOwner)
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
func (kyberRouterv2 *KyberRouterv2) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return kyberRouterv2.abi.Pack("transferOwnership", newOwner)
}

// PackUpdateWhitelist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x33320de3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateWhitelist(address[] addr, bool[] value) returns()
func (kyberRouterv2 *KyberRouterv2) PackUpdateWhitelist(addr []common.Address, value []bool) []byte {
	enc, err := kyberRouterv2.abi.Pack("updateWhitelist", addr, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateWhitelist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x33320de3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateWhitelist(address[] addr, bool[] value) returns()
func (kyberRouterv2 *KyberRouterv2) TryPackUpdateWhitelist(addr []common.Address, value []bool) ([]byte, error) {
	return kyberRouterv2.abi.Pack("updateWhitelist", addr, value)
}

// KyberRouterv2ClientData represents a ClientData event raised by the KyberRouterv2 contract.
type KyberRouterv2ClientData struct {
	ClientData []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const KyberRouterv2ClientDataEventName = "ClientData"

// ContractEventName returns the user-defined event name.
func (KyberRouterv2ClientData) ContractEventName() string {
	return KyberRouterv2ClientDataEventName
}

// UnpackClientDataEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ClientData(bytes clientData)
func (kyberRouterv2 *KyberRouterv2) UnpackClientDataEvent(log *types.Log) (*KyberRouterv2ClientData, error) {
	event := "ClientData"
	if log.Topics[0] != kyberRouterv2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(KyberRouterv2ClientData)
	if len(log.Data) > 0 {
		if err := kyberRouterv2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range kyberRouterv2.abi.Events[event].Inputs {
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

// KyberRouterv2Error represents a Error event raised by the KyberRouterv2 contract.
type KyberRouterv2Error struct {
	Reason string
	Raw    *types.Log // Blockchain specific contextual infos
}

const KyberRouterv2ErrorEventName = "Error"

// ContractEventName returns the user-defined event name.
func (KyberRouterv2Error) ContractEventName() string {
	return KyberRouterv2ErrorEventName
}

// UnpackErrorEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Error(string reason)
func (kyberRouterv2 *KyberRouterv2) UnpackErrorEvent(log *types.Log) (*KyberRouterv2Error, error) {
	event := "Error"
	if log.Topics[0] != kyberRouterv2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(KyberRouterv2Error)
	if len(log.Data) > 0 {
		if err := kyberRouterv2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range kyberRouterv2.abi.Events[event].Inputs {
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

// KyberRouterv2Exchange represents a Exchange event raised by the KyberRouterv2 contract.
type KyberRouterv2Exchange struct {
	Pair      common.Address
	AmountOut *big.Int
	Output    common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const KyberRouterv2ExchangeEventName = "Exchange"

// ContractEventName returns the user-defined event name.
func (KyberRouterv2Exchange) ContractEventName() string {
	return KyberRouterv2ExchangeEventName
}

// UnpackExchangeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Exchange(address pair, uint256 amountOut, address output)
func (kyberRouterv2 *KyberRouterv2) UnpackExchangeEvent(log *types.Log) (*KyberRouterv2Exchange, error) {
	event := "Exchange"
	if log.Topics[0] != kyberRouterv2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(KyberRouterv2Exchange)
	if len(log.Data) > 0 {
		if err := kyberRouterv2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range kyberRouterv2.abi.Events[event].Inputs {
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

// KyberRouterv2Fee represents a Fee event raised by the KyberRouterv2 contract.
type KyberRouterv2Fee struct {
	Token       common.Address
	TotalAmount *big.Int
	TotalFee    *big.Int
	Recipients  []common.Address
	Amounts     []*big.Int
	IsBps       bool
	Raw         *types.Log // Blockchain specific contextual infos
}

const KyberRouterv2FeeEventName = "Fee"

// ContractEventName returns the user-defined event name.
func (KyberRouterv2Fee) ContractEventName() string {
	return KyberRouterv2FeeEventName
}

// UnpackFeeEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Fee(address token, uint256 totalAmount, uint256 totalFee, address[] recipients, uint256[] amounts, bool isBps)
func (kyberRouterv2 *KyberRouterv2) UnpackFeeEvent(log *types.Log) (*KyberRouterv2Fee, error) {
	event := "Fee"
	if log.Topics[0] != kyberRouterv2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(KyberRouterv2Fee)
	if len(log.Data) > 0 {
		if err := kyberRouterv2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range kyberRouterv2.abi.Events[event].Inputs {
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

// KyberRouterv2OwnershipTransferred represents a OwnershipTransferred event raised by the KyberRouterv2 contract.
type KyberRouterv2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const KyberRouterv2OwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (KyberRouterv2OwnershipTransferred) ContractEventName() string {
	return KyberRouterv2OwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (kyberRouterv2 *KyberRouterv2) UnpackOwnershipTransferredEvent(log *types.Log) (*KyberRouterv2OwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != kyberRouterv2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(KyberRouterv2OwnershipTransferred)
	if len(log.Data) > 0 {
		if err := kyberRouterv2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range kyberRouterv2.abi.Events[event].Inputs {
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

// KyberRouterv2Swapped represents a Swapped event raised by the KyberRouterv2 contract.
type KyberRouterv2Swapped struct {
	Sender       common.Address
	SrcToken     common.Address
	DstToken     common.Address
	DstReceiver  common.Address
	SpentAmount  *big.Int
	ReturnAmount *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const KyberRouterv2SwappedEventName = "Swapped"

// ContractEventName returns the user-defined event name.
func (KyberRouterv2Swapped) ContractEventName() string {
	return KyberRouterv2SwappedEventName
}

// UnpackSwappedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (kyberRouterv2 *KyberRouterv2) UnpackSwappedEvent(log *types.Log) (*KyberRouterv2Swapped, error) {
	event := "Swapped"
	if log.Topics[0] != kyberRouterv2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(KyberRouterv2Swapped)
	if len(log.Data) > 0 {
		if err := kyberRouterv2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range kyberRouterv2.abi.Events[event].Inputs {
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
