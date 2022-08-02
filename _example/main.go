package main

import "github.com/disgoorg/validate"

type Foo struct {
	Bar string
	Baz int
}

func (f Foo) Validate() error {
	return validate.Validate(
		validate.New(f.Bar, validate.Required[string], validate.StringRange(0, 10)),
		validate.New(f.Baz, validate.Required[int], validate.NumberRange(-5, 5)),
	)
}

func main() {
	f := Foo{
		Bar: "",
		Baz: -1,
	}
	if err := f.Validate(); err != nil {
		panic(err)
	}
}