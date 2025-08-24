package btc

import (
	"bytes"
	"fmt"
	"math/big"
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
}

func (b *Btc) validateOutputs(rule *types.Rule, tx *wire.MsgTx) error {
	outputs := make(map[int]*outputConstraints)

	// Only process output constraints
	for _, constraint := range rule.GetParameterConstraints() {
		name := constraint.GetParameterName()

		if index, isAddress, err := b.parseConstraintName(name); err != nil {
			return fmt.Errorf("failed to parse constraint name: %w", err)
		} else {
			// All constraints are output constraints now
			b.setConstraint(outputs, index, constraint, isAddress)
		}
	}

	if err := b.validateOutputConstraintCounts(outputs, tx); err != nil {
		return err
	}

	return b.validateOutputConstraints(outputs, tx)
}

func (b *Btc) parseConstraintName(name string) (index int, isAddress bool, err error) {
	if strings.HasPrefix(name, "output_address_") {
		indexStr := strings.TrimPrefix(name, "output_address_")
		ind, er := strconv.Atoi(indexStr)
		if er != nil {
			return 0, false, fmt.Errorf("invalid constraint name: %s", name)
		}
		return ind, true, nil
	}

	if strings.HasPrefix(name, "output_value_") {
		indexStr := strings.TrimPrefix(name, "output_value_")
		ind, er := strconv.Atoi(indexStr)
		if er != nil {
			return 0, false, fmt.Errorf("invalid constraint name: %s", name)
		}
		return ind, false, nil
	}

	return 0, false, fmt.Errorf("unsupported constraint parameter name (only output_* supported): %s", name)
}

func (b *Btc) setConstraint(constraints map[int]*outputConstraints, index int, constraint *types.ParameterConstraint, isAddress bool) {
	if constraints[index] == nil {
		constraints[index] = &outputConstraints{}
	}

	if isAddress {
		constraints[index].address = constraint
	} else {
		constraints[index].value = constraint
	}
}

func (b *Btc) validateOutputConstraintCounts(outputConstraints map[int]*outputConstraints, tx *wire.MsgTx) error {
	if len(outputConstraints) != len(tx.TxOut) {
		return fmt.Errorf("output count mismatch: rule has %d outputs, tx has %d outputs", len(outputConstraints), len(tx.TxOut))
	}

	for i := 0; i < len(tx.TxOut); i += 1 {
		if constraints, exists := outputConstraints[i]; !exists || constraints.address == nil || constraints.value == nil {
			return fmt.Errorf("missing address or value constraint for output %d", i)
		}
	}

	return nil
}

func (b *Btc) validateOutputConstraints(outputConstraints map[int]*outputConstraints, tx *wire.MsgTx) error {
	for i, txOut := range tx.TxOut {
		constraints := outputConstraints[i]

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

	default:
		return fmt.Errorf("unknown constraint type: %s", kind.String())
	}
}
