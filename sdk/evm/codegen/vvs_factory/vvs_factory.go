// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vvs_factory

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

// VvsFactoryMetaData contains all meta data concerning the VvsFactory contract.
var VvsFactoryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_CODE_PAIR_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "VvsFactory",
}

// VvsFactory is an auto generated Go binding around an Ethereum contract.
type VvsFactory struct {
	abi abi.ABI
}

// NewVvsFactory creates a new instance of VvsFactory.
func NewVvsFactory() *VvsFactory {
	parsed, err := VvsFactoryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &VvsFactory{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *VvsFactory) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _feeToSetter) returns()
func (vvsFactory *VvsFactory) PackConstructor(_feeToSetter common.Address) []byte {
	enc, err := vvsFactory.abi.Pack("", _feeToSetter)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackINITCODEPAIRHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5855a25a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (vvsFactory *VvsFactory) PackINITCODEPAIRHASH() []byte {
	enc, err := vvsFactory.abi.Pack("INIT_CODE_PAIR_HASH")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackINITCODEPAIRHASH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5855a25a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (vvsFactory *VvsFactory) TryPackINITCODEPAIRHASH() ([]byte, error) {
	return vvsFactory.abi.Pack("INIT_CODE_PAIR_HASH")
}

// UnpackINITCODEPAIRHASH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5855a25a.
//
// Solidity: function INIT_CODE_PAIR_HASH() view returns(bytes32)
func (vvsFactory *VvsFactory) UnpackINITCODEPAIRHASH(data []byte) ([32]byte, error) {
	out, err := vvsFactory.abi.Unpack("INIT_CODE_PAIR_HASH", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackAllPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e3dd18b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (vvsFactory *VvsFactory) PackAllPairs(arg0 *big.Int) []byte {
	enc, err := vvsFactory.abi.Pack("allPairs", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllPairs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e3dd18b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (vvsFactory *VvsFactory) TryPackAllPairs(arg0 *big.Int) ([]byte, error) {
	return vvsFactory.abi.Pack("allPairs", arg0)
}

// UnpackAllPairs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (vvsFactory *VvsFactory) UnpackAllPairs(data []byte) (common.Address, error) {
	out, err := vvsFactory.abi.Unpack("allPairs", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackAllPairsLength is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x574f2ba3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allPairsLength() view returns(uint256)
func (vvsFactory *VvsFactory) PackAllPairsLength() []byte {
	enc, err := vvsFactory.abi.Pack("allPairsLength")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllPairsLength is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x574f2ba3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allPairsLength() view returns(uint256)
func (vvsFactory *VvsFactory) TryPackAllPairsLength() ([]byte, error) {
	return vvsFactory.abi.Pack("allPairsLength")
}

// UnpackAllPairsLength is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (vvsFactory *VvsFactory) UnpackAllPairsLength(data []byte) (*big.Int, error) {
	out, err := vvsFactory.abi.Unpack("allPairsLength", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCreatePair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc9c65396.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (vvsFactory *VvsFactory) PackCreatePair(tokenA common.Address, tokenB common.Address) []byte {
	enc, err := vvsFactory.abi.Pack("createPair", tokenA, tokenB)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreatePair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc9c65396.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (vvsFactory *VvsFactory) TryPackCreatePair(tokenA common.Address, tokenB common.Address) ([]byte, error) {
	return vvsFactory.abi.Pack("createPair", tokenA, tokenB)
}

// UnpackCreatePair is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (vvsFactory *VvsFactory) UnpackCreatePair(data []byte) (common.Address, error) {
	out, err := vvsFactory.abi.Unpack("createPair", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFeeTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x017e7e58.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function feeTo() view returns(address)
func (vvsFactory *VvsFactory) PackFeeTo() []byte {
	enc, err := vvsFactory.abi.Pack("feeTo")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFeeTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x017e7e58.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function feeTo() view returns(address)
func (vvsFactory *VvsFactory) TryPackFeeTo() ([]byte, error) {
	return vvsFactory.abi.Pack("feeTo")
}

// UnpackFeeTo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (vvsFactory *VvsFactory) UnpackFeeTo(data []byte) (common.Address, error) {
	out, err := vvsFactory.abi.Unpack("feeTo", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFeeToSetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x094b7415.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function feeToSetter() view returns(address)
func (vvsFactory *VvsFactory) PackFeeToSetter() []byte {
	enc, err := vvsFactory.abi.Pack("feeToSetter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFeeToSetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x094b7415.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function feeToSetter() view returns(address)
func (vvsFactory *VvsFactory) TryPackFeeToSetter() ([]byte, error) {
	return vvsFactory.abi.Pack("feeToSetter")
}

// UnpackFeeToSetter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (vvsFactory *VvsFactory) UnpackFeeToSetter(data []byte) (common.Address, error) {
	out, err := vvsFactory.abi.Unpack("feeToSetter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe6a43905.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPair(address , address ) view returns(address)
func (vvsFactory *VvsFactory) PackGetPair(arg0 common.Address, arg1 common.Address) []byte {
	enc, err := vvsFactory.abi.Pack("getPair", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe6a43905.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPair(address , address ) view returns(address)
func (vvsFactory *VvsFactory) TryPackGetPair(arg0 common.Address, arg1 common.Address) ([]byte, error) {
	return vvsFactory.abi.Pack("getPair", arg0, arg1)
}

// UnpackGetPair is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (vvsFactory *VvsFactory) UnpackGetPair(data []byte) (common.Address, error) {
	out, err := vvsFactory.abi.Unpack("getPair", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetFeeTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf46901ed.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (vvsFactory *VvsFactory) PackSetFeeTo(feeTo common.Address) []byte {
	enc, err := vvsFactory.abi.Pack("setFeeTo", feeTo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFeeTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf46901ed.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (vvsFactory *VvsFactory) TryPackSetFeeTo(feeTo common.Address) ([]byte, error) {
	return vvsFactory.abi.Pack("setFeeTo", feeTo)
}

// PackSetFeeToSetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa2e74af6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (vvsFactory *VvsFactory) PackSetFeeToSetter(feeToSetter common.Address) []byte {
	enc, err := vvsFactory.abi.Pack("setFeeToSetter", feeToSetter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFeeToSetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa2e74af6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (vvsFactory *VvsFactory) TryPackSetFeeToSetter(feeToSetter common.Address) ([]byte, error) {
	return vvsFactory.abi.Pack("setFeeToSetter", feeToSetter)
}

// VvsFactoryPairCreated represents a PairCreated event raised by the VvsFactory contract.
type VvsFactoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const VvsFactoryPairCreatedEventName = "PairCreated"

// ContractEventName returns the user-defined event name.
func (VvsFactoryPairCreated) ContractEventName() string {
	return VvsFactoryPairCreatedEventName
}

// UnpackPairCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (vvsFactory *VvsFactory) UnpackPairCreatedEvent(log *types.Log) (*VvsFactoryPairCreated, error) {
	event := "PairCreated"
	if log.Topics[0] != vvsFactory.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VvsFactoryPairCreated)
	if len(log.Data) > 0 {
		if err := vvsFactory.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vvsFactory.abi.Events[event].Inputs {
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
