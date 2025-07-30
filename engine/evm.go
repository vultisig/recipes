package engine

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/vultisig/recipes/ethereum"
	"github.com/vultisig/recipes/resolver"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
)

type evm struct{}

func newEvm() *evm {
	return &evm{}
}

const magicAssetIdDefault = "default"

func (e *evm) evaluate(rule *types.Rule, txBytes []byte) error {
	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules suppoted, got: %s", rule.GetEffect().String())
	}

	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}

	filepath := path.Join("..", "abi", r.ProtocolId+".json")

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

	txData, err := ethereum.DecodeUnsignedPayload(txBytes)
	if err != nil {
		return fmt.Errorf("failed to decode tx payload: %w", err)
	}
	tx := etypes.NewTx(txData)

	targetKind := rule.GetTarget().GetTargetType()
	switch targetKind {
	case types.TargetType_TARGET_TYPE_ADDRESS:
		if tx.To() == nil || !addrEqual(*tx.To(), common.HexToAddress(rule.GetTarget().GetAddress())) {
			toHex := "nil"
			if tx.To() != nil {
				toHex = tx.To().Hex()
			}
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_target_address=%s",
				toHex,
				rule.GetTarget().GetAddress(),
			)
		}
	case types.TargetType_TARGET_TYPE_MAGIC_CONSTANT:
		resolve, er := resolver.NewMagicConstantRegistry().GetResolver(rule.GetTarget().GetMagicConstant())
		if er != nil {
			return fmt.Errorf(
				"failed to get resolver: magic_const=%s",
				rule.GetTarget().GetMagicConstant().String(),
			)
		}

		resolvedAddr, _, er := resolve.Resolve(rule.GetTarget().GetMagicConstant(), r.ChainId, "default")
		if er != nil {
			return fmt.Errorf(
				"failed to resolve magic const: %s",
				rule.GetTarget().GetMagicConstant().String(),
			)
		}
		if tx.To() == nil || !addrEqual(*tx.To(), common.HexToAddress(resolvedAddr)) {
			toHex := "nil"
			if tx.To() != nil {
				toHex = tx.To().Hex()
			}
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_magic_const_resolved=%s",
				toHex,
				resolvedAddr,
			)
		}
	}

	const dataOffset = 4
	args, err := method.Inputs.Unpack(tx.Data()[dataOffset:])
	if err != nil {
		return fmt.Errorf("failed to unpack abi args: %w", err)
	}

	for i, arg := range args {
		input := method.Inputs[i]
		switch actual := arg.(type) {
		case common.Address:
			er := assertArg[common.Address](
				r.ChainId,
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				func(_expected string) (common.Address, error) {
					return common.HexToAddress(_expected), nil
				},
				addrEqual,
				// 'min' and 'max' should always fail for 'address' type
				doFalse[common.Address],
				doFalse[common.Address],
				addrEqual,
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case []common.Address:
			er := assertArg[[]common.Address](
				r.ChainId,
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				func(_expected string) ([]common.Address, error) {
					var addrs []common.Address
					for _, s := range strings.Split(_expected, ",") {
						addrs = append(addrs, common.HexToAddress(s))
					}
					return addrs, nil
				},
				func(_expected, _actual []common.Address) bool {
					if len(_expected) != len(_actual) {
						return false
					}
					for _i := range _expected {
						if !addrEqual(_expected[_i], _actual[_i]) {
							return false
						}
					}
					return true
				},
				// 'min', 'max', 'magic' should always fail for '[]address' type
				doFalse[[]common.Address],
				doFalse[[]common.Address],
				doFalse[[]common.Address],
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case *big.Int:
			er := assertArg[*big.Int](
				r.ChainId,
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
				func(_expected, _actual *big.Int) bool {
					return _expected.Cmp(_actual) == 0
				},
				func(_expected, _actual *big.Int) bool {
					return _expected.Cmp(_actual) == 1
				},
				func(_expected, _actual *big.Int) bool {
					return _expected.Cmp(_actual) == -1
				},
				// 'magic' should always fail for 'big.Int' type
				doFalse,
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case uint8:
			er := assertArg[uint8](
				r.ChainId,
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				func(_expected string) (uint8, error) {
					v, er := strconv.ParseUint(_expected, 10, 8)
					if er != nil {
						return 0, fmt.Errorf("failed to parse string to uint8: %s", _expected)
					}
					return uint8(v), nil
				},
				equal,
				func(_expected, _actual uint8) bool {
					return _expected <= _actual
				},
				func(_expected, _actual uint8) bool {
					return _expected >= _actual
				},
				// 'magic' should always fail for 'uint8' type
				doFalse,
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case bool:
			er := assertArg[bool](
				r.ChainId,
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				strconv.ParseBool,
				equal,
				// 'min', 'max', 'magic' should always fail for 'bool' type
				doFalse[bool],
				doFalse[bool],
				doFalse[bool],
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case [32]byte:
			er := assertArg[[32]byte](
				r.ChainId,
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				func(_expected string) ([32]byte, error) {
					b, er := base64.StdEncoding.DecodeString(_expected)
					if er != nil {
						return [32]byte{}, fmt.Errorf("failed to decode b64 string: %w", er)
					}
					if len(b) != 32 {
						return [32]byte{}, fmt.Errorf("len must be 32, got: %d", len(b))
					}
					return [32]byte(b), nil
				},
				func(_expected, _actual [32]byte) bool {
					return bytes.Equal(_expected[:], _actual[:])
				},
				// 'min', 'max', 'magic' should always fail for '[32]byte' type
				doFalse[[32]byte],
				doFalse[[32]byte],
				doFalse[[32]byte],
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

func addrEqual(a, b common.Address) bool {
	return bytes.Equal(a[:], b[:])
}

func equal[T comparable](a, b T) bool {
	return a == b
}

func assertArg[expectedT, actualT any](
	chain string,
	constraints []*types.ParameterConstraint,
	name string,
	actual actualT,
	makeExpectedFromString func(string) (expectedT, error),
	assertFixed func(expectedT, actualT) bool,
	assertMin func(expectedT, actualT) bool,
	assertMax func(expectedT, actualT) bool,
	assertMagic func(expectedT, actualT) bool,
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
				if assertFixed(expected, actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare fixed values: expected=%v, actual=%v",
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
				if assertMin(expected, actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare min values: expected=%v, actual=%v",
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
				if assertMax(expected, actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare max values: expected=%v, actual=%v",
					expected,
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

				expected, err := makeExpectedFromString(resolvedAddr)
				if err != nil {
					return fmt.Errorf(
						"failed to build exact type from magic_const: resolved=%s",
						resolvedAddr,
					)
				}
				if assertMagic(expected, actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare magic values: expected(resolved magic addr)=%v, actual(in tx)=%v",
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
