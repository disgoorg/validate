package validate

import (
	"fmt"
	"regexp"
)

func ErrStringRange(mix int, max int) error {
	return fmt.Errorf("string needs to be between %d and %d characters", mix, max)
}

func StringRange(min int, max int) ValidatorFunc[string] {
	return func(v string) error {
		if len(v) < min || len(v) > max {
			return ErrStringRange(min, max)
		}
		return nil
	}
}

func ErrStringMatchRegex(regexp *regexp.Regexp) error {
	return fmt.Errorf("string needs to match regex: %s", regexp.String())
}

func StringMatchRegex(regexp *regexp.Regexp) ValidatorFunc[string] {
	return func(v string) error {
		if !regexp.MatchString(v) {
			return ErrStringMatchRegex(regexp)
		}
		return nil
	}
}
