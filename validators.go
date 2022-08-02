package validate

import (
	"errors"
	"fmt"
)

var ErrRequired = errors.New("value is required")

func Required[T comparable](v T) error {
	var t T
	if v == t {
		return ErrRequired
	}
	return nil
}

func ErrEnum[T comparable](a ...T) error {
	return fmt.Errorf("string needs to be one of the following: %v", a)
}

func Enum[T comparable](a ...T) ValidatorFunc[T] {
	return func(v T) error {
		for _, e := range a {
			if v == e {
				return nil
			}
		}
		return ErrEnum(a...)
	}
}
