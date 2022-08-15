package validate

import (
	"fmt"
	"reflect"
)

func ErrMaxLen(max int) error {
	return fmt.Errorf("slice cannot have more than %d elements", max)
}

func SliceMaxLen[T any](max int) ValidatorFunc[[]T] {
	return func(v []T) error {
		if len(v) > max {
			return ErrMaxLen(max)
		}
		return nil
	}
}

func ErrNilElement(i int) error {
	return fmt.Errorf("slice cannot contain nil elements. index: %d", i)
}

func SliceNoneNil[T any](v []T) error {
	var t T
	switch reflect.TypeOf(t).Kind() {
	case reflect.Interface, reflect.Slice, reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer:
	default:
		return nil
	}

	for i, e := range v {
		value := reflect.ValueOf(e)
		if value.IsNil() {
			return ErrNilElement(i)
		}
	}
	return nil
}
