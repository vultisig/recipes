package aavev3

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	pool "github.com/vultisig/recipes/sdk/evm/codegen/aavev3_pool"
	dp "github.com/vultisig/recipes/sdk/evm/codegen/aavev3_dataprovider"
)

var MaxUint256 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))

type UserAccountData struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	LTV                         *big.Int
	HealthFactor                *big.Int
}

type ReserveData struct {
	LiquidityRate      *big.Int
	VariableBorrowRate *big.Int
}

type ReserveConfigData struct {
	Decimals                 *big.Int
	LTV                     *big.Int
	LiquidationThreshold    *big.Int
	LiquidationBonus        *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	IsActive                 bool
	IsFrozen                 bool
}

func userAccountDataFromCodegen(out pool.GetUserAccountDataOutput) *UserAccountData {
	return &UserAccountData{
		TotalCollateralBase:         out.TotalCollateralBase,
		TotalDebtBase:               out.TotalDebtBase,
		AvailableBorrowsBase:        out.AvailableBorrowsBase,
		CurrentLiquidationThreshold: out.CurrentLiquidationThreshold,
		LTV:                         out.Ltv,
		HealthFactor:                out.HealthFactor,
	}
}

func reserveDataFromCodegen(out pool.Struct1) *ReserveData {
	return &ReserveData{
		LiquidityRate:      out.CurrentLiquidityRate,
		VariableBorrowRate: out.CurrentVariableBorrowRate,
	}
}

func reserveConfigDataFromCodegen(out dp.GetReserveConfigurationDataOutput) *ReserveConfigData {
	return &ReserveConfigData{
		Decimals:                 out.Decimals,
		LTV:                     out.Ltv,
		LiquidationThreshold:    out.LiquidationThreshold,
		LiquidationBonus:        out.LiquidationBonus,
		ReserveFactor:            out.ReserveFactor,
		UsageAsCollateralEnabled: out.UsageAsCollateralEnabled,
		BorrowingEnabled:         out.BorrowingEnabled,
		IsActive:                 out.IsActive,
		IsFrozen:                 out.IsFrozen,
	}
}

func RayToAPY(ray *big.Int) float64 {
	f := new(big.Float).SetInt(ray)
	divisor := new(big.Float).SetFloat64(1e25)
	pct, _ := new(big.Float).Quo(f, divisor).Float64()
	return math.Round(pct*100) / 100
}

func ParseAmount(s string, decimals int) (*big.Int, error) {
	s = strings.TrimSpace(s)
	if strings.EqualFold(s, "max") {
		return new(big.Int).Set(MaxUint256), nil
	}

	parts := strings.SplitN(s, ".", 2)
	if len(parts) == 0 || parts[0] == "" && (len(parts) < 2 || parts[1] == "") {
		return nil, fmt.Errorf("invalid amount: %q", s)
	}

	wholePart := parts[0]
	if wholePart == "" {
		wholePart = "0"
	}

	fracPart := ""
	if len(parts) == 2 {
		fracPart = parts[1]
	}

	if len(fracPart) > decimals {
		fracPart = fracPart[:decimals]
	}
	for len(fracPart) < decimals {
		fracPart += "0"
	}

	combined := wholePart + fracPart
	combined = strings.TrimLeft(combined, "0")
	if combined == "" {
		combined = "0"
	}

	result, ok := new(big.Int).SetString(combined, 10)
	if !ok {
		return nil, fmt.Errorf("invalid amount: %q", s)
	}

	return result, nil
}
