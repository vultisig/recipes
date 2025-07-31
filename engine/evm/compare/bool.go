package compare

import (
	"strconv"
)

type Bool struct {
	Falsy[bool]
	inner bool
}

func NewBool(raw string) (Compare[bool], error) {
	v, err := strconv.ParseBool(raw)
	if err != nil {
		return nil, err
	}
	return &Bool{
		inner: v,
	}, nil
}

func (b *Bool) Fixed(actual bool) bool {
	return b.inner == actual
}
