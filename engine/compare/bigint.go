package compare

import (
	"fmt"
	"math/big"
)

type BigInt struct {
	Falsy[*big.Int]
	inner *big.Int
}

func NewBigInt(raw string) (Compare[*big.Int], error) {
	v, parseOk := new(big.Int).SetString(raw, 10)
	if !parseOk {
		return nil, fmt.Errorf("failed to create big int: %s", raw)
	}
	return &BigInt{
		inner: v,
	}, nil
}

func (b *BigInt) Fixed(actual *big.Int) bool {
	return b.inner.Cmp(actual) == 0
}

func (b *BigInt) Min(actual *big.Int) bool {
	cmp := b.inner.Cmp(actual)
	return cmp == -1 || cmp == 0
}

func (b *BigInt) Max(actual *big.Int) bool {
	cmp := b.inner.Cmp(actual)
	return cmp == 1 || cmp == 0
}
