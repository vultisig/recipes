// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mayachain_router

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

// THORChainRouterCoin is an auto generated low-level Go binding around an user-defined struct.
type THORChainRouterCoin struct {
	Asset  common.Address
	Amount *big.Int
}

// MayachainRouterMetaData contains all meta data concerning the MayachainRouter contract.
var MayachainRouterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rune\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"TransferAllowance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"TransferOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"finalAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"TransferOutAndCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTHORChain_Router.Coin[]\",\"name\":\"coins\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"VaultTransfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RUNE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"depositWithExpiry\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"asgard\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTHORChain_Router.Coin[]\",\"name\":\"coins\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"returnVaultAssets\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"finalToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferOutAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"vaultAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "MayachainRouter",
}

// MayachainRouter is an auto generated Go binding around an Ethereum contract.
type MayachainRouter struct {
	abi abi.ABI
}

// NewMayachainRouter creates a new instance of MayachainRouter.
func NewMayachainRouter() *MayachainRouter {
	parsed, err := MayachainRouterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MayachainRouter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MayachainRouter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address rune) returns()
func (mayachainRouter *MayachainRouter) PackConstructor(rune common.Address) []byte {
	enc, err := mayachainRouter.abi.Pack("", rune)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRUNE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x93e4eaa9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function RUNE() view returns(address)
func (mayachainRouter *MayachainRouter) PackRUNE() []byte {
	enc, err := mayachainRouter.abi.Pack("RUNE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRUNE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x93e4eaa9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function RUNE() view returns(address)
func (mayachainRouter *MayachainRouter) TryPackRUNE() ([]byte, error) {
	return mayachainRouter.abi.Pack("RUNE")
}

// UnpackRUNE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x93e4eaa9.
//
// Solidity: function RUNE() view returns(address)
func (mayachainRouter *MayachainRouter) UnpackRUNE(data []byte) (common.Address, error) {
	out, err := mayachainRouter.abi.Unpack("RUNE", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1fece7b4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deposit(address vault, address asset, uint256 amount, string memo) payable returns()
func (mayachainRouter *MayachainRouter) PackDeposit(vault common.Address, asset common.Address, amount *big.Int, memo string) []byte {
	enc, err := mayachainRouter.abi.Pack("deposit", vault, asset, amount, memo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1fece7b4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deposit(address vault, address asset, uint256 amount, string memo) payable returns()
func (mayachainRouter *MayachainRouter) TryPackDeposit(vault common.Address, asset common.Address, amount *big.Int, memo string) ([]byte, error) {
	return mayachainRouter.abi.Pack("deposit", vault, asset, amount, memo)
}

// PackDepositWithExpiry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44bc937b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function depositWithExpiry(address vault, address asset, uint256 amount, string memo, uint256 expiration) payable returns()
func (mayachainRouter *MayachainRouter) PackDepositWithExpiry(vault common.Address, asset common.Address, amount *big.Int, memo string, expiration *big.Int) []byte {
	enc, err := mayachainRouter.abi.Pack("depositWithExpiry", vault, asset, amount, memo, expiration)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDepositWithExpiry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44bc937b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function depositWithExpiry(address vault, address asset, uint256 amount, string memo, uint256 expiration) payable returns()
func (mayachainRouter *MayachainRouter) TryPackDepositWithExpiry(vault common.Address, asset common.Address, amount *big.Int, memo string, expiration *big.Int) ([]byte, error) {
	return mayachainRouter.abi.Pack("depositWithExpiry", vault, asset, amount, memo, expiration)
}

// PackReturnVaultAssets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2923e82e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function returnVaultAssets(address router, address asgard, (address,uint256)[] coins, string memo) payable returns()
func (mayachainRouter *MayachainRouter) PackReturnVaultAssets(router common.Address, asgard common.Address, coins []THORChainRouterCoin, memo string) []byte {
	enc, err := mayachainRouter.abi.Pack("returnVaultAssets", router, asgard, coins, memo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReturnVaultAssets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2923e82e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function returnVaultAssets(address router, address asgard, (address,uint256)[] coins, string memo) payable returns()
func (mayachainRouter *MayachainRouter) TryPackReturnVaultAssets(router common.Address, asgard common.Address, coins []THORChainRouterCoin, memo string) ([]byte, error) {
	return mayachainRouter.abi.Pack("returnVaultAssets", router, asgard, coins, memo)
}

// PackTransferAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1b738b32.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferAllowance(address router, address newVault, address asset, uint256 amount, string memo) returns()
func (mayachainRouter *MayachainRouter) PackTransferAllowance(router common.Address, newVault common.Address, asset common.Address, amount *big.Int, memo string) []byte {
	enc, err := mayachainRouter.abi.Pack("transferAllowance", router, newVault, asset, amount, memo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1b738b32.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferAllowance(address router, address newVault, address asset, uint256 amount, string memo) returns()
func (mayachainRouter *MayachainRouter) TryPackTransferAllowance(router common.Address, newVault common.Address, asset common.Address, amount *big.Int, memo string) ([]byte, error) {
	return mayachainRouter.abi.Pack("transferAllowance", router, newVault, asset, amount, memo)
}

// PackTransferOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x574da717.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOut(address to, address asset, uint256 amount, string memo) payable returns()
func (mayachainRouter *MayachainRouter) PackTransferOut(to common.Address, asset common.Address, amount *big.Int, memo string) []byte {
	enc, err := mayachainRouter.abi.Pack("transferOut", to, asset, amount, memo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x574da717.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOut(address to, address asset, uint256 amount, string memo) payable returns()
func (mayachainRouter *MayachainRouter) TryPackTransferOut(to common.Address, asset common.Address, amount *big.Int, memo string) ([]byte, error) {
	return mayachainRouter.abi.Pack("transferOut", to, asset, amount, memo)
}

// PackTransferOutAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4039fd4b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOutAndCall(address aggregator, address finalToken, address to, uint256 amountOutMin, string memo) payable returns()
func (mayachainRouter *MayachainRouter) PackTransferOutAndCall(aggregator common.Address, finalToken common.Address, to common.Address, amountOutMin *big.Int, memo string) []byte {
	enc, err := mayachainRouter.abi.Pack("transferOutAndCall", aggregator, finalToken, to, amountOutMin, memo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferOutAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4039fd4b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferOutAndCall(address aggregator, address finalToken, address to, uint256 amountOutMin, string memo) payable returns()
func (mayachainRouter *MayachainRouter) TryPackTransferOutAndCall(aggregator common.Address, finalToken common.Address, to common.Address, amountOutMin *big.Int, memo string) ([]byte, error) {
	return mayachainRouter.abi.Pack("transferOutAndCall", aggregator, finalToken, to, amountOutMin, memo)
}

// PackVaultAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x03b6a673.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function vaultAllowance(address vault, address token) view returns(uint256 amount)
func (mayachainRouter *MayachainRouter) PackVaultAllowance(vault common.Address, token common.Address) []byte {
	enc, err := mayachainRouter.abi.Pack("vaultAllowance", vault, token)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVaultAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x03b6a673.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function vaultAllowance(address vault, address token) view returns(uint256 amount)
func (mayachainRouter *MayachainRouter) TryPackVaultAllowance(vault common.Address, token common.Address) ([]byte, error) {
	return mayachainRouter.abi.Pack("vaultAllowance", vault, token)
}

// UnpackVaultAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x03b6a673.
//
// Solidity: function vaultAllowance(address vault, address token) view returns(uint256 amount)
func (mayachainRouter *MayachainRouter) UnpackVaultAllowance(data []byte) (*big.Int, error) {
	out, err := mayachainRouter.abi.Unpack("vaultAllowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// MayachainRouterDeposit represents a Deposit event raised by the MayachainRouter contract.
type MayachainRouterDeposit struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Memo   string
	Raw    *types.Log // Blockchain specific contextual infos
}

const MayachainRouterDepositEventName = "Deposit"

// ContractEventName returns the user-defined event name.
func (MayachainRouterDeposit) ContractEventName() string {
	return MayachainRouterDepositEventName
}

// UnpackDepositEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Deposit(address indexed to, address indexed asset, uint256 amount, string memo)
func (mayachainRouter *MayachainRouter) UnpackDepositEvent(log *types.Log) (*MayachainRouterDeposit, error) {
	event := "Deposit"
	if len(log.Topics) == 0 || log.Topics[0] != mayachainRouter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MayachainRouterDeposit)
	if len(log.Data) > 0 {
		if err := mayachainRouter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range mayachainRouter.abi.Events[event].Inputs {
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

// MayachainRouterTransferAllowance represents a TransferAllowance event raised by the MayachainRouter contract.
type MayachainRouterTransferAllowance struct {
	OldVault common.Address
	NewVault common.Address
	Asset    common.Address
	Amount   *big.Int
	Memo     string
	Raw      *types.Log // Blockchain specific contextual infos
}

const MayachainRouterTransferAllowanceEventName = "TransferAllowance"

// ContractEventName returns the user-defined event name.
func (MayachainRouterTransferAllowance) ContractEventName() string {
	return MayachainRouterTransferAllowanceEventName
}

// UnpackTransferAllowanceEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferAllowance(address indexed oldVault, address indexed newVault, address asset, uint256 amount, string memo)
func (mayachainRouter *MayachainRouter) UnpackTransferAllowanceEvent(log *types.Log) (*MayachainRouterTransferAllowance, error) {
	event := "TransferAllowance"
	if len(log.Topics) == 0 || log.Topics[0] != mayachainRouter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MayachainRouterTransferAllowance)
	if len(log.Data) > 0 {
		if err := mayachainRouter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range mayachainRouter.abi.Events[event].Inputs {
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

// MayachainRouterTransferOut represents a TransferOut event raised by the MayachainRouter contract.
type MayachainRouterTransferOut struct {
	Vault  common.Address
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Memo   string
	Raw    *types.Log // Blockchain specific contextual infos
}

const MayachainRouterTransferOutEventName = "TransferOut"

// ContractEventName returns the user-defined event name.
func (MayachainRouterTransferOut) ContractEventName() string {
	return MayachainRouterTransferOutEventName
}

// UnpackTransferOutEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferOut(address indexed vault, address indexed to, address asset, uint256 amount, string memo)
func (mayachainRouter *MayachainRouter) UnpackTransferOutEvent(log *types.Log) (*MayachainRouterTransferOut, error) {
	event := "TransferOut"
	if len(log.Topics) == 0 || log.Topics[0] != mayachainRouter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MayachainRouterTransferOut)
	if len(log.Data) > 0 {
		if err := mayachainRouter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range mayachainRouter.abi.Events[event].Inputs {
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

// MayachainRouterTransferOutAndCall represents a TransferOutAndCall event raised by the MayachainRouter contract.
type MayachainRouterTransferOutAndCall struct {
	Vault        common.Address
	Target       common.Address
	Amount       *big.Int
	FinalAsset   common.Address
	To           common.Address
	AmountOutMin *big.Int
	Memo         string
	Raw          *types.Log // Blockchain specific contextual infos
}

const MayachainRouterTransferOutAndCallEventName = "TransferOutAndCall"

// ContractEventName returns the user-defined event name.
func (MayachainRouterTransferOutAndCall) ContractEventName() string {
	return MayachainRouterTransferOutAndCallEventName
}

// UnpackTransferOutAndCallEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferOutAndCall(address indexed vault, address target, uint256 amount, address finalAsset, address to, uint256 amountOutMin, string memo)
func (mayachainRouter *MayachainRouter) UnpackTransferOutAndCallEvent(log *types.Log) (*MayachainRouterTransferOutAndCall, error) {
	event := "TransferOutAndCall"
	if len(log.Topics) == 0 || log.Topics[0] != mayachainRouter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MayachainRouterTransferOutAndCall)
	if len(log.Data) > 0 {
		if err := mayachainRouter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range mayachainRouter.abi.Events[event].Inputs {
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

// MayachainRouterVaultTransfer represents a VaultTransfer event raised by the MayachainRouter contract.
type MayachainRouterVaultTransfer struct {
	OldVault common.Address
	NewVault common.Address
	Coins    []THORChainRouterCoin
	Memo     string
	Raw      *types.Log // Blockchain specific contextual infos
}

const MayachainRouterVaultTransferEventName = "VaultTransfer"

// ContractEventName returns the user-defined event name.
func (MayachainRouterVaultTransfer) ContractEventName() string {
	return MayachainRouterVaultTransferEventName
}

// UnpackVaultTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VaultTransfer(address indexed oldVault, address indexed newVault, (address,uint256)[] coins, string memo)
func (mayachainRouter *MayachainRouter) UnpackVaultTransferEvent(log *types.Log) (*MayachainRouterVaultTransfer, error) {
	event := "VaultTransfer"
	if len(log.Topics) == 0 || log.Topics[0] != mayachainRouter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MayachainRouterVaultTransfer)
	if len(log.Data) > 0 {
		if err := mayachainRouter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range mayachainRouter.abi.Events[event].Inputs {
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
