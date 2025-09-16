package solana

import (
	"encoding/json"
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
}

type idlAccount struct {
	Name     string `json:"name"`
	IsMut    bool   `json:"isMut"`
	IsSigner bool   `json:"isSigner"`
}

type argType string

const (
	argU8        argType = "u8"
	argU64       argType = "u64"
	argPublicKey argType = "string"
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
