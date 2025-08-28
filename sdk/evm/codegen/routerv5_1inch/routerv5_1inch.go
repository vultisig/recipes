// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package routerv5_1inch

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

// GenericRouterSwapDescription is an auto generated low-level Go binding around an user-defined struct.
type GenericRouterSwapDescription struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceiver     common.Address
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
}

// OrderLibOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderLibOrder struct {
	Salt          *big.Int
	MakerAsset    common.Address
	TakerAsset    common.Address
	Maker         common.Address
	Receiver      common.Address
	AllowedSender common.Address
	MakingAmount  *big.Int
	TakingAmount  *big.Int
	Offsets       *big.Int
	Interactions  []byte
}

// OrderRFQLibOrderRFQ is an auto generated low-level Go binding around an user-defined struct.
type OrderRFQLibOrderRFQ struct {
	Info          *big.Int
	MakerAsset    common.Address
	TakerAsset    common.Address
	Maker         common.Address
	AllowedSender common.Address
	MakingAmount  *big.Int
	TakingAmount  *big.Int
}

// Routerv51inchMetaData contains all meta data concerning the Routerv51inch contract.
var Routerv51inchMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AdvanceNonceFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArbitraryStaticCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthDepositRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetAmountCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidatedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOneAmountShouldBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermitLengthTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PredicateIsNotTrue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrivateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQBadSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQPrivateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQSwapWithZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQZeroTargetIsForbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyDetected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemainingAmountIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReservesCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReturnAmountIsNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafePermitBadLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFromFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"name\":\"SimulationResults\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapAmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapWithZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountIncreased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromMakerToTakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromTakerToMakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongGetter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMinReturn\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroReturnAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroTargetIsForbidden\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingRaw\",\"type\":\"uint256\"}],\"name\":\"OrderCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilledRFQ\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"amount\",\"type\":\"uint8\"}],\"name\":\"advanceNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"and\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"arbitraryStaticCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderRemaining\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalMask\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"checkPredicate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"clipperSwapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"eq\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQCompact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filledMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filledTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderRFQTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filledMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filledTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"fillOrderRFQToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"fillOrderToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"gt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"invalidatorForOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"lt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"name\":\"nonceEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"or\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remainingRaw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"remainingsRaw\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"executor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structGenericRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"timestampBelow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNonceAccount\",\"type\":\"uint256\"}],\"name\":\"timestampBelowAndNonceEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3Swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3SwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"unoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"unoswapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"unoswapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "Routerv51inch",
}

// Routerv51inch is an auto generated Go binding around an Ethereum contract.
type Routerv51inch struct {
	abi abi.ABI
}

// NewRouterv51inch creates a new instance of Routerv51inch.
func NewRouterv51inch() *Routerv51inch {
	parsed, err := Routerv51inchMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Routerv51inch{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Routerv51inch) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address weth) returns()
func (routerv51inch *Routerv51inch) PackConstructor(weth common.Address) []byte {
	enc, err := routerv51inch.abi.Pack("", weth)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAdvanceNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72c244a8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (routerv51inch *Routerv51inch) PackAdvanceNonce(amount uint8) []byte {
	enc, err := routerv51inch.abi.Pack("advanceNonce", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAdvanceNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x72c244a8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (routerv51inch *Routerv51inch) TryPackAdvanceNonce(amount uint8) ([]byte, error) {
	return routerv51inch.abi.Pack("advanceNonce", amount)
}

// PackAnd is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbfa75143.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) PackAnd(offsets *big.Int, data []byte) []byte {
	enc, err := routerv51inch.abi.Pack("and", offsets, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAnd is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbfa75143.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) TryPackAnd(offsets *big.Int, data []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("and", offsets, data)
}

// UnpackAnd is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) UnpackAnd(data []byte) (bool, error) {
	out, err := routerv51inch.abi.Unpack("and", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackArbitraryStaticCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf15fcd8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (routerv51inch *Routerv51inch) PackArbitraryStaticCall(target common.Address, data []byte) []byte {
	enc, err := routerv51inch.abi.Pack("arbitraryStaticCall", target, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackArbitraryStaticCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf15fcd8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (routerv51inch *Routerv51inch) TryPackArbitraryStaticCall(target common.Address, data []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("arbitraryStaticCall", target, data)
}

// UnpackArbitraryStaticCall is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (routerv51inch *Routerv51inch) UnpackArbitraryStaticCall(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("arbitraryStaticCall", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCancelOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d9a56f6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (routerv51inch *Routerv51inch) PackCancelOrder(order OrderLibOrder) []byte {
	enc, err := routerv51inch.abi.Pack("cancelOrder", order)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d9a56f6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (routerv51inch *Routerv51inch) TryPackCancelOrder(order OrderLibOrder) ([]byte, error) {
	return routerv51inch.abi.Pack("cancelOrder", order)
}

// CancelOrderOutput serves as a container for the return parameters of contract
// method CancelOrder.
type CancelOrderOutput struct {
	OrderRemaining *big.Int
	OrderHash      [32]byte
}

// UnpackCancelOrder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d9a56f6.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (routerv51inch *Routerv51inch) UnpackCancelOrder(data []byte) (CancelOrderOutput, error) {
	out, err := routerv51inch.abi.Unpack("cancelOrder", data)
	outstruct := new(CancelOrderOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.OrderRemaining = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.OrderHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackCancelOrderRFQ is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x825caba1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (routerv51inch *Routerv51inch) PackCancelOrderRFQ(orderInfo *big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("cancelOrderRFQ", orderInfo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelOrderRFQ is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x825caba1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (routerv51inch *Routerv51inch) TryPackCancelOrderRFQ(orderInfo *big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("cancelOrderRFQ", orderInfo)
}

// PackCancelOrderRFQ0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbddccd35.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (routerv51inch *Routerv51inch) PackCancelOrderRFQ0(orderInfo *big.Int, additionalMask *big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("cancelOrderRFQ0", orderInfo, additionalMask)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelOrderRFQ0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbddccd35.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (routerv51inch *Routerv51inch) TryPackCancelOrderRFQ0(orderInfo *big.Int, additionalMask *big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("cancelOrderRFQ0", orderInfo, additionalMask)
}

// PackCheckPredicate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6c838250.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (routerv51inch *Routerv51inch) PackCheckPredicate(order OrderLibOrder) []byte {
	enc, err := routerv51inch.abi.Pack("checkPredicate", order)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCheckPredicate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6c838250.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (routerv51inch *Routerv51inch) TryPackCheckPredicate(order OrderLibOrder) ([]byte, error) {
	return routerv51inch.abi.Pack("checkPredicate", order)
}

// UnpackCheckPredicate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6c838250.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (routerv51inch *Routerv51inch) UnpackCheckPredicate(data []byte) (bool, error) {
	out, err := routerv51inch.abi.Unpack("checkPredicate", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackClipperSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84bd6d29.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) PackClipperSwap(clipperExchange common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) []byte {
	enc, err := routerv51inch.abi.Pack("clipperSwap", clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClipperSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84bd6d29.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) TryPackClipperSwap(clipperExchange common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) ([]byte, error) {
	return routerv51inch.abi.Pack("clipperSwap", clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// UnpackClipperSwap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84bd6d29.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) UnpackClipperSwap(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("clipperSwap", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackClipperSwapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x093d4fa5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) PackClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) []byte {
	enc, err := routerv51inch.abi.Pack("clipperSwapTo", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClipperSwapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x093d4fa5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) TryPackClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) ([]byte, error) {
	return routerv51inch.abi.Pack("clipperSwapTo", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// UnpackClipperSwapTo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x093d4fa5.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) UnpackClipperSwapTo(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("clipperSwapTo", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackClipperSwapToWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc805a666.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) PackClipperSwapToWithPermit(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte, permit []byte) []byte {
	enc, err := routerv51inch.abi.Pack("clipperSwapToWithPermit", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs, permit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClipperSwapToWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc805a666.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) TryPackClipperSwapToWithPermit(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte, permit []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("clipperSwapToWithPermit", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs, permit)
}

// UnpackClipperSwapToWithPermit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc805a666.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) UnpackClipperSwapToWithPermit(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("clipperSwapToWithPermit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackDestroy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83197ef0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function destroy() returns()
func (routerv51inch *Routerv51inch) PackDestroy() []byte {
	enc, err := routerv51inch.abi.Pack("destroy")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDestroy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83197ef0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function destroy() returns()
func (routerv51inch *Routerv51inch) TryPackDestroy() ([]byte, error) {
	return routerv51inch.abi.Pack("destroy")
}

// PackEq is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6fe7b0ba.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) PackEq(value *big.Int, data []byte) []byte {
	enc, err := routerv51inch.abi.Pack("eq", value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEq is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6fe7b0ba.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) TryPackEq(value *big.Int, data []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("eq", value, data)
}

// UnpackEq is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) UnpackEq(data []byte) (bool, error) {
	out, err := routerv51inch.abi.Unpack("eq", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackFillOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62e238bb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) PackFillOrder(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("fillOrder", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62e238bb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) TryPackFillOrder(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("fillOrder", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount)
}

// FillOrderOutput serves as a container for the return parameters of contract
// method FillOrder.
type FillOrderOutput struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 [32]byte
}

// UnpackFillOrder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62e238bb.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) UnpackFillOrder(data []byte) (FillOrderOutput, error) {
	out, err := routerv51inch.abi.Unpack("fillOrder", data)
	outstruct := new(FillOrderOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Arg2 = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackFillOrderRFQ is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3eca9c0a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) PackFillOrderRFQ(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("fillOrderRFQ", order, signature, flagsAndAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrderRFQ is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3eca9c0a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) TryPackFillOrderRFQ(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("fillOrderRFQ", order, signature, flagsAndAmount)
}

// FillOrderRFQOutput serves as a container for the return parameters of contract
// method FillOrderRFQ.
type FillOrderRFQOutput struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 [32]byte
}

// UnpackFillOrderRFQ is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3eca9c0a.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) UnpackFillOrderRFQ(data []byte) (FillOrderRFQOutput, error) {
	out, err := routerv51inch.abi.Unpack("fillOrderRFQ", data)
	outstruct := new(FillOrderRFQOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Arg2 = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackFillOrderRFQCompact is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9570eeee.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (routerv51inch *Routerv51inch) PackFillOrderRFQCompact(order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("fillOrderRFQCompact", order, r, vs, flagsAndAmount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrderRFQCompact is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9570eeee.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (routerv51inch *Routerv51inch) TryPackFillOrderRFQCompact(order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("fillOrderRFQCompact", order, r, vs, flagsAndAmount)
}

// FillOrderRFQCompactOutput serves as a container for the return parameters of contract
// method FillOrderRFQCompact.
type FillOrderRFQCompactOutput struct {
	FilledMakingAmount *big.Int
	FilledTakingAmount *big.Int
	OrderHash          [32]byte
}

// UnpackFillOrderRFQCompact is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9570eeee.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (routerv51inch *Routerv51inch) UnpackFillOrderRFQCompact(data []byte) (FillOrderRFQCompactOutput, error) {
	out, err := routerv51inch.abi.Unpack("fillOrderRFQCompact", data)
	outstruct := new(FillOrderRFQCompactOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.FilledMakingAmount = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.FilledTakingAmount = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.OrderHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackFillOrderRFQTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a099843.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (routerv51inch *Routerv51inch) PackFillOrderRFQTo(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) []byte {
	enc, err := routerv51inch.abi.Pack("fillOrderRFQTo", order, signature, flagsAndAmount, target)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrderRFQTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a099843.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (routerv51inch *Routerv51inch) TryPackFillOrderRFQTo(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) ([]byte, error) {
	return routerv51inch.abi.Pack("fillOrderRFQTo", order, signature, flagsAndAmount, target)
}

// FillOrderRFQToOutput serves as a container for the return parameters of contract
// method FillOrderRFQTo.
type FillOrderRFQToOutput struct {
	FilledMakingAmount *big.Int
	FilledTakingAmount *big.Int
	OrderHash          [32]byte
}

// UnpackFillOrderRFQTo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5a099843.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (routerv51inch *Routerv51inch) UnpackFillOrderRFQTo(data []byte) (FillOrderRFQToOutput, error) {
	out, err := routerv51inch.abi.Unpack("fillOrderRFQTo", data)
	outstruct := new(FillOrderRFQToOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.FilledMakingAmount = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.FilledTakingAmount = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.OrderHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackFillOrderRFQToWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70ccbd31.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) PackFillOrderRFQToWithPermit(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) []byte {
	enc, err := routerv51inch.abi.Pack("fillOrderRFQToWithPermit", order, signature, flagsAndAmount, target, permit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrderRFQToWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70ccbd31.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) TryPackFillOrderRFQToWithPermit(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("fillOrderRFQToWithPermit", order, signature, flagsAndAmount, target, permit)
}

// FillOrderRFQToWithPermitOutput serves as a container for the return parameters of contract
// method FillOrderRFQToWithPermit.
type FillOrderRFQToWithPermitOutput struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 [32]byte
}

// UnpackFillOrderRFQToWithPermit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70ccbd31.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) UnpackFillOrderRFQToWithPermit(data []byte) (FillOrderRFQToWithPermitOutput, error) {
	out, err := routerv51inch.abi.Unpack("fillOrderRFQToWithPermit", data)
	outstruct := new(FillOrderRFQToWithPermitOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Arg2 = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackFillOrderTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe5d7bde6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (routerv51inch *Routerv51inch) PackFillOrderTo(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address) []byte {
	enc, err := routerv51inch.abi.Pack("fillOrderTo", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrderTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe5d7bde6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (routerv51inch *Routerv51inch) TryPackFillOrderTo(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address) ([]byte, error) {
	return routerv51inch.abi.Pack("fillOrderTo", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target)
}

// FillOrderToOutput serves as a container for the return parameters of contract
// method FillOrderTo.
type FillOrderToOutput struct {
	ActualMakingAmount *big.Int
	ActualTakingAmount *big.Int
	OrderHash          [32]byte
}

// UnpackFillOrderTo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe5d7bde6.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (routerv51inch *Routerv51inch) UnpackFillOrderTo(data []byte) (FillOrderToOutput, error) {
	out, err := routerv51inch.abi.Unpack("fillOrderTo", data)
	outstruct := new(FillOrderToOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ActualMakingAmount = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.ActualTakingAmount = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.OrderHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackFillOrderToWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd365c695.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) PackFillOrderToWithPermit(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address, permit []byte) []byte {
	enc, err := routerv51inch.abi.Pack("fillOrderToWithPermit", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target, permit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrderToWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd365c695.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) TryPackFillOrderToWithPermit(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address, permit []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("fillOrderToWithPermit", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target, permit)
}

// FillOrderToWithPermitOutput serves as a container for the return parameters of contract
// method FillOrderToWithPermit.
type FillOrderToWithPermitOutput struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 [32]byte
}

// UnpackFillOrderToWithPermit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd365c695.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (routerv51inch *Routerv51inch) UnpackFillOrderToWithPermit(data []byte) (FillOrderToWithPermitOutput, error) {
	out, err := routerv51inch.abi.Unpack("fillOrderToWithPermit", data)
	outstruct := new(FillOrderToWithPermitOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Arg2 = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackGt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f38e2b8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) PackGt(value *big.Int, data []byte) []byte {
	enc, err := routerv51inch.abi.Pack("gt", value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f38e2b8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) TryPackGt(value *big.Int, data []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("gt", value, data)
}

// UnpackGt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) UnpackGt(data []byte) (bool, error) {
	out, err := routerv51inch.abi.Unpack("gt", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHashOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x37e7316f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (routerv51inch *Routerv51inch) PackHashOrder(order OrderLibOrder) []byte {
	enc, err := routerv51inch.abi.Pack("hashOrder", order)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHashOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x37e7316f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (routerv51inch *Routerv51inch) TryPackHashOrder(order OrderLibOrder) ([]byte, error) {
	return routerv51inch.abi.Pack("hashOrder", order)
}

// UnpackHashOrder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x37e7316f.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (routerv51inch *Routerv51inch) UnpackHashOrder(data []byte) ([32]byte, error) {
	out, err := routerv51inch.abi.Unpack("hashOrder", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackIncreaseNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc53a0292.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function increaseNonce() returns()
func (routerv51inch *Routerv51inch) PackIncreaseNonce() []byte {
	enc, err := routerv51inch.abi.Pack("increaseNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIncreaseNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc53a0292.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function increaseNonce() returns()
func (routerv51inch *Routerv51inch) TryPackIncreaseNonce() ([]byte, error) {
	return routerv51inch.abi.Pack("increaseNonce")
}

// PackInvalidatorForOrderRFQ is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x56f16124.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (routerv51inch *Routerv51inch) PackInvalidatorForOrderRFQ(maker common.Address, slot *big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("invalidatorForOrderRFQ", maker, slot)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInvalidatorForOrderRFQ is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x56f16124.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (routerv51inch *Routerv51inch) TryPackInvalidatorForOrderRFQ(maker common.Address, slot *big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("invalidatorForOrderRFQ", maker, slot)
}

// UnpackInvalidatorForOrderRFQ is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (routerv51inch *Routerv51inch) UnpackInvalidatorForOrderRFQ(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("invalidatorForOrderRFQ", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackLt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca4ece22.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) PackLt(value *big.Int, data []byte) []byte {
	enc, err := routerv51inch.abi.Pack("lt", value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca4ece22.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) TryPackLt(value *big.Int, data []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("lt", value, data)
}

// UnpackLt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) UnpackLt(data []byte) (bool, error) {
	out, err := routerv51inch.abi.Unpack("lt", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70ae92d2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nonce(address ) view returns(uint256)
func (routerv51inch *Routerv51inch) PackNonce(arg0 common.Address) []byte {
	enc, err := routerv51inch.abi.Pack("nonce", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70ae92d2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function nonce(address ) view returns(uint256)
func (routerv51inch *Routerv51inch) TryPackNonce(arg0 common.Address) ([]byte, error) {
	return routerv51inch.abi.Pack("nonce", arg0)
}

// UnpackNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (routerv51inch *Routerv51inch) UnpackNonce(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("nonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackNonceEquals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcf6fc6e3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (routerv51inch *Routerv51inch) PackNonceEquals(makerAddress common.Address, makerNonce *big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("nonceEquals", makerAddress, makerNonce)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNonceEquals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcf6fc6e3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (routerv51inch *Routerv51inch) TryPackNonceEquals(makerAddress common.Address, makerNonce *big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("nonceEquals", makerAddress, makerNonce)
}

// UnpackNonceEquals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (routerv51inch *Routerv51inch) UnpackNonceEquals(data []byte) (bool, error) {
	out, err := routerv51inch.abi.Unpack("nonceEquals", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackOr is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74261145.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) PackOr(offsets *big.Int, data []byte) []byte {
	enc, err := routerv51inch.abi.Pack("or", offsets, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOr is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74261145.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) TryPackOr(offsets *big.Int, data []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("or", offsets, data)
}

// UnpackOr is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (routerv51inch *Routerv51inch) UnpackOr(data []byte) (bool, error) {
	out, err := routerv51inch.abi.Unpack("or", data)
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
func (routerv51inch *Routerv51inch) PackOwner() []byte {
	enc, err := routerv51inch.abi.Pack("owner")
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
func (routerv51inch *Routerv51inch) TryPackOwner() ([]byte, error) {
	return routerv51inch.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (routerv51inch *Routerv51inch) UnpackOwner(data []byte) (common.Address, error) {
	out, err := routerv51inch.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRemaining is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc1ed74c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (routerv51inch *Routerv51inch) PackRemaining(orderHash [32]byte) []byte {
	enc, err := routerv51inch.abi.Pack("remaining", orderHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemaining is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc1ed74c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (routerv51inch *Routerv51inch) TryPackRemaining(orderHash [32]byte) ([]byte, error) {
	return routerv51inch.abi.Pack("remaining", orderHash)
}

// UnpackRemaining is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (routerv51inch *Routerv51inch) UnpackRemaining(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("remaining", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRemainingRaw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7e54f092.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (routerv51inch *Routerv51inch) PackRemainingRaw(orderHash [32]byte) []byte {
	enc, err := routerv51inch.abi.Pack("remainingRaw", orderHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemainingRaw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7e54f092.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (routerv51inch *Routerv51inch) TryPackRemainingRaw(orderHash [32]byte) ([]byte, error) {
	return routerv51inch.abi.Pack("remainingRaw", orderHash)
}

// UnpackRemainingRaw is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (routerv51inch *Routerv51inch) UnpackRemainingRaw(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("remainingRaw", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRemainingsRaw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x942461bb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (routerv51inch *Routerv51inch) PackRemainingsRaw(orderHashes [][32]byte) []byte {
	enc, err := routerv51inch.abi.Pack("remainingsRaw", orderHashes)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemainingsRaw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x942461bb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (routerv51inch *Routerv51inch) TryPackRemainingsRaw(orderHashes [][32]byte) ([]byte, error) {
	return routerv51inch.abi.Pack("remainingsRaw", orderHashes)
}

// UnpackRemainingsRaw is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (routerv51inch *Routerv51inch) UnpackRemainingsRaw(data []byte) ([]*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("remainingsRaw", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (routerv51inch *Routerv51inch) PackRenounceOwnership() []byte {
	enc, err := routerv51inch.abi.Pack("renounceOwnership")
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
func (routerv51inch *Routerv51inch) TryPackRenounceOwnership() ([]byte, error) {
	return routerv51inch.abi.Pack("renounceOwnership")
}

// PackRescueFunds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78e3214f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (routerv51inch *Routerv51inch) PackRescueFunds(token common.Address, amount *big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("rescueFunds", token, amount)
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
func (routerv51inch *Routerv51inch) TryPackRescueFunds(token common.Address, amount *big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("rescueFunds", token, amount)
}

// PackSimulate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd61951d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function simulate(address target, bytes data) returns()
func (routerv51inch *Routerv51inch) PackSimulate(target common.Address, data []byte) []byte {
	enc, err := routerv51inch.abi.Pack("simulate", target, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSimulate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd61951d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function simulate(address target, bytes data) returns()
func (routerv51inch *Routerv51inch) TryPackSimulate(target common.Address, data []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("simulate", target, data)
}

// PackSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12aa3caf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (routerv51inch *Routerv51inch) PackSwap(executor common.Address, desc GenericRouterSwapDescription, permit []byte, data []byte) []byte {
	enc, err := routerv51inch.abi.Pack("swap", executor, desc, permit, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12aa3caf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (routerv51inch *Routerv51inch) TryPackSwap(executor common.Address, desc GenericRouterSwapDescription, permit []byte, data []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("swap", executor, desc, permit, data)
}

// SwapOutput serves as a container for the return parameters of contract
// method Swap.
type SwapOutput struct {
	ReturnAmount *big.Int
	SpentAmount  *big.Int
}

// UnpackSwap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x12aa3caf.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (routerv51inch *Routerv51inch) UnpackSwap(data []byte) (SwapOutput, error) {
	out, err := routerv51inch.abi.Unpack("swap", data)
	outstruct := new(SwapOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ReturnAmount = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.SpentAmount = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackTimestampBelow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x63592c2b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (routerv51inch *Routerv51inch) PackTimestampBelow(time *big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("timestampBelow", time)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTimestampBelow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x63592c2b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (routerv51inch *Routerv51inch) TryPackTimestampBelow(time *big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("timestampBelow", time)
}

// UnpackTimestampBelow is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (routerv51inch *Routerv51inch) UnpackTimestampBelow(data []byte) (bool, error) {
	out, err := routerv51inch.abi.Unpack("timestampBelow", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTimestampBelowAndNonceEquals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2cc2878d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (routerv51inch *Routerv51inch) PackTimestampBelowAndNonceEquals(timeNonceAccount *big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("timestampBelowAndNonceEquals", timeNonceAccount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTimestampBelowAndNonceEquals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2cc2878d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (routerv51inch *Routerv51inch) TryPackTimestampBelowAndNonceEquals(timeNonceAccount *big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("timestampBelowAndNonceEquals", timeNonceAccount)
}

// UnpackTimestampBelowAndNonceEquals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2cc2878d.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (routerv51inch *Routerv51inch) UnpackTimestampBelowAndNonceEquals(data []byte) (bool, error) {
	out, err := routerv51inch.abi.Unpack("timestampBelowAndNonceEquals", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (routerv51inch *Routerv51inch) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := routerv51inch.abi.Pack("transferOwnership", newOwner)
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
func (routerv51inch *Routerv51inch) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return routerv51inch.abi.Pack("transferOwnership", newOwner)
}

// PackUniswapV3Swap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe449022e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) PackUniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("uniswapV3Swap", amount, minReturn, pools)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUniswapV3Swap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe449022e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) TryPackUniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("uniswapV3Swap", amount, minReturn, pools)
}

// UnpackUniswapV3Swap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) UnpackUniswapV3Swap(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("uniswapV3Swap", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUniswapV3SwapCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa461e33.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (routerv51inch *Routerv51inch) PackUniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) []byte {
	enc, err := routerv51inch.abi.Pack("uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUniswapV3SwapCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa461e33.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (routerv51inch *Routerv51inch) TryPackUniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// PackUniswapV3SwapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc80f1a8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) PackUniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("uniswapV3SwapTo", recipient, amount, minReturn, pools)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUniswapV3SwapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbc80f1a8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) TryPackUniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("uniswapV3SwapTo", recipient, amount, minReturn, pools)
}

// UnpackUniswapV3SwapTo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) UnpackUniswapV3SwapTo(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("uniswapV3SwapTo", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUniswapV3SwapToWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2521b930.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) PackUniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) []byte {
	enc, err := routerv51inch.abi.Pack("uniswapV3SwapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUniswapV3SwapToWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2521b930.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) TryPackUniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("uniswapV3SwapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
}

// UnpackUniswapV3SwapToWithPermit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) UnpackUniswapV3SwapToWithPermit(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("uniswapV3SwapToWithPermit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUnoswap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0502b1c5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) PackUnoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("unoswap", srcToken, amount, minReturn, pools)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnoswap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0502b1c5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) TryPackUnoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("unoswap", srcToken, amount, minReturn, pools)
}

// UnpackUnoswap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0502b1c5.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) UnpackUnoswap(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("unoswap", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUnoswapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf78dc253.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) PackUnoswapTo(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) []byte {
	enc, err := routerv51inch.abi.Pack("unoswapTo", recipient, srcToken, amount, minReturn, pools)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnoswapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf78dc253.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) TryPackUnoswapTo(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) ([]byte, error) {
	return routerv51inch.abi.Pack("unoswapTo", recipient, srcToken, amount, minReturn, pools)
}

// UnpackUnoswapTo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf78dc253.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) UnpackUnoswapTo(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("unoswapTo", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUnoswapToWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c15fd91.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) PackUnoswapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) []byte {
	enc, err := routerv51inch.abi.Pack("unoswapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnoswapToWithPermit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c15fd91.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) TryPackUnoswapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) ([]byte, error) {
	return routerv51inch.abi.Pack("unoswapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
}

// UnpackUnoswapToWithPermit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3c15fd91.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (routerv51inch *Routerv51inch) UnpackUnoswapToWithPermit(data []byte) (*big.Int, error) {
	out, err := routerv51inch.abi.Unpack("unoswapToWithPermit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// Routerv51inchNonceIncreased represents a NonceIncreased event raised by the Routerv51inch contract.
type Routerv51inchNonceIncreased struct {
	Maker    common.Address
	NewNonce *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const Routerv51inchNonceIncreasedEventName = "NonceIncreased"

// ContractEventName returns the user-defined event name.
func (Routerv51inchNonceIncreased) ContractEventName() string {
	return Routerv51inchNonceIncreasedEventName
}

// UnpackNonceIncreasedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (routerv51inch *Routerv51inch) UnpackNonceIncreasedEvent(log *types.Log) (*Routerv51inchNonceIncreased, error) {
	event := "NonceIncreased"
	if log.Topics[0] != routerv51inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv51inchNonceIncreased)
	if len(log.Data) > 0 {
		if err := routerv51inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv51inch.abi.Events[event].Inputs {
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

// Routerv51inchOrderCanceled represents a OrderCanceled event raised by the Routerv51inch contract.
type Routerv51inchOrderCanceled struct {
	Maker        common.Address
	OrderHash    [32]byte
	RemainingRaw *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const Routerv51inchOrderCanceledEventName = "OrderCanceled"

// ContractEventName returns the user-defined event name.
func (Routerv51inchOrderCanceled) ContractEventName() string {
	return Routerv51inchOrderCanceledEventName
}

// UnpackOrderCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OrderCanceled(address indexed maker, bytes32 orderHash, uint256 remainingRaw)
func (routerv51inch *Routerv51inch) UnpackOrderCanceledEvent(log *types.Log) (*Routerv51inchOrderCanceled, error) {
	event := "OrderCanceled"
	if log.Topics[0] != routerv51inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv51inchOrderCanceled)
	if len(log.Data) > 0 {
		if err := routerv51inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv51inch.abi.Events[event].Inputs {
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

// Routerv51inchOrderFilled represents a OrderFilled event raised by the Routerv51inch contract.
type Routerv51inchOrderFilled struct {
	Maker     common.Address
	OrderHash [32]byte
	Remaining *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const Routerv51inchOrderFilledEventName = "OrderFilled"

// ContractEventName returns the user-defined event name.
func (Routerv51inchOrderFilled) ContractEventName() string {
	return Routerv51inchOrderFilledEventName
}

// UnpackOrderFilledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (routerv51inch *Routerv51inch) UnpackOrderFilledEvent(log *types.Log) (*Routerv51inchOrderFilled, error) {
	event := "OrderFilled"
	if log.Topics[0] != routerv51inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv51inchOrderFilled)
	if len(log.Data) > 0 {
		if err := routerv51inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv51inch.abi.Events[event].Inputs {
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

// Routerv51inchOrderFilledRFQ represents a OrderFilledRFQ event raised by the Routerv51inch contract.
type Routerv51inchOrderFilledRFQ struct {
	OrderHash    [32]byte
	MakingAmount *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const Routerv51inchOrderFilledRFQEventName = "OrderFilledRFQ"

// ContractEventName returns the user-defined event name.
func (Routerv51inchOrderFilledRFQ) ContractEventName() string {
	return Routerv51inchOrderFilledRFQEventName
}

// UnpackOrderFilledRFQEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (routerv51inch *Routerv51inch) UnpackOrderFilledRFQEvent(log *types.Log) (*Routerv51inchOrderFilledRFQ, error) {
	event := "OrderFilledRFQ"
	if log.Topics[0] != routerv51inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv51inchOrderFilledRFQ)
	if len(log.Data) > 0 {
		if err := routerv51inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv51inch.abi.Events[event].Inputs {
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

// Routerv51inchOwnershipTransferred represents a OwnershipTransferred event raised by the Routerv51inch contract.
type Routerv51inchOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const Routerv51inchOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (Routerv51inchOwnershipTransferred) ContractEventName() string {
	return Routerv51inchOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (routerv51inch *Routerv51inch) UnpackOwnershipTransferredEvent(log *types.Log) (*Routerv51inchOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != routerv51inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv51inchOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := routerv51inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv51inch.abi.Events[event].Inputs {
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
func (routerv51inch *Routerv51inch) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["AccessDenied"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackAccessDeniedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["AdvanceNonceFailed"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackAdvanceNonceFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["AlreadyFilled"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackAlreadyFilledError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["ArbitraryStaticCallFailed"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackArbitraryStaticCallFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["BadPool"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackBadPoolError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["BadSignature"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackBadSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["ETHTransferFailed"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackETHTransferFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["EmptyPools"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackEmptyPoolsError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["EthDepositRejected"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackEthDepositRejectedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["GetAmountCallFailed"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackGetAmountCallFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["IncorrectDataLength"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackIncorrectDataLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["InvalidMsgValue"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackInvalidMsgValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["InvalidatedOrder"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackInvalidatedOrderError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["MakingAmountExceeded"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackMakingAmountExceededError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["MakingAmountTooLow"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackMakingAmountTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["OnlyOneAmountShouldBeZero"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackOnlyOneAmountShouldBeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["OrderExpired"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackOrderExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["PermitLengthTooLow"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackPermitLengthTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["PredicateIsNotTrue"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackPredicateIsNotTrueError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["PrivateOrder"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackPrivateOrderError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["RFQBadSignature"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackRFQBadSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["RFQPrivateOrder"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackRFQPrivateOrderError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["RFQSwapWithZeroAmount"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackRFQSwapWithZeroAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["RFQZeroTargetIsForbidden"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackRFQZeroTargetIsForbiddenError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["ReentrancyDetected"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackReentrancyDetectedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["RemainingAmountIsZero"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackRemainingAmountIsZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["ReservesCallFailed"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackReservesCallFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["ReturnAmountIsNotEnough"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackReturnAmountIsNotEnoughError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["SafePermitBadLength"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackSafePermitBadLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["SafeTransferFailed"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackSafeTransferFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["SafeTransferFromFailed"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackSafeTransferFromFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["SimulationResults"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackSimulationResultsError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["SwapAmountTooLarge"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackSwapAmountTooLargeError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["SwapWithZeroAmount"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackSwapWithZeroAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["TakingAmountExceeded"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackTakingAmountExceededError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["TakingAmountIncreased"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackTakingAmountIncreasedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["TakingAmountTooHigh"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackTakingAmountTooHighError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["TransferFromMakerToTakerFailed"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackTransferFromMakerToTakerFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["TransferFromTakerToMakerFailed"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackTransferFromTakerToMakerFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["UnknownOrder"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackUnknownOrderError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["WrongAmount"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackWrongAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["WrongGetter"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackWrongGetterError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["ZeroAddress"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["ZeroMinReturn"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackZeroMinReturnError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["ZeroReturnAmount"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackZeroReturnAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv51inch.abi.Errors["ZeroTargetIsForbidden"].ID.Bytes()[:4]) {
		return routerv51inch.UnpackZeroTargetIsForbiddenError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// Routerv51inchAccessDenied represents a AccessDenied error raised by the Routerv51inch contract.
type Routerv51inchAccessDenied struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccessDenied()
func Routerv51inchAccessDeniedErrorID() common.Hash {
	return common.HexToHash("0x4ca888678577e39e3ac9ac6cfe78050ee33f0c0eec2c686bc222b2cf4140a62c")
}

// UnpackAccessDeniedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccessDenied()
func (routerv51inch *Routerv51inch) UnpackAccessDeniedError(raw []byte) (*Routerv51inchAccessDenied, error) {
	out := new(Routerv51inchAccessDenied)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "AccessDenied", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchAdvanceNonceFailed represents a AdvanceNonceFailed error raised by the Routerv51inch contract.
type Routerv51inchAdvanceNonceFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AdvanceNonceFailed()
func Routerv51inchAdvanceNonceFailedErrorID() common.Hash {
	return common.HexToHash("0xbd71636dddb2a118a14ef51205c3cfec4113d075279ddd7e4c7c0ad6d5997216")
}

// UnpackAdvanceNonceFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AdvanceNonceFailed()
func (routerv51inch *Routerv51inch) UnpackAdvanceNonceFailedError(raw []byte) (*Routerv51inchAdvanceNonceFailed, error) {
	out := new(Routerv51inchAdvanceNonceFailed)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "AdvanceNonceFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchAlreadyFilled represents a AlreadyFilled error raised by the Routerv51inch contract.
type Routerv51inchAlreadyFilled struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyFilled()
func Routerv51inchAlreadyFilledErrorID() common.Hash {
	return common.HexToHash("0x41a26a633154f27ef57c08231424bd8d87c4557db15f2b3e790dd2b1f50cc123")
}

// UnpackAlreadyFilledError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyFilled()
func (routerv51inch *Routerv51inch) UnpackAlreadyFilledError(raw []byte) (*Routerv51inchAlreadyFilled, error) {
	out := new(Routerv51inchAlreadyFilled)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "AlreadyFilled", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchArbitraryStaticCallFailed represents a ArbitraryStaticCallFailed error raised by the Routerv51inch contract.
type Routerv51inchArbitraryStaticCallFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ArbitraryStaticCallFailed()
func Routerv51inchArbitraryStaticCallFailedErrorID() common.Hash {
	return common.HexToHash("0x1f1b8f6185f3387f7dd1c7ff816bd1dd2b3f59e21a9f1807a270349ee1bd307c")
}

// UnpackArbitraryStaticCallFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ArbitraryStaticCallFailed()
func (routerv51inch *Routerv51inch) UnpackArbitraryStaticCallFailedError(raw []byte) (*Routerv51inchArbitraryStaticCallFailed, error) {
	out := new(Routerv51inchArbitraryStaticCallFailed)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "ArbitraryStaticCallFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchBadPool represents a BadPool error raised by the Routerv51inch contract.
type Routerv51inchBadPool struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BadPool()
func Routerv51inchBadPoolErrorID() common.Hash {
	return common.HexToHash("0xb2c02722cf230da6a05b6ae0e22f42ed25be4bf9b34cb4514ebd83ff4a53308a")
}

// UnpackBadPoolError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BadPool()
func (routerv51inch *Routerv51inch) UnpackBadPoolError(raw []byte) (*Routerv51inchBadPool, error) {
	out := new(Routerv51inchBadPool)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "BadPool", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchBadSignature represents a BadSignature error raised by the Routerv51inch contract.
type Routerv51inchBadSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BadSignature()
func Routerv51inchBadSignatureErrorID() common.Hash {
	return common.HexToHash("0x5cd5d2335541c4f2ed05fbe44f397e8b79f8e2333157122d2dab06e378ef7685")
}

// UnpackBadSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BadSignature()
func (routerv51inch *Routerv51inch) UnpackBadSignatureError(raw []byte) (*Routerv51inchBadSignature, error) {
	out := new(Routerv51inchBadSignature)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "BadSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchETHTransferFailed represents a ETHTransferFailed error raised by the Routerv51inch contract.
type Routerv51inchETHTransferFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ETHTransferFailed()
func Routerv51inchETHTransferFailedErrorID() common.Hash {
	return common.HexToHash("0xb12d13ebe76e15b5fdb7bf52f0daba617b83ebcc560b0666c44fcdcd71f4362b")
}

// UnpackETHTransferFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ETHTransferFailed()
func (routerv51inch *Routerv51inch) UnpackETHTransferFailedError(raw []byte) (*Routerv51inchETHTransferFailed, error) {
	out := new(Routerv51inchETHTransferFailed)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "ETHTransferFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchEmptyPools represents a EmptyPools error raised by the Routerv51inch contract.
type Routerv51inchEmptyPools struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmptyPools()
func Routerv51inchEmptyPoolsErrorID() common.Hash {
	return common.HexToHash("0x67e7c0f67442677d81e86252d9ba68271e533da1767e0e19e8bd809ab180362f")
}

// UnpackEmptyPoolsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmptyPools()
func (routerv51inch *Routerv51inch) UnpackEmptyPoolsError(raw []byte) (*Routerv51inchEmptyPools, error) {
	out := new(Routerv51inchEmptyPools)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "EmptyPools", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchEthDepositRejected represents a EthDepositRejected error raised by the Routerv51inch contract.
type Routerv51inchEthDepositRejected struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EthDepositRejected()
func Routerv51inchEthDepositRejectedErrorID() common.Hash {
	return common.HexToHash("0x1b10b0f9ae66bdc9d7cb534137e3b02811ba15910619f3dcc6f5f5e2f8e91721")
}

// UnpackEthDepositRejectedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EthDepositRejected()
func (routerv51inch *Routerv51inch) UnpackEthDepositRejectedError(raw []byte) (*Routerv51inchEthDepositRejected, error) {
	out := new(Routerv51inchEthDepositRejected)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "EthDepositRejected", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchGetAmountCallFailed represents a GetAmountCallFailed error raised by the Routerv51inch contract.
type Routerv51inchGetAmountCallFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GetAmountCallFailed()
func Routerv51inchGetAmountCallFailedErrorID() common.Hash {
	return common.HexToHash("0x110b8e73833370cda9e049f868aaedce79f6c9b116437870e0b7c46794443778")
}

// UnpackGetAmountCallFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GetAmountCallFailed()
func (routerv51inch *Routerv51inch) UnpackGetAmountCallFailedError(raw []byte) (*Routerv51inchGetAmountCallFailed, error) {
	out := new(Routerv51inchGetAmountCallFailed)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "GetAmountCallFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchIncorrectDataLength represents a IncorrectDataLength error raised by the Routerv51inch contract.
type Routerv51inchIncorrectDataLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error IncorrectDataLength()
func Routerv51inchIncorrectDataLengthErrorID() common.Hash {
	return common.HexToHash("0xef356d7ab93249f82a1c67e73dfee93500a1c7c4329d2f22ab59bb72810a1326")
}

// UnpackIncorrectDataLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error IncorrectDataLength()
func (routerv51inch *Routerv51inch) UnpackIncorrectDataLengthError(raw []byte) (*Routerv51inchIncorrectDataLength, error) {
	out := new(Routerv51inchIncorrectDataLength)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "IncorrectDataLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchInsufficientBalance represents a InsufficientBalance error raised by the Routerv51inch contract.
type Routerv51inchInsufficientBalance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance()
func Routerv51inchInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xf4d678b8ce6b5157126b1484a53523762a93571537a7d5ae97d8014a44715c94")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance()
func (routerv51inch *Routerv51inch) UnpackInsufficientBalanceError(raw []byte) (*Routerv51inchInsufficientBalance, error) {
	out := new(Routerv51inchInsufficientBalance)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchInvalidMsgValue represents a InvalidMsgValue error raised by the Routerv51inch contract.
type Routerv51inchInvalidMsgValue struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidMsgValue()
func Routerv51inchInvalidMsgValueErrorID() common.Hash {
	return common.HexToHash("0x1841b4e1b5bc2b6ee2d929f61c1a6d695028c1b47aa99b13e72b417bfaebc3cd")
}

// UnpackInvalidMsgValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidMsgValue()
func (routerv51inch *Routerv51inch) UnpackInvalidMsgValueError(raw []byte) (*Routerv51inchInvalidMsgValue, error) {
	out := new(Routerv51inchInvalidMsgValue)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "InvalidMsgValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchInvalidatedOrder represents a InvalidatedOrder error raised by the Routerv51inch contract.
type Routerv51inchInvalidatedOrder struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidatedOrder()
func Routerv51inchInvalidatedOrderErrorID() common.Hash {
	return common.HexToHash("0xf71fbda25ab7ea5b8d7774a17207b5659639b9c5360eda3ff0782c89adfd8eeb")
}

// UnpackInvalidatedOrderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidatedOrder()
func (routerv51inch *Routerv51inch) UnpackInvalidatedOrderError(raw []byte) (*Routerv51inchInvalidatedOrder, error) {
	out := new(Routerv51inchInvalidatedOrder)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "InvalidatedOrder", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchMakingAmountExceeded represents a MakingAmountExceeded error raised by the Routerv51inch contract.
type Routerv51inchMakingAmountExceeded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MakingAmountExceeded()
func Routerv51inchMakingAmountExceededErrorID() common.Hash {
	return common.HexToHash("0xaa34b696169d5d4e9afa93ef29b31f257603f48d1af97d965867463be83438a4")
}

// UnpackMakingAmountExceededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MakingAmountExceeded()
func (routerv51inch *Routerv51inch) UnpackMakingAmountExceededError(raw []byte) (*Routerv51inchMakingAmountExceeded, error) {
	out := new(Routerv51inchMakingAmountExceeded)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "MakingAmountExceeded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchMakingAmountTooLow represents a MakingAmountTooLow error raised by the Routerv51inch contract.
type Routerv51inchMakingAmountTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MakingAmountTooLow()
func Routerv51inchMakingAmountTooLowErrorID() common.Hash {
	return common.HexToHash("0x481ea392150ef4334793b50f5628cfd068e177118e6993d7c20fd36fdfccc9c1")
}

// UnpackMakingAmountTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MakingAmountTooLow()
func (routerv51inch *Routerv51inch) UnpackMakingAmountTooLowError(raw []byte) (*Routerv51inchMakingAmountTooLow, error) {
	out := new(Routerv51inchMakingAmountTooLow)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "MakingAmountTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchOnlyOneAmountShouldBeZero represents a OnlyOneAmountShouldBeZero error raised by the Routerv51inch contract.
type Routerv51inchOnlyOneAmountShouldBeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOneAmountShouldBeZero()
func Routerv51inchOnlyOneAmountShouldBeZeroErrorID() common.Hash {
	return common.HexToHash("0x00e2a5226c33830c06ec02b98eb8b645c7dd4d809eca88f7a1fd7e0f746e0b9d")
}

// UnpackOnlyOneAmountShouldBeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOneAmountShouldBeZero()
func (routerv51inch *Routerv51inch) UnpackOnlyOneAmountShouldBeZeroError(raw []byte) (*Routerv51inchOnlyOneAmountShouldBeZero, error) {
	out := new(Routerv51inchOnlyOneAmountShouldBeZero)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "OnlyOneAmountShouldBeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchOrderExpired represents a OrderExpired error raised by the Routerv51inch contract.
type Routerv51inchOrderExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OrderExpired()
func Routerv51inchOrderExpiredErrorID() common.Hash {
	return common.HexToHash("0xc56873bac2ec2d1dd6d40a2a9b432692a2b77a4253c07d6f7ef4929f60ced89c")
}

// UnpackOrderExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OrderExpired()
func (routerv51inch *Routerv51inch) UnpackOrderExpiredError(raw []byte) (*Routerv51inchOrderExpired, error) {
	out := new(Routerv51inchOrderExpired)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "OrderExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchPermitLengthTooLow represents a PermitLengthTooLow error raised by the Routerv51inch contract.
type Routerv51inchPermitLengthTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PermitLengthTooLow()
func Routerv51inchPermitLengthTooLowErrorID() common.Hash {
	return common.HexToHash("0xd9e1c6dcf9e4601acfbad747aebca8f3749c1d143a4bbbc1682bcfb1e681629a")
}

// UnpackPermitLengthTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PermitLengthTooLow()
func (routerv51inch *Routerv51inch) UnpackPermitLengthTooLowError(raw []byte) (*Routerv51inchPermitLengthTooLow, error) {
	out := new(Routerv51inchPermitLengthTooLow)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "PermitLengthTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchPredicateIsNotTrue represents a PredicateIsNotTrue error raised by the Routerv51inch contract.
type Routerv51inchPredicateIsNotTrue struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PredicateIsNotTrue()
func Routerv51inchPredicateIsNotTrueErrorID() common.Hash {
	return common.HexToHash("0xb6629c02d7b61b434160b9a6b56d99a2aa5c959e50d9876456b23faf101fdacf")
}

// UnpackPredicateIsNotTrueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PredicateIsNotTrue()
func (routerv51inch *Routerv51inch) UnpackPredicateIsNotTrueError(raw []byte) (*Routerv51inchPredicateIsNotTrue, error) {
	out := new(Routerv51inchPredicateIsNotTrue)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "PredicateIsNotTrue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchPrivateOrder represents a PrivateOrder error raised by the Routerv51inch contract.
type Routerv51inchPrivateOrder struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PrivateOrder()
func Routerv51inchPrivateOrderErrorID() common.Hash {
	return common.HexToHash("0xd4dfdafe06416a471ed7b8bd430201111ccc397ff3d5636080df9597d0d8626c")
}

// UnpackPrivateOrderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PrivateOrder()
func (routerv51inch *Routerv51inch) UnpackPrivateOrderError(raw []byte) (*Routerv51inchPrivateOrder, error) {
	out := new(Routerv51inchPrivateOrder)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "PrivateOrder", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchRFQBadSignature represents a RFQBadSignature error raised by the Routerv51inch contract.
type Routerv51inchRFQBadSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RFQBadSignature()
func Routerv51inchRFQBadSignatureErrorID() common.Hash {
	return common.HexToHash("0x17c2b1f11b44b97e23e5551e40c8eb362edef005b1ac3c4273f80e9a34dae4f1")
}

// UnpackRFQBadSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RFQBadSignature()
func (routerv51inch *Routerv51inch) UnpackRFQBadSignatureError(raw []byte) (*Routerv51inchRFQBadSignature, error) {
	out := new(Routerv51inchRFQBadSignature)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "RFQBadSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchRFQPrivateOrder represents a RFQPrivateOrder error raised by the Routerv51inch contract.
type Routerv51inchRFQPrivateOrder struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RFQPrivateOrder()
func Routerv51inchRFQPrivateOrderErrorID() common.Hash {
	return common.HexToHash("0xe8c66321f8c8ebefab371e1717aeb615498db263ddef031cbf8d9573384a713b")
}

// UnpackRFQPrivateOrderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RFQPrivateOrder()
func (routerv51inch *Routerv51inch) UnpackRFQPrivateOrderError(raw []byte) (*Routerv51inchRFQPrivateOrder, error) {
	out := new(Routerv51inchRFQPrivateOrder)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "RFQPrivateOrder", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchRFQSwapWithZeroAmount represents a RFQSwapWithZeroAmount error raised by the Routerv51inch contract.
type Routerv51inchRFQSwapWithZeroAmount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RFQSwapWithZeroAmount()
func Routerv51inchRFQSwapWithZeroAmountErrorID() common.Hash {
	return common.HexToHash("0x07b6e79f27d20791bce85b8a2eb6aa85b6424ca10a08d5e92515cf5bb8560d42")
}

// UnpackRFQSwapWithZeroAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RFQSwapWithZeroAmount()
func (routerv51inch *Routerv51inch) UnpackRFQSwapWithZeroAmountError(raw []byte) (*Routerv51inchRFQSwapWithZeroAmount, error) {
	out := new(Routerv51inchRFQSwapWithZeroAmount)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "RFQSwapWithZeroAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchRFQZeroTargetIsForbidden represents a RFQZeroTargetIsForbidden error raised by the Routerv51inch contract.
type Routerv51inchRFQZeroTargetIsForbidden struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RFQZeroTargetIsForbidden()
func Routerv51inchRFQZeroTargetIsForbiddenErrorID() common.Hash {
	return common.HexToHash("0x692e45e060706118d0fca1024e0421af623dee6b574578be5c300bcae04029cf")
}

// UnpackRFQZeroTargetIsForbiddenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RFQZeroTargetIsForbidden()
func (routerv51inch *Routerv51inch) UnpackRFQZeroTargetIsForbiddenError(raw []byte) (*Routerv51inchRFQZeroTargetIsForbidden, error) {
	out := new(Routerv51inchRFQZeroTargetIsForbidden)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "RFQZeroTargetIsForbidden", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchReentrancyDetected represents a ReentrancyDetected error raised by the Routerv51inch contract.
type Routerv51inchReentrancyDetected struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyDetected()
func Routerv51inchReentrancyDetectedErrorID() common.Hash {
	return common.HexToHash("0xc5f2be51ec4ec0ad8a7972d497da993a6fcbb89cf72c05f97d654ed81ce53492")
}

// UnpackReentrancyDetectedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyDetected()
func (routerv51inch *Routerv51inch) UnpackReentrancyDetectedError(raw []byte) (*Routerv51inchReentrancyDetected, error) {
	out := new(Routerv51inchReentrancyDetected)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "ReentrancyDetected", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchRemainingAmountIsZero represents a RemainingAmountIsZero error raised by the Routerv51inch contract.
type Routerv51inchRemainingAmountIsZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RemainingAmountIsZero()
func Routerv51inchRemainingAmountIsZeroErrorID() common.Hash {
	return common.HexToHash("0xecef3664cb9d2bc973778b02b151f52318010a6bc6add1a0a26dd0ed6ced3cec")
}

// UnpackRemainingAmountIsZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RemainingAmountIsZero()
func (routerv51inch *Routerv51inch) UnpackRemainingAmountIsZeroError(raw []byte) (*Routerv51inchRemainingAmountIsZero, error) {
	out := new(Routerv51inchRemainingAmountIsZero)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "RemainingAmountIsZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchReservesCallFailed represents a ReservesCallFailed error raised by the Routerv51inch contract.
type Routerv51inchReservesCallFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReservesCallFailed()
func Routerv51inchReservesCallFailedErrorID() common.Hash {
	return common.HexToHash("0x85cd58dcd62237b1af0231b4a8943a369462ba5bafe0a59c9a3b7637f8e6b2ac")
}

// UnpackReservesCallFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReservesCallFailed()
func (routerv51inch *Routerv51inch) UnpackReservesCallFailedError(raw []byte) (*Routerv51inchReservesCallFailed, error) {
	out := new(Routerv51inchReservesCallFailed)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "ReservesCallFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchReturnAmountIsNotEnough represents a ReturnAmountIsNotEnough error raised by the Routerv51inch contract.
type Routerv51inchReturnAmountIsNotEnough struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReturnAmountIsNotEnough()
func Routerv51inchReturnAmountIsNotEnoughErrorID() common.Hash {
	return common.HexToHash("0xf32bec2f2b35bcc476d8e33c53406cb18350e239c4a66a9fec60187b5910cabb")
}

// UnpackReturnAmountIsNotEnoughError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReturnAmountIsNotEnough()
func (routerv51inch *Routerv51inch) UnpackReturnAmountIsNotEnoughError(raw []byte) (*Routerv51inchReturnAmountIsNotEnough, error) {
	out := new(Routerv51inchReturnAmountIsNotEnough)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "ReturnAmountIsNotEnough", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchSafePermitBadLength represents a SafePermitBadLength error raised by the Routerv51inch contract.
type Routerv51inchSafePermitBadLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafePermitBadLength()
func Routerv51inchSafePermitBadLengthErrorID() common.Hash {
	return common.HexToHash("0x682758570ff056bd9b78ea8324a2e9cfc01bd4d65b5bfcae33c2ad311fae3abd")
}

// UnpackSafePermitBadLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafePermitBadLength()
func (routerv51inch *Routerv51inch) UnpackSafePermitBadLengthError(raw []byte) (*Routerv51inchSafePermitBadLength, error) {
	out := new(Routerv51inchSafePermitBadLength)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "SafePermitBadLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchSafeTransferFailed represents a SafeTransferFailed error raised by the Routerv51inch contract.
type Routerv51inchSafeTransferFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeTransferFailed()
func Routerv51inchSafeTransferFailedErrorID() common.Hash {
	return common.HexToHash("0xfb7f50796995f43b6f601c7bdc661fff3554a3197898a14efb94c31f671088ce")
}

// UnpackSafeTransferFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeTransferFailed()
func (routerv51inch *Routerv51inch) UnpackSafeTransferFailedError(raw []byte) (*Routerv51inchSafeTransferFailed, error) {
	out := new(Routerv51inchSafeTransferFailed)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "SafeTransferFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchSafeTransferFromFailed represents a SafeTransferFromFailed error raised by the Routerv51inch contract.
type Routerv51inchSafeTransferFromFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeTransferFromFailed()
func Routerv51inchSafeTransferFromFailedErrorID() common.Hash {
	return common.HexToHash("0xf4059071351ef2c2ce34e72012fb887b861f62f289b68b0c4e171c82a6b2023d")
}

// UnpackSafeTransferFromFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeTransferFromFailed()
func (routerv51inch *Routerv51inch) UnpackSafeTransferFromFailedError(raw []byte) (*Routerv51inchSafeTransferFromFailed, error) {
	out := new(Routerv51inchSafeTransferFromFailed)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "SafeTransferFromFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchSimulationResults represents a SimulationResults error raised by the Routerv51inch contract.
type Routerv51inchSimulationResults struct {
	Success bool
	Res     []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SimulationResults(bool success, bytes res)
func Routerv51inchSimulationResultsErrorID() common.Hash {
	return common.HexToHash("0x1934afc82dd770c05787a7a944999e73cb2b7d95aae2db5cf2dd220cffb39cf7")
}

// UnpackSimulationResultsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SimulationResults(bool success, bytes res)
func (routerv51inch *Routerv51inch) UnpackSimulationResultsError(raw []byte) (*Routerv51inchSimulationResults, error) {
	out := new(Routerv51inchSimulationResults)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "SimulationResults", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchSwapAmountTooLarge represents a SwapAmountTooLarge error raised by the Routerv51inch contract.
type Routerv51inchSwapAmountTooLarge struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SwapAmountTooLarge()
func Routerv51inchSwapAmountTooLargeErrorID() common.Hash {
	return common.HexToHash("0xcf0b4d3a884e114cb306f42fb35bbbe89e97b367986376e531cdaff45419959c")
}

// UnpackSwapAmountTooLargeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SwapAmountTooLarge()
func (routerv51inch *Routerv51inch) UnpackSwapAmountTooLargeError(raw []byte) (*Routerv51inchSwapAmountTooLarge, error) {
	out := new(Routerv51inchSwapAmountTooLarge)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "SwapAmountTooLarge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchSwapWithZeroAmount represents a SwapWithZeroAmount error raised by the Routerv51inch contract.
type Routerv51inchSwapWithZeroAmount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SwapWithZeroAmount()
func Routerv51inchSwapWithZeroAmountErrorID() common.Hash {
	return common.HexToHash("0xfba5a276a6a0bc3e7a381b7cdc29e9d356a2baf3ab47924f586f2e1eb54db7f9")
}

// UnpackSwapWithZeroAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SwapWithZeroAmount()
func (routerv51inch *Routerv51inch) UnpackSwapWithZeroAmountError(raw []byte) (*Routerv51inchSwapWithZeroAmount, error) {
	out := new(Routerv51inchSwapWithZeroAmount)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "SwapWithZeroAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchTakingAmountExceeded represents a TakingAmountExceeded error raised by the Routerv51inch contract.
type Routerv51inchTakingAmountExceeded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TakingAmountExceeded()
func Routerv51inchTakingAmountExceededErrorID() common.Hash {
	return common.HexToHash("0x7f902a93c23e1e7f62cc5268147eaca1f0798d33dd25f87e80c63d453e0b5913")
}

// UnpackTakingAmountExceededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TakingAmountExceeded()
func (routerv51inch *Routerv51inch) UnpackTakingAmountExceededError(raw []byte) (*Routerv51inchTakingAmountExceeded, error) {
	out := new(Routerv51inchTakingAmountExceeded)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "TakingAmountExceeded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchTakingAmountIncreased represents a TakingAmountIncreased error raised by the Routerv51inch contract.
type Routerv51inchTakingAmountIncreased struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TakingAmountIncreased()
func Routerv51inchTakingAmountIncreasedErrorID() common.Hash {
	return common.HexToHash("0x939c4204c9cd28b71b02327fd6a2bb0aaf55023e1b918084e2fa50bc26a04d77")
}

// UnpackTakingAmountIncreasedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TakingAmountIncreased()
func (routerv51inch *Routerv51inch) UnpackTakingAmountIncreasedError(raw []byte) (*Routerv51inchTakingAmountIncreased, error) {
	out := new(Routerv51inchTakingAmountIncreased)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "TakingAmountIncreased", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchTakingAmountTooHigh represents a TakingAmountTooHigh error raised by the Routerv51inch contract.
type Routerv51inchTakingAmountTooHigh struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TakingAmountTooHigh()
func Routerv51inchTakingAmountTooHighErrorID() common.Hash {
	return common.HexToHash("0xfb8ae129bb15dd7d41422f958c6ce0c12a8a6756e86abf2d1f1edaa920524a69")
}

// UnpackTakingAmountTooHighError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TakingAmountTooHigh()
func (routerv51inch *Routerv51inch) UnpackTakingAmountTooHighError(raw []byte) (*Routerv51inchTakingAmountTooHigh, error) {
	out := new(Routerv51inchTakingAmountTooHigh)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "TakingAmountTooHigh", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchTransferFromMakerToTakerFailed represents a TransferFromMakerToTakerFailed error raised by the Routerv51inch contract.
type Routerv51inchTransferFromMakerToTakerFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferFromMakerToTakerFailed()
func Routerv51inchTransferFromMakerToTakerFailedErrorID() common.Hash {
	return common.HexToHash("0x70a03f480a56f2ba97a5ae95726db5c4aa88d6b043211db6f6841f90c71390ac")
}

// UnpackTransferFromMakerToTakerFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferFromMakerToTakerFailed()
func (routerv51inch *Routerv51inch) UnpackTransferFromMakerToTakerFailedError(raw []byte) (*Routerv51inchTransferFromMakerToTakerFailed, error) {
	out := new(Routerv51inchTransferFromMakerToTakerFailed)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "TransferFromMakerToTakerFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchTransferFromTakerToMakerFailed represents a TransferFromTakerToMakerFailed error raised by the Routerv51inch contract.
type Routerv51inchTransferFromTakerToMakerFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferFromTakerToMakerFailed()
func Routerv51inchTransferFromTakerToMakerFailedErrorID() common.Hash {
	return common.HexToHash("0x478a5205e5ee1fb19e08efddef0ae39549de8d23a39befdc0b65f352d124ffd3")
}

// UnpackTransferFromTakerToMakerFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferFromTakerToMakerFailed()
func (routerv51inch *Routerv51inch) UnpackTransferFromTakerToMakerFailedError(raw []byte) (*Routerv51inchTransferFromTakerToMakerFailed, error) {
	out := new(Routerv51inchTransferFromTakerToMakerFailed)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "TransferFromTakerToMakerFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchUnknownOrder represents a UnknownOrder error raised by the Routerv51inch contract.
type Routerv51inchUnknownOrder struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnknownOrder()
func Routerv51inchUnknownOrderErrorID() common.Hash {
	return common.HexToHash("0xb838de9615eead93e8f6f4f2b287f3edc63a6b81847f55d5b28ba54e8d10b7bc")
}

// UnpackUnknownOrderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnknownOrder()
func (routerv51inch *Routerv51inch) UnpackUnknownOrderError(raw []byte) (*Routerv51inchUnknownOrder, error) {
	out := new(Routerv51inchUnknownOrder)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "UnknownOrder", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchWrongAmount represents a WrongAmount error raised by the Routerv51inch contract.
type Routerv51inchWrongAmount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WrongAmount()
func Routerv51inchWrongAmountErrorID() common.Hash {
	return common.HexToHash("0x49986e73bdd0fbe5e1cceb3af2382adb7426855b5defd5cf6a2531f6c9e41ced")
}

// UnpackWrongAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WrongAmount()
func (routerv51inch *Routerv51inch) UnpackWrongAmountError(raw []byte) (*Routerv51inchWrongAmount, error) {
	out := new(Routerv51inchWrongAmount)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "WrongAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchWrongGetter represents a WrongGetter error raised by the Routerv51inch contract.
type Routerv51inchWrongGetter struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WrongGetter()
func Routerv51inchWrongGetterErrorID() common.Hash {
	return common.HexToHash("0xbec74c85987581ab4fbe0a96f7e74947d6dddde03d6a23964b3f44535efb6c6e")
}

// UnpackWrongGetterError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WrongGetter()
func (routerv51inch *Routerv51inch) UnpackWrongGetterError(raw []byte) (*Routerv51inchWrongGetter, error) {
	out := new(Routerv51inchWrongGetter)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "WrongGetter", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchZeroAddress represents a ZeroAddress error raised by the Routerv51inch contract.
type Routerv51inchZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroAddress()
func Routerv51inchZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xd92e233df2717d4a40030e20904abd27b68fcbeede117eaaccbbdac9618c8c73")
}

// UnpackZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroAddress()
func (routerv51inch *Routerv51inch) UnpackZeroAddressError(raw []byte) (*Routerv51inchZeroAddress, error) {
	out := new(Routerv51inchZeroAddress)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "ZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchZeroMinReturn represents a ZeroMinReturn error raised by the Routerv51inch contract.
type Routerv51inchZeroMinReturn struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroMinReturn()
func Routerv51inchZeroMinReturnErrorID() common.Hash {
	return common.HexToHash("0x0262dde406f45de042c06c7df2790374a9d79ccf79e8ab35852b326859e71236")
}

// UnpackZeroMinReturnError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroMinReturn()
func (routerv51inch *Routerv51inch) UnpackZeroMinReturnError(raw []byte) (*Routerv51inchZeroMinReturn, error) {
	out := new(Routerv51inchZeroMinReturn)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "ZeroMinReturn", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchZeroReturnAmount represents a ZeroReturnAmount error raised by the Routerv51inch contract.
type Routerv51inchZeroReturnAmount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroReturnAmount()
func Routerv51inchZeroReturnAmountErrorID() common.Hash {
	return common.HexToHash("0x28ebf24706d62ddb5e44af4b07a60fd49f9f979ea8fe3d9ce63597ba678e0cf8")
}

// UnpackZeroReturnAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroReturnAmount()
func (routerv51inch *Routerv51inch) UnpackZeroReturnAmountError(raw []byte) (*Routerv51inchZeroReturnAmount, error) {
	out := new(Routerv51inchZeroReturnAmount)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "ZeroReturnAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv51inchZeroTargetIsForbidden represents a ZeroTargetIsForbidden error raised by the Routerv51inch contract.
type Routerv51inchZeroTargetIsForbidden struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroTargetIsForbidden()
func Routerv51inchZeroTargetIsForbiddenErrorID() common.Hash {
	return common.HexToHash("0xb0c4d05f88cb3cf5b9ac58bde5c7946fada07a6c8c763015c89e868e19fc6dce")
}

// UnpackZeroTargetIsForbiddenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroTargetIsForbidden()
func (routerv51inch *Routerv51inch) UnpackZeroTargetIsForbiddenError(raw []byte) (*Routerv51inchZeroTargetIsForbidden, error) {
	out := new(Routerv51inchZeroTargetIsForbidden)
	if err := routerv51inch.abi.UnpackIntoInterface(out, "ZeroTargetIsForbidden", raw); err != nil {
		return nil, err
	}
	return out, nil
}
