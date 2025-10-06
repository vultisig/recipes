package xrpl

import (
	"context"
	"encoding/hex"
	"slices"
	"strings"
	"testing"

	"github.com/vultisig/mobile-tss-lib/tss"
	xrpgo "github.com/xyield/xrpl-go/binary-codec"
)

// Test constants for XRP transactions
const (
	// Test XRP addresses (valid format but not real addresses)
	testFromAddress = "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH"
	testToAddress   = "rDNhvtTNPjjUfvsTD1dFGcXxRvGHfWZm4z"

	// Test transaction parameters
	testAmount             = uint64(1000000) // 1 XRP in drops
	testSequence           = uint32(123)
	testFee                = uint64(12) // 12 drops
	testLastLedgerSequence = uint32(99999999)

	// Test compressed public key (33 bytes)
	testPubKeyHex = "0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"

	// Mock signature values for testing (32 bytes each)
	mockRValue = "1234567890123456789012345678901234567890123456789012345678901234"
	mockSValue = "abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd"

	// Test derived key for signature mapping
	testDerivedKey = "test-derived-key"
)

// Mock RPC client for testing
type mockRPCClient struct {
	submitError   error
	submitCalled  bool
	lastTxBlob    string
	shouldSucceed bool
}

func (m *mockRPCClient) SubmitTransaction(ctx context.Context, txBlob string) error {
	m.submitCalled = true
	m.lastTxBlob = txBlob

	if m.submitError != nil {
		return m.submitError
	}

	return nil
}

// Helper function to create test unsigned transaction using XRPL binary codec directly
func createTestUnsignedTransaction(t *testing.T) []byte {
	// Build the JSON-like map the XRPL binary codec expects
	jsonMap := map[string]any{
		"Account":            testFromAddress,
		"TransactionType":    "Payment",
		"Amount":             "1000000", // drops, as string
		"Destination":        testToAddress,
		"Fee":                "12", // drops, as string
		"Sequence":           int(testSequence),
		"LastLedgerSequence": int(testLastLedgerSequence),
	}

	// Encode with the XRPL binary codec
	hexStr, err := xrpgo.Encode(jsonMap)
	if err != nil {
		t.Fatalf("Failed to encode test transaction: %v", err)
	}

	// Convert to bytes
	txBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("Failed to decode hex string: %v", err)
	}

	return txBytes
}

// Helper function to create test signatures map
func createTestSignatures() map[string]tss.KeysignResponse {
	return map[string]tss.KeysignResponse{
		testDerivedKey: {
			Msg:          "test-message",
			R:            mockRValue,
			S:            mockSValue,
			DerSignature: "mock-der-signature",
			RecoveryID:   "00",
		},
	}
}

// Helper function to create test public key bytes
func createTestPubKey(t *testing.T) []byte {
	pubKeyBytes, err := hex.DecodeString(testPubKeyHex)
	if err != nil {
		t.Fatalf("Failed to decode test public key: %v", err)
	}
	return pubKeyBytes
}

func TestNewSDK(t *testing.T) {
	mockClient := &mockRPCClient{}
	sdk := NewSDK(mockClient)

	if sdk == nil {
		t.Fatal("NewSDK returned nil")
	}

	if sdk.rpcClient != mockClient {
		t.Error("SDK rpcClient not set correctly")
	}
}

func TestSDK_Sign_Success(t *testing.T) {
	sdk := &SDK{}

	// Create test unsigned transaction
	unsignedTx := createTestUnsignedTransaction(t)
	signatures := createTestSignatures()
	pubKey := createTestPubKey(t)

	// Note: This test will fail signature verification since we're using mock values
	// XRPL SDK does cryptographic verification unlike BTC SDK which trusts TSS signatures
	_, err := sdk.Sign(unsignedTx, signatures, pubKey)

	// We expect this to fail at signature verification with mock data
	if err == nil {
		t.Error("Expected signature verification to fail with mock data")
	}

	// Verify it's failing at the right place (signature verification)
	if !strings.Contains(err.Error(), "signature") {
		t.Errorf("Expected signature verification error, got: %v", err)
	}
}

func TestSDK_Sign_EmptySignatures(t *testing.T) {
	sdk := &SDK{}

	unsignedTx := createTestUnsignedTransaction(t)
	emptySignatures := make(map[string]tss.KeysignResponse)
	pubKey := createTestPubKey(t)

	_, err := sdk.Sign(unsignedTx, emptySignatures, pubKey)

	if err == nil {
		t.Error("Expected error for empty signatures")
	}

	if !strings.Contains(err.Error(), "no signatures provided") {
		t.Errorf("Expected 'no signatures provided' error, got: %v", err)
	}
}

func TestSDK_Sign_InvalidPubKey(t *testing.T) {
	sdk := &SDK{}

	unsignedTx := createTestUnsignedTransaction(t)
	signatures := createTestSignatures()
	invalidPubKey := []byte{0x01, 0x02} // Too short

	_, err := sdk.Sign(unsignedTx, signatures, invalidPubKey)

	if err == nil {
		t.Error("Expected error for invalid pubkey length")
	}

	if !strings.Contains(err.Error(), "pubkey must be 33 bytes") {
		t.Errorf("Expected pubkey length error, got: %v", err)
	}
}

func TestSDK_Sign_InvalidRSLength(t *testing.T) {
	sdk := &SDK{}

	unsignedTx := createTestUnsignedTransaction(t)
	pubKey := createTestPubKey(t)

	// Test invalid R length
	invalidSignatures := map[string]tss.KeysignResponse{
		testDerivedKey: {
			R: "1234", // Too short
			S: mockSValue,
		},
	}

	_, err := sdk.Sign(unsignedTx, invalidSignatures, pubKey)

	if err == nil {
		t.Error("Expected error for invalid R length")
	}

	if !strings.Contains(err.Error(), "r must be 32 bytes") {
		t.Errorf("Expected R length error, got: %v", err)
	}

	// Test invalid S length
	invalidSignatures = map[string]tss.KeysignResponse{
		testDerivedKey: {
			R: mockRValue,
			S: "abcd", // Too short
		},
	}

	_, err = sdk.Sign(unsignedTx, invalidSignatures, pubKey)

	if err == nil {
		t.Error("Expected error for invalid S length")
	}

	if !strings.Contains(err.Error(), "s must be 32 bytes") {
		t.Errorf("Expected S length error, got: %v", err)
	}
}

func TestSDK_Sign_AlreadySigned(t *testing.T) {
	sdk := &SDK{}

	// Create unsigned transaction and manually add TxnSignature to test error handling
	unsignedTx := createTestUnsignedTransaction(t)
	txHex := hex.EncodeToString(unsignedTx)

	// Decode, add both SigningPubKey and TxnSignature to simulate fully-signed tx
	decoded, err := xrpgo.Decode(strings.ToUpper(txHex))
	if err != nil {
		t.Fatalf("Failed to decode test transaction: %v", err)
	}

	decoded["SigningPubKey"] = testPubKeyHex
	decoded["TxnSignature"] = "3045022100ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890ABCDEF123456789002200123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF"
	signedHex, err := xrpgo.Encode(decoded)
	if err != nil {
		t.Fatalf("Failed to encode test transaction: %v", err)
	}

	signedTx, err := hex.DecodeString(signedHex)
	if err != nil {
		t.Fatalf("Failed to decode signed transaction hex: %v", err)
	}

	signatures := createTestSignatures()
	pubKey := createTestPubKey(t)

	_, err = sdk.Sign(signedTx, signatures, pubKey)

	if err == nil {
		t.Error("Expected error for already signed transaction")
	}

	if !strings.Contains(err.Error(), "already contains TxnSignature") {
		t.Errorf("Expected 'already signed' error, got: %v", err)
	}
}

func TestSDK_Broadcast_Success(t *testing.T) {
	mockClient := &mockRPCClient{shouldSucceed: true}
	sdk := NewSDK(mockClient)

	testTxBytes := []byte{0x12, 0x34, 0x56, 0x78}

	err := sdk.Broadcast(context.Background(), testTxBytes)

	if err != nil {
		t.Errorf("Broadcast failed: %v", err)
	}

	if !mockClient.submitCalled {
		t.Error("Submit was not called")
	}

	expectedHex := hex.EncodeToString(testTxBytes)
	if mockClient.lastTxBlob != expectedHex {
		t.Errorf("Expected tx blob %s, got %s", expectedHex, mockClient.lastTxBlob)
	}
}

func TestSDK_Send_Success(t *testing.T) {
	mockClient := &mockRPCClient{shouldSucceed: true}
	sdk := NewSDK(mockClient)

	// Note: This will fail at signature verification with mock data
	unsignedTx := createTestUnsignedTransaction(t)
	signatures := createTestSignatures()
	pubKey := createTestPubKey(t)

	err := sdk.Send(context.Background(), unsignedTx, signatures, pubKey)

	// Should fail at signature verification
	if err == nil {
		t.Error("Expected signature verification to fail with mock data")
	}

	if !strings.Contains(err.Error(), "failed to sign transaction") {
		t.Errorf("Expected signing error, got: %v", err)
	}
}

func TestSDK_derEncodeRS(t *testing.T) {
	sdk := &SDK{}

	// Test with valid 32-byte R and S values
	r, err := hex.DecodeString(mockRValue)
	if err != nil {
		t.Fatalf("Failed to decode mock R: %v", err)
	}

	s, err := hex.DecodeString(mockSValue)
	if err != nil {
		t.Fatalf("Failed to decode mock S: %v", err)
	}

	der := sdk.derEncodeRS(r, s)

	if len(der) == 0 {
		t.Error("DER encoding returned empty bytes")
	}

	// DER should start with 0x30 (SEQUENCE)
	if der[0] != 0x30 {
		t.Errorf("Expected DER to start with 0x30, got 0x%02x", der[0])
	}

	// Should contain 0x02 markers for INTEGERs
	if !slices.Contains(der, 0x02) {
		t.Error("DER encoding should contain INTEGER markers (0x02)")
	}
}

func TestSDK_verifyRS_InvalidPubKey(t *testing.T) {
	sdk := &SDK{}

	digest := make([]byte, 32)
	r := make([]byte, 32)
	s := make([]byte, 32)
	invalidPubKey := []byte{0x01, 0x02} // Invalid public key

	_, err := sdk.verifyRS(digest, r, s, invalidPubKey)

	if err == nil {
		t.Error("Expected error for invalid public key")
	}

	if !strings.Contains(err.Error(), "parse pubkey") {
		t.Errorf("Expected pubkey parsing error, got: %v", err)
	}
}

// Test that demonstrates the complete workflow structure
func TestXRPLWorkflow_Structure(t *testing.T) {
	t.Log("=== XRPL SDK Workflow Structure Test ===")

	// 1. Create SDK with mock RPC client
	mockClient := &mockRPCClient{shouldSucceed: true}
	sdk := NewSDK(mockClient)
	t.Log("✓ SDK created with RPC client")

	// 2. Create unsigned transaction (using helper function)
	unsignedTx := createTestUnsignedTransaction(t)
	t.Logf("✓ Unsigned transaction created: %d bytes", len(unsignedTx))

	// 3. Create signatures (mock data)
	signatures := createTestSignatures()
	pubKey := createTestPubKey(t)
	t.Log("✓ Mock signatures and pubkey prepared")

	// 4. Sign transaction (will fail verification with mock data, but tests structure)
	_, err := sdk.Sign(unsignedTx, signatures, pubKey)
	// We expect this to fail with mock signatures
	if err != nil && strings.Contains(err.Error(), "signature") {
		t.Log("✓ Signing process reached verification step (expected with mock data)")
	} else if err == nil {
		t.Error("Unexpected success with mock signatures")
	} else {
		t.Errorf("Unexpected error type: %v", err)
	}

	t.Log("✓ XRPL SDK workflow structure test completed")
}
