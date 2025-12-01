package util

import (
	"fmt"
	"strings"

	"github.com/vultisig/recipes/types"
)

func ParseResource(resource string) (*types.ResourcePath, error) {
	parts := strings.Split(resource, ".")
	switch len(parts) {
	case 3:
		return &types.ResourcePath{
			ChainId:    parts[0],
			ProtocolId: parts[1],
			FunctionId: parts[2],
		}, nil
	case 2:
		return &types.ResourcePath{
			ChainId:    parts[0],
			ProtocolId: parts[1],
			FunctionId: "",
		}, nil
	default:
		return nil, fmt.Errorf("invalid resource format: %s", resource)
	}
}
