package evm_test

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/mobile-tss-lib/tss"

	"github.com/vultisig/recipes/chain/evm/arbitrum"
	"github.com/vultisig/recipes/chain/evm/avalanche"
	"github.com/vultisig/recipes/chain/evm/base"
	"github.com/vultisig/recipes/chain/evm/bsc"
	"github.com/vultisig/recipes/chain/evm/ethereum"
	vultisigTypes "github.com/vultisig/recipes/types"
)

// DynamicFeeTxWithoutSignature mirrors the ethereum package struct for test encoding
type DynamicFeeTxWithoutSignature struct {
	ChainID    *big.Int
	Nonce      uint64
	GasTipCap  *big.Int
	GasFeeCap  *big.Int
	Gas        uint64
	To         *common.Address `rlp:"nil"`
	Value      *big.Int
	Data       []byte
	AccessList types.AccessList
}

// chainTestCase defines test parameters for each EVM chain
type chainTestCase struct {
	name           string
	chainID        int64
	newChain       func() vultisigTypes.Chain
	expectedID     string
	nativeProtocol string
}

var evmChains = []chainTestCase{
	{
		name:           "Ethereum",
		chainID:        1,
		newChain:       ethereum.NewEthereum,
		expectedID:     "ethereum",
		nativeProtocol: "eth",
	},
	{
		name:           "Arbitrum",
		chainID:        42161,
		newChain:       arbitrum.NewArbitrum,
		expectedID:     "arbitrum",
		nativeProtocol: "eth",
	},
	{
		name:           "Avalanche",
		chainID:        43114,
		newChain:       avalanche.NewAvalanche,
		expectedID:     "avalanche",
		nativeProtocol: "avax",
	},
	{
		name:           "Base",
		chainID:        8453,
		newChain:       base.NewBase,
		expectedID:     "base",
		nativeProtocol: "eth",
	},
	{
		name:           "BSC",
		chainID:        56,
		newChain:       bsc.NewBsc,
		expectedID:     "bsc",
		nativeProtocol: "bnb",
	},
}

func TestEVMChains_BasicProperties(t *testing.T) {
	for _, tc := range evmChains {
		t.Run(tc.name, func(t *testing.T) {
			chain := tc.newChain()

			// Test ID
			require.Equal(t, tc.expectedID, chain.ID(), "Chain ID mismatch")

			// Test Name is non-empty
			require.NotEmpty(t, chain.Name(), "Chain name should not be empty")

			// Test Description is non-empty
			require.NotEmpty(t, chain.Description(), "Chain description should not be empty")

			// Test supported protocols includes native protocol
			protocols := chain.SupportedProtocols()
			require.Contains(t, protocols, tc.nativeProtocol, "Should support native protocol")
		})
	}
}

func TestEVMChains_GetNativeProtocol(t *testing.T) {
	for _, tc := range evmChains {
		t.Run(tc.name, func(t *testing.T) {
			chain := tc.newChain()

			protocol, err := chain.GetProtocol(tc.nativeProtocol)
			require.NoError(t, err, "Should get native protocol")
			require.NotNil(t, protocol, "Protocol should not be nil")
			require.Equal(t, tc.nativeProtocol, protocol.ID(), "Protocol ID mismatch")

			// Check transfer function exists
			fn, err := protocol.GetFunction("transfer")
			require.NoError(t, err, "Should have transfer function")
			require.NotNil(t, fn, "Transfer function should not be nil")
		})
	}
}

func TestEVMChains_ParseTransaction(t *testing.T) {
	for _, tc := range evmChains {
		t.Run(tc.name, func(t *testing.T) {
			chain := tc.newChain()

			// Build a test transaction
			to := common.HexToAddress("0x1111111111111111111111111111111111111111")
			txBytes := buildTestTx(t, tc.chainID, &to, big.NewInt(1000000000000000000), nil)

			// Parse the transaction
			decodedTx, err := chain.ParseTransaction(common.Bytes2Hex(txBytes))
			require.NoError(t, err, "Should parse transaction")
			require.NotNil(t, decodedTx, "Decoded tx should not be nil")
			require.Equal(t, to.Hex(), decodedTx.To(), "To address mismatch")
		})
	}
}

func TestEVMChains_ComputeTxHash(t *testing.T) {
	// Test vectors for each chain - using real-world verified transactions where possible
	testCases := []struct {
		name           string
		chainID        int64
		newChain       func() vultisigTypes.Chain
		nonce          uint64
		gasTipCap      *big.Int
		gasFeeCap      *big.Int
		gas            uint64
		to             string
		value          *big.Int
		r              string
		s              string
		recoveryID     string
		expectedTxHash string
	}{
		{
			// Real Ethereum mainnet tx: https://etherscan.io/tx/0xfb58176cf444f9887751a07f749549799b9e6e0a398faa4e29a5cd9a90fa295a
			name:           "Ethereum",
			chainID:        1,
			newChain:       ethereum.NewEthereum,
			nonce:          2553547,
			gasTipCap:      big.NewInt(0),
			gasFeeCap:      big.NewInt(5714758749),
			gas:            23100,
			to:             "0x087b027b0573d4f01345ef8d081e0e7d3b378d14",
			value:          big.NewInt(25767654731246261),
			r:              "d55e81731a80a10a66475fb52021b03b9173359a3220c10db76739b674355f7a",
			s:              "6ebf679597e97da64d048e28fe418b2ca0ef08c2a0583b97d11703dc11cd727b",
			recoveryID:     "01",
			expectedTxHash: "0xfb58176cf444f9887751a07f749549799b9e6e0a398faa4e29a5cd9a90fa295a",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chain := tc.newChain()

			buf := &bytes.Buffer{}
			err := buf.WriteByte(types.DynamicFeeTxType)
			require.NoError(t, err, "WriteByte")

			to := common.HexToAddress(tc.to)
			err = rlp.Encode(buf, &DynamicFeeTxWithoutSignature{
				ChainID:   big.NewInt(tc.chainID),
				Nonce:     tc.nonce,
				GasTipCap: tc.gasTipCap,
				GasFeeCap: tc.gasFeeCap,
				Gas:       tc.gas,
				To:        &to,
				Value:     tc.value,
			})
			require.NoError(t, err, "rlp.Encode")

			txHash, err := chain.ComputeTxHash(buf.Bytes(), []tss.KeysignResponse{{
				R:          tc.r,
				S:          tc.s,
				RecoveryID: tc.recoveryID,
			}})
			require.NoError(t, err, "ComputeTxHash")
			require.Equal(t, tc.expectedTxHash, txHash, "Tx hash mismatch")
		})
	}
}

func TestEVMChains_ComputeTxHash_MultipleSignaturesError(t *testing.T) {
	for _, tc := range evmChains {
		t.Run(tc.name, func(t *testing.T) {
			chain := tc.newChain()

			to := common.HexToAddress("0x1111111111111111111111111111111111111111")
			txBytes := buildTestTx(t, tc.chainID, &to, big.NewInt(1000000000000000000), nil)

			// Try with multiple signatures - should error
			_, err := chain.ComputeTxHash(txBytes, []tss.KeysignResponse{
				{R: "abc", S: "def", RecoveryID: "00"},
				{R: "123", S: "456", RecoveryID: "01"},
			})
			require.Error(t, err, "Should error with multiple signatures")
			require.Contains(t, err.Error(), "expected exactly one signature")
		})
	}
}

// buildTestTx builds a test EIP-1559 transaction
func buildTestTx(t *testing.T, chainID int64, to *common.Address, value *big.Int, data []byte) []byte {
	t.Helper()

	buf := &bytes.Buffer{}
	err := buf.WriteByte(types.DynamicFeeTxType)
	require.NoError(t, err)

	err = rlp.Encode(buf, &DynamicFeeTxWithoutSignature{
		ChainID:    big.NewInt(chainID),
		Nonce:      0,
		GasTipCap:  big.NewInt(2_000_000_000),
		GasFeeCap:  big.NewInt(20_000_000_000),
		Gas:        21000,
		To:         to,
		Value:      value,
		Data:       data,
		AccessList: nil,
	})
	require.NoError(t, err)

	return buf.Bytes()
}

