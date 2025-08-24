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

	if err := b.validateInputsOutputs(rule, tx); err != nil {
		return fmt.Errorf("failed to validate inputs/outputs: %w", err)
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

type ioConstraints struct {
	address *types.ParameterConstraint
	value   *types.ParameterConstraint
}

func (b *Btc) validateInputsOutputs(rule *types.Rule, tx *wire.MsgTx) error {
	inputConstraints := make(map[int]*ioConstraints)
	outputConstraints := make(map[int]*ioConstraints)

	inputCount := 0
	for _, constraint := range rule.GetParameterConstraints() {
		name := constraint.GetParameterName()
		if _, isInput, _, err := b.parseConstraintName(name); err != nil {
			return err
		} else if isInput {
			if index, _, _, _ := b.parseConstraintName(name); index+1 > inputCount {
				inputCount = index + 1
			}
		}
	}

	for _, constraint := range rule.GetParameterConstraints() {
		name := constraint.GetParameterName()

		if index, isInput, isAddress, err := b.parseConstraintName(name); err != nil {
			return err
		} else if isInput {
			b.setConstraint(inputConstraints, index, constraint, isAddress)
		} else {
			outputIndex := index - inputCount
			b.setConstraint(outputConstraints, outputIndex, constraint, isAddress)
		}
	}

	if err := b.validateConstraintCounts(inputConstraints, outputConstraints, tx); err != nil {
		return err
	}

	if err := b.validateInputs(inputConstraints, tx); err != nil {
		return err
	}

	return b.validateOutputs(outputConstraints, tx)
}

func (b *Btc) parseConstraintName(name string) (index int, isInput bool, isAddress bool, err error) {
	prefixes := map[string]struct{ isInput, isAddress bool }{
		"input_address_":  {true, true},
		"input_value_":    {true, false},
		"output_address_": {false, true},
		"output_value_":   {false, false},
	}

	for prefix, props := range prefixes {
		if strings.HasPrefix(name, prefix) {
			indexStr := strings.TrimPrefix(name, prefix)
			ind, er := strconv.Atoi(indexStr)
			if er != nil {
				return 0, false, false, fmt.Errorf("invalid constraint name: %s", name)
			}
			return ind, props.isInput, props.isAddress, nil
		}
	}

	return 0, false, false, fmt.Errorf("unknown constraint parameter name: %s", name)
}

func (b *Btc) setConstraint(constraints map[int]*ioConstraints, index int, constraint *types.ParameterConstraint, isAddress bool) {
	if constraints[index] == nil {
		constraints[index] = &ioConstraints{}
	}

	if isAddress {
		constraints[index].address = constraint
	} else {
		constraints[index].value = constraint
	}
}

func (b *Btc) validateConstraintCounts(inputConstraints, outputConstraints map[int]*ioConstraints, tx *wire.MsgTx) error {
	if len(inputConstraints) != len(tx.TxIn) {
		return fmt.Errorf("input count mismatch: rule has %d inputs, tx has %d inputs", len(inputConstraints), len(tx.TxIn))
	}

	if len(outputConstraints) != len(tx.TxOut) {
		return fmt.Errorf("output count mismatch: rule has %d outputs, tx has %d outputs", len(outputConstraints), len(tx.TxOut))
	}

	for i := 0; i < len(tx.TxIn); i += 1 {
		if constraints, exists := inputConstraints[i]; !exists || constraints.address == nil || constraints.value == nil {
			return fmt.Errorf("missing address or value constraint for input %d", i)
		}
	}

	for i := 0; i < len(tx.TxOut); i++ {
		if constraints, exists := outputConstraints[i]; !exists || constraints.address == nil || constraints.value == nil {
			return fmt.Errorf("missing address or value constraint for output %d", i)
		}
	}

	return nil
}

func (b *Btc) validateInputs(inputConstraints map[int]*ioConstraints, tx *wire.MsgTx) error {
	for i := range tx.TxIn {
		constraints := inputConstraints[i]
		addressConstraint := constraints.address.GetConstraint().GetFixedValue()
		if addressConstraint == "" {
			return fmt.Errorf("input %d address constraint is empty", i)
		}

		valueConstraint := constraints.value.GetConstraint().GetFixedValue()
		if valueConstraint == "" {
			return fmt.Errorf("input %d value constraint is empty", i)
		}
		if _, ok := big.NewInt(0).SetString(valueConstraint, 10); !ok {
			return fmt.Errorf("input %d value constraint is not a valid number: %s", i, valueConstraint)
		}
	}
	return nil
}

func (b *Btc) validateOutputs(outputConstraints map[int]*ioConstraints, tx *wire.MsgTx) error {
	for i, txOut := range tx.TxOut {
		constraints := outputConstraints[i]

		outputAddress, err := b.extractAddress(txOut)
		if err != nil {
			return fmt.Errorf("failed to extract address from output %d: %w", i, err)
		}

		outputAmount := big.NewInt(txOut.Value)

		if err := validateConstraint(constraints.address, outputAddress, compare.NewString); err != nil {
			return fmt.Errorf("output %d address validation failed: %w", i, err)
		}

		if err := validateConstraint(constraints.value, outputAmount, compare.NewBigInt); err != nil {
			return fmt.Errorf("output %d value validation failed: %w", i, err)
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
