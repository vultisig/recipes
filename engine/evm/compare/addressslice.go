package compare

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type AddressSlice struct {
	Falsy[[]common.Address]
	inner []common.Address
}

func NewAddressSlice(raw string) (Compare[[]common.Address], error) {
	var addrs []common.Address
	for _, s := range strings.Split(raw, ",") {
		addrs = append(addrs, common.HexToAddress(s))
	}
	return &AddressSlice{
		inner: addrs,
	}, nil
}

func (a *AddressSlice) Fixed(actual []common.Address) bool {
	if len(a.inner) != len(actual) {
		return false
	}
	for i := range a.inner {
		if !bytes.Equal(a.inner[i][:], actual[i][:]) {
			return false
		}
	}
	return true
}
