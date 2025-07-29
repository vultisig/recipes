package engine

import (
	"fmt"
	"math/big"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
)

type evm struct {
	chainID *big.Int
}

func newEvm(chainID *big.Int) *evm {
	return &evm{
		chainID: chainID,
	}
}

func (e *evm) evaluate(rule *types.Rule, tx []byte) error {
	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}

	filepath := path.Join("abi", r.ProtocolId+".json")
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open abi json: path=%s, err=%w", filepath, err)
	}
	defer func() {
		_ = file.Close()
	}()

	a, err := abi.JSON(file)
	if err != nil {
		return fmt.Errorf("failed to parse abi json: %w", err)
	}

	method, ok := a.Methods[r.FunctionId]
	if !ok {
		return fmt.Errorf("failed to find abi method: %s", r.FunctionId)
	}

	args, err := method.Inputs.Unpack(tx[4:])
	if err != nil {
		return fmt.Errorf("failed to unpack abi args: %w", err)
	}

	for i, arg := range args {
		input := method.Inputs[i]
		switch actual := arg.(type) {
		case common.Address:
			er := assertArg[common.Address](
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				func(_expected string) (common.Address, error) {
					return common.HexToAddress(_expected), nil
				},
				func(_expected common.Address, _actual common.Address) bool {
					return _expected.Cmp(_actual) == 0
				},
				// 'min' and 'max' should always fail for 'address' type
				doFalse[common.Address],
				doFalse[common.Address],
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case *big.Int:
			er := assertArg[*big.Int](
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				func(_expected string) (*big.Int, error) {
					v, parseOk := new(big.Int).SetString(_expected, 10)
					if !parseOk {
						return nil, fmt.Errorf("failed to create big int: %s", _expected)
					}
					return v, nil
				},
				func(_expected *big.Int, _actual *big.Int) bool {
					return _expected.Cmp(_actual) == 0
				},
				func(_expected *big.Int, _actual *big.Int) bool {
					return _expected.Cmp(_actual) == 1
				},
				func(_expected *big.Int, _actual *big.Int) bool {
					return _expected.Cmp(_actual) == -1
				},
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		default:
			return fmt.Errorf("unsupported arg type: %s", input.Type.String())
		}
	}
	return nil
}

func doFalse[T any](_, _ T) bool {
	return false
}

func assertArg[expectedT, actualT any](
	constraints []*types.ParameterConstraint,
	name string,
	actual actualT,
	makeExpectedFromString func(string) (expectedT, error),
	cmpFixed func(expectedT, actualT) bool,
	cmpMin func(expectedT, actualT) bool,
	cmpMax func(expectedT, actualT) bool,
) error {
	for _, constraint := range constraints {
		if constraint.GetParameterName() == name {
			kind := constraint.GetConstraint().GetType()

			switch kind {
			case types.ConstraintType_CONSTRAINT_TYPE_FIXED:
				expected, err := makeExpectedFromString(constraint.GetConstraint().GetFixedValue())
				if err != nil {
					return fmt.Errorf(
						"failed to build exact fixed type from constraint: %s",
						constraint.GetConstraint().GetValue(),
					)
				}
				if cmpFixed(expected, actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare fixed values: expected=%s, actual=%s",
					expected,
					actual,
				)

			case types.ConstraintType_CONSTRAINT_TYPE_MIN:
				expected, err := makeExpectedFromString(constraint.GetConstraint().GetMinValue())
				if err != nil {
					return fmt.Errorf(
						"failed to build exact min type from constraint: %s",
						constraint.GetConstraint().GetMinValue(),
					)
				}
				if cmpMin(expected, actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare min values: expected=%s, actual=%s",
					expected,
					actual,
				)

			case types.ConstraintType_CONSTRAINT_TYPE_MAX:
				expected, err := makeExpectedFromString(constraint.GetConstraint().GetMaxValue())
				if err != nil {
					return fmt.Errorf(
						"failed to build exact max type from constraint: %s",
						constraint.GetConstraint().GetMaxValue(),
					)
				}
				if cmpMax(expected, actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare max values: expected=%s, actual=%s",
					expected,
					actual,
				)

			default:
				return fmt.Errorf("unknown constraint type: %s", constraint.GetConstraint().GetType())
			}
		}
	}
	return fmt.Errorf("arg not found: %s", name)
}
