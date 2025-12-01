package xrpl

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	stdcompare "github.com/vultisig/recipes/engine/compare"
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

	// Validate parameter constraints for XRP payments
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
	for _, constraint := range constraints {
		paramName := constraint.GetParameterName()

		// Extract the actual value from the payment based on parameter name
		value, err := x.extractParameterValue(paramName, payment)
		if err != nil {
			return fmt.Errorf("failed to extract parameter %s: %w", paramName, err)
		}

		// Use type-based constraint validation
		if err := x.assertArgsByType(resource.ChainId, paramName, value, constraints); err != nil {
			return fmt.Errorf("constraint validation failed for parameter %s: %w", paramName, err)
		}
	}
	return nil
}

// extractParameterValue extracts the actual value from payment for the given parameter name
// Returns appropriate Go types: *big.Int for amounts, string for others
func (x *XRPL) extractParameterValue(paramName string, payment *transactions.Payment) (interface{}, error) {
	switch paramName {
	case "recipient":
		return string(payment.Destination), nil
	case "amount":
		// Return amount as *big.Int for numeric comparisons
		return x.extractCurrencyAmountAsBigInt(payment.Amount)
	case "memo":
		return ExtractMemoFromXRPPayment(payment)
	default:
		return nil, fmt.Errorf("unsupported parameter: %s", paramName)
	}
}

// ExtractMemoFromXRPPayment extracts memo data from XRPL Payment transaction
func ExtractMemoFromXRPPayment(payment *transactions.Payment) (string, error) {
	if len(payment.Memos) == 0 {
		return "", fmt.Errorf("memo required but not found in payment transaction")
	}

	// XRPL memos are typically hex-encoded, need to decode
	memo := payment.Memos[0]
	if memo.Memo.MemoData == "" {
		return "", fmt.Errorf("memo field present but empty")
	}

	// Decode hex to string (THORChain memos are text)
	memoBytes, err := hex.DecodeString(memo.Memo.MemoData)
	if err != nil {
		// If not hex, treat as plain string
		return memo.Memo.MemoData, nil
	}

	return string(memoBytes), nil
}

// extractCurrencyAmountAsBigInt converts a CurrencyAmount to *big.Int for numeric comparisons
// For now, only XRP native tokens are supported
func (x *XRPL) extractCurrencyAmountAsBigInt(amount xrptypes.CurrencyAmount) (*big.Int, error) {
	if amount == nil {
		return nil, fmt.Errorf("amount is nil")
	}

	xrpAmount, ok := amount.(xrptypes.XRPCurrencyAmount)
	if !ok {
		return nil, fmt.Errorf("only XRP amounts are supported, got: %T", amount)
	}

	return big.NewInt(int64(xrpAmount)), nil
}

// assertArgsByType validates constraints using the appropriate comparator based on Go type
func (x *XRPL) assertArgsByType(chainId, inputName string, arg interface{}, constraints []*types.ParameterConstraint) error {
	switch actual := arg.(type) {
	case string:
		err := stdcompare.AssertArg(
			chainId,
			constraints,
			inputName,
			actual,
			stdcompare.NewString,
		)
		if err != nil {
			return fmt.Errorf("failed to assert string parameter: %w", err)
		}

	case *big.Int:
		err := stdcompare.AssertArg(
			chainId,
			constraints,
			inputName,
			actual,
			stdcompare.NewBigInt,
		)
		if err != nil {
			return fmt.Errorf("failed to assert big.Int parameter: %w", err)
		}

	default:
		return fmt.Errorf("unsupported parameter type: %T", actual)
	}
	return nil
}
