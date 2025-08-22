package compare

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

type Bytes32 struct {
	Falsy[[32]byte]
	inner [32]byte
}

func NewBytes32(raw string) (Compare[[32]byte], error) {
	b, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, fmt.Errorf("failed to decode b64 string: %w", err)
	}
	if len(b) != 32 {
		return nil, fmt.Errorf("len must be 32, got: %d", len(b))
	}

	var result [32]byte
	copy(result[:], b)

	return &Bytes32{
		inner: result,
	}, nil
}

func (b *Bytes32) Fixed(actual [32]byte) bool {
	return bytes.Equal(b.inner[:], actual[:])
}
