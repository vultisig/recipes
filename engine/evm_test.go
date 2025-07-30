package engine

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
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

func TestERC20Transfer(t *testing.T) {
	e := newEvm()
	erc20Instance := erc20.NewErc20()

	// Test cases
	testCases := []struct {
		name        string
		rule        *types.Rule
		recipient   common.Address
		amount      *big.Int
		shouldError bool
	}{
		{
			name: "Fixed constraint - matching recipient and amount",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "recipient",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value:    &types.Constraint_FixedValue{FixedValue: "0x1111111111111111111111111111111111111111"},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value:    &types.Constraint_FixedValue{FixedValue: "1000000000000000000"},
							Required: true,
						},
					},
				},
			},
			recipient:   common.HexToAddress("0x1111111111111111111111111111111111111111"),
			amount:      big.NewInt(1000000000000000000), // 1 token
			shouldError: false,
		},
		{
			name: "Min constraint - amount above minimum",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "recipient",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value:    &types.Constraint_FixedValue{FixedValue: "0x1111111111111111111111111111111111111111"},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_MIN,
							Value:    &types.Constraint_MinValue{MinValue: "500000000000000000"},
							Required: true,
						},
					},
				},
			},
			recipient:   common.HexToAddress("0x1111111111111111111111111111111111111111"),
			amount:      big.NewInt(1000000000000000000), // 1 token, above 0.5 minimum
			shouldError: false,
		},
		{
			name: "Max constraint - amount below maximum",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "recipient",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value:    &types.Constraint_FixedValue{FixedValue: "0x1111111111111111111111111111111111111111"},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_MAX,
							Value:    &types.Constraint_MaxValue{MaxValue: "2000000000000000000"},
							Required: true,
						},
					},
				},
			},
			recipient:   common.HexToAddress("0x1111111111111111111111111111111111111111"),
			amount:      big.NewInt(1000000000000000000), // 1 token, below 2 maximum
			shouldError: false,
		},
		{
			name: "Max constraint - amount above maximum",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "recipient",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value:    &types.Constraint_FixedValue{FixedValue: "0x1111111111111111111111111111111111111111"},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_MAX,
							Value:    &types.Constraint_MaxValue{MaxValue: "1000000000000000000"},
							Required: true,
						},
					},
				},
			},
			recipient:   common.HexToAddress("0x1111111111111111111111111111111111111111"),
			amount:      big.NewInt(2000000000000000000), // 2 token, below 1 maximum
			shouldError: true,
		},
		{
			name: "Min constraint - amount below minimum",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "ethereum.erc20.transfer",
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "recipient",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value:    &types.Constraint_FixedValue{FixedValue: "0x1111111111111111111111111111111111111111"},
							Required: true,
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type:     types.ConstraintType_CONSTRAINT_TYPE_MIN,
							Value:    &types.Constraint_MaxValue{MaxValue: "2000000000000000000"},
							Required: true,
						},
					},
				},
			},
			recipient:   common.HexToAddress("0x1111111111111111111111111111111111111111"),
			amount:      big.NewInt(1000000000000000000), // 1 token, below 2 minimum
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Pack the transfer function call
			data := erc20Instance.PackTransfer(tc.recipient, tc.amount)

			// Create a transaction with the packed data
			tokenAddress := common.HexToAddress("0x2222222222222222222222222222222222222222")
			txBytes := buildUnsignedTx(tokenAddress, data, big.NewInt(0))

			// Evaluate the transaction against the rule
			err := e.evaluate(tc.rule, txBytes)

			if tc.shouldError && err == nil {
				t.Errorf("Expected error but got none")
			}

			if !tc.shouldError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
