package evm

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

type Evm struct {
	nativeSymbol string
}

func NewEvm(nativeSymbol string) *Evm {
	return &Evm{
		nativeSymbol: strings.ToLower(nativeSymbol),
	}
}

func (e *Evm) Evaluate(rule *types.Rule, txBytes []byte) error {
	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules suppoted, got: %s", rule.GetEffect().String())
	}

	r, err := util.ParseResource(rule.GetResource())
	if err != nil {
		return fmt.Errorf("failed to parse rule resource: %w", err)
	}

	txData, err := ethereum.DecodeUnsignedPayload(txBytes)
	if err != nil {
		return fmt.Errorf("failed to decode tx payload: %w", err)
	}
	tx := etypes.NewTx(txData)

	err = assertTarget(r, rule.GetTarget(), tx.To())
	if err != nil {
		return fmt.Errorf("failed to assert target: %w", err)
	}

	if r.ProtocolId == e.nativeSymbol {
		er := evaluateArgsNative(r, rule, tx)
		if er != nil {
			return fmt.Errorf("failed to Evaluate native: symbol=%s, error=%w", e.nativeSymbol, er)
		}
		return nil
	}

	er := assertArgsAbi(r, rule, tx.Data())
	if er != nil {
		return fmt.Errorf("failed to Evaluate ABI: %w", er)
	}
	return nil
}

func evaluateArgsNative(resource *types.ResourcePath, rule *types.Rule, tx *etypes.Transaction) error {
	err := assertArg[*big.Int](
		resource.ChainId,
		rule.GetParameterConstraints(),
		"amount",
		tx.Value(),
		bigIntFromString,
	)
}

func assertTarget(resource *types.ResourcePath, target *types.Target, to *common.Address) error {
	targetKind := target.GetTargetType()
	switch targetKind {
	case types.TargetType_TARGET_TYPE_ADDRESS:
		if to == nil || !addrEqual(*to, common.HexToAddress(target.GetAddress())) {
			toHex := "nil"
			if to != nil {
				toHex = to.Hex()
			}
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_target_address=%s",
				toHex,
				target.GetAddress(),
			)
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
		if to == nil || !addrEqual(*to, common.HexToAddress(resolvedAddr)) {
			toHex := "nil"
			if to != nil {
				toHex = to.Hex()
			}
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_magic_const_resolved=%s",
				toHex,
				resolvedAddr,
			)
		}
	}
	return fmt.Errorf("unknow target type: %s", targetKind.String())
}

func assertArgsAbi(resource *types.ResourcePath, rule *types.Rule, data []byte) error {
	filepath := path.Join("..", "abi", resource.ProtocolId+".json")

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

	method, ok := a.Methods[resource.FunctionId]
	if !ok {
		return fmt.Errorf("failed to find abi method: %s", resource.FunctionId)
	}

	const dataOffset = 4
	args, err := method.Inputs.Unpack(data[dataOffset:])
	if err != nil {
		return fmt.Errorf("failed to unpack abi args: %w", err)
	}

	for i, arg := range args {
		input := method.Inputs[i]
		switch actual := arg.(type) {
		case common.Address:
			er := assertArg[common.Address](
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				addrFromString,
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
				resource.GetChainId(),
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
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				bigIntFromString,
				func(_expected, _actual *big.Int) bool {
					return _expected.Cmp(_actual) == 0
				},
				func(_expected, _actual *big.Int) bool {
					cmp := _expected.Cmp(_actual)
					return cmp == -1 || cmp == 0
				},
				func(_expected, _actual *big.Int) bool {
					cmp := _expected.Cmp(_actual)
					return cmp == 1 || cmp == 0
				},
				// 'magic' should always fail for 'big.Int' type
				doFalse,
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case uint8:
			er := assertArg[uint8](
				resource.GetChainId(),
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
				resource.GetChainId(),
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
				resource.GetChainId(),
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

func addrFromString(hex string) (common.Address, error) {
	return common.HexToAddress(hex), nil
}

func bigIntFromString(val string) (*big.Int, error) {
	v, parseOk := new(big.Int).SetString(val, 10)
	if !parseOk {
		return nil, fmt.Errorf("failed to create big int: %s", val)
	}
	return v, nil
}

func equal[T comparable](a, b T) bool {
	return a == b
}

func assertArg[expectedT, actualT any](
	chain string,
	expectedList []*types.ParameterConstraint,
	expectedName string,
	actual actualT,
	makeExpectedFromString func(string) (expectedT, error),
	assertFixed func(expectedT, actualT) bool,
	assertMin func(expectedT, actualT) bool,
	assertMax func(expectedT, actualT) bool,
	assertMagic func(expectedT, actualT) bool,
) error {
	const magicAssetIdDefault = "default"

	for _, constraint := range expectedList {
		if constraint.GetParameterName() == expectedName {
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
	return fmt.Errorf("arg not found: %s", expectedName)
}
