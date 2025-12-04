package cosmos

// BroadcastTxRequest is the request payload for broadcasting transactions
type BroadcastTxRequest struct {
	TxBytes []byte `json:"tx_bytes"`
	Mode    string `json:"mode"`
}

// BroadcastTxResponse is the response from broadcasting a transaction
type BroadcastTxResponse struct {
	TxResponse *TxResponse `json:"tx_response"`
}

// TxResponse contains the result of a transaction broadcast
type TxResponse struct {
	Code      uint32 `json:"code"`
	Codespace string `json:"codespace"`
	Data      string `json:"data"`
	RawLog    string `json:"raw_log"`
	TxHash    string `json:"txhash"`
}

