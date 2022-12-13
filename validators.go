package validate

import (
	"errors"
	"fmt"
	"golang.org/x/exp/slices"
)

var ErrRequired = errors.New("value is required")

func Required[T comparable](v T) error {
	var t T
	if v == t {
		return ErrRequired
	}
	return nil
}

func ErrFlags[T ~int64](a ...T) error {
	var flags string
	for i := range a {
		flags += fmt.Sprintf("%s, ", a[i])
	}
	return fmt.Errorf("flags can't include the following flags: %s", a)
}

func AllowedFlags[T int64 | ~int64](v T, flags ...T) error {
	var disallowedFlags []T
	for flag := T(1); flag < (1 << 8); flag *= 2 {
		if !slices.Contains(flags, flag) && (v&flag) == flag {
			disallowedFlags = append(disallowedFlags, flag)
		}
	}

	return ErrFlags(disallowedFlags...)
}

func ErrEnum[T comparable](a ...T) error {
	return fmt.Errorf("string needs to be one of the following: %v", a)
}

func Enum[T comparable](a ...T) ValueValidateFunc[T] {
	return func(v T) error {
		for _, e := range a {
			if v == e {
				return nil
			}
		}
		return ErrEnum(a...)
	}
}

func ErrEnumInMapValues[K comparable, V comparable](a map[K]V) error {
	var values string
	for k := range a {
		values += fmt.Sprintf("%v, ", a[k])
	}
	return fmt.Errorf("string needs to be one of the following: %v", values)
}

func EnumInMapValues[K comparable, V comparable](a map[K]V) ValueValidateFunc[V] {
	return func(v V) error {
		for k := range a {
			if v == a[k] {
				return nil
			}
		}
		return ErrEnumInMapValues(a)
	}
}

func ErrEnumInMapKeys[K comparable, V any](a map[K]V) error {
	var keys string
	for k := range a {
		keys += fmt.Sprintf("%v, ", k)
	}
	return fmt.Errorf("string needs to be one of the following: %v", keys)
}

func EnumInMapKeys[K comparable, V any](a map[K]V) ValueValidateFunc[K] {
	return func(v K) error {
		_, ok := a[v]
		if !ok {
			EnumInMapKeys(a)
		}
		return nil
	}
}
