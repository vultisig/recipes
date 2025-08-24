package btc

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vultisig/recipes/types"
)

// Multiple outputs mainnet tx
// https://www.blockchain.com/en/explorer/transactions/btc/8ae9dbdef0078520915c0eb384e063c3c36863e349af81a01eaf59981a504972
const testTxHex = "02000000000102479f9baa3bd517d9fdf6a30b80c206467e33f256d6294e7c8e1b82d413278072010000000000000000f51e16954b6881bfc800a5d7a904b35f630603a04b51ab99c5de93e4fd8163e40100000000000000000240420f0000000000160014753bf12681e2aff9ee7551640052a0695ac4abaa5fd1480000000000160014fd34aaf1265db5e17ec3aa862bd420fd9200a13702483045022100d791955bb0682828c6ef49c520c87c283757cb2f9bcbcb67b8efffcc7c24864402203d65e362e57732de9eed7b5eb707aae546793e2e7e167f8ad7a5579053c5255d012103cf69d9d7b585f2ac5a8e325e2781a890a2a7885e21b53c48e33e345c747671b2024730440220496ca174a3e5409442c75a373a634107e5935f69318bd7412f78f293b8a2961f02205cd0809abc56df2c65031245a2b739c858cd66ce9818d27c328880746ecb366a012103c11745e0d3d6f55f133945974eaaa62e71fa8c8dd1847d573efdb8e88c65f74c00000000"

func newFixed(index int, address, value string) []*types.ParameterConstraint {
	return []*types.ParameterConstraint{{
		ParameterName: fmt.Sprintf("output_address_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
			Value: &types.Constraint_FixedValue{
				FixedValue: address,
			},
			Required: true,
		},
	}, {
		ParameterName: fmt.Sprintf("output_value_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
			Value: &types.Constraint_FixedValue{
				FixedValue: value,
			},
			Required: true,
		},
	}}
}

func newMin(index int, address, minValue string) []*types.ParameterConstraint {
	return []*types.ParameterConstraint{{
		ParameterName: fmt.Sprintf("output_address_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
			Value: &types.Constraint_FixedValue{
				FixedValue: address,
			},
			Required: true,
		},
	}, {
		ParameterName: fmt.Sprintf("output_value_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_MIN,
			Value: &types.Constraint_MinValue{
				MinValue: minValue,
			},
			Required: true,
		},
	}}
}

func newMax(index int, address, maxValue string) []*types.ParameterConstraint {
	return []*types.ParameterConstraint{{
		ParameterName: fmt.Sprintf("output_address_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_FIXED,
			Value: &types.Constraint_FixedValue{
				FixedValue: address,
			},
			Required: true,
		},
	}, {
		ParameterName: fmt.Sprintf("output_value_%d", index),
		Constraint: &types.Constraint{
			Type: types.ConstraintType_CONSTRAINT_TYPE_MAX,
			Value: &types.Constraint_MaxValue{
				MaxValue: maxValue,
			},
			Required: true,
		},
	}}
}

func TestBtc_Evaluate_Fixed(t *testing.T) {
	txBytes, err := hex.DecodeString(testTxHex)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newFixed(
		0,
		"bc1qw5alzf5pu2hlnmn429jqq54qd9dvf2a2jjvvv0",
		"1000000",
	)...)
	params = append(params, newFixed(
		1,
		"bc1ql5624ufxtk67zlkr42rzh4pqlkfqpgfh220msa",
		"4772191",
	)...)

	err = NewBtc().Evaluate(&types.Rule{
		Resource:             "bitcoin.btc.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}

func TestBtc_Evaluate_MinConstraints(t *testing.T) {
	txBytes, err := hex.DecodeString(testTxHex)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newMin(
		0,
		"bc1qw5alzf5pu2hlnmn429jqq54qd9dvf2a2jjvvv0",
		"500000",
	)...) // min 500k, actual 1M
	params = append(params, newMin(
		1,
		"bc1ql5624ufxtk67zlkr42rzh4pqlkfqpgfh220msa",
		"4000000",
	)...) // min 4M, actual 4.77M

	err = NewBtc().Evaluate(&types.Rule{
		Resource:             "bitcoin.btc.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}

func TestBtc_Evaluate_MinConstraints_ShouldFail(t *testing.T) {
	txBytes, err := hex.DecodeString(testTxHex)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newMin(
		0,
		"bc1qw5alzf5pu2hlnmn429jqq54qd9dvf2a2jjvvv0",
		"2000000",
	)...) // min 2M, actual 1M - should fail
	params = append(params, newFixed(
		1,
		"bc1ql5624ufxtk67zlkr42rzh4pqlkfqpgfh220msa",
		"4772191",
	)...)

	err = NewBtc().Evaluate(&types.Rule{
		Resource:             "bitcoin.btc.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "min value constraint failed")
}

func TestBtc_Evaluate_MaxConstraints(t *testing.T) {
	txBytes, err := hex.DecodeString(testTxHex)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newMax(
		0,
		"bc1qw5alzf5pu2hlnmn429jqq54qd9dvf2a2jjvvv0",
		"1500000",
	)...) // max 1.5M, actual 1M
	params = append(params, newMax(
		1,
		"bc1ql5624ufxtk67zlkr42rzh4pqlkfqpgfh220msa",
		"5000000",
	)...) // max 5M, actual 4.77M

	err = NewBtc().Evaluate(&types.Rule{
		Resource:             "bitcoin.btc.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.NoError(t, err)
}

func TestBtc_Evaluate_MaxConstraints_ShouldFail(t *testing.T) {
	txBytes, err := hex.DecodeString(testTxHex)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint
	params = append(params, newMax(
		0,
		"bc1qw5alzf5pu2hlnmn429jqq54qd9dvf2a2jjvvv0",
		"500000",
	)...) // max 500k, actual 1M - should fail
	params = append(params, newFixed(
		1,
		"bc1ql5624ufxtk67zlkr42rzh4pqlkfqpgfh220msa",
		"4772191",
	)...)

	err = NewBtc().Evaluate(&types.Rule{
		Resource:             "bitcoin.btc.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "max value constraint failed")
}

func TestBtc_Evaluate_WrongAddresses_ShouldFail(t *testing.T) {
	txBytes, err := hex.DecodeString(testTxHex)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint

	// Wrong address for output 0 - should fail
	params = append(params, newFixed(
		0,
		"bc1qwrongaddresshere",
		"1000000",
	)...)
	params = append(params, newFixed(
		1,
		"bc1ql5624ufxtk67zlkr42rzh4pqlkfqpgfh220msa",
		"4772191",
	)...)

	err = NewBtc().Evaluate(&types.Rule{
		Resource:             "bitcoin.btc.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fixed value constraint failed")
}

func TestBtc_Evaluate_WrongValues_ShouldFail(t *testing.T) {
	txBytes, err := hex.DecodeString(testTxHex)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint

	// Wrong value for output 0 - should fail
	params = append(params, newFixed(
		0,
		"bc1qw5alzf5pu2hlnmn429jqq54qd9dvf2a2jjvvv0",
		"999999",
	)...) // wrong value
	params = append(params, newFixed(
		1,
		"bc1ql5624ufxtk67zlkr42rzh4pqlkfqpgfh220msa",
		"4772191",
	)...)

	err = NewBtc().Evaluate(&types.Rule{
		Resource:             "bitcoin.btc.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fixed value constraint failed")
}

func TestBtc_Evaluate_MismatchedOutputCounts_ShouldFail(t *testing.T) {
	txBytes, err := hex.DecodeString(testTxHex)
	assert.NoError(t, err)

	var params []*types.ParameterConstraint

	// Only one output constraint instead of two - should fail
	params = append(params, newFixed(
		0,
		"bc1qw5alzf5pu2hlnmn429jqq54qd9dvf2a2jjvvv0",
		"1000000",
	)...)

	err = NewBtc().Evaluate(&types.Rule{
		Resource:             "bitcoin.btc.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, txBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "output count mismatch")
}

func TestBtc_Evaluate_InvalidTxBytes_ShouldFail(t *testing.T) {
	invalidTxBytes := []byte("invalid transaction data")

	var params []*types.ParameterConstraint
	params = append(params, newFixed(
		0,
		"bc1q9sgyna73eew32elmrczk99mu4rrn53sa49zfjr",
		"447175",
	)...)

	err := NewBtc().Evaluate(&types.Rule{
		Resource:             "bitcoin.btc.transfer",
		Effect:               types.Effect_EFFECT_ALLOW,
		ParameterConstraints: params,
	}, invalidTxBytes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse bitcoin transaction")
}
