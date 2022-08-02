package validate

import (
	"errors"
	"regexp"
)

var ErrStringRange = errors.New("string is not in range")

func StringRange(min int, max int) ValidatorFunc[string] {
	return func(v string) error {
		if len(v) < min || len(v) > max {
			return ErrStringRange
		}
		return nil
	}
}

var ErrStringMatchRegex = errors.New("string does not match regex")

func StringMatchRegex(regexp *regexp.Regexp) ValidatorFunc[string] {
	return func(v string) error {
		if !regexp.MatchString(v) {
			return ErrStringMatchRegex
		}
		return nil
	}
}
