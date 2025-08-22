package compare

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vultisig/recipes/engine/compare"
)

type Address struct {
	compare.Falsy[common.Address]
	inner common.Address
}

func NewAddress(raw string) (compare.Compare[common.Address], error) {
	return &Address{
		inner: common.HexToAddress(raw),
	}, nil
}

func (a *Address) Fixed(v common.Address) bool {
	return bytes.Equal(a.inner[:], v[:])
}

func (a *Address) Magic(v common.Address) bool {
	return bytes.Equal(a.inner[:], v[:])
}
