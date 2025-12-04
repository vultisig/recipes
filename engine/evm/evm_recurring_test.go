package evm

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vultisig/recipes/metarule"
	"github.com/vultisig/recipes/sdk/evm/codegen/erc20"
	"github.com/vultisig/recipes/sdk/evm/codegen/routerv6_1inch"
	"github.com/vultisig/recipes/types"
	vgcommon "github.com/vultisig/vultisig-go/common"
)

// Test constants
const (
	testFromAddress      = "0xcB9B049B9c937acFDB87EeCfAa9e7f2c51E754f5"
	testRecipientAddress = "0x1234567890abcdef1234567890abcdef12345678"
	testUSDCAddress      = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	testWETHAddress      = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	test1inchRouter      = "0x111111125421ca6dc452d289314280a0f8842a65"
)

// buildEVMTestTx creates an unsigned EVM transaction for testing
func buildEVMTestTx(chainID int64, to common.Address, data []byte, value *big.Int) []byte {
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
		ChainID:    big.NewInt(chainID),
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

// TestRecurringSend_Native tests end-to-end native ETH recurring send flow
func TestRecurringSend_Native(t *testing.T) {
	testCases := []struct {
		name        string
		chain       vgcommon.Chain
		chainID     int64
		toAddress   string
		amount      *big.Int
		shouldPass  bool
	}{
		{
			name:       "Ethereum native transfer - valid",
			chain:      vgcommon.Ethereum,
			chainID:    1,
			toAddress:  testRecipientAddress,
			amount:     big.NewInt(1_000_000_000_000_000_000), // 1 ETH
			shouldPass: true,
		},
		{
			name:       "BSC native transfer - valid",
			chain:      vgcommon.BscChain,
			chainID:    56,
			toAddress:  testRecipientAddress,
			amount:     big.NewInt(1_000_000_000_000_000_000), // 1 BNB
			shouldPass: true,
		},
		{
			name:       "Arbitrum native transfer - valid",
			chain:      vgcommon.Arbitrum,
			chainID:    42161,
			toAddress:  testRecipientAddress,
			amount:     big.NewInt(500_000_000_000_000_000), // 0.5 ETH
			shouldPass: true,
		},
		{
			name:       "Avalanche native transfer - valid",
			chain:      vgcommon.Avalanche,
			chainID:    43114,
			toAddress:  testRecipientAddress,
			amount:     big.NewInt(2_000_000_000_000_000_000), // 2 AVAX
			shouldPass: true,
		},
		{
			name:       "Wrong amount - should fail",
			chain:      vgcommon.Ethereum,
			chainID:    1,
			toAddress:  testRecipientAddress,
			amount:     big.NewInt(2_000_000_000_000_000_000), // 2 ETH instead of expected 1 ETH
			shouldPass: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Step 1: Create meta-rule for send operation
			metaRuleHandler := metarule.NewMetaRule()

			chainName := strings.ToLower(tc.chain.String())
			nativeSymbol, err := tc.chain.NativeSymbol()
			require.NoError(t, err)

			expectedAmount := big.NewInt(1_000_000_000_000_000_000).String() // Expected 1 unit
			if tc.chain == vgcommon.Avalanche {
				expectedAmount = "2000000000000000000" // 2 AVAX for Avalanche test
			} else if tc.chain == vgcommon.Arbitrum {
				expectedAmount = "500000000000000000" // 0.5 ETH for Arbitrum test
			}

			sendRule := &types.Rule{
				Resource: fmt.Sprintf("%s.send", chainName),
				Effect:   types.Effect_EFFECT_ALLOW,
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "asset",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_ANY, // Native token (empty)
						},
					},
					{
						ParameterName: "from_address",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: testFromAddress,
							},
						},
					},
					{
						ParameterName: "to_address",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: tc.toAddress,
							},
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: expectedAmount,
							},
						},
					},
				},
			}

			// Step 2: Transform meta-rule to concrete rules
			rules, err := metaRuleHandler.TryFormat(sendRule)
			require.NoError(t, err)
			require.Len(t, rules, 1, "Native send should produce 1 rule")

			// Verify transformed resource
			expectedResource := fmt.Sprintf("%s.%s.transfer", chainName, strings.ToLower(nativeSymbol))
			assert.Equal(t, expectedResource, rules[0].Resource)

			// Step 3: Build transaction
			txBytes := buildEVMTestTx(
				tc.chainID,
				common.HexToAddress(tc.toAddress),
				nil, // No data for native transfer
				tc.amount,
			)

			// Step 4: Evaluate with engine
			evmEngine, err := NewEvm(nativeSymbol)
			require.NoError(t, err)

			err = evmEngine.Evaluate(rules[0], txBytes)
			if tc.shouldPass {
				assert.NoError(t, err, "Expected transaction to pass validation")
			} else {
				assert.Error(t, err, "Expected transaction to fail validation")
			}
		})
	}
}

// TestRecurringSend_ERC20 tests end-to-end ERC20 recurring send flow
func TestRecurringSend_ERC20(t *testing.T) {
	testCases := []struct {
		name         string
		chain        vgcommon.Chain
		chainID      int64
		tokenAddress string
		recipient    string
		amount       *big.Int
		shouldPass   bool
	}{
		{
			name:         "Ethereum USDC transfer - valid",
			chain:        vgcommon.Ethereum,
			chainID:      1,
			tokenAddress: testUSDCAddress,
			recipient:    testRecipientAddress,
			amount:       big.NewInt(1_000_000_000), // 1000 USDC (6 decimals)
			shouldPass:   true,
		},
		{
			name:         "BSC token transfer - valid",
			chain:        vgcommon.BscChain,
			chainID:      56,
			tokenAddress: testUSDCAddress,
			recipient:    testRecipientAddress,
			amount:       big.NewInt(500_000_000), // 500 USDC
			shouldPass:   true,
		},
		{
			name:         "Wrong recipient - should fail",
			chain:        vgcommon.Ethereum,
			chainID:      1,
			tokenAddress: testUSDCAddress,
			recipient:    "0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef", // Wrong recipient
			amount:       big.NewInt(1_000_000_000),
			shouldPass:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Step 1: Create meta-rule for ERC20 send
			metaRuleHandler := metarule.NewMetaRule()

			chainName := strings.ToLower(tc.chain.String())
			nativeSymbol, err := tc.chain.NativeSymbol()
			require.NoError(t, err)

			sendRule := &types.Rule{
				Resource: fmt.Sprintf("%s.send", chainName),
				Effect:   types.Effect_EFFECT_ALLOW,
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "asset",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: tc.tokenAddress,
							},
						},
					},
					{
						ParameterName: "from_address",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: testFromAddress,
							},
						},
					},
					{
						ParameterName: "to_address",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: testRecipientAddress, // Expected recipient
							},
						},
					},
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: tc.amount.String(),
							},
						},
					},
				},
			}

			// Step 2: Transform to concrete rules
			rules, err := metaRuleHandler.TryFormat(sendRule)
			require.NoError(t, err)
			require.Len(t, rules, 1, "ERC20 send should produce 1 rule")

			// Verify transformed resource
			expectedResource := fmt.Sprintf("%s.erc20.transfer", chainName)
			assert.Equal(t, expectedResource, rules[0].Resource)

			// Step 3: Build ERC20 transfer transaction
			erc20Data := erc20.NewErc20().PackTransfer(
				common.HexToAddress(tc.recipient),
				tc.amount,
			)

			txBytes := buildEVMTestTx(
				tc.chainID,
				common.HexToAddress(tc.tokenAddress),
				erc20Data,
				big.NewInt(0), // No ETH value for ERC20 transfer
			)

			// Step 4: Evaluate with engine
			evmEngine, err := NewEvm(nativeSymbol)
			require.NoError(t, err)

			err = evmEngine.Evaluate(rules[0], txBytes)
			if tc.shouldPass {
				assert.NoError(t, err, "Expected transaction to pass validation")
			} else {
				assert.Error(t, err, "Expected transaction to fail validation")
			}
		})
	}
}

// TestRecurringSwap_SameChain_1inch tests same-chain swap via 1inch
func TestRecurringSwap_SameChain_1inch(t *testing.T) {
	testCases := []struct {
		name       string
		chain      vgcommon.Chain
		chainID    int64
		fromAsset  string // Empty for native
		toAsset    string
		amount     *big.Int
		shouldPass bool
	}{
		{
			name:       "Ethereum USDC to WETH swap - valid",
			chain:      vgcommon.Ethereum,
			chainID:    1,
			fromAsset:  testUSDCAddress,
			toAsset:    testWETHAddress,
			amount:     big.NewInt(1_000_000_000), // 1000 USDC
			shouldPass: true,
		},
		{
			name:       "Wrong source token - should fail",
			chain:      vgcommon.Ethereum,
			chainID:    1,
			fromAsset:  testWETHAddress, // Different from expected USDC
			toAsset:    testWETHAddress,
			amount:     big.NewInt(1_000_000_000),
			shouldPass: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Step 1: Create meta-rule for same-chain swap
			metaRuleHandler := metarule.NewMetaRule()

			chainName := strings.ToLower(tc.chain.String())
			nativeSymbol, err := tc.chain.NativeSymbol()
			require.NoError(t, err)

			// Expected values in rule
			expectedFromAsset := testUSDCAddress

			swapRule := &types.Rule{
				Resource: fmt.Sprintf("%s.swap", chainName),
				Effect:   types.Effect_EFFECT_ALLOW,
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "from_asset",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: expectedFromAsset,
							},
						},
					},
					{
						ParameterName: "from_address",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: testFromAddress,
							},
						},
					},
					{
						ParameterName: "from_amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: tc.amount.String(),
							},
						},
					},
					{
						ParameterName: "to_chain",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: tc.chain.String(), // Same chain = 1inch
							},
						},
					},
					{
						ParameterName: "to_asset",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: tc.toAsset,
							},
						},
					},
					{
						ParameterName: "to_address",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: testFromAddress, // Receive to same address
							},
						},
					},
				},
			}

			// Step 2: Transform to concrete rules
			rules, err := metaRuleHandler.TryFormat(swapRule)
			require.NoError(t, err)
			require.GreaterOrEqual(t, len(rules), 1, "Swap should produce at least 1 rule")

			// Find the swap rule (not approve)
			var swapConcreteRule *types.Rule
			for _, r := range rules {
				if strings.Contains(r.Resource, "routerV6_1inch.swap") {
					swapConcreteRule = r
					break
				}
			}
			require.NotNil(t, swapConcreteRule, "Should have 1inch swap rule")

			// Step 3: Build 1inch swap transaction
			swapDesc := routerv6_1inch.GenericRouterSwapDescription{
				SrcToken:        common.HexToAddress(tc.fromAsset),
				DstToken:        common.HexToAddress(tc.toAsset),
				SrcReceiver:     common.HexToAddress(testFromAddress),
				DstReceiver:     common.HexToAddress(testFromAddress),
				Amount:          tc.amount,
				MinReturnAmount: big.NewInt(1),
				Flags:           big.NewInt(0),
			}

			swapData := routerv6_1inch.NewRouterv61inch().PackSwap(
				common.HexToAddress(testFromAddress), // executor
				swapDesc,
				[]byte{0x01, 0x02, 0x03, 0x04}, // data
			)

			txBytes := buildEVMTestTx(
				tc.chainID,
				common.HexToAddress(test1inchRouter),
				swapData,
				big.NewInt(0),
			)

			// Step 4: Evaluate with engine
			evmEngine, err := NewEvm(nativeSymbol)
			require.NoError(t, err)

			err = evmEngine.Evaluate(swapConcreteRule, txBytes)
			if tc.shouldPass {
				assert.NoError(t, err, "Expected swap transaction to pass validation")
			} else {
				assert.Error(t, err, "Expected swap transaction to fail validation")
			}
		})
	}
}

// TestRecurringSwap_CrossChain_ThorChain tests cross-chain swap meta-rule transformation
// Note: Full transaction evaluation requires network access to resolve THORCHAIN_VAULT magic constant
func TestRecurringSwap_CrossChain_ThorChain(t *testing.T) {
	testCases := []struct {
		name      string
		chain     vgcommon.Chain
		fromAsset string
		toChain   string
		toAsset   string
		toAddress string
		amount    *big.Int
	}{
		{
			name:      "Ethereum ETH to Bitcoin BTC swap",
			chain:     vgcommon.Ethereum,
			fromAsset: "", // Native ETH
			toChain:   "Bitcoin",
			toAsset:   "", // Native BTC
			toAddress: "bc1qw589q7vva3wxju9zxz8gt59pfz2frwsazglsj8",
			amount:    big.NewInt(1_000_000_000_000_000_000), // 1 ETH
		},
		{
			name:      "Ethereum USDC to BSC BNB swap",
			chain:     vgcommon.Ethereum,
			fromAsset: testUSDCAddress,
			toChain:   "BSC",
			toAsset:   "", // Native BNB
			toAddress: "0x1234567890abcdef1234567890abcdef12345678",
			amount:    big.NewInt(1_000_000_000),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test meta-rule transformation only (no network access needed)
			metaRuleHandler := metarule.NewMetaRule()

			chainName := strings.ToLower(tc.chain.String())

			swapRule := &types.Rule{
				Resource: fmt.Sprintf("%s.swap", chainName),
				Effect:   types.Effect_EFFECT_ALLOW,
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "from_asset",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: tc.fromAsset,
							},
						},
					},
					{
						ParameterName: "from_address",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: testFromAddress,
							},
						},
					},
					{
						ParameterName: "from_amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: tc.amount.String(),
							},
						},
					},
					{
						ParameterName: "to_chain",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: tc.toChain,
							},
						},
					},
					{
						ParameterName: "to_asset",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: tc.toAsset,
							},
						},
					},
					{
						ParameterName: "to_address",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: tc.toAddress,
							},
						},
					},
				},
			}

			// Transform to concrete rules
			rules, err := metaRuleHandler.TryFormat(swapRule)
			require.NoError(t, err)

			// Verify cross-chain swap produces ThorChain router rule
			var thorchainRule *types.Rule
			var approveRule *types.Rule
			for _, r := range rules {
				if strings.Contains(r.Resource, "thorchain_router.depositWithExpiry") {
					thorchainRule = r
				}
				if strings.Contains(r.Resource, "erc20.approve") {
					approveRule = r
				}
			}

			require.NotNil(t, thorchainRule, "Cross-chain swap should produce ThorChain router rule")

			// Verify rule structure
			assert.Equal(t, types.Effect_EFFECT_ALLOW, thorchainRule.Effect)
			assert.Equal(t, fmt.Sprintf("%s.thorchain_router.depositWithExpiry", chainName), thorchainRule.Resource)

			// Verify target is THORCHAIN_ROUTER magic constant
			assert.Equal(t, types.TargetType_TARGET_TYPE_MAGIC_CONSTANT, thorchainRule.Target.TargetType)
			assert.Equal(t, types.MagicConstant_THORCHAIN_ROUTER, thorchainRule.Target.GetMagicConstant())

			// Verify parameter constraints
			paramByName := make(map[string]*types.ParameterConstraint)
			for _, param := range thorchainRule.ParameterConstraints {
				paramByName[param.ParameterName] = param
			}

			// vault should be THORCHAIN_VAULT magic constant
			assert.Contains(t, paramByName, "vault")
			assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_MAGIC_CONSTANT, paramByName["vault"].Constraint.Type)
			assert.Equal(t, types.MagicConstant_THORCHAIN_VAULT, paramByName["vault"].Constraint.GetMagicConstantValue())

			// memo should be regexp
			assert.Contains(t, paramByName, "memo")
			assert.Equal(t, types.ConstraintType_CONSTRAINT_TYPE_REGEXP, paramByName["memo"].Constraint.Type)

			// For ERC20 from_asset, should also have approve rule
			if tc.fromAsset != "" {
				require.NotNil(t, approveRule, "ERC20 cross-chain swap should have approve rule")
				assert.Contains(t, approveRule.Resource, "erc20.approve")
			}
		})
	}
}

// TestRecurringSend_AllEVMChains tests send across all supported EVM chains
func TestRecurringSend_AllEVMChains(t *testing.T) {
	evmChains := []struct {
		chain   vgcommon.Chain
		chainID int64
	}{
		{vgcommon.Ethereum, 1},
		{vgcommon.BscChain, 56},
		{vgcommon.Arbitrum, 42161},
		{vgcommon.Avalanche, 43114},
		{vgcommon.Base, 8453},
		{vgcommon.Blast, 81457},
		{vgcommon.CronosChain, 25},
		{vgcommon.Optimism, 10},
		{vgcommon.Polygon, 137},
		{vgcommon.Zksync, 324},
	}

	for _, tc := range evmChains {
		t.Run(tc.chain.String(), func(t *testing.T) {
			nativeSymbol, err := tc.chain.NativeSymbol()
			require.NoError(t, err)

			chainName := strings.ToLower(tc.chain.String())
			amount := big.NewInt(1_000_000_000_000_000_000) // 1 native token

			// Create and transform meta-rule
			metaRuleHandler := metarule.NewMetaRule()
			sendRule := &types.Rule{
				Resource: fmt.Sprintf("%s.send", chainName),
				Effect:   types.Effect_EFFECT_ALLOW,
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
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: testFromAddress,
							},
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
								FixedValue: amount.String(),
							},
						},
					},
				},
			}

			rules, err := metaRuleHandler.TryFormat(sendRule)
			require.NoError(t, err)
			require.Len(t, rules, 1)

			// Build and evaluate transaction
			txBytes := buildEVMTestTx(
				tc.chainID,
				common.HexToAddress(testRecipientAddress),
				nil,
				amount,
			)

			evmEngine, err := NewEvm(nativeSymbol)
			require.NoError(t, err)

			err = evmEngine.Evaluate(rules[0], txBytes)
			assert.NoError(t, err, "Native send should pass for %s", tc.chain.String())
		})
	}
}

// TestMetaRuleTransformation_EVM verifies meta-rule transformations produce correct output
func TestMetaRuleTransformation_EVM(t *testing.T) {
	metaRuleHandler := metarule.NewMetaRule()

	t.Run("Native send produces correct resource", func(t *testing.T) {
		rule := &types.Rule{
			Resource: "ethereum.send",
			Effect:   types.Effect_EFFECT_ALLOW,
			ParameterConstraints: []*types.ParameterConstraint{
				{ParameterName: "asset", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_ANY}},
				{ParameterName: "from_address", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: testFromAddress}}},
				{ParameterName: "to_address", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: testRecipientAddress}}},
				{ParameterName: "amount", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: "1000000000000000000"}}},
			},
		}

		rules, err := metaRuleHandler.TryFormat(rule)
		require.NoError(t, err)
		assert.Equal(t, "ethereum.eth.transfer", rules[0].Resource)
	})

	t.Run("ERC20 send produces correct resource", func(t *testing.T) {
		rule := &types.Rule{
			Resource: "ethereum.send",
			Effect:   types.Effect_EFFECT_ALLOW,
			ParameterConstraints: []*types.ParameterConstraint{
				{ParameterName: "asset", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: testUSDCAddress}}},
				{ParameterName: "from_address", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: testFromAddress}}},
				{ParameterName: "to_address", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: testRecipientAddress}}},
				{ParameterName: "amount", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: "1000000000"}}},
			},
		}

		rules, err := metaRuleHandler.TryFormat(rule)
		require.NoError(t, err)
		assert.Equal(t, "ethereum.erc20.transfer", rules[0].Resource)
		assert.Equal(t, testUSDCAddress, rules[0].Target.GetAddress())
	})

	t.Run("Same-chain swap produces 1inch rule", func(t *testing.T) {
		rule := &types.Rule{
			Resource: "ethereum.swap",
			Effect:   types.Effect_EFFECT_ALLOW,
			ParameterConstraints: []*types.ParameterConstraint{
				{ParameterName: "from_asset", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: testUSDCAddress}}},
				{ParameterName: "from_address", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: testFromAddress}}},
				{ParameterName: "from_amount", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: "1000000000"}}},
				{ParameterName: "to_chain", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: "Ethereum"}}},
				{ParameterName: "to_asset", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: testWETHAddress}}},
				{ParameterName: "to_address", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: testFromAddress}}},
			},
		}

		rules, err := metaRuleHandler.TryFormat(rule)
		require.NoError(t, err)

		// Should have 1inch swap rule and ERC20 approve rule
		hasSwapRule := false
		hasApproveRule := false
		for _, r := range rules {
			if strings.Contains(r.Resource, "routerV6_1inch.swap") {
				hasSwapRule = true
			}
			if strings.Contains(r.Resource, "erc20.approve") {
				hasApproveRule = true
			}
		}
		assert.True(t, hasSwapRule, "Should have 1inch swap rule")
		assert.True(t, hasApproveRule, "Should have ERC20 approve rule")
	})

	t.Run("Cross-chain swap produces ThorChain rule", func(t *testing.T) {
		rule := &types.Rule{
			Resource: "ethereum.swap",
			Effect:   types.Effect_EFFECT_ALLOW,
			ParameterConstraints: []*types.ParameterConstraint{
				{ParameterName: "from_asset", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: ""}}}, // Native ETH
				{ParameterName: "from_address", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: testFromAddress}}},
				{ParameterName: "from_amount", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: "1000000000000000000"}}},
				{ParameterName: "to_chain", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: "Bitcoin"}}},
				{ParameterName: "to_asset", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: ""}}}, // Native BTC
				{ParameterName: "to_address", Constraint: &types.Constraint{Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED, Value: &types.Constraint_FixedValue{FixedValue: "bc1qw589q7vva3wxju9zxz8gt59pfz2frwsazglsj8"}}},
			},
		}

		rules, err := metaRuleHandler.TryFormat(rule)
		require.NoError(t, err)

		// Should have ThorChain router rule
		hasThorChainRule := false
		for _, r := range rules {
			if strings.Contains(r.Resource, "thorchain_router.depositWithExpiry") {
				hasThorChainRule = true
			}
		}
		assert.True(t, hasThorChainRule, "Should have ThorChain router rule for cross-chain swap")
	})
}

