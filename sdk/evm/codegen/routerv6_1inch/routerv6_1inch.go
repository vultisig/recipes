// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package routerv6_1inch

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

// IOrderMixinOrder is an auto generated low-level Go binding around an user-defined struct.
type IOrderMixinOrder struct {
	Salt         *big.Int
	Maker        *big.Int
	Receiver     *big.Int
	MakerAsset   *big.Int
	TakerAsset   *big.Int
	MakingAmount *big.Int
	TakingAmount *big.Int
	MakerTraits  *big.Int
}

// Routerv61inchMetaData contains all meta data concerning the Routerv61inch contract.
var Routerv61inchMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdvanceEpochFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArbitraryStaticCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadCurveSwapSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BitInvalidatedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EpochManagerAndBitInvalidatorsAreIncompatible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthDepositRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPermit2Transfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidatedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchArraysLengths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderIsNotSuitableForMassInvalidation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Permit2TransferAmountTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PredicateIsNotTrue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrivateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyDetected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemainingInvalidatedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReservesCallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"}],\"name\":\"ReturnAmountIsNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFromFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"name\":\"SimulationResults\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapWithZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromMakerToTakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromTakerToMakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongSeriesNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMinReturn\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotValue\",\"type\":\"uint256\"}],\"name\":\"BitInvalidatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"series\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newEpoch\",\"type\":\"uint256\"}],\"name\":\"EpochIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"series\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"advanceEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"and\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"arbitraryStaticCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"bitInvalidatorForOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalMask\",\"type\":\"uint256\"}],\"name\":\"bitsInvalidateForOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"MakerTraits[]\",\"name\":\"makerTraits\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"predicate\",\"type\":\"bytes\"}],\"name\":\"checkPredicate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchange\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"Address\",\"name\":\"srcToken\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchange\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"Address\",\"name\":\"srcToken\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inCoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"curveSwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"series\",\"type\":\"uint96\"}],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"series\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerEpoch\",\"type\":\"uint256\"}],\"name\":\"epochEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"eq\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"}],\"name\":\"ethUnoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"}],\"name\":\"ethUnoswap2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex3\",\"type\":\"uint256\"}],\"name\":\"ethUnoswap3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"}],\"name\":\"ethUnoswapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"}],\"name\":\"ethUnoswapTo2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex3\",\"type\":\"uint256\"}],\"name\":\"ethUnoswapTo3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"receiver\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderMixin.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"TakerTraits\",\"name\":\"takerTraits\",\"type\":\"uint256\"}],\"name\":\"fillContractOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"receiver\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderMixin.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"TakerTraits\",\"name\":\"takerTraits\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"fillContractOrderArgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"receiver\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderMixin.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"TakerTraits\",\"name\":\"takerTraits\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"receiver\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderMixin.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"TakerTraits\",\"name\":\"takerTraits\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"fillOrderArgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"gt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"maker\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"receiver\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"makerAsset\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"takerAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"MakerTraits\",\"name\":\"makerTraits\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderMixin.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"series\",\"type\":\"uint96\"}],\"name\":\"increaseEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"lt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"not\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"or\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"action\",\"type\":\"bytes\"}],\"name\":\"permitAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"rawRemainingInvalidatorForOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remainingInvalidatorForOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"executor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structGenericRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"}],\"name\":\"unoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"}],\"name\":\"unoswap2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex3\",\"type\":\"uint256\"}],\"name\":\"unoswap3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"}],\"name\":\"unoswapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"}],\"name\":\"unoswapTo2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Address\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex2\",\"type\":\"uint256\"},{\"internalType\":\"Address\",\"name\":\"dex3\",\"type\":\"uint256\"}],\"name\":\"unoswapTo3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "Routerv61inch",
}

// Routerv61inch is an auto generated Go binding around an Ethereum contract.
type Routerv61inch struct {
	abi abi.ABI
}

// NewRouterv61inch creates a new instance of Routerv61inch.
func NewRouterv61inch() *Routerv61inch {
	parsed, err := Routerv61inchMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Routerv61inch{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Routerv61inch) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address weth) returns()
func (routerv61inch *Routerv61inch) PackConstructor(weth common.Address) []byte {
	enc, err := routerv61inch.abi.Pack("", weth)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAdvanceEpoch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d2c7c16.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function advanceEpoch(uint96 series, uint256 amount) returns()
func (routerv61inch *Routerv61inch) PackAdvanceEpoch(series *big.Int, amount *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("advanceEpoch", series, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAdvanceEpoch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d2c7c16.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function advanceEpoch(uint96 series, uint256 amount) returns()
func (routerv61inch *Routerv61inch) TryPackAdvanceEpoch(series *big.Int, amount *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("advanceEpoch", series, amount)
}

// PackAnd is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbfa75143.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (routerv61inch *Routerv61inch) PackAnd(offsets *big.Int, data []byte) []byte {
	enc, err := routerv61inch.abi.Pack("and", offsets, data)
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
func (routerv61inch *Routerv61inch) TryPackAnd(offsets *big.Int, data []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("and", offsets, data)
}

// UnpackAnd is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (routerv61inch *Routerv61inch) UnpackAnd(data []byte) (bool, error) {
	out, err := routerv61inch.abi.Unpack("and", data)
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
func (routerv61inch *Routerv61inch) PackArbitraryStaticCall(target common.Address, data []byte) []byte {
	enc, err := routerv61inch.abi.Pack("arbitraryStaticCall", target, data)
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
func (routerv61inch *Routerv61inch) TryPackArbitraryStaticCall(target common.Address, data []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("arbitraryStaticCall", target, data)
}

// UnpackArbitraryStaticCall is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (routerv61inch *Routerv61inch) UnpackArbitraryStaticCall(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("arbitraryStaticCall", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBitInvalidatorForOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x143e86a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bitInvalidatorForOrder(address maker, uint256 slot) view returns(uint256)
func (routerv61inch *Routerv61inch) PackBitInvalidatorForOrder(maker common.Address, slot *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("bitInvalidatorForOrder", maker, slot)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBitInvalidatorForOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x143e86a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bitInvalidatorForOrder(address maker, uint256 slot) view returns(uint256)
func (routerv61inch *Routerv61inch) TryPackBitInvalidatorForOrder(maker common.Address, slot *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("bitInvalidatorForOrder", maker, slot)
}

// UnpackBitInvalidatorForOrder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x143e86a7.
//
// Solidity: function bitInvalidatorForOrder(address maker, uint256 slot) view returns(uint256)
func (routerv61inch *Routerv61inch) UnpackBitInvalidatorForOrder(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("bitInvalidatorForOrder", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBitsInvalidateForOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x05b1ea03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bitsInvalidateForOrder(uint256 makerTraits, uint256 additionalMask) returns()
func (routerv61inch *Routerv61inch) PackBitsInvalidateForOrder(makerTraits *big.Int, additionalMask *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("bitsInvalidateForOrder", makerTraits, additionalMask)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBitsInvalidateForOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x05b1ea03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bitsInvalidateForOrder(uint256 makerTraits, uint256 additionalMask) returns()
func (routerv61inch *Routerv61inch) TryPackBitsInvalidateForOrder(makerTraits *big.Int, additionalMask *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("bitsInvalidateForOrder", makerTraits, additionalMask)
}

// PackCancelOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb68fb020.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelOrder(uint256 makerTraits, bytes32 orderHash) returns()
func (routerv61inch *Routerv61inch) PackCancelOrder(makerTraits *big.Int, orderHash [32]byte) []byte {
	enc, err := routerv61inch.abi.Pack("cancelOrder", makerTraits, orderHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb68fb020.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelOrder(uint256 makerTraits, bytes32 orderHash) returns()
func (routerv61inch *Routerv61inch) TryPackCancelOrder(makerTraits *big.Int, orderHash [32]byte) ([]byte, error) {
	return routerv61inch.abi.Pack("cancelOrder", makerTraits, orderHash)
}

// PackCancelOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x89e7c650.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelOrders(uint256[] makerTraits, bytes32[] orderHashes) returns()
func (routerv61inch *Routerv61inch) PackCancelOrders(makerTraits []*big.Int, orderHashes [][32]byte) []byte {
	enc, err := routerv61inch.abi.Pack("cancelOrders", makerTraits, orderHashes)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelOrders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x89e7c650.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelOrders(uint256[] makerTraits, bytes32[] orderHashes) returns()
func (routerv61inch *Routerv61inch) TryPackCancelOrders(makerTraits []*big.Int, orderHashes [][32]byte) ([]byte, error) {
	return routerv61inch.abi.Pack("cancelOrders", makerTraits, orderHashes)
}

// PackCheckPredicate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x15169dec.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function checkPredicate(bytes predicate) view returns(bool)
func (routerv61inch *Routerv61inch) PackCheckPredicate(predicate []byte) []byte {
	enc, err := routerv61inch.abi.Pack("checkPredicate", predicate)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCheckPredicate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x15169dec.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function checkPredicate(bytes predicate) view returns(bool)
func (routerv61inch *Routerv61inch) TryPackCheckPredicate(predicate []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("checkPredicate", predicate)
}

// UnpackCheckPredicate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x15169dec.
//
// Solidity: function checkPredicate(bytes predicate) view returns(bool)
func (routerv61inch *Routerv61inch) UnpackCheckPredicate(data []byte) (bool, error) {
	out, err := routerv61inch.abi.Unpack("checkPredicate", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackClipperSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2d374e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clipperSwap(address clipperExchange, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackClipperSwap(clipperExchange common.Address, srcToken *big.Int, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) []byte {
	enc, err := routerv61inch.abi.Pack("clipperSwap", clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClipperSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2d374e5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clipperSwap(address clipperExchange, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackClipperSwap(clipperExchange common.Address, srcToken *big.Int, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) ([]byte, error) {
	return routerv61inch.abi.Pack("clipperSwap", clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// UnpackClipperSwap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd2d374e5.
//
// Solidity: function clipperSwap(address clipperExchange, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackClipperSwap(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("clipperSwap", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackClipperSwapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d652af.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken *big.Int, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) []byte {
	enc, err := routerv61inch.abi.Pack("clipperSwapTo", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClipperSwapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d652af.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken *big.Int, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) ([]byte, error) {
	return routerv61inch.abi.Pack("clipperSwapTo", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// UnpackClipperSwapTo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc4d652af.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, uint256 srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackClipperSwapTo(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("clipperSwapTo", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCurveSwapCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe413f48d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function curveSwapCallback(address , address , address inCoin, uint256 dx, uint256 ) returns()
func (routerv61inch *Routerv61inch) PackCurveSwapCallback(arg0 common.Address, arg1 common.Address, inCoin common.Address, dx *big.Int, arg4 *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("curveSwapCallback", arg0, arg1, inCoin, dx, arg4)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCurveSwapCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe413f48d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function curveSwapCallback(address , address , address inCoin, uint256 dx, uint256 ) returns()
func (routerv61inch *Routerv61inch) TryPackCurveSwapCallback(arg0 common.Address, arg1 common.Address, inCoin common.Address, dx *big.Int, arg4 *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("curveSwapCallback", arg0, arg1, inCoin, dx, arg4)
}

// PackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (routerv61inch *Routerv61inch) PackEip712Domain() []byte {
	enc, err := routerv61inch.abi.Pack("eip712Domain")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEip712Domain is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84b0196e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (routerv61inch *Routerv61inch) TryPackEip712Domain() ([]byte, error) {
	return routerv61inch.abi.Pack("eip712Domain")
}

// Eip712DomainOutput serves as a container for the return parameters of contract
// method Eip712Domain.
type Eip712DomainOutput struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}

// UnpackEip712Domain is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (routerv61inch *Routerv61inch) UnpackEip712Domain(data []byte) (Eip712DomainOutput, error) {
	out, err := routerv61inch.abi.Unpack("eip712Domain", data)
	outstruct := new(Eip712DomainOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	return *outstruct, nil
}

// PackEpoch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfcea9e4e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function epoch(address maker, uint96 series) view returns(uint256)
func (routerv61inch *Routerv61inch) PackEpoch(maker common.Address, series *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("epoch", maker, series)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEpoch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfcea9e4e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function epoch(address maker, uint96 series) view returns(uint256)
func (routerv61inch *Routerv61inch) TryPackEpoch(maker common.Address, series *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("epoch", maker, series)
}

// UnpackEpoch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfcea9e4e.
//
// Solidity: function epoch(address maker, uint96 series) view returns(uint256)
func (routerv61inch *Routerv61inch) UnpackEpoch(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("epoch", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackEpochEquals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce3d710a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function epochEquals(address maker, uint256 series, uint256 makerEpoch) view returns(bool)
func (routerv61inch *Routerv61inch) PackEpochEquals(maker common.Address, series *big.Int, makerEpoch *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("epochEquals", maker, series, makerEpoch)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEpochEquals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce3d710a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function epochEquals(address maker, uint256 series, uint256 makerEpoch) view returns(bool)
func (routerv61inch *Routerv61inch) TryPackEpochEquals(maker common.Address, series *big.Int, makerEpoch *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("epochEquals", maker, series, makerEpoch)
}

// UnpackEpochEquals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xce3d710a.
//
// Solidity: function epochEquals(address maker, uint256 series, uint256 makerEpoch) view returns(bool)
func (routerv61inch *Routerv61inch) UnpackEpochEquals(data []byte) (bool, error) {
	out, err := routerv61inch.abi.Unpack("epochEquals", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackEq is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6fe7b0ba.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (routerv61inch *Routerv61inch) PackEq(value *big.Int, data []byte) []byte {
	enc, err := routerv61inch.abi.Pack("eq", value, data)
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
func (routerv61inch *Routerv61inch) TryPackEq(value *big.Int, data []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("eq", value, data)
}

// UnpackEq is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (routerv61inch *Routerv61inch) UnpackEq(data []byte) (bool, error) {
	out, err := routerv61inch.abi.Unpack("eq", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackEthUnoswap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa76dfc3b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ethUnoswap(uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackEthUnoswap(minReturn *big.Int, dex *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("ethUnoswap", minReturn, dex)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEthUnoswap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa76dfc3b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ethUnoswap(uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackEthUnoswap(minReturn *big.Int, dex *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("ethUnoswap", minReturn, dex)
}

// UnpackEthUnoswap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa76dfc3b.
//
// Solidity: function ethUnoswap(uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackEthUnoswap(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("ethUnoswap", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackEthUnoswap2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x89af926a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ethUnoswap2(uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackEthUnoswap2(minReturn *big.Int, dex *big.Int, dex2 *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("ethUnoswap2", minReturn, dex, dex2)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEthUnoswap2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x89af926a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ethUnoswap2(uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackEthUnoswap2(minReturn *big.Int, dex *big.Int, dex2 *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("ethUnoswap2", minReturn, dex, dex2)
}

// UnpackEthUnoswap2 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x89af926a.
//
// Solidity: function ethUnoswap2(uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackEthUnoswap2(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("ethUnoswap2", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackEthUnoswap3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x188ac35d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ethUnoswap3(uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackEthUnoswap3(minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("ethUnoswap3", minReturn, dex, dex2, dex3)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEthUnoswap3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x188ac35d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ethUnoswap3(uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackEthUnoswap3(minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("ethUnoswap3", minReturn, dex, dex2, dex3)
}

// UnpackEthUnoswap3 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x188ac35d.
//
// Solidity: function ethUnoswap3(uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackEthUnoswap3(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("ethUnoswap3", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackEthUnoswapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x175accdc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ethUnoswapTo(uint256 to, uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackEthUnoswapTo(to *big.Int, minReturn *big.Int, dex *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("ethUnoswapTo", to, minReturn, dex)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEthUnoswapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x175accdc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ethUnoswapTo(uint256 to, uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackEthUnoswapTo(to *big.Int, minReturn *big.Int, dex *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("ethUnoswapTo", to, minReturn, dex)
}

// UnpackEthUnoswapTo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x175accdc.
//
// Solidity: function ethUnoswapTo(uint256 to, uint256 minReturn, uint256 dex) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackEthUnoswapTo(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("ethUnoswapTo", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackEthUnoswapTo2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f449d71.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ethUnoswapTo2(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackEthUnoswapTo2(to *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("ethUnoswapTo2", to, minReturn, dex, dex2)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEthUnoswapTo2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f449d71.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ethUnoswapTo2(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackEthUnoswapTo2(to *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("ethUnoswapTo2", to, minReturn, dex, dex2)
}

// UnpackEthUnoswapTo2 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0f449d71.
//
// Solidity: function ethUnoswapTo2(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackEthUnoswapTo2(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("ethUnoswapTo2", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackEthUnoswapTo3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x493189f0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ethUnoswapTo3(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackEthUnoswapTo3(to *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("ethUnoswapTo3", to, minReturn, dex, dex2, dex3)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEthUnoswapTo3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x493189f0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ethUnoswapTo3(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackEthUnoswapTo3(to *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("ethUnoswapTo3", to, minReturn, dex, dex2, dex3)
}

// UnpackEthUnoswapTo3 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x493189f0.
//
// Solidity: function ethUnoswapTo3(uint256 to, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) payable returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackEthUnoswapTo3(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("ethUnoswapTo3", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackFillContractOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcc713a04.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillContractOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits) returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) PackFillContractOrder(order IOrderMixinOrder, signature []byte, amount *big.Int, takerTraits *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("fillContractOrder", order, signature, amount, takerTraits)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillContractOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcc713a04.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillContractOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits) returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) TryPackFillContractOrder(order IOrderMixinOrder, signature []byte, amount *big.Int, takerTraits *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("fillContractOrder", order, signature, amount, takerTraits)
}

// FillContractOrderOutput serves as a container for the return parameters of contract
// method FillContractOrder.
type FillContractOrderOutput struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 [32]byte
}

// UnpackFillContractOrder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcc713a04.
//
// Solidity: function fillContractOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits) returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) UnpackFillContractOrder(data []byte) (FillContractOrderOutput, error) {
	out, err := routerv61inch.abi.Unpack("fillContractOrder", data)
	outstruct := new(FillContractOrderOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Arg2 = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackFillContractOrderArgs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x56a75868.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillContractOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits, bytes args) returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) PackFillContractOrderArgs(order IOrderMixinOrder, signature []byte, amount *big.Int, takerTraits *big.Int, args []byte) []byte {
	enc, err := routerv61inch.abi.Pack("fillContractOrderArgs", order, signature, amount, takerTraits, args)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillContractOrderArgs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x56a75868.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillContractOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits, bytes args) returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) TryPackFillContractOrderArgs(order IOrderMixinOrder, signature []byte, amount *big.Int, takerTraits *big.Int, args []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("fillContractOrderArgs", order, signature, amount, takerTraits, args)
}

// FillContractOrderArgsOutput serves as a container for the return parameters of contract
// method FillContractOrderArgs.
type FillContractOrderArgsOutput struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 [32]byte
}

// UnpackFillContractOrderArgs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x56a75868.
//
// Solidity: function fillContractOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes signature, uint256 amount, uint256 takerTraits, bytes args) returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) UnpackFillContractOrderArgs(data []byte) (FillContractOrderArgsOutput, error) {
	out, err := routerv61inch.abi.Unpack("fillContractOrderArgs", data)
	outstruct := new(FillContractOrderArgsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Arg2 = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackFillOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fda64bd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits) payable returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) PackFillOrder(order IOrderMixinOrder, r [32]byte, vs [32]byte, amount *big.Int, takerTraits *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("fillOrder", order, r, vs, amount, takerTraits)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fda64bd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits) payable returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) TryPackFillOrder(order IOrderMixinOrder, r [32]byte, vs [32]byte, amount *big.Int, takerTraits *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("fillOrder", order, r, vs, amount, takerTraits)
}

// FillOrderOutput serves as a container for the return parameters of contract
// method FillOrder.
type FillOrderOutput struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 [32]byte
}

// UnpackFillOrder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9fda64bd.
//
// Solidity: function fillOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits) payable returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) UnpackFillOrder(data []byte) (FillOrderOutput, error) {
	out, err := routerv61inch.abi.Unpack("fillOrder", data)
	outstruct := new(FillOrderOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Arg2 = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	return *outstruct, nil
}

// PackFillOrderArgs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf497df75.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fillOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits, bytes args) payable returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) PackFillOrderArgs(order IOrderMixinOrder, r [32]byte, vs [32]byte, amount *big.Int, takerTraits *big.Int, args []byte) []byte {
	enc, err := routerv61inch.abi.Pack("fillOrderArgs", order, r, vs, amount, takerTraits, args)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFillOrderArgs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf497df75.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fillOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits, bytes args) payable returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) TryPackFillOrderArgs(order IOrderMixinOrder, r [32]byte, vs [32]byte, amount *big.Int, takerTraits *big.Int, args []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("fillOrderArgs", order, r, vs, amount, takerTraits, args)
}

// FillOrderArgsOutput serves as a container for the return parameters of contract
// method FillOrderArgs.
type FillOrderArgsOutput struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 [32]byte
}

// UnpackFillOrderArgs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf497df75.
//
// Solidity: function fillOrderArgs((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 amount, uint256 takerTraits, bytes args) payable returns(uint256, uint256, bytes32)
func (routerv61inch *Routerv61inch) UnpackFillOrderArgs(data []byte) (FillOrderArgsOutput, error) {
	out, err := routerv61inch.abi.Unpack("fillOrderArgs", data)
	outstruct := new(FillOrderArgsOutput)
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
func (routerv61inch *Routerv61inch) PackGt(value *big.Int, data []byte) []byte {
	enc, err := routerv61inch.abi.Pack("gt", value, data)
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
func (routerv61inch *Routerv61inch) TryPackGt(value *big.Int, data []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("gt", value, data)
}

// UnpackGt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (routerv61inch *Routerv61inch) UnpackGt(data []byte) (bool, error) {
	out, err := routerv61inch.abi.Unpack("gt", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackHashOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x802b2ef1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function hashOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order) view returns(bytes32)
func (routerv61inch *Routerv61inch) PackHashOrder(order IOrderMixinOrder) []byte {
	enc, err := routerv61inch.abi.Pack("hashOrder", order)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackHashOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x802b2ef1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function hashOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order) view returns(bytes32)
func (routerv61inch *Routerv61inch) TryPackHashOrder(order IOrderMixinOrder) ([]byte, error) {
	return routerv61inch.abi.Pack("hashOrder", order)
}

// UnpackHashOrder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x802b2ef1.
//
// Solidity: function hashOrder((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) order) view returns(bytes32)
func (routerv61inch *Routerv61inch) UnpackHashOrder(data []byte) ([32]byte, error) {
	out, err := routerv61inch.abi.Unpack("hashOrder", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackIncreaseEpoch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3cf8043.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function increaseEpoch(uint96 series) returns()
func (routerv61inch *Routerv61inch) PackIncreaseEpoch(series *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("increaseEpoch", series)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIncreaseEpoch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc3cf8043.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function increaseEpoch(uint96 series) returns()
func (routerv61inch *Routerv61inch) TryPackIncreaseEpoch(series *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("increaseEpoch", series)
}

// PackLt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca4ece22.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (routerv61inch *Routerv61inch) PackLt(value *big.Int, data []byte) []byte {
	enc, err := routerv61inch.abi.Pack("lt", value, data)
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
func (routerv61inch *Routerv61inch) TryPackLt(value *big.Int, data []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("lt", value, data)
}

// UnpackLt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (routerv61inch *Routerv61inch) UnpackLt(data []byte) (bool, error) {
	out, err := routerv61inch.abi.Unpack("lt", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackNot is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf797959.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function not(bytes data) view returns(bool)
func (routerv61inch *Routerv61inch) PackNot(data []byte) []byte {
	enc, err := routerv61inch.abi.Pack("not", data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNot is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf797959.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function not(bytes data) view returns(bool)
func (routerv61inch *Routerv61inch) TryPackNot(data []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("not", data)
}

// UnpackNot is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf797959.
//
// Solidity: function not(bytes data) view returns(bool)
func (routerv61inch *Routerv61inch) UnpackNot(data []byte) (bool, error) {
	out, err := routerv61inch.abi.Unpack("not", data)
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
func (routerv61inch *Routerv61inch) PackOr(offsets *big.Int, data []byte) []byte {
	enc, err := routerv61inch.abi.Pack("or", offsets, data)
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
func (routerv61inch *Routerv61inch) TryPackOr(offsets *big.Int, data []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("or", offsets, data)
}

// UnpackOr is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (routerv61inch *Routerv61inch) UnpackOr(data []byte) (bool, error) {
	out, err := routerv61inch.abi.Unpack("or", data)
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
func (routerv61inch *Routerv61inch) PackOwner() []byte {
	enc, err := routerv61inch.abi.Pack("owner")
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
func (routerv61inch *Routerv61inch) TryPackOwner() ([]byte, error) {
	return routerv61inch.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (routerv61inch *Routerv61inch) UnpackOwner(data []byte) (common.Address, error) {
	out, err := routerv61inch.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pause() returns()
func (routerv61inch *Routerv61inch) PackPause() []byte {
	enc, err := routerv61inch.abi.Pack("pause")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8456cb59.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pause() returns()
func (routerv61inch *Routerv61inch) TryPackPause() ([]byte, error) {
	return routerv61inch.abi.Pack("pause")
}

// PackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function paused() view returns(bool)
func (routerv61inch *Routerv61inch) PackPaused() []byte {
	enc, err := routerv61inch.abi.Pack("paused")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function paused() view returns(bool)
func (routerv61inch *Routerv61inch) TryPackPaused() ([]byte, error) {
	return routerv61inch.abi.Pack("paused")
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (routerv61inch *Routerv61inch) UnpackPaused(data []byte) (bool, error) {
	out, err := routerv61inch.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPermitAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5816d723.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function permitAndCall(bytes permit, bytes action) payable returns()
func (routerv61inch *Routerv61inch) PackPermitAndCall(permit []byte, action []byte) []byte {
	enc, err := routerv61inch.abi.Pack("permitAndCall", permit, action)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPermitAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5816d723.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function permitAndCall(bytes permit, bytes action) payable returns()
func (routerv61inch *Routerv61inch) TryPackPermitAndCall(permit []byte, action []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("permitAndCall", permit, action)
}

// PackRawRemainingInvalidatorForOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc2a40753.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rawRemainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (routerv61inch *Routerv61inch) PackRawRemainingInvalidatorForOrder(maker common.Address, orderHash [32]byte) []byte {
	enc, err := routerv61inch.abi.Pack("rawRemainingInvalidatorForOrder", maker, orderHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRawRemainingInvalidatorForOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc2a40753.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rawRemainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (routerv61inch *Routerv61inch) TryPackRawRemainingInvalidatorForOrder(maker common.Address, orderHash [32]byte) ([]byte, error) {
	return routerv61inch.abi.Pack("rawRemainingInvalidatorForOrder", maker, orderHash)
}

// UnpackRawRemainingInvalidatorForOrder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc2a40753.
//
// Solidity: function rawRemainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (routerv61inch *Routerv61inch) UnpackRawRemainingInvalidatorForOrder(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("rawRemainingInvalidatorForOrder", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRemainingInvalidatorForOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x435b9789.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function remainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (routerv61inch *Routerv61inch) PackRemainingInvalidatorForOrder(maker common.Address, orderHash [32]byte) []byte {
	enc, err := routerv61inch.abi.Pack("remainingInvalidatorForOrder", maker, orderHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemainingInvalidatorForOrder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x435b9789.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function remainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (routerv61inch *Routerv61inch) TryPackRemainingInvalidatorForOrder(maker common.Address, orderHash [32]byte) ([]byte, error) {
	return routerv61inch.abi.Pack("remainingInvalidatorForOrder", maker, orderHash)
}

// UnpackRemainingInvalidatorForOrder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x435b9789.
//
// Solidity: function remainingInvalidatorForOrder(address maker, bytes32 orderHash) view returns(uint256)
func (routerv61inch *Routerv61inch) UnpackRemainingInvalidatorForOrder(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("remainingInvalidatorForOrder", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function renounceOwnership() returns()
func (routerv61inch *Routerv61inch) PackRenounceOwnership() []byte {
	enc, err := routerv61inch.abi.Pack("renounceOwnership")
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
func (routerv61inch *Routerv61inch) TryPackRenounceOwnership() ([]byte, error) {
	return routerv61inch.abi.Pack("renounceOwnership")
}

// PackRescueFunds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x78e3214f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (routerv61inch *Routerv61inch) PackRescueFunds(token common.Address, amount *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("rescueFunds", token, amount)
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
func (routerv61inch *Routerv61inch) TryPackRescueFunds(token common.Address, amount *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("rescueFunds", token, amount)
}

// PackSimulate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd61951d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function simulate(address target, bytes data) returns()
func (routerv61inch *Routerv61inch) PackSimulate(target common.Address, data []byte) []byte {
	enc, err := routerv61inch.abi.Pack("simulate", target, data)
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
func (routerv61inch *Routerv61inch) TryPackSimulate(target common.Address, data []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("simulate", target, data)
}

// PackSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07ed2379.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (routerv61inch *Routerv61inch) PackSwap(executor common.Address, desc GenericRouterSwapDescription, data []byte) []byte {
	enc, err := routerv61inch.abi.Pack("swap", executor, desc, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07ed2379.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (routerv61inch *Routerv61inch) TryPackSwap(executor common.Address, desc GenericRouterSwapDescription, data []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("swap", executor, desc, data)
}

// SwapOutput serves as a container for the return parameters of contract
// method Swap.
type SwapOutput struct {
	ReturnAmount *big.Int
	SpentAmount  *big.Int
}

// UnpackSwap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x07ed2379.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (routerv61inch *Routerv61inch) UnpackSwap(data []byte) (SwapOutput, error) {
	out, err := routerv61inch.abi.Unpack("swap", data)
	outstruct := new(SwapOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ReturnAmount = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.SpentAmount = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (routerv61inch *Routerv61inch) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := routerv61inch.abi.Pack("transferOwnership", newOwner)
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
func (routerv61inch *Routerv61inch) TryPackTransferOwnership(newOwner common.Address) ([]byte, error) {
	return routerv61inch.abi.Pack("transferOwnership", newOwner)
}

// PackUniswapV3SwapCallback is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfa461e33.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (routerv61inch *Routerv61inch) PackUniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) []byte {
	enc, err := routerv61inch.abi.Pack("uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
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
func (routerv61inch *Routerv61inch) TryPackUniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) ([]byte, error) {
	return routerv61inch.abi.Pack("uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// PackUnoswap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83800a8e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unoswap(uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackUnoswap(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("unoswap", token, amount, minReturn, dex)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnoswap is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83800a8e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unoswap(uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackUnoswap(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("unoswap", token, amount, minReturn, dex)
}

// UnpackUnoswap is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x83800a8e.
//
// Solidity: function unoswap(uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackUnoswap(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("unoswap", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUnoswap2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8770ba91.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unoswap2(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackUnoswap2(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("unoswap2", token, amount, minReturn, dex, dex2)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnoswap2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8770ba91.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unoswap2(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackUnoswap2(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("unoswap2", token, amount, minReturn, dex, dex2)
}

// UnpackUnoswap2 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8770ba91.
//
// Solidity: function unoswap2(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackUnoswap2(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("unoswap2", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUnoswap3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x19367472.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unoswap3(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackUnoswap3(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("unoswap3", token, amount, minReturn, dex, dex2, dex3)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnoswap3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x19367472.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unoswap3(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackUnoswap3(token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("unoswap3", token, amount, minReturn, dex, dex2, dex3)
}

// UnpackUnoswap3 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x19367472.
//
// Solidity: function unoswap3(uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackUnoswap3(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("unoswap3", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUnoswapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2c95c82.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unoswapTo(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackUnoswapTo(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("unoswapTo", to, token, amount, minReturn, dex)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnoswapTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2c95c82.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unoswapTo(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackUnoswapTo(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("unoswapTo", to, token, amount, minReturn, dex)
}

// UnpackUnoswapTo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe2c95c82.
//
// Solidity: function unoswapTo(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackUnoswapTo(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("unoswapTo", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUnoswapTo2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea76dddf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unoswapTo2(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackUnoswapTo2(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("unoswapTo2", to, token, amount, minReturn, dex, dex2)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnoswapTo2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea76dddf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unoswapTo2(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackUnoswapTo2(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("unoswapTo2", to, token, amount, minReturn, dex, dex2)
}

// UnpackUnoswapTo2 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xea76dddf.
//
// Solidity: function unoswapTo2(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackUnoswapTo2(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("unoswapTo2", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUnoswapTo3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf7a70056.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unoswapTo3(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) PackUnoswapTo3(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) []byte {
	enc, err := routerv61inch.abi.Pack("unoswapTo3", to, token, amount, minReturn, dex, dex2, dex3)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnoswapTo3 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf7a70056.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unoswapTo3(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) TryPackUnoswapTo3(to *big.Int, token *big.Int, amount *big.Int, minReturn *big.Int, dex *big.Int, dex2 *big.Int, dex3 *big.Int) ([]byte, error) {
	return routerv61inch.abi.Pack("unoswapTo3", to, token, amount, minReturn, dex, dex2, dex3)
}

// UnpackUnoswapTo3 is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf7a70056.
//
// Solidity: function unoswapTo3(uint256 to, uint256 token, uint256 amount, uint256 minReturn, uint256 dex, uint256 dex2, uint256 dex3) returns(uint256 returnAmount)
func (routerv61inch *Routerv61inch) UnpackUnoswapTo3(data []byte) (*big.Int, error) {
	out, err := routerv61inch.abi.Unpack("unoswapTo3", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unpause() returns()
func (routerv61inch *Routerv61inch) PackUnpause() []byte {
	enc, err := routerv61inch.abi.Pack("unpause")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnpause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f4ba83a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unpause() returns()
func (routerv61inch *Routerv61inch) TryPackUnpause() ([]byte, error) {
	return routerv61inch.abi.Pack("unpause")
}

// Routerv61inchBitInvalidatorUpdated represents a BitInvalidatorUpdated event raised by the Routerv61inch contract.
type Routerv61inchBitInvalidatorUpdated struct {
	Maker     common.Address
	SlotIndex *big.Int
	SlotValue *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const Routerv61inchBitInvalidatorUpdatedEventName = "BitInvalidatorUpdated"

// ContractEventName returns the user-defined event name.
func (Routerv61inchBitInvalidatorUpdated) ContractEventName() string {
	return Routerv61inchBitInvalidatorUpdatedEventName
}

// UnpackBitInvalidatorUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BitInvalidatorUpdated(address indexed maker, uint256 slotIndex, uint256 slotValue)
func (routerv61inch *Routerv61inch) UnpackBitInvalidatorUpdatedEvent(log *types.Log) (*Routerv61inchBitInvalidatorUpdated, error) {
	event := "BitInvalidatorUpdated"
	if log.Topics[0] != routerv61inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv61inchBitInvalidatorUpdated)
	if len(log.Data) > 0 {
		if err := routerv61inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv61inch.abi.Events[event].Inputs {
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

// Routerv61inchEIP712DomainChanged represents a EIP712DomainChanged event raised by the Routerv61inch contract.
type Routerv61inchEIP712DomainChanged struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const Routerv61inchEIP712DomainChangedEventName = "EIP712DomainChanged"

// ContractEventName returns the user-defined event name.
func (Routerv61inchEIP712DomainChanged) ContractEventName() string {
	return Routerv61inchEIP712DomainChangedEventName
}

// UnpackEIP712DomainChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EIP712DomainChanged()
func (routerv61inch *Routerv61inch) UnpackEIP712DomainChangedEvent(log *types.Log) (*Routerv61inchEIP712DomainChanged, error) {
	event := "EIP712DomainChanged"
	if log.Topics[0] != routerv61inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv61inchEIP712DomainChanged)
	if len(log.Data) > 0 {
		if err := routerv61inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv61inch.abi.Events[event].Inputs {
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

// Routerv61inchEpochIncreased represents a EpochIncreased event raised by the Routerv61inch contract.
type Routerv61inchEpochIncreased struct {
	Maker    common.Address
	Series   *big.Int
	NewEpoch *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const Routerv61inchEpochIncreasedEventName = "EpochIncreased"

// ContractEventName returns the user-defined event name.
func (Routerv61inchEpochIncreased) ContractEventName() string {
	return Routerv61inchEpochIncreasedEventName
}

// UnpackEpochIncreasedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EpochIncreased(address indexed maker, uint256 series, uint256 newEpoch)
func (routerv61inch *Routerv61inch) UnpackEpochIncreasedEvent(log *types.Log) (*Routerv61inchEpochIncreased, error) {
	event := "EpochIncreased"
	if log.Topics[0] != routerv61inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv61inchEpochIncreased)
	if len(log.Data) > 0 {
		if err := routerv61inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv61inch.abi.Events[event].Inputs {
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

// Routerv61inchOrderCancelled represents a OrderCancelled event raised by the Routerv61inch contract.
type Routerv61inchOrderCancelled struct {
	OrderHash [32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const Routerv61inchOrderCancelledEventName = "OrderCancelled"

// ContractEventName returns the user-defined event name.
func (Routerv61inchOrderCancelled) ContractEventName() string {
	return Routerv61inchOrderCancelledEventName
}

// UnpackOrderCancelledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OrderCancelled(bytes32 orderHash)
func (routerv61inch *Routerv61inch) UnpackOrderCancelledEvent(log *types.Log) (*Routerv61inchOrderCancelled, error) {
	event := "OrderCancelled"
	if log.Topics[0] != routerv61inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv61inchOrderCancelled)
	if len(log.Data) > 0 {
		if err := routerv61inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv61inch.abi.Events[event].Inputs {
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

// Routerv61inchOrderFilled represents a OrderFilled event raised by the Routerv61inch contract.
type Routerv61inchOrderFilled struct {
	OrderHash       [32]byte
	RemainingAmount *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const Routerv61inchOrderFilledEventName = "OrderFilled"

// ContractEventName returns the user-defined event name.
func (Routerv61inchOrderFilled) ContractEventName() string {
	return Routerv61inchOrderFilledEventName
}

// UnpackOrderFilledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OrderFilled(bytes32 orderHash, uint256 remainingAmount)
func (routerv61inch *Routerv61inch) UnpackOrderFilledEvent(log *types.Log) (*Routerv61inchOrderFilled, error) {
	event := "OrderFilled"
	if log.Topics[0] != routerv61inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv61inchOrderFilled)
	if len(log.Data) > 0 {
		if err := routerv61inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv61inch.abi.Events[event].Inputs {
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

// Routerv61inchOwnershipTransferred represents a OwnershipTransferred event raised by the Routerv61inch contract.
type Routerv61inchOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const Routerv61inchOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (Routerv61inchOwnershipTransferred) ContractEventName() string {
	return Routerv61inchOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (routerv61inch *Routerv61inch) UnpackOwnershipTransferredEvent(log *types.Log) (*Routerv61inchOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != routerv61inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv61inchOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := routerv61inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv61inch.abi.Events[event].Inputs {
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

// Routerv61inchPaused represents a Paused event raised by the Routerv61inch contract.
type Routerv61inchPaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const Routerv61inchPausedEventName = "Paused"

// ContractEventName returns the user-defined event name.
func (Routerv61inchPaused) ContractEventName() string {
	return Routerv61inchPausedEventName
}

// UnpackPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Paused(address account)
func (routerv61inch *Routerv61inch) UnpackPausedEvent(log *types.Log) (*Routerv61inchPaused, error) {
	event := "Paused"
	if log.Topics[0] != routerv61inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv61inchPaused)
	if len(log.Data) > 0 {
		if err := routerv61inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv61inch.abi.Events[event].Inputs {
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

// Routerv61inchUnpaused represents a Unpaused event raised by the Routerv61inch contract.
type Routerv61inchUnpaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const Routerv61inchUnpausedEventName = "Unpaused"

// ContractEventName returns the user-defined event name.
func (Routerv61inchUnpaused) ContractEventName() string {
	return Routerv61inchUnpausedEventName
}

// UnpackUnpausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Unpaused(address account)
func (routerv61inch *Routerv61inch) UnpackUnpausedEvent(log *types.Log) (*Routerv61inchUnpaused, error) {
	event := "Unpaused"
	if log.Topics[0] != routerv61inch.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Routerv61inchUnpaused)
	if len(log.Data) > 0 {
		if err := routerv61inch.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range routerv61inch.abi.Events[event].Inputs {
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
func (routerv61inch *Routerv61inch) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["AdvanceEpochFailed"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackAdvanceEpochFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["ArbitraryStaticCallFailed"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackArbitraryStaticCallFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["BadCurveSwapSelector"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackBadCurveSwapSelectorError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["BadPool"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackBadPoolError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["BadSignature"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackBadSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["BitInvalidatedOrder"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackBitInvalidatedOrderError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["ETHTransferFailed"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackETHTransferFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["EnforcedPause"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackEnforcedPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["EpochManagerAndBitInvalidatorsAreIncompatible"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackEpochManagerAndBitInvalidatorsAreIncompatibleError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["EthDepositRejected"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackEthDepositRejectedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["ExpectedPause"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackExpectedPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["InsufficientBalance"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackInsufficientBalanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["InvalidMsgValue"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackInvalidMsgValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["InvalidPermit2Transfer"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackInvalidPermit2TransferError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["InvalidShortString"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackInvalidShortStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["InvalidatedOrder"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackInvalidatedOrderError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["MakingAmountTooLow"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackMakingAmountTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["MismatchArraysLengths"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackMismatchArraysLengthsError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["OrderExpired"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackOrderExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["OrderIsNotSuitableForMassInvalidation"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackOrderIsNotSuitableForMassInvalidationError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["PartialFillNotAllowed"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackPartialFillNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["Permit2TransferAmountTooHigh"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackPermit2TransferAmountTooHighError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["PredicateIsNotTrue"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackPredicateIsNotTrueError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["PrivateOrder"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackPrivateOrderError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["ReentrancyDetected"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackReentrancyDetectedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["RemainingInvalidatedOrder"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackRemainingInvalidatedOrderError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["ReservesCallFailed"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackReservesCallFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["ReturnAmountIsNotEnough"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackReturnAmountIsNotEnoughError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["SafeTransferFailed"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackSafeTransferFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["SafeTransferFromFailed"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackSafeTransferFromFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["SimulationResults"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackSimulationResultsError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["StringTooLong"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackStringTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["SwapWithZeroAmount"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackSwapWithZeroAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["TakingAmountExceeded"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackTakingAmountExceededError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["TakingAmountTooHigh"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackTakingAmountTooHighError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["TransferFromMakerToTakerFailed"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackTransferFromMakerToTakerFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["TransferFromTakerToMakerFailed"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackTransferFromTakerToMakerFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["WrongSeriesNonce"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackWrongSeriesNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["ZeroAddress"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], routerv61inch.abi.Errors["ZeroMinReturn"].ID.Bytes()[:4]) {
		return routerv61inch.UnpackZeroMinReturnError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// Routerv61inchAdvanceEpochFailed represents a AdvanceEpochFailed error raised by the Routerv61inch contract.
type Routerv61inchAdvanceEpochFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AdvanceEpochFailed()
func Routerv61inchAdvanceEpochFailedErrorID() common.Hash {
	return common.HexToHash("0x555fbbbfeabfd971ce92d463313ac8e97eef15283bea456e0b07b53b772f21fb")
}

// UnpackAdvanceEpochFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AdvanceEpochFailed()
func (routerv61inch *Routerv61inch) UnpackAdvanceEpochFailedError(raw []byte) (*Routerv61inchAdvanceEpochFailed, error) {
	out := new(Routerv61inchAdvanceEpochFailed)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "AdvanceEpochFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchArbitraryStaticCallFailed represents a ArbitraryStaticCallFailed error raised by the Routerv61inch contract.
type Routerv61inchArbitraryStaticCallFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ArbitraryStaticCallFailed()
func Routerv61inchArbitraryStaticCallFailedErrorID() common.Hash {
	return common.HexToHash("0x1f1b8f6185f3387f7dd1c7ff816bd1dd2b3f59e21a9f1807a270349ee1bd307c")
}

// UnpackArbitraryStaticCallFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ArbitraryStaticCallFailed()
func (routerv61inch *Routerv61inch) UnpackArbitraryStaticCallFailedError(raw []byte) (*Routerv61inchArbitraryStaticCallFailed, error) {
	out := new(Routerv61inchArbitraryStaticCallFailed)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "ArbitraryStaticCallFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchBadCurveSwapSelector represents a BadCurveSwapSelector error raised by the Routerv61inch contract.
type Routerv61inchBadCurveSwapSelector struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BadCurveSwapSelector()
func Routerv61inchBadCurveSwapSelectorErrorID() common.Hash {
	return common.HexToHash("0xa231cb8233d288e430caedc48c61916542b547fdf9f0d1e9c5e7bee2c74d3f6a")
}

// UnpackBadCurveSwapSelectorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BadCurveSwapSelector()
func (routerv61inch *Routerv61inch) UnpackBadCurveSwapSelectorError(raw []byte) (*Routerv61inchBadCurveSwapSelector, error) {
	out := new(Routerv61inchBadCurveSwapSelector)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "BadCurveSwapSelector", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchBadPool represents a BadPool error raised by the Routerv61inch contract.
type Routerv61inchBadPool struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BadPool()
func Routerv61inchBadPoolErrorID() common.Hash {
	return common.HexToHash("0xb2c02722cf230da6a05b6ae0e22f42ed25be4bf9b34cb4514ebd83ff4a53308a")
}

// UnpackBadPoolError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BadPool()
func (routerv61inch *Routerv61inch) UnpackBadPoolError(raw []byte) (*Routerv61inchBadPool, error) {
	out := new(Routerv61inchBadPool)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "BadPool", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchBadSignature represents a BadSignature error raised by the Routerv61inch contract.
type Routerv61inchBadSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BadSignature()
func Routerv61inchBadSignatureErrorID() common.Hash {
	return common.HexToHash("0x5cd5d2335541c4f2ed05fbe44f397e8b79f8e2333157122d2dab06e378ef7685")
}

// UnpackBadSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BadSignature()
func (routerv61inch *Routerv61inch) UnpackBadSignatureError(raw []byte) (*Routerv61inchBadSignature, error) {
	out := new(Routerv61inchBadSignature)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "BadSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchBitInvalidatedOrder represents a BitInvalidatedOrder error raised by the Routerv61inch contract.
type Routerv61inchBitInvalidatedOrder struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BitInvalidatedOrder()
func Routerv61inchBitInvalidatedOrderErrorID() common.Hash {
	return common.HexToHash("0xa4f62a96cee3aa300c0bbbbcb25528165e928f62f3c82489f7cbe98a289031ff")
}

// UnpackBitInvalidatedOrderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BitInvalidatedOrder()
func (routerv61inch *Routerv61inch) UnpackBitInvalidatedOrderError(raw []byte) (*Routerv61inchBitInvalidatedOrder, error) {
	out := new(Routerv61inchBitInvalidatedOrder)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "BitInvalidatedOrder", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchETHTransferFailed represents a ETHTransferFailed error raised by the Routerv61inch contract.
type Routerv61inchETHTransferFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ETHTransferFailed()
func Routerv61inchETHTransferFailedErrorID() common.Hash {
	return common.HexToHash("0xb12d13ebe76e15b5fdb7bf52f0daba617b83ebcc560b0666c44fcdcd71f4362b")
}

// UnpackETHTransferFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ETHTransferFailed()
func (routerv61inch *Routerv61inch) UnpackETHTransferFailedError(raw []byte) (*Routerv61inchETHTransferFailed, error) {
	out := new(Routerv61inchETHTransferFailed)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "ETHTransferFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchEnforcedPause represents a EnforcedPause error raised by the Routerv61inch contract.
type Routerv61inchEnforcedPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EnforcedPause()
func Routerv61inchEnforcedPauseErrorID() common.Hash {
	return common.HexToHash("0xd93c0665d6c96d04a8f174024fc4ddd66c250604aff22bbec808de86dd3637e3")
}

// UnpackEnforcedPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EnforcedPause()
func (routerv61inch *Routerv61inch) UnpackEnforcedPauseError(raw []byte) (*Routerv61inchEnforcedPause, error) {
	out := new(Routerv61inchEnforcedPause)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "EnforcedPause", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchEpochManagerAndBitInvalidatorsAreIncompatible represents a EpochManagerAndBitInvalidatorsAreIncompatible error raised by the Routerv61inch contract.
type Routerv61inchEpochManagerAndBitInvalidatorsAreIncompatible struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EpochManagerAndBitInvalidatorsAreIncompatible()
func Routerv61inchEpochManagerAndBitInvalidatorsAreIncompatibleErrorID() common.Hash {
	return common.HexToHash("0x9e744e25935f4a5fccfeb474ae30bf3df59df3b9ef8da0ded0c4fc30f56a4942")
}

// UnpackEpochManagerAndBitInvalidatorsAreIncompatibleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EpochManagerAndBitInvalidatorsAreIncompatible()
func (routerv61inch *Routerv61inch) UnpackEpochManagerAndBitInvalidatorsAreIncompatibleError(raw []byte) (*Routerv61inchEpochManagerAndBitInvalidatorsAreIncompatible, error) {
	out := new(Routerv61inchEpochManagerAndBitInvalidatorsAreIncompatible)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "EpochManagerAndBitInvalidatorsAreIncompatible", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchEthDepositRejected represents a EthDepositRejected error raised by the Routerv61inch contract.
type Routerv61inchEthDepositRejected struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EthDepositRejected()
func Routerv61inchEthDepositRejectedErrorID() common.Hash {
	return common.HexToHash("0x1b10b0f9ae66bdc9d7cb534137e3b02811ba15910619f3dcc6f5f5e2f8e91721")
}

// UnpackEthDepositRejectedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EthDepositRejected()
func (routerv61inch *Routerv61inch) UnpackEthDepositRejectedError(raw []byte) (*Routerv61inchEthDepositRejected, error) {
	out := new(Routerv61inchEthDepositRejected)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "EthDepositRejected", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchExpectedPause represents a ExpectedPause error raised by the Routerv61inch contract.
type Routerv61inchExpectedPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExpectedPause()
func Routerv61inchExpectedPauseErrorID() common.Hash {
	return common.HexToHash("0x8dfc202bcfe9a735b559bee70674422512bc5c30f687046ae8778315fb81da44")
}

// UnpackExpectedPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExpectedPause()
func (routerv61inch *Routerv61inch) UnpackExpectedPauseError(raw []byte) (*Routerv61inchExpectedPause, error) {
	out := new(Routerv61inchExpectedPause)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "ExpectedPause", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchInsufficientBalance represents a InsufficientBalance error raised by the Routerv61inch contract.
type Routerv61inchInsufficientBalance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientBalance()
func Routerv61inchInsufficientBalanceErrorID() common.Hash {
	return common.HexToHash("0xf4d678b8ce6b5157126b1484a53523762a93571537a7d5ae97d8014a44715c94")
}

// UnpackInsufficientBalanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientBalance()
func (routerv61inch *Routerv61inch) UnpackInsufficientBalanceError(raw []byte) (*Routerv61inchInsufficientBalance, error) {
	out := new(Routerv61inchInsufficientBalance)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "InsufficientBalance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchInvalidMsgValue represents a InvalidMsgValue error raised by the Routerv61inch contract.
type Routerv61inchInvalidMsgValue struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidMsgValue()
func Routerv61inchInvalidMsgValueErrorID() common.Hash {
	return common.HexToHash("0x1841b4e1b5bc2b6ee2d929f61c1a6d695028c1b47aa99b13e72b417bfaebc3cd")
}

// UnpackInvalidMsgValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidMsgValue()
func (routerv61inch *Routerv61inch) UnpackInvalidMsgValueError(raw []byte) (*Routerv61inchInvalidMsgValue, error) {
	out := new(Routerv61inchInvalidMsgValue)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "InvalidMsgValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchInvalidPermit2Transfer represents a InvalidPermit2Transfer error raised by the Routerv61inch contract.
type Routerv61inchInvalidPermit2Transfer struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPermit2Transfer()
func Routerv61inchInvalidPermit2TransferErrorID() common.Hash {
	return common.HexToHash("0x2aefd060ad0cabed1a25fad3d53e9f3a00a3cc41eab754c8ca1c4e004c79199e")
}

// UnpackInvalidPermit2TransferError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPermit2Transfer()
func (routerv61inch *Routerv61inch) UnpackInvalidPermit2TransferError(raw []byte) (*Routerv61inchInvalidPermit2Transfer, error) {
	out := new(Routerv61inchInvalidPermit2Transfer)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "InvalidPermit2Transfer", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchInvalidShortString represents a InvalidShortString error raised by the Routerv61inch contract.
type Routerv61inchInvalidShortString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidShortString()
func Routerv61inchInvalidShortStringErrorID() common.Hash {
	return common.HexToHash("0xb3512b0c6163e5f0bafab72bb631b9d58cd7a731b082f910338aa21c83d5c274")
}

// UnpackInvalidShortStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidShortString()
func (routerv61inch *Routerv61inch) UnpackInvalidShortStringError(raw []byte) (*Routerv61inchInvalidShortString, error) {
	out := new(Routerv61inchInvalidShortString)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "InvalidShortString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchInvalidatedOrder represents a InvalidatedOrder error raised by the Routerv61inch contract.
type Routerv61inchInvalidatedOrder struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidatedOrder()
func Routerv61inchInvalidatedOrderErrorID() common.Hash {
	return common.HexToHash("0xf71fbda25ab7ea5b8d7774a17207b5659639b9c5360eda3ff0782c89adfd8eeb")
}

// UnpackInvalidatedOrderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidatedOrder()
func (routerv61inch *Routerv61inch) UnpackInvalidatedOrderError(raw []byte) (*Routerv61inchInvalidatedOrder, error) {
	out := new(Routerv61inchInvalidatedOrder)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "InvalidatedOrder", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchMakingAmountTooLow represents a MakingAmountTooLow error raised by the Routerv61inch contract.
type Routerv61inchMakingAmountTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MakingAmountTooLow()
func Routerv61inchMakingAmountTooLowErrorID() common.Hash {
	return common.HexToHash("0x481ea392150ef4334793b50f5628cfd068e177118e6993d7c20fd36fdfccc9c1")
}

// UnpackMakingAmountTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MakingAmountTooLow()
func (routerv61inch *Routerv61inch) UnpackMakingAmountTooLowError(raw []byte) (*Routerv61inchMakingAmountTooLow, error) {
	out := new(Routerv61inchMakingAmountTooLow)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "MakingAmountTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchMismatchArraysLengths represents a MismatchArraysLengths error raised by the Routerv61inch contract.
type Routerv61inchMismatchArraysLengths struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MismatchArraysLengths()
func Routerv61inchMismatchArraysLengthsErrorID() common.Hash {
	return common.HexToHash("0xd97cd9d8b91f1c3a5f1a62cd435081a750a2efbcd98b364537f060998e4c9d2a")
}

// UnpackMismatchArraysLengthsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MismatchArraysLengths()
func (routerv61inch *Routerv61inch) UnpackMismatchArraysLengthsError(raw []byte) (*Routerv61inchMismatchArraysLengths, error) {
	out := new(Routerv61inchMismatchArraysLengths)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "MismatchArraysLengths", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchOrderExpired represents a OrderExpired error raised by the Routerv61inch contract.
type Routerv61inchOrderExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OrderExpired()
func Routerv61inchOrderExpiredErrorID() common.Hash {
	return common.HexToHash("0xc56873bac2ec2d1dd6d40a2a9b432692a2b77a4253c07d6f7ef4929f60ced89c")
}

// UnpackOrderExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OrderExpired()
func (routerv61inch *Routerv61inch) UnpackOrderExpiredError(raw []byte) (*Routerv61inchOrderExpired, error) {
	out := new(Routerv61inchOrderExpired)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "OrderExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchOrderIsNotSuitableForMassInvalidation represents a OrderIsNotSuitableForMassInvalidation error raised by the Routerv61inch contract.
type Routerv61inchOrderIsNotSuitableForMassInvalidation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OrderIsNotSuitableForMassInvalidation()
func Routerv61inchOrderIsNotSuitableForMassInvalidationErrorID() common.Hash {
	return common.HexToHash("0x86bffacaa6490b46f621a9fb7cd99cfd35a57e3c842dc6b21f5ab4017991881d")
}

// UnpackOrderIsNotSuitableForMassInvalidationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OrderIsNotSuitableForMassInvalidation()
func (routerv61inch *Routerv61inch) UnpackOrderIsNotSuitableForMassInvalidationError(raw []byte) (*Routerv61inchOrderIsNotSuitableForMassInvalidation, error) {
	out := new(Routerv61inchOrderIsNotSuitableForMassInvalidation)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "OrderIsNotSuitableForMassInvalidation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the Routerv61inch contract.
type Routerv61inchOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func Routerv61inchOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (routerv61inch *Routerv61inch) UnpackOwnableInvalidOwnerError(raw []byte) (*Routerv61inchOwnableInvalidOwner, error) {
	out := new(Routerv61inchOwnableInvalidOwner)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the Routerv61inch contract.
type Routerv61inchOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func Routerv61inchOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (routerv61inch *Routerv61inch) UnpackOwnableUnauthorizedAccountError(raw []byte) (*Routerv61inchOwnableUnauthorizedAccount, error) {
	out := new(Routerv61inchOwnableUnauthorizedAccount)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchPartialFillNotAllowed represents a PartialFillNotAllowed error raised by the Routerv61inch contract.
type Routerv61inchPartialFillNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PartialFillNotAllowed()
func Routerv61inchPartialFillNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x8ef0017c7e87b3e0bc22365d7355061577475bd21ddecd5075ef6878171adaee")
}

// UnpackPartialFillNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PartialFillNotAllowed()
func (routerv61inch *Routerv61inch) UnpackPartialFillNotAllowedError(raw []byte) (*Routerv61inchPartialFillNotAllowed, error) {
	out := new(Routerv61inchPartialFillNotAllowed)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "PartialFillNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchPermit2TransferAmountTooHigh represents a Permit2TransferAmountTooHigh error raised by the Routerv61inch contract.
type Routerv61inchPermit2TransferAmountTooHigh struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Permit2TransferAmountTooHigh()
func Routerv61inchPermit2TransferAmountTooHighErrorID() common.Hash {
	return common.HexToHash("0x8112e1191b8d69967e637b46dd14d8cb26f6fe6b51228e930f72cb67b0b3b292")
}

// UnpackPermit2TransferAmountTooHighError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Permit2TransferAmountTooHigh()
func (routerv61inch *Routerv61inch) UnpackPermit2TransferAmountTooHighError(raw []byte) (*Routerv61inchPermit2TransferAmountTooHigh, error) {
	out := new(Routerv61inchPermit2TransferAmountTooHigh)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "Permit2TransferAmountTooHigh", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchPredicateIsNotTrue represents a PredicateIsNotTrue error raised by the Routerv61inch contract.
type Routerv61inchPredicateIsNotTrue struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PredicateIsNotTrue()
func Routerv61inchPredicateIsNotTrueErrorID() common.Hash {
	return common.HexToHash("0xb6629c02d7b61b434160b9a6b56d99a2aa5c959e50d9876456b23faf101fdacf")
}

// UnpackPredicateIsNotTrueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PredicateIsNotTrue()
func (routerv61inch *Routerv61inch) UnpackPredicateIsNotTrueError(raw []byte) (*Routerv61inchPredicateIsNotTrue, error) {
	out := new(Routerv61inchPredicateIsNotTrue)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "PredicateIsNotTrue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchPrivateOrder represents a PrivateOrder error raised by the Routerv61inch contract.
type Routerv61inchPrivateOrder struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PrivateOrder()
func Routerv61inchPrivateOrderErrorID() common.Hash {
	return common.HexToHash("0xd4dfdafe06416a471ed7b8bd430201111ccc397ff3d5636080df9597d0d8626c")
}

// UnpackPrivateOrderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PrivateOrder()
func (routerv61inch *Routerv61inch) UnpackPrivateOrderError(raw []byte) (*Routerv61inchPrivateOrder, error) {
	out := new(Routerv61inchPrivateOrder)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "PrivateOrder", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchReentrancyDetected represents a ReentrancyDetected error raised by the Routerv61inch contract.
type Routerv61inchReentrancyDetected struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReentrancyDetected()
func Routerv61inchReentrancyDetectedErrorID() common.Hash {
	return common.HexToHash("0xc5f2be51ec4ec0ad8a7972d497da993a6fcbb89cf72c05f97d654ed81ce53492")
}

// UnpackReentrancyDetectedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReentrancyDetected()
func (routerv61inch *Routerv61inch) UnpackReentrancyDetectedError(raw []byte) (*Routerv61inchReentrancyDetected, error) {
	out := new(Routerv61inchReentrancyDetected)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "ReentrancyDetected", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchRemainingInvalidatedOrder represents a RemainingInvalidatedOrder error raised by the Routerv61inch contract.
type Routerv61inchRemainingInvalidatedOrder struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RemainingInvalidatedOrder()
func Routerv61inchRemainingInvalidatedOrderErrorID() common.Hash {
	return common.HexToHash("0xaa3eef95872d48f4732bb71c93c18b36b9106bca38221b603d3a1fd8c04962a4")
}

// UnpackRemainingInvalidatedOrderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RemainingInvalidatedOrder()
func (routerv61inch *Routerv61inch) UnpackRemainingInvalidatedOrderError(raw []byte) (*Routerv61inchRemainingInvalidatedOrder, error) {
	out := new(Routerv61inchRemainingInvalidatedOrder)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "RemainingInvalidatedOrder", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchReservesCallFailed represents a ReservesCallFailed error raised by the Routerv61inch contract.
type Routerv61inchReservesCallFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReservesCallFailed()
func Routerv61inchReservesCallFailedErrorID() common.Hash {
	return common.HexToHash("0x85cd58dcd62237b1af0231b4a8943a369462ba5bafe0a59c9a3b7637f8e6b2ac")
}

// UnpackReservesCallFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReservesCallFailed()
func (routerv61inch *Routerv61inch) UnpackReservesCallFailedError(raw []byte) (*Routerv61inchReservesCallFailed, error) {
	out := new(Routerv61inchReservesCallFailed)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "ReservesCallFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchReturnAmountIsNotEnough represents a ReturnAmountIsNotEnough error raised by the Routerv61inch contract.
type Routerv61inchReturnAmountIsNotEnough struct {
	Result    *big.Int
	MinReturn *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReturnAmountIsNotEnough(uint256 result, uint256 minReturn)
func Routerv61inchReturnAmountIsNotEnoughErrorID() common.Hash {
	return common.HexToHash("0x064a4ec669ebbcf350a2a63a12812ea0f8abc55c4d6f719192fa02339bfb5783")
}

// UnpackReturnAmountIsNotEnoughError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReturnAmountIsNotEnough(uint256 result, uint256 minReturn)
func (routerv61inch *Routerv61inch) UnpackReturnAmountIsNotEnoughError(raw []byte) (*Routerv61inchReturnAmountIsNotEnough, error) {
	out := new(Routerv61inchReturnAmountIsNotEnough)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "ReturnAmountIsNotEnough", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchSafeTransferFailed represents a SafeTransferFailed error raised by the Routerv61inch contract.
type Routerv61inchSafeTransferFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeTransferFailed()
func Routerv61inchSafeTransferFailedErrorID() common.Hash {
	return common.HexToHash("0xfb7f50796995f43b6f601c7bdc661fff3554a3197898a14efb94c31f671088ce")
}

// UnpackSafeTransferFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeTransferFailed()
func (routerv61inch *Routerv61inch) UnpackSafeTransferFailedError(raw []byte) (*Routerv61inchSafeTransferFailed, error) {
	out := new(Routerv61inchSafeTransferFailed)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "SafeTransferFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchSafeTransferFromFailed represents a SafeTransferFromFailed error raised by the Routerv61inch contract.
type Routerv61inchSafeTransferFromFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeTransferFromFailed()
func Routerv61inchSafeTransferFromFailedErrorID() common.Hash {
	return common.HexToHash("0xf4059071351ef2c2ce34e72012fb887b861f62f289b68b0c4e171c82a6b2023d")
}

// UnpackSafeTransferFromFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeTransferFromFailed()
func (routerv61inch *Routerv61inch) UnpackSafeTransferFromFailedError(raw []byte) (*Routerv61inchSafeTransferFromFailed, error) {
	out := new(Routerv61inchSafeTransferFromFailed)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "SafeTransferFromFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchSimulationResults represents a SimulationResults error raised by the Routerv61inch contract.
type Routerv61inchSimulationResults struct {
	Success bool
	Res     []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SimulationResults(bool success, bytes res)
func Routerv61inchSimulationResultsErrorID() common.Hash {
	return common.HexToHash("0x1934afc82dd770c05787a7a944999e73cb2b7d95aae2db5cf2dd220cffb39cf7")
}

// UnpackSimulationResultsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SimulationResults(bool success, bytes res)
func (routerv61inch *Routerv61inch) UnpackSimulationResultsError(raw []byte) (*Routerv61inchSimulationResults, error) {
	out := new(Routerv61inchSimulationResults)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "SimulationResults", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchStringTooLong represents a StringTooLong error raised by the Routerv61inch contract.
type Routerv61inchStringTooLong struct {
	Str string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StringTooLong(string str)
func Routerv61inchStringTooLongErrorID() common.Hash {
	return common.HexToHash("0x305a27a93f8e33b7392df0a0f91d6fc63847395853c45991eec52dbf24d72381")
}

// UnpackStringTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StringTooLong(string str)
func (routerv61inch *Routerv61inch) UnpackStringTooLongError(raw []byte) (*Routerv61inchStringTooLong, error) {
	out := new(Routerv61inchStringTooLong)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "StringTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchSwapWithZeroAmount represents a SwapWithZeroAmount error raised by the Routerv61inch contract.
type Routerv61inchSwapWithZeroAmount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SwapWithZeroAmount()
func Routerv61inchSwapWithZeroAmountErrorID() common.Hash {
	return common.HexToHash("0xfba5a276a6a0bc3e7a381b7cdc29e9d356a2baf3ab47924f586f2e1eb54db7f9")
}

// UnpackSwapWithZeroAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SwapWithZeroAmount()
func (routerv61inch *Routerv61inch) UnpackSwapWithZeroAmountError(raw []byte) (*Routerv61inchSwapWithZeroAmount, error) {
	out := new(Routerv61inchSwapWithZeroAmount)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "SwapWithZeroAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchTakingAmountExceeded represents a TakingAmountExceeded error raised by the Routerv61inch contract.
type Routerv61inchTakingAmountExceeded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TakingAmountExceeded()
func Routerv61inchTakingAmountExceededErrorID() common.Hash {
	return common.HexToHash("0x7f902a93c23e1e7f62cc5268147eaca1f0798d33dd25f87e80c63d453e0b5913")
}

// UnpackTakingAmountExceededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TakingAmountExceeded()
func (routerv61inch *Routerv61inch) UnpackTakingAmountExceededError(raw []byte) (*Routerv61inchTakingAmountExceeded, error) {
	out := new(Routerv61inchTakingAmountExceeded)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "TakingAmountExceeded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchTakingAmountTooHigh represents a TakingAmountTooHigh error raised by the Routerv61inch contract.
type Routerv61inchTakingAmountTooHigh struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TakingAmountTooHigh()
func Routerv61inchTakingAmountTooHighErrorID() common.Hash {
	return common.HexToHash("0xfb8ae129bb15dd7d41422f958c6ce0c12a8a6756e86abf2d1f1edaa920524a69")
}

// UnpackTakingAmountTooHighError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TakingAmountTooHigh()
func (routerv61inch *Routerv61inch) UnpackTakingAmountTooHighError(raw []byte) (*Routerv61inchTakingAmountTooHigh, error) {
	out := new(Routerv61inchTakingAmountTooHigh)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "TakingAmountTooHigh", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchTransferFromMakerToTakerFailed represents a TransferFromMakerToTakerFailed error raised by the Routerv61inch contract.
type Routerv61inchTransferFromMakerToTakerFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferFromMakerToTakerFailed()
func Routerv61inchTransferFromMakerToTakerFailedErrorID() common.Hash {
	return common.HexToHash("0x70a03f480a56f2ba97a5ae95726db5c4aa88d6b043211db6f6841f90c71390ac")
}

// UnpackTransferFromMakerToTakerFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferFromMakerToTakerFailed()
func (routerv61inch *Routerv61inch) UnpackTransferFromMakerToTakerFailedError(raw []byte) (*Routerv61inchTransferFromMakerToTakerFailed, error) {
	out := new(Routerv61inchTransferFromMakerToTakerFailed)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "TransferFromMakerToTakerFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchTransferFromTakerToMakerFailed represents a TransferFromTakerToMakerFailed error raised by the Routerv61inch contract.
type Routerv61inchTransferFromTakerToMakerFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TransferFromTakerToMakerFailed()
func Routerv61inchTransferFromTakerToMakerFailedErrorID() common.Hash {
	return common.HexToHash("0x478a5205e5ee1fb19e08efddef0ae39549de8d23a39befdc0b65f352d124ffd3")
}

// UnpackTransferFromTakerToMakerFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TransferFromTakerToMakerFailed()
func (routerv61inch *Routerv61inch) UnpackTransferFromTakerToMakerFailedError(raw []byte) (*Routerv61inchTransferFromTakerToMakerFailed, error) {
	out := new(Routerv61inchTransferFromTakerToMakerFailed)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "TransferFromTakerToMakerFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchWrongSeriesNonce represents a WrongSeriesNonce error raised by the Routerv61inch contract.
type Routerv61inchWrongSeriesNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WrongSeriesNonce()
func Routerv61inchWrongSeriesNonceErrorID() common.Hash {
	return common.HexToHash("0xe3e8b052cdfe2ab6e42c1d9cfa30e8e68c2d8c9a154f1b861b829b2ca1c65101")
}

// UnpackWrongSeriesNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WrongSeriesNonce()
func (routerv61inch *Routerv61inch) UnpackWrongSeriesNonceError(raw []byte) (*Routerv61inchWrongSeriesNonce, error) {
	out := new(Routerv61inchWrongSeriesNonce)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "WrongSeriesNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchZeroAddress represents a ZeroAddress error raised by the Routerv61inch contract.
type Routerv61inchZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroAddress()
func Routerv61inchZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xd92e233df2717d4a40030e20904abd27b68fcbeede117eaaccbbdac9618c8c73")
}

// UnpackZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroAddress()
func (routerv61inch *Routerv61inch) UnpackZeroAddressError(raw []byte) (*Routerv61inchZeroAddress, error) {
	out := new(Routerv61inchZeroAddress)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "ZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Routerv61inchZeroMinReturn represents a ZeroMinReturn error raised by the Routerv61inch contract.
type Routerv61inchZeroMinReturn struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroMinReturn()
func Routerv61inchZeroMinReturnErrorID() common.Hash {
	return common.HexToHash("0x0262dde406f45de042c06c7df2790374a9d79ccf79e8ab35852b326859e71236")
}

// UnpackZeroMinReturnError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroMinReturn()
func (routerv61inch *Routerv61inch) UnpackZeroMinReturnError(raw []byte) (*Routerv61inchZeroMinReturn, error) {
	out := new(Routerv61inchZeroMinReturn)
	if err := routerv61inch.abi.UnpackIntoInterface(out, "ZeroMinReturn", raw); err != nil {
		return nil, err
	}
	return out, nil
}
