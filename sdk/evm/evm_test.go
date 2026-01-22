package evm

import (
	"bytes"
	"context"
	"crypto/sha256"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	reth "github.com/vultisig/recipes/chain/evm/ethereum"
	rsdk "github.com/vultisig/recipes/sdk"
	"github.com/vultisig/recipes/sdk/evm/codegen/erc20"
)

func TestSDK_MakeTx_unit(t *testing.T) {
	ctx := context.Background()

	from := common.HexToAddress("0x74823625c397ae30B0a549B9931DF49DF24E3a48")
	to := common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
	value := big.NewInt(0)
	data := erc20.NewErc20().PackTransfer(
		common.HexToAddress("0x9756752DC9f0a947366B6C91bb0487D6c6Bf4d17"),
		big.NewInt(9767000),
	)

	mockRpcClient := newMock_rpcClient(t)
	gasLimit := uint64(84000)
	mockRpcClient.On("EstimateGas", ctx, ethereum.CallMsg{
		From:  from,
		To:    &to,
		Data:  data,
		Value: value,
	}).Return(gasLimit, nil)

	gasTipCap := big.NewInt(1000)
	mockRpcClient.On("SuggestGasTipCap", ctx).Return(gasTipCap, nil)

	feeHistory := &ethereum.FeeHistory{
		BaseFee: []*big.Int{big.NewInt(1000)},
	}
	mockRpcClient.On(
		"FeeHistory",
		ctx,
		uint64(1),
		mock.Anything,
		mock.Anything,
	).Return(feeHistory, nil)

	nonce := uint64(1)
	mockRpcClient.On("PendingNonceAt", ctx, from).Return(nonce, nil)

	mockRpcClientRaw := newMock_rpcClientRaw(t)
	gasFeeCap := new(big.Int).Add(addGas(gasTipCap, 2), addGas(feeHistory.BaseFee[0], 2))
	mockRpcClientRaw.On(
		"CallContext",
		ctx,
		mock.Anything,
		"eth_createAccessList",
		[]interface{}{
			createAccessListArgs{
				From:                 from.Hex(),
				To:                   to.Hex(),
				Gas:                  hexutil.EncodeUint64(gasLimit + gasLimit/2),
				MaxPriorityFeePerGas: hexutil.EncodeBig(addGas(gasTipCap, 2)),
				MaxFeePerGas:         hexutil.EncodeBig(gasFeeCap),
				Value:                hexutil.EncodeBig(value),
				Data:                 hexutil.Encode(data),
			},
			"latest",
		},
	).Return(nil)

	sdk := NewSDK(big.NewInt(1), mockRpcClient, mockRpcClientRaw)
	tx, err := sdk.MakeTx(
		ctx,
		from,
		to,
		value,
		data,
		0, // nonceOffset
	)
	require.NoError(t, err)

	decoded, err := reth.DecodeUnsignedPayload(tx)
	require.NoError(t, err)
	decodedTx := types.NewTx(decoded)
	require.Equal(t, &to, decodedTx.To())
	require.Equal(t, value, decodedTx.Value())
	require.Equal(t, data, decodedTx.Data())
	require.Equal(t, gasLimit+gasLimit/2, decodedTx.Gas())
	require.Equal(t, addGas(gasTipCap, 2), decodedTx.GasTipCap())
	require.Equal(t, gasFeeCap, decodedTx.GasFeeCap())
	require.Equal(t, nonce, decodedTx.Nonce())

	mockRpcClient.AssertExpectations(t)
	mockRpcClientRaw.AssertExpectations(t)
}

func TestSDK_MakeTx_e2e(t *testing.T) {
	if os.Getenv("INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test")
	}

	const rpcURL = "https://eth.llamarpc.com"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rpc, err := ethclient.DialContext(ctx, rpcURL)
	require.NoError(t, err)

	from := common.HexToAddress("0x74823625c397ae30B0a549B9931DF49DF24E3a48")
	to := common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
	value := big.NewInt(0)
	data := erc20.NewErc20().PackTransfer(
		common.HexToAddress("0x9756752DC9f0a947366B6C91bb0487D6c6Bf4d17"),
		big.NewInt(9767000),
	)

	sdk := NewSDK(big.NewInt(1), rpc, rpc.Client())
	tx, err := sdk.MakeTx(
		ctx,
		from,
		to,
		value,
		data,
		0, // nonceOffset
	)
	require.NoError(t, err)

	decoded, err := reth.DecodeUnsignedPayload(tx)
	require.NoError(t, err)
	decodedTx := types.NewTx(decoded)
	require.Equal(t, &to, decodedTx.To())
	require.Equal(t, value, decodedTx.Value())
	require.Equal(t, data, decodedTx.Data())
}

func TestSDK_MakeTx_WithNonceOffset(t *testing.T) {
	ctx := context.Background()

	from := common.HexToAddress("0x74823625c397ae30B0a549B9931DF49DF24E3a48")
	to := common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
	value := big.NewInt(0)
	data := erc20.NewErc20().PackTransfer(
		common.HexToAddress("0x9756752DC9f0a947366B6C91bb0487D6c6Bf4d17"),
		big.NewInt(9767000),
	)

	mockRpcClient := newMock_rpcClient(t)
	gasLimit := uint64(84000)
	mockRpcClient.On("EstimateGas", ctx, ethereum.CallMsg{
		From:  from,
		To:    &to,
		Data:  data,
		Value: value,
	}).Return(gasLimit, nil)

	gasTipCap := big.NewInt(1000)
	mockRpcClient.On("SuggestGasTipCap", ctx).Return(gasTipCap, nil)

	feeHistory := &ethereum.FeeHistory{
		BaseFee: []*big.Int{big.NewInt(1000)},
	}
	mockRpcClient.On(
		"FeeHistory",
		ctx,
		uint64(1),
		mock.Anything,
		mock.Anything,
	).Return(feeHistory, nil)

	baseNonce := uint64(5)
	mockRpcClient.On("PendingNonceAt", ctx, from).Return(baseNonce, nil)

	mockRpcClientRaw := newMock_rpcClientRaw(t)
	mockRpcClientRaw.On(
		"CallContext",
		ctx,
		mock.Anything,
		"eth_createAccessList",
		mock.Anything,
	).Return(nil)

	sdk := NewSDK(big.NewInt(1), mockRpcClient, mockRpcClientRaw)

	// Test with nonce offset of 3
	nonceOffset := uint64(3)
	tx, err := sdk.MakeTx(
		ctx,
		from,
		to,
		value,
		data,
		nonceOffset,
	)
	require.NoError(t, err)

	decoded, err := reth.DecodeUnsignedPayload(tx)
	require.NoError(t, err)
	decodedTx := types.NewTx(decoded)

	// Verify nonce is baseNonce + nonceOffset
	expectedNonce := baseNonce + nonceOffset
	require.Equal(t, expectedNonce, decodedTx.Nonce(), "nonce should be baseNonce + nonceOffset")

	mockRpcClient.AssertExpectations(t)
	mockRpcClientRaw.AssertExpectations(t)
}

// createTestEIP1559TxBytes creates unsigned EIP-1559 transaction bytes for testing
func createTestEIP1559TxBytes(t *testing.T, chainID *big.Int) []byte {
	to := common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
	txData := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     1,
		GasTipCap: big.NewInt(1000000000),  // 1 gwei
		GasFeeCap: big.NewInt(50000000000), // 50 gwei
		Gas:       21000,
		To:        &to,
		Value:     big.NewInt(1000000000000000000), // 1 ETH
		Data:      nil,
	}

	// RLP encode without signature (matches SDK's encode format)
	encoded, err := rlp.EncodeToBytes([]interface{}{
		txData.ChainID,
		txData.Nonce,
		txData.GasTipCap,
		txData.GasFeeCap,
		txData.Gas,
		txData.To,
		txData.Value,
		txData.Data,
		txData.AccessList,
	})
	require.NoError(t, err)

	// Prepend tx type byte
	return append([]byte{types.DynamicFeeTxType}, encoded...)
}

// createTestLegacyTxBytes creates unsigned legacy transaction bytes for testing
func createTestLegacyTxBytes(t *testing.T, chainID *big.Int) []byte {
	to := common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
	txData := &types.LegacyTx{
		Nonce:    1,
		GasPrice: big.NewInt(50000000000), // 50 gwei
		Gas:      21000,
		To:       &to,
		Value:    big.NewInt(1000000000000000000), // 1 ETH
		Data:     nil,
	}

	// For legacy tx, we need to RLP encode with v=chainID, r=0, s=0 for EIP-155
	// to get the signing hash
	encoded, err := rlp.EncodeToBytes([]interface{}{
		txData.Nonce,
		txData.GasPrice,
		txData.Gas,
		txData.To,
		txData.Value,
		txData.Data,
		chainID,    // V = chainID for EIP-155 signing
		uint64(0),  // R = 0
		uint64(0),  // S = 0
	})
	require.NoError(t, err)

	// Legacy txs don't have a type prefix, but SDK expects one
	// Actually for legacy, let's use access list type which is simpler
	return encoded
}

func TestDeriveSigningHashes_ValidEIP1559Transaction(t *testing.T) {
	chainID := big.NewInt(1) // Ethereum mainnet
	sdk := NewSDK(chainID, nil, nil)

	txBytes := createTestEIP1559TxBytes(t, chainID)

	hashes, err := sdk.DeriveSigningHashes(txBytes, rsdk.DeriveOptions{})
	require.NoError(t, err)

	// Should return exactly 1 hash for EVM transaction
	require.Len(t, hashes, 1)

	// Hash should be 32 bytes (SHA256 output)
	assert.Len(t, hashes[0].Hash, 32)

	// Message should be 32 bytes (Keccak256 signing hash)
	assert.Len(t, hashes[0].Message, 32)

	// Hash should be SHA256 of Message
	expectedHash := sha256.Sum256(hashes[0].Message)
	assert.True(t, bytes.Equal(hashes[0].Hash, expectedHash[:]))
}

func TestDeriveSigningHashes_DifferentChainIDs(t *testing.T) {
	testCases := []struct {
		name    string
		chainID *big.Int
	}{
		{"ethereum_mainnet", big.NewInt(1)},
		{"arbitrum", big.NewInt(42161)},
		{"polygon", big.NewInt(137)},
		{"base", big.NewInt(8453)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sdk := NewSDK(tc.chainID, nil, nil)
			txBytes := createTestEIP1559TxBytes(t, tc.chainID)

			hashes, err := sdk.DeriveSigningHashes(txBytes, rsdk.DeriveOptions{})
			require.NoError(t, err)
			require.Len(t, hashes, 1)
			assert.Len(t, hashes[0].Hash, 32)
			assert.Len(t, hashes[0].Message, 32)
		})
	}
}

func TestDeriveSigningHashes_InvalidTransactionBytes(t *testing.T) {
	sdk := NewSDK(big.NewInt(1), nil, nil)

	testCases := []struct {
		name  string
		input []byte
	}{
		{"random_bytes", []byte{0x01, 0x02, 0x03, 0x04}},
		{"empty_bytes", []byte{}},
		{"single_byte", []byte{0x02}},
		{"invalid_type_prefix", []byte{0xFF, 0x01, 0x02, 0x03}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := sdk.DeriveSigningHashes(tc.input, rsdk.DeriveOptions{})
			require.Error(t, err)
		})
	}
}

func TestDeriveSigningHashes_ConsistentResults(t *testing.T) {
	chainID := big.NewInt(1)
	sdk := NewSDK(chainID, nil, nil)
	txBytes := createTestEIP1559TxBytes(t, chainID)

	// Call multiple times
	hashes1, err := sdk.DeriveSigningHashes(txBytes, rsdk.DeriveOptions{})
	require.NoError(t, err)

	hashes2, err := sdk.DeriveSigningHashes(txBytes, rsdk.DeriveOptions{})
	require.NoError(t, err)

	// Results should be identical
	require.Len(t, hashes1, 1)
	require.Len(t, hashes2, 1)
	assert.True(t, bytes.Equal(hashes1[0].Hash, hashes2[0].Hash))
	assert.True(t, bytes.Equal(hashes1[0].Message, hashes2[0].Message))
}

func TestDeriveSigningHashes_MalformedRLP(t *testing.T) {
	sdk := NewSDK(big.NewInt(1), nil, nil)

	// Valid type prefix but malformed RLP
	malformedTx := append([]byte{types.DynamicFeeTxType}, []byte{0xc0, 0xc0, 0xff}...)

	_, err := sdk.DeriveSigningHashes(malformedTx, rsdk.DeriveOptions{})
	require.Error(t, err)
}
