package resolver

import (
	"fmt"

	"github.com/vultisig/vultisig-go/common"
)

func getThorChainSymbol(chain common.Chain) (string, error) {
	switch chain {
	case common.Bitcoin:
		return "BTC", nil
	case common.Ethereum:
		return "ETH", nil
	case common.Avalanche:
		return "AVAX", nil
	case common.BscChain:
		return "BSC", nil
	case common.Base:
		return "BASE", nil
	case common.Litecoin:
		return "LTC", nil
	case common.Dogecoin:
		return "DOGE", nil
	case common.BitcoinCash:
		return "BCH", nil
	case common.XRP:
		return "XRP", nil
	case common.GaiaChain:
		return "GAIA", nil
	case common.Tron:
		return "TRON", nil
	case common.Zcash:
		return "ZEC", nil
	default:
		return "", fmt.Errorf("chain %s not supported by ThorChain", chain.String())
	}
}
