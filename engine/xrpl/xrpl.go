package xrpl

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/vultisig/recipes/types"
	xrpgo "github.com/xyield/xrpl-go/binary-codec"
	"github.com/xyield/xrpl-go/model/transactions"
)

// XRPLTransaction represents a parsed XRPL Payment transaction
type XRPLTransaction struct {
	Account         string              `json:"Account"`
	TransactionType transactions.TxType `json:"TransactionType"`
	Destination     string              `json:"Destination"`
	Amount          string              `json:"Amount"` // For XRP (drops) or object for tokens
	DestinationTag  *uint32             `json:"DestinationTag,omitempty"`
	Fee             string              `json:"Fee"`
	Sequence        uint32              `json:"Sequence"`
}

// XRPL represents the XRP Ledger engine implementation
type XRPL struct{}

// NewXRPL creates a new XRPL engine instance
func NewXRPL() *XRPL {
	return &XRPL{}
}

// Evaluate validates an XRPL transaction against policy rules
// This is the main entry point called by the main engine
func (x *XRPL) Evaluate(rule *types.Rule, txBytes []byte) error {
	// Validate rule effect is ALLOW (following existing pattern from BTC/EVM engines)
	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
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
	if err := x.validateTarget(rule.GetTarget(), tx); err != nil {
		return fmt.Errorf("failed to validate target: %w", err)
	}

	// Validate parameter constraints
	if err := x.validateParameterConstraints(rule.GetParameterConstraints(), tx); err != nil {
		return fmt.Errorf("failed to validate parameter constraints: %w", err)
	}

	return nil
}

// parseTransaction parses XRPL transaction bytes into a structured format
func (x *XRPL) parseTransaction(txBytes []byte) (*XRPLTransaction, error) {
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

	// Unmarshal into our struct
	var tx XRPLTransaction
	if err := json.Unmarshal(jsonBytes, &tx); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XRPL transaction: %w", err)
	}

	// Validate required fields for Payment transactions
	if tx.Account == "" {
		return nil, fmt.Errorf("account field is required")
	}
	if tx.TransactionType == "" {
		return nil, fmt.Errorf("transactionType field is required")
	}
	if tx.Destination == "" {
		return nil, fmt.Errorf("destination field is required")
	}
	if tx.Amount == "" {
		return nil, fmt.Errorf("amount field is required")
	}

	return &tx, nil
}

// validateTarget validates the transaction target against the rule target
func (x *XRPL) validateTarget(target *types.Target, tx *XRPLTransaction) error {
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
		if !strings.EqualFold(tx.Destination, expectedAddress) {
			return fmt.Errorf("target address mismatch: expected=%s, actual=%s",
				expectedAddress, tx.Destination)
		}
	default:
		return fmt.Errorf("unsupported target type: %s", target.GetTargetType())
	}

	return nil
}

// validateParameterConstraints validates all parameter constraints
func (x *XRPL) validateParameterConstraints(constraints []*types.ParameterConstraint, tx *XRPLTransaction) error {
	for _, constraint := range constraints {
		paramName := constraint.GetParameterName()

		var err error
		switch paramName {
		case "recipient":
			err = x.validateRecipientConstraint(constraint, tx.Destination)
		case "amount":
			err = x.validateAmountConstraint(constraint, tx.Amount)
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
func (x *XRPL) validateRecipientConstraint(constraint *types.ParameterConstraint, recipient string) error {
	constraintType := constraint.GetConstraint().GetType()

	switch constraintType {
	case types.ConstraintType_CONSTRAINT_TYPE_ANY:
		return nil

	case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
		expectedValue := constraint.GetConstraint().GetFixedValue()
		// Case-insensitive comparison for addresses
		if !strings.EqualFold(recipient, expectedValue) {
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
		switch magicConstant {
		case types.MagicConstant_VULTISIG_TREASURY:
			// TODO: Replace with actual Vultisig treasury XRPL address
			treasuryAddress := "rVultisigTreasuryAddress123456789"
			if !strings.EqualFold(recipient, treasuryAddress) {
				return fmt.Errorf("magic constant VULTISIG_TREASURY constraint failed: expected=%s, actual=%s",
					treasuryAddress, recipient)
			}
		case types.MagicConstant_THORCHAIN_VAULT:
			// TODO: Replace with actual THORChain vault XRPL address
			thorchainAddress := "rThorchainVaultAddress123456789"
			if !strings.EqualFold(recipient, thorchainAddress) {
				return fmt.Errorf("magic constant THORCHAIN_VAULT constraint failed: expected=%s, actual=%s",
					thorchainAddress, recipient)
			}
		default:
			return fmt.Errorf("unsupported magic constant: %s", magicConstant)
		}

	default:
		return fmt.Errorf("unsupported constraint type for recipient: %s", constraintType)
	}

	return nil
}

// validateAmountConstraint validates amount constraints (in drops)
func (x *XRPL) validateAmountConstraint(constraint *types.ParameterConstraint, amount string) error {
	// Convert amount to big.Int for numeric comparisons
	amountBigInt := new(big.Int)
	if _, ok := amountBigInt.SetString(amount, 10); !ok {
		return fmt.Errorf("invalid amount format: %s", amount)
	}

	constraintType := constraint.GetConstraint().GetType()

	switch constraintType {
	case types.ConstraintType_CONSTRAINT_TYPE_ANY:
		return nil

	case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
		expectedAmount := constraint.GetConstraint().GetFixedValue()
		if amount != expectedAmount {
			return fmt.Errorf("fixed amount constraint failed: expected=%s, actual=%s",
				expectedAmount, amount)
		}

	case types.ConstraintType_CONSTRAINT_TYPE_MIN:
		minValue := constraint.GetConstraint().GetMinValue()
		minBigInt := new(big.Int)
		if _, ok := minBigInt.SetString(minValue, 10); !ok {
			return fmt.Errorf("invalid min constraint value: %s", minValue)
		}
		if amountBigInt.Cmp(minBigInt) < 0 {
			return fmt.Errorf("min amount constraint failed: expected>=%s, actual=%s",
				minValue, amount)
		}

	case types.ConstraintType_CONSTRAINT_TYPE_MAX:
		maxValue := constraint.GetConstraint().GetMaxValue()
		maxBigInt := new(big.Int)
		if _, ok := maxBigInt.SetString(maxValue, 10); !ok {
			return fmt.Errorf("invalid max constraint value: %s", maxValue)
		}
		if amountBigInt.Cmp(maxBigInt) > 0 {
			return fmt.Errorf("max amount constraint failed: expected<=%s, actual=%s",
				maxValue, amount)
		}

	case types.ConstraintType_CONSTRAINT_TYPE_REGEXP:
		pattern := constraint.GetConstraint().GetRegexpValue()
		if pattern == "" {
			return fmt.Errorf("regexp pattern cannot be empty")
		}
		matched, err := regexp.MatchString(pattern, amount)
		if err != nil {
			return fmt.Errorf("invalid regexp pattern: %w", err)
		}
		if !matched {
			return fmt.Errorf("regexp constraint failed: pattern=%s, value=%s", pattern, amount)
		}

	default:
		return fmt.Errorf("unsupported constraint type for amount: %s", constraintType)
	}

	return nil
}
