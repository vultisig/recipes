package thorchain

import (
	"context"
	"testing"

	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/mobile-tss-lib/tss"
)

// MockRPCClient for testing
type MockRPCClient struct {
	mock.Mock
}

func (m *MockRPCClient) BroadcastTxSync(ctx context.Context, txBytes []byte) (*coretypes.ResultBroadcastTx, error) {
	args := m.Called(ctx, txBytes)
	return args.Get(0).(*coretypes.ResultBroadcastTx), args.Error(1)
}

// Test vectors with deterministic signatures for reproducible tests
var testSignatureVectors = map[string]tss.KeysignResponse{
	"valid_signature": {
		R: "c6fb0eb99c6d0b0b8f1f3e3f3f3f3f3f3f3f3f3f3f3f3f3f3f3f3f3f3f3f3f3f",
		S: "4d2c6e2c6e2c6e2c6e2c6e2c6e2c6e2c6e2c6e2c6e2c6e2c6e2c6e2c6e2c6e2c",
	},
	"high_s_value": {
		R: "a1b2c3d4e5f6789aa1b2c3d4e5f6789aa1b2c3d4e5f6789aa1b2c3d4e5f6789a",
		S: "ffff000000000000ffff000000000000ffff000000000000ffff000000000000",
	},
	"with_hex_prefix": {
		R: "0x123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0",
		S: "0xfedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321",
	},
}

// buildRealisticTHORChainTx creates a valid Cosmos SDK transaction for testing
func buildRealisticTHORChainTx(t *testing.T) []byte {
	t.Helper()

	// Create a bank MsgSend message (most common THORChain tx type)
	msgSend := &banktypes.MsgSend{
		FromAddress: "thor1htrqlgcqc8lexctrx7c2kppq4vnphkatgaj932",
		ToAddress:   "thor1qvlul0ujfrq27ja7uxrp8r7my9juegz0ug3nsg",
		Amount:      sdk.NewCoins(sdk.NewInt64Coin("rune", 100000000)),
	}

	// Pack the message into Any type
	msgAny, err := codectypes.NewAnyWithValue(msgSend)
	require.NoError(t, err)

	// Create transaction body
	txBody := &tx.TxBody{
		Messages:                    []*codectypes.Any{msgAny},
		Memo:                        "test transaction",
		TimeoutHeight:               0,
		ExtensionOptions:            []*codectypes.Any{},
		NonCriticalExtensionOptions: []*codectypes.Any{},
	}

	// Create a test secp256k1 public key
	testPrivKey := secp256k1.GenPrivKey()
	testPubKey := testPrivKey.PubKey()
	pubKeyAny, err := codectypes.NewAnyWithValue(testPubKey)
	require.NoError(t, err)

	// Create auth info with signer info (required for MessageHash)
	authInfo := &tx.AuthInfo{
		SignerInfos: []*tx.SignerInfo{
			{
				PublicKey: pubKeyAny,
				ModeInfo: &tx.ModeInfo{
					Sum: &tx.ModeInfo_Single_{
						Single: &tx.ModeInfo_Single{
							Mode: signing.SignMode_SIGN_MODE_DIRECT,
						},
					},
				},
				Sequence: 0, // Will be updated in MessageHash method
			},
		},
		Fee: &tx.Fee{
			Amount:   []sdk.Coin{},
			GasLimit: 200000,
			Payer:    "",
			Granter:  "",
		},
	}

	// Create unsigned transaction
	unsignedTx := &tx.Tx{
		Body:       txBody,
		AuthInfo:   authInfo,
		Signatures: [][]byte{}, // Empty for unsigned tx
	}

	// Marshal using the SDK codec
	thorSDK := NewSDK(nil)
	txBytes, err := thorSDK.codec.Marshal(unsignedTx)
	require.NoError(t, err)

	return txBytes
}

// buildInvalidProtobufTx creates malformed protobuf data for error testing
func buildInvalidProtobufTx() []byte {
	return []byte{0xFF, 0xFF, 0xFF, 0xFF} // Invalid protobuf data
}

func TestNewSDK(t *testing.T) {
	rpcClient := &MockRPCClient{}
	sdk := NewSDK(rpcClient)

	assert.NotNil(t, sdk)
	assert.Equal(t, rpcClient, sdk.rpcClient)
}

func TestNewCometBFTRPCClient(t *testing.T) {
	// Test with a valid HTTP endpoint format
	client, err := NewCometBFTRPCClient("http://localhost:26657")
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.NotNil(t, client.client)
}

func TestSDK_Sign_WithRealisticTransaction(t *testing.T) {
	sdk := NewSDK(nil)

	// Use realistic THORChain transaction data
	unsignedTxBytes := buildRealisticTHORChainTx(t)

	// Use deterministic signature test vector
	signatures := map[string]tss.KeysignResponse{
		"test_key": testSignatureVectors["valid_signature"],
	}

	signedTxBytes, err := sdk.Sign(unsignedTxBytes, signatures)
	require.NoError(t, err)
	assert.NotEmpty(t, signedTxBytes)

	// Verify the signed transaction can be unmarshaled
	var signedTx tx.Tx
	err = sdk.codec.Unmarshal(signedTxBytes, &signedTx)
	require.NoError(t, err)

	// Verify signature was added
	assert.Len(t, signedTx.Signatures, 1)
	assert.NotEmpty(t, signedTx.Signatures[0])

	// Verify transaction body is preserved
	assert.NotNil(t, signedTx.Body)
	assert.Len(t, signedTx.Body.Messages, 1)
	assert.Equal(t, "test transaction", signedTx.Body.Memo)
}

func TestSDK_Sign_NoSignatures(t *testing.T) {
	sdk := NewSDK(nil)

	unsignedTxBytes := buildRealisticTHORChainTx(t)
	signatures := map[string]tss.KeysignResponse{}

	_, err := sdk.Sign(unsignedTxBytes, signatures)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no signatures provided")
}

func TestSDK_Sign_MultipleSignatures(t *testing.T) {
	sdk := NewSDK(nil)

	unsignedTxBytes := buildRealisticTHORChainTx(t)
	signatures := map[string]tss.KeysignResponse{
		"sig1": {R: "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", S: "fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321"},
		"sig2": {R: "abcd1234567890efabcd1234567890efabcd1234567890efabcd1234567890ef", S: "5678fedcba0987125678fedcba0987125678fedcba0987125678fedcba098712"},
	}

	_, err := sdk.Sign(unsignedTxBytes, signatures)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "expected 1 signature, got 2")
}

// Removed - redundant with TestSDK_Sign_EdgeCases

func TestSDK_Sign_InvalidProtobufData(t *testing.T) {
	sdk := NewSDK(nil)

	// Test with malformed protobuf data
	invalidTxBytes := buildInvalidProtobufTx()
	signatures := map[string]tss.KeysignResponse{
		"test_key": testSignatureVectors["valid_signature"],
	}

	_, err := sdk.Sign(invalidTxBytes, signatures)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal unsigned transaction")
}

func TestSDK_Broadcast(t *testing.T) {
	mockRPC := &MockRPCClient{}
	sdk := NewSDK(mockRPC)

	signedTx := []byte("mock_signed_transaction")
	ctx := context.Background()

	// Mock successful broadcast result
	mockResult := &coretypes.ResultBroadcastTx{
		Code: 0,
		Hash: []byte("mock_hash"),
	}

	mockRPC.On("BroadcastTxSync", ctx, signedTx).Return(mockResult, nil)

	err := sdk.Broadcast(ctx, signedTx)
	assert.NoError(t, err)

	mockRPC.AssertExpectations(t)
}

func TestSDK_Broadcast_NoRPCClient(t *testing.T) {
	sdk := NewSDK(nil)

	signedTx := []byte("mock_signed_transaction")
	ctx := context.Background()

	err := sdk.Broadcast(ctx, signedTx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "rpc client not configured")
}

func TestSDK_Send_Success(t *testing.T) {
	mockRPC := &MockRPCClient{}
	sdk := NewSDK(mockRPC)

	unsignedTxBytes := buildRealisticTHORChainTx(t)
	signatures := map[string]tss.KeysignResponse{
		"test_key": testSignatureVectors["valid_signature"],
	}
	ctx := context.Background()

	// Mock successful broadcast result
	mockResult := &coretypes.ResultBroadcastTx{
		Code: 0,
		Hash: []byte("mock_hash"),
	}

	// Expect the RPC call with any signed transaction bytes
	mockRPC.On("BroadcastTxSync", ctx, mock.AnythingOfType("[]uint8")).Return(mockResult, nil)

	err := sdk.Send(ctx, unsignedTxBytes, signatures)
	assert.NoError(t, err)

	mockRPC.AssertExpectations(t)
}

func TestSDK_Send_SigningError(t *testing.T) {
	mockRPC := &MockRPCClient{}
	sdk := NewSDK(mockRPC)

	// Use invalid protobuf data to trigger signing error
	invalidTxBytes := buildInvalidProtobufTx()
	signatures := map[string]tss.KeysignResponse{
		"test_key": testSignatureVectors["valid_signature"],
	}
	ctx := context.Background()

	err := sdk.Send(ctx, invalidTxBytes, signatures)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to sign transaction")

	// Should not have called broadcast
	mockRPC.AssertNotCalled(t, "BroadcastTxSync")
}

func TestSDK_MessageHash(t *testing.T) {
	sdk := NewSDK(nil)

	// Use realistic transaction data
	unsignedTxBytes := buildRealisticTHORChainTx(t)

	// Test with sample account number and sequence
	accountNumber := uint64(147625)
	sequence := uint64(42)

	hash, err := sdk.MessageHash(unsignedTxBytes, accountNumber, sequence)
	require.NoError(t, err)
	assert.Len(t, hash, 32, "SHA256 hash should be 32 bytes")

	// Test deterministic hashing with same parameters
	hash2, err := sdk.MessageHash(unsignedTxBytes, accountNumber, sequence)
	require.NoError(t, err)
	assert.Equal(t, hash, hash2, "Hash should be deterministic for same inputs")

	// Different account numbers should produce different hashes
	differentHash, err := sdk.MessageHash(unsignedTxBytes, accountNumber+1, sequence)
	require.NoError(t, err)
	assert.NotEqual(t, hash, differentHash, "Different account numbers should produce different hashes")

	// Different sequences should produce different hashes
	differentSeqHash, err := sdk.MessageHash(unsignedTxBytes, accountNumber, sequence+1)
	require.NoError(t, err)
	assert.NotEqual(t, hash, differentSeqHash, "Different sequences should produce different hashes")

	// Different transactions should produce different hashes
	differentTxBytes := buildRealisticTHORChainTx(t)
	// Modify the memo to make it different
	var tx tx.Tx
	err = sdk.codec.Unmarshal(differentTxBytes, &tx)
	require.NoError(t, err)
	tx.Body.Memo = "different memo"
	differentTxBytes, err = sdk.codec.Marshal(&tx)
	require.NoError(t, err)

	differentTxHash, err := sdk.MessageHash(differentTxBytes, accountNumber, sequence)
	require.NoError(t, err)
	assert.NotEqual(t, hash, differentTxHash, "Different transactions should have different hashes")
}

func TestSDK_Sign_RawSignatureFormat(t *testing.T) {
	sdk := NewSDK(nil)
	unsignedTxBytes := buildRealisticTHORChainTx(t)

	tests := []struct {
		name string
		sig  tss.KeysignResponse
	}{
		{
			name: "without hex prefix",
			sig: tss.KeysignResponse{
				R: "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
				S: "fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321",
			},
		},
		{
			name: "with hex prefix",
			sig:  testSignatureVectors["with_hex_prefix"],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signatures := map[string]tss.KeysignResponse{"test_key": tt.sig}

			signedTxBytes, err := sdk.Sign(unsignedTxBytes, signatures)
			require.NoError(t, err)

			// Verify the signed transaction structure
			var signedTx tx.Tx
			err = sdk.codec.Unmarshal(signedTxBytes, &signedTx)
			require.NoError(t, err)

			// Should have exactly one signature
			require.Len(t, signedTx.Signatures, 1)

			// Signature should be 64 bytes (32 R + 32 S)
			signature := signedTx.Signatures[0]
			assert.Len(t, signature, 64, "Signature should be 64 bytes (raw R||S format)")
		})
	}
}

func TestSDK_Integration_Testnet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	// Test connection to THORChain testnet
	client, err := NewCometBFTRPCClient("https://testnet.rpc.thorchain.info")
	if err != nil {
		t.Skipf("unable to connect to testnet: %v", err)
	}

	sdk := NewSDK(client)
	assert.NotNil(t, sdk)
	assert.NotNil(t, sdk.rpcClient)

	// Test message hash computation
	unsignedTxBytes := buildRealisticTHORChainTx(t)
	hash, err := sdk.MessageHash(unsignedTxBytes, 147625, 42)
	require.NoError(t, err)
	assert.Len(t, hash, 32)

	// Note: We don't actually broadcast in tests to avoid spamming the network
}

func TestSDK_Sign_EdgeCases(t *testing.T) {
	sdk := NewSDK(nil)

	tests := []struct {
		name        string
		rHex, sHex  string
		shouldPass  bool
		description string
	}{
		{
			name:        "all zeros",
			rHex:        "0000000000000000000000000000000000000000000000000000000000000000",
			sHex:        "0000000000000000000000000000000000000000000000000000000000000000",
			shouldPass:  false,
			description: "Zero S values are invalid (not in [1, N-1])",
		},
		{
			name:        "high values",
			rHex:        "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			sHex:        "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			shouldPass:  false,
			description: "S values >= curve order are invalid",
		},
		{
			name:        "mixed values",
			rHex:        "123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0",
			sHex:        "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
			shouldPass:  true,
			description: "Mixed byte patterns should work",
		},
		{
			name:        "short R value",
			rHex:        "1234", // Too short
			sHex:        "fedcba0987654321fedcba0987654321fedcba0987654321fedcba0987654321",
			shouldPass:  false,
			description: "Short R value should fail",
		},
	}

	unsignedTxBytes := buildRealisticTHORChainTx(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signatures := map[string]tss.KeysignResponse{
				"test_key": {
					R: tt.rHex,
					S: tt.sHex,
				},
			}

			signedTxBytes, err := sdk.Sign(unsignedTxBytes, signatures)

			if tt.shouldPass {
				require.NoError(t, err, "Expected signing to succeed")
				assert.NotEmpty(t, signedTxBytes)

				// Verify signature format
				var signedTx tx.Tx
				err = sdk.codec.Unmarshal(signedTxBytes, &signedTx)
				require.NoError(t, err)
				require.Len(t, signedTx.Signatures, 1)
				assert.Len(t, signedTx.Signatures[0], 64, "Signature should be 64 bytes")
			} else {
				assert.Error(t, err, "Expected signing to fail")
			}
		})
	}
}

func TestSDK_SignerInfoValidation(t *testing.T) {
	sdk := NewSDK(nil)

	tests := []struct {
		name        string
		modifyTx    func(*tx.Tx)
		expectError string
	}{
		{
			name: "missing AuthInfo",
			modifyTx: func(txData *tx.Tx) {
				txData.AuthInfo = nil
			},
			expectError: "transaction missing AuthInfo",
		},
		{
			name: "missing SignerInfos",
			modifyTx: func(txData *tx.Tx) {
				txData.AuthInfo.SignerInfos = nil
			},
			expectError: "transaction missing SignerInfos",
		},
		{
			name: "missing PublicKey",
			modifyTx: func(txData *tx.Tx) {
				txData.AuthInfo.SignerInfos[0].PublicKey = nil
			},
			expectError: "signer missing PublicKey",
		},
		{
			name: "missing ModeInfo",
			modifyTx: func(txData *tx.Tx) {
				txData.AuthInfo.SignerInfos[0].ModeInfo = nil
			},
			expectError: "signer missing ModeInfo",
		},
		{
			name: "wrong sign mode",
			modifyTx: func(txData *tx.Tx) {
				txData.AuthInfo.SignerInfos[0].ModeInfo = &tx.ModeInfo{
					Sum: &tx.ModeInfo_Single_{
						Single: &tx.ModeInfo_Single{
							Mode: signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
						},
					},
				}
			},
			expectError: "expected SIGN_MODE_DIRECT, got: SIGN_MODE_LEGACY_AMINO_JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build a transaction with proper signer info first
			unsignedTxBytes := buildRealisticTHORChainTx(t)

			// Parse it back to modify
			var unsignedTx tx.Tx
			err := sdk.codec.Unmarshal(unsignedTxBytes, &unsignedTx)
			require.NoError(t, err)

			// Apply the modification
			tt.modifyTx(&unsignedTx)

			// Marshal it back
			modifiedTxBytes, err := sdk.codec.Marshal(&unsignedTx)
			require.NoError(t, err)

			// Try to sign - should fail validation
			signatures := map[string]tss.KeysignResponse{
				"test_key": testSignatureVectors["valid_signature"],
			}

			_, err = sdk.Sign(modifiedTxBytes, signatures)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectError)
		})
	}
}

