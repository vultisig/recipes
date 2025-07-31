package evm

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	rcommon "github.com/vultisig/recipes/common"
	"github.com/vultisig/recipes/sdk/evm/codegen/erc20"
	"github.com/vultisig/recipes/types"
)

// Helper function to build an unsigned transaction
func buildUnsignedTx(to common.Address, data []byte, value *big.Int) []byte {
	unsigned := struct {
		ChainID    *big.Int
		Nonce      uint64
		GasTipCap  *big.Int
		GasFeeCap  *big.Int
		Gas        uint64
		To         *common.Address `rlp:"nil"`
		Value      *big.Int
		Data       []byte
		AccessList etypes.AccessList
	}{
		ChainID:    big.NewInt(1),
		Nonce:      0,
		GasTipCap:  big.NewInt(2_000_000_000),  // 2 gwei
		GasFeeCap:  big.NewInt(20_000_000_000), // 20 gwei
		Gas:        300_000,
		To:         &to,
		Value:      value,
		Data:       data,
		AccessList: nil,
	}
	payload, err := rlp.EncodeToBytes(unsigned)
	if err != nil {
		panic(err)
	}
	return append([]byte{etypes.DynamicFeeTxType}, payload...)
}

func TestERC20Transfer_assertArg(t *testing.T) {
	const (
		magicConstTreasury = "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9"
		usdc               = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
		dumb1              = "0x1111111111111111111111111111111111111111"
		dumb2              = "0x2222222222222222222222222222222222222222"
	)

	testCases := []struct {
		name        string
		rule        *types.Rule
		to          common.Address
		recipient   common.Address
		amount      *big.Int
		shouldError bool
	}{
		{
			name: "Fixed constraint: magic const and fixed amount",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
					Target: &types.Target_MagicConstant{
						MagicConstant: types.MagicConstant_VULTISIG_TREASURY,
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
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "1000000000000000000",
							},
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(magicConstTreasury),
			recipient:   common.HexToAddress(magicConstTreasury),
			amount:      big.NewInt(1000000000000000000), // 1
			shouldError: false,
		},
		{
			name: "Fixed constraint: fixed recipient and amount",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: usdc,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "recipient",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: dumb1,
							},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "1000000000000000000",
							},
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(usdc),
			recipient:   common.HexToAddress(dumb1),
			amount:      big.NewInt(1000000000000000000), // 1
			shouldError: false,
		},
		{
			name: "Min constraint: amount above minimum",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: usdc,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "recipient",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: dumb1,
							},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
							Value: &types.Constraint_MinValue{
								MinValue: "500000000000000000",
							},
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(usdc),
			recipient:   common.HexToAddress(dumb1),
			amount:      big.NewInt(1000000000000000000), // 1 token > 0.5 minimum
			shouldError: false,
		},
		{
			name: "Max constraint: amount below maximum",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: usdc,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "recipient",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: dumb1,
							},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
							Value: &types.Constraint_MaxValue{
								MaxValue: "2000000000000000000",
							},
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(usdc),
			recipient:   common.HexToAddress(dumb1),
			amount:      big.NewInt(1000000000000000000), // 1 token < 2 maximum
			shouldError: false,
		},
		{
			name: "Max constraint: amount above maximum",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: usdc,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "recipient",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: dumb1,
							},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
							Value: &types.Constraint_MaxValue{
								MaxValue: "1000000000000000000",
							},
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(usdc),
			recipient:   common.HexToAddress(dumb1),
			amount:      big.NewInt(2000000000000000000), // 2 token > 1 maximum
			shouldError: true,
		},
		{
			name: "Min constraint: amount below minimum",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: usdc,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "recipient",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: dumb1,
							},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
							Value: &types.Constraint_MaxValue{
								MaxValue: "2000000000000000000",
							},
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(usdc),
			recipient:   common.HexToAddress(dumb1),
			amount:      big.NewInt(1000000000000000000), // 1 token < 2 minimum
			shouldError: true,
		},
		{
			name: "Target: wrong address",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: usdc,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "recipient",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: dumb1,
							},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_MaxValue{
								MaxValue: "2000000000000000000",
							},
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(dumb2),
			recipient:   common.HexToAddress(dumb1),
			amount:      big.NewInt(2000000000000000000),
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		label := "[positive]"
		if tc.shouldError {
			label = "[negative]"
		}
		t.Run(fmt.Sprintf("%s %s", label, tc.name), func(t *testing.T) {
			data := erc20.NewErc20().PackTransfer(tc.recipient, tc.amount)
			txBytes := buildUnsignedTx(tc.to, data, big.NewInt(0))

			native, _ := rcommon.Ethereum.NativeSymbol()
			err := NewEvm(native).Evaluate(tc.rule, txBytes)

			if tc.shouldError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
