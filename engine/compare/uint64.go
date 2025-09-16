package compare

import (
	"fmt"
	"strconv"
)

type Uint64 struct {
	Falsy[uint64]
	expected uint64
}

func NewUint64(raw string) (Compare[uint64], error) {
	expected, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uint64: %w", err)
	}

	return &Uint64{
		expected: expected,
	}, nil
}

func (u *Uint64) Fixed(actual uint64) bool {
	return u.expected == actual
}

func (u *Uint64) Min(actual uint64) bool {
	return actual >= u.expected
}

func (u *Uint64) Max(actual uint64) bool {
	return actual <= u.expected
}
