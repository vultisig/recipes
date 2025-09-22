package metarule

import (
	"fmt"
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
				ParameterName: "amount",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: "1000000",
					},
				},
			},
			{
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
					Value: &types.Constraint_FixedValue{
						FixedValue: testAddress,
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
		Resource: "bitcoin.send", // Unsupported chain
	}

	_, err := metaRule.TryFormat(rule)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "got meta format (bitcoin.send) but chain not supported: Bitcoin")
}

func TestTryFormat_SolanaSOLTransfer(t *testing.T) {
	metaRule := NewMetaRule()

	// Test native SOL transfer (target is system program)
	rule := &types.Rule{}
	rule.Resource = "solana.send" // Meta-rule format with empty function ID
	rule.Target = &types.Target{
		TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
		Target: &types.Target_Address{
			Address: solana.SystemProgramID.String(), // Native SOL transfer
		},
	}
	rule.ParameterConstraints = []*types.ParameterConstraint{
		{
			ParameterName: "recipient",
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
	assert.Equal(t, testAddress, result[0].Target.GetAddress())
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

	// Test SPL token transfer (target is not system program)
	const tokenMintAddress = "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" // USDC mint
	rule := &types.Rule{
		Resource: "solana.send",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: tokenMintAddress, // SPL token transfer
			},
		},
		ParameterConstraints: []*types.ParameterConstraint{
			{
				ParameterName: "recipient",
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
	require.Len(t, result, 1)
	assert.Equal(t, "solana.spl_token.transfer", result[0].Resource)
	assert.Equal(t, tokenMintAddress, result[0].Target.GetAddress())
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
			// Missing recipient constraint
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
	assert.Contains(t, err.Error(), "failed to parse `recipient`")
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
				ParameterName: "recipient",
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
	assert.Contains(t, err.Error(), "failed to parse `amount`")
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
				ParameterName: "recipient",
				Constraint: &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_ANY, // Invalid for recipient
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
	assert.Contains(t, err.Error(), "invalid constraint type for `recipient`")
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
				ParameterName: "recipient",
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
	expectedResource := fmt.Sprintf("%s.%s.transfer", testEVMChain, nativeSymbol)
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
				ParameterName: "recipient",
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
				ParameterName: "recipient",
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
	assert.Contains(t, err.Error(), "failed to parse `amount`")
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
	assert.Contains(t, err.Error(), "failed to parse `recipient`")
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
				ParameterName: "recipient",
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
				ParameterName: "recipient",
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
	assert.Contains(t, err.Error(), "invalid constraint type for `recipient`")
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
				ParameterName: "recipient",
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
	expectedResource := fmt.Sprintf("%s.%s.transfer", testEVMChain, nativeSymbol)
	assert.Equal(t, expectedResource, result[0].Resource)
	assert.Equal(t, types.MagicConstant_VULTISIG_TREASURY, result[0].Target.GetMagicConstant())
}
