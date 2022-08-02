package validate

import (
	"errors"
)

var ErrIntRange = errors.New("int is not in range")

func IntRange(min int, max int) ValidatorFunc[int] {
	return func(v int) error {
		if v < min || v > max {
			return ErrIntRange
		}
		return nil
	}
}
