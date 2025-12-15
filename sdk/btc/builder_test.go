package btc

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/btcutil/psbt"
	"github.com/btcsuite/btcd/wire"
)

// Test public key (compressed, 33 bytes)
var testPubKey, _ = hex.DecodeString("02a1633cafcc01ebfb6d78e39f687a1f0995c62fc95f51ead10a02ee0be551b5dc")

func TestNewBuilder(t *testing.T) {
	builder := Mainnet()
	if builder == nil {
		t.Fatal("expected non-nil builder")
	}
	if builder.DustLimit != 546 {
		t.Errorf("expected dust limit 546, got %d", builder.DustLimit)
	}
}

func TestBuild_Send(t *testing.T) {
	builder := Testnet()

	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 100000, PkScript: p2wpkhScript},
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000002", Index: 1, Value: 50000, PkScript: p2wpkhScript},
	}

	// Create outputs: recipient + change (change with Value=0)
	recipientScript, _ := hex.DecodeString("0014751e76e8199196d454941c45d1b3a323f1433bd6") // tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx
	changeScript, _ := hex.DecodeString("0014" + "1234567890abcdef1234567890abcdef12345678")

	outputs := []*wire.TxOut{
		{Value: 50000, PkScript: recipientScript},
		{Value: 0, PkScript: changeScript}, // change output
	}

	result, err := builder.Build(utxos, outputs, 1, 10, testPubKey)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	if result == nil {
		t.Fatal("expected non-nil result")
	}

	// Verify PSBT packet exists
	if result.Packet == nil {
		t.Error("expected non-nil PSBT packet")
	}

	// Verify fee is reasonable
	if result.Fee < 500 || result.Fee > 3000 {
		t.Errorf("fee seems unreasonable: %d sats", result.Fee)
	}

	// Verify change was set
	if result.ChangeIndex != 1 {
		t.Errorf("expected change index 1, got %d", result.ChangeIndex)
	}

	if result.ChangeAmount <= 0 {
		t.Error("expected positive change amount")
	}

	// Verify selected UTXOs
	if len(result.SelectedUTXOs) == 0 {
		t.Error("expected at least one selected UTXO")
	}
}

func TestBuild_InsufficientFunds(t *testing.T) {
	builder := Testnet()

	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 1000, PkScript: p2wpkhScript},
	}

	recipientScript, _ := hex.DecodeString("0014751e76e8199196d454941c45d1b3a323f1433bd6")
	changeScript, _ := hex.DecodeString("0014" + "1234567890abcdef1234567890abcdef12345678")

	outputs := []*wire.TxOut{
		{Value: 50000, PkScript: recipientScript},
		{Value: 0, PkScript: changeScript},
	}

	_, err := builder.Build(utxos, outputs, 1, 10, testPubKey)
	if err == nil {
		t.Fatal("expected error for insufficient funds")
	}
}

func TestBuild_NoUTXOs(t *testing.T) {
	builder := Testnet()

	recipientScript, _ := hex.DecodeString("0014751e76e8199196d454941c45d1b3a323f1433bd6")
	changeScript, _ := hex.DecodeString("0014" + "1234567890abcdef1234567890abcdef12345678")

	outputs := []*wire.TxOut{
		{Value: 50000, PkScript: recipientScript},
		{Value: 0, PkScript: changeScript},
	}

	_, err := builder.Build([]UTXO{}, outputs, 1, 10, testPubKey)
	if err == nil {
		t.Fatal("expected error for no UTXOs")
	}
}

func TestBuild_InvalidPubKey(t *testing.T) {
	builder := Testnet()

	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 100000, PkScript: p2wpkhScript},
	}

	recipientScript, _ := hex.DecodeString("0014751e76e8199196d454941c45d1b3a323f1433bd6")
	changeScript, _ := hex.DecodeString("0014" + "1234567890abcdef1234567890abcdef12345678")

	outputs := []*wire.TxOut{
		{Value: 50000, PkScript: recipientScript},
		{Value: 0, PkScript: changeScript},
	}

	_, err := builder.Build(utxos, outputs, 1, 10, []byte{0x01, 0x02})
	if err == nil {
		t.Fatal("expected error for invalid pubkey")
	}
}

func TestBuild_Swap(t *testing.T) {
	builder := Testnet()

	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 200000, PkScript: p2wpkhScript},
	}

	// Swap outputs: vault address + OP_RETURN memo + change
	vaultScript, _ := hex.DecodeString("0014751e76e8199196d454941c45d1b3a323f1433bd6")
	// OP_RETURN script: 0x6a (OP_RETURN) + 0x24 (push 36 bytes) + memo
	opReturnScript, _ := hex.DecodeString("6a24" + hex.EncodeToString([]byte("=:ETH.ETH:0x1234567890abcdef:0/1/0")))
	changeScript, _ := hex.DecodeString("0014" + "1234567890abcdef1234567890abcdef12345678")

	outputs := []*wire.TxOut{
		{Value: 100000, PkScript: vaultScript},
		{Value: 0, PkScript: opReturnScript},
		{Value: 0, PkScript: changeScript}, // change output
	}

	result, err := builder.Build(utxos, outputs, 2, 15, testPubKey)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	// Should have 3 outputs: vault, memo (OP_RETURN), change
	if len(result.Packet.UnsignedTx.TxOut) != 3 {
		t.Errorf("expected 3 outputs, got %d", len(result.Packet.UnsignedTx.TxOut))
	}

	// Verify OP_RETURN output exists
	hasOpReturn := false
	for _, out := range result.Packet.UnsignedTx.TxOut {
		if len(out.PkScript) > 0 && out.PkScript[0] == 0x6a { // OP_RETURN
			hasOpReturn = true
			break
		}
	}
	if !hasOpReturn {
		t.Error("expected OP_RETURN output for memo")
	}
}

func TestSelectUTXOs_LargestFirst(t *testing.T) {
	builder := Testnet()

	utxos := []UTXO{
		{TxHash: "tx1", Index: 0, Value: 10000},
		{TxHash: "tx2", Index: 0, Value: 50000},
		{TxHash: "tx3", Index: 0, Value: 30000},
	}

	selected, total, err := builder.selectUTXOs(utxos, 40000, 10, 2, 0)
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

func TestEstimateTxVBytes(t *testing.T) {
	tests := []struct {
		name       string
		numInputs  int
		numOutputs int
		minSize    int
		maxSize    int
	}{
		{"1-in-1-out", 1, 1, 100, 150},
		{"1-in-2-out", 1, 2, 130, 180},
		{"2-in-2-out", 2, 2, 190, 250},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := EstimateTxVBytes(tt.numInputs, tt.numOutputs, 0)
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

func TestPopulatePSBTMetadata(t *testing.T) {
	builder := Testnet()

	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 100000, PkScript: p2wpkhScript},
	}

	recipientScript, _ := hex.DecodeString("0014751e76e8199196d454941c45d1b3a323f1433bd6")
	changeScript, _ := hex.DecodeString("0014" + "1234567890abcdef1234567890abcdef12345678")

	outputs := []*wire.TxOut{
		{Value: 50000, PkScript: recipientScript},
		{Value: 0, PkScript: changeScript},
	}

	result, err := builder.Build(utxos, outputs, 1, 10, testPubKey)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	// Create a mock fetcher (not needed since we have PkScript)
	mockFetcher := &mockPrevTxFetcher{}

	err = PopulatePSBTMetadata(result, mockFetcher)
	if err != nil {
		t.Fatalf("PopulatePSBTMetadata failed: %v", err)
	}

	// Verify WitnessUtxo was populated
	if result.Packet.Inputs[0].WitnessUtxo == nil {
		t.Error("expected WitnessUtxo to be populated")
	}
}

type mockPrevTxFetcher struct{}

func (m *mockPrevTxFetcher) GetRawTransaction(txHash string) ([]byte, error) {
	return nil, nil
}

func TestCreatePSBT(t *testing.T) {
	builder := Testnet()

	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 100000, PkScript: p2wpkhScript},
	}

	recipientScript, _ := hex.DecodeString("0014751e76e8199196d454941c45d1b3a323f1433bd6")
	changeScript, _ := hex.DecodeString("0014" + "1234567890abcdef1234567890abcdef12345678")

	outputs := []*wire.TxOut{
		{Value: 50000, PkScript: recipientScript},
		{Value: 0, PkScript: changeScript},
	}

	result, err := builder.Build(utxos, outputs, 1, 10, testPubKey)
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	// Serialize and re-parse PSBT
	var buf bytes.Buffer
	if err := result.Packet.Serialize(&buf); err != nil {
		t.Fatalf("failed to serialize PSBT: %v", err)
	}

	_, err = psbt.NewFromRawBytes(&buf, false)
	if err != nil {
		t.Fatalf("failed to parse serialized PSBT: %v", err)
	}
}

func TestBuild_InvalidChangeIndex(t *testing.T) {
	builder := Testnet()

	p2wpkhScript, _ := hex.DecodeString("0014" + "89abcdefabbaabbaabbaabbaabbaabbaabbaabba")

	utxos := []UTXO{
		{TxHash: "0000000000000000000000000000000000000000000000000000000000000001", Index: 0, Value: 100000, PkScript: p2wpkhScript},
	}

	recipientScript, _ := hex.DecodeString("0014751e76e8199196d454941c45d1b3a323f1433bd6")

	outputs := []*wire.TxOut{
		{Value: 50000, PkScript: recipientScript},
	}

	// Invalid change index (out of bounds)
	_, err := builder.Build(utxos, outputs, 5, 10, testPubKey)
	if err == nil {
		t.Fatal("expected error for invalid change index")
	}
}
