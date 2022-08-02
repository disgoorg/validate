package validate

import (
	"errors"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type animalType string

const (
	Cat   = "CAT"
	Dog   = "DOG"
	Other = "OTHER"
)

var ErrInvalidAnimalType = errors.New("invalid animal type")

func validAnimalType(v animalType) error {
	if v == Cat || v == Dog || v == Other {
		return nil
	}
	return ErrInvalidAnimalType
}

var nameRegex = regexp.MustCompile("^[a-zA-Z]+$")

type animal struct {
	Name string
	Age  int
	Type animalType
}

func (v animal) Validate() error {
	return Validate(
		New(v.Name, StringRange(3, 10), StringMatchRegex(nameRegex)),
		New(v.Age, IntRange(0, 100)),
		New(v.Type, Required[animalType], validAnimalType),
	)
}

func TestValidate(t *testing.T) {
	v := []struct {
		v animal
		e error
	}{
		{
			v: animal{
				Name: "test",
				Age:  10,
				Type: Cat,
			},
			e: nil,
		},
		{
			v: animal{
				Name: "t",
				Age:  10,
				Type: Dog,
			},
			e: ErrStringRange,
		},
		{
			v: animal{
				Name: "test",
				Age:  -1,
				Type: Other,
			},
			e: ErrIntRange,
		},
		{
			v: animal{
				Name: "test",
				Age:  10,
				Type: "fish",
			},
			e: ErrInvalidAnimalType,
		},
		{
			v: animal{
				Name: "####",
				Age:  10,
				Type: "fish",
			},
			e: ErrStringMatchRegex,
		},
		{
			v: animal{
				Name: "test",
				Age:  10,
				Type: "",
			},
			e: ErrRequired,
		},
	}

	for _, tt := range v {
		err := tt.v.Validate()
		assert.Equal(t, tt.e, err)
	}
}
