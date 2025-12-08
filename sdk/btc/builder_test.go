package btc

import (
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/btcsuite/btcd/chaincfg"
)

// Test public key (compressed, 33 bytes)
var testPubKey, _ = hex.DecodeString("02a1633cafcc01ebfb6d78e39f687a1f0995c62fc95f51ead10a02ee0be551b5dc")

func TestNewTransactionBuilder(t *testing.T) {
	builder := NewTransactionBuilder(MainnetConfig)
	if builder == nil {
		t.Fatal("expected non-nil builder")
	}
	if builder.config.Network != &chaincfg.MainNetParams {
		t.Error("expected mainnet params")
	}
	if builder.config.DustLimit != 546 {
		t.Errorf("expected dust limit 546, got %d", builder.config.DustLimit)
	}
}

func TestBuildSendTransaction(t *testing.T) {
	builder := NewTransactionBuilder(TestnetConfig)

	// Create test UTXOs with PkScript for P2WPKH
	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 100000, PkScript: p2wpkhScript},
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000002", Index: 1, Value: 50000, PkScript: p2wpkhScript},
	}

	result, err := builder.BuildSendTransaction(SendParams{
		UTXOs:      utxos,
		FeeRate:    10,
		FromPubKey: testPubKey,
		ToAddress:  "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
		Amount:     50000,
	})

	if err != nil {
		t.Fatalf("BuildSendTransaction failed: %v", err)
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}

	// Verify PSBT is valid
	if len(result.PSBT) == 0 {
		t.Error("expected non-empty PSBT")
	}

	// Verify we can parse the PSBT
	_, err = psbt.NewFromRawBytes(newBytesReader(result.PSBT), false)
	if err != nil {
		t.Errorf("failed to parse PSBT: %v", err)
	}

	// Verify fee is reasonable (should be around 1100-1500 sats for 2-in-2-out at 10 sat/vbyte)
	if result.Fee < 500 || result.Fee > 3000 {
		t.Errorf("fee seems unreasonable: %d sats", result.Fee)
	}

	// Verify change exists
	if result.ChangeIndex == -1 {
		t.Error("expected change output")
	}

	if result.ChangeAmount <= 0 {
		t.Error("expected positive change amount")
	}

	// Verify selected UTXOs
	if len(result.SelectedUTXOs) == 0 {
		t.Error("expected at least one selected UTXO")
	}
}

func TestBuildTransaction_InsufficientFunds(t *testing.T) {
	builder := NewTransactionBuilder(TestnetConfig)

	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 1000, PkScript: p2wpkhScript},
	}

	_, err := builder.BuildSendTransaction(SendParams{
		UTXOs:      utxos,
		FeeRate:    10,
		FromPubKey: testPubKey,
		ToAddress:  "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
		Amount:     50000, // More than we have
	})

	if err == nil {
		t.Fatal("expected error for insufficient funds")
	}
}

func TestBuildTransaction_NoUTXOs(t *testing.T) {
	builder := NewTransactionBuilder(TestnetConfig)

	_, err := builder.BuildSendTransaction(SendParams{
		UTXOs:      []UTXO{},
		FeeRate:    10,
		FromPubKey: testPubKey,
		ToAddress:  "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
		Amount:     50000,
	})

	if err == nil {
		t.Fatal("expected error for no UTXOs")
	}
}

func TestBuildTransaction_InvalidPubKey(t *testing.T) {
	builder := NewTransactionBuilder(TestnetConfig)

	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 100000, PkScript: p2wpkhScript},
	}

	_, err := builder.BuildSendTransaction(SendParams{
		UTXOs:      utxos,
		FeeRate:    10,
		FromPubKey: []byte{0x01, 0x02}, // Invalid: too short
		ToAddress:  "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
		Amount:     50000,
	})

	if err == nil {
		t.Fatal("expected error for invalid pubkey")
	}
}

func TestBuildSwapTransaction(t *testing.T) {
	builder := NewTransactionBuilder(TestnetConfig)

	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 200000, PkScript: p2wpkhScript},
	}

	result, err := builder.BuildSwapTransaction(SwapParams{
		UTXOs:        utxos,
		FeeRate:      15,
		FromPubKey:   testPubKey,
		VaultAddress: "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
		Amount:       100000,
		Memo:         "=:ETH.ETH:0x1234567890abcdef:0/1/0",
	})

	if err != nil {
		t.Fatalf("BuildSwapTransaction failed: %v", err)
	}

	// Verify PSBT is valid
	pkt, err := psbt.NewFromRawBytes(newBytesReader(result.PSBT), false)
	if err != nil {
		t.Fatalf("failed to parse PSBT: %v", err)
	}

	// Should have 3 outputs: vault, memo (OP_RETURN), change
	if len(pkt.UnsignedTx.TxOut) != 3 {
		t.Errorf("expected 3 outputs, got %d", len(pkt.UnsignedTx.TxOut))
	}

	// Verify OP_RETURN output exists (should be output 1)
	hasOpReturn := false
	for _, out := range pkt.UnsignedTx.TxOut {
		if len(out.PkScript) > 0 && out.PkScript[0] == 0x6a { // OP_RETURN
			hasOpReturn = true
			break
		}
	}
	if !hasOpReturn {
		t.Error("expected OP_RETURN output for memo")
	}
}

func TestBuildSwapTransaction_NoMemo(t *testing.T) {
	builder := NewTransactionBuilder(TestnetConfig)

	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 200000, PkScript: p2wpkhScript},
	}

	_, err := builder.BuildSwapTransaction(SwapParams{
		UTXOs:        utxos,
		FeeRate:      15,
		FromPubKey:   testPubKey,
		VaultAddress: "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
		Amount:       100000,
		Memo:         "", // Empty memo
	})

	if err == nil {
		t.Fatal("expected error for empty memo")
	}
}

func TestSelectUTXOs_LargestFirst(t *testing.T) {
	builder := NewTransactionBuilder(TestnetConfig)

	utxos := []UTXO{
		{TxHash: "tx1", Index: 0, Value: 10000},
		{TxHash: "tx2", Index: 0, Value: 50000},
		{TxHash: "tx3", Index: 0, Value: 30000},
	}

	selected, total, err := builder.selectUTXOs(utxos, 40000, 10, 2, false)
	if err != nil {
		t.Fatalf("selectUTXOs failed: %v", err)
	}

	// Should select the largest UTXO first (50000)
	if len(selected) != 1 {
		t.Errorf("expected 1 UTXO selected, got %d", len(selected))
	}

	if selected[0].Value != 50000 {
		t.Errorf("expected largest UTXO (50000) to be selected, got %d", selected[0].Value)
	}

	if total != 50000 {
		t.Errorf("expected total 50000, got %d", total)
	}
}

func TestSelectUTXOs_SelectAll(t *testing.T) {
	builder := NewTransactionBuilder(TestnetConfig)

	utxos := []UTXO{
		{TxHash: "tx1", Index: 0, Value: 10000},
		{TxHash: "tx2", Index: 0, Value: 50000},
		{TxHash: "tx3", Index: 0, Value: 30000},
	}

	selected, total, err := builder.selectUTXOs(utxos, 10000, 10, 2, true)
	if err != nil {
		t.Fatalf("selectUTXOs failed: %v", err)
	}

	if len(selected) != 3 {
		t.Errorf("expected all 3 UTXOs selected, got %d", len(selected))
	}

	if total != 90000 {
		t.Errorf("expected total 90000, got %d", total)
	}
}

func TestEstimateTxSize(t *testing.T) {
	tests := []struct {
		name       string
		numInputs  int
		numOutputs int
		isSegwit   bool
		minSize    int
		maxSize    int
	}{
		{"1-in-1-out segwit", 1, 1, true, 100, 150},
		{"1-in-2-out segwit", 1, 2, true, 130, 180},
		{"2-in-2-out segwit", 2, 2, true, 190, 250},
		{"1-in-1-out legacy", 1, 1, false, 180, 230},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := EstimateTxSize(tt.numInputs, tt.numOutputs, tt.isSegwit)
			if size < tt.minSize || size > tt.maxSize {
				t.Errorf("size %d outside expected range [%d, %d]", size, tt.minSize, tt.maxSize)
			}
		})
	}
}

func TestCalculateFee(t *testing.T) {
	tests := []struct {
		vbytes      int
		satsPerByte uint64
		expected    uint64
	}{
		{100, 1, 100},
		{100, 10, 1000},
		{250, 15, 3750},
	}

	for _, tt := range tests {
		fee := CalculateFee(tt.vbytes, tt.satsPerByte)
		if fee != tt.expected {
			t.Errorf("CalculateFee(%d, %d) = %d, expected %d", tt.vbytes, tt.satsPerByte, fee, tt.expected)
		}
	}
}

func TestPubKeyToP2WPKHAddress(t *testing.T) {
	// Test with known public key and expected address
	addr, err := PubKeyToP2WPKHAddress(testPubKey, &chaincfg.MainNetParams)
	if err != nil {
		t.Fatalf("PubKeyToP2WPKHAddress failed: %v", err)
	}

	// Should be a valid bech32 address starting with bc1
	if len(addr) == 0 {
		t.Error("expected non-empty address")
	}
	if addr[0:3] != "bc1" {
		t.Errorf("expected mainnet bech32 address starting with bc1, got %s", addr)
	}
}

func TestCreateOPReturnOutput(t *testing.T) {
	data := []byte("test memo data")
	out, err := CreateOPReturnOutput(data)
	if err != nil {
		t.Fatalf("CreateOPReturnOutput failed: %v", err)
	}

	if out.Value != 0 {
		t.Errorf("expected 0 value for OP_RETURN, got %d", out.Value)
	}

	// Check script starts with OP_RETURN (0x6a)
	if len(out.PkScript) == 0 || out.PkScript[0] != 0x6a {
		t.Error("expected OP_RETURN script")
	}
}

func TestCreateOPReturnOutput_TooLong(t *testing.T) {
	data := make([]byte, 81) // > 80 bytes
	_, err := CreateOPReturnOutput(data)
	if err == nil {
		t.Fatal("expected error for data > 80 bytes")
	}
}

func TestIsWitnessOutput(t *testing.T) {
	tests := []struct {
		name     string
		script   string
		expected bool
	}{
		{"P2WPKH", "0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba", true},
		{"P2WSH", "0020" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabbaabba", true},
		{"P2PKH", "76a914" + "89abcdefabbaabbaabbaabbaabbaabba" + "88ac", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			script, _ := hex.DecodeString(tt.script)
			result := IsWitnessOutput(script)
			if result != tt.expected {
				t.Errorf("IsWitnessOutput() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// Helper to create bytes.Reader
func newBytesReader(b []byte) *bytesReader {
	return &bytesReader{data: b, pos: 0}
}

type bytesReader struct {
	data []byte
	pos  int
}

func (r *bytesReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, nil
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}
