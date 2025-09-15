package conv

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

func CopyProto[T proto.Message](in T) (T, error) {
	if in == nil {
		return *new(T), errors.New("nil input")
	}
	return proto.Clone(in).(T), nil
}
