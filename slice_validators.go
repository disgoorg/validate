package validate

import (
	"fmt"
)

func ErrLength(max int) error {
	return fmt.Errorf("slice cannot have more than %d elements", max)
}

func SliceLength[T any](max int) ValidatorFunc[[]T] {
	return func(v []T) error {
		if len(v) > max {
			return ErrLength(max)
		}
		return nil
	}
}

func ErrNilElement(i int) error {
	return fmt.Errorf("slice cannot contain nil elements. index: %d", i)
}

func SliceNoneNil[T any](v []T) error {
	for i, e := range v {
		if e == nil {
			return ErrNilElement(i)
		}
	}
	return nil
}
