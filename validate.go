package validate

import "errors"

type (
	// Validator is an interface that can be implemented to validate a value
	Validator interface {
		Validate() error
	}

	// ValueValidateFunc is a function that validates a value
	ValueValidateFunc[V any] func(v V) error

	// SliceValidateFunc is a function that validates a slice index and value
	SliceValidateFunc[V any] func(i int, v V) error

	// MapValidateFunc is a function that validates a map key and value
	MapValidateFunc[K comparable, V any] func(k K, v V) error
)

// AsSlice converts the ValueValidateFunc to a SliceValidateFunc
func (f ValueValidateFunc[V]) AsSlice() SliceValidateFunc[V] {
	return func(_ int, t V) error {
		return f(t)
	}
}

// AsMap converts the ValueValidateFunc to a MapValidateFunc
func (f ValueValidateFunc[V]) AsMap() MapValidateFunc[any, V] {
	return func(_ any, v V) error {
		return f(v)
	}
}

// Validate validates all the given Validator(s) and returns all validation errors as a single error
func Validate(validators ...Validator) error {
	var errs []error
	for i := range validators {
		if err := validators[i].Validate(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

// Value validates the given value with the given ValueValidateFunc(s) and checks if the value implements the Validator interface and calls it's Validate method as well
func Value[V any](v V, validateFuncs ...ValueValidateFunc[V]) Validator {
	return &valueValidator[V]{
		value: v,
		validateFunc: func(v V) error {
			var errs []error
			for i := range validateFuncs {
				if err := validateFuncs[i](v); err != nil {
					errs = append(errs, err)
				}
			}
			return errors.Join(errs...)
		},
	}
}

// Slice validates the given slice of values with the given SliceValidateFunc(s) and checks if the value implements the Validator interface and calls it's Validate method as well
func Slice[V any](v []V, validateFuncs ...SliceValidateFunc[V]) Validator {
	return &sliceValidator[V]{
		values: v,
		validateFunc: func(i int, v V) error {
			var errs []error
			for i := range validateFuncs {
				if err := validateFuncs[i](i, v); err != nil {
					errs = append(errs, err)
				}
			}
			return errors.Join(errs...)
		},
	}
}

// Map validates the given map of key, values with the given MapValidateFunc(s) and checks if the value implements the Validator interface and calls it's Validate method as well
func Map[K comparable, V any](v map[K]V, validateFuncs ...MapValidateFunc[K, V]) Validator {
	return &mapValidator[K, V]{
		values: v,
		validateFunc: func(k K, v V) error {
			for i := range validateFuncs {
				if err := validateFuncs[i](k, v); err != nil {
					return err
				}
			}
			return nil
		},
	}
}

type valueValidator[V any] struct {
	value        V
	validateFunc ValueValidateFunc[V]
}

func (v valueValidator[V]) Validate() error {
	var errs []error
	if err := v.validateFunc(v.value); err != nil {
		errs = append(errs, err)
	}
	if validator, ok := any(v.value).(Validator); ok {
		if err := validator.Validate(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

type sliceValidator[V any] struct {
	values       []V
	validateFunc SliceValidateFunc[V]
}

func (v sliceValidator[V]) Validate() error {
	var errs []error
	for i := range v.values {
		if err := v.validateFunc(i, v.values[i]); err != nil {
			errs = append(errs, err)
		}
		if validator, ok := any(v.values[i]).(Validator); ok {
			if err := validator.Validate(); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return errors.Join(errs...)
}

type mapValidator[K comparable, V any] struct {
	values       map[K]V
	validateFunc MapValidateFunc[K, V]
}

func (v mapValidator[K, V]) Validate() error {
	var errs []error
	for k := range v.values {
		if err := v.validateFunc(k, v.values[k]); err != nil {
			errs = append(errs, err)
		}
		if validator, ok := any(v.values[k]).(Validator); ok {
			if err := validator.Validate(); err != nil {
				errs = append(errs, err)
			}
		}
	}
	return nil
}
