// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aerodrome_router

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

// IRouterRoute is an auto generated low-level Go binding around an user-defined struct.
type IRouterRoute struct {
	From    common.Address
	To      common.Address
	Stable  bool
	Factory common.Address
}

// IRouterZap is an auto generated low-level Go binding around an user-defined struct.
type IRouterZap struct {
	TokenA        common.Address
	TokenB        common.Address
	Stable        bool
	Factory       common.Address
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}

// AerodromeRouterMetaData contains all meta data concerning the AerodromeRouter contract.
var AerodromeRouterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factoryRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountADesired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountAOptimal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountB\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAmountBDesired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientOutputAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmountInForETHDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRouteA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRouteB\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenInForETHDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolFactoryDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"UNSAFE_swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountInA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInB\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"}],\"name\":\"generateZapInParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"}],\"name\":\"generateZapOutParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"poolFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"}],\"name\":\"quoteAddLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"quoteRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"quoteStableLiquidityRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"sortTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routes\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountInA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInB\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"internalType\":\"structIRouter.Zap\",\"name\":\"zapInPool\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stake\",\"type\":\"bool\"}],\"name\":\"zapIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"internalType\":\"structIRouter.Zap\",\"name\":\"zapOutPool\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routesA\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"}],\"internalType\":\"structIRouter.Route[]\",\"name\":\"routesB\",\"type\":\"tuple[]\"}],\"name\":\"zapOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "AerodromeRouter",
}

// AerodromeRouter is an auto generated Go binding around an Ethereum contract.
type AerodromeRouter struct {
	abi abi.ABI
}

// NewAerodromeRouter creates a new instance of AerodromeRouter.
func NewAerodromeRouter() *AerodromeRouter {
	parsed, err := AerodromeRouterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &AerodromeRouter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *AerodromeRouter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _forwarder, address _factoryRegistry, address _factory, address _voter, address _weth) returns()
func (aerodromeRouter *AerodromeRouter) PackConstructor(_forwarder common.Address, _factoryRegistry common.Address, _factory common.Address, _voter common.Address, _weth common.Address) []byte {
	enc, err := aerodromeRouter.abi.Pack("", _forwarder, _factoryRegistry, _factory, _voter, _weth)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackETHER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42cb1fbc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ETHER() view returns(address)
func (aerodromeRouter *AerodromeRouter) PackETHER() []byte {
	enc, err := aerodromeRouter.abi.Pack("ETHER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackETHER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42cb1fbc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ETHER() view returns(address)
func (aerodromeRouter *AerodromeRouter) TryPackETHER() ([]byte, error) {
	return aerodromeRouter.abi.Pack("ETHER")
}

// UnpackETHER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x42cb1fbc.
//
// Solidity: function ETHER() view returns(address)
func (aerodromeRouter *AerodromeRouter) UnpackETHER(data []byte) (common.Address, error) {
	out, err := aerodromeRouter.abi.Unpack("ETHER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackUNSAFESwapExactTokensForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4111d597.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[])
func (aerodromeRouter *AerodromeRouter) PackUNSAFESwapExactTokensForTokens(amounts []*big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("UNSAFE_swapExactTokensForTokens", amounts, routes, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUNSAFESwapExactTokensForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4111d597.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[])
func (aerodromeRouter *AerodromeRouter) TryPackUNSAFESwapExactTokensForTokens(amounts []*big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("UNSAFE_swapExactTokensForTokens", amounts, routes, to, deadline)
}

// UnpackUNSAFESwapExactTokensForTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4111d597.
//
// Solidity: function UNSAFE_swapExactTokensForTokens(uint256[] amounts, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[])
func (aerodromeRouter *AerodromeRouter) UnpackUNSAFESwapExactTokensForTokens(data []byte) ([]*big.Int, error) {
	out, err := aerodromeRouter.abi.Unpack("UNSAFE_swapExactTokensForTokens", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackAddLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a47ddc3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) PackAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("addLiquidity", tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a47ddc3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) TryPackAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("addLiquidity", tokenA, tokenB, stable, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityOutput serves as a container for the return parameters of contract
// method AddLiquidity.
type AddLiquidityOutput struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}

// UnpackAddLiquidity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5a47ddc3.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, bool stable, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) UnpackAddLiquidity(data []byte) (AddLiquidityOutput, error) {
	out, err := aerodromeRouter.abi.Unpack("addLiquidity", data)
	outstruct := new(AddLiquidityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Liquidity = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackAddLiquidityETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb7e0d4c0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) PackAddLiquidityETH(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("addLiquidityETH", token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddLiquidityETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb7e0d4c0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) TryPackAddLiquidityETH(token common.Address, stable bool, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("addLiquidityETH", token, stable, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline)
}

// AddLiquidityETHOutput serves as a container for the return parameters of contract
// method AddLiquidityETH.
type AddLiquidityETHOutput struct {
	AmountToken *big.Int
	AmountETH   *big.Int
	Liquidity   *big.Int
}

// UnpackAddLiquidityETH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb7e0d4c0.
//
// Solidity: function addLiquidityETH(address token, bool stable, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) UnpackAddLiquidityETH(data []byte) (AddLiquidityETHOutput, error) {
	out, err := aerodromeRouter.abi.Unpack("addLiquidityETH", data)
	outstruct := new(AddLiquidityETHOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountToken = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountETH = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Liquidity = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackDefaultFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd4b6846d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function defaultFactory() view returns(address)
func (aerodromeRouter *AerodromeRouter) PackDefaultFactory() []byte {
	enc, err := aerodromeRouter.abi.Pack("defaultFactory")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDefaultFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd4b6846d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function defaultFactory() view returns(address)
func (aerodromeRouter *AerodromeRouter) TryPackDefaultFactory() ([]byte, error) {
	return aerodromeRouter.abi.Pack("defaultFactory")
}

// UnpackDefaultFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd4b6846d.
//
// Solidity: function defaultFactory() view returns(address)
func (aerodromeRouter *AerodromeRouter) UnpackDefaultFactory(data []byte) (common.Address, error) {
	out, err := aerodromeRouter.abi.Unpack("defaultFactory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFactoryRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3bf0c9fb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function factoryRegistry() view returns(address)
func (aerodromeRouter *AerodromeRouter) PackFactoryRegistry() []byte {
	enc, err := aerodromeRouter.abi.Pack("factoryRegistry")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFactoryRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3bf0c9fb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function factoryRegistry() view returns(address)
func (aerodromeRouter *AerodromeRouter) TryPackFactoryRegistry() ([]byte, error) {
	return aerodromeRouter.abi.Pack("factoryRegistry")
}

// UnpackFactoryRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3bf0c9fb.
//
// Solidity: function factoryRegistry() view returns(address)
func (aerodromeRouter *AerodromeRouter) UnpackFactoryRegistry(data []byte) (common.Address, error) {
	out, err := aerodromeRouter.abi.Unpack("factoryRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGenerateZapInParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07db50fa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function generateZapInParams(address tokenA, address tokenB, bool stable, address _factory, uint256 amountInA, uint256 amountInB, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (aerodromeRouter *AerodromeRouter) PackGenerateZapInParams(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address, amountInA *big.Int, amountInB *big.Int, routesA []IRouterRoute, routesB []IRouterRoute) []byte {
	enc, err := aerodromeRouter.abi.Pack("generateZapInParams", tokenA, tokenB, stable, factory, amountInA, amountInB, routesA, routesB)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGenerateZapInParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07db50fa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function generateZapInParams(address tokenA, address tokenB, bool stable, address _factory, uint256 amountInA, uint256 amountInB, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (aerodromeRouter *AerodromeRouter) TryPackGenerateZapInParams(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address, amountInA *big.Int, amountInB *big.Int, routesA []IRouterRoute, routesB []IRouterRoute) ([]byte, error) {
	return aerodromeRouter.abi.Pack("generateZapInParams", tokenA, tokenB, stable, factory, amountInA, amountInB, routesA, routesB)
}

// GenerateZapInParamsOutput serves as a container for the return parameters of contract
// method GenerateZapInParams.
type GenerateZapInParamsOutput struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}

// UnpackGenerateZapInParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x07db50fa.
//
// Solidity: function generateZapInParams(address tokenA, address tokenB, bool stable, address _factory, uint256 amountInA, uint256 amountInB, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (aerodromeRouter *AerodromeRouter) UnpackGenerateZapInParams(data []byte) (GenerateZapInParamsOutput, error) {
	out, err := aerodromeRouter.abi.Unpack("generateZapInParams", data)
	outstruct := new(GenerateZapInParamsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountOutMinA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountOutMinB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.AmountAMin = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	outstruct.AmountBMin = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGenerateZapOutParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7539d413.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function generateZapOutParams(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (aerodromeRouter *AerodromeRouter) PackGenerateZapOutParams(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address, liquidity *big.Int, routesA []IRouterRoute, routesB []IRouterRoute) []byte {
	enc, err := aerodromeRouter.abi.Pack("generateZapOutParams", tokenA, tokenB, stable, factory, liquidity, routesA, routesB)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGenerateZapOutParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7539d413.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function generateZapOutParams(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (aerodromeRouter *AerodromeRouter) TryPackGenerateZapOutParams(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address, liquidity *big.Int, routesA []IRouterRoute, routesB []IRouterRoute) ([]byte, error) {
	return aerodromeRouter.abi.Pack("generateZapOutParams", tokenA, tokenB, stable, factory, liquidity, routesA, routesB)
}

// GenerateZapOutParamsOutput serves as a container for the return parameters of contract
// method GenerateZapOutParams.
type GenerateZapOutParamsOutput struct {
	AmountOutMinA *big.Int
	AmountOutMinB *big.Int
	AmountAMin    *big.Int
	AmountBMin    *big.Int
}

// UnpackGenerateZapOutParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7539d413.
//
// Solidity: function generateZapOutParams(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) view returns(uint256 amountOutMinA, uint256 amountOutMinB, uint256 amountAMin, uint256 amountBMin)
func (aerodromeRouter *AerodromeRouter) UnpackGenerateZapOutParams(data []byte) (GenerateZapOutParamsOutput, error) {
	out, err := aerodromeRouter.abi.Unpack("generateZapOutParams", data)
	outstruct := new(GenerateZapOutParamsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountOutMinA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountOutMinB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.AmountAMin = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	outstruct.AmountBMin = abi.ConvertType(out[3], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGetAmountsOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5509a1ac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool,address)[] routes) view returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) PackGetAmountsOut(amountIn *big.Int, routes []IRouterRoute) []byte {
	enc, err := aerodromeRouter.abi.Pack("getAmountsOut", amountIn, routes)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAmountsOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5509a1ac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool,address)[] routes) view returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) TryPackGetAmountsOut(amountIn *big.Int, routes []IRouterRoute) ([]byte, error) {
	return aerodromeRouter.abi.Pack("getAmountsOut", amountIn, routes)
}

// UnpackGetAmountsOut is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5509a1ac.
//
// Solidity: function getAmountsOut(uint256 amountIn, (address,address,bool,address)[] routes) view returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) UnpackGetAmountsOut(data []byte) ([]*big.Int, error) {
	out, err := aerodromeRouter.abi.Unpack("getAmountsOut", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackGetReserves is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8c0037dc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable, address _factory) view returns(uint256 reserveA, uint256 reserveB)
func (aerodromeRouter *AerodromeRouter) PackGetReserves(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address) []byte {
	enc, err := aerodromeRouter.abi.Pack("getReserves", tokenA, tokenB, stable, factory)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetReserves is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8c0037dc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable, address _factory) view returns(uint256 reserveA, uint256 reserveB)
func (aerodromeRouter *AerodromeRouter) TryPackGetReserves(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address) ([]byte, error) {
	return aerodromeRouter.abi.Pack("getReserves", tokenA, tokenB, stable, factory)
}

// GetReservesOutput serves as a container for the return parameters of contract
// method GetReserves.
type GetReservesOutput struct {
	ReserveA *big.Int
	ReserveB *big.Int
}

// UnpackGetReserves is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8c0037dc.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable, address _factory) view returns(uint256 reserveA, uint256 reserveB)
func (aerodromeRouter *AerodromeRouter) UnpackGetReserves(data []byte) (GetReservesOutput, error) {
	out, err := aerodromeRouter.abi.Unpack("getReserves", data)
	outstruct := new(GetReservesOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.ReserveA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.ReserveB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackIsTrustedForwarder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x572b6c05.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (aerodromeRouter *AerodromeRouter) PackIsTrustedForwarder(forwarder common.Address) []byte {
	enc, err := aerodromeRouter.abi.Pack("isTrustedForwarder", forwarder)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsTrustedForwarder is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x572b6c05.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (aerodromeRouter *AerodromeRouter) TryPackIsTrustedForwarder(forwarder common.Address) ([]byte, error) {
	return aerodromeRouter.abi.Pack("isTrustedForwarder", forwarder)
}

// UnpackIsTrustedForwarder is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (aerodromeRouter *AerodromeRouter) UnpackIsTrustedForwarder(data []byte) (bool, error) {
	out, err := aerodromeRouter.abi.Unpack("isTrustedForwarder", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPoolFor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x874029d9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function poolFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (aerodromeRouter *AerodromeRouter) PackPoolFor(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address) []byte {
	enc, err := aerodromeRouter.abi.Pack("poolFor", tokenA, tokenB, stable, factory)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPoolFor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x874029d9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function poolFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (aerodromeRouter *AerodromeRouter) TryPackPoolFor(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address) ([]byte, error) {
	return aerodromeRouter.abi.Pack("poolFor", tokenA, tokenB, stable, factory)
}

// UnpackPoolFor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x874029d9.
//
// Solidity: function poolFor(address tokenA, address tokenB, bool stable, address _factory) view returns(address pool)
func (aerodromeRouter *AerodromeRouter) UnpackPoolFor(data []byte) (common.Address, error) {
	out, err := aerodromeRouter.abi.Unpack("poolFor", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackQuoteAddLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce700c29.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) PackQuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address, amountADesired *big.Int, amountBDesired *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("quoteAddLiquidity", tokenA, tokenB, stable, factory, amountADesired, amountBDesired)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuoteAddLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xce700c29.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) TryPackQuoteAddLiquidity(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address, amountADesired *big.Int, amountBDesired *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("quoteAddLiquidity", tokenA, tokenB, stable, factory, amountADesired, amountBDesired)
}

// QuoteAddLiquidityOutput serves as a container for the return parameters of contract
// method QuoteAddLiquidity.
type QuoteAddLiquidityOutput struct {
	AmountA   *big.Int
	AmountB   *big.Int
	Liquidity *big.Int
}

// UnpackQuoteAddLiquidity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xce700c29.
//
// Solidity: function quoteAddLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 amountADesired, uint256 amountBDesired) view returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) UnpackQuoteAddLiquidity(data []byte) (QuoteAddLiquidityOutput, error) {
	out, err := aerodromeRouter.abi.Unpack("quoteAddLiquidity", data)
	outstruct := new(QuoteAddLiquidityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Liquidity = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackQuoteRemoveLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc92de3ec.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (aerodromeRouter *AerodromeRouter) PackQuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address, liquidity *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("quoteRemoveLiquidity", tokenA, tokenB, stable, factory, liquidity)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuoteRemoveLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc92de3ec.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (aerodromeRouter *AerodromeRouter) TryPackQuoteRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, factory common.Address, liquidity *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("quoteRemoveLiquidity", tokenA, tokenB, stable, factory, liquidity)
}

// QuoteRemoveLiquidityOutput serves as a container for the return parameters of contract
// method QuoteRemoveLiquidity.
type QuoteRemoveLiquidityOutput struct {
	AmountA *big.Int
	AmountB *big.Int
}

// UnpackQuoteRemoveLiquidity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc92de3ec.
//
// Solidity: function quoteRemoveLiquidity(address tokenA, address tokenB, bool stable, address _factory, uint256 liquidity) view returns(uint256 amountA, uint256 amountB)
func (aerodromeRouter *AerodromeRouter) UnpackQuoteRemoveLiquidity(data []byte) (QuoteRemoveLiquidityOutput, error) {
	out, err := aerodromeRouter.abi.Unpack("quoteRemoveLiquidity", data)
	outstruct := new(QuoteRemoveLiquidityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackQuoteStableLiquidityRatio is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5ba53c7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function quoteStableLiquidityRatio(address tokenA, address tokenB, address _factory) view returns(uint256 ratio)
func (aerodromeRouter *AerodromeRouter) PackQuoteStableLiquidityRatio(tokenA common.Address, tokenB common.Address, factory common.Address) []byte {
	enc, err := aerodromeRouter.abi.Pack("quoteStableLiquidityRatio", tokenA, tokenB, factory)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackQuoteStableLiquidityRatio is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5ba53c7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function quoteStableLiquidityRatio(address tokenA, address tokenB, address _factory) view returns(uint256 ratio)
func (aerodromeRouter *AerodromeRouter) TryPackQuoteStableLiquidityRatio(tokenA common.Address, tokenB common.Address, factory common.Address) ([]byte, error) {
	return aerodromeRouter.abi.Pack("quoteStableLiquidityRatio", tokenA, tokenB, factory)
}

// UnpackQuoteStableLiquidityRatio is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf5ba53c7.
//
// Solidity: function quoteStableLiquidityRatio(address tokenA, address tokenB, address _factory) view returns(uint256 ratio)
func (aerodromeRouter *AerodromeRouter) UnpackQuoteStableLiquidityRatio(data []byte) (*big.Int, error) {
	out, err := aerodromeRouter.abi.Unpack("quoteStableLiquidityRatio", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRemoveLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0dede6c4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (aerodromeRouter *AerodromeRouter) PackRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("removeLiquidity", tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0dede6c4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (aerodromeRouter *AerodromeRouter) TryPackRemoveLiquidity(tokenA common.Address, tokenB common.Address, stable bool, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("removeLiquidity", tokenA, tokenB, stable, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityOutput serves as a container for the return parameters of contract
// method RemoveLiquidity.
type RemoveLiquidityOutput struct {
	AmountA *big.Int
	AmountB *big.Int
}

// UnpackRemoveLiquidity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0dede6c4.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, bool stable, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (aerodromeRouter *AerodromeRouter) UnpackRemoveLiquidity(data []byte) (RemoveLiquidityOutput, error) {
	out, err := aerodromeRouter.abi.Unpack("removeLiquidity", data)
	outstruct := new(RemoveLiquidityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountA = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountB = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackRemoveLiquidityETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd7b0e0a5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (aerodromeRouter *AerodromeRouter) PackRemoveLiquidityETH(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("removeLiquidityETH", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidityETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd7b0e0a5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (aerodromeRouter *AerodromeRouter) TryPackRemoveLiquidityETH(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("removeLiquidityETH", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// RemoveLiquidityETHOutput serves as a container for the return parameters of contract
// method RemoveLiquidityETH.
type RemoveLiquidityETHOutput struct {
	AmountToken *big.Int
	AmountETH   *big.Int
}

// UnpackRemoveLiquidityETH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd7b0e0a5.
//
// Solidity: function removeLiquidityETH(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountETH)
func (aerodromeRouter *AerodromeRouter) UnpackRemoveLiquidityETH(data []byte) (RemoveLiquidityETHOutput, error) {
	out, err := aerodromeRouter.abi.Unpack("removeLiquidityETH", data)
	outstruct := new(RemoveLiquidityETHOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AmountToken = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.AmountETH = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackRemoveLiquidityETHSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe411f14.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (aerodromeRouter *AerodromeRouter) PackRemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("removeLiquidityETHSupportingFeeOnTransferTokens", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveLiquidityETHSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe411f14.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (aerodromeRouter *AerodromeRouter) TryPackRemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, stable bool, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("removeLiquidityETHSupportingFeeOnTransferTokens", token, stable, liquidity, amountTokenMin, amountETHMin, to, deadline)
}

// UnpackRemoveLiquidityETHSupportingFeeOnTransferTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfe411f14.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, bool stable, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline) returns(uint256 amountETH)
func (aerodromeRouter *AerodromeRouter) UnpackRemoveLiquidityETHSupportingFeeOnTransferTokens(data []byte) (*big.Int, error) {
	out, err := aerodromeRouter.abi.Unpack("removeLiquidityETHSupportingFeeOnTransferTokens", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSortTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x544caa56.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (aerodromeRouter *AerodromeRouter) PackSortTokens(tokenA common.Address, tokenB common.Address) []byte {
	enc, err := aerodromeRouter.abi.Pack("sortTokens", tokenA, tokenB)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSortTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x544caa56.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (aerodromeRouter *AerodromeRouter) TryPackSortTokens(tokenA common.Address, tokenB common.Address) ([]byte, error) {
	return aerodromeRouter.abi.Pack("sortTokens", tokenA, tokenB)
}

// SortTokensOutput serves as a container for the return parameters of contract
// method SortTokens.
type SortTokensOutput struct {
	Token0 common.Address
	Token1 common.Address
}

// UnpackSortTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (aerodromeRouter *AerodromeRouter) UnpackSortTokens(data []byte) (SortTokensOutput, error) {
	out, err := aerodromeRouter.abi.Unpack("sortTokens", data)
	outstruct := new(SortTokensOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Token0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackSwapExactETHForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x903638a4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) PackSwapExactETHForTokens(amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("swapExactETHForTokens", amountOutMin, routes, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactETHForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x903638a4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) TryPackSwapExactETHForTokens(amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("swapExactETHForTokens", amountOutMin, routes, to, deadline)
}

// UnpackSwapExactETHForTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x903638a4.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) UnpackSwapExactETHForTokens(data []byte) ([]*big.Int, error) {
	out, err := aerodromeRouter.abi.Unpack("swapExactETHForTokens", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackSwapExactETHForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3da5acba.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns()
func (aerodromeRouter *AerodromeRouter) PackSwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, routes, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactETHForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3da5acba.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) payable returns()
func (aerodromeRouter *AerodromeRouter) TryPackSwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, routes, to, deadline)
}

// PackSwapExactTokensForETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc6b7f1b6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) PackSwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("swapExactTokensForETH", amountIn, amountOutMin, routes, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactTokensForETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc6b7f1b6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) TryPackSwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("swapExactTokensForETH", amountIn, amountOutMin, routes, to, deadline)
}

// UnpackSwapExactTokensForETH is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc6b7f1b6.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) UnpackSwapExactTokensForETH(data []byte) ([]*big.Int, error) {
	out, err := aerodromeRouter.abi.Unpack("swapExactTokensForETH", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackSwapExactTokensForETHSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12bc3aca.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (aerodromeRouter *AerodromeRouter) PackSwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactTokensForETHSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12bc3aca.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (aerodromeRouter *AerodromeRouter) TryPackSwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
}

// PackSwapExactTokensForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcac88ea9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) PackSwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("swapExactTokensForTokens", amountIn, amountOutMin, routes, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactTokensForTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcac88ea9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) TryPackSwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("swapExactTokensForTokens", amountIn, amountOutMin, routes, to, deadline)
}

// UnpackSwapExactTokensForTokens is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcac88ea9.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns(uint256[] amounts)
func (aerodromeRouter *AerodromeRouter) UnpackSwapExactTokensForTokens(data []byte) ([]*big.Int, error) {
	out, err := aerodromeRouter.abi.Unpack("swapExactTokensForTokens", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackSwapExactTokensForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x88cd821e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (aerodromeRouter *AerodromeRouter) PackSwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) []byte {
	enc, err := aerodromeRouter.abi.Pack("swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwapExactTokensForTokensSupportingFeeOnTransferTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x88cd821e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, (address,address,bool,address)[] routes, address to, uint256 deadline) returns()
func (aerodromeRouter *AerodromeRouter) TryPackSwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, routes []IRouterRoute, to common.Address, deadline *big.Int) ([]byte, error) {
	return aerodromeRouter.abi.Pack("swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, routes, to, deadline)
}

// PackVoter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46c96aac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function voter() view returns(address)
func (aerodromeRouter *AerodromeRouter) PackVoter() []byte {
	enc, err := aerodromeRouter.abi.Pack("voter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVoter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46c96aac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function voter() view returns(address)
func (aerodromeRouter *AerodromeRouter) TryPackVoter() ([]byte, error) {
	return aerodromeRouter.abi.Pack("voter")
}

// UnpackVoter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (aerodromeRouter *AerodromeRouter) UnpackVoter(data []byte) (common.Address, error) {
	out, err := aerodromeRouter.abi.Unpack("voter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackWeth is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3fc8cef3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function weth() view returns(address)
func (aerodromeRouter *AerodromeRouter) PackWeth() []byte {
	enc, err := aerodromeRouter.abi.Pack("weth")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWeth is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3fc8cef3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function weth() view returns(address)
func (aerodromeRouter *AerodromeRouter) TryPackWeth() ([]byte, error) {
	return aerodromeRouter.abi.Pack("weth")
}

// UnpackWeth is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (aerodromeRouter *AerodromeRouter) UnpackWeth(data []byte) (common.Address, error) {
	out, err := aerodromeRouter.abi.Unpack("weth", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackZapIn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfb49bafd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function zapIn(address tokenIn, uint256 amountInA, uint256 amountInB, (address,address,bool,address,uint256,uint256,uint256,uint256) zapInPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB, address to, bool stake) payable returns(uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) PackZapIn(tokenIn common.Address, amountInA *big.Int, amountInB *big.Int, zapInPool IRouterZap, routesA []IRouterRoute, routesB []IRouterRoute, to common.Address, stake bool) []byte {
	enc, err := aerodromeRouter.abi.Pack("zapIn", tokenIn, amountInA, amountInB, zapInPool, routesA, routesB, to, stake)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackZapIn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfb49bafd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function zapIn(address tokenIn, uint256 amountInA, uint256 amountInB, (address,address,bool,address,uint256,uint256,uint256,uint256) zapInPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB, address to, bool stake) payable returns(uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) TryPackZapIn(tokenIn common.Address, amountInA *big.Int, amountInB *big.Int, zapInPool IRouterZap, routesA []IRouterRoute, routesB []IRouterRoute, to common.Address, stake bool) ([]byte, error) {
	return aerodromeRouter.abi.Pack("zapIn", tokenIn, amountInA, amountInB, zapInPool, routesA, routesB, to, stake)
}

// UnpackZapIn is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfb49bafd.
//
// Solidity: function zapIn(address tokenIn, uint256 amountInA, uint256 amountInB, (address,address,bool,address,uint256,uint256,uint256,uint256) zapInPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB, address to, bool stake) payable returns(uint256 liquidity)
func (aerodromeRouter *AerodromeRouter) UnpackZapIn(data []byte) (*big.Int, error) {
	out, err := aerodromeRouter.abi.Unpack("zapIn", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackZapOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa81b9159.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function zapOut(address tokenOut, uint256 liquidity, (address,address,bool,address,uint256,uint256,uint256,uint256) zapOutPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) returns()
func (aerodromeRouter *AerodromeRouter) PackZapOut(tokenOut common.Address, liquidity *big.Int, zapOutPool IRouterZap, routesA []IRouterRoute, routesB []IRouterRoute) []byte {
	enc, err := aerodromeRouter.abi.Pack("zapOut", tokenOut, liquidity, zapOutPool, routesA, routesB)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackZapOut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa81b9159.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function zapOut(address tokenOut, uint256 liquidity, (address,address,bool,address,uint256,uint256,uint256,uint256) zapOutPool, (address,address,bool,address)[] routesA, (address,address,bool,address)[] routesB) returns()
func (aerodromeRouter *AerodromeRouter) TryPackZapOut(tokenOut common.Address, liquidity *big.Int, zapOutPool IRouterZap, routesA []IRouterRoute, routesB []IRouterRoute) ([]byte, error) {
	return aerodromeRouter.abi.Pack("zapOut", tokenOut, liquidity, zapOutPool, routesA, routesB)
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (aerodromeRouter *AerodromeRouter) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["ETHTransferFailed"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackETHTransferFailedError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["Expired"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InsufficientAmount"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInsufficientAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InsufficientAmountA"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInsufficientAmountAError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InsufficientAmountADesired"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInsufficientAmountADesiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InsufficientAmountAOptimal"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInsufficientAmountAOptimalError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InsufficientAmountB"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInsufficientAmountBError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InsufficientAmountBDesired"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInsufficientAmountBDesiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InsufficientLiquidity"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInsufficientLiquidityError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InsufficientOutputAmount"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInsufficientOutputAmountError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InvalidAmountInForETHDeposit"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInvalidAmountInForETHDepositError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InvalidPath"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInvalidPathError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InvalidRouteA"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInvalidRouteAError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InvalidRouteB"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInvalidRouteBError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["InvalidTokenInForETHDeposit"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackInvalidTokenInForETHDepositError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["OnlyWETH"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackOnlyWETHError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["PoolDoesNotExist"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackPoolDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["PoolFactoryDoesNotExist"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackPoolFactoryDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["SameAddresses"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackSameAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], aerodromeRouter.abi.Errors["ZeroAddress"].ID.Bytes()[:4]) {
		return aerodromeRouter.UnpackZeroAddressError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// AerodromeRouterETHTransferFailed represents a ETHTransferFailed error raised by the AerodromeRouter contract.
type AerodromeRouterETHTransferFailed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ETHTransferFailed()
func AerodromeRouterETHTransferFailedErrorID() common.Hash {
	return common.HexToHash("0xb12d13ebe76e15b5fdb7bf52f0daba617b83ebcc560b0666c44fcdcd71f4362b")
}

// UnpackETHTransferFailedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ETHTransferFailed()
func (aerodromeRouter *AerodromeRouter) UnpackETHTransferFailedError(raw []byte) (*AerodromeRouterETHTransferFailed, error) {
	out := new(AerodromeRouterETHTransferFailed)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "ETHTransferFailed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterExpired represents a Expired error raised by the AerodromeRouter contract.
type AerodromeRouterExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Expired()
func AerodromeRouterExpiredErrorID() common.Hash {
	return common.HexToHash("0x203d82d8d99f63bfecc8335216735e0271df4249ea752b030f9ab305b94e5afe")
}

// UnpackExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Expired()
func (aerodromeRouter *AerodromeRouter) UnpackExpiredError(raw []byte) (*AerodromeRouterExpired, error) {
	out := new(AerodromeRouterExpired)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "Expired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInsufficientAmount represents a InsufficientAmount error raised by the AerodromeRouter contract.
type AerodromeRouterInsufficientAmount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientAmount()
func AerodromeRouterInsufficientAmountErrorID() common.Hash {
	return common.HexToHash("0x5945ea56efb769109dee1ed59908b3ae737ed062a0b113d04594c2dac7318b76")
}

// UnpackInsufficientAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientAmount()
func (aerodromeRouter *AerodromeRouter) UnpackInsufficientAmountError(raw []byte) (*AerodromeRouterInsufficientAmount, error) {
	out := new(AerodromeRouterInsufficientAmount)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InsufficientAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInsufficientAmountA represents a InsufficientAmountA error raised by the AerodromeRouter contract.
type AerodromeRouterInsufficientAmountA struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientAmountA()
func AerodromeRouterInsufficientAmountAErrorID() common.Hash {
	return common.HexToHash("0x8f66ec146bd346df5433ee69855986e98f67306bf016fbb7df873dc70fe8e0bd")
}

// UnpackInsufficientAmountAError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientAmountA()
func (aerodromeRouter *AerodromeRouter) UnpackInsufficientAmountAError(raw []byte) (*AerodromeRouterInsufficientAmountA, error) {
	out := new(AerodromeRouterInsufficientAmountA)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InsufficientAmountA", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInsufficientAmountADesired represents a InsufficientAmountADesired error raised by the AerodromeRouter contract.
type AerodromeRouterInsufficientAmountADesired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientAmountADesired()
func AerodromeRouterInsufficientAmountADesiredErrorID() common.Hash {
	return common.HexToHash("0xdc6b2ef2a8f4532aa7caa0f91aacf4caaaaebd43f82df98ff18413fcc8545f95")
}

// UnpackInsufficientAmountADesiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientAmountADesired()
func (aerodromeRouter *AerodromeRouter) UnpackInsufficientAmountADesiredError(raw []byte) (*AerodromeRouterInsufficientAmountADesired, error) {
	out := new(AerodromeRouterInsufficientAmountADesired)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InsufficientAmountADesired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInsufficientAmountAOptimal represents a InsufficientAmountAOptimal error raised by the AerodromeRouter contract.
type AerodromeRouterInsufficientAmountAOptimal struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientAmountAOptimal()
func AerodromeRouterInsufficientAmountAOptimalErrorID() common.Hash {
	return common.HexToHash("0xfe496df65555c7d5aed7aef9ceeb5f1493c7fdc410f450e4a5b87ffa226265b1")
}

// UnpackInsufficientAmountAOptimalError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientAmountAOptimal()
func (aerodromeRouter *AerodromeRouter) UnpackInsufficientAmountAOptimalError(raw []byte) (*AerodromeRouterInsufficientAmountAOptimal, error) {
	out := new(AerodromeRouterInsufficientAmountAOptimal)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InsufficientAmountAOptimal", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInsufficientAmountB represents a InsufficientAmountB error raised by the AerodromeRouter contract.
type AerodromeRouterInsufficientAmountB struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientAmountB()
func AerodromeRouterInsufficientAmountBErrorID() common.Hash {
	return common.HexToHash("0x34c906248d67a78152acc7acf0aa395e3ffe5a09bd503e8aeb4381022329773a")
}

// UnpackInsufficientAmountBError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientAmountB()
func (aerodromeRouter *AerodromeRouter) UnpackInsufficientAmountBError(raw []byte) (*AerodromeRouterInsufficientAmountB, error) {
	out := new(AerodromeRouterInsufficientAmountB)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InsufficientAmountB", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInsufficientAmountBDesired represents a InsufficientAmountBDesired error raised by the AerodromeRouter contract.
type AerodromeRouterInsufficientAmountBDesired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientAmountBDesired()
func AerodromeRouterInsufficientAmountBDesiredErrorID() common.Hash {
	return common.HexToHash("0xacee051367ba2af7da207ccde359f48b8c6bc9b2df743c2b572b62d4c618290c")
}

// UnpackInsufficientAmountBDesiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientAmountBDesired()
func (aerodromeRouter *AerodromeRouter) UnpackInsufficientAmountBDesiredError(raw []byte) (*AerodromeRouterInsufficientAmountBDesired, error) {
	out := new(AerodromeRouterInsufficientAmountBDesired)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InsufficientAmountBDesired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInsufficientLiquidity represents a InsufficientLiquidity error raised by the AerodromeRouter contract.
type AerodromeRouterInsufficientLiquidity struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientLiquidity()
func AerodromeRouterInsufficientLiquidityErrorID() common.Hash {
	return common.HexToHash("0xbb55fd27c46b5ba9f88ff2cb2222216afeb0f193423b26615497b3020ab61f8e")
}

// UnpackInsufficientLiquidityError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientLiquidity()
func (aerodromeRouter *AerodromeRouter) UnpackInsufficientLiquidityError(raw []byte) (*AerodromeRouterInsufficientLiquidity, error) {
	out := new(AerodromeRouterInsufficientLiquidity)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InsufficientLiquidity", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInsufficientOutputAmount represents a InsufficientOutputAmount error raised by the AerodromeRouter contract.
type AerodromeRouterInsufficientOutputAmount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientOutputAmount()
func AerodromeRouterInsufficientOutputAmountErrorID() common.Hash {
	return common.HexToHash("0x42301c237a29c3e090da1b147e7b504f08db56c9600985b34af188477d350fa7")
}

// UnpackInsufficientOutputAmountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientOutputAmount()
func (aerodromeRouter *AerodromeRouter) UnpackInsufficientOutputAmountError(raw []byte) (*AerodromeRouterInsufficientOutputAmount, error) {
	out := new(AerodromeRouterInsufficientOutputAmount)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InsufficientOutputAmount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInvalidAmountInForETHDeposit represents a InvalidAmountInForETHDeposit error raised by the AerodromeRouter contract.
type AerodromeRouterInvalidAmountInForETHDeposit struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAmountInForETHDeposit()
func AerodromeRouterInvalidAmountInForETHDepositErrorID() common.Hash {
	return common.HexToHash("0x70a3fb92ea4aa1c50dc3977e4b7b97d2fa8209b030b0eb8f4530ef6f728b273b")
}

// UnpackInvalidAmountInForETHDepositError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAmountInForETHDeposit()
func (aerodromeRouter *AerodromeRouter) UnpackInvalidAmountInForETHDepositError(raw []byte) (*AerodromeRouterInvalidAmountInForETHDeposit, error) {
	out := new(AerodromeRouterInvalidAmountInForETHDeposit)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InvalidAmountInForETHDeposit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInvalidPath represents a InvalidPath error raised by the AerodromeRouter contract.
type AerodromeRouterInvalidPath struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPath()
func AerodromeRouterInvalidPathErrorID() common.Hash {
	return common.HexToHash("0x20db826769e3af8321787ba1a381ae0528298df9eba27cc33194b460ff8a3602")
}

// UnpackInvalidPathError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPath()
func (aerodromeRouter *AerodromeRouter) UnpackInvalidPathError(raw []byte) (*AerodromeRouterInvalidPath, error) {
	out := new(AerodromeRouterInvalidPath)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InvalidPath", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInvalidRouteA represents a InvalidRouteA error raised by the AerodromeRouter contract.
type AerodromeRouterInvalidRouteA struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRouteA()
func AerodromeRouterInvalidRouteAErrorID() common.Hash {
	return common.HexToHash("0x4ea0e338ad492f5b4e51c4be1a94f2de7c7766a3e72c24dd03fc629d2101cb07")
}

// UnpackInvalidRouteAError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRouteA()
func (aerodromeRouter *AerodromeRouter) UnpackInvalidRouteAError(raw []byte) (*AerodromeRouterInvalidRouteA, error) {
	out := new(AerodromeRouterInvalidRouteA)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InvalidRouteA", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInvalidRouteB represents a InvalidRouteB error raised by the AerodromeRouter contract.
type AerodromeRouterInvalidRouteB struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRouteB()
func AerodromeRouterInvalidRouteBErrorID() common.Hash {
	return common.HexToHash("0xcac9040c66a09bad7e682f1905903e8f694918e8d1fe59ffb29a2bd3a5152e1f")
}

// UnpackInvalidRouteBError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRouteB()
func (aerodromeRouter *AerodromeRouter) UnpackInvalidRouteBError(raw []byte) (*AerodromeRouterInvalidRouteB, error) {
	out := new(AerodromeRouterInvalidRouteB)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InvalidRouteB", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterInvalidTokenInForETHDeposit represents a InvalidTokenInForETHDeposit error raised by the AerodromeRouter contract.
type AerodromeRouterInvalidTokenInForETHDeposit struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidTokenInForETHDeposit()
func AerodromeRouterInvalidTokenInForETHDepositErrorID() common.Hash {
	return common.HexToHash("0xae6d566f54709b9c86ec4796d9b738ec7a514d1e64576e65118c62438d41cd81")
}

// UnpackInvalidTokenInForETHDepositError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidTokenInForETHDeposit()
func (aerodromeRouter *AerodromeRouter) UnpackInvalidTokenInForETHDepositError(raw []byte) (*AerodromeRouterInvalidTokenInForETHDeposit, error) {
	out := new(AerodromeRouterInvalidTokenInForETHDeposit)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "InvalidTokenInForETHDeposit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterOnlyWETH represents a OnlyWETH error raised by the AerodromeRouter contract.
type AerodromeRouterOnlyWETH struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyWETH()
func AerodromeRouterOnlyWETHErrorID() common.Hash {
	return common.HexToHash("0x01f180c9ae17cb7cc92ad1b0413ba1cc7b9ad02bd1c5029c7c87f2b8d3f8eb29")
}

// UnpackOnlyWETHError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyWETH()
func (aerodromeRouter *AerodromeRouter) UnpackOnlyWETHError(raw []byte) (*AerodromeRouterOnlyWETH, error) {
	out := new(AerodromeRouterOnlyWETH)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "OnlyWETH", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterPoolDoesNotExist represents a PoolDoesNotExist error raised by the AerodromeRouter contract.
type AerodromeRouterPoolDoesNotExist struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PoolDoesNotExist()
func AerodromeRouterPoolDoesNotExistErrorID() common.Hash {
	return common.HexToHash("0x9c8787c02ab43cbce7796a08dac3b3f4309a6bc09e8b8f916a07724aca347119")
}

// UnpackPoolDoesNotExistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PoolDoesNotExist()
func (aerodromeRouter *AerodromeRouter) UnpackPoolDoesNotExistError(raw []byte) (*AerodromeRouterPoolDoesNotExist, error) {
	out := new(AerodromeRouterPoolDoesNotExist)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "PoolDoesNotExist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterPoolFactoryDoesNotExist represents a PoolFactoryDoesNotExist error raised by the AerodromeRouter contract.
type AerodromeRouterPoolFactoryDoesNotExist struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PoolFactoryDoesNotExist()
func AerodromeRouterPoolFactoryDoesNotExistErrorID() common.Hash {
	return common.HexToHash("0x9a73ab4679c4e32da77911cbda015e7dc4ec96659161fdd4a07a8bab63f9710d")
}

// UnpackPoolFactoryDoesNotExistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PoolFactoryDoesNotExist()
func (aerodromeRouter *AerodromeRouter) UnpackPoolFactoryDoesNotExistError(raw []byte) (*AerodromeRouterPoolFactoryDoesNotExist, error) {
	out := new(AerodromeRouterPoolFactoryDoesNotExist)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "PoolFactoryDoesNotExist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterSameAddresses represents a SameAddresses error raised by the AerodromeRouter contract.
type AerodromeRouterSameAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SameAddresses()
func AerodromeRouterSameAddressesErrorID() common.Hash {
	return common.HexToHash("0xca57cff4baca38e2fe41e555c4f9633eb8a0fe0c15135ed0d64322d294b4932b")
}

// UnpackSameAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SameAddresses()
func (aerodromeRouter *AerodromeRouter) UnpackSameAddressesError(raw []byte) (*AerodromeRouterSameAddresses, error) {
	out := new(AerodromeRouterSameAddresses)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "SameAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// AerodromeRouterZeroAddress represents a ZeroAddress error raised by the AerodromeRouter contract.
type AerodromeRouterZeroAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ZeroAddress()
func AerodromeRouterZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xd92e233df2717d4a40030e20904abd27b68fcbeede117eaaccbbdac9618c8c73")
}

// UnpackZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ZeroAddress()
func (aerodromeRouter *AerodromeRouter) UnpackZeroAddressError(raw []byte) (*AerodromeRouterZeroAddress, error) {
	out := new(AerodromeRouterZeroAddress)
	if err := aerodromeRouter.abi.UnpackIntoInterface(out, "ZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}
