package validate

import (
	"errors"
	"reflect"
)

var ErrRequired = errors.New("value is required")

func Required[T any](v T) error {
	var t T
	if reflect.DeepEqual(v, t) {
		return ErrRequired
	}
	return nil
}
