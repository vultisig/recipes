package tron

import (
	"bytes"
	"context"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/mobile-tss-lib/tss"
	rsdk "github.com/vultisig/recipes/sdk"
)

// MockRPCClient implements RPCClient for testing
type MockRPCClient struct {
	BroadcastFunc    func(ctx context.Context, signedTx *SignedTransaction) (*BroadcastResponse, error)
	BroadcastHexFunc func(ctx context.Context, txHex string) (*BroadcastResponse, error)
}

func (m *MockRPCClient) BroadcastTransaction(ctx context.Context, signedTx *SignedTransaction) (*BroadcastResponse, error) {
	if m.BroadcastFunc != nil {
		return m.BroadcastFunc(ctx, signedTx)
	}
	return &BroadcastResponse{
		Result: true,
		TxID:   "mock_txid",
	}, nil
}

func (m *MockRPCClient) BroadcastHex(ctx context.Context, txHex string) (*BroadcastResponse, error) {
	if m.BroadcastHexFunc != nil {
		return m.BroadcastHexFunc(ctx, txHex)
	}
	return &BroadcastResponse{
		Result: true,
		TxID:   "mock_txid_hex",
	}, nil
}

func TestNewSDK(t *testing.T) {
	mockClient := &MockRPCClient{}
	sdk := NewSDK(mockClient)

	assert.NotNil(t, sdk)
	assert.NotNil(t, sdk.rpcClient)
}

func TestNewHTTPRPCClient(t *testing.T) {
	endpoints := []string{"https://api.trongrid.io"}
	client := NewHTTPRPCClient(endpoints)

	assert.NotNil(t, client)
	assert.Equal(t, endpoints, client.endpoints)
	assert.NotNil(t, client.client)
}

func TestSDK_Sign_NoSignatures(t *testing.T) {
	mockClient := &MockRPCClient{}
	sdk := NewSDK(mockClient)

	_, err := sdk.Sign([]byte("tx"), map[string]tss.KeysignResponse{}, make([]byte, 33))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no signatures provided")
}

func TestSDK_Sign_InvalidPubKeyLength(t *testing.T) {
	mockClient := &MockRPCClient{}
	sdk := NewSDK(mockClient)

	sigs := map[string]tss.KeysignResponse{
		"test": {R: "00", S: "00"},
	}

	_, err := sdk.Sign([]byte("tx"), sigs, make([]byte, 32)) // Wrong length
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "pubkey must be 33 or 65 bytes")
}

func TestSDK_Broadcast_NoClient(t *testing.T) {
	sdk := &SDK{rpcClient: nil}

	_, err := sdk.Broadcast(context.Background(), []byte(`{"txID":"test"}`))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "rpc client not configured")
}

func TestSDK_Broadcast_Success(t *testing.T) {
	mockClient := &MockRPCClient{
		BroadcastHexFunc: func(ctx context.Context, txHex string) (*BroadcastResponse, error) {
			// Verify we receive hex-encoded protobuf bytes
			assert.NotEmpty(t, txHex)
			return &BroadcastResponse{
				Result: true,
				TxID:   "TEST_TXID_123",
			}, nil
		},
	}

	sdk := NewSDK(mockClient)

	// Broadcast now expects protobuf-serialized signed transaction bytes
	signedTxBytes := []byte{0x0a, 0x02, 0x01, 0x02, 0x12, 0x41} // minimal protobuf structure
	resp, err := sdk.Broadcast(context.Background(), signedTxBytes)

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Result)
	assert.Equal(t, "TEST_TXID_123", resp.TxID)
}

func TestSDK_ComputeTxHash(t *testing.T) {
	mockClient := &MockRPCClient{}
	sdk := NewSDK(mockClient)

	txBytes := []byte("test transaction")
	hash1 := sdk.ComputeTxHash(txBytes)
	hash2 := sdk.ComputeTxHash(txBytes)

	assert.Equal(t, hash1, hash2)
	assert.Len(t, hash1, 64) // SHA256 produces 32 bytes = 64 hex chars
}

func TestNormalizeLowS(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		shouldError bool
	}{
		{
			name:        "empty input",
			input:       []byte{},
			shouldError: true,
		},
		{
			name:        "valid small S",
			input:       []byte{0x01},
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := normalizeLowS(tt.input)
			if tt.shouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, result, 32)
			}
		})
	}
}

func TestCleanHex(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0xabc123", "abc123"},
		{"0Xabc123", "abc123"},
		{"abc123", "abc123"},
		{"  0xabc123  ", "abc123"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := cleanHex(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMainnetEndpoints(t *testing.T) {
	assert.NotEmpty(t, MainnetEndpoints)
	for _, endpoint := range MainnetEndpoints {
		assert.Contains(t, endpoint, "http")
	}
}

func TestTestnetEndpoints(t *testing.T) {
	assert.NotEmpty(t, TestnetEndpoints)
	for _, endpoint := range TestnetEndpoints {
		assert.Contains(t, endpoint, "http")
	}
}

func TestDeriveSigningHashes_ValidTransaction(t *testing.T) {
	sdk := NewSDK(nil)

	// Sample raw_data bytes (protobuf-serialized transaction body)
	txBytes := []byte{0x0a, 0x02, 0x01, 0x02, 0x12, 0x41, 0x00, 0x01, 0x02, 0x03}

	hashes, err := sdk.DeriveSigningHashes(txBytes, rsdk.DeriveOptions{})
	require.NoError(t, err)

	// Should return exactly 1 hash for TRON transaction
	require.Len(t, hashes, 1)

	// Hash should be 32 bytes (SHA256)
	assert.Len(t, hashes[0].Hash, 32)

	// Message should be 32 bytes (same as hash for TRON)
	assert.Len(t, hashes[0].Message, 32)

	// For TRON, Message and Hash should be identical (both are SHA256 of raw_data)
	assert.True(t, bytes.Equal(hashes[0].Hash, hashes[0].Message))

	// Verify hash is correct SHA256
	expectedHash := sha256.Sum256(txBytes)
	assert.True(t, bytes.Equal(hashes[0].Hash, expectedHash[:]))
}

func TestDeriveSigningHashes_EmptyTransaction(t *testing.T) {
	sdk := NewSDK(nil)

	_, err := sdk.DeriveSigningHashes([]byte{}, rsdk.DeriveOptions{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "empty transaction bytes")
}

func TestDeriveSigningHashes_ConsistentResults(t *testing.T) {
	sdk := NewSDK(nil)

	txBytes := []byte{0x0a, 0x10, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

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

func TestDeriveSigningHashes_DifferentTransactions_DifferentHashes(t *testing.T) {
	sdk := NewSDK(nil)

	txBytes1 := []byte{0x0a, 0x02, 0x01, 0x02}
	txBytes2 := []byte{0x0a, 0x02, 0x03, 0x04} // Different payload

	hashes1, err := sdk.DeriveSigningHashes(txBytes1, rsdk.DeriveOptions{})
	require.NoError(t, err)

	hashes2, err := sdk.DeriveSigningHashes(txBytes2, rsdk.DeriveOptions{})
	require.NoError(t, err)

	// Hashes should be different
	assert.False(t, bytes.Equal(hashes1[0].Hash, hashes2[0].Hash))
}

