package compare

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
)

type Address struct {
	Falsy[common.Address]
	inner common.Address
}

func NewAddress(raw string) (*Address, error) {
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
