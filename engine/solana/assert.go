package solana

import (
	"bytes"
	"fmt"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/vultisig/recipes/engine/compare"
	solcmp "github.com/vultisig/recipes/engine/solana/compare"
	"github.com/vultisig/recipes/resolver"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/vultisig-go/common"
)

func assertTarget(
	resource *types.ResourcePath,
	targetRule *types.Target,
	actual solana.PublicKey,
) error {
	targetKind := targetRule.GetTargetType()
	switch targetKind {
	case types.TargetType_TARGET_TYPE_ADDRESS:
		address := targetRule.GetAddress()
		if address == "" {
			return fmt.Errorf("address cannot be empty for TARGET_TYPE_ADDRESS")
		}

		expectedRuleTarget, err := solana.PublicKeyFromBase58(address)
		if err != nil {
			return fmt.Errorf("failed to parse `targetRule` address: %w", err)
		}

		if !actual.Equals(expectedRuleTarget) {
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_target_address=%s",
				actual.String(),
				expectedRuleTarget.String(),
			)
		}
		return nil

	case types.TargetType_TARGET_TYPE_MAGIC_CONSTANT:
		resolve, er := resolver.NewMagicConstantRegistry().GetResolver(targetRule.GetMagicConstant())
		if er != nil {
			return fmt.Errorf(
				"failed to get resolver: magic_const=%s",
				targetRule.GetMagicConstant().String(),
			)
		}

		resolvedAddr, _, er := resolve.Resolve(
			targetRule.GetMagicConstant(),
			resource.ChainId,
			"default",
		)
		if er != nil {
			return fmt.Errorf(
				"failed to resolve magic const: value=%s, error=%w",
				targetRule.GetMagicConstant().String(),
				er,
			)
		}
		resolvedTarget, er := solana.PublicKeyFromBase58(resolvedAddr)
		if er != nil {
			return fmt.Errorf("failed to parse resolved targetRule: %w", er)
		}

		if !actual.Equals(resolvedTarget) {
			return fmt.Errorf(
				"tx target is wrong: tx_to=%s, rule_magic_const_resolved=%s",
				actual.String(),
				resolvedTarget.String(),
			)
		}
		return nil

	default:
		return fmt.Errorf("unknown targetRule type: %s", targetKind.String())
	}
}

func assertFuncSelector(expected []byte, data solana.Base58) error {
	if len(expected) == 0 {
		return fmt.Errorf("discriminator cannot be empty")
	}

	if len(data) < len(expected) {
		return fmt.Errorf(
			"instruction data too short: expected at least %d bytes for discriminator, got %d",
			len(expected),
			len(data),
		)
	}

	actual := []byte(data[:len(expected)])
	if !bytes.Equal(expected, actual) {
		return fmt.Errorf("function discriminator mismatch: expected %x, got %x", expected, actual)
	}

	return nil
}

func assertAccounts(constraints []*types.ParameterConstraint, msg solana.Message, accs []idlAccount) error {
	const constraintPrefix = "account_"

	inst := msg.Instructions[0]
	if len(accs) != len(inst.Accounts) {
		return fmt.Errorf(
			"account count mismatch: IDL has %d accounts, tx has %d accounts",
			len(accs),
			len(inst.Accounts),
		)
	}

	for i, acc := range accs {
		name := constraintPrefix + acc.Name

		actual, err := msg.Account(inst.Accounts[i])
		if err != nil {
			return fmt.Errorf("failed to get account %d: %w", i, err)
		}

		err = compare.AssertArg(
			common.Solana.String(),
			constraints,
			name,
			actual,
			solcmp.NewPubKey,
		)
		if err != nil {
			return fmt.Errorf("failed to assert: name=%s: %w", name, err)
		}
	}
	return nil
}

func assertArgs(
	constraints []*types.ParameterConstraint,
	data solana.Base58,
	args []idlArgument,
	discriminator []byte,
) error {
	const constraintPrefix = "arg_"

	err := assertFuncSelector(discriminator, data)
	if err != nil {
		return fmt.Errorf("discriminator validation failed: %w", err)
	}

	decoder := bin.NewBorshDecoder(data[len(discriminator):])
	for _, arg := range args {
		name := constraintPrefix + arg.Name

		switch arg.Type {
		case argU8:
			er := decodeAndAssert(decoder, constraints, name, compare.NewUint8)
			if er != nil {
				return fmt.Errorf("failed to decode & assert: %w", er)
			}

		case argU16:
			er := decodeAndAssert(decoder, constraints, name, compare.NewUint16)
			if er != nil {
				return fmt.Errorf("failed to decode & assert: %w", er)
			}

		case argU64:
			er := decodeAndAssert(decoder, constraints, name, compare.NewUint64)
			if er != nil {
				return fmt.Errorf("failed to decode & assert: %w", er)
			}

		case argPublicKey:
			er := decodeAndAssert(decoder, constraints, name, solcmp.NewPubKey)
			if er != nil {
				return fmt.Errorf("failed to decode & assert: %w", er)
			}

		// For vector types, decode the length first and then skip the vector elements
		// Falsy comparer means assert would pass only on ANY rule-constraint, for example, for FIXED it would fail
		case argVec:
			// add correct padding to the decoder offset
			var vecLen uint32
			er := decoder.Decode(&vecLen)
			if er != nil {
				return fmt.Errorf("failed to decode vector length for %s: %w", name, er)
			}

			er = compare.AssertArg(
				common.Solana.String(),
				constraints,
				name,
				struct{}{},
				compare.NewFalsy,
			)
			if er != nil {
				return fmt.Errorf("failed to assert vector: %w", er)
			}

		default:
			return fmt.Errorf("unsupported argument type: %s (name=%s)", arg.Type, arg.Name)
		}
	}
	return nil
}

func decodeAndAssert[T any](
	decoder *bin.Decoder,
	expectedList []*types.ParameterConstraint,
	expectedName string,
	makeComparer compare.Constructor[T],
) error {
	var err error
	var actual T
	err = decoder.Decode(&actual)
	if err != nil {
		return fmt.Errorf("failed to decode: name=%s: %w", expectedName, err)
	}

	err = compare.AssertArg(
		common.Solana.String(),
		expectedList,
		expectedName,
		actual,
		makeComparer,
	)
	if err != nil {
		return fmt.Errorf("failed to assert: name=%s: %w", expectedName, err)
	}
	return nil
}
