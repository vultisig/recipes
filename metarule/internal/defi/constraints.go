package defi

import (
	"fmt"

	"github.com/vultisig/recipes/types"
)

// LPConstraints holds parsed LP meta-protocol constraints
type LPConstraints struct {
	Action     *types.Constraint
	Protocol   *types.Constraint
	Pool       *types.Constraint
	Token0     *types.Constraint
	Token1     *types.Constraint
	Amount0    *types.Constraint
	Amount1    *types.Constraint
	MinAmount0 *types.Constraint
	MinAmount1 *types.Constraint
	Recipient  *types.Constraint
}

// ParseLPConstraints extracts LP constraints from a rule
func ParseLPConstraints(rule *types.Rule) (LPConstraints, error) {
	res := LPConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "action":
			res.Action = c.GetConstraint()
		case "protocol":
			res.Protocol = c.GetConstraint()
		case "pool":
			res.Pool = c.GetConstraint()
		case "token0":
			res.Token0 = c.GetConstraint()
		case "token1":
			res.Token1 = c.GetConstraint()
		case "amount0":
			res.Amount0 = c.GetConstraint()
		case "amount1":
			res.Amount1 = c.GetConstraint()
		case "min_amount0":
			res.MinAmount0 = c.GetConstraint()
		case "min_amount1":
			res.MinAmount1 = c.GetConstraint()
		case "recipient":
			res.Recipient = c.GetConstraint()
		}
	}

	if res.Action == nil {
		return res, fmt.Errorf("missing required constraint: action")
	}
	if res.Protocol == nil {
		return res, fmt.Errorf("missing required constraint: protocol")
	}

	return res, nil
}

// LendConstraints holds parsed lending meta-protocol constraints
type LendConstraints struct {
	Action     *types.Constraint
	Protocol   *types.Constraint
	Asset      *types.Constraint
	Amount     *types.Constraint
	OnBehalfOf *types.Constraint
	Collateral *types.Constraint
}

// ParseLendConstraints extracts lending constraints from a rule
func ParseLendConstraints(rule *types.Rule) (LendConstraints, error) {
	res := LendConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "action":
			res.Action = c.GetConstraint()
		case "protocol":
			res.Protocol = c.GetConstraint()
		case "asset":
			res.Asset = c.GetConstraint()
		case "amount":
			res.Amount = c.GetConstraint()
		case "on_behalf_of":
			res.OnBehalfOf = c.GetConstraint()
		case "collateral":
			res.Collateral = c.GetConstraint()
		}
	}

	if res.Action == nil {
		return res, fmt.Errorf("missing required constraint: action")
	}
	if res.Protocol == nil {
		return res, fmt.Errorf("missing required constraint: protocol")
	}
	if res.Asset == nil {
		return res, fmt.Errorf("missing required constraint: asset")
	}
	if res.Amount == nil {
		return res, fmt.Errorf("missing required constraint: amount")
	}

	return res, nil
}

// PerpsConstraints holds parsed perps meta-protocol constraints
type PerpsConstraints struct {
	Action          *types.Constraint
	Protocol        *types.Constraint
	Market          *types.Constraint
	SizeDelta       *types.Constraint
	CollateralDelta *types.Constraint
	CollateralToken *types.Constraint
	AcceptablePrice *types.Constraint
	ExecutionFee    *types.Constraint
	Leverage        *types.Constraint
}

// ParsePerpsConstraints extracts perps constraints from a rule
func ParsePerpsConstraints(rule *types.Rule) (PerpsConstraints, error) {
	res := PerpsConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "action":
			res.Action = c.GetConstraint()
		case "protocol":
			res.Protocol = c.GetConstraint()
		case "market":
			res.Market = c.GetConstraint()
		case "size_delta":
			res.SizeDelta = c.GetConstraint()
		case "collateral_delta":
			res.CollateralDelta = c.GetConstraint()
		case "collateral_token":
			res.CollateralToken = c.GetConstraint()
		case "acceptable_price":
			res.AcceptablePrice = c.GetConstraint()
		case "execution_fee":
			res.ExecutionFee = c.GetConstraint()
		case "leverage":
			res.Leverage = c.GetConstraint()
		}
	}

	if res.Action == nil {
		return res, fmt.Errorf("missing required constraint: action")
	}
	if res.Protocol == nil {
		return res, fmt.Errorf("missing required constraint: protocol")
	}

	return res, nil
}

// Helper functions for creating constraints

// Fixed creates a fixed value constraint
func Fixed(value string) *types.Constraint {
	return &types.Constraint{
		Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
		Value: &types.Constraint_FixedValue{
			FixedValue: value,
		},
	}
}

// Any creates an any-value constraint
func Any() *types.Constraint {
	return &types.Constraint{
		Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
	}
}

// Min creates a minimum value constraint
func Min(value string) *types.Constraint {
	return &types.Constraint{
		Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
		Value: &types.Constraint_MinValue{
			MinValue: value,
		},
	}
}

// Max creates a maximum value constraint
func Max(value string) *types.Constraint {
	return &types.Constraint{
		Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
		Value: &types.Constraint_MaxValue{
			MaxValue: value,
		},
	}
}

// BetConstraints holds parsed bet meta-protocol constraints
type BetConstraints struct {
	Action   *types.Constraint // buy, sell, cancel
	Protocol *types.Constraint // polymarket
	Market   *types.Constraint // Polymarket CTF tokenId (outcome token)
	Amount   *types.Constraint // fillAmount for CTFExchange.fillOrder
	Price    *types.Constraint // Limit price in basis points (0-10000 = 0-100%)
	Maker    *types.Constraint // Maker address
}

// ParseBetConstraints extracts bet constraints from a rule
func ParseBetConstraints(rule *types.Rule) (BetConstraints, error) {
	res := BetConstraints{}

	for _, c := range rule.GetParameterConstraints() {
		switch c.GetParameterName() {
		case "action":
			res.Action = c.GetConstraint()
		case "protocol":
			res.Protocol = c.GetConstraint()
		case "market":
			res.Market = c.GetConstraint()
		case "amount":
			res.Amount = c.GetConstraint()
		case "price":
			res.Price = c.GetConstraint()
		case "maker":
			res.Maker = c.GetConstraint()
		}
	}

	if res.Action == nil {
		return res, fmt.Errorf("missing required constraint: action")
	}
	if res.Protocol == nil {
		return res, fmt.Errorf("missing required constraint: protocol")
	}

	return res, nil
}

