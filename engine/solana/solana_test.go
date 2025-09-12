package solana

import (
	"math/big"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	stdcompare "github.com/vultisig/recipes/engine/compare"
	solanautil "github.com/vultisig/recipes/solana"
	"github.com/vultisig/recipes/types"
)

// buildMockSystemTransferTx creates a real Solana system transfer transaction
func buildMockSystemTransferTx(from, to solana.PublicKey, amount uint64) []byte {
	// Create a real system transfer instruction using solana-go
	transferInst := system.NewTransferInstruction(
		amount,
		from,
		to,
	).Build()

	// Create a minimal transaction
	tx, err := solana.NewTransaction(
		[]solana.Instruction{transferInst},
		solana.Hash{}, // Recent blockhash (empty for test)
		solana.TransactionPayer(from),
	)
	if err != nil {
		panic(err)
	}

	// Serialize the transaction
	data, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return data
}

// buildMockMultiInstructionTx creates a transaction with multiple instructions
func buildMockMultiInstructionTx(from, to solana.PublicKey, amount uint64) []byte {
	// Create two transfer instructions
	transferInst1 := system.NewTransferInstruction(
		amount,
		from,
		to,
	).Build()

	transferInst2 := system.NewTransferInstruction(
		amount,
		from,
		to,
	).Build()

	// Create a transaction with multiple instructions
	tx, err := solana.NewTransaction(
		[]solana.Instruction{transferInst1, transferInst2},
		solana.Hash{}, // Recent blockhash (empty for test)
		solana.TransactionPayer(from),
	)
	if err != nil {
		panic(err)
	}

	// Serialize the transaction
	data, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return data
}

func TestEvaluate_InstructionCountValidation(t *testing.T) {
	const lamports = uint64(1000000)

	// Create test keypairs
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()

	engine, err := NewSolana("sol")
	require.NoError(t, err)

	rule := &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: "solana.sol.transfer",
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: toKey.PublicKey().String(),
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
					Required: true,
				},
			},
		},
	}

	tests := []struct {
		name        string
		txBytes     []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "single instruction - should pass",
			txBytes:     buildMockSystemTransferTx(fromKey.PublicKey(), toKey.PublicKey(), lamports),
			expectError: false,
		},
		{
			name:        "multiple instructions - should fail",
			txBytes:     buildMockMultiInstructionTx(fromKey.PublicKey(), toKey.PublicKey(), lamports),
			expectError: true,
			errorMsg:    "only single instruction transactions are allowed, got 2 instructions",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := engine.Evaluate(rule, tt.txBytes)
			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNewSolana(t *testing.T) {
	tests := []struct {
		name         string
		nativeSymbol string
		expectError  bool
	}{
		{
			name:         "valid solana creation",
			nativeSymbol: "sol",
			expectError:  false,
		},
		{
			name:         "valid solana creation uppercase",
			nativeSymbol: "SOL",
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine, err := NewSolana(tt.nativeSymbol)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, engine)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, engine)
				assert.Equal(t, "sol", engine.nativeSymbol) // Should always be lowercase
			}
		})
	}
}

func TestEvaluate_NativeSOLTransfer(t *testing.T) {
	const lamports = uint64(1000000) // 0.001 SOL

	// Create test keypairs
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()

	engine, err := NewSolana("sol")
	require.NoError(t, err)

	txBytes := buildMockSystemTransferTx(fromKey.PublicKey(), toKey.PublicKey(), lamports)

	tests := []struct {
		name        string
		rule        *types.Rule
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid native transfer with fixed amount",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.sol.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: toKey.PublicKey().String(),
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
							Required: true,
						},
					},
				},
			},
			expectError: false,
		},
		{
			name: "valid native transfer with min amount",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.sol.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: toKey.PublicKey().String(),
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
							Value: &types.Constraint_MinValue{
								MinValue: "500000",
							},
							Required: true,
						},
					},
				},
			},
			expectError: false,
		},
		{
			name: "invalid native transfer - wrong amount",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.sol.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: toKey.PublicKey().String(),
					},
				},
				ParameterConstraints: []*types.ParameterConstraint{
					{
						ParameterName: "amount",
						Constraint: &types.Constraint{
							Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
							Value: &types.Constraint_FixedValue{
								FixedValue: "2000000",
							},
							Required: true,
						},
					},
				},
			},
			expectError: true,
			errorMsg:    "failed to compare fixed values",
		},
		{
			name: "invalid native transfer - deny effect",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_DENY,
				Resource: "solana.sol.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: toKey.PublicKey().String(),
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
							Required: true,
						},
					},
				},
			},
			expectError: true,
			errorMsg:    "only allow rules supported",
		},
		{
			name: "invalid - multiple instructions",
			rule: &types.Rule{
				Effect:   types.Effect_EFFECT_ALLOW,
				Resource: "solana.sol.transfer",
				Target: &types.Target{
					TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
					Target: &types.Target_Address{
						Address: toKey.PublicKey().String(),
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
							Required: true,
						},
					},
				},
			},
			expectError: true,
			errorMsg:    "only single instruction transactions are allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use multi-instruction transaction for the multiple instructions test
			var testTxBytes []byte
			if tt.name == "invalid - multiple instructions" {
				testTxBytes = buildMockMultiInstructionTx(fromKey.PublicKey(), toKey.PublicKey(), lamports)
			} else {
				testTxBytes = txBytes
			}

			err := engine.Evaluate(tt.rule, testTxBytes)
			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseTransferAmount(t *testing.T) {
	const lamports = uint64(1000000)

	// Create test keypairs
	fromKey := solana.NewWallet()
	toKey := solana.NewWallet()

	tests := []struct {
		name           string
		setupTx        func() (solana.CompiledInstruction, []solana.PublicKey)
		expectedAmount uint64
		expectError    bool
		errorMsg       string
	}{
		{
			name: "valid system transfer instruction",
			setupTx: func() (solana.CompiledInstruction, []solana.PublicKey) {
				// Build a real system transfer instruction
				transferInst := system.NewTransferInstruction(
					lamports,
					fromKey.PublicKey(),
					toKey.PublicKey(),
				).Build()

				// Create accounts array
				accounts := []solana.PublicKey{
					fromKey.PublicKey(),
					toKey.PublicKey(),
					solanautil.SystemProgramID,
				}

				// Build compiled instruction
				data, _ := transferInst.Data() // Get the instruction data
				compiledInst := solana.CompiledInstruction{
					ProgramIDIndex: 2,              // System program is at index 2
					Accounts:       []uint16{0, 1}, // From and to accounts
					Data:           data,
				}

				return compiledInst, accounts
			},
			expectedAmount: lamports,
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compiledInst, accounts := tt.setupTx()
			amount, err := solanautil.ParseTransferAmount(compiledInst, accounts)
			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
				assert.Equal(t, uint64(0), amount)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedAmount, amount)
			}
		})
	}
}

func TestAssertArg(t *testing.T) {
	tests := []struct {
		name        string
		chain       string
		constraints []*types.ParameterConstraint
		paramName   string
		actualValue interface{}
		expectError bool
		errorMsg    string
	}{
		{
			name:  "any constraint passes",
			chain: "solana",
			constraints: []*types.ParameterConstraint{
				{
					ParameterName: "amount",
					Constraint: &types.Constraint{
						Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
					},
				},
			},
			paramName:   "amount",
			actualValue: big.NewInt(1000),
			expectError: false,
		},
		{
			name:  "parameter not found",
			chain: "solana",
			constraints: []*types.ParameterConstraint{
				{
					ParameterName: "different_param",
					Constraint: &types.Constraint{
						Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
					},
				},
			},
			paramName:   "amount",
			actualValue: big.NewInt(1000),
			expectError: true,
			errorMsg:    "arg not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			switch v := tt.actualValue.(type) {
			case *big.Int:
				err = assertArg(tt.chain, tt.constraints, tt.paramName, v, func(raw string) (stdcompare.Compare[*big.Int], error) {
					val, success := new(big.Int).SetString(raw, 10)
					if !success {
						return nil, assert.AnError
					}
					return &mockBigIntComparer{expected: val}, nil
				})
			case uint64:
				err = assertArg(tt.chain, tt.constraints, tt.paramName, v, func(raw string) (stdcompare.Compare[uint64], error) {
					return &mockU64Comparer{}, nil
				})
			}

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// Mock comparers for testing
type mockBigIntComparer struct {
	expected *big.Int
}

func (m *mockBigIntComparer) Fixed(actual *big.Int) bool { return m.expected.Cmp(actual) == 0 }
func (m *mockBigIntComparer) Min(actual *big.Int) bool   { return actual.Cmp(m.expected) >= 0 }
func (m *mockBigIntComparer) Max(actual *big.Int) bool   { return actual.Cmp(m.expected) <= 0 }
func (m *mockBigIntComparer) Magic(actual *big.Int) bool { return m.Fixed(actual) }

type mockU64Comparer struct{}

func (m *mockU64Comparer) Fixed(actual uint64) bool { return true }
func (m *mockU64Comparer) Min(actual uint64) bool   { return true }
func (m *mockU64Comparer) Max(actual uint64) bool   { return true }
func (m *mockU64Comparer) Magic(actual uint64) bool { return true }
