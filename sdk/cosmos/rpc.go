package cosmos

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// RPCClient interface for Cosmos JSON-RPC calls
type RPCClient interface {
	BroadcastTx(ctx context.Context, txBytes []byte) (*BroadcastTxResponse, error)
}

// HTTPRPCClient implements RPCClient using HTTP
type HTTPRPCClient struct {
	endpoints []string
	client    *http.Client
}

// NewHTTPRPCClient creates a new HTTP RPC client with the given endpoints
func NewHTTPRPCClient(endpoints []string) *HTTPRPCClient {
	return &HTTPRPCClient{
		endpoints: endpoints,
		client:    &http.Client{Timeout: 30 * time.Second},
	}
}

// BroadcastTx broadcasts a signed transaction to the Cosmos network
func (c *HTTPRPCClient) BroadcastTx(ctx context.Context, txBytes []byte) (*BroadcastTxResponse, error) {
	txBase64 := base64.StdEncoding.EncodeToString(txBytes)

	requestBody := map[string]interface{}{
		"tx_bytes": txBase64,
		"mode":     "BROADCAST_MODE_SYNC",
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	var lastErr error
	for _, endpoint := range c.endpoints {
		url := endpoint + "/cosmos/tx/v1beta1/txs"

		req, reqErr := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(requestJSON))
		if reqErr != nil {
			lastErr = fmt.Errorf("failed to create request for %s: %w", endpoint, reqErr)
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("failed to send request to %s: %w", endpoint, err)
			continue
		}

		body, readErr := io.ReadAll(resp.Body)
		_ = resp.Body.Close()

		if readErr != nil {
			lastErr = fmt.Errorf("failed to read response from %s: %w", endpoint, readErr)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("HTTP error from %s: %d, body: %s", endpoint, resp.StatusCode, string(body))
			continue
		}

		var broadcastResp BroadcastTxResponse
		if err := json.Unmarshal(body, &broadcastResp); err != nil {
			lastErr = fmt.Errorf("failed to parse response from %s: %w", endpoint, err)
			continue
		}

		if broadcastResp.TxResponse != nil && broadcastResp.TxResponse.Code != 0 {
			lastErr = fmt.Errorf("broadcast failed at %s: code=%d, log=%s",
				endpoint, broadcastResp.TxResponse.Code, broadcastResp.TxResponse.RawLog)
			continue
		}

		return &broadcastResp, nil
	}

	return nil, fmt.Errorf("all endpoints failed, last error: %w", lastErr)
}

