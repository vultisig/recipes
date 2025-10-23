package jsonschema

func Definitions() map[string]interface{} {
	asset := NewAsset()

	return map[string]interface{}{
		asset.Name(): asset.Schema(),
	}
}

func id(label string) string {
	return "https://github.com/vultisig/recipes/jsonschema/types/" + label
}
