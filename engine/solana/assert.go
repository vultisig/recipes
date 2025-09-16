package solana

import (
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
	expectedRuleTarget, err := solana.PublicKeyFromBase58(targetRule.GetAddress())
	if err != nil {
		return fmt.Errorf("failed to parse `targetRule` address: %w", err)
	}

	targetKind := targetRule.GetTargetType()
	switch targetKind {
	case types.TargetType_TARGET_TYPE_ADDRESS:
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

func assertAccounts(constraints []*types.ParameterConstraint, msg solana.Message, accs []idlAccount) error {
	const constraintPrefix = "account_"
	for i, acc := range accs {
		actual, err := msg.Account(uint16(i))
		if err != nil {
			return fmt.Errorf("failed to get account %d: %w", i, err)
		}

		err = compare.AssertArg(
			common.Solana.String(),
			constraints,
			constraintPrefix+acc.Name,
			actual,
			solcmp.NewPubKey,
		)
		if err != nil {
			return fmt.Errorf("failed to assert: name=%s: %w", acc.Name, err)
		}
	}
	return nil
}

func assertArgs(constraints []*types.ParameterConstraint, data solana.Base58, args []idlArgument) error {
	const constraintPrefix = "arg_"
	decoder := bin.NewBorshDecoder(data)
	for _, arg := range args {
		switch arg.Type {
		case argU8:
			err := decodeAndAssert(decoder, constraints, constraintPrefix+arg.Name, compare.NewUint8)
			if err != nil {
				return fmt.Errorf("failed to decode & assert: %w", err)
			}

		case argU64:
			err := decodeAndAssert(decoder, constraints, constraintPrefix+arg.Name, compare.NewUint64)
			if err != nil {
				return fmt.Errorf("failed to decode & assert: %w", err)
			}

		case argPublicKey:
			err := decodeAndAssert(decoder, constraints, constraintPrefix+arg.Name, solcmp.NewPubKey)
			if err != nil {
				return fmt.Errorf("failed to decode & assert: %w", err)
			}
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
