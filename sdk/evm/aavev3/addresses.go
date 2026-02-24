package aavev3

import (
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type Deployment struct {
	Pool         ethcommon.Address
	DataProvider ethcommon.Address
}

var deployments = map[uint64]Deployment{
	1: {
		Pool:         ethcommon.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"),
		DataProvider: ethcommon.HexToAddress("0x0a16f2FCC0D44FaE41cc54e079281D84A363bECD"),
	},
	10: {
		Pool:         ethcommon.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD"),
		DataProvider: ethcommon.HexToAddress("0x243Aa95cAC2a25651eda86e80bEe66114413c43b"),
	},
	56: {
		Pool:         ethcommon.HexToAddress("0x6807dc923806fE8Fd134338EABCA509979a7e0cB"),
		DataProvider: ethcommon.HexToAddress("0xc90Df74A7c16245c5F5C5870327Ceb38Fe5d5328"),
	},
	137: {
		Pool:         ethcommon.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD"),
		DataProvider: ethcommon.HexToAddress("0x243Aa95cAC2a25651eda86e80bEe66114413c43b"),
	},
	324: {
		Pool:         ethcommon.HexToAddress("0x78e30497a3c7527d953c6B1E3541b021A98Ac43c"),
		DataProvider: ethcommon.HexToAddress("0x9057ac7b2D35606F8AD5aE2FCBafcD94E58D9927"),
	},
	5000: {
		Pool:         ethcommon.HexToAddress("0xCFbFa83332bB1A3154FA4BA4febedf5c94bDA7c0"),
		DataProvider: ethcommon.HexToAddress("0x487c5c669D9eee6057C44973207101276cf73b68"),
	},
	8453: {
		Pool:         ethcommon.HexToAddress("0xA238Dd80C259a72e81d7e4664a9801593F98d1c5"),
		DataProvider: ethcommon.HexToAddress("0x0F43731EB8d45A581f4a36DD74F5f358bc90C73A"),
	},
	42161: {
		Pool:         ethcommon.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD"),
		DataProvider: ethcommon.HexToAddress("0x243Aa95cAC2a25651eda86e80bEe66114413c43b"),
	},
	43114: {
		Pool:         ethcommon.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD"),
		DataProvider: ethcommon.HexToAddress("0x243Aa95cAC2a25651eda86e80bEe66114413c43b"),
	},
}

func GetDeployment(chainID *big.Int) (Deployment, bool) {
	d, ok := deployments[chainID.Uint64()]
	return d, ok
}

func SupportedChainIDs() []uint64 {
	ids := make([]uint64, 0, len(deployments))
	for id := range deployments {
		ids = append(ids, id)
	}
	return ids
}
