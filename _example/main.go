package main

import "github.com/disgoorg/validate"

type Foo struct {
	Bar string
	Baz int
	Bay []*int
}

func (f Foo) Validate() error {
	return validate.Validate(
		validate.Value(f.Bar, validate.Required[string], validate.StringRange(0, 10)),
		validate.Value(f.Baz, validate.Required[int], validate.NumberRange(-5, 5)),
		validate.Value(f.Bay, validate.SliceMaxLen[*int](5), validate.SliceNoneNil[*int]),
	)
}

func main() {
	f := Foo{
		Bar: "a",
		Baz: -1,
		Bay: []*int{nil},
	}
	if err := f.Validate(); err != nil {
		panic(err)
	}
}
