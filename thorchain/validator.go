package thorchain

import (
	"fmt"
	"math/big"
)

// ThorchainValidator provides basic validation for RUNE/TCY transfers
type ThorchainValidator struct{}

// NewThorchainValidator creates a new Thorchain validator
func NewThorchainValidator() *ThorchainValidator {
	return &ThorchainValidator{}
}

// ValidateTransfer validates basic RUNE/TCY transfer parameters
func (v *ThorchainValidator) ValidateTransfer(params map[string]interface{}) error {
	// Validate recipient address
	if recipient, ok := params["recipient"]; ok {
		if recipientStr, ok := recipient.(string); ok {
			if err := ValidateThorchainAddress(recipientStr); err != nil {
				return fmt.Errorf("invalid recipient address: %w", err)
			}
		} else {
			return fmt.Errorf("recipient must be a string address")
		}
	} else {
		return fmt.Errorf("recipient address is required for transfers")
	}

	// Get denomination from params, default to "rune"
	denom := "rune"
	if denomParam, ok := params["denom"]; ok {
		if denomStr, ok := denomParam.(string); ok {
			denom = denomStr
		}
	}

	// Validate denomination is supported
	if denom != "rune" && denom != "tcy" {
		return fmt.Errorf("unsupported denomination: %s (must be 'rune' or 'tcy')", denom)
	}

	// Validate amount
	if amount, ok := params["amount"]; ok {
		if amountBig, ok := amount.(*big.Int); ok {
			if err := ValidateAmount(amountBig, denom); err != nil {
				return fmt.Errorf("invalid transfer amount: %w", err)
			}
		} else {
			return fmt.Errorf("amount must be a big.Int")
		}
	} else {
		return fmt.Errorf("amount is required for transfers")
	}

	// Validate memo if present (optional)
	if memo, ok := params["memo"]; ok {
		if memoStr, ok := memo.(string); ok {
			if err := ValidateMemo(memoStr); err != nil {
				return fmt.Errorf("invalid memo: %w", err)
			}
		}
	}

	return nil
}
