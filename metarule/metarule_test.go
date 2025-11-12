package metarule

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/recipes/sdk/evm"
	"github.com/vultisig/recipes/types"
	"github.com/vultisig/recipes/util"
	"github.com/vultisig/vultisig-go/common"
)

const testAddress = "4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM"

func TestTryFormat_NonMetaRule(t *testing.T) {
	metaRule := NewMetaRule()

	// Test with a complete rule that has function ID (not a meta-rule)
	rule := &types.Rule{
		Resource: "solana.sol.transfer", // Already has function ID
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, rule, result[0]) // Should return unchanged
}

func TestTryFormat_UnsupportedChain(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "unsupported.send", // Actually unsupported chain
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse chain id")
}

func TestTryFormat_SolanaSOLTransfer(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{}
	rule.Resource = "solana.send"
	rule.Target = &types.Target{
		TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
		Target: &types.Target_Address{
			Address: solana.SystemProgramID.String(),
		},
	}
	rule.ParameterConstraints = []*types.ParameterConstraint{
		{
			ParameterName: "asset",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
			},
		},
		{
			ParameterName: "from_address",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
			},
		},
		{
			ParameterName: "to_address",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				Value: &types.Constraint_FixedValue{
					FixedValue: testAddress,
				},
			},
		},
		{
			ParameterName: "amount",
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				Value: &types.Constraint_FixedValue{
					FixedValue: "1000000",
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	require.Len(t, result, 1)
	assert.Equal(t, "solana.system.transfer", result[0].Resource)
	assert.Equal(t, solana.SystemProgramID.String(), result[0].Target.GetAddress())
	assert.Len(t, result[0].ParameterConstraints, 3)

	paramNames := make([]string, len(result[0].ParameterConstraints))
	for i, param := range result[0].ParameterConstraints {
		paramNames[i] = param.ParameterName
	}
	assert.Contains(t, paramNames, "account_from")
	assert.Contains(t, paramNames, "account_to")
	assert.Contains(t, paramNames, "arg_lamports")
}

func TestTryFormat_SolanaSPLTokenTransfer(t *testing.T) {
	metaRule := NewMetaRule()

	const tokenMintAddress = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"
	rule := &types.Rule{
		Resource: "solana.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: tokenMintAddress,
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: tokenMintAddress,
					},
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testAddress,
					},
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testAddress,
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	require.Len(t, result, 2)
	assert.Equal(t, "solana.spl_token.transfer", result[0].Resource)
	assert.Equal(t, solana.TokenProgramID.String(), result[0].Target.GetAddress())
	assert.Len(t, result[0].ParameterConstraints, 4)
}

func TestTryFormat_SolanaMissingRecipientConstraint(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			// Missing to_address constraint
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find constraint: to_address")
}

func TestTryFormat_SolanaMissingAmountConstraint(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "RecipientAddress123",
					},
				},
			},
			// Missing amount constraint
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find constraint: amount")
}

func TestTryFormat_SolanaUnsupportedProtocol(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.stake", // Unsupported protocol
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: solana.SystemProgramID.String(),
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol id: stake")
}

func TestTryFormat_SolanaInvalidRecipientConstraintType(t *testing.T) {
	metaRule := NewMetaRule()

	const tokenMintAddress = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"
	rule := &types.Rule{
		Resource: "solana.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: tokenMintAddress,
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: tokenMintAddress,
					},
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testAddress,
					},
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "`to_address` must be fixed constraint for spl token transfer")
}

const testEVMChain = "ethereum"
const testRecipientAddress = "0x1234567890abcdef1234567890abcdef12345678"
const testTokenAddress = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"

func TestHandleEVM_NativeTransfer(t *testing.T) {
	metaRule := NewMetaRule()

	in := &types.Rule{
		Resource: fmt.Sprintf("%s.send", testEVMChain),
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: evm.ZeroAddress.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testRecipientAddress,
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1",
					},
				},
			},
		},
	}

	r, err := util.ParseResource(in.GetResource())
	require.NoError(t, err)

	result, err := metaRule.handleEVM(in, r)
	require.NoError(t, err)
	assert.NotNil(t, result)
	require.Len(t, result, 1)

	chain, _ := common.FromString(testEVMChain)
	nativeSymbol, _ := chain.NativeSymbol()
	expectedResource := fmt.Sprintf("%s.%s.transfer", testEVMChain, strings.ToLower(nativeSymbol))
	assert.Equal(t, expectedResource, result[0].Resource)
	assert.Equal(t, testRecipientAddress, result[0].Target.GetAddress())
	assert.Equal(t, "amount", result[0].ParameterConstraints[0].ParameterName)
}

func TestHandleEVM_ERC20Transfer(t *testing.T) {
	metaRule := NewMetaRule()

	in := &types.Rule{
		Resource: fmt.Sprintf("%s.send", testEVMChain),
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: testTokenAddress,
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testTokenAddress,
					},
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testRecipientAddress,
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1",
					},
				},
			},
		},
	}

	r, err := util.ParseResource(in.GetResource())
	require.NoError(t, err)

	result, err := metaRule.handleEVM(in, r)
	require.NoError(t, err)
	assert.NotNil(t, result)
	require.Len(t, result, 1)

	expectedResource := fmt.Sprintf("%s.erc20.transfer", testEVMChain)
	assert.Equal(t, expectedResource, result[0].Resource)
	assert.Equal(t, in.Target, result[0].Target)

	paramNames := make([]string, len(result[0].ParameterConstraints))
	for i, param := range result[0].ParameterConstraints {
		paramNames[i] = param.ParameterName
	}
	assert.Contains(t, paramNames, "recipient")
	assert.Contains(t, paramNames, "amount")
}

func TestTryFormat_EVM_NonMetaRule(t *testing.T) {
	metaRule := NewMetaRule()

	in := &types.Rule{
		Resource: fmt.Sprintf("%s.uniswapV2_router.swap", testEVMChain),
	}

	result, err := metaRule.TryFormat(in)
	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, in, result[0]) // Returned unchanged
}

func TestHandleEVM_MissingAmount(t *testing.T) {
	metaRule := NewMetaRule()

	in := &types.Rule{
		Resource: fmt.Sprintf("%s.send", testEVMChain),
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: testRecipientAddress,
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testRecipientAddress,
					},
				},
			},
		},
	}

	r, err := util.ParseResource(in.GetResource())
	require.NoError(t, err)

	_, err = metaRule.handleEVM(in, r)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find constraint: amount")
}

func TestHandleEVM_MissingRecipient(t *testing.T) {
	metaRule := NewMetaRule()

	in := &types.Rule{
		Resource: fmt.Sprintf("%s.send", testEVMChain),
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: testRecipientAddress,
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1",
					},
				},
			},
		},
	}

	r, err := util.ParseResource(in.GetResource())
	require.NoError(t, err)

	_, err = metaRule.handleEVM(in, r)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find constraint: to_address")
}

func TestHandleEVM_InvalidChain(t *testing.T) {
	metaRule := NewMetaRule()

	in := &types.Rule{
		Resource: "invalidchain.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: evm.ZeroAddress.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1",
					},
				},
			},
		},
	}

	r, err := util.ParseResource(in.GetResource())
	require.NoError(t, err)

	_, err = metaRule.handleEVM(in, r)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid chainID")
}

func TestHandleEVM_InvalidTargetTypeForNative(t *testing.T) {
	metaRule := NewMetaRule()

	in := &types.Rule{
		Resource: fmt.Sprintf("%s.send", testEVMChain),
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: evm.ZeroAddress.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_UNSPECIFIED,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1",
					},
				},
			},
		},
	}

	r, err := util.ParseResource(in.GetResource())
	require.NoError(t, err)

	_, err = metaRule.handleEVM(in, r)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid constraint type for `to_address`")
}

func TestHandleEVM_MagicConstantTargetForNative(t *testing.T) {
	metaRule := NewMetaRule()

	in := &types.Rule{
		Resource: fmt.Sprintf("%s.send", testEVMChain),
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: evm.ZeroAddress.String(),
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT,
					Value: &types.Constraint_MagicConstantValue{
						MagicConstantValue: types.MagicConstant_VULTISIG_TREASURY,
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1",
					},
				},
			},
		},
	}

	r, err := util.ParseResource(in.GetResource())
	require.NoError(t, err)

	result, err := metaRule.handleEVM(in, r)
	require.NoError(t, err)
	assert.NotNil(t, result)
	require.Len(t, result, 1)

	chain, _ := common.FromString(testEVMChain)
	nativeSymbol, _ := chain.NativeSymbol()
	expectedResource := fmt.Sprintf("%s.%s.transfer", testEVMChain, strings.ToLower(nativeSymbol))
	assert.Equal(t, expectedResource, result[0].Resource)
	assert.Equal(t, types.MagicConstant_VULTISIG_TREASURY, result[0].Target.GetMagicConstant())
}

func TestTryFormat_EvmSwap(t *testing.T) {
	const (
		fromAddress       = "0xcB9B049B9c937acFDB87EeCfAa9e7f2c51E754f5"
		fromAmount        = "1000000"
		weth              = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" // WETH
		usdc              = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" // USDC
		routerAddr        = "0x111111125421ca6dc452d289314280a0f8842a65"
		expectedResource0 = "ethereum.erc20.approve"
		expectedResource1 = "ethereum.routerV6_1inch.swap"
	)

	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ethereum.swap",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "from_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: usdc,
					},
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAddress,
					},
				},
			},
			{
				ParameterName: "from_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAmount,
					},
				},
			},
			{
				ParameterName: "to_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: weth,
					},
				},
			},
			{
				ParameterName: "to_chain",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: common.Ethereum.String(),
					},
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAddress,
					},
				},
			},
		},
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: routerAddr,
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 2)

	assert.Equal(t, expectedResource1, result[0].Resource)
	assert.Equal(t, types.TargetType_TARGET_TYPE_ADDRESS, result[0].Target.TargetType)
	assert.Equal(t, routerAddr, result[0].GetTarget().GetAddress())
	require.Len(t, result[0].ParameterConstraints, 9)

	paramByName := make(map[string]*types.ParameterConstraint)
	for _, param := range result[0].ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	assert.Contains(t, paramByName, "desc.srcToken")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["desc.srcToken"].Constraint.Type)
	assert.Equal(t, usdc, paramByName["desc.srcToken"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "desc.dstToken")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["desc.dstToken"].Constraint.Type)
	assert.Equal(t, weth, paramByName["desc.dstToken"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "desc.srcReceiver")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["desc.srcReceiver"].Constraint.Type)

	assert.Contains(t, paramByName, "desc.dstReceiver")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["desc.dstReceiver"].Constraint.Type)
	assert.Equal(t, fromAddress, paramByName["desc.dstReceiver"].Constraint.GetFixedValue())

	assert.Equal(t, expectedResource0, result[1].Resource)
	assert.Equal(t, types.TargetType_TARGET_TYPE_ADDRESS, result[1].Target.TargetType)
	assert.Equal(t, usdc, result[1].GetTarget().GetAddress())
	require.Len(t, result[1].ParameterConstraints, 2)

	paramByName = make(map[string]*types.ParameterConstraint)
	for _, param := range result[1].ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	assert.Contains(t, paramByName, "amount")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_MIN, paramByName["amount"].Constraint.Type)
	assert.Equal(t, fromAmount, paramByName["amount"].Constraint.GetMinValue())

	assert.Contains(t, paramByName, "spender")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["spender"].Constraint.Type)
	assert.Equal(t, routerAddr, paramByName["spender"].Constraint.GetFixedValue())
}

func TestTryFormat_BitcoinSwap(t *testing.T) {
	const (
		fromAddress      = "bc1qw589q7vva3wxju9zxz8gt59pfz2frwsazglsj8"
		toAddress        = "0xcB9B049B9c937acFDB87EeCfAa9e7f2c51E754f5"
		fromAmount       = "1000000"
		toChain          = "ethereum"
		toAsset          = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
		expectedResource = "bitcoin.btc.transfer"
	)

	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "bitcoin.swap",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "from_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "",
					},
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAddress,
					},
				},
			},
			{
				ParameterName: "from_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAmount,
					},
				},
			},
			{
				ParameterName: "to_chain",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toChain,
					},
				},
			},
			{
				ParameterName: "to_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toAsset,
					},
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toAddress,
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, expectedResource, result[0].Resource)
	assert.Equal(t, types.TargetType_TARGET_TYPE_UNSPECIFIED, result[0].Target.TargetType)
	require.Len(t, result[0].ParameterConstraints, 5)

	paramByName := make(map[string]*types.ParameterConstraint)
	for _, param := range result[0].ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	assert.Contains(t, paramByName, "output_address_0")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT, paramByName["output_address_0"].Constraint.Type)
	assert.Equal(t, types.MagicConstant_THORCHAIN_VAULT, paramByName["output_address_0"].Constraint.GetMagicConstantValue())

	assert.Contains(t, paramByName, "output_value_0")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["output_value_0"].Constraint.Type)
	assert.Equal(t, fromAmount, paramByName["output_value_0"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "output_address_1")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["output_address_1"].Constraint.Type)
	assert.Equal(t, fromAddress, paramByName["output_address_1"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "output_value_1")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["output_value_1"].Constraint.Type)

	assert.Contains(t, paramByName, "output_data_2")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_REGEXP, paramByName["output_data_2"].Constraint.Type)
	regexpValue := paramByName["output_data_2"].Constraint.GetRegexpValue()
	assert.Equal(t, regexpValue, fmt.Sprintf(
		`^=:ETH\.USDC:%s:.*`,
		toAddress,
	))
}

func TestTryFormat_BitcoinSend(t *testing.T) {
	const (
		changeAddress    = "bc1qchange123456789abcdef1234567890abcdef12"
		recipientAddress = "bc1qrecipient123456789abcdef1234567890abcd"
		amount           = "500000"
		asset            = "BTC"
		expectedResource = "bitcoin.btc.transfer"
	)

	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "bitcoin.send",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: asset,
					},
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: recipientAddress,
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: amount,
					},
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: changeAddress,
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, expectedResource, result[0].Resource)
	assert.Equal(t, types.TargetType_TARGET_TYPE_UNSPECIFIED, result[0].Target.TargetType)
	require.Len(t, result[0].ParameterConstraints, 4)

	paramByName := make(map[string]*types.ParameterConstraint)
	for _, param := range result[0].ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	assert.Contains(t, paramByName, "output_address_0")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["output_address_0"].Constraint.Type)
	assert.Equal(t, recipientAddress, paramByName["output_address_0"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "output_value_0")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["output_value_0"].Constraint.Type)
	assert.Equal(t, amount, paramByName["output_value_0"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "output_address_1")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["output_address_1"].Constraint.Type)
	assert.Equal(t, changeAddress, paramByName["output_address_1"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "output_value_1")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["output_value_1"].Constraint.Type)
}

func TestTryFormat_SolanaSwap(t *testing.T) {
	const (
		fromAsset           = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // USDC mint
		fromAddress         = "4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM"
		fromAmount          = "1000000"
		toChain             = "solana"
		toAsset             = "So11111111111111111111111111111111111111112" // WSOL
		toAddress           = "5w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKN"
		jupiterAddress      = "JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4"
		jupiterEvent        = "D8cy77BBepLMngZx6ZukaTff5hCt1HrWyKk3Hnd9oitf"
		tokenProgramAddress = "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"
	)

	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.swap",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "from_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAsset,
					},
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAddress,
					},
				},
			},
			{
				ParameterName: "from_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAmount,
					},
				},
			},
			{
				ParameterName: "to_chain",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toChain,
					},
				},
			},
			{
				ParameterName: "to_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toAsset,
					},
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toAddress,
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 9) // 2 system transfers + 2 ATA create (source + dest) + 1 syncNative + 1 SPL token approve + 1 closeAccount + 2 Jupiter (route + shared_accounts_route)

	// First two rules should be system transfers
	assert.Equal(t, "solana.system.transfer", result[0].Resource)
	assert.Equal(t, "solana.system.transfer", result[1].Resource)

	// Third rule should be source ATA create (USDC)
	sourceAtaRule := result[2]
	assert.Equal(t, "solana.associated_token_account.create", sourceAtaRule.Resource)
	assert.Equal(t, types.Effect_EFFECT_ALLOW, sourceAtaRule.Effect)
	assert.Equal(t, "ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL", sourceAtaRule.Target.GetAddress())

	// Fourth rule should be destination ATA create (WSOL)
	destAtaRule := result[3]
	assert.Equal(t, "solana.associated_token_account.create", destAtaRule.Resource)
	assert.Equal(t, types.Effect_EFFECT_ALLOW, destAtaRule.Effect)
	assert.Equal(t, "ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL", destAtaRule.Target.GetAddress())

	// Fifth rule should be syncNative
	assert.Equal(t, "solana.spl_token.syncNative", result[4].Resource)

	// Sixth rule should be SPL token approve
	approveRule := result[5]
	assert.Equal(t, "solana.spl_token.approve", approveRule.Resource)
	assert.Equal(t, types.Effect_EFFECT_ALLOW, approveRule.Effect)
	assert.Equal(t, "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", approveRule.Target.GetAddress())
	require.Len(t, approveRule.ParameterConstraints, 4)

	paramByName := make(map[string]*types.ParameterConstraint)
	for _, param := range approveRule.ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	assert.Contains(t, paramByName, "account_source")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_source"].Constraint.Type)
	expectedSourceATA, err := DeriveATA(fromAddress, fromAsset)
	require.NoError(t, err)
	assert.Equal(t, expectedSourceATA, paramByName["account_source"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "account_delegate")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_delegate"].Constraint.Type)
	assert.Equal(t, jupiterAddress, paramByName["account_delegate"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "account_owner")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_owner"].Constraint.Type)
	assert.Equal(t, fromAddress, paramByName["account_owner"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "arg_amount")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["arg_amount"].Constraint.Type)
	assert.Equal(t, fromAmount, paramByName["arg_amount"].Constraint.GetFixedValue())

	// Seventh rule should be SPL token closeAccount
	closeAccountRule := result[6]
	assert.Equal(t, "solana.spl_token.closeAccount", closeAccountRule.Resource)
	assert.Equal(t, tokenProgramAddress, closeAccountRule.Target.GetAddress())
	require.Len(t, closeAccountRule.ParameterConstraints, 3)

	// Eighth rule should be Jupiter route
	jupiterRouteRule := result[7]
	assert.Equal(t, "solana.jupiter_aggregatorv6.route", jupiterRouteRule.Resource)
	assert.Equal(t, jupiterAddress, jupiterRouteRule.Target.GetAddress())
	require.Len(t, jupiterRouteRule.ParameterConstraints, 14)

	// Ninth rule should be Jupiter shared_accounts_route
	jupiterSharedAccountsRouteRule := result[8]
	assert.Equal(t, "solana.jupiter_aggregatorv6.shared_accounts_route", jupiterSharedAccountsRouteRule.Resource)
	assert.Equal(t, jupiterAddress, jupiterSharedAccountsRouteRule.Target.GetAddress())
	require.Len(t, jupiterSharedAccountsRouteRule.ParameterConstraints, 19)

	// Verify route rule parameters
	jupiterRule := jupiterRouteRule

	paramByName = make(map[string]*types.ParameterConstraint)
	for _, param := range jupiterRule.ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	assert.Contains(t, paramByName, "account_token_program")
	assert.Equal(t, tokenProgramAddress, paramByName["account_token_program"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "account_user_transfer_authority")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["account_user_transfer_authority"].Constraint.Type)

	assert.Contains(t, paramByName, "account_user_source_token_account")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_user_source_token_account"].Constraint.Type)
	expectedUserSourceATA, err := DeriveATA(fromAddress, fromAsset)
	require.NoError(t, err)
	assert.Equal(t, expectedUserSourceATA, paramByName["account_user_source_token_account"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "account_user_destination_token_account")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_user_destination_token_account"].Constraint.Type)
	expectedUserDestATA, err := DeriveATA(toAddress, toAsset)
	require.NoError(t, err)
	assert.Equal(t, expectedUserDestATA, paramByName["account_user_destination_token_account"].Constraint.GetFixedValue())

	// Verify the newly added account_destination_token_account constraint
	// This is optional and can be Jupiter address, user's ATA, or other values
	assert.Contains(t, paramByName, "account_destination_token_account")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["account_destination_token_account"].Constraint.Type)

	assert.Contains(t, paramByName, "account_destination_mint")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_destination_mint"].Constraint.Type)
	assert.Equal(t, toAsset, paramByName["account_destination_mint"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "account_event_authority")
	assert.Equal(t, jupiterEvent, paramByName["account_event_authority"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "account_program")
	assert.Equal(t, jupiterAddress, paramByName["account_program"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "arg_route_plan")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["arg_route_plan"].Constraint.Type)

	assert.Contains(t, paramByName, "arg_in_amount")
	assert.Equal(t, fromAmount, paramByName["arg_in_amount"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "arg_quoted_out_amount")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["arg_quoted_out_amount"].Constraint.Type)

	assert.Contains(t, paramByName, "arg_slippage_bps")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["arg_slippage_bps"].Constraint.Type)

	assert.Contains(t, paramByName, "arg_platform_fee_bps")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["arg_platform_fee_bps"].Constraint.Type)
}

func TestTryFormat_SolanaSwapNativeAsset(t *testing.T) {
	const (
		fromAsset   = "" // Empty string for native SOL
		fromAddress = "4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM"
		fromAmount  = "1000000"
		toChain     = "solana"
		toAsset     = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // USDC mint
		toAddress   = "5w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKN"
	)

	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.swap",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "from_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAsset,
					},
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAddress,
					},
				},
			},
			{
				ParameterName: "from_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAmount,
					},
				},
			},
			{
				ParameterName: "to_chain",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toChain,
					},
				},
			},
			{
				ParameterName: "to_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toAsset,
					},
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toAddress,
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 9) // 2 system transfers + 2 ATA create + 1 syncNative + 1 WSOL approve + 1 closeAccount + 2 Jupiter (route + shared_accounts_route) for native SOL to SPL token

	// First two rules should be system transfers
	assert.Equal(t, "solana.system.transfer", result[0].Resource)
	assert.Equal(t, "solana.system.transfer", result[1].Resource)

	// Third rule should be source ATA create (WSOL)
	assert.Equal(t, "solana.associated_token_account.create", result[2].Resource)

	// Fourth rule should be destination ATA create (USDC - no source ATA needed for native SOL)
	ataRule := result[3]
	assert.Equal(t, "solana.associated_token_account.create", ataRule.Resource)
	assert.Equal(t, "ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL", ataRule.Target.GetAddress())

	// Fifth rule should be syncNative for WSOL
	assert.Equal(t, "solana.spl_token.syncNative", result[4].Resource)

	// Sixth rule should be WSOL approve
	approveRule := result[5]
	assert.Equal(t, "solana.spl_token.approve", approveRule.Resource)
	assert.Equal(t, "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", approveRule.Target.GetAddress())

	// Seventh rule should be SPL token closeAccount
	closeAccountRule := result[6]
	assert.Equal(t, "solana.spl_token.closeAccount", closeAccountRule.Resource)

	// Eighth rule should be Jupiter route
	jupiterRouteRule := result[7]
	assert.Equal(t, "solana.jupiter_aggregatorv6.route", jupiterRouteRule.Resource)

	// Ninth rule should be Jupiter shared_accounts_route
	jupiterSharedAccountsRouteRule := result[8]
	assert.Equal(t, "solana.jupiter_aggregatorv6.shared_accounts_route", jupiterSharedAccountsRouteRule.Resource)
}

const testXRPAddress = "rw2ciyaNshpHe7bCHo4bRWq6pqqynnWKQg"
const testXRPVaultAddress = "rfunGxj8FWbK3iYuxQvYMA9LGhJ9mYFuss"

func TestTryFormat_XRPSend(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ripple.send",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testXRPAddress,
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, "ripple.xrp.transfer", result[0].Resource)
	assert.Equal(t, types.TargetType_TARGET_TYPE_UNSPECIFIED, result[0].Target.TargetType)
	require.Len(t, result[0].ParameterConstraints, 2)

	paramByName := make(map[string]*types.ParameterConstraint)
	for _, param := range result[0].ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	assert.Contains(t, paramByName, "recipient")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["recipient"].Constraint.Type)
	assert.Equal(t, testXRPAddress, paramByName["recipient"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "amount")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["amount"].Constraint.Type)
	assert.Equal(t, "1000000", paramByName["amount"].Constraint.GetFixedValue())
}

func TestTryFormat_XRPSwap(t *testing.T) {
	const (
		fromAddress      = testXRPAddress
		fromAmount       = "2000000" // 2 XRP in drops
		toChain          = "ethereum"
		toAsset          = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" // USDC
		toAddress        = "0x742d35Cc6634C0532925a3b8D5c9E0B0Cf8a6b"
		expectedResource = "ripple.swap"
	)

	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ripple.swap",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "from_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "", // Empty for native XRP
					},
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAddress,
					},
				},
			},
			{
				ParameterName: "from_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAmount,
					},
				},
			},
			{
				ParameterName: "to_chain",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toChain,
					},
				},
			},
			{
				ParameterName: "to_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toAsset,
					},
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toAddress,
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, expectedResource, result[0].Resource)
	assert.Equal(t, types.TargetType_TARGET_TYPE_MAGIC_CONSTANT, result[0].Target.TargetType)
	assert.Equal(t, types.MagicConstant_THORCHAIN_VAULT, result[0].Target.GetMagicConstant())
	require.Len(t, result[0].ParameterConstraints, 3)

	paramByName := make(map[string]*types.ParameterConstraint)
	for _, param := range result[0].ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	assert.Contains(t, paramByName, "recipient")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT, paramByName["recipient"].Constraint.Type)
	assert.Equal(t, types.MagicConstant_THORCHAIN_VAULT, paramByName["recipient"].Constraint.GetMagicConstantValue())

	assert.Contains(t, paramByName, "amount")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["amount"].Constraint.Type)
	assert.Equal(t, fromAmount, paramByName["amount"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "memo")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_REGEXP, paramByName["memo"].Constraint.Type)
	regexpValue := paramByName["memo"].Constraint.GetRegexpValue()
	assert.Equal(t, fmt.Sprintf("^=:ETH\\.USDC:%s:.*", toAddress), regexpValue)
}

func TestTryFormat_XRPSend_MissingRecipient(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ripple.send",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find constraint: to_address")
}

func TestTryFormat_XRPSend_MissingAmount(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ripple.send",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testXRPAddress,
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find constraint: amount")
}

func TestTryFormat_XRPSwap_MissingFromAsset(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ripple.swap",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testXRPAddress,
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find constraint: from_asset")
}

func TestTryFormat_XRPSend_MagicConstantRecipient(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ripple.send",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT,
					Value: &types.Constraint_MagicConstantValue{
						MagicConstantValue: types.MagicConstant_VULTISIG_TREASURY,
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, "ripple.xrp.transfer", result[0].Resource)
	assert.Equal(t, types.TargetType_TARGET_TYPE_UNSPECIFIED, result[0].Target.TargetType)
	require.Len(t, result[0].ParameterConstraints, 2)

	paramByName := make(map[string]*types.ParameterConstraint)
	for _, param := range result[0].ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	assert.Contains(t, paramByName, "recipient")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT, paramByName["recipient"].Constraint.Type)
	assert.Equal(t, types.MagicConstant_VULTISIG_TREASURY, paramByName["recipient"].Constraint.GetMagicConstantValue())
}

func TestTryFormat_XRP_NonMetaRule(t *testing.T) {
	metaRule := NewMetaRule()

	// Test with a complete XRP rule that has function ID (not a meta-rule)
	rule := &types.Rule{
		Resource: "ripple.thorchain_swap.swap", // Already has function ID
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: testXRPVaultAddress,
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "2000000",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Equal(t, rule, result[0]) // Should return unchanged
}

func TestTryFormat_XRP_UnsupportedProtocol(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "ripple.stake", // Unsupported protocol
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol id for XRP: stake")
}

func TestDeriveATA(t *testing.T) {
	// https://solscan.io/account/EDT9FrASLP4gRKFnka5h4vgVBHzrmTZxgxmGa4G4vect
	owner := "4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM"
	mint := "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"
	expectedATA := "EDT9FrASLP4gRKFnka5h4vgVBHzrmTZxgxmGa4G4vect"

	ata, err := DeriveATA(owner, mint)
	require.NoError(t, err)
	assert.Equal(t, expectedATA, ata)
}

func TestDeriveATA_WSOLDestination(t *testing.T) {
	owner := "4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM"
	mint := "So11111111111111111111111111111111111111112"
	expectedATA := "R97cgCoxcqrUaaW7wg8drNBiLkicyQQHmct7pY8tdMR"

	ata, err := DeriveATA(owner, mint)
	require.NoError(t, err)
	assert.Equal(t, expectedATA, ata)
}

func TestCreateJupiterRule_StrictConstraints(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "solana.swap",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "from_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "",
					},
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM",
					},
				},
			},
			{
				ParameterName: "from_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "50000000",
					},
				},
			},
			{
				ParameterName: "to_chain",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "solana",
					},
				},
			},
			{
				ParameterName: "to_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
					},
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 9, "should have 2 system transfers + 2 ATA creates + 1 syncNative + 1 approve + 1 closeAccount + 2 Jupiter (route + shared_accounts_route)")

	// First two rules should be system transfers for source and destination ATA funding
	assert.Equal(t, "solana.system.transfer", result[0].Resource)
	assert.Equal(t, "solana.system.transfer", result[1].Resource)

	// Next two rules should be ATA create for source (WSOL) and destination (USDC)
	assert.Equal(t, "solana.associated_token_account.create", result[2].Resource)
	assert.Equal(t, "solana.associated_token_account.create", result[3].Resource)

	// Fifth rule should be syncNative for WSOL
	assert.Equal(t, "solana.spl_token.syncNative", result[4].Resource)

	// Sixth rule should be SPL token approve for WSOL
	approveRule := result[5]
	assert.Equal(t, "solana.spl_token.approve", approveRule.Resource)
	assert.Equal(t, "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", approveRule.Target.GetAddress())

	// Seventh rule should be SPL token closeAccount
	closeAccountRule := result[6]
	assert.Equal(t, "solana.spl_token.closeAccount", closeAccountRule.Resource)
	assert.Equal(t, "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", closeAccountRule.Target.GetAddress())

	// Find route rule
	var routeRule *types.Rule
	for _, r := range result {
		if r.Resource == "solana.jupiter_aggregatorv6.route" {
			routeRule = r
			break
		}
	}
	require.NotNil(t, routeRule, "should have route rule")

	// Build parameter map for easier assertions
	paramByName := make(map[string]*types.ParameterConstraint)
	for _, param := range routeRule.ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	// Verify all account constraints are present
	assert.Contains(t, paramByName, "account_token_program")
	assert.Contains(t, paramByName, "account_user_transfer_authority")
	assert.Contains(t, paramByName, "account_user_source_token_account")
	assert.Contains(t, paramByName, "account_user_destination_token_account")
	assert.Contains(t, paramByName, "account_destination_token_account")
	assert.Contains(t, paramByName, "account_destination_mint")
	assert.Contains(t, paramByName, "account_platform_fee_account")
	assert.Contains(t, paramByName, "account_event_authority")
	assert.Contains(t, paramByName, "account_program")

	// Verify argument constraints are present
	assert.Contains(t, paramByName, "arg_route_plan")
	assert.Contains(t, paramByName, "arg_in_amount")
	assert.Contains(t, paramByName, "arg_quoted_out_amount")
	assert.Contains(t, paramByName, "arg_slippage_bps")
	assert.Contains(t, paramByName, "arg_platform_fee_bps")

	// Verify FIXED constraints
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_token_program"].Constraint.Type)
	assert.Equal(t, "TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA", paramByName["account_token_program"].Constraint.GetFixedValue())

	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_destination_mint"].Constraint.Type)
	assert.Equal(t, "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v", paramByName["account_destination_mint"].Constraint.GetFixedValue())

	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_event_authority"].Constraint.Type)
	assert.Equal(t, "D8cy77BBepLMngZx6ZukaTff5hCt1HrWyKk3Hnd9oitf", paramByName["account_event_authority"].Constraint.GetFixedValue())

	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_program"].Constraint.Type)
	assert.Equal(t, "JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4", paramByName["account_program"].Constraint.GetFixedValue())

	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["arg_in_amount"].Constraint.Type)
	assert.Equal(t, "50000000", paramByName["arg_in_amount"].Constraint.GetFixedValue())

	// Verify ANY constraints exist for dynamic fields (userTransferAuthority can be any signer)
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["account_user_transfer_authority"].Constraint.Type)
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["account_destination_token_account"].Constraint.Type)
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_ANY, paramByName["account_platform_fee_account"].Constraint.Type)

	// Verify FIXED constraints for user ATAs
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_user_source_token_account"].Constraint.Type)
	expectedSourceATA, err := DeriveATA("4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM", "So11111111111111111111111111111111111111112")
	require.NoError(t, err)
	assert.Equal(t, expectedSourceATA, paramByName["account_user_source_token_account"].Constraint.GetFixedValue())

	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["account_user_destination_token_account"].Constraint.Type)
	expectedDestATA, err := DeriveATA("4w3VdMehnFqFTNEg9jZtKS76n4pNcVjaDZK9TQtw9jKM", "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")
	require.NoError(t, err)
	assert.Equal(t, expectedDestATA, paramByName["account_user_destination_token_account"].Constraint.GetFixedValue())
}

const testTHORChainAddress = "thor1abc123def456ghi789"

func TestTryFormat_THORChainSend(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "thorchain.send",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testTHORChainAddress,
					},
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, "thorchain.rune.transfer", result[0].Resource)
	assert.Equal(t, types.TargetType_TARGET_TYPE_UNSPECIFIED, result[0].Target.TargetType)
	require.Len(t, result[0].ParameterConstraints, 2)

	paramByName := make(map[string]*types.ParameterConstraint)
	for _, param := range result[0].ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	assert.Contains(t, paramByName, "recipient")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["recipient"].Constraint.Type)
	assert.Equal(t, testTHORChainAddress, paramByName["recipient"].Constraint.GetFixedValue())

	assert.Contains(t, paramByName, "amount")
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["amount"].Constraint.Type)
	assert.Equal(t, "1000000", paramByName["amount"].Constraint.GetFixedValue())
}

func TestTryFormat_THORChainSwap(t *testing.T) {
	const (
		fromAddress      = testTHORChainAddress
		fromAmount       = "2000000" // 2 RUNE
		toChain          = "ethereum"
		toAsset          = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" // USDC
		toAddress        = "0x742d35Cc6634C0532925a3b8D5c9E0B0Cf8a6b"
		expectedResource = "thorchain.thorchain_swap.swap"
	)

	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "thorchain.swap",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "from_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "", // Empty for native RUNE
					},
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAddress,
					},
				},
			},
			{
				ParameterName: "from_amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: fromAmount,
					},
				},
			},
			{
				ParameterName: "to_chain",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toChain,
					},
				},
			},
			{
				ParameterName: "to_asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toAsset,
					},
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: toAddress,
					},
				},
			},
		},
	}

	result, err := metaRule.TryFormat(rule)
	require.NoError(t, err)
	require.Len(t, result, 1)

	assert.Equal(t, expectedResource, result[0].Resource)
	assert.Equal(t, types.TargetType_TARGET_TYPE_UNSPECIFIED, result[0].Target.TargetType)
	require.Len(t, result[0].ParameterConstraints, 3)

	paramByName := make(map[string]*types.ParameterConstraint)
	for _, param := range result[0].ParameterConstraints {
		paramByName[param.ParameterName] = param
	}

	// Verify all required parameters are present
	assert.Contains(t, paramByName, "amount")
	assert.Contains(t, paramByName, "from_asset")
	assert.Contains(t, paramByName, "memo")

	// Verify constraint values
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["amount"].Constraint.Type)
	assert.Equal(t, fromAmount, paramByName["amount"].Constraint.GetFixedValue())

	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_FIXED, paramByName["from_asset"].Constraint.Type)
	assert.Equal(t, "RUNE", paramByName["from_asset"].Constraint.GetFixedValue()) // Defaults to RUNE for empty from_asset

	// Verify memo constraint (regexp for THORChain swap format)
	assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_REGEXP, paramByName["memo"].Constraint.Type)
	regexpValue := paramByName["memo"].Constraint.GetRegexpValue()
	expectedPattern := fmt.Sprintf("^=:ETH\\.USDC:%s:.*", regexp.QuoteMeta(toAddress))
	assert.Equal(t, expectedPattern, regexpValue)
}

func TestTryFormat_THORChainSend_MissingRecipient(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "thorchain.send",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find constraint: to_address")
}

func TestTryFormat_THORChainSwap_MissingFromAsset(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "thorchain.swap",
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testTHORChainAddress,
					},
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to find constraint: from_asset")
}

func TestTryFormat_THORChain_UnsupportedProtocol(t *testing.T) {
	metaRule := NewMetaRule()

	rule := &types.Rule{
		Resource: "thorchain.stake", // Unsupported protocol
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "asset",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "from_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
				},
			},
			{
				ParameterName: "to_address",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				},
			},
			{
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				},
			},
		},
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol id for THORChain: stake")
}
