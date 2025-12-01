package compare

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

type Bytes struct {
	Falsy[[]byte]
	inner []byte
}

func NewBytes(raw string) (Compare[[]byte], error) {
	b, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, fmt.Errorf("failed to decode b64 string: %w", err)
	}

	result := make([]byte, len(b))
	copy(result[:], b)

	return &Bytes{
		inner: result,
	}, nil
}

func (b *Bytes) Fixed(actual []byte) bool {
	return bytes.Equal(b.inner[:], actual[:])
}
