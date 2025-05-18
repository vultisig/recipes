package util

import (
	"fmt"
	"strings"

	"github.com/vultisig/recipes/types"
)

func ParseResource(resource string) (*types.ResourcePath, error) {
	parts := strings.Split(resource, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid resource format: %s", resource)
	}

	return &types.ResourcePath{
		ChainId:    parts[0],
		ProtocolId: parts[1],
		FunctionId: parts[2],
	}, nil
}
