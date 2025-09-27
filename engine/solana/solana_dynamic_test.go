package solana

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/recipes/types"
)

func TestIDL(t *testing.T) {
	idls, err := loadIDLDir()
	require.NoError(t, err)
	require.NotEmpty(t, idls)

	engine, err := NewSolana()
	require.NoError(t, err)

	for protocolName, idlData := range idls {
		t.Run(fmt.Sprintf("Protocol(%s)", protocolName), func(t *testing.T) {
			for _, instruction := range idlData.Instructions {
				t.Run(fmt.Sprintf("Instruction(%s)", instruction.Name), func(t *testing.T) {
					runTestSuite(t, engine, protocolName, instruction)
				})
			}
		})
	}
}

func runTestSuite(t *testing.T, engine *Solana, protocolName string, instruction idlInstruction) {
	programID := solana.NewWallet().PublicKey()

	accs, args, err := mockAccsAndArgs(instruction)
	assert.NoErrorf(t, err, "Failed to generate accounts and args: %v", err)

	t.Run("PositiveCase", func(t *testing.T) {
		rule := buildPositiveRule(protocolName, instruction, accs, args, programID)
		txBytes, er := genTx(instruction, accs, args, programID)
		assert.NoErrorf(t, er, "Failed to generate transaction: %v", er)

		er = engine.Evaluate(rule, txBytes)
		assert.NoError(t, er, "Positive case should pass")
	})

	t.Run("NegativeCases", func(t *testing.T) {
		t.Run("WrongProgramID", func(t *testing.T) {
			wrongProgramID := solana.NewWallet().PublicKey()
			rule := buildPositiveRule(protocolName, instruction, accs, args, wrongProgramID)
			txBytes, er := genTx(instruction, accs, args, programID)
			assert.NoErrorf(t, er, "Failed to generate transaction: %v", er)

			er = engine.Evaluate(rule, txBytes)
			assert.Error(t, er, "Should fail with wrong program ID")
		})

		if len(instruction.Accounts) > 0 {
			t.Run("WrongAccount", func(t *testing.T) {
				wrongWallets, _, er := mockAccsAndArgs(instruction)
				assert.NoErrorf(t, er, "Failed to generate wrong accounts and args: %v", er)

				rule := buildPositiveRule(protocolName, instruction, wrongWallets, args, programID)
				txBytes, er := genTx(instruction, accs, args, programID)
				assert.NoErrorf(t, er, "Failed to generate transaction: %v", er)

				er = engine.Evaluate(rule, txBytes)
				assert.Error(t, er, "Should fail with wrong account")
			})
		}
	})

	t.Run("MinMaxConstraints", func(t *testing.T) {
		for _, arg := range instruction.Args {
			if arg.Type == argU64 || arg.Type == argU16 || arg.Type == argU8 {
				var minValuePass, maxValuePass, minValueFail, maxValueFail string

				if arg.Type == argU8 {
					minValuePass = "3"  // testValue=6, so min=3 should pass
					maxValuePass = "10" // testValue=6, so max=10 should pass
					minValueFail = "10" // testValue=6, so min=10 should fail
					maxValueFail = "3"  // testValue=6, so max=3 should fail
				} else if arg.Type == argU16 {
					minValuePass = "500"  // testValue=1000, so min=500 should pass
					maxValuePass = "2000" // testValue=1000, so max=2000 should pass
					minValueFail = "2000" // testValue=1000, so min=2000 should fail
					maxValueFail = "500"  // testValue=1000, so max=500 should fail
				} else {
					minValuePass = "500000"  // testValue=1000000, so min=500000 should pass
					maxValuePass = "2000000" // testValue=1000000, so max=2000000 should pass
					minValueFail = "2000000" // testValue=1000000, so min=2000000 should fail
					maxValueFail = "500000"  // testValue=1000000, so max=500000 should fail
				}

				t.Run(fmt.Sprintf("Min_%s", arg.Name), func(t *testing.T) {
					rule := buildMinRule(protocolName, instruction, accs, args, programID, arg.Name, minValuePass)
					txBytes, er := genTx(instruction, accs, args, programID)
					assert.NoErrorf(t, er, "Failed to generate transaction: %v", er)

					er = engine.Evaluate(rule, txBytes)
					assert.NoError(t, er, "Min constraint should pass when value is above minimum")
				})

				t.Run(fmt.Sprintf("Max_%s", arg.Name), func(t *testing.T) {
					rule := buildMaxRule(protocolName, instruction, accs, args, programID, arg.Name, maxValuePass)
					txBytes, er := genTx(instruction, accs, args, programID)
					assert.NoErrorf(t, er, "Failed to generate transaction: %v", er)

					er = engine.Evaluate(rule, txBytes)
					assert.NoError(t, er, "Max constraint should pass when value is below maximum")
				})

				t.Run(fmt.Sprintf("MinFail_%s", arg.Name), func(t *testing.T) {
					rule := buildMinRule(protocolName, instruction, accs, args, programID, arg.Name, minValueFail)
					txBytes, er := genTx(instruction, accs, args, programID)
					assert.NoErrorf(t, er, "Failed to generate transaction: %v", er)

					er = engine.Evaluate(rule, txBytes)
					assert.Error(t, er, "Min constraint should fail when value is below minimum")
				})

				t.Run(fmt.Sprintf("MaxFail_%s", arg.Name), func(t *testing.T) {
					rule := buildMaxRule(protocolName, instruction, accs, args, programID, arg.Name, maxValueFail)
					txBytes, er := genTx(instruction, accs, args, programID)
					assert.NoErrorf(t, er, "Failed to generate transaction: %v", er)

					er = engine.Evaluate(rule, txBytes)
					assert.Error(t, er, "Max constraint should fail when value is above maximum")
				})
			}
		}
	})
}

type argument struct {
	kind  argType
	value string
}

func mockAccsAndArgs(instruction idlInstruction) ([]*solana.Wallet, map[string]argument, error) {
	wallets := make([]*solana.Wallet, len(instruction.Accounts))
	for i := 0; i < len(instruction.Accounts); i++ {
		wallets[i] = solana.NewWallet()
	}

	values := make(map[string]argument)
	for _, a := range instruction.Args {
		switch a.Type {
		case argU64:
			values[a.Name] = argument{kind: a.Type, value: "1000000"}
		case argU16:
			values[a.Name] = argument{kind: a.Type, value: "1000"}
		case argU8:
			values[a.Name] = argument{kind: a.Type, value: "6"}
		case argPublicKey:
			values[a.Name] = argument{kind: a.Type, value: solana.NewWallet().PublicKey().String()}
		case argVec:
			values[a.Name] = argument{kind: a.Type, value: "any"}
		default:
			return nil, nil, fmt.Errorf("unsupported argument type %s for arg %s", a.Type, a.Name)
		}
	}

	return wallets, values, nil
}

func buildPositiveRule(
	protocolName string,
	instruction idlInstruction,
	wallets []*solana.Wallet,
	testValues map[string]argument,
	programID solana.PublicKey,
) *types.Rule {
	resource := fmt.Sprintf("solana.%s.%s", protocolName, instruction.Name)

	constraints := make([]*types.ParameterConstraint, 0)

	for i, account := range instruction.Accounts {
		constraints = append(constraints, &types.ParameterConstraint{
			ParameterName: fmt.Sprintf("account_%s", account.Name),
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				Value: &types.Constraint_FixedValue{
					FixedValue: wallets[i].PublicKey().String(),
				},
				Required: true,
			},
		})
	}

	for _, a := range instruction.Args {
		if v, ok := testValues[a.Name]; ok {
			if a.Type == argVec {
				constraints = append(constraints, &types.ParameterConstraint{
					ParameterName: fmt.Sprintf("arg_%s", a.Name),
					Constraint: &types.Constraint{
						Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
					},
				})
			} else {
				constraints = append(constraints, &types.ParameterConstraint{
					ParameterName: fmt.Sprintf("arg_%s", a.Name),
					Constraint: &types.Constraint{
						Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
						Value: &types.Constraint_FixedValue{
							FixedValue: v.value,
						},
						Required: true,
					},
				})
			}

		}
	}

	return &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: resource,
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: programID.String(),
			},
		},
		ParameterConstraints: constraints,
	}
}

func buildMinRule(
	protocolName string,
	instruction idlInstruction,
	wallets []*solana.Wallet,
	testValues map[string]argument,
	programID solana.PublicKey,
	argName, minValue string,
) *types.Rule {
	resource := fmt.Sprintf("solana.%s.%s", protocolName, instruction.Name)
	constraints := make([]*types.ParameterConstraint, 0)

	for i, account := range instruction.Accounts {
		constraints = append(constraints, &types.ParameterConstraint{
			ParameterName: fmt.Sprintf("account_%s", account.Name),
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				Value: &types.Constraint_FixedValue{
					FixedValue: wallets[i].PublicKey().String(),
				},
				Required: true,
			},
		})
	}

	for _, a := range instruction.Args {
		if v, ok := testValues[a.Name]; ok {
			var constraint *types.Constraint
			if a.Name == argName {
				constraint = &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
					Value: &types.Constraint_MinValue{
						MinValue: minValue,
					},
					Required: true,
				}
			} else {
				if a.Type == argVec {
					constraint = &types.Constraint{
						Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
					}
				} else {
					constraint = &types.Constraint{
						Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
						Value: &types.Constraint_FixedValue{
							FixedValue: v.value,
						},
						Required: true,
					}
				}
			}

			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: fmt.Sprintf("arg_%s", a.Name),
				Constraint:    constraint,
			})
		}
	}

	return &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: resource,
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: programID.String(),
			},
		},
		ParameterConstraints: constraints,
	}
}

func buildMaxRule(
	protocolName string,
	instruction idlInstruction,
	wallets []*solana.Wallet,
	testValues map[string]argument,
	programID solana.PublicKey,
	argName, maxValue string,
) *types.Rule {
	resource := fmt.Sprintf("solana.%s.%s", protocolName, instruction.Name)
	constraints := make([]*types.ParameterConstraint, 0)

	for i, account := range instruction.Accounts {
		constraints = append(constraints, &types.ParameterConstraint{
			ParameterName: fmt.Sprintf("account_%s", account.Name),
			Constraint: &types.Constraint{
				Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
				Value: &types.Constraint_FixedValue{
					FixedValue: wallets[i].PublicKey().String(),
				},
				Required: true,
			},
		})
	}

	for _, a := range instruction.Args {
		if v, ok := testValues[a.Name]; ok {
			var constraint *types.Constraint
			if a.Name == argName {
				constraint = &types.Constraint{
					Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
					Value: &types.Constraint_MaxValue{
						MaxValue: maxValue,
					},
					Required: true,
				}
			} else {
				if a.Type == argVec {
					constraint = &types.Constraint{
						Type: types.ConstraintType_CONSTRAINT_TYPE_ANY,
					}
				} else {
					constraint = &types.Constraint{
						Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
						Value: &types.Constraint_FixedValue{
							FixedValue: v.value,
						},
						Required: true,
					}
				}
			}

			constraints = append(constraints, &types.ParameterConstraint{
				ParameterName: fmt.Sprintf("arg_%s", a.Name),
				Constraint:    constraint,
			})
		}
	}

	return &types.Rule{
		Effect:   types.Effect_EFFECT_ALLOW,
		Resource: resource,
		Target: &types.Target{
			TargetType: types.TargetType_TARGET_TYPE_ADDRESS,
			Target: &types.Target_Address{
				Address: programID.String(),
			},
		},
		ParameterConstraints: constraints,
	}
}

func genTx(
	instruction idlInstruction,
	wallets []*solana.Wallet,
	testValues map[string]argument,
	programID solana.PublicKey,
) ([]byte, error) {
	accounts := make(solana.AccountMetaSlice, 0)
	for i, account := range instruction.Accounts {
		publicKey := wallets[i].PublicKey()
		accounts = append(accounts, &solana.AccountMeta{
			PublicKey:  publicKey,
			IsWritable: account.IsMut,
			IsSigner:   account.IsSigner,
		})
	}

	discriminator := instruction.Metadata.Discriminator

	instructionData, err := buildInstructionData(instruction, testValues, discriminator)
	if err != nil {
		return nil, fmt.Errorf("failed to build instruction data: %w", err)
	}

	inst := solana.NewInstruction(programID, accounts, instructionData)
	payer := solana.NewWallet().PublicKey()
	if len(accounts) > 0 && accounts[0].IsSigner {
		payer = accounts[0].PublicKey
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{inst},
		solana.Hash{},
		solana.TransactionPayer(payer),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	data, err := tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal transaction: %w", err)
	}
	return data, nil
}

func getArgType(args []idlArgument, argName string) argType {
	for _, arg := range args {
		if arg.Name == argName {
			return arg.Type
		}
	}
	return ""
}

func buildInstructionData(
	instruction idlInstruction,
	testValues map[string]argument,
	discriminator []byte,
) ([]byte, error) {
	var buf bytes.Buffer
	encoder := bin.NewBorshEncoder(&buf)
	buf.Write(discriminator)

	for _, arg := range instruction.Args {
		v, ok := testValues[arg.Name]
		if !ok {
			return nil, fmt.Errorf("missing test value for arg %s", arg.Name)
		}

		switch arg.Type {
		case argU64:
			val, err := strconv.ParseUint(v.value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse u64 value %q for arg %s: %w", v, arg.Name, err)
			}
			if er := encoder.WriteUint64(val, bin.LE); er != nil {
				return nil, fmt.Errorf("failed to encode u64 value for arg %s: %w", arg.Name, er)
			}
		case argU16:
			val, err := strconv.ParseUint(v.value, 10, 16)
			if err != nil {
				return nil, fmt.Errorf("failed to parse u16 value %q for arg %s: %w", v, arg.Name, err)
			}
			if er := encoder.WriteUint16(uint16(val), bin.LE); er != nil {
				return nil, fmt.Errorf("failed to encode u16 value for arg %s: %w", arg.Name, er)
			}
		case argU8:
			val, err := strconv.ParseUint(v.value, 10, 8)
			if err != nil {
				return nil, fmt.Errorf("failed to parse u8 value %q for arg %s: %w", v, arg.Name, err)
			}
			if er := encoder.WriteUint8(uint8(val)); er != nil {
				return nil, fmt.Errorf("failed to encode u8 value for arg %s: %w", arg.Name, er)
			}
		case argPublicKey:
			pubkey, err := solana.PublicKeyFromBase58(v.value)
			if err != nil {
				return nil, fmt.Errorf(
					"failed to parse publicKey value %q for arg %s: %w",
					v,
					arg.Name,
					err,
				)
			}
			if er := encoder.WriteBytes(pubkey.Bytes(), false); er != nil {
				return nil, fmt.Errorf("failed to encode publicKey value for arg %s: %w", arg.Name, er)
			}
		case argVec:
			if er := encoder.WriteUint32(0, bin.LE); er != nil {
				return nil, fmt.Errorf("failed to encode empty vector for arg %s: %w", arg.Name, er)
			}
		default:
			return nil, fmt.Errorf("unsupported argument type %s for arg %s", arg.Type, arg.Name)
		}
	}

	return buf.Bytes(), nil
}
