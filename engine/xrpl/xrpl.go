package xrpl

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"regexp"

	"github.com/vultisig/recipes/engine/compare"
	"github.com/vultisig/recipes/resolver"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
	xrpgo "github.com/xyield/xrpl-go/binary-codec"
	"github.com/xyield/xrpl-go/model/transactions"
	xrptypes "github.com/xyield/xrpl-go/model/transactions/types"
)

// XRPL represents the XRP Ledger engine implementation
type XRPL struct{}

// NewXRPL creates a new XRPL engine instance
func NewXRPL() *XRPL {
	return &XRPL{}
}

// Supports returns true if this engine supports the given chain
func (x *XRPL) Supports(chain common.Chain) bool {
	return chain == common.XRP
}

// Evaluate validates an XRPL transaction against policy rules
// This is the main entry point called by the main engine
func (x *XRPL) Evaluate(rule *types.Rule, txBytes []byte) error {
	// Validate rule effect is ALLOW (following existing pattern from BTC/EVM engines)
	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
	}

	// Parse resource to extract protocol and function information
	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}
	// Basic resource checks
	if r.ChainId != "xrp" {
		return fmt.Errorf("unsupported chain in resource: %s", r.ChainId)
	}

	// Parse XRPL transaction from txBytes using binary codec
	tx, err := x.parseTransaction(txBytes)
	if err != nil {
		return fmt.Errorf("failed to parse XRPL transaction: %w", err)
	}

	// Validate it's a Payment transaction
	if tx.TransactionType != transactions.PaymentTx {
		return fmt.Errorf("only Payment transactions are supported, got: %s", tx.TransactionType)
	}

	// Validate target if specified
	if err := x.validateTarget(r, rule.GetTarget(), tx); err != nil {
		return fmt.Errorf("failed to validate target: %w", err)
	}

	// Validate parameter constraints
	if err := x.validateParameterConstraints(r, rule.GetParameterConstraints(), tx); err != nil {
		return fmt.Errorf("failed to validate parameter constraints: %w", err)
	}

	return nil
}

// parseTransaction parses XRPL transaction bytes into a Payment transaction
func (x *XRPL) parseTransaction(txBytes []byte) (*transactions.Payment, error) {
	// Convert bytes to hex string for binary codec
	hexStr := hex.EncodeToString(txBytes)

	// Use XRPL binary codec to decode hex to JSON
	jsonData, err := xrpgo.Decode(hexStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode XRPL binary format: %w", err)
	}

	// Convert map to JSON bytes for unmarshaling
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal decoded JSON: %w", err)
	}

	// Unmarshal into Payment struct
	var payment transactions.Payment
	if err := json.Unmarshal(jsonBytes, &payment); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XRPL Payment transaction: %w", err)
	}

	// Validate required fields for Payment transactions
	if string(payment.Account) == "" {
		return nil, fmt.Errorf("account field is required")
	}
	if payment.TransactionType != transactions.PaymentTx {
		return nil, fmt.Errorf("expected Payment transaction, got: %s", payment.TransactionType)
	}
	if string(payment.Destination) == "" {
		return nil, fmt.Errorf("destination field is required")
	}
	if payment.Amount == nil {
		return nil, fmt.Errorf("amount field is required")
	}

	return &payment, nil
}

// validateTarget validates the transaction target against the rule target
func (x *XRPL) validateTarget(resource *types.ResourcePath, target *types.Target, payment *transactions.Payment) error {
	if target == nil || target.GetTargetType() == types.TargetType_TARGET_TYPE_UNSPECIFIED {
		return nil // No target validation required
	}

	switch target.GetTargetType() {
	case types.TargetType_TARGET_TYPE_ADDRESS:
		expectedAddress := target.GetAddress()
		if expectedAddress == "" {
			return fmt.Errorf("target address cannot be empty")
		}
		// For XRPL, we validate against the Destination (recipient)
		actualDestination := string(payment.Destination)
		if actualDestination != expectedAddress {
			return fmt.Errorf("target address mismatch: expected=%s, actual=%s",
				expectedAddress, actualDestination)
		}
	case types.TargetType_TARGET_TYPE_MAGIC_CONSTANT:
		resolve, err := resolver.NewMagicConstantRegistry().GetResolver(target.GetMagicConstant())
		if err != nil {
			return fmt.Errorf(
				"failed to get resolver: magic_const=%s",
				target.GetMagicConstant().String(),
			)
		}

		resolvedAddr, _, err := resolve.Resolve(
			target.GetMagicConstant(),
			resource.ChainId,
			"default",
		)
		if err != nil {
			return fmt.Errorf(
				"failed to resolve magic const: value=%s, error=%w",
				target.GetMagicConstant().String(),
				err,
			)
		}
		actualDestination := string(payment.Destination)
		if actualDestination != resolvedAddr {
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_magic_const_resolved=%s",
				actualDestination,
				resolvedAddr,
			)
		}
		return nil

	default:
		return fmt.Errorf("unsupported target type: %s", target.GetTargetType())
	}

	return nil
}

// validateParameterConstraints validates all parameter constraints
func (x *XRPL) validateParameterConstraints(resource *types.ResourcePath, constraints []*types.ParameterConstraint, payment *transactions.Payment) error {
	chain := resource.ChainId
	for _, constraint := range constraints {
		paramName := constraint.GetParameterName()

		var err error
		switch paramName {
		case "recipient":
			err = x.validateRecipientConstraint(chain, constraint, string(payment.Destination))
		case "amount":
			err = x.validateAmountConstraint(constraint, payment)
		default:
			err = fmt.Errorf("unsupported parameter: %s", paramName)
		}

		if err != nil {
			return fmt.Errorf("constraint validation failed for parameter %s: %w", paramName, err)
		}
	}
	return nil
}

// validateRecipientConstraint validates recipient address constraints
func (x *XRPL) validateRecipientConstraint(chain string, constraint *types.ParameterConstraint, recipient string) error {
	constraintType := constraint.GetConstraint().GetType()

	switch constraintType {
	case types.ConstraintType_CONSTRAINT_TYPE_ANY:
		return nil

	case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
		expectedValue := constraint.GetConstraint().GetFixedValue()

		if recipient != expectedValue {
			return fmt.Errorf("fixed recipient constraint failed: expected=%s, actual=%s",
				expectedValue, recipient)
		}

	case types.ConstraintType_CONSTRAINT_TYPE_REGEXP:
		pattern := constraint.GetConstraint().GetRegexpValue()
		if pattern == "" {
			return fmt.Errorf("regexp pattern cannot be empty")
		}
		matched, err := regexp.MatchString(pattern, recipient)
		if err != nil {
			return fmt.Errorf("invalid regexp pattern: %w", err)
		}
		if !matched {
			return fmt.Errorf("regexp constraint failed: pattern=%s, value=%s", pattern, recipient)
		}

	case types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT:
		magicConstant := constraint.GetConstraint().GetMagicConstantValue()
		resolve, err := resolver.NewMagicConstantRegistry().GetResolver(magicConstant)
		if err != nil {
			return fmt.Errorf(
				"failed to get resolver: magic_const=%s",
				magicConstant.String(),
			)
		}

		resolvedAddr, _, err := resolve.Resolve(
			magicConstant,
			chain,
			"",
		)
		if err != nil {
			return fmt.Errorf(
				"failed to resolve magic const: value=%s, error=%w",
				magicConstant.String(),
				err,
			)
		}
		if recipient != resolvedAddr {
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_magic_const_resolved=%s",
				recipient,
				resolvedAddr,
			)
		}
	default:
		return fmt.Errorf("unsupported constraint type for recipient: %s", constraintType)
	}

	return nil
}

// validateAmountConstraint validates amount constraints (XRP only)
func (x *XRPL) validateAmountConstraint(constraint *types.ParameterConstraint, payment *transactions.Payment) error {
	// Check for partial payment flag (tfPartialPayment = 131072)
	const tfPartialPayment uint = 131072
	if payment.Flags&tfPartialPayment != 0 {
		// For now, reject partial payments as they could bypass amount constraints
		// The actual delivered amount could be less than payment.Amount
		return fmt.Errorf("partial payments are not supported for policy validation")
	}
	// Convert amount to string for comparison
	amountStr, err := x.formatCurrencyAmount(payment.Amount)
	if err != nil {
		return fmt.Errorf("failed to format currency amount: %w", err)
	}

	// Convert to big.Int for numeric comparisons
	amountBigInt := new(big.Int)
	if _, ok := amountBigInt.SetString(amountStr, 10); !ok {
		return fmt.Errorf("invalid XRP amount format: %s", amountStr)
	}

	constraintType := constraint.GetConstraint().GetType()

	switch constraintType {
	case types.ConstraintType_CONSTRAINT_TYPE_ANY:
		return nil

	case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
		expectedAmount := constraint.GetConstraint().GetFixedValue()
		comparator, err := compare.NewBigInt(expectedAmount)
		if err != nil {
			return fmt.Errorf("invalid fixed constraint value: %s", expectedAmount)
		}
		if !comparator.Fixed(amountBigInt) {
			return fmt.Errorf("fixed amount constraint failed: expected=%s, actual=%s",
				expectedAmount, amountStr)
		}

	case types.ConstraintType_CONSTRAINT_TYPE_MIN:
		minValue := constraint.GetConstraint().GetMinValue()
		comparator, err := compare.NewBigInt(minValue)
		if err != nil {
			return fmt.Errorf("invalid min constraint value: %s", minValue)
		}
		if !comparator.Min(amountBigInt) {
			return fmt.Errorf("min amount constraint failed: expected>=%s, actual=%s",
				minValue, amountStr)
		}

	case types.ConstraintType_CONSTRAINT_TYPE_MAX:
		maxValue := constraint.GetConstraint().GetMaxValue()
		comparator, err := compare.NewBigInt(maxValue)
		if err != nil {
			return fmt.Errorf("invalid max constraint value: %s", maxValue)
		}
		if !comparator.Max(amountBigInt) {
			return fmt.Errorf("max amount constraint failed: expected<=%s, actual=%s",
				maxValue, amountStr)
		}

	case types.ConstraintType_CONSTRAINT_TYPE_REGEXP:
		pattern := constraint.GetConstraint().GetRegexpValue()
		if pattern == "" {
			return fmt.Errorf("regexp pattern cannot be empty")
		}
		matched, err := regexp.MatchString(pattern, amountStr)
		if err != nil {
			return fmt.Errorf("invalid regexp pattern: %w", err)
		}
		if !matched {
			return fmt.Errorf("regexp constraint failed: pattern=%s, value=%s", pattern, amountStr)
		}

	default:
		return fmt.Errorf("unsupported constraint type for amount: %s", constraintType)
	}

	return nil
}

// formatCurrencyAmount converts a CurrencyAmount to string for comparison
// For now, only XRP native tokens are supported
func (x *XRPL) formatCurrencyAmount(amount xrptypes.CurrencyAmount) (string, error) {
	if amount == nil {
		return "", fmt.Errorf("amount is nil")
	}

	xrpAmount, ok := amount.(xrptypes.XRPCurrencyAmount)
	if !ok {
		return "", fmt.Errorf("only XRP amounts are supported, got: %T", amount)
	}

	// Convert XRP amount (drops) to string
	return fmt.Sprintf("%d", uint64(xrpAmount)), nil
}
