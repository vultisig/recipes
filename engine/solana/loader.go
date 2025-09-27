package solana

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strings"

	idl_embed "github.com/vultisig/recipes/idl"
)

type idl struct {
	Instructions []idlInstruction `json:"instructions"`
	Accounts     []idlAccount     `json:"accounts,omitempty"`
	Types        []idlType        `json:"types,omitempty"`
	Name         string           `json:"name"`
	Version      string           `json:"version"`
}

type idlInstruction struct {
	Name     string        `json:"name"`
	Accounts []idlAccount  `json:"accounts"`
	Args     []idlArgument `json:"args"`
	Metadata idlMetadata   `json:"metadata"`
}

type idlMetadata struct {
	Discriminator []byte `json:"discriminator,omitempty"`
}

type idlAccount struct {
	Name     string `json:"name"`
	IsMut    bool   `json:"isMut"`
	IsSigner bool   `json:"isSigner"`
}

type argType string

func (t *argType) UnmarshalJSON(data []byte) error {
	// Try to unmarshal as a simple string first
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		// Simple string like "u8", "u16", "publicKey", etc
		*t = argType(str)
		return nil
	}

	extra := map[string]interface{}{}
	err := json.Unmarshal(data, &extra)
	if err != nil {
		return fmt.Errorf("`type` field failed to unmarshal to map: %w", err)
	}

	// always map with 1 key
	for key := range extra {
		*t = argType(key)
		return nil
	}
	return errors.New("`type` field: unexpected empty map")
}

const (
	argU8        argType = "u8"
	argU16       argType = "u16"
	argU64       argType = "u64"
	argPublicKey argType = "publicKey"

	// nested types
	argVec argType = "vec"
)

type idlArgument struct {
	Name string  `json:"name"`
	Type argType `json:"type"`
}

type idlType struct {
	Name string                 `json:"name"`
	Type map[string]interface{} `json:"type"`
}

type protocolID = string

func loadIDLDir() (map[protocolID]idl, error) {
	base := "."

	entries, err := idl_embed.Dir.ReadDir(base)
	if err != nil {
		return nil, fmt.Errorf("failed to read idl dir: err=%w", err)
	}

	idls := make(map[protocolID]idl)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		filepath := path.Join(base, entry.Name())
		file, er := idl_embed.Dir.Open(filepath)
		if er != nil {
			return nil, fmt.Errorf("failed to open idl json: path=%s, err=%w", filepath, er)
		}

		var idlItem idl
		er = json.NewDecoder(file).Decode(&idlItem)
		_ = file.Close()
		if er != nil {
			return nil, fmt.Errorf("failed to parse idl json: %w", er)
		}

		idls[strings.TrimSuffix(entry.Name(), ".json")] = idlItem
	}
	return idls, nil
}
