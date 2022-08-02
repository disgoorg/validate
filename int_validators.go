package validate

import (
	"fmt"
)

type Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func ErrNumberRange[T Number](min T, max T) error {
	return fmt.Errorf("number needs to be between %v and %v", min, max)
}

func NumberRange[T Number](min T, max T) ValidatorFunc[T] {
	return func(v T) error {
		if v < min || v > max {
			return ErrNumberRange(min, max)
		}
		return nil
	}
}
