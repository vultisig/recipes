package thorchain

import (
	"errors"
	"fmt"
	"strings"

	"github.com/vultisig/vultisig-go/common"
)

type metadata struct {
	asset string
}

// asset names mapping is static, we shouldn't call any external API in metarule
var pools = []metadata{
	{asset: "AVAX.AVAX"},
	{asset: "AVAX.SOL-0XFE6B19286885A4F7F55ADAD09C3CD1F906D2478F"},
	{asset: "AVAX.USDC-0XB97EF9EF8734C71904D8002F8B6BC66DD9C48A6E"},
	{asset: "AVAX.USDT-0X9702230A8EA53601F5CD2DC00FDBC13D4DF4A8C7"},
	{asset: "BASE.CBBTC-0XCBB7C0000AB88B473B1F5AFD9EF808440EED33BF"},
	{asset: "BASE.ETH"},
	{asset: "BASE.USDC-0X833589FCD6EDB6E08F4C7C32D4F71B54BDA02913"},
	{asset: "BASE.VVV-0XACFE6019ED1A7DC6F7B508C02D1B04EC88CC21BF"},
	{asset: "BCH.BCH"},
	{asset: "BSC.BNB"},
	{asset: "BSC.BTCB-0X7130D2A12B9BCBFAE4F2634D864A1EE1CE3EAD9C"},
	{asset: "BSC.BUSD-0XE9E7CEA3DEDCA5984780BAFC599BD69ADD087D56"},
	{asset: "BSC.ETH-0X2170ED0880AC9A755FD29B2688956BD959F933F8"},
	{asset: "BSC.TWT-0X4B0F1812E5DF2A09796481FF14017E6005508003"},
	{asset: "BSC.USDC-0X8AC76A51CC950D9822D68B83FE1AD97B32CD580D"},
	{asset: "BSC.USDT-0X55D398326F99059FF775485246999027B3197955"},
	{asset: "BTC.BTC"},
	{asset: "DOGE.DOGE"},
	{asset: "ETH.AAVE-0X7FC66500C84A76AD7E9C93437BFC5AC33E2DDAE9"},
	{asset: "ETH.DAI-0X6B175474E89094C44DA98B954EEDEAC495271D0F"},
	{asset: "ETH.DPI-0X1494CA1F11D487C2BBE4543E90080AEBA4BA3C2B"},
	{asset: "ETH.ETH"},
	{asset: "ETH.FOX-0XC770EEFAD204B5180DF6A14EE197D99D808EE52D"},
	{asset: "ETH.GUSD-0X056FD409E1D7A124BD7017459DFEA2F387B6D5CD"},
	{asset: "ETH.LINK-0X514910771AF9CA656AF840DFF83E8264ECF986CA"},
	{asset: "ETH.LUSD-0X5F98805A4E8BE255A32880FDEC7F6728C6568BA0"},
	{asset: "ETH.RAZE-0X5EAA69B29F99C84FE5DE8200340B4E9B4AB38EAC"},
	{asset: "ETH.SNX-0XC011A73EE8576FB46F5E1C5751CA3B9FE0AF2A6F"},
	{asset: "ETH.TGT-0X108A850856DB3F85D0269A2693D896B394C80325"},
	{asset: "ETH.THOR-0XA5F2211B9B8170F694421F2046281775E8468044"},
	{asset: "ETH.USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48"},
	{asset: "ETH.USDP-0X8E870D67F660D95D5BE530380D0EC0BD388289E1"},
	{asset: "ETH.USDT-0XDAC17F958D2EE523A2206206994597C13D831EC7"},
	{asset: "ETH.VTHOR-0X815C23ECA83261B6EC689B60CC4A58B54BC24D8D"},
	{asset: "ETH.WBTC-0X2260FAC5E5542A773AA44FBCFEDF7C193BC2C599"},
	{asset: "ETH.XDEFI-0X72B886D09C117654AB7DA13A14D603001DE0B777"},
	{asset: "ETH.XRUNE-0X69FA0FEE221AD11012BAB0FDB45D444D3D2CE71C"},
	{asset: "ETH.YFI-0X0BC529C00C6401AEF6D220BE8C6EA1667F6AD93E"},
	{asset: "GAIA.ATOM"},
	{asset: "LTC.LTC"},
	{asset: "THOR.RUNE"},
	{asset: "THOR.NAMI"},
	{asset: "THOR.RUJI"},
	{asset: "THOR.TCY"},
	{asset: "XRP.XRP"},
	{asset: "TRON.TRX"},
}

type network string

const (
	btc       network = "BTC"
	eth       network = "ETH"
	bsc       network = "BSC"
	base      network = "BASE"
	avax      network = "AVAX"
	xrp       network = "XRP"
	thorchain network = "THOR"
	gaia      network = "GAIA"
	tron      network = "TRON"
)

func parseNetwork(c common.Chain) (network, error) {
	switch c {
	case common.Bitcoin:
		return btc, nil
	case common.Ethereum:
		return eth, nil
	case common.BscChain:
		return bsc, nil
	case common.Base:
		return base, nil
	case common.Avalanche:
		return avax, nil
	case common.XRP:
		return xrp, nil
	case common.THORChain:
		return thorchain, nil
	case common.GaiaChain:
		return gaia, nil
	case common.Tron:
		return tron, nil
	default:
		return "", errors.New("unknown chain")
	}
}

func MakeAsset(chain common.Chain, asset string) (string, error) {
	thorNet, err := parseNetwork(chain)
	if err != nil {
		return "", fmt.Errorf("unsupported chain: %w", err)
	}

	// Check if asset is native token
	if asset == "" {
		// Native token format: Network.TokenSymbol (e.g., AVAX.AVAX)
		nativeSymbol, er := chain.NativeSymbol()
		if er != nil {
			return "", fmt.Errorf("failed to get native symbol for chain %s: %w", chain, er)
		}
		return string(thorNet) + "." + nativeSymbol, nil
	}

	networkPrefix := string(thorNet) + "."
	targetAsset := strings.ToUpper(asset)

	for _, pp := range pools {
		// Check if pool belongs to our network
		if !strings.HasPrefix(pp.asset, networkPrefix) {
			continue
		}

		// Split the asset into parts: Network.TokenSymbol-Asset
		parts := strings.Split(pp.asset, ".")
		if len(parts) != 2 {
			continue
		}

		// Check if the second part contains a dash (indicating token with address)
		tokenPart := parts[1]
		if !strings.Contains(tokenPart, "-") {
			// If it doesnt have a dash, try to match as native asset
			if strings.EqualFold(tokenPart, targetAsset) {
				return pp.asset, nil
			}
			continue
		}

		// Split token part: TokenSymbol-Address
		tokenParts := strings.Split(tokenPart, "-")
		if len(tokenParts) != 2 {
			continue
		}

		// Check if the address matches (case-insensitive)
		poolAddress := tokenParts[1]
		if strings.EqualFold(poolAddress, targetAsset) {
			return parts[0] + "." + tokenParts[0], nil
		}
	}

	return "", fmt.Errorf("asset not found in THORChain pools for chain %s and asset %s", thorNet, asset)
}

// ShortCode returns the short code for the asset.
func ShortCode(asset string) string {
	switch asset {
	case "THOR.RUNE":
		return "r"
	case "BTC.BTC":
		return "b"
	case "ETH.ETH":
		return "e"
	case "GAIA.ATOM":
		return "g"
	case "DOGE.DOGE":
		return "d"
	case "LTC.LTC":
		return "l"
	case "BCH.BCH":
		return "c"
	case "AVAX.AVAX":
		return "a"
	case "BSC.BNB":
		return "s"
	case "BASE.ETH":
		return "f"
	case "TRON.TRX":
		return "tr"
	case "XRP.XRP":
		return "x"
	case "SOL.SOL":
		return "o"
	default:
		return ""
	}
}
