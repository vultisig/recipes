package conv

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func CopyProto[T proto.Message](in T) (T, error) {
	b, err := protojson.Marshal(in)
	if err != nil {
		return *new(T), fmt.Errorf("failed to marshal proto: %w", err)
	}

	out := *new(T)
	v := reflect.ValueOf(out)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		// For pointer types, we need to create a new instance
		out = proto.Clone(in).(T)
		proto.Reset(out)
	}

	err = protojson.Unmarshal(b, out)
	if err != nil {
		return *new(T), fmt.Errorf("failed to unmarshal proto: %w", err)
	}

	return out, nil
}
