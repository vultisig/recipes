package evm

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/vultisig/recipes/sdk/evm/codegen/erc20"
	"github.com/vultisig/recipes/sdk/evm/codegen/routerv6_1inch"
	"github.com/vultisig/recipes/sdk/evm/codegen/uniswapv2_router"
	"github.com/vultisig/recipes/types"
	vgcommon "github.com/vultisig/vultisig-go/common"
)

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

func TestEvaluate_ERC20Transfer(t *testing.T) {
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
							Value: &types.Constraint_MinValue{
								MinValue: "2000000000000000000",
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
			name: "Any constraint: any recipient and amount",
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
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(usdc),
			recipient:   common.HexToAddress(dumb2),
			amount:      big.NewInt(3000000000000000000),
			shouldError: false,
		},
		{
			name: "Mixed constraints: any recipient with fixed amount",
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
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
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
			recipient:   common.HexToAddress(dumb2),
			amount:      big.NewInt(1000000000000000000),
			shouldError: false,
		},
		{
			name: "Mixed constraints: any recipient with incorrect fixed amount",
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
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
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
			recipient:   common.HexToAddress(dumb2),
			amount:      big.NewInt(2000000000000000000),
			shouldError: true,
		},
		{
			name: "Mixed constraints: fixed recipient with any amount",
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
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(usdc),
			recipient:   common.HexToAddress(dumb1),      // Must match fixed value
			amount:      big.NewInt(5000000000000000000), // Any amount should be accepted
			shouldError: false,
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

			native, _ := vgcommon.Ethereum.NativeSymbol()
			evm, err := NewEvm(native)
			if err != nil {
				t.Fatalf("Failed to create EVM: %v", err)
			}
			err = evm.Evaluate(tc.rule, txBytes)

			if tc.shouldError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestEvaluate_NativeTransfer(t *testing.T) {
	const (
		magicConstTreasury = "0x8E247a480449c84a5fDD25974A8501f3EFa4ABb9"
		dumb1              = "0x1111111111111111111111111111111111111111"
		dumb2              = "0x2222222222222222222222222222222222222222"
	)

	native, err := vgcommon.Ethereum.NativeSymbol()
	if err != nil {
		t.Fatalf("Failed to get native symbol: %v", err)
	}

	testCases := []struct {
		name        string
		rule        *types.Rule
		to          common.Address
		value       *big.Int
		shouldError bool
	}{
		{
			name: "Fixed constraint: amount equals fixed value",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("ethereum.%s.transfer", strings.ToLower(native)),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: dumb1,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
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
			to:          common.HexToAddress(dumb1),
			value:       big.NewInt(1000000000000000000), // 1 ETH
			shouldError: false,
		},
		{
			name: "Fixed constraint: amount not equal to fixed value",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("ethereum.%s.transfer", strings.ToLower(native)),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: dumb1,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
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
			to:          common.HexToAddress(dumb1),
			value:       big.NewInt(2000000000000000000), // 2 ETH
			shouldError: true,
		},
		{
			name: "Magic constant target",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("ethereum.%s.transfer", strings.ToLower(native)),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_MAGIC_CONSTANT,
					Target: &types.Target_MagicConstant{
						MagicConstant: types.MagicConstant_VULTISIG_TREASURY,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
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
			value:       big.NewInt(1000000000000000000), // 1 ETH
			shouldError: false,
		},
		{
			name: "Magic constant constraint (should fail - magic constants resolve to addresses, not numeric values)",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("ethereum.%s.transfer", strings.ToLower(native)),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: dumb1,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT,
							Value: &types.Constraint_MagicConstantValue{
								MagicConstantValue: types.MagicConstant_VULTISIG_TREASURY,
							},
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(dumb1),
			value:       big.NewInt(1000000000000000000), // 1 ETH
			shouldError: true,
		},
		{
			name: "Wrong target address",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("ethereum.%s.transfer", strings.ToLower(native)),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: dumb1,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
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
			to:          common.HexToAddress(dumb2),
			value:       big.NewInt(1000000000000000000), // 1 ETH
			shouldError: true,
		},
		{
			name: "Invalid function ID",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("ethereum.%s.invalid", strings.ToLower(native)),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: dumb1,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
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
			to:          common.HexToAddress(dumb1),
			value:       big.NewInt(1000000000000000000), // 1 ETH
			shouldError: true,
		},
		{
			name: "Wrong number of parameter constraints",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("ethereum.%s.transfer", strings.ToLower(native)),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: dumb1,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
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
					{
						ParameterName: "extra",
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
			to:          common.HexToAddress(dumb1),
			value:       big.NewInt(1000000000000000000), // 1 ETH
			shouldError: true,
		},
		{
			name: "Missing amount parameter constraint",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("ethereum.%s.transfer", strings.ToLower(native)),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: dumb1,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "wrong_name",
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
			to:          common.HexToAddress(dumb1),
			value:       big.NewInt(1000000000000000000), // 1 ETH
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		label := "[positive]"
		if tc.shouldError {
			label = "[negative]"
		}
		t.Run(fmt.Sprintf("%s %s", label, tc.name), func(t *testing.T) {
			txBytes := buildUnsignedTx(tc.to, []byte{}, tc.value)

			evm, err := NewEvm(native)
			if err != nil {
				t.Fatalf("Failed to create EVM: %v", err)
			}
			er := evm.Evaluate(tc.rule, txBytes)

			if tc.shouldError && er == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.shouldError && er != nil {
				t.Errorf("Expected no error but got: %v", er)
			}
		})
	}
}

func TestEvaluate_UniswapAddressArray(t *testing.T) {
	const (
		uniswapRouter = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D" // Uniswap V2 Router
		weth          = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" // WETH
		usdc          = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" // USDC
		dai           = "0x6B175474E89094C44Da98b954EedeAC495271d0F" // DAI
		dumb1         = "0x1111111111111111111111111111111111111111"
		dumb2         = "0x2222222222222222222222222222222222222222"
	)

	testCases := []struct {
		name         string
		rule         *types.Rule
		to           common.Address
		amountIn     *big.Int
		amountOutMin *big.Int
		path         []common.Address
		recipient    common.Address
		deadline     *big.Int
		shouldError  bool
	}{
		{
			name: "Fixed constraint: path with fixed addresses",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.uniswapV2_router.swapExactTokensForTokens",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: uniswapRouter,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "amountIn",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "1000000000000000000",
							},
							Required: true,
						},
					},
					{
						ParameterName: "amountOutMin",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "500000000000000000",
							},
							Required: true,
						},
					},
					{
						ParameterName: "path",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: weth + "," + usdc,
							},
							Required: true,
						},
					},
					{
						ParameterName: "to",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: dumb1,
							},
							Required: true,
						},
					},
					{
						ParameterName: "deadline",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "1700000000",
							},
							Required: true,
						},
					},
				},
			},
			to:           common.HexToAddress(uniswapRouter),
			amountIn:     big.NewInt(1000000000000000000), // 1 ETH
			amountOutMin: big.NewInt(500000000000000000),  // 0.5 ETH
			path:         []common.Address{common.HexToAddress(weth), common.HexToAddress(usdc)},
			recipient:    common.HexToAddress(dumb1),
			deadline:     big.NewInt(1700000000),
			shouldError:  false,
		},
	}

	for _, tc := range testCases {
		label := "[positive]"
		if tc.shouldError {
			label = "[negative]"
		}
		t.Run(fmt.Sprintf("%s %s", label, tc.name), func(t *testing.T) {
			data := uniswapv2_router.NewUniswapv2Router().PackSwapExactTokensForTokens(
				tc.amountIn,
				tc.amountOutMin,
				tc.path,
				tc.recipient,
				tc.deadline,
			)
			txBytes := buildUnsignedTx(tc.to, data, big.NewInt(0))

			native, _ := vgcommon.Ethereum.NativeSymbol()
			evm, err := NewEvm(native)
			if err != nil {
				t.Fatalf("Failed to create EVM: %v", err)
			}
			err = evm.Evaluate(tc.rule, txBytes)

			if tc.shouldError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestEvaluate_UniswapUint8AndBytes32(t *testing.T) {
	const (
		uniswapRouter = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D" // Uniswap V2 Router
		token         = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" // USDC
		dumb1         = "0x1111111111111111111111111111111111111111"
	)

	// Create sample bytes32 values
	rBytes := [32]byte{}
	sBytes := [32]byte{}
	for i := 0; i < 32; i++ {
		rBytes[i] = byte(i)
		sBytes[i] = byte(32 - i)
	}

	// Base64 encode the bytes32 values
	rBase64 := base64.StdEncoding.EncodeToString(rBytes[:])
	sBase64 := base64.StdEncoding.EncodeToString(sBytes[:])

	testCases := []struct {
		name           string
		rule           *types.Rule
		to             common.Address
		token          common.Address
		liquidity      *big.Int
		amountTokenMin *big.Int
		amountETHMin   *big.Int
		recipient      common.Address
		deadline       *big.Int
		approveMax     bool
		v              uint8
		r              [32]byte
		s              [32]byte
		shouldError    bool
	}{
		{
			name: "Fixed constraint: uint8 and bytes32 parameters",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.uniswapV2_router.removeLiquidityETHWithPermit",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: uniswapRouter,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "token",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: token,
							},
							Required: true,
						},
					},
					{
						ParameterName: "liquidity",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "1000000000000000000",
							},
							Required: true,
						},
					},
					{
						ParameterName: "amountTokenMin",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "500000000000000000",
							},
							Required: true,
						},
					},
					{
						ParameterName: "amountETHMin",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "500000000000000000",
							},
							Required: true,
						},
					},
					{
						ParameterName: "to",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: dumb1,
							},
							Required: true,
						},
					},
					{
						ParameterName: "deadline",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "1700000000",
							},
							Required: true,
						},
					},
					{
						ParameterName: "approveMax",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "true",
							},
							Required: true,
						},
					},
					{
						ParameterName: "v",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "27",
							},
							Required: true,
						},
					},
					{
						ParameterName: "r",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: rBase64,
							},
							Required: true,
						},
					},
					{
						ParameterName: "s",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: sBase64,
							},
							Required: true,
						},
					},
				},
			},
			to:             common.HexToAddress(uniswapRouter),
			token:          common.HexToAddress(token),
			liquidity:      big.NewInt(1000000000000000000), // 1 ETH
			amountTokenMin: big.NewInt(500000000000000000),  // 0.5 ETH
			amountETHMin:   big.NewInt(500000000000000000),  // 0.5 ETH
			recipient:      common.HexToAddress(dumb1),
			deadline:       big.NewInt(1700000000),
			approveMax:     true,
			v:              27,
			r:              rBytes,
			s:              sBytes,
			shouldError:    false,
		},
	}

	for _, tc := range testCases {
		label := "[positive]"
		if tc.shouldError {
			label = "[negative]"
		}
		t.Run(fmt.Sprintf("%s %s", label, tc.name), func(t *testing.T) {
			data := uniswapv2_router.NewUniswapv2Router().PackRemoveLiquidityETHWithPermit(
				tc.token,
				tc.liquidity,
				tc.amountTokenMin,
				tc.amountETHMin,
				tc.recipient,
				tc.deadline,
				tc.approveMax,
				tc.v,
				tc.r,
				tc.s,
			)
			txBytes := buildUnsignedTx(tc.to, data, big.NewInt(0))

			native, _ := vgcommon.Ethereum.NativeSymbol()
			evm, err := NewEvm(native)
			if err != nil {
				t.Fatalf("Failed to create EVM: %v", err)
			}
			err = evm.Evaluate(tc.rule, txBytes)

			if tc.shouldError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestEvaluate_1inchSwap(t *testing.T) {
	const (
		router1inch   = "0x111111125421cA6dc452d289314280a0f8842A65" // Uniswap V2 Router
		weth          = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" // WETH
		uniswapRouter = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D" // Uniswap V2 Router
		token         = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" // USDC
		dumb1         = "0x1111111111111111111111111111111111111111"
	)

	// Create sample bytes32 values
	testBytes := make([]byte, 4)
	testBytes2 := make([]byte, 4)
	for i := 0; i < 4; i++ {
		testBytes[i] = byte(i)
		testBytes2[i] = byte(4 - i)
	}

	// Base64 encode the bytes32 values
	testBase64 := base64.StdEncoding.EncodeToString(testBytes)

	testCases := []struct {
		name        string
		rule        *types.Rule
		to          common.Address
		swap        *big.Int
		executor    common.Address
		desc        routerv6_1inch.GenericRouterSwapDescription
		data        []byte
		shouldError bool
	}{
		{
			name: "valid tuple in constraint",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.routerV6_1inch.swap",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: router1inch,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "executor",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Value: &types.Constraint_FixedValue{
								FixedValue: "500000000000000000",
							},
							Required: true,
						},
					},
					{
						ParameterName: "desc.srcToken",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.dstToken",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.srcReceiver",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.dstReceiver",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.minReturnAmount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.flags",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "data",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: testBase64,
							},
							Required: true,
						},
					},
				},
			},
			to:       common.HexToAddress(router1inch),
			swap:     big.NewInt(1000000000000000000), // 1 ETH
			executor: common.HexToAddress(uniswapRouter),
			desc: routerv6_1inch.GenericRouterSwapDescription{
				SrcToken:        common.HexToAddress(token),
				DstToken:        common.HexToAddress(weth),
				SrcReceiver:     common.HexToAddress(dumb1),
				DstReceiver:     common.HexToAddress(dumb1),
				Amount:          big.NewInt(1000000000000000000),
				MinReturnAmount: big.NewInt(1),
				Flags:           new(big.Int),
			},
			data:        testBytes,
			shouldError: false,
		},
		{
			name: "incomplete tuple in constraints",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.routerV6_1inch.swap",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: router1inch,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "executor",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Value: &types.Constraint_FixedValue{
								FixedValue: "500000000000000000",
							},
							Required: true,
						},
					},
					{
						ParameterName: "desc.srcToken",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "data",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
				},
			},
			to:       common.HexToAddress(router1inch),
			swap:     big.NewInt(1000000000000000000), // 1 ETH
			executor: common.HexToAddress(uniswapRouter),
			desc: routerv6_1inch.GenericRouterSwapDescription{
				SrcToken:        common.HexToAddress(token),
				DstToken:        common.HexToAddress(weth),
				SrcReceiver:     common.HexToAddress(dumb1),
				DstReceiver:     common.HexToAddress(dumb1),
				Amount:          big.NewInt(1000000000000000000),
				MinReturnAmount: big.NewInt(1),
				Flags:           new(big.Int),
			},
			data:        []byte{},
			shouldError: true,
		},
		{
			name: "failed dstToken check",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.routerV6_1inch.swap",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: router1inch,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "executor",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Value: &types.Constraint_FixedValue{
								FixedValue: "500000000000000000",
							},
							Required: true,
						},
					},
					{
						ParameterName: "desc.srcToken",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.dstToken",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: weth,
							},
							Required: true,
						},
					},
					{
						ParameterName: "desc.srcReceiver",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.dstReceiver",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.minReturnAmount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.flags",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "data",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: testBase64,
							},
							Required: true,
						},
					},
				},
			},
			to:       common.HexToAddress(router1inch),
			swap:     big.NewInt(1000000000000000000), // 1 ETH
			executor: common.HexToAddress(uniswapRouter),
			desc: routerv6_1inch.GenericRouterSwapDescription{
				SrcToken:        common.HexToAddress(weth),
				DstToken:        common.HexToAddress(token),
				SrcReceiver:     common.HexToAddress(dumb1),
				DstReceiver:     common.HexToAddress(dumb1),
				Amount:          big.NewInt(1000000000000000000),
				MinReturnAmount: big.NewInt(1),
				Flags:           new(big.Int),
			},
			data:        testBytes,
			shouldError: true,
		},
		{
			name: "failed data(bytes) check",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.routerV6_1inch.swap",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: router1inch,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "executor",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.srcToken",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.dstToken",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: weth,
							},
							Required: true,
						},
					},
					{
						ParameterName: "desc.srcReceiver",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.dstReceiver",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.minReturnAmount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "desc.flags",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "data",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: testBase64,
							},
							Required: true,
						},
					},
				},
			},
			to:       common.HexToAddress(router1inch),
			swap:     big.NewInt(1000000000000000000), // 1 ETH
			executor: common.HexToAddress(uniswapRouter),
			desc: routerv6_1inch.GenericRouterSwapDescription{
				SrcToken:        common.HexToAddress(token),
				DstToken:        common.HexToAddress(weth),
				SrcReceiver:     common.HexToAddress(dumb1),
				DstReceiver:     common.HexToAddress(dumb1),
				Amount:          big.NewInt(1000000000000000000),
				MinReturnAmount: big.NewInt(1),
				Flags:           new(big.Int),
			},
			data:        testBytes2,
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.name), func(t *testing.T) {
			data := routerv6_1inch.NewRouterv61inch().PackSwap(
				tc.executor,
				tc.desc,
				tc.data,
			)
			txBytes := buildUnsignedTx(tc.to, data, big.NewInt(0))

			native, _ := vgcommon.Ethereum.NativeSymbol()
			evm, err := NewEvm(native)
			if err != nil {
				t.Fatalf("Failed to create EVM: %v", err)
			}
			err = evm.Evaluate(tc.rule, txBytes)

			if tc.shouldError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestEvaluate_ErrorCases(t *testing.T) {
	const (
		dumb1 = "0x1111111111111111111111111111111111111111"
	)

	native, err := vgcommon.Ethereum.NativeSymbol()
	if err != nil {
		t.Fatalf("Failed to get native symbol: %v", err)
	}

	testCases := []struct {
		name        string
		rule        *types.Rule
		txBytes     []byte
		shouldError bool
	}{
		{
			name: "Invalid resource format",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "invalid_resource",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: dumb1,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
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
			txBytes:     buildUnsignedTx(common.HexToAddress(dumb1), []byte{}, big.NewInt(1000000000000000000)),
			shouldError: true,
		},
		{
			name: "Invalid effect (not ALLOW)",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_DENY,
				Resource: fmt.Sprintf("ethereum.%s.transfer", strings.ToLower(native)),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: dumb1,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
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
			txBytes:     buildUnsignedTx(common.HexToAddress(dumb1), []byte{}, big.NewInt(1000000000000000000)),
			shouldError: true,
		},
		{
			name: "Invalid transaction payload",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: fmt.Sprintf("ethereum.%s.transfer", strings.ToLower(native)),
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: dumb1,
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
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
			txBytes:     []byte{0x01, 0x02, 0x03}, // Invalid transaction payload
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			evm, err := NewEvm(native)
			if err != nil {
				t.Fatalf("Failed to create EVM: %v", err)
			}
			er := evm.Evaluate(tc.rule, tc.txBytes)

			if tc.shouldError && er == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.shouldError && er != nil {
				t.Errorf("Expected no error but got: %v", er)
			}
		})
	}
}

func TestEvaluate_MethodDescriptorValidation(t *testing.T) {
	const (
		usdc  = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
		dumb1 = "0x1111111111111111111111111111111111111111"
	)

	testCases := []struct {
		name          string
		rule          *types.Rule
		to            common.Address
		recipient     common.Address
		amount        *big.Int
		customData    []byte
		shouldError   bool
		expectedError string
	}{
		{
			name: "Valid method descriptor",
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
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(usdc),
			recipient:   common.HexToAddress(dumb1),
			amount:      big.NewInt(1000000000000000000),
			shouldError: false,
		},
		{
			name: "Invalid method descriptor - wrong first 4 bytes",
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
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
				},
			},
			to:            common.HexToAddress(usdc),
			recipient:     common.HexToAddress(dumb1),
			amount:        big.NewInt(1000000000000000000),
			customData:    append([]byte{0x12, 0x34, 0x56, 0x78}, make([]byte, 64)...),
			shouldError:   true,
			expectedError: "method descriptor mismatch",
		},
		{
			name: "Calldata too short",
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
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Required: true,
						},
					},
				},
			},
			to:            common.HexToAddress(usdc),
			recipient:     common.HexToAddress(dumb1),
			amount:        big.NewInt(1000000000000000000),
			customData:    []byte{0x12, 0x34},
			shouldError:   true,
			expectedError: "calldata too short",
		},
	}

	for _, tc := range testCases {
		label := "[positive]"
		if tc.shouldError {
			label = "[negative]"
		}
		t.Run(fmt.Sprintf("%s %s", label, tc.name), func(t *testing.T) {
			var data []byte
			if tc.customData != nil {
				data = tc.customData
			} else {
				data = erc20.NewErc20().PackTransfer(tc.recipient, tc.amount)
			}
			txBytes := buildUnsignedTx(tc.to, data, big.NewInt(0))

			native, _ := vgcommon.Ethereum.NativeSymbol()
			evm, err := NewEvm(native)
			if err != nil {
				t.Fatalf("Failed to create EVM: %v", err)
			}
			err = evm.Evaluate(tc.rule, txBytes)

			if tc.shouldError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
			if tc.shouldError && err != nil && tc.expectedError != "" {
				if !strings.Contains(err.Error(), tc.expectedError) {
					t.Errorf("Expected error containing '%s', but got: %v", tc.expectedError, err)
				}
			}
		})
	}
}

func TestEvaluate_RegexpConstraints(t *testing.T) {
	const (
		usdc  = "0xA0b86a33E6842C040F93B1DE7dd5aEf8E71bcE64"
		dumb1 = "0x742d35Cc6634C0532925a3b8D9C9d9f8C4669bE5"
		dumb2 = "0x843d35Cc6634C0532925a3b8D9C9d9f8C4669bE5"
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
			name: "Regexp constraint: address pattern match",
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
							Type: types.ConstraintType_CONSTRAINT_TYPE_REGEXP,
							Value: &types.Constraint_RegexpValue{
								RegexpValue: "(?i)0x742d.*", // Case-insensitive address pattern
							},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
							Value: &types.Constraint_FixedValue{
								FixedValue: "1000000000000000000",
							},
							Required: true,
						},
					},
				},
			},
			to:          common.HexToAddress(usdc),
			recipient:   common.HexToAddress(dumb1), // Matches (?i)0x742d.*
			amount:      big.NewInt(1000000000000000000),
			shouldError: false,
		},
		{
			name: "Regexp constraint: address pattern mismatch",
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
							Type: types.ConstraintType_CONSTRAINT_TYPE_REGEXP,
							Value: &types.Constraint_RegexpValue{
								RegexpValue: "0x843d.*", // Different address pattern
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
			recipient:   common.HexToAddress(dumb1), // Doesnt match 0x843d.*
			amount:      big.NewInt(1000000000000000000),
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

			evm, err := NewEvm("ethereum")
			if err != nil {
				t.Fatalf("Failed to create EVM: %v", err)
			}
			err = evm.Evaluate(tc.rule, txBytes)

			if tc.shouldError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
