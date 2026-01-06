package evm

import (
	"fmt"

	"github.com/vultisig/vultisig-go/common"
)

// DeFi protocol addresses per chain

// UniswapV3 NonfungiblePositionManager addresses
var UniswapV3PositionManager = map[common.Chain]string{
	common.Ethereum:  "0xC36442b4a4522E871399CD717aBDD847Ab11FE88",
	common.Arbitrum:  "0xC36442b4a4522E871399CD717aBDD847Ab11FE88",
	common.Optimism:  "0xC36442b4a4522E871399CD717aBDD847Ab11FE88",
	common.Base:      "0x03a520b32C04BF3bEEf7BEb72E919cf822Ed34f1",
	common.Polygon:   "0xC36442b4a4522E871399CD717aBDD847Ab11FE88",
	common.BscChain:  "0x7b8A01B39D58278b5DE7e48c8449c9f4F5170613",
	common.Avalanche: "0x655C406EBFa14EE2006250925e54ec43AD184f8B",
}

// AAVE V3 Pool addresses
var AaveV3Pool = map[common.Chain]string{
	common.Ethereum:  "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2",
	common.Arbitrum:  "0x794a61358D6845594F94dc1DB02A252b5b4814aD",
	common.Optimism:  "0x794a61358D6845594F94dc1DB02A252b5b4814aD",
	common.Base:      "0xA238Dd80C259a72e81d7e4664a9801593F98d1c5",
	common.Polygon:   "0x794a61358D6845594F94dc1DB02A252b5b4814aD",
	common.Avalanche: "0x794a61358D6845594F94dc1DB02A252b5b4814aD",
}

// Compound V3 Comet (USDC market) addresses
var CompoundV3CometUSDC = map[common.Chain]string{
	common.Ethereum: "0xc3d688B66703497DAA19211EEdff47f25384cdc3",
	common.Arbitrum: "0xA5EDBDD9646f8dFF606d7448e414884C7d905dCA",
	common.Base:     "0xb125E6687d4313864e53df431d5425969c15Eb2F",
	common.Polygon:  "0xF25212E676D1F7F89Cd72fFEe66158f541246445",
}

// Compound V3 Comet (WETH market) addresses
var CompoundV3CometWETH = map[common.Chain]string{
	common.Ethereum: "0xA17581A9E3356d9A858b789D68B4d866e593aE94",
}

// GMX V2 Exchange Router addresses
var GMXV2ExchangeRouter = map[common.Chain]string{
	common.Arbitrum:  "0x7C68C7866A64FA2160F78EEaE12217FFbf871fa8",
	common.Avalanche: "0x79be2F4eC8A4143BaF963206cF133f3710856D0a",
}

// GMX V2 Order Vault addresses
var GMXV2OrderVault = map[common.Chain]string{
	common.Arbitrum:  "0x31eF83a530Fde1B38EE9A18093A333D8Bbbc40D5",
	common.Avalanche: "0xD3D60D22d415aD43b7E64b510b2B9B8B9Dc3C9E6",
}

// Hyperliquid L1 Bridge address (Arbitrum bridge for deposits/withdrawals)
// Note: Hyperliquid is a custom L1 - this is the Arbitrum bridge contract
var HyperliquidBridge = map[common.Chain]string{
	common.Arbitrum: "0x2Df1c51E09aECF9cacB7bc98cB1742757f163dF7",
}

// GetUniswapV3PositionManager returns the UniswapV3 Position Manager address for a chain
func GetUniswapV3PositionManager(chain common.Chain) (string, error) {
	addr, ok := UniswapV3PositionManager[chain]
	if !ok {
		return "", fmt.Errorf("UniswapV3 Position Manager not available on chain: %s", chain.String())
	}
	return addr, nil
}

// GetAaveV3Pool returns the AAVE V3 Pool address for a chain
func GetAaveV3Pool(chain common.Chain) (string, error) {
	addr, ok := AaveV3Pool[chain]
	if !ok {
		return "", fmt.Errorf("AAVE V3 Pool not available on chain: %s", chain.String())
	}
	return addr, nil
}

// GetCompoundV3Comet returns the Compound V3 Comet address for a chain and market
func GetCompoundV3Comet(chain common.Chain, market string) (string, error) {
	var registry map[common.Chain]string
	switch market {
	case "USDC", "usdc":
		registry = CompoundV3CometUSDC
	case "WETH", "weth":
		registry = CompoundV3CometWETH
	default:
		return "", fmt.Errorf("unknown Compound V3 market: %s", market)
	}

	addr, ok := registry[chain]
	if !ok {
		return "", fmt.Errorf("compound V3 %s market not available on chain: %s", market, chain.String())
	}
	return addr, nil
}

// GetGMXV2ExchangeRouter returns the GMX V2 Exchange Router address for a chain
func GetGMXV2ExchangeRouter(chain common.Chain) (string, error) {
	addr, ok := GMXV2ExchangeRouter[chain]
	if !ok {
		return "", fmt.Errorf("GMX V2 Exchange Router not available on chain: %s", chain.String())
	}
	return addr, nil
}

// GetGMXV2OrderVault returns the GMX V2 Order Vault address for a chain
func GetGMXV2OrderVault(chain common.Chain) (string, error) {
	addr, ok := GMXV2OrderVault[chain]
	if !ok {
		return "", fmt.Errorf("GMX V2 Order Vault not available on chain: %s", chain.String())
	}
	return addr, nil
}

// GetHyperliquidBridge returns the Hyperliquid bridge address for a chain
func GetHyperliquidBridge(chain common.Chain) (string, error) {
	addr, ok := HyperliquidBridge[chain]
	if !ok {
		return "", fmt.Errorf("hyperliquid bridge not available on chain: %s", chain.String())
	}
	return addr, nil
}

// Polymarket CTF Exchange address (Polygon only)
// This is the main exchange contract for trading conditional tokens
var PolymarketCTFExchange = map[common.Chain]string{
	common.Polygon: "0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E",
}

// Polymarket Neg Risk CTF Exchange address (Polygon only)
// Used for negative risk markets (multi-outcome markets)
var PolymarketNegRiskCTFExchange = map[common.Chain]string{
	common.Polygon: "0xC5d563A36AE78145C45a50134d48A1215220f80a",
}

// Polymarket Conditional Token Framework (CTF) address
var PolymarketCTF = map[common.Chain]string{
	common.Polygon: "0x4D97DCd97eC945f40cF65F87097ACe5EA0476045",
}

// Polymarket USDC.e collateral token address
var PolymarketUSDC = map[common.Chain]string{
	common.Polygon: "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174",
}

// GetPolymarketCTFExchange returns the Polymarket CTF Exchange address for a chain
func GetPolymarketCTFExchange(chain common.Chain) (string, error) {
	addr, ok := PolymarketCTFExchange[chain]
	if !ok {
		return "", fmt.Errorf("polymarket CTF Exchange not available on chain: %s", chain.String())
	}
	return addr, nil
}

// GetPolymarketNegRiskCTFExchange returns the Polymarket Neg Risk CTF Exchange address for a chain
func GetPolymarketNegRiskCTFExchange(chain common.Chain) (string, error) {
	addr, ok := PolymarketNegRiskCTFExchange[chain]
	if !ok {
		return "", fmt.Errorf("polymarket Neg Risk CTF Exchange not available on chain: %s", chain.String())
	}
	return addr, nil
}

// GetPolymarketCTF returns the Polymarket CTF address for a chain
func GetPolymarketCTF(chain common.Chain) (string, error) {
	addr, ok := PolymarketCTF[chain]
	if !ok {
		return "", fmt.Errorf("polymarket CTF not available on chain: %s", chain.String())
	}
	return addr, nil
}

// GetPolymarketUSDC returns the Polymarket USDC address for a chain
func GetPolymarketUSDC(chain common.Chain) (string, error) {
	addr, ok := PolymarketUSDC[chain]
	if !ok {
		return "", fmt.Errorf("polymarket USDC not available on chain: %s", chain.String())
	}
	return addr, nil
}

