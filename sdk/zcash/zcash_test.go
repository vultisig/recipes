package zcash

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/vultisig/mobile-tss-lib/tss"
)

// Test public key (compressed)
const testPubKeyHex = "0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"

// Mock signature values for testing (not cryptographically valid)
const (
	mockDerSignature = "3044022012345678901234567890123456789012345678901234567890123456789012340220abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	mockRValue       = "1234567890123456789012345678901234567890123456789012345678901234"
	mockSValue       = "abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	mockRecoveryID   = "00"
)

// Test transaction values
const (
	testPrevTxHash   = "9b1a2f3e4d5c6b7a8f9e1f0b4b4d5b2b4b8d3e0c8050b5b0e3f7650145cdabcd"
	testInputValue   = uint64(100000000) // 1 ZEC in zatoshis
	testOutputValue  = int64(99990000)   // 0.9999 ZEC (minus fee)
	testBroadcastTx  = "test-broadcast-hash"
)

// P2PKH script for t1Rv4exT7bqhZqi2j7xz8bUHDMxwosrjADU (example)
var testP2PKHScript = []byte{
	0x76, 0xa9, 0x14, // OP_DUP OP_HASH160 PUSH20
	0xab, 0xcd, 0xef, 0xab, 0xcd, 0xef, 0xab, 0xcd,
	0xef, 0xab, 0xcd, 0xef, 0xab, 0xcd, 0xef, 0xab,
	0xcd, 0xef, 0xab, 0xcd, // 20-byte hash
	0x88, 0xac, // OP_EQUALVERIFY OP_CHECKSIG
}

// Mock broadcaster for testing
type mockBroadcaster struct {
	lastTx    []byte
	returnErr error
}

func (m *mockBroadcaster) BroadcastTransaction(signedTx []byte) (string, error) {
	m.lastTx = signedTx
	if m.returnErr != nil {
		return "", m.returnErr
	}
	return testBroadcastTx, nil
}

func createTestUnsignedTx(t *testing.T) *UnsignedTx {
	t.Helper()

	pubKey, err := hex.DecodeString(testPubKeyHex)
	if err != nil {
		t.Fatalf("Failed to decode pubkey: %v", err)
	}

	inputs := []TxInput{
		{
			TxHash:   testPrevTxHash,
			Index:    0,
			Value:    testInputValue,
			Script:   testP2PKHScript,
			Sequence: 0xffffffff,
		},
	}

	outputs := []*TxOutput{
		{
			Value:  testOutputValue,
			Script: testP2PKHScript,
		},
	}

	sdk := &SDK{}

	// Serialize unsigned tx
	rawBytes, err := sdk.SerializeUnsignedTx(inputs, outputs)
	if err != nil {
		t.Fatalf("Failed to serialize unsigned tx: %v", err)
	}

	// Calculate sig hashes
	sigHashes := make([][]byte, len(inputs))
	for i := range inputs {
		sigHash, err := sdk.CalculateSigHash(inputs, outputs, i)
		if err != nil {
			t.Fatalf("Failed to calculate sig hash for input %d: %v", i, err)
		}
		sigHashes[i] = sigHash
	}

	return &UnsignedTx{
		Inputs:    inputs,
		Outputs:   outputs,
		PubKey:    pubKey,
		RawBytes:  rawBytes,
		SigHashes: sigHashes,
	}
}

func TestSDK_DeriveKeyFromMessage(t *testing.T) {
	sdk := &SDK{}

	testMessage := []byte("test message hash")
	derivedKey := sdk.DeriveKeyFromMessage(testMessage)

	// Verify the key derivation process manually
	hash := sha256.Sum256(testMessage)
	expectedKey := base64.StdEncoding.EncodeToString(hash[:])

	if derivedKey != expectedKey {
		t.Errorf("Key derivation mismatch. Expected %s, got %s", expectedKey, derivedKey)
	}

	// Test that same message produces same key
	derivedKey2 := sdk.DeriveKeyFromMessage(testMessage)
	if derivedKey != derivedKey2 {
		t.Error("Same message should produce same derived key")
	}

	// Test that different messages produce different keys
	differentMessage := []byte("different message")
	differentKey := sdk.DeriveKeyFromMessage(differentMessage)
	if derivedKey == differentKey {
		t.Error("Different messages should produce different keys")
	}
}

func TestSDK_CalculateSigHash(t *testing.T) {
	sdk := &SDK{}

	inputs := []TxInput{
		{
			TxHash:   testPrevTxHash,
			Index:    0,
			Value:    testInputValue,
			Script:   testP2PKHScript,
			Sequence: 0xffffffff,
		},
	}

	outputs := []*TxOutput{
		{
			Value:  testOutputValue,
			Script: testP2PKHScript,
		},
	}

	// Test calculating sig hash for valid input
	sigHash, err := sdk.CalculateSigHash(inputs, outputs, 0)
	if err != nil {
		t.Fatalf("CalculateSigHash failed: %v", err)
	}

	if len(sigHash) != 32 {
		t.Errorf("Expected signature hash length 32, got %d", len(sigHash))
	}

	// Test with invalid input index
	_, err = sdk.CalculateSigHash(inputs, outputs, 1)
	if err == nil {
		t.Error("Expected error for invalid input index")
	}

	_, err = sdk.CalculateSigHash(inputs, outputs, -1)
	if err == nil {
		t.Error("Expected error for negative input index")
	}

	// Test that same inputs produce same hash
	sigHash2, err := sdk.CalculateSigHash(inputs, outputs, 0)
	if err != nil {
		t.Fatalf("Second CalculateSigHash failed: %v", err)
	}
	if !bytes.Equal(sigHash, sigHash2) {
		t.Error("Same inputs should produce same sig hash")
	}
}

func TestSDK_SerializeUnsignedTx(t *testing.T) {
	sdk := &SDK{}

	inputs := []TxInput{
		{
			TxHash:   testPrevTxHash,
			Index:    0,
			Value:    testInputValue,
			Script:   testP2PKHScript,
			Sequence: 0xffffffff,
		},
	}

	outputs := []*TxOutput{
		{
			Value:  testOutputValue,
			Script: testP2PKHScript,
		},
	}

	rawBytes, err := sdk.SerializeUnsignedTx(inputs, outputs)
	if err != nil {
		t.Fatalf("SerializeUnsignedTx failed: %v", err)
	}

	// Verify minimum expected size:
	// - Header (4) + Version group ID (4) = 8
	// - Input count (1) + inputs (32+4+1+4 = 41) = 42
	// - Output count (1) + outputs (8+1+25 = 34) = 35
	// - Lock time (4) + Expiry height (4) + Value balance (8) = 16
	// - Shielded counts (3) = 3
	// Total minimum: 8 + 42 + 35 + 16 + 3 = 104 bytes
	minExpected := 100 // Approximate minimum
	if len(rawBytes) < minExpected {
		t.Errorf("Serialized tx too small: got %d bytes, expected at least %d", len(rawBytes), minExpected)
	}

	// Verify header bytes (v4 with overwintered flag)
	expectedVersion := []byte{0x04, 0x00, 0x00, 0x80}
	if !bytes.Equal(rawBytes[:4], expectedVersion) {
		t.Errorf("Version mismatch: got %x, expected %x", rawBytes[:4], expectedVersion)
	}

	// Verify version group ID (Sapling)
	expectedVersionGroup := []byte{0x85, 0x20, 0x2f, 0x89}
	if !bytes.Equal(rawBytes[4:8], expectedVersionGroup) {
		t.Errorf("Version group ID mismatch: got %x, expected %x", rawBytes[4:8], expectedVersionGroup)
	}
}

func TestSDK_SerializeUnsignedTx_InvalidInput(t *testing.T) {
	sdk := &SDK{}

	// Test with invalid tx hash
	inputs := []TxInput{
		{
			TxHash: "invalid-hex",
			Index:  0,
			Value:  testInputValue,
			Script: testP2PKHScript,
		},
	}

	outputs := []*TxOutput{
		{
			Value:  testOutputValue,
			Script: testP2PKHScript,
		},
	}

	_, err := sdk.SerializeUnsignedTx(inputs, outputs)
	if err == nil {
		t.Error("Expected error for invalid tx hash hex")
	}

	// Test with negative output value
	validInputs := []TxInput{
		{
			TxHash: testPrevTxHash,
			Index:  0,
			Value:  testInputValue,
			Script: testP2PKHScript,
		},
	}

	negativeOutputs := []*TxOutput{
		{
			Value:  -1,
			Script: testP2PKHScript,
		},
	}

	_, err = sdk.SerializeUnsignedTx(validInputs, negativeOutputs)
	if err == nil {
		t.Error("Expected error for negative output value")
	}
}

func TestSDK_Sign(t *testing.T) {
	broadcaster := &mockBroadcaster{}
	sdk := NewSDK(broadcaster)

	unsignedTx := createTestUnsignedTx(t)

	// Create signatures map with correct derived key
	derivedKey := sdk.DeriveKeyFromMessage(unsignedTx.SigHashes[0])
	signatures := map[string]tss.KeysignResponse{
		derivedKey: {
			Msg:          base64.StdEncoding.EncodeToString(unsignedTx.SigHashes[0]),
			DerSignature: mockDerSignature,
			R:            mockRValue,
			S:            mockSValue,
			RecoveryID:   mockRecoveryID,
		},
	}

	// Sign the transaction
	signedTx, err := sdk.Sign(unsignedTx, signatures)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	if len(signedTx) == 0 {
		t.Error("Signed transaction is empty")
	}

	// Signed tx should be larger than unsigned (includes scriptsig)
	if len(signedTx) <= len(unsignedTx.RawBytes) {
		t.Error("Signed transaction should be larger than unsigned")
	}

	// Verify header bytes are still correct
	expectedVersion := []byte{0x04, 0x00, 0x00, 0x80}
	if !bytes.Equal(signedTx[:4], expectedVersion) {
		t.Errorf("Signed tx version mismatch: got %x, expected %x", signedTx[:4], expectedVersion)
	}

	t.Logf("✓ Transaction signed successfully, %d bytes", len(signedTx))
}

func TestSDK_Sign_MissingSignature(t *testing.T) {
	sdk := NewSDK(nil)

	unsignedTx := createTestUnsignedTx(t)

	// Empty signatures map
	signatures := map[string]tss.KeysignResponse{}

	_, err := sdk.Sign(unsignedTx, signatures)
	if err == nil {
		t.Error("Expected error for missing signature")
	}

	if !bytes.Contains([]byte(err.Error()), []byte("missing signature")) {
		t.Errorf("Expected 'missing signature' error, got: %v", err)
	}
}

func TestSDK_Sign_InvalidSignature(t *testing.T) {
	sdk := NewSDK(nil)

	unsignedTx := createTestUnsignedTx(t)

	// Create signatures map with invalid DER signature
	derivedKey := sdk.DeriveKeyFromMessage(unsignedTx.SigHashes[0])
	signatures := map[string]tss.KeysignResponse{
		derivedKey: {
			DerSignature: "invalid-hex",
		},
	}

	_, err := sdk.Sign(unsignedTx, signatures)
	if err == nil {
		t.Error("Expected error for invalid DER signature")
	}
}

func TestSDK_Broadcast(t *testing.T) {
	broadcaster := &mockBroadcaster{}
	sdk := NewSDK(broadcaster)

	signedTx := []byte{0x01, 0x02, 0x03, 0x04}

	txHash, err := sdk.Broadcast(signedTx)
	if err != nil {
		t.Fatalf("Broadcast failed: %v", err)
	}

	if txHash != testBroadcastTx {
		t.Errorf("Expected tx hash %s, got %s", testBroadcastTx, txHash)
	}

	if !bytes.Equal(broadcaster.lastTx, signedTx) {
		t.Error("Broadcaster received wrong transaction")
	}
}

func TestSDK_Broadcast_NoBroadcaster(t *testing.T) {
	sdk := NewSDK(nil)

	_, err := sdk.Broadcast([]byte{0x01})
	if err == nil {
		t.Error("Expected error when broadcaster is nil")
	}
}

func TestSDK_Send(t *testing.T) {
	broadcaster := &mockBroadcaster{}
	sdk := NewSDK(broadcaster)

	unsignedTx := createTestUnsignedTx(t)

	// Create valid signatures
	derivedKey := sdk.DeriveKeyFromMessage(unsignedTx.SigHashes[0])
	signatures := map[string]tss.KeysignResponse{
		derivedKey: {
			DerSignature: mockDerSignature,
		},
	}

	txHash, err := sdk.Send(unsignedTx, signatures)
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}

	if txHash != testBroadcastTx {
		t.Errorf("Expected tx hash %s, got %s", testBroadcastTx, txHash)
	}

	if broadcaster.lastTx == nil {
		t.Error("Broadcaster should have received transaction")
	}
}

func TestSDK_ComputeTxHash(t *testing.T) {
	sdk := &SDK{}

	// Test with a sample transaction
	sampleTx := []byte{
		0x04, 0x00, 0x00, 0x80, // version
		0x85, 0x20, 0x2f, 0x89, // version group id
		0x00,                   // input count
		0x01,                   // output count
		0x00, 0x00, 0x00, 0x00, // value (8 bytes)
		0x00, 0x00, 0x00, 0x00,
		0x00,                   // script length
		0x00, 0x00, 0x00, 0x00, // lock time
		0x00, 0x00, 0x00, 0x00, // expiry height
		0x00, 0x00, 0x00, 0x00, // value balance (8 bytes)
		0x00, 0x00, 0x00, 0x00,
		0x00, // shielded spends
		0x00, // shielded outputs
		0x00, // joinsplits
	}

	txHash := sdk.ComputeTxHash(sampleTx)

	// Hash should be 64 hex characters
	if len(txHash) != 64 {
		t.Errorf("Expected tx hash length 64, got %d", len(txHash))
	}

	// Same input should produce same hash
	txHash2 := sdk.ComputeTxHash(sampleTx)
	if txHash != txHash2 {
		t.Error("Same transaction should produce same hash")
	}
}

func TestSDK_MultipleInputs(t *testing.T) {
	sdk := &SDK{}

	pubKey, err := hex.DecodeString(testPubKeyHex)
	if err != nil {
		t.Fatalf("Failed to decode pubkey: %v", err)
	}

	// Create transaction with multiple inputs
	inputs := []TxInput{
		{
			TxHash:   testPrevTxHash,
			Index:    0,
			Value:    50000000,
			Script:   testP2PKHScript,
			Sequence: 0xffffffff,
		},
		{
			TxHash:   testPrevTxHash,
			Index:    1,
			Value:    50000000,
			Script:   testP2PKHScript,
			Sequence: 0xffffffff,
		},
	}

	outputs := []*TxOutput{
		{
			Value:  99990000,
			Script: testP2PKHScript,
		},
	}

	// Calculate sig hashes for both inputs
	sigHash0, err := sdk.CalculateSigHash(inputs, outputs, 0)
	if err != nil {
		t.Fatalf("Failed to calculate sig hash for input 0: %v", err)
	}

	sigHash1, err := sdk.CalculateSigHash(inputs, outputs, 1)
	if err != nil {
		t.Fatalf("Failed to calculate sig hash for input 1: %v", err)
	}

	// Sig hashes should be different for different inputs
	if bytes.Equal(sigHash0, sigHash1) {
		t.Error("Different inputs should have different sig hashes")
	}

	// Create unsigned tx
	rawBytes, err := sdk.SerializeUnsignedTx(inputs, outputs)
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}

	unsignedTx := &UnsignedTx{
		Inputs:    inputs,
		Outputs:   outputs,
		PubKey:    pubKey,
		RawBytes:  rawBytes,
		SigHashes: [][]byte{sigHash0, sigHash1},
	}

	// Create signatures for both inputs
	signatures := map[string]tss.KeysignResponse{
		sdk.DeriveKeyFromMessage(sigHash0): {DerSignature: mockDerSignature},
		sdk.DeriveKeyFromMessage(sigHash1): {DerSignature: mockDerSignature},
	}

	// Sign should succeed
	signedTx, err := sdk.Sign(unsignedTx, signatures)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	if len(signedTx) == 0 {
		t.Error("Signed transaction is empty")
	}

	t.Logf("✓ Multi-input transaction signed successfully, %d bytes", len(signedTx))
}

func TestWriteCompactSize(t *testing.T) {
	tests := []struct {
		name     string
		value    uint64
		expected []byte
	}{
		{"zero", 0, []byte{0x00}},
		{"small", 100, []byte{0x64}},
		{"max_one_byte", 252, []byte{0xfc}},
		{"two_byte_min", 253, []byte{0xfd, 0xfd, 0x00}},
		{"two_byte", 1000, []byte{0xfd, 0xe8, 0x03}},
		{"four_byte_min", 0x10000, []byte{0xfe, 0x00, 0x00, 0x01, 0x00}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			writeCompactSize(&buf, tt.value)
			if !bytes.Equal(buf.Bytes(), tt.expected) {
				t.Errorf("writeCompactSize(%d) = %x, want %x", tt.value, buf.Bytes(), tt.expected)
			}
		})
	}
}

func TestTrim0x(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0xabcd", "abcd"},
		{"0Xabcd", "abcd"},
		{"abcd", "abcd"},
		{"0x", ""},
		{"", ""},
	}

	for _, tt := range tests {
		result := trim0x(tt.input)
		if result != tt.expected {
			t.Errorf("trim0x(%s) = %s, want %s", tt.input, result, tt.expected)
		}
	}
}

