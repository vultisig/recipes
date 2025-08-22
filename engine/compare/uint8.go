package compare

import (
	"fmt"
	"strconv"
)

type Uint8 struct {
	Falsy[uint8]
	inner uint8
}

func NewUint8(raw string) (Compare[uint8], error) {
	v, err := strconv.ParseUint(raw, 10, 8)
	if err != nil {
		return nil, fmt.Errorf("failed to parse string to uint8: %s", raw)
	}
	return &Uint8{
		inner: uint8(v),
	}, nil
}

func (u *Uint8) Fixed(actual uint8) bool {
	return u.inner == actual
}

func (u *Uint8) Min(actual uint8) bool {
	return u.inner <= actual
}

func (u *Uint8) Max(actual uint8) bool {
	return u.inner >= actual
}
