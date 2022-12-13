package validate

import (
	"fmt"
	"reflect"
)

func ErrMaxLen(max int) error {
	return fmt.Errorf("slice cannot have more than %d elements", max)
}

func SliceMaxLen[T any](max int) ValueValidateFunc[[]T] {
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
	for i, e := range v {
		typeOf := reflect.TypeOf(e)
		if typeOf == nil {
			return ErrNilElement(i)
		}
		switch typeOf.Kind() {
		case reflect.Interface, reflect.Slice, reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer:
			if reflect.ValueOf(e).IsNil() {
				return ErrNilElement(i)
			}
		}
	}
	return nil
}
