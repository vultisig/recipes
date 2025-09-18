package compare

import (
	"fmt"
	"regexp"

	"github.com/vultisig/recipes/resolver"
	"github.com/vultisig/recipes/types"
)

type Constructor[expectedT any] func(raw string) (Compare[expectedT], error)

type Compare[expectedT any] interface {
	Fixed(actual expectedT) bool
	Min(actual expectedT) bool
	Max(actual expectedT) bool
}

// Falsy embed to rewrite only required compare methods: for example, min/max for address must be false
type Falsy[T any] struct{}

func (f *Falsy[T]) Fixed(_ T) bool {
	return false
}
func (f *Falsy[T]) Min(_ T) bool {
	return false
}
func (f *Falsy[T]) Max(_ T) bool {
	return false
}

func AssertArg[T any](
	chain string,
	expectedList []*types.ParameterConstraint,
	expectedName string,
	actual T,
	makeComparer Constructor[T],
) error {
	const magicAssetIdDefault = "default"

	for _, constraint := range expectedList {
		if constraint.GetParameterName() == expectedName {
			kind := constraint.GetConstraint().GetType()

			switch kind {
			case types.ConstraintType_CONSTRAINT_TYPE_ANY:
				return nil

			case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
				comparer, err := makeComparer(constraint.GetConstraint().GetFixedValue())
				if err != nil {
					return fmt.Errorf(
						"failed to build exact fixed type from constraint: %s",
						constraint.GetConstraint().GetFixedValue(),
					)
				}
				if comparer.Fixed(actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare fixed values: expected=%v, actual=%v",
					constraint.GetConstraint().GetFixedValue(),
					actual,
				)

			case types.ConstraintType_CONSTRAINT_TYPE_MIN:
				comparer, err := makeComparer(constraint.GetConstraint().GetMinValue())
				if err != nil {
					return fmt.Errorf(
						"failed to build exact min type from constraint: %s",
						constraint.GetConstraint().GetMinValue(),
					)
				}
				if comparer.Min(actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare min values: expected=%v, actual=%v",
					constraint.GetConstraint().GetMinValue(),
					actual,
				)

			case types.ConstraintType_CONSTRAINT_TYPE_MAX:
				comparer, err := makeComparer(constraint.GetConstraint().GetMaxValue())
				if err != nil {
					return fmt.Errorf(
						"failed to build exact max type from constraint: %s",
						constraint.GetConstraint().GetMaxValue(),
					)
				}
				if comparer.Max(actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare max values: expected=%v, actual=%v",
					constraint.GetConstraint().GetMaxValue(),
					actual,
				)

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

				resolvedAddr, _, err := resolve.Resolve(
					constraint.GetConstraint().GetMagicConstantValue(),
					chain,
					magicAssetIdDefault,
				)
				if err != nil {
					return fmt.Errorf(
						"failed to resolve magic const: magic_const=%s",
						constraint.GetConstraint().GetMagicConstantValue().String(),
					)
				}

				comparer, err := makeComparer(resolvedAddr)
				if err != nil {
					return fmt.Errorf(
						"failed to build exact type from magic_const: resolved=%s",
						resolvedAddr,
					)
				}
				if comparer.Fixed(actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare magic values: expected(resolved magic addr)=%v, actual(in tx)=%v",
					resolvedAddr,
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
				return fmt.Errorf("unknown constraint type: %s", constraint.GetConstraint().GetType())
			}
		}
	}
	return fmt.Errorf("arg not found: %s", expectedName)
}
