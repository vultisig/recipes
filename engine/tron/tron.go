package tron

import (
	"fmt"
	"math/big"

	chaintron "github.com/vultisig/recipes/chain/tron"
	stdcompare "github.com/vultisig/recipes/engine/compare"
	"github.com/vultisig/recipes/resolver"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
)

// Tron represents the TRON engine implementation
type Tron struct {
	chain *chaintron.Chain
}

// NewTron creates a new Tron engine instance
func NewTron() *Tron {
	return &Tron{
		chain: chaintron.NewChain(),
	}
}

// Supports returns true if this engine supports the given chain
func (t *Tron) Supports(chain common.Chain) bool {
	return chain == common.Tron
}

// Evaluate validates a TRON transaction against policy rules
func (t *Tron) Evaluate(rule *types.Rule, txBytes []byte) error {
	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
	}

	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}

	decodedTx, err := t.chain.ParseTransactionBytes(txBytes)
	if err != nil {
		return fmt.Errorf("failed to parse TRON transaction: %w", err)
	}

	parsedTx, ok := decodedTx.(*chaintron.ParsedTronTransaction)
	if !ok {
		return fmt.Errorf("unexpected transaction type: %T", decodedTx)
	}

	rawData := parsedTx.GetRawData()
	if rawData == nil {
		return fmt.Errorf("transaction raw data is nil")
	}

	if len(rawData.Contract) == 0 {
		return fmt.Errorf("transaction has no contracts")
	}

	if len(rawData.Contract) != 1 {
		return fmt.Errorf("only single-contract transactions supported, got %d contracts", len(rawData.Contract))
	}

	contract := rawData.Contract[0]
	if contract.Type != "TransferContract" {
		return fmt.Errorf("only TransferContract is supported, got: %s", contract.Type)
	}

	if err := t.validateTarget(r, rule.GetTarget(), parsedTx); err != nil {
		return fmt.Errorf("failed to validate target: %w", err)
	}

	if err := t.validateParameterConstraints(r, rule.GetParameterConstraints(), parsedTx); err != nil {
		return fmt.Errorf("failed to validate parameter constraints: %w", err)
	}

	return nil
}

// validateTarget validates the transaction target against the rule target
func (t *Tron) validateTarget(resource *types.ResourcePath, target *types.Target, tx *chaintron.ParsedTronTransaction) error {
	if target == nil || target.GetTargetType() == types.TargetType_TARGET_TYPE_UNSPECIFIED {
		return nil
	}

	actualDestination := tx.To()

	switch target.GetTargetType() {
	case types.TargetType_TARGET_TYPE_ADDRESS:
		expectedAddress := target.GetAddress()
		if expectedAddress == "" {
			return fmt.Errorf("target address cannot be empty")
		}
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

		if actualDestination != resolvedAddr {
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_magic_const_resolved=%s",
				actualDestination,
				resolvedAddr,
			)
		}

	default:
		return fmt.Errorf("unsupported target type: %s", target.GetTargetType())
	}

	return nil
}

// validateParameterConstraints validates all parameter constraints
func (t *Tron) validateParameterConstraints(resource *types.ResourcePath, constraints []*types.ParameterConstraint, tx *chaintron.ParsedTronTransaction) error {
	for _, constraint := range constraints {
		paramName := constraint.GetParameterName()

		value, err := t.extractParameterValue(paramName, tx)
		if err != nil {
			return fmt.Errorf("failed to extract parameter %s: %w", paramName, err)
		}

		if err := t.assertArgsByType(resource.ChainId, paramName, value, constraints); err != nil {
			return fmt.Errorf("constraint validation failed for parameter %s: %w", paramName, err)
		}
	}
	return nil
}

// extractParameterValue extracts the actual value from transaction for the given parameter name
func (t *Tron) extractParameterValue(paramName string, tx *chaintron.ParsedTronTransaction) (interface{}, error) {
	switch paramName {
	case "recipient":
		return tx.To(), nil
	case "amount":
		return tx.GetAmount(), nil
	case "memo":
		return tx.GetMemo(), nil
	default:
		return nil, fmt.Errorf("unsupported parameter: %s", paramName)
	}
}

// assertArgsByType validates constraints using the appropriate comparator based on Go type
func (t *Tron) assertArgsByType(chainId, inputName string, arg interface{}, constraints []*types.ParameterConstraint) error {
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

