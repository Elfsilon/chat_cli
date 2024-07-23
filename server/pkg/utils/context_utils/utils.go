package ctxutil

import (
	"context"
	"fmt"
	"reflect"
)

type ContextValueKey string

const UserID ContextValueKey = "user_id"

func errNilContextValue(key string, value any, valType string) error {
	return fmt.Errorf(
		"failed converting context's value with key %v, possibly nil: value = %v, type = %v",
		key,
		value,
		valType,
	)
}

func GetValue[T any](ctx context.Context, key ContextValueKey) (T, error) {
	raw := ctx.Value(key)

	var val T
	val, ok := raw.(T)
	if !ok {
		var t string
		if raw == nil {
			t = "nil"
		} else {
			t = reflect.ValueOf(raw).Kind().String()
		}

		return val, errNilContextValue(string(key), raw, t)
	}

	return val, nil
}
