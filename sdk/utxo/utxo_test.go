package utxo

import (
	"testing"
)

func TestChainParams_GetInputSize(t *testing.T) {
	tests := []struct {
		name      string
		params    ChainParams
		inputType string
		want      int
	}{
		{"Bitcoin P2WPKH", BitcoinParams, "p2wpkh", 68},
		{"Bitcoin P2PKH", BitcoinParams, "p2pkh", 148},
		{"Bitcoin P2TR", BitcoinParams, "p2tr", 58},
		{"Bitcoin default", BitcoinParams, "unknown", 68}, // Defaults to P2WPKH for SegWit chains
		{"Dogecoin P2PKH", DogecoinParams, "p2pkh", 148},
		{"Dogecoin default", DogecoinParams, "unknown", 148}, // Defaults to P2PKH for non-SegWit
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.GetInputSize(tt.inputType)
			if got != tt.want {
				t.Errorf("GetInputSize(%s) = %d, want %d", tt.inputType, got, tt.want)
			}
		})
	}
}

func TestChainParams_GetOutputSize(t *testing.T) {
	tests := []struct {
		name       string
		params     ChainParams
		outputType string
		want       int
	}{
		{"Bitcoin P2WPKH", BitcoinParams, "p2wpkh", 31},
		{"Bitcoin P2PKH", BitcoinParams, "p2pkh", 34},
		{"Bitcoin OP_RETURN", BitcoinParams, "opreturn", 11},
		{"Bitcoin default", BitcoinParams, "unknown", 31}, // Defaults to P2WPKH for SegWit
		{"Dogecoin P2PKH", DogecoinParams, "p2pkh", 34},
		{"Dogecoin default", DogecoinParams, "unknown", 34}, // Defaults to P2PKH
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.GetOutputSize(tt.outputType)
			if got != tt.want {
				t.Errorf("GetOutputSize(%s) = %d, want %d", tt.outputType, got, tt.want)
			}
		})
	}
}

func TestSelect_LargestFirst(t *testing.T) {
	utxos := []UTXO{
		{TxHash: "tx1", Index: 0, Value: 10000},
		{TxHash: "tx2", Index: 0, Value: 50000},
		{TxHash: "tx3", Index: 0, Value: 30000},
	}

	result, err := Select(SelectionParams{
		UTXOs:        utxos,
		TargetAmount: 40000,
		FeeRate:      10,
		NumOutputs:   2,
		Strategy:     LargestFirst,
		ChainParams:  BitcoinParams,
	})

	if err != nil {
		t.Fatalf("Select failed: %v", err)
	}

	// Should select the largest UTXO first (50000)
	if len(result.Selected) != 1 {
		t.Errorf("expected 1 UTXO selected, got %d", len(result.Selected))
	}

	if result.Selected[0].Value != 50000 {
		t.Errorf("expected largest UTXO (50000), got %d", result.Selected[0].Value)
	}
}

func TestSelect_SmallestFirst(t *testing.T) {
	utxos := []UTXO{
		{TxHash: "tx1", Index: 0, Value: 10000},
		{TxHash: "tx2", Index: 0, Value: 50000},
		{TxHash: "tx3", Index: 0, Value: 30000},
	}

	result, err := Select(SelectionParams{
		UTXOs:        utxos,
		TargetAmount: 5000,
		FeeRate:      10,
		NumOutputs:   2,
		Strategy:     SmallestFirst,
		ChainParams:  BitcoinParams,
	})

	if err != nil {
		t.Fatalf("Select failed: %v", err)
	}

	// Should select the smallest UTXO first (10000)
	if len(result.Selected) != 1 {
		t.Errorf("expected 1 UTXO selected, got %d", len(result.Selected))
	}

	if result.Selected[0].Value != 10000 {
		t.Errorf("expected smallest UTXO (10000), got %d", result.Selected[0].Value)
	}
}

func TestSelect_SelectAll(t *testing.T) {
	utxos := []UTXO{
		{TxHash: "tx1", Index: 0, Value: 10000},
		{TxHash: "tx2", Index: 0, Value: 50000},
		{TxHash: "tx3", Index: 0, Value: 30000},
	}

	result, err := Select(SelectionParams{
		UTXOs:        utxos,
		TargetAmount: 10000,
		FeeRate:      10,
		NumOutputs:   2,
		Strategy:     SelectAll,
		ChainParams:  BitcoinParams,
	})

	if err != nil {
		t.Fatalf("Select failed: %v", err)
	}

	if len(result.Selected) != 3 {
		t.Errorf("expected all 3 UTXOs, got %d", len(result.Selected))
	}

	if result.TotalValue != 90000 {
		t.Errorf("expected total 90000, got %d", result.TotalValue)
	}
}

func TestSelect_InsufficientFunds(t *testing.T) {
	utxos := []UTXO{
		{TxHash: "tx1", Index: 0, Value: 1000},
	}

	_, err := Select(SelectionParams{
		UTXOs:        utxos,
		TargetAmount: 50000,
		FeeRate:      10,
		NumOutputs:   2,
		Strategy:     LargestFirst,
		ChainParams:  BitcoinParams,
	})

	if err == nil {
		t.Fatal("expected error for insufficient funds")
	}
}

func TestSelect_NoUTXOs(t *testing.T) {
	_, err := Select(SelectionParams{
		UTXOs:        []UTXO{},
		TargetAmount: 10000,
		FeeRate:      10,
		NumOutputs:   2,
		Strategy:     LargestFirst,
		ChainParams:  BitcoinParams,
	})

	if err == nil {
		t.Fatal("expected error for no UTXOs")
	}
}

func TestSelectWithDustHandling(t *testing.T) {
	// Create a scenario where change would be dust
	utxos := []UTXO{
		{TxHash: "tx1", Index: 0, Value: 10000},
	}

	// Target amount that leaves dust change (less than 546)
	result, err := SelectWithDustHandling(SelectionParams{
		UTXOs:        utxos,
		TargetAmount: 9000, // Leaves ~500 after fee (dust)
		FeeRate:      5,
		NumOutputs:   2,
		Strategy:     LargestFirst,
		ChainParams:  BitcoinParams,
	})

	if err != nil {
		t.Fatalf("SelectWithDustHandling failed: %v", err)
	}

	// Dust should be added to fee, change should be 0
	if result.Change != 0 && result.Change < BitcoinParams.DustLimit {
		t.Errorf("expected dust to be added to fee, but change is %d", result.Change)
	}
}

func TestEstimateTxSize(t *testing.T) {
	tests := []struct {
		name    string
		params  EstimateFeeParams
		minSize int
		maxSize int
	}{
		{
			name: "1-in-2-out Bitcoin SegWit",
			params: EstimateFeeParams{
				NumInputs:   1,
				NumOutputs:  2,
				ChainParams: BitcoinParams,
			},
			minSize: 100,
			maxSize: 150,
		},
		{
			name: "2-in-2-out Bitcoin SegWit",
			params: EstimateFeeParams{
				NumInputs:   2,
				NumOutputs:  2,
				ChainParams: BitcoinParams,
			},
			minSize: 160,
			maxSize: 220,
		},
		{
			name: "1-in-2-out Dogecoin (no SegWit)",
			params: EstimateFeeParams{
				NumInputs:   1,
				NumOutputs:  2,
				ChainParams: DogecoinParams,
			},
			minSize: 200,
			maxSize: 260,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := EstimateTxSize(tt.params)
			if size < tt.minSize || size > tt.maxSize {
				t.Errorf("size %d outside expected range [%d, %d]", size, tt.minSize, tt.maxSize)
			}
		})
	}
}

func TestCalculateFee(t *testing.T) {
	tests := []struct {
		size int
		rate uint64
		want uint64
	}{
		{100, 1, 100},
		{100, 10, 1000},
		{250, 15, 3750},
	}

	for _, tt := range tests {
		got := CalculateFee(tt.size, tt.rate)
		if got != tt.want {
			t.Errorf("CalculateFee(%d, %d) = %d, want %d", tt.size, tt.rate, got, tt.want)
		}
	}
}

func TestChainParamsPresets(t *testing.T) {
	// Verify all presets have required fields
	presets := []struct {
		name   string
		params ChainParams
	}{
		{"Bitcoin", BitcoinParams},
		{"Litecoin", LitecoinParams},
		{"Dogecoin", DogecoinParams},
		{"BitcoinCash", BitcoinCashParams},
		{"Zcash", ZcashParams},
	}

	for _, tt := range presets {
		t.Run(tt.name, func(t *testing.T) {
			if tt.params.Name == "" {
				t.Error("Name is empty")
			}
			if tt.params.Ticker == "" {
				t.Error("Ticker is empty")
			}
			if tt.params.DustLimit <= 0 {
				t.Error("DustLimit should be positive")
			}
			if tt.params.TxOverhead <= 0 {
				t.Error("TxOverhead should be positive")
			}
			if len(tt.params.InputSizes) == 0 {
				t.Error("InputSizes is empty")
			}
			if len(tt.params.OutputSizes) == 0 {
				t.Error("OutputSizes is empty")
			}
		})
	}
}
