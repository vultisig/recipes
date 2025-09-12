package compare

import (
	"fmt"
	"strconv"

	"github.com/vultisig/recipes/engine/compare"
)

type U64 struct {
	compare.Falsy[uint64]
	expected uint64
}

func NewU64(raw string) (compare.Compare[uint64], error) {
	expected, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uint64: %w", err)
	}

	return &U64{
		expected: expected,
	}, nil
}

func (u *U64) Fixed(actual uint64) bool {
	return u.expected == actual
}

func (u *U64) Min(actual uint64) bool {
	return actual >= u.expected
}

func (u *U64) Max(actual uint64) bool {
	return actual <= u.expected
}

func (u *U64) Magic(actual uint64) bool {
	return u.Fixed(actual)
}
