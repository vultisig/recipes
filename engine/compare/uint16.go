package compare

import (
	"fmt"
	"strconv"
)

type Uint16 struct {
	Falsy[uint16]
	inner uint16
}

func NewUint16(raw string) (Compare[uint16], error) {
	v, err := strconv.ParseUint(raw, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("failed to parse string to uint16: %s", raw)
	}
	return &Uint16{
		inner: uint16(v),
	}, nil
}

func (u *Uint16) Fixed(actual uint16) bool {
	return u.inner == actual
}

func (u *Uint16) Min(actual uint16) bool {
	return u.inner <= actual
}

func (u *Uint16) Max(actual uint16) bool {
	return u.inner >= actual
}
