package validate

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

type animalType string

const (
	Cat   animalType = "CAT"
	Dog   animalType = "DOG"
	Other animalType = "OTHER"
)

var nameRegex = regexp.MustCompile("^[a-zA-Z]+$")

type animal struct {
	Name string
	Age  int
	Type animalType
}

func (v animal) Validate() error {
	return Validate(
		Value(v.Name, StringRange(3, 10), StringMatchRegex(nameRegex)),
		Value(v.Age, NumberRange(0, 100)),
		Value(v.Type, Required[animalType], Enum(Cat, Dog, Other)),
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
			e: ErrStringRange(3, 10),
		},
		{
			v: animal{
				Name: "test",
				Age:  -1,
				Type: Other,
			},
			e: ErrNumberRange(0, 100),
		},
		{
			v: animal{
				Name: "test",
				Age:  10,
				Type: "fish",
			},
			e: ErrEnum(Cat, Dog, Other),
		},
		{
			v: animal{
				Name: "####",
				Age:  10,
				Type: "fish",
			},
			e: ErrStringMatchRegex(nameRegex),
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
		fmt.Println(err)
	}
}
