package validate

import "errors"

var ErrInvalidValidationType = errors.New("invalid validation value type")

type Validator interface {
	Validate() error
}

type ValidatorFunc[T any] func(t T) error

func New[T any](v T, validatorFuncs ...ValidatorFunc[T]) Validator {
	return &validator[T]{
		v: v,
		f: func(t T) error {
			for _, f := range validatorFuncs {
				if err := f(t); err != nil {
					return err
				}
			}
			return nil
		},
	}
}

type validator[T any] struct {
	v any
	f func(t T) error
}

func (v validator[T]) Validate() error {
	if t, ok := v.v.(T); ok {
		return v.f(t)
	}
	return ErrInvalidValidationType
}

func Validate(validators ...Validator) error {
	for _, v := range validators {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
