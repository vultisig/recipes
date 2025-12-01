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

func (inst *idlInstruction) UnmarshalJSON(data []byte) error {
	type Alias idlInstruction
	aux := &struct {
		Discriminator []byte `json:"discriminator"`
		*Alias
	}{
		Alias: (*Alias)(inst),
	}

	err := json.Unmarshal(data, aux)
	if err != nil {
		return err
	}

	if len(aux.Discriminator) > 0 {
		inst.Metadata.Discriminator = aux.Discriminator
	}

	return nil
}

type idlMetadata struct {
	Discriminator []byte `json:"discriminator,omitempty"`
}

type idlAccount struct {
	Name       string `json:"name"`
	IsMut      bool   `json:"isMut"`
	IsSigner   bool   `json:"isSigner"`
	IsOptional bool   `json:"isOptional"`
}

func (acc *idlAccount) UnmarshalJSON(data []byte) error {
	type Alias idlAccount
	aux := &struct {
		Writable *bool `json:"writable"`
		Signer   *bool `json:"signer"`
		Optional *bool `json:"optional"`
		*Alias
	}{
		Alias: (*Alias)(acc),
	}

	err := json.Unmarshal(data, aux)
	if err != nil {
		return err
	}

	if aux.Writable != nil {
		acc.IsMut = *aux.Writable
	}

	if aux.Signer != nil {
		acc.IsSigner = *aux.Signer
	}

	if aux.Optional != nil {
		acc.IsOptional = *aux.Optional
	}

	return nil
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
	argBool      argType = "bool"
	argPublicKey argType = "publicKey"

	// nested types
	argVec     argType = "vec"
	argDefined argType = "defined"
)

type idlArgument struct {
	Name      string                 `json:"name"`
	Type      argType                `json:"type"`
	TypeInfo  map[string]interface{} `json:"-"`
	IsComplex bool                   `json:"-"`
}

func (arg *idlArgument) UnmarshalJSON(data []byte) error {
	type Alias idlArgument
	aux := &struct {
		Type json.RawMessage `json:"type"`
		*Alias
	}{
		Alias: (*Alias)(arg),
	}

	err := json.Unmarshal(data, aux)
	if err != nil {
		return err
	}

	err = json.Unmarshal(aux.Type, &arg.Type)
	if err != nil {
		return err
	}

	var typeInfo map[string]interface{}
	err = json.Unmarshal(aux.Type, &typeInfo)
	if err == nil && len(typeInfo) > 0 {
		arg.TypeInfo = typeInfo
		arg.IsComplex = isComplexType(typeInfo)
	}

	return nil
}

func isComplexType(typeInfo map[string]interface{}) bool {
	if len(typeInfo) == 0 {
		return false
	}

	for key, value := range typeInfo {
		switch key {
		case "vec":
			vecContent, ok := value.(map[string]interface{})
			if !ok {
				return false
			}
			if _, hasDefined := vecContent["defined"]; hasDefined {
				return true
			}
			return isComplexType(vecContent)
		case "defined":
			return true
		case "option":
			optionContent, ok := value.(map[string]interface{})
			if !ok {
				return false
			}
			return isComplexType(optionContent)
		}
	}

	return false
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
