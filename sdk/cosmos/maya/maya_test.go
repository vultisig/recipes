package maya

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vultisig/mobile-tss-lib/tss"
	cosmossdk "github.com/vultisig/recipes/sdk/cosmos"
)

// MockRPCClient implements RPCClient for testing
type MockRPCClient struct {
	BroadcastFunc func(ctx context.Context, txBytes []byte) (*cosmossdk.BroadcastTxResponse, error)
}

func (m *MockRPCClient) BroadcastTx(ctx context.Context, txBytes []byte) (*cosmossdk.BroadcastTxResponse, error) {
	if m.BroadcastFunc != nil {
		return m.BroadcastFunc(ctx, txBytes)
	}
	return &cosmossdk.BroadcastTxResponse{
		TxResponse: &cosmossdk.TxResponse{
			Code:   0,
			TxHash: "MOCK_TX_HASH",
		},
	}, nil
}

func TestNewSDK(t *testing.T) {
	mockClient := &MockRPCClient{}
	sdk := NewSDK(mockClient)

	assert.NotNil(t, sdk)
	assert.NotNil(t, sdk.SDK)
}

func TestNewMainnetSDK(t *testing.T) {
	sdk := NewMainnetSDK()
	assert.NotNil(t, sdk)
}

func TestNewStagenetSDK(t *testing.T) {
	sdk := NewStagenetSDK()
	assert.NotNil(t, sdk)
}

func TestNewHTTPRPCClient(t *testing.T) {
	endpoints := []string{"https://mayanode.mayachain.info"}
	client := NewHTTPRPCClient(endpoints)

	assert.NotNil(t, client)
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

func TestSDK_Sign_InvalidRLength(t *testing.T) {
	mockClient := &MockRPCClient{}
	sdk := NewSDK(mockClient)

	sigs := map[string]tss.KeysignResponse{
		"test": {R: "1234", S: "abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd"},
	}

	_, err := sdk.Sign([]byte("tx"), sigs, make([]byte, 33))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "r must be 32 bytes")
}

func TestSDK_Sign_InvalidSLength(t *testing.T) {
	mockClient := &MockRPCClient{}
	sdk := NewSDK(mockClient)

	sigs := map[string]tss.KeysignResponse{
		"test": {R: "1234567890123456789012345678901234567890123456789012345678901234", S: "abcd"},
	}

	_, err := sdk.Sign([]byte("tx"), sigs, make([]byte, 33))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "s must be 32 bytes")
}

func TestSDK_Broadcast_NoClient(t *testing.T) {
	sdk := &SDK{SDK: cosmossdk.NewSDK(nil)}

	_, err := sdk.Broadcast(context.Background(), []byte{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "rpc client not configured")
}

func TestSDK_Broadcast_Success(t *testing.T) {
	mockClient := &MockRPCClient{
		BroadcastFunc: func(ctx context.Context, txBytes []byte) (*cosmossdk.BroadcastTxResponse, error) {
			return &cosmossdk.BroadcastTxResponse{
				TxResponse: &cosmossdk.TxResponse{
					Code:   0,
					TxHash: "TEST_MAYA_TX_HASH",
				},
			}, nil
		},
	}

	sdk := NewSDK(mockClient)

	resp, err := sdk.Broadcast(context.Background(), []byte{0x01, 0x02, 0x03})

	require.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint32(0), resp.TxResponse.Code)
	assert.Equal(t, "TEST_MAYA_TX_HASH", resp.TxResponse.TxHash)
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

func TestGetPubKeyFromBytes(t *testing.T) {
	// Valid 33-byte compressed pubkey
	validPubKey := make([]byte, 33)
	validPubKey[0] = 0x02 // Compressed pubkey prefix

	pubKey, err := GetPubKeyFromBytes(validPubKey)
	assert.NoError(t, err)
	assert.NotNil(t, pubKey)

	// Invalid length
	_, err = GetPubKeyFromBytes(make([]byte, 32))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid pubkey length")
}

func TestMainnetEndpoints(t *testing.T) {
	assert.NotEmpty(t, MainnetEndpoints)
	for _, endpoint := range MainnetEndpoints {
		assert.Contains(t, endpoint, "http")
	}
}

func TestStagenetEndpoints(t *testing.T) {
	assert.NotEmpty(t, StagenetEndpoints)
	for _, endpoint := range StagenetEndpoints {
		assert.Contains(t, endpoint, "http")
	}
}

