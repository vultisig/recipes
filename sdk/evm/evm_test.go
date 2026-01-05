package evm

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	reth "github.com/vultisig/recipes/chain/evm/ethereum"
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
