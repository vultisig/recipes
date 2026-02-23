// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polymarket_ctf_exchange

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

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Salt          *big.Int
	Maker         common.Address
	Signer        common.Address
	Taker         common.Address
	TokenId       *big.Int
	MakerAmount   *big.Int
	TakerAmount   *big.Int
	Expiration    *big.Int
	Nonce         *big.Int
	FeeRateBps    *big.Int
	Side          uint8
	SignatureType uint8
	Signature     []byte
}

// PolymarketCtfExchangeMetaData contains all meta data concerning the PolymarketCtfExchange contract.
var PolymarketCtfExchangeMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fillAmounts\",\"type\":\"uint256[]\"}],\"name\":\"fillOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "PolymarketCtfExchange",
}

// PolymarketCtfExchange is an auto generated Go binding around an Ethereum contract.
type PolymarketCtfExchange struct {
	abi abi.ABI
}

// NewPolymarketCtfExchange creates a new instance of PolymarketCtfExchange.
func NewPolymarketCtfExchange() *PolymarketCtfExchange {
	parsed, err := PolymarketCtfExchangeMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &PolymarketCtfExchange{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *PolymarketCtfExchange) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCancelOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa6dfcf86.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (polymarketCtfExchange *PolymarketCtfExchange) PackCancelOrder(order Order) []byte {
	enc, err := polymarketCtfExchange.abi.Pack("cancelOrder", order)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa6dfcf86.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (polymarketCtfExchange *PolymarketCtfExchange) TryPackCancelOrder(order Order) ([]byte, error) {
	return polymarketCtfExchange.abi.Pack("cancelOrder", order)
}

// PackCancelOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa950b48.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (polymarketCtfExchange *PolymarketCtfExchange) PackCancelOrders(orders []Order) []byte {
	enc, err := polymarketCtfExchange.abi.Pack("cancelOrders", orders)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa950b48.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (polymarketCtfExchange *PolymarketCtfExchange) TryPackCancelOrders(orders []Order) ([]byte, error) {
	return polymarketCtfExchange.abi.Pack("cancelOrders", orders)
}

// PackFillOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe729aaf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (polymarketCtfExchange *PolymarketCtfExchange) PackFillOrder(order Order, fillAmount *big.Int) []byte {
	enc, err := polymarketCtfExchange.abi.Pack("fillOrder", order, fillAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe729aaf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (polymarketCtfExchange *PolymarketCtfExchange) TryPackFillOrder(order Order, fillAmount *big.Int) ([]byte, error) {
	return polymarketCtfExchange.abi.Pack("fillOrder", order, fillAmount)
}

// PackFillOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd798eff6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (polymarketCtfExchange *PolymarketCtfExchange) PackFillOrders(orders []Order, fillAmounts []*big.Int) []byte {
	enc, err := polymarketCtfExchange.abi.Pack("fillOrders", orders, fillAmounts)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd798eff6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (polymarketCtfExchange *PolymarketCtfExchange) TryPackFillOrders(orders []Order, fillAmounts []*big.Int) ([]byte, error) {
	return polymarketCtfExchange.abi.Pack("fillOrders", orders, fillAmounts)
}
