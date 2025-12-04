package mayachain

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
// MAYAChain supported pools
var pools = []metadata{
	{asset: "ARB.ETH"},
	{asset: "ARB.USDC-0XAF88D065E77C8CC2239327C5EDB3A432268E5831"},
	{asset: "ARB.USDT-0XFD086BC7CD5C481DCC9C85EBE478A1C0B69FCBB9"},
	{asset: "ARB.WBTC-0X2F2A2543B76A4166549F7AAB2E75BEF0AEFC5B0F"},
	{asset: "ARB.WSTETH-0X5979D7B546E38E414F7E9822514BE443A4800529"},
	{asset: "BTC.BTC"},
	{asset: "DASH.DASH"},
	{asset: "ETH.ETH"},
	{asset: "ETH.USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48"},
	{asset: "ETH.USDT-0XDAC17F958D2EE523A2206206994597C13D831EC7"},
	{asset: "ETH.WSTETH-0X7F39C581F595B53C5CB19BD0B3F8DA6C935E2CA0"},
	{asset: "KUJI.KUJI"},
	{asset: "KUJI.USK"},
	{asset: "MAYA.CACAO"},
	{asset: "THOR.RUNE"},
	{asset: "XRD.XRD"},
	{asset: "ZEC.ZEC"},
}

type network string

const (
	btc       network = "BTC"
	eth       network = "ETH"
	arb       network = "ARB"
	kuji      network = "KUJI"
	dash      network = "DASH"
	thor      network = "THOR"
	maya      network = "MAYA"
	xrd       network = "XRD"
	zec       network = "ZEC"
)

func parseNetwork(c common.Chain) (network, error) {
	switch c {
	case common.Bitcoin:
		return btc, nil
	case common.Ethereum:
		return eth, nil
	case common.Arbitrum:
		return arb, nil
	case common.Kujira:
		return kuji, nil
	case common.Dash:
		return dash, nil
	case common.THORChain:
		return thor, nil
	case common.MayaChain:
		return maya, nil
	case common.Zcash:
		return zec, nil
	default:
		return "", errors.New("unknown chain")
	}
}

func MakeAsset(chain common.Chain, asset string) (string, error) {
	mayaNet, err := parseNetwork(chain)
	if err != nil {
		return "", fmt.Errorf("unsupported chain: %w", err)
	}

	// Check if asset is native token
	if asset == "" {
		nativeSymbol, er := chain.NativeSymbol()
		if er != nil {
			return "", fmt.Errorf("failed to get native symbol for chain %s: %w", chain, er)
		}
		return string(mayaNet) + "." + nativeSymbol, nil
	}

	networkPrefix := string(mayaNet) + "."
	targetAsset := strings.ToUpper(asset)

	for _, pp := range pools {
		if !strings.HasPrefix(pp.asset, networkPrefix) {
			continue
		}

		parts := strings.Split(pp.asset, ".")
		if len(parts) != 2 {
			continue
		}

		tokenPart := parts[1]
		if !strings.Contains(tokenPart, "-") {
			if strings.EqualFold(tokenPart, targetAsset) {
				return pp.asset, nil
			}
			continue
		}

		tokenParts := strings.Split(tokenPart, "-")
		if len(tokenParts) != 2 {
			continue
		}

		poolAddress := tokenParts[1]
		if strings.EqualFold(poolAddress, targetAsset) {
			return parts[0] + "." + tokenParts[0], nil
		}
	}

	return "", fmt.Errorf("asset not found in MAYAChain pools for chain %s and asset %s", mayaNet, asset)
}

// ShortCode returns the short code for the asset on MAYAChain.
func ShortCode(asset string) string {
	switch asset {
	case "MAYA.CACAO":
		return "c"
	case "THOR.RUNE":
		return "r"
	case "BTC.BTC":
		return "b"
	case "ETH.ETH":
		return "e"
	case "DASH.DASH":
		return "d"
	case "KUJI.KUJI":
		return "k"
	case "XRD.XRD":
		return "x"
	case "ZEC.ZEC":
		return "z"
	case "ARB.ETH":
		return "a"
	default:
		return ""
	}
}

