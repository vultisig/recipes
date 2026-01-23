package btc

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/vultisig/mobile-tss-lib/tss"
	rsdk "github.com/vultisig/recipes/sdk"
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
	testPrevTxHash    = "9b1a2f3e4d5c6b7a8f9e1f0b4b4d5b2b4b8d3e0c8050b5b0e3f7650145cd"
	testOutputValue   = 6300000 // 0.063 BTC
	testInputValue    = 1000000 // 0.01 BTC
	testOutputScript  = "0014abcdefabcdefabcdefabcdefabcdefabcdefabcd"
	testWitnessScript = "0014fedcbafedcbafedcbafedcbafedcbafedcbafedc"
	testBroadcastHash = "9b1a2f3e4d5c6b7a8f9e1f0b4b4d5b2b4b8d3e0c8050b5b0e3f765014f5cd"
)

// Mock RPC client for testing
type mockRPCClient struct {
	broadcastTx    *wire.MsgTx
	broadcastError error
}

func (m *mockRPCClient) SendRawTransaction(tx *wire.MsgTx, allowHighFees bool) (*chainhash.Hash, error) {
	m.broadcastTx = tx
	if m.broadcastError != nil {
		return nil, m.broadcastError
	}
	txHash := tx.TxHash()
	return &txHash, nil
}

// createTestPSBT creates a minimal valid PSBT for testing
func createTestPSBT() (*psbt.Packet, error) {
	// Create unsigned transaction
	tx := wire.NewMsgTx(2)

	// Add input (previous transaction output)
	prevHash, _ := chainhash.NewHashFromStr(testPrevTxHash)
	prevOut := wire.NewOutPoint(prevHash, 0)
	txIn := wire.NewTxIn(prevOut, nil, nil)
	tx.AddTxIn(txIn)

	// Add output
	scriptHash, _ := hex.DecodeString(testOutputScript)
	txOut := wire.NewTxOut(testOutputValue, scriptHash)
	tx.AddTxOut(txOut)

	// Create PSBT packet
	psbtPacket, err := psbt.NewFromUnsignedTx(tx)
	if err != nil {
		return nil, err
	}

	// Add witness UTXO for input 0
	witnessUtxo := &wire.TxOut{
		Value:    testInputValue,
		PkScript: []byte{0x00, 0x14}, // P2WPKH prefix
	}
	witnessScript, err := hex.DecodeString(testWitnessScript)
	if err != nil {
		return nil, err
	}
	witnessUtxo.PkScript = witnessScript
	psbtPacket.Inputs[0].WitnessUtxo = witnessUtxo

	// Add BIP32 derivation for the test public key
	pubKeyBytes, err := hex.DecodeString(testPubKeyHex)
	if err != nil {
		return nil, err
	}
	derivation := &psbt.Bip32Derivation{
		PubKey:               pubKeyBytes,
		MasterKeyFingerprint: 0x12345678,
		Bip32Path:            []uint32{0x80000000, 1, 0}, // m/0'/1/0
	}
	psbtPacket.Inputs[0].Bip32Derivation = append(psbtPacket.Inputs[0].Bip32Derivation, derivation)

	return psbtPacket, nil
}

func TestSDK_extractPubkeyForInput(t *testing.T) {
	sdk := &SDK{}

	// Test with valid BIP32 derivation
	pubKeyBytes, err := hex.DecodeString(testPubKeyHex)
	if err != nil {
		t.Fatalf("invalid pubkey hex: %v", err)
	}
	derivation := &psbt.Bip32Derivation{
		PubKey: pubKeyBytes,
	}

	input := &psbt.PInput{
		Bip32Derivation: []*psbt.Bip32Derivation{derivation},
	}

	pubkey, err := sdk.extractPubkeyForInput(input)
	if err != nil {
		t.Fatalf("extractPubkeyForInput failed: %v", err)
	}

	if !bytes.Equal(pubkey, pubKeyBytes) {
		t.Errorf("Expected pubkey %x, got %x", pubKeyBytes, pubkey)
	}

	// Test with no derivation info
	emptyInput := &psbt.PInput{}
	_, err = sdk.extractPubkeyForInput(emptyInput)
	if err == nil {
		t.Error("Expected error for input with no public key")
	}
}

func TestSDK_calculateInputSignatureHash(t *testing.T) {
	sdk := &SDK{}

	// Create test PSBT with witness input
	psbtPacket, err := createTestPSBT()
	if err != nil {
		t.Fatalf("Failed to create test PSBT: %v", err)
	}

	// Test calculateInputSignatureHash for witness input
	sigHash, err := sdk.CalculateInputSignatureHash(psbtPacket, 0)
	if err != nil {
		t.Fatalf("calculateInputSignatureHash failed: %v", err)
	}

	if len(sigHash) != 32 {
		t.Errorf("Expected signature hash length 32, got %d", len(sigHash))
	}

	// Test with invalid input index
	_, err = sdk.CalculateInputSignatureHash(psbtPacket, 1)
	if err == nil {
		t.Error("Expected error for invalid input index")
	}

	// Test with input missing UTXO data
	tx := wire.NewMsgTx(2)
	tx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(&chainhash.Hash{}, 0), nil, nil))
	emptyPSBT, err := psbt.NewFromUnsignedTx(tx)
	if err != nil {
		t.Fatalf("Failed to create empty PSBT: %v", err)
	}
	_, err = sdk.CalculateInputSignatureHash(emptyPSBT, 0)
	if err == nil {
		t.Error("Expected error for input with no UTXO data")
	}
}

func TestSDK_deriveKeyFromMessage(t *testing.T) {
	sdk := &SDK{}

	testMessage := []byte("test message hash")
	derivedKey := sdk.deriveKeyFromMessage(testMessage)

	// Verify the key derivation process manually
	encodedMsg := base64.StdEncoding.EncodeToString(testMessage)
	decodedMsg, _ := base64.StdEncoding.DecodeString(encodedMsg)
	hash := sha256.Sum256(decodedMsg)
	expectedKey := base64.StdEncoding.EncodeToString(hash[:])

	if derivedKey != expectedKey {
		t.Errorf("Key derivation mismatch. Expected %s, got %s", expectedKey, derivedKey)
	}

	// Test that same message produces same key
	derivedKey2 := sdk.deriveKeyFromMessage(testMessage)
	if derivedKey != derivedKey2 {
		t.Error("Same message should produce same derived key")
	}

	// Test that different messages produce different keys
	differentMessage := []byte("different message")
	differentKey := sdk.deriveKeyFromMessage(differentMessage)
	if derivedKey == differentKey {
		t.Error("Different messages should produce different keys")
	}
}

func TestSDK_Sign_WithSignatureHashes(t *testing.T) {
	sdk := &SDK{
		rpcClient: &mockRPCClient{},
	}

	// Create test PSBT
	psbtPacket, err := createTestPSBT()
	if err != nil {
		t.Fatalf("Failed to create test PSBT: %v", err)
	}

	// Verify PSBT starts with no signatures
	if len(psbtPacket.Inputs[0].PartialSigs) != 0 {
		t.Fatal("PSBT should start with no partial signatures")
	}

	// Calculate the signature hash for input 0
	sigHash, err := sdk.CalculateInputSignatureHash(psbtPacket, 0)
	if err != nil {
		t.Fatalf("Failed to calculate signature hash: %v", err)
	}
	t.Logf("✓ Signature hash calculated: %x", sigHash)

	// Derive the key that would be used in the signatures map
	derivedKey := sdk.deriveKeyFromMessage(sigHash)
	t.Logf("✓ Derived key: %s", derivedKey)

	// Serialize PSBT to bytes
	var buf bytes.Buffer
	err = psbtPacket.Serialize(&buf)
	if err != nil {
		t.Fatalf("Failed to serialize PSBT: %v", err)
	}
	psbtBytes := buf.Bytes()

	// Create signatures map with the correct derived key
	signatures := map[string]tss.KeysignResponse{
		derivedKey: {
			Msg:          base64.StdEncoding.EncodeToString(sigHash),
			DerSignature: mockDerSignature,
			R:            mockRValue,
			S:            mockSValue,
			RecoveryID:   mockRecoveryID,
		},
	}
	t.Logf("✓ Signatures map created with key: %s", derivedKey)

	// Test signing - should succeed
	signedTxBytes, err := sdk.Sign(psbtBytes, signatures)
	if err != nil {
		t.Fatalf("Signing failed: %v", err)
	}

	t.Logf("✓ Transaction signed successfully, %d bytes", len(signedTxBytes))

	// Verify the signed transaction can be parsed
	var signedTx wire.MsgTx
	err = signedTx.Deserialize(bytes.NewReader(signedTxBytes))
	if err != nil {
		t.Fatalf("Failed to parse signed transaction: %v", err)
	}

	t.Logf("✓ Signed transaction hash: %s", signedTx.TxHash().String())
	t.Logf("✓ PSBT signing test completed successfully")
}

func TestSDK_Sign_MissingSignature(t *testing.T) {
	sdk := &SDK{
		rpcClient: &mockRPCClient{},
	}

	// Create test PSBT
	psbtPacket, err := createTestPSBT()
	if err != nil {
		t.Fatalf("Failed to create test PSBT: %v", err)
	}

	// Serialize PSBT to bytes
	var buf bytes.Buffer
	err = psbtPacket.Serialize(&buf)
	if err != nil {
		t.Fatalf("Failed to serialize PSBT: %v", err)
	}
	psbtBytes := buf.Bytes()

	// Create empty signatures map
	signatures := map[string]tss.KeysignResponse{}

	// Test signing - should fail with missing signature error
	_, err = sdk.Sign(psbtBytes, signatures)
	if err == nil {
		t.Error("Expected error for missing signature")
	}

	if !bytes.Contains([]byte(err.Error()), []byte("missing signature")) {
		t.Errorf("Expected 'missing signature' error, got: %v", err)
	}
}

// TestPSBTStructure validates that our PSBT creation and parsing works correctly
func TestPSBTStructure(t *testing.T) {
	// Create test PSBT
	psbtPacket, err := createTestPSBT()
	if err != nil {
		t.Fatalf("Failed to create test PSBT: %v", err)
	}

	// Validate basic structure
	if psbtPacket.UnsignedTx == nil {
		t.Fatal("PSBT missing unsigned transaction")
	}

	if len(psbtPacket.Inputs) != 1 {
		t.Fatalf("Expected 1 input, got %d", len(psbtPacket.Inputs))
	}

	if len(psbtPacket.UnsignedTx.TxOut) != 1 {
		t.Fatalf("Expected 1 output, got %d", len(psbtPacket.UnsignedTx.TxOut))
	}

	// Validate input has required fields
	input := psbtPacket.Inputs[0]
	if input.WitnessUtxo == nil {
		t.Error("Input missing witness UTXO")
	}

	if len(input.Bip32Derivation) == 0 {
		t.Error("Input missing BIP32 derivation info")
	}

	// Test serialization and deserialization
	var buf bytes.Buffer
	err = psbtPacket.Serialize(&buf)
	if err != nil {
		t.Fatalf("Failed to serialize PSBT: %v", err)
	}

	// Parse it back
	parsedPacket, err := psbt.NewFromRawBytes(bytes.NewReader(buf.Bytes()), false)
	if err != nil {
		t.Fatalf("Failed to parse PSBT: %v", err)
	}

	// Validate parsed structure matches original
	if parsedPacket.UnsignedTx.TxHash() != psbtPacket.UnsignedTx.TxHash() {
		t.Error("PSBT transaction hash mismatch after serialization roundtrip")
	}
}

// TestSignatureApplication tests that signatures are properly applied to PSBT
func TestSignatureApplication(t *testing.T) {
	// Create PSBT
	psbtPacket, err := createTestPSBT()
	if err != nil {
		t.Fatalf("Failed to create PSBT: %v", err)
	}

	// Verify input initially has no signatures
	if len(psbtPacket.Inputs[0].PartialSigs) != 0 {
		t.Error("Input should start with no signatures")
	}

	// Simulate applying a signature manually (bypass signature hash calculation)
	pubKey, err := hex.DecodeString(testPubKeyHex)
	if err != nil {
		t.Fatalf("invalid pubkey hex: %v", err)
	}
	mockSig, err := hex.DecodeString(mockDerSignature)
	if err != nil {
		t.Fatalf("invalid DER signature hex: %v", err)
	}
	fullSig := append(mockSig, 0x01) // SIGHASH_ALL

	partialSig := &psbt.PartialSig{
		PubKey:    pubKey,
		Signature: fullSig,
	}
	psbtPacket.Inputs[0].PartialSigs = append(psbtPacket.Inputs[0].PartialSigs, partialSig)

	// Verify signature was added
	if len(psbtPacket.Inputs[0].PartialSigs) != 1 {
		t.Error("Signature was not added to input")
	}

	if !bytes.Equal(psbtPacket.Inputs[0].PartialSigs[0].PubKey, pubKey) {
		t.Error("Public key mismatch in partial signature")
	}

	if !bytes.Equal(psbtPacket.Inputs[0].PartialSigs[0].Signature, fullSig) {
		t.Error("Signature mismatch in partial signature")
	}
}

// createTestPSBTWithMultipleInputs creates a PSBT with multiple inputs for testing
func createTestPSBTWithMultipleInputs(t *testing.T, numInputs int) (*psbt.Packet, error) {
	tx := wire.NewMsgTx(2)

	// Add multiple inputs
	for i := 0; i < numInputs; i++ {
		prevHash, _ := chainhash.NewHashFromStr(testPrevTxHash)
		prevOut := wire.NewOutPoint(prevHash, uint32(i))
		txIn := wire.NewTxIn(prevOut, nil, nil)
		tx.AddTxIn(txIn)
	}

	// Add output
	scriptHash, _ := hex.DecodeString(testOutputScript)
	txOut := wire.NewTxOut(testOutputValue, scriptHash)
	tx.AddTxOut(txOut)

	// Create PSBT packet
	psbtPacket, err := psbt.NewFromUnsignedTx(tx)
	if err != nil {
		return nil, err
	}

	// Add witness UTXO and BIP32 derivation for each input
	for i := 0; i < numInputs; i++ {
		witnessScript, _ := hex.DecodeString(testWitnessScript)
		witnessUtxo := &wire.TxOut{
			Value:    testInputValue,
			PkScript: witnessScript,
		}
		psbtPacket.Inputs[i].WitnessUtxo = witnessUtxo

		pubKeyBytes, _ := hex.DecodeString(testPubKeyHex)
		derivation := &psbt.Bip32Derivation{
			PubKey:               pubKeyBytes,
			MasterKeyFingerprint: 0x12345678,
			Bip32Path:            []uint32{0x80000000, 1, uint32(i)},
		}
		psbtPacket.Inputs[i].Bip32Derivation = append(psbtPacket.Inputs[i].Bip32Derivation, derivation)
	}

	return psbtPacket, nil
}

func TestDeriveSigningHashes_ValidPSBT_SingleInput(t *testing.T) {
	sdk := &SDK{}

	// Create test PSBT with single input
	psbtPacket, err := createTestPSBT()
	if err != nil {
		t.Fatalf("Failed to create test PSBT: %v", err)
	}

	// Serialize PSBT to bytes
	var buf bytes.Buffer
	err = psbtPacket.Serialize(&buf)
	if err != nil {
		t.Fatalf("Failed to serialize PSBT: %v", err)
	}

	// Derive signing hashes
	hashes, err := sdk.DeriveSigningHashes(buf.Bytes(), rsdk.DeriveOptions{})
	if err != nil {
		t.Fatalf("DeriveSigningHashes failed: %v", err)
	}

	// Should return 1 hash for 1 input
	if len(hashes) != 1 {
		t.Fatalf("Expected 1 hash, got %d", len(hashes))
	}

	// Hash should be 32 bytes (SHA256 output)
	if len(hashes[0].Hash) != 32 {
		t.Errorf("Expected hash length 32, got %d", len(hashes[0].Hash))
	}

	// Message should be 32 bytes (signature hash)
	if len(hashes[0].Message) != 32 {
		t.Errorf("Expected message length 32, got %d", len(hashes[0].Message))
	}

	// Hash should be SHA256 of Message
	expectedHash := sha256.Sum256(hashes[0].Message)
	if !bytes.Equal(hashes[0].Hash, expectedHash[:]) {
		t.Error("Hash should be SHA256 of Message")
	}
}

func TestDeriveSigningHashes_ValidPSBT_MultipleInputs(t *testing.T) {
	sdk := &SDK{}

	testCases := []int{2, 3, 5}

	for _, numInputs := range testCases {
		t.Run(fmt.Sprintf("%d_inputs", numInputs), func(t *testing.T) {
			psbtPacket, err := createTestPSBTWithMultipleInputs(t, numInputs)
			if err != nil {
				t.Fatalf("Failed to create test PSBT: %v", err)
			}

			var buf bytes.Buffer
			err = psbtPacket.Serialize(&buf)
			if err != nil {
				t.Fatalf("Failed to serialize PSBT: %v", err)
			}

			hashes, err := sdk.DeriveSigningHashes(buf.Bytes(), rsdk.DeriveOptions{})
			if err != nil {
				t.Fatalf("DeriveSigningHashes failed: %v", err)
			}

			// Should return one hash per input
			if len(hashes) != numInputs {
				t.Fatalf("Expected %d hashes, got %d", numInputs, len(hashes))
			}

			// Each hash should be unique and properly formatted
			seen := make(map[string]bool)
			for i, h := range hashes {
				if len(h.Hash) != 32 {
					t.Errorf("Input %d: expected hash length 32, got %d", i, len(h.Hash))
				}
				if len(h.Message) != 32 {
					t.Errorf("Input %d: expected message length 32, got %d", i, len(h.Message))
				}

				hashKey := base64.StdEncoding.EncodeToString(h.Hash)
				if seen[hashKey] {
					t.Errorf("Input %d: duplicate hash detected", i)
				}
				seen[hashKey] = true
			}
		})
	}
}

func TestDeriveSigningHashes_InvalidPSBT(t *testing.T) {
	sdk := &SDK{}

	testCases := []struct {
		name        string
		input       []byte
		errContains string
	}{
		{
			name:        "random_bytes",
			input:       []byte{0x01, 0x02, 0x03, 0x04},
			errContains: "psbt",
		},
		{
			name:        "empty_bytes",
			input:       []byte{},
			errContains: "psbt",
		},
		{
			name:        "truncated_psbt_magic",
			input:       []byte{0x70, 0x73, 0x62, 0x74}, // "psbt" without 0xff
			errContains: "psbt",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := sdk.DeriveSigningHashes(tc.input, rsdk.DeriveOptions{})
			if err == nil {
				t.Error("Expected error for invalid PSBT")
			}
		})
	}
}

func TestDeriveSigningHashes_PSBTWithoutUnsignedTx(t *testing.T) {
	sdk := &SDK{}

	// Create a minimal PSBT with just the magic bytes but no content
	// PSBT magic: 0x70736274ff followed by minimal global section
	invalidPSBT := []byte{0x70, 0x73, 0x62, 0x74, 0xff, 0x00}

	_, err := sdk.DeriveSigningHashes(invalidPSBT, rsdk.DeriveOptions{})
	if err == nil {
		t.Error("Expected error for PSBT without unsigned transaction")
	}
}

func TestDeriveSigningHashes_ConsistentResults(t *testing.T) {
	sdk := &SDK{}

	psbtPacket, err := createTestPSBT()
	if err != nil {
		t.Fatalf("Failed to create test PSBT: %v", err)
	}

	var buf bytes.Buffer
	err = psbtPacket.Serialize(&buf)
	if err != nil {
		t.Fatalf("Failed to serialize PSBT: %v", err)
	}
	psbtBytes := buf.Bytes()

	// Call DeriveSigningHashes multiple times
	hashes1, err := sdk.DeriveSigningHashes(psbtBytes, rsdk.DeriveOptions{})
	if err != nil {
		t.Fatalf("First call failed: %v", err)
	}

	hashes2, err := sdk.DeriveSigningHashes(psbtBytes, rsdk.DeriveOptions{})
	if err != nil {
		t.Fatalf("Second call failed: %v", err)
	}

	// Results should be identical
	if len(hashes1) != len(hashes2) {
		t.Fatalf("Inconsistent hash counts: %d vs %d", len(hashes1), len(hashes2))
	}

	for i := range hashes1 {
		if !bytes.Equal(hashes1[i].Hash, hashes2[i].Hash) {
			t.Errorf("Hash %d is not consistent across calls", i)
		}
		if !bytes.Equal(hashes1[i].Message, hashes2[i].Message) {
			t.Errorf("Message %d is not consistent across calls", i)
		}
	}
}
