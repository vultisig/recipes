package evm

import (
	"bytes"
	"fmt"
	"math/big"
	"path"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	abi_embed "github.com/vultisig/recipes/abi"
	"github.com/vultisig/recipes/engine/evm/compare"
	"github.com/vultisig/recipes/ethereum"
	"github.com/vultisig/recipes/resolver"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
)

type Evm struct {
	nativeSymbol string
	abi          map[protocolID]abi.ABI
}

func NewEvm(nativeSymbol string) (*Evm, error) {
	abis, err := loadAbiDir()
	if err != nil {
		return nil, fmt.Errorf("failed to load abi dir: %w", err)
	}

	return &Evm{
		nativeSymbol: strings.ToLower(nativeSymbol),
		abi:          abis,
	}, nil
}

func (e *Evm) Evaluate(rule *types.Rule, txBytes []byte) error {
	if rule.GetEffect().String() != types.Effect_EFFECT_ALLOW.String() {
		return fmt.Errorf("only allow rules supported, got: %s", rule.GetEffect().String())
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
		er := assertArgsNative(r, rule, tx)
		if er != nil {
			return fmt.Errorf("failed to Evaluate native: symbol=%s, error=%w", e.nativeSymbol, er)
		}
		return nil
	}

	if tx.Value() != nil && tx.Value().Sign() != 0 {
		return fmt.Errorf(
			"tx value must be zero for non-native: abi=%s, tx_value=%s",
			r.ProtocolId,
			tx.Value().String(),
		)
	}
	er := e.assertArgsAbi(r, rule, tx.Data())
	if er != nil {
		return fmt.Errorf("failed to Evaluate ABI: %w", er)
	}
	return nil
}

func assertArgsNative(resource *types.ResourcePath, rule *types.Rule, tx *etypes.Transaction) error {
	if resource.FunctionId != "transfer" {
		return fmt.Errorf(
			"only 'transfer' function supported for native: symbol=%s, function_id=%s",
			resource.ProtocolId,
			resource.FunctionId,
		)
	}

	if len(rule.GetParameterConstraints()) != 1 {
		return fmt.Errorf("expected 1 parameter constraint, got: %d", len(rule.GetParameterConstraints()))
	}

	err := assertArg(
		resource.ChainId,
		rule.GetParameterConstraints(),
		"amount",
		tx.Value(),
		compare.NewBigInt,
	)
	if err != nil {
		return fmt.Errorf("failed to assert amount arg (tx value): %w", err)
	}
	return nil
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
		return nil

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
		return nil

	default:
		return fmt.Errorf("unknown target type: %s", targetKind.String())
	}
}

type protocolID = string

func loadAbiDir() (map[protocolID]abi.ABI, error) {
	base := path.Join(".")

	entries, err := abi_embed.Dir.ReadDir(base)
	if err != nil {
		return nil, fmt.Errorf("failed to read abi dir: err=%w", err)
	}

	abis := make(map[string]abi.ABI)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		filepath := path.Join(base, entry.Name())
		file, er := abi_embed.Dir.Open(filepath)
		if er != nil {
			return nil, fmt.Errorf("failed to open abi json: path=%s, err=%w", filepath, er)
		}

		a, er := abi.JSON(file)
		_ = file.Close()
		if er != nil {
			return nil, fmt.Errorf("failed to parse abi json: %w", er)
		}

		abis[strings.TrimSuffix(entry.Name(), ".json")] = a
	}
	return abis, nil
}

func (e *Evm) assertArgsAbi(resource *types.ResourcePath, rule *types.Rule, data []byte) error {
	a, ok := e.abi[resource.ProtocolId]
	if !ok {
		return fmt.Errorf("failed to get abi: protocolId=%s", resource.ProtocolId)
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

	if len(rule.GetParameterConstraints()) != len(args) {
		// if some arg not found by name, assertArg returns the error below during assertion,
		// so there 2 (len check and get later) it's enough to determine that lists are equal
		return fmt.Errorf(
			"constraints must be same list as ABI args: constraints_len=%d, abi_args_len=%d",
			len(rule.GetParameterConstraints()),
			len(args),
		)
	}

	for i, arg := range args {
		input := method.Inputs[i]
		switch actual := arg.(type) {
		case string:
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				compare.NewString,
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case common.Address:
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				compare.NewAddress,
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case []common.Address:
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				compare.NewAddressSlice,
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case *big.Int:
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				compare.NewBigInt,
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case uint8:
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				compare.NewUint8,
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case bool:
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				compare.NewBool,
			)
			if er != nil {
				return fmt.Errorf("failed to assert: %w", er)
			}

		case [32]byte:
			er := assertArg(
				resource.GetChainId(),
				rule.GetParameterConstraints(),
				input.Name,
				actual,
				compare.NewBytes32,
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

func addrEqual(a, b common.Address) bool {
	return bytes.Equal(a[:], b[:])
}

func assertArg[T any](
	chain string,
	expectedList []*types.ParameterConstraint,
	expectedName string,
	actual T,
	makeComparer compare.Constructor[T],
) error {
	const magicAssetIdDefault = "default"

	for _, constraint := range expectedList {
		if constraint.GetParameterName() == expectedName {
			kind := constraint.GetConstraint().GetType()

			switch kind {
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
				if comparer.Magic(actual) {
					return nil
				}
				return fmt.Errorf(
					"failed to compare magic values: expected(resolved magic addr)=%v, actual(in tx)=%v",
					resolvedAddr,
					actual,
				)

			default:
				return fmt.Errorf("unknown constraint type: %s", constraint.GetConstraint().GetType())
			}
		}
	}
	return fmt.Errorf("arg not found: %s", expectedName)
}
