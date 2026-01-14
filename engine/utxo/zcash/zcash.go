package zcash

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/txscript"
	"github.com/vultisig/recipes/chain/utxo/zcash"
	stdcompare "github.com/vultisig/recipes/engine/compare"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

type Zcash struct{}

func NewZcash() *Zcash {
	return &Zcash{}
}

// Supports returns true if this engine supports the given chain
func (z *Zcash) Supports(chain common.Chain) bool {
	return chain == common.Zcash
}

func (z *Zcash) Evaluate(rule *types.Rule, txBytes []byte) error {
	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
	}
	if rule.GetTarget().GetTargetType() != types.TargetType_TARGET_TYPE_UNSPECIFIED {
		return fmt.Errorf("target type must be unspecified for Zcash, got: %s", rule.GetTarget().GetTargetType().String())
	}

	tx, err := z.parseTx(txBytes)
	if err != nil {
		return fmt.Errorf("failed to parse zcash transaction: %w", err)
	}

	if err := z.validateOutputs(rule, tx); err != nil {
		return fmt.Errorf("failed to validate outputs: %w", err)
	}

	return nil
}

func (z *Zcash) parseTx(txBytes []byte) (*zcash.ZcashTransaction, error) {
	parsed, err := zcash.ParseZcashTransactionWithNetwork(
		fmt.Sprintf("%x", txBytes),
		zcash.ZcashMainNetParams,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to parse transaction: %w", err)
	}

	// Type assert to get the underlying transaction
	zcashParsed, ok := parsed.(*zcash.ParsedZcashTransaction)
	if !ok {
		return nil, fmt.Errorf("unexpected transaction type")
	}

	return zcashParsed.GetTransaction(), nil
}

type outputConstraints struct {
	address *types.ParameterConstraint
	value   *types.ParameterConstraint
	data    *types.ParameterConstraint
}

func (z *Zcash) validateOutputs(rule *types.Rule, tx *zcash.ZcashTransaction) error {
	outputs := make(map[int]*outputConstraints)

	for _, constraint := range rule.GetParameterConstraints() {
		name := constraint.GetParameterName()

		if index, constrType, err := z.parseConstraintName(name); err != nil {
			return fmt.Errorf("failed to parse constraint name: %w", err)
		} else {
			z.setConstraint(outputs, index, constraint, constrType)
		}
	}

	if err := z.validateOutputConstraintCounts(outputs, tx); err != nil {
		return fmt.Errorf("failed to validate output constraint counts: %w", err)
	}

	return z.validateOutputConstraints(rule.GetParameterConstraints(), outputs, tx)
}

type constraintType string

const (
	address constraintType = "address"
	value   constraintType = "value"
	data    constraintType = "data"
)

func (z *Zcash) parseConstraintName(name string) (index int, constraintType constraintType, err error) {
	if strings.HasPrefix(name, "output_address_") {
		indexStr := strings.TrimPrefix(name, "output_address_")
		ind, er := strconv.Atoi(indexStr)
		if er != nil {
			return 0, "", fmt.Errorf("invalid constraint name: %s", name)
		}
		return ind, address, nil
	}

	if strings.HasPrefix(name, "output_value_") {
		indexStr := strings.TrimPrefix(name, "output_value_")
		ind, er := strconv.Atoi(indexStr)
		if er != nil {
			return 0, "", fmt.Errorf("invalid constraint name: %s", name)
		}
		return ind, value, nil
	}

	if strings.HasPrefix(name, "output_data_") {
		indexStr := strings.TrimPrefix(name, "output_data_")
		ind, er := strconv.Atoi(indexStr)
		if er != nil {
			return 0, "", fmt.Errorf("invalid constraint name: %s", name)
		}
		return ind, data, nil
	}

	return 0, "", fmt.Errorf("unsupported constraint parameter name (only output_* supported): %s", name)
}

func (z *Zcash) setConstraint(
	constraints map[int]*outputConstraints,
	index int,
	constraint *types.ParameterConstraint,
	constraintType constraintType,
) {
	if constraints[index] == nil {
		constraints[index] = &outputConstraints{}
	}

	switch constraintType {
	case address:
		constraints[index].address = constraint
	case value:
		constraints[index].value = constraint
	case data:
		constraints[index].data = constraint
	}
}

func (z *Zcash) validateOutputConstraintCounts(outputConstraints map[int]*outputConstraints, tx *zcash.ZcashTransaction) error {
	if len(outputConstraints) != len(tx.Outputs) {
		return fmt.Errorf("output count mismatch: rule has %d outputs, tx has %d outputs", len(outputConstraints), len(tx.Outputs))
	}

	for i := 0; i < len(tx.Outputs); i++ {
		constraints, exists := outputConstraints[i]
		if !exists {
			return fmt.Errorf("missing constraints for output %d", i)
		}

		// Exclusivity logic: output must be either data OR (address+value), but not both
		hasData := constraints.data != nil
		hasAddressValue := constraints.address != nil && constraints.value != nil

		if hasData && hasAddressValue {
			return fmt.Errorf("output %d cannot have both data and address+value constraints", i)
		}

		if !hasData && !hasAddressValue {
			return fmt.Errorf("output %d must have either data constraint or both address and value constraints", i)
		}

		// If using address+value format, both must be present
		if !hasData && (constraints.address == nil || constraints.value == nil) {
			return fmt.Errorf("output %d with non-data format must have both address and value constraints", i)
		}
	}

	return nil
}

func (z *Zcash) validateOutputConstraints(
	constraintList []*types.ParameterConstraint,
	outputConstraints map[int]*outputConstraints,
	tx *zcash.ZcashTransaction,
) error {
	const chainID = "zcash"

	for i, txOut := range tx.Outputs {
		constraints := outputConstraints[i]

		if constraints.data != nil {
			// Data constraint validation - validate against OP_RETURN data
			if len(txOut.PkScript) < 2 || txOut.PkScript[0] != txscript.OP_RETURN {
				return fmt.Errorf("output %d is not an OP_RETURN script", i)
			}

			// Extract data from OP_RETURN script using txscript.PushedData
			// which handles all PUSHDATA variants (OP_DATA_1-75, OP_PUSHDATA1/2/4)
			pushedData, _ := txscript.PushedData(txOut.PkScript[1:]) // Skip OP_RETURN
			var dataBytes []byte
			if len(pushedData) > 0 {
				dataBytes = pushedData[0]
			}

			// Use raw bytes as string for regexp matching (ASCII data)
			dataStr := string(dataBytes)

			if er := stdcompare.AssertArg(chainID, constraintList, fmt.Sprintf("output_data_%d", i), dataStr, stdcompare.NewString); er != nil {
				return fmt.Errorf("output %d data validation failed: %w", i, er)
			}
		} else {
			// Address+value constraint validation
			outputAddress, err := z.extractAddress(txOut)
			if err != nil {
				return fmt.Errorf("failed to extract address from output %d: %w", i, err)
			}

			outputAmount := big.NewInt(txOut.Value)

			if er := stdcompare.AssertArg(chainID, constraintList, fmt.Sprintf("output_address_%d", i), outputAddress, stdcompare.NewString); er != nil {
				return fmt.Errorf("output %d address validation failed: %w", i, er)
			}

			if er := stdcompare.AssertArg(chainID, constraintList, fmt.Sprintf("output_value_%d", i), outputAmount, stdcompare.NewBigInt); er != nil {
				return fmt.Errorf("output %d value validation failed: %w", i, er)
			}
		}
	}
	return nil
}

func (z *Zcash) extractAddress(txOut *zcash.ZcashOutput) (string, error) {
	addr, err := zcash.ExtractZcashAddress(txOut.PkScript, zcash.ZcashMainNetParams)
	if err != nil {
		return "", fmt.Errorf("failed to extract address from script: %w", err)
	}
	return addr, nil
}
