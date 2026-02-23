package aavev3

import (
	"context"
	"math/big"
	"testing"

	goeth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

type mockRPC struct {
	calls    []goeth.CallMsg
	results  map[string][]byte
	callFunc func(goeth.CallMsg) ([]byte, error)
}

func (m *mockRPC) CallContract(ctx context.Context, call goeth.CallMsg, blockNumber *big.Int) ([]byte, error) {
	m.calls = append(m.calls, call)
	if m.callFunc != nil {
		return m.callFunc(call)
	}
	key := call.To.Hex() + ":" + ethcommon.Bytes2Hex(call.Data[:4])
	if result, ok := m.results[key]; ok {
		return result, nil
	}
	return nil, nil
}

func decimalsResponse(d uint8) []byte {
	b := make([]byte, 32)
	b[31] = d
	return b
}

func symbolResponse(s string) []byte {
	offset := make([]byte, 32)
	offset[31] = 0x20

	length := make([]byte, 32)
	length[31] = byte(len(s))

	padded := make([]byte, 32)
	copy(padded, []byte(s))

	result := make([]byte, 0, 96)
	result = append(result, offset...)
	result = append(result, length...)
	result = append(result, padded...)
	return result
}

func newMockClient() (*Client, *mockRPC) {
	deploy := Deployment{
		Pool:         ethcommon.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2"),
		DataProvider: ethcommon.HexToAddress("0x7B4EB56E7CD4b454BA8ff71E4518426c9B8bFe4B"),
	}

	asset := ethcommon.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	rpc := &mockRPC{
		results: map[string][]byte{
			asset.Hex() + ":" + ethcommon.Bytes2Hex(erc20Codec.PackDecimals()[:4]): decimalsResponse(6),
			asset.Hex() + ":" + ethcommon.Bytes2Hex(erc20Codec.PackSymbol()[:4]):   symbolResponse("USDC"),
		},
	}

	return NewClient(rpc, deploy), rpc
}

func TestBuildDepositTx(t *testing.T) {
	client, _ := newMockClient()
	asset := ethcommon.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	user := ethcommon.HexToAddress("0x1111111111111111111111111111111111111111")

	txs, err := BuildDepositTx(context.Background(), client, asset, "100", user)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(txs) != 2 {
		t.Fatalf("expected 2 transactions, got %d", len(txs))
	}

	if txs[0].To != asset {
		t.Errorf("approve tx should target asset contract, got %s", txs[0].To.Hex())
	}

	if txs[1].To != client.PoolAddress() {
		t.Errorf("supply tx should target pool, got %s", txs[1].To.Hex())
	}

	if len(txs[0].Data) != 68 {
		t.Errorf("approve calldata length = %d, want 68", len(txs[0].Data))
	}

	if len(txs[1].Data) != 132 {
		t.Errorf("supply calldata length = %d, want 132", len(txs[1].Data))
	}

	if txs[0].Value.Sign() != 0 || txs[1].Value.Sign() != 0 {
		t.Error("value should be 0 for token operations")
	}
}

func TestBuildWithdrawTx(t *testing.T) {
	client, _ := newMockClient()
	asset := ethcommon.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	user := ethcommon.HexToAddress("0x1111111111111111111111111111111111111111")

	txs, err := BuildWithdrawTx(context.Background(), client, asset, "50", user)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(txs) != 1 {
		t.Fatalf("expected 1 transaction, got %d", len(txs))
	}

	if txs[0].To != client.PoolAddress() {
		t.Errorf("withdraw tx should target pool, got %s", txs[0].To.Hex())
	}

	if len(txs[0].Data) != 100 {
		t.Errorf("withdraw calldata length = %d, want 100", len(txs[0].Data))
	}
}

func TestBuildWithdrawTxMax(t *testing.T) {
	client, _ := newMockClient()
	asset := ethcommon.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	user := ethcommon.HexToAddress("0x1111111111111111111111111111111111111111")

	txs, err := BuildWithdrawTx(context.Background(), client, asset, "max", user)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(txs) != 1 {
		t.Fatalf("expected 1 transaction, got %d", len(txs))
	}

	withdrawData := txs[0].Data
	amountBytes := withdrawData[36:68]
	amount := new(big.Int).SetBytes(amountBytes)
	if amount.Cmp(MaxUint256) != 0 {
		t.Errorf("max withdraw should use MaxUint256, got %s", amount)
	}
}

func TestBuildBorrowTx(t *testing.T) {
	client, _ := newMockClient()
	asset := ethcommon.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	user := ethcommon.HexToAddress("0x1111111111111111111111111111111111111111")

	txs, err := BuildBorrowTx(context.Background(), client, asset, "1000", user)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(txs) != 1 {
		t.Fatalf("expected 1 transaction, got %d", len(txs))
	}

	if txs[0].To != client.PoolAddress() {
		t.Errorf("borrow tx should target pool, got %s", txs[0].To.Hex())
	}

	if len(txs[0].Data) != 164 {
		t.Errorf("borrow calldata length = %d, want 164", len(txs[0].Data))
	}
}

func TestBuildRepayTx(t *testing.T) {
	client, _ := newMockClient()
	asset := ethcommon.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	user := ethcommon.HexToAddress("0x1111111111111111111111111111111111111111")

	txs, err := BuildRepayTx(context.Background(), client, asset, "500", user)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(txs) != 2 {
		t.Fatalf("expected 2 transactions, got %d", len(txs))
	}

	if txs[0].To != asset {
		t.Errorf("approve tx should target asset contract, got %s", txs[0].To.Hex())
	}

	if txs[1].To != client.PoolAddress() {
		t.Errorf("repay tx should target pool, got %s", txs[1].To.Hex())
	}

	if len(txs[0].Data) != 68 {
		t.Errorf("approve calldata length = %d, want 68", len(txs[0].Data))
	}

	if len(txs[1].Data) != 132 {
		t.Errorf("repay calldata length = %d, want 132", len(txs[1].Data))
	}
}
