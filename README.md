[![Go Reference](https://pkg.go.dev/badge/github.com/disgoorg/validate.svg)](https://pkg.go.dev/github.com/disgoorg/validate)
[![Go Report](https://goreportcard.com/badge/github.com/disgoorg/validate)](https://goreportcard.com/report/github.com/disgoorg/validate)
[![Go Version](https://img.shields.io/github/go-mod/go-version/disgoorg/validate)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/disgoorg/validate/blob/master/LICENSE)
[![DisGo Version](https://img.shields.io/github/v/tag/disgoorg/validate?label=release)](https://github.com/disgoorg/validate/releases/latest)
[![DisGo Discord](https://discord.com/api/guilds/817327181659111454/widget.png)](https://discord.gg/TewhTfDpvW)


# validate

Validate is a simple type validation library built with Go 1.18 generics. It provides an alternative approach to using struct tags to validate structs or other types.

## Getting Started

### Installing

```sh
$ go get github.com/disgoorg/validate
```

## Usage

Usage is simple. Create a new type and implement the `validate.Validator` interface. Then, use the `validate.Validate` function with `validate.New` and pass the field or value to validate as first parameter.
After this you can pass as many `validate.ValidateFunc[any]` functions as you want to validate the value. The functions will be executed in order and the first one that returns an error will stop the validation.

```go
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
```

## Custom Validators

Validate comes with a few predefined validators, but you can also implement your own validators.
For this you can use functions with closures to pass parameters to the validator func.

Here is an example:
```go
func StringRange(min int, max int) ValidatorFunc[string] {
	return func(v string) error {
		if len(v) < min || len(v) > max {
			return fmt.Errorf("string must be between %d and %d characters", min, max)
		}
		return nil
	}
}
```

## Documentation

Documentation can be found here

* [![Go Reference](https://pkg.go.dev/badge/github.com/disgoorg/disgo.svg)](https://pkg.go.dev/github.com/disgoorg/validate)

## Examples

You can find examples [here](https://github.com/disgoorg/validate/tree/master/_example)

# Troubleshooting

For help feel free to open an issues or reach out on [Discord](https://discord.gg/TewhTfDpvW)

## Contributing

Contributions are welcomed but for bigger changes we recommend first reaching out via [Discord](https://discord.gg/TewhTfDpvW) or create an issue to discuss your problems, intentions and ideas.

## License

Distributed under the [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/disgoorg/validate/blob/master/LICENSE). See LICENSE for more information.

