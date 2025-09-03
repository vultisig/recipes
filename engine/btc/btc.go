package btc

import (
	"bytes"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/vultisig/recipes/engine/compare"
	"github.com/vultisig/recipes/resolver"
	"github.com/vultisig/recipes/types"
)

type Btc struct{}

func NewBtc() *Btc {
	return &Btc{}
}

func (b *Btc) Evaluate(rule *types.Rule, txBytes []byte) error {
	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
	}
	if rule.GetTarget() != nil {
		return fmt.Errorf("target must be nil for BTC, got: %s", rule.GetTarget().String())
	}

	tx, err := b.parseTx(txBytes)
	if err != nil {
		return fmt.Errorf("failed to parse bitcoin transaction: %w", err)
	}

	if err := b.validateOutputs(rule, tx); err != nil {
		return fmt.Errorf("failed to validate outputs: %w", err)
	}

	return nil
}

func (b *Btc) parseTx(txBytes []byte) (*wire.MsgTx, error) {
	tx := &wire.MsgTx{}
	err := tx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize transaction: %w", err)
	}
	return tx, nil
}

type outputConstraints struct {
	address *types.ParameterConstraint
	value   *types.ParameterConstraint
	data    *types.ParameterConstraint
}

func (b *Btc) validateOutputs(rule *types.Rule, tx *wire.MsgTx) error {
	outputs := make(map[int]*outputConstraints)

	for _, constraint := range rule.GetParameterConstraints() {
		name := constraint.GetParameterName()

		if index, constrType, err := b.parseConstraintName(name); err != nil {
			return fmt.Errorf("failed to parse constraint name: %w", err)
		} else {
			b.setConstraint(outputs, index, constraint, constrType)
		}
	}

	if err := b.validateOutputConstraintCounts(outputs, tx); err != nil {
		return fmt.Errorf("failed to validate output constraint counts: %w", err)
	}

	return b.validateOutputConstraints(outputs, tx)
}

type constraintType string

const (
	address constraintType = "address"
	value   constraintType = "value"
	data    constraintType = "data"
)

func (b *Btc) parseConstraintName(name string) (index int, constraintType constraintType, err error) {
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

func (b *Btc) setConstraint(
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

func (b *Btc) validateOutputConstraintCounts(outputConstraints map[int]*outputConstraints, tx *wire.MsgTx) error {
	if len(outputConstraints) != len(tx.TxOut) {
		return fmt.Errorf("output count mismatch: rule has %d outputs, tx has %d outputs", len(outputConstraints), len(tx.TxOut))
	}

	for i := 0; i < len(tx.TxOut); i += 1 {
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

func (b *Btc) validateOutputConstraints(outputConstraints map[int]*outputConstraints, tx *wire.MsgTx) error {
	for i, txOut := range tx.TxOut {
		constraints := outputConstraints[i]

		if constraints.data != nil {
			// Data constraint validation - validate against OP_RETURN data
			if len(txOut.PkScript) < 2 || txOut.PkScript[0] != 0x6a {
				return fmt.Errorf("output %d is not an OP_RETURN script", i)
			}

			// Extract data from OP_RETURN script: 0x6a <data_len> <data>
			var dataBytes []byte
			if len(txOut.PkScript) > 2 {
				dataLen := int(txOut.PkScript[1])
				if len(txOut.PkScript) >= 2+dataLen {
					dataBytes = txOut.PkScript[2 : 2+dataLen]
				}
			}

			// Use raw bytes as string for regexp matching (ASCII data)
			dataStr := string(dataBytes)

			if er := validateConstraint(constraints.data, dataStr, compare.NewString); er != nil {
				return fmt.Errorf("output %d data validation failed: %w", i, er)
			}
		} else {
			// Address+value constraint validation
			outputAddress, err := b.extractAddress(txOut)
			if err != nil {
				return fmt.Errorf("failed to extract address from output %d: %w", i, err)
			}

			outputAmount := big.NewInt(txOut.Value)

			if er := validateConstraint(constraints.address, outputAddress, compare.NewString); er != nil {
				return fmt.Errorf("output %d address validation failed: %w", i, er)
			}

			if er := validateConstraint(constraints.value, outputAmount, compare.NewBigInt); er != nil {
				return fmt.Errorf("output %d value validation failed: %w", i, er)
			}
		}
	}
	return nil
}

func (b *Btc) extractAddress(txOut *wire.TxOut) (string, error) {
	_, addrs, _, err := txscript.ExtractPkScriptAddrs(txOut.PkScript, &chaincfg.MainNetParams)
	if err != nil {
		return "", fmt.Errorf("failed to extract address from script: %w", err)
	}

	if len(addrs) == 0 {
		return "", fmt.Errorf("no address found in script")
	}

	return addrs[0].EncodeAddress(), nil
}

func validateConstraint[T any](
	constraint *types.ParameterConstraint,
	actual T,
	makeComparer compare.Constructor[T],
) error {
	kind := constraint.GetConstraint().GetType()

	switch kind {
	case types.ConstraintType_CONSTRAINT_TYPE_ANY:
		return nil

	case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
		comparer, err := makeComparer(constraint.GetConstraint().GetFixedValue())
		if err != nil {
			return fmt.Errorf("failed to build fixed comparer: %w", err)
		}
		if comparer.Fixed(actual) {
			return nil
		}
		return fmt.Errorf("fixed value constraint failed: expected=%v, actual=%v",
			constraint.GetConstraint().GetFixedValue(), actual)

	case types.ConstraintType_CONSTRAINT_TYPE_MIN:
		comparer, err := makeComparer(constraint.GetConstraint().GetMinValue())
		if err != nil {
			return fmt.Errorf("failed to build min comparer: %w", err)
		}
		if comparer.Min(actual) {
			return nil
		}
		return fmt.Errorf("min value constraint failed: expected>=%v, actual=%v",
			constraint.GetConstraint().GetMinValue(), actual)

	case types.ConstraintType_CONSTRAINT_TYPE_MAX:
		comparer, err := makeComparer(constraint.GetConstraint().GetMaxValue())
		if err != nil {
			return fmt.Errorf("failed to build max comparer: %w", err)
		}
		if comparer.Max(actual) {
			return nil
		}
		return fmt.Errorf("max value constraint failed: expected<=%v, actual=%v",
			constraint.GetConstraint().GetMaxValue(), actual)

	case types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT:
		resolve, err := resolver.NewMagicConstantRegistry().GetResolver(
			constraint.GetConstraint().GetMagicConstantValue(),
		)
		if err != nil {
			return fmt.Errorf(
				"failed to get magic const resolver: magic_const=%s",
				constraint.GetConstraint().GetMagicConstantValue().String(),
			)
		}

		resolvedValue, _, err := resolve.Resolve(
			constraint.GetConstraint().GetMagicConstantValue(),
			"bitcoin",
			"default",
		)
		if err != nil {
			return fmt.Errorf(
				"failed to resolve magic const: magic_const=%s",
				constraint.GetConstraint().GetMagicConstantValue().String(),
			)
		}

		comparer, err := makeComparer(resolvedValue)
		if err != nil {
			return fmt.Errorf(
				"failed to build magic comparer: resolved=%s",
				resolvedValue,
			)
		}
		if comparer.Magic(actual) {
			return nil
		}
		return fmt.Errorf(
			"magic value constraint failed: expected(resolved)=%v, actual=%v",
			resolvedValue,
			actual,
		)

	case types.ConstraintType_CONSTRAINT_TYPE_REGEXP:
		strVal := fmt.Sprintf("%v", actual)
		ok, err := regexp.MatchString(
			constraint.GetConstraint().GetRegexpValue(),
			strVal,
		)
		if err != nil {
			return fmt.Errorf("regexp match failed: expected=%v, actual=%v",
				constraint.GetConstraint().GetRegexpValue(), actual)
		}
		if ok {
			return nil
		}
		return fmt.Errorf("regexp value constraint failed: expected=%v, actual=%v",
			constraint.GetConstraint().GetRegexpValue(), actual)

	default:
		return fmt.Errorf("unknown constraint type: %s", kind.String())
	}
}
