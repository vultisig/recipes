// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package thorchain_router

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

// ThorchainRouterMetaData contains all meta data concerning the ThorchainRouter contract.
var ThorchainRouterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rune\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"TransferAllowance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"TransferOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"finalAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"TransferOutAndCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structTHORChain_Router.Coin[]\",\"name\":\"coins\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"VaultTransfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RUNE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"depositWithExpiry\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"asgard\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTHORChain_Router.Coin[]\",\"name\":\"coins\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"returnVaultAssets\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferOut\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"aggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"finalToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferOutAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"vaultAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "ThorchainRouter",
}

// ThorchainRouter is an auto generated Go binding around an Ethereum contract.
type ThorchainRouter struct {
	abi abi.ABI
}

// NewThorchainRouter creates a new instance of ThorchainRouter.
func NewThorchainRouter() *ThorchainRouter {
	parsed, err := ThorchainRouterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ThorchainRouter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ThorchainRouter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address rune) returns()
func (thorchainRouter *ThorchainRouter) PackConstructor(rune common.Address) []byte {
	enc, err := thorchainRouter.abi.Pack("", rune)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRUNE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x93e4eaa9.
//
// Solidity: function RUNE() view returns(address)
func (thorchainRouter *ThorchainRouter) PackRUNE() []byte {
	enc, err := thorchainRouter.abi.Pack("RUNE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRUNE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x93e4eaa9.
//
// Solidity: function RUNE() view returns(address)
func (thorchainRouter *ThorchainRouter) UnpackRUNE(data []byte) (common.Address, error) {
	out, err := thorchainRouter.abi.Unpack("RUNE", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1fece7b4.
//
// Solidity: function deposit(address vault, address asset, uint256 amount, string memo) payable returns()
func (thorchainRouter *ThorchainRouter) PackDeposit(vault common.Address, asset common.Address, amount *big.Int, memo string) []byte {
	enc, err := thorchainRouter.abi.Pack("deposit", vault, asset, amount, memo)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDepositWithExpiry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44bc937b.
//
// Solidity: function depositWithExpiry(address vault, address asset, uint256 amount, string memo, uint256 expiration) payable returns()
func (thorchainRouter *ThorchainRouter) PackDepositWithExpiry(vault common.Address, asset common.Address, amount *big.Int, memo string, expiration *big.Int) []byte {
	enc, err := thorchainRouter.abi.Pack("depositWithExpiry", vault, asset, amount, memo, expiration)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackReturnVaultAssets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2923e82e.
//
// Solidity: function returnVaultAssets(address router, address asgard, (address,uint256)[] coins, string memo) payable returns()
func (thorchainRouter *ThorchainRouter) PackReturnVaultAssets(router common.Address, asgard common.Address, coins []THORChainRouterCoin, memo string) []byte {
	enc, err := thorchainRouter.abi.Pack("returnVaultAssets", router, asgard, coins, memo)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1b738b32.
//
// Solidity: function transferAllowance(address router, address newVault, address asset, uint256 amount, string memo) returns()
func (thorchainRouter *ThorchainRouter) PackTransferAllowance(router common.Address, newVault common.Address, asset common.Address, amount *big.Int, memo string) []byte {
	enc, err := thorchainRouter.abi.Pack("transferAllowance", router, newVault, asset, amount, memo)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x574da717.
//
// Solidity: function transferOut(address to, address asset, uint256 amount, string memo) payable returns()
func (thorchainRouter *ThorchainRouter) PackTransferOut(to common.Address, asset common.Address, amount *big.Int, memo string) []byte {
	enc, err := thorchainRouter.abi.Pack("transferOut", to, asset, amount, memo)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferOutAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4039fd4b.
//
// Solidity: function transferOutAndCall(address aggregator, address finalToken, address to, uint256 amountOutMin, string memo) payable returns()
func (thorchainRouter *ThorchainRouter) PackTransferOutAndCall(aggregator common.Address, finalToken common.Address, to common.Address, amountOutMin *big.Int, memo string) []byte {
	enc, err := thorchainRouter.abi.Pack("transferOutAndCall", aggregator, finalToken, to, amountOutMin, memo)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackVaultAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x03b6a673.
//
// Solidity: function vaultAllowance(address vault, address token) view returns(uint256 amount)
func (thorchainRouter *ThorchainRouter) PackVaultAllowance(vault common.Address, token common.Address) []byte {
	enc, err := thorchainRouter.abi.Pack("vaultAllowance", vault, token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackVaultAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x03b6a673.
//
// Solidity: function vaultAllowance(address vault, address token) view returns(uint256 amount)
func (thorchainRouter *ThorchainRouter) UnpackVaultAllowance(data []byte) (*big.Int, error) {
	out, err := thorchainRouter.abi.Unpack("vaultAllowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// ThorchainRouterDeposit represents a Deposit event raised by the ThorchainRouter contract.
type ThorchainRouterDeposit struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Memo   string
	Raw    *types.Log // Blockchain specific contextual infos
}

const ThorchainRouterDepositEventName = "Deposit"

// ContractEventName returns the user-defined event name.
func (ThorchainRouterDeposit) ContractEventName() string {
	return ThorchainRouterDepositEventName
}

// UnpackDepositEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Deposit(address indexed to, address indexed asset, uint256 amount, string memo)
func (thorchainRouter *ThorchainRouter) UnpackDepositEvent(log *types.Log) (*ThorchainRouterDeposit, error) {
	event := "Deposit"
	if log.Topics[0] != thorchainRouter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ThorchainRouterDeposit)
	if len(log.Data) > 0 {
		if err := thorchainRouter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range thorchainRouter.abi.Events[event].Inputs {
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

// ThorchainRouterTransferAllowance represents a TransferAllowance event raised by the ThorchainRouter contract.
type ThorchainRouterTransferAllowance struct {
	OldVault common.Address
	NewVault common.Address
	Asset    common.Address
	Amount   *big.Int
	Memo     string
	Raw      *types.Log // Blockchain specific contextual infos
}

const ThorchainRouterTransferAllowanceEventName = "TransferAllowance"

// ContractEventName returns the user-defined event name.
func (ThorchainRouterTransferAllowance) ContractEventName() string {
	return ThorchainRouterTransferAllowanceEventName
}

// UnpackTransferAllowanceEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferAllowance(address indexed oldVault, address indexed newVault, address asset, uint256 amount, string memo)
func (thorchainRouter *ThorchainRouter) UnpackTransferAllowanceEvent(log *types.Log) (*ThorchainRouterTransferAllowance, error) {
	event := "TransferAllowance"
	if log.Topics[0] != thorchainRouter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ThorchainRouterTransferAllowance)
	if len(log.Data) > 0 {
		if err := thorchainRouter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range thorchainRouter.abi.Events[event].Inputs {
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

// ThorchainRouterTransferOut represents a TransferOut event raised by the ThorchainRouter contract.
type ThorchainRouterTransferOut struct {
	Vault  common.Address
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	Memo   string
	Raw    *types.Log // Blockchain specific contextual infos
}

const ThorchainRouterTransferOutEventName = "TransferOut"

// ContractEventName returns the user-defined event name.
func (ThorchainRouterTransferOut) ContractEventName() string {
	return ThorchainRouterTransferOutEventName
}

// UnpackTransferOutEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferOut(address indexed vault, address indexed to, address asset, uint256 amount, string memo)
func (thorchainRouter *ThorchainRouter) UnpackTransferOutEvent(log *types.Log) (*ThorchainRouterTransferOut, error) {
	event := "TransferOut"
	if log.Topics[0] != thorchainRouter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ThorchainRouterTransferOut)
	if len(log.Data) > 0 {
		if err := thorchainRouter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range thorchainRouter.abi.Events[event].Inputs {
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

// ThorchainRouterTransferOutAndCall represents a TransferOutAndCall event raised by the ThorchainRouter contract.
type ThorchainRouterTransferOutAndCall struct {
	Vault        common.Address
	Target       common.Address
	Amount       *big.Int
	FinalAsset   common.Address
	To           common.Address
	AmountOutMin *big.Int
	Memo         string
	Raw          *types.Log // Blockchain specific contextual infos
}

const ThorchainRouterTransferOutAndCallEventName = "TransferOutAndCall"

// ContractEventName returns the user-defined event name.
func (ThorchainRouterTransferOutAndCall) ContractEventName() string {
	return ThorchainRouterTransferOutAndCallEventName
}

// UnpackTransferOutAndCallEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferOutAndCall(address indexed vault, address target, uint256 amount, address finalAsset, address to, uint256 amountOutMin, string memo)
func (thorchainRouter *ThorchainRouter) UnpackTransferOutAndCallEvent(log *types.Log) (*ThorchainRouterTransferOutAndCall, error) {
	event := "TransferOutAndCall"
	if log.Topics[0] != thorchainRouter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ThorchainRouterTransferOutAndCall)
	if len(log.Data) > 0 {
		if err := thorchainRouter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range thorchainRouter.abi.Events[event].Inputs {
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

// ThorchainRouterVaultTransfer represents a VaultTransfer event raised by the ThorchainRouter contract.
type ThorchainRouterVaultTransfer struct {
	OldVault common.Address
	NewVault common.Address
	Coins    []THORChainRouterCoin
	Memo     string
	Raw      *types.Log // Blockchain specific contextual infos
}

const ThorchainRouterVaultTransferEventName = "VaultTransfer"

// ContractEventName returns the user-defined event name.
func (ThorchainRouterVaultTransfer) ContractEventName() string {
	return ThorchainRouterVaultTransferEventName
}

// UnpackVaultTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VaultTransfer(address indexed oldVault, address indexed newVault, (address,uint256)[] coins, string memo)
func (thorchainRouter *ThorchainRouter) UnpackVaultTransferEvent(log *types.Log) (*ThorchainRouterVaultTransfer, error) {
	event := "VaultTransfer"
	if log.Topics[0] != thorchainRouter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ThorchainRouterVaultTransfer)
	if len(log.Data) > 0 {
		if err := thorchainRouter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range thorchainRouter.abi.Events[event].Inputs {
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
