package btc

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
)

type UtxoFetcher interface {
	GetTxOut(txHash *chainhash.Hash, index uint32, includeMempool bool) (*btcutil.Tx, *wire.TxOut, error)
}

type RpcClient struct {
	client *rpcclient.Client
}

func NewRpcClient(host, user, pass string, useTLS bool) (*RpcClient, error) {
	config := &rpcclient.ConnConfig{
		Host:         host,
		User:         user,
		Pass:         pass,
		DisableTLS:   !useTLS,
		HTTPPostMode: true,
	}

	client, err := rpcclient.New(config, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %w", err)
	}

	return &RpcClient{client: client}, nil
}

func (c *RpcClient) GetTxOut(txHash *chainhash.Hash, index uint32, includeMempool bool) (*btcutil.Tx, *wire.TxOut, error) {
	txOutResult, err := c.client.GetTxOut(txHash, index, includeMempool)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get transaction output: %w", err)
	}

	if txOutResult == nil {
		return nil, nil, fmt.Errorf("transaction output not found or already spent")
	}

	value, err := btcutil.NewAmount(txOutResult.Value)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid amount: %w", err)
	}

	scriptBytes, err := hex.DecodeString(txOutResult.ScriptPubKey.Hex)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to decode script: %w", err)
	}

	txOut := &wire.TxOut{
		Value:    int64(value),
		PkScript: scriptBytes,
	}

	return nil, txOut, nil
}
