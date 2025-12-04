package gaia

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/mobile-tss-lib/tss"
)

// MockRPCClient implements RPCClient for testing
type MockRPCClient struct {
	BroadcastTxFunc func(ctx context.Context, txBytes []byte) (*BroadcastTxResponse, error)
}

func (m *MockRPCClient) BroadcastTx(ctx context.Context, txBytes []byte) (*BroadcastTxResponse, error) {
	if m.BroadcastTxFunc != nil {
		return m.BroadcastTxFunc(ctx, txBytes)
	}
	return &BroadcastTxResponse{
		TxResponse: &TxResponse{
			Code:   0,
			TxHash: "ABC123",
		},
	}, nil
}

func TestNewSDK(t *testing.T) {
	mockClient := &MockRPCClient{}
	sdk := NewSDK(mockClient)

	assert.NotNil(t, sdk)
	assert.NotNil(t, sdk.rpcClient)
	assert.NotNil(t, sdk.cdc)
}

func TestNewHTTPRPCClient(t *testing.T) {
	endpoints := []string{"https://example.com"}
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
	assert.Contains(t, err.Error(), "pubkey must be 33 bytes")
}

func TestSDK_Broadcast_NoClient(t *testing.T) {
	sdk := &SDK{rpcClient: nil}

	_, err := sdk.Broadcast(context.Background(), []byte("tx"))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "rpc client not configured")
}

func TestSDK_Broadcast_Success(t *testing.T) {
	mockClient := &MockRPCClient{
		BroadcastTxFunc: func(ctx context.Context, txBytes []byte) (*BroadcastTxResponse, error) {
			return &BroadcastTxResponse{
				TxResponse: &TxResponse{
					Code:   0,
					TxHash: "TEST_HASH_123",
				},
			}, nil
		},
	}

	sdk := NewSDK(mockClient)
	resp, err := sdk.Broadcast(context.Background(), []byte("signed_tx"))

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "TEST_HASH_123", resp.TxResponse.TxHash)
}

func TestSDK_ComputeTxHash(t *testing.T) {
	mockClient := &MockRPCClient{}
	sdk := NewSDK(mockClient)

	// Test that hash computation returns consistent results
	txBytes := []byte("test transaction")
	hash1 := sdk.ComputeTxHash(txBytes)
	hash2 := sdk.ComputeTxHash(txBytes)

	assert.Equal(t, hash1, hash2)
	assert.Len(t, hash1, 64) // SHA256 produces 32 bytes = 64 hex chars
}

func TestGetPubKeyFromBytes(t *testing.T) {
	tests := []struct {
		name        string
		pubKeyBytes []byte
		shouldError bool
	}{
		{
			name:        "valid compressed pubkey",
			pubKeyBytes: make([]byte, 33),
			shouldError: false,
		},
		{
			name:        "invalid length - too short",
			pubKeyBytes: make([]byte, 32),
			shouldError: true,
		},
		{
			name:        "invalid length - too long",
			pubKeyBytes: make([]byte, 65),
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pubKey, err := GetPubKeyFromBytes(tt.pubKeyBytes)
			if tt.shouldError {
				assert.Error(t, err)
				assert.Nil(t, pubKey)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, pubKey)
			}
		})
	}
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
			input:       []byte{0x01}, // Small S value
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
				assert.Len(t, result, 32) // Should be padded to 32 bytes
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

