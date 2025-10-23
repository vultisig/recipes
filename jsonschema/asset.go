package jsonschema

type Asset struct{}

func NewAsset() Asset {
	return Asset{}
}

func (a Asset) Name() string {
	return "asset"
}

func (a Asset) Schema() map[string]interface{} {
	return map[string]interface{}{
		"$schema":     "https://json-schema.org/draft/2020-12/schema",
		"$id":         id("asset"),
		"title":       "Asset",
		"description": "An asset schema representing a blockchain asset with token, chain, and address information",
		"type":        "object",
		"properties": map[string]interface{}{
			"token": map[string]interface{}{
				"type":        "string",
				"description": "Smart-contract address. Empty for native asset",
			},
			"chain": map[string]interface{}{
				"type":        "string",
				"description": "Blockchain network identifier as common.Chain string representation",
			},
			"address": map[string]interface{}{
				"type":        "string",
				"description": "User address on the blockchain",
			},
		},
		"required":             []string{"chain", "address"},
		"additionalProperties": false,
	}
}
