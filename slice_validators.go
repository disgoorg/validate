package validate

import (
	"fmt"
)

func ErrLength(max int) error {
	return fmt.Errorf("slice cannot have more than %d elements", max)
}

func SliceLength(max int) ValidatorFunc[[]any] {
	return func(v []any) error {
		if len(v) > max {
			return ErrLength(max)
		}
		return nil
	}
}

func ErrorNilElement(i int) error {
	return fmt.Errorf("slice cannot contain nil elements. index: %d", i)
}

func SliceNoneNil() ValidatorFunc[[]any] {
	return func(v []any) error {
		for i, e := range v {
			if e == nil {
				return ErrorNilElement(i)
			}
		}
		return nil
	}
}
