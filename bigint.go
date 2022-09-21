package bigint

import "errors"

// Bigint ...
type Bigint struct {
	value string
}

var (
	// ErrorNotNumber used when input consists of not only numbers
	ErrorNotNumber = errors.New("input not a number")
)

func validation(num string) error {
	// TODO
	return nil
}

// NewInt returns new Bigint value based on input
func NewInt(num string) (Bigint, error) {
	// TODO
	// step 1: validation
	// step 2: clean up 0000123 -> 123 | +123 -> 123 | +-+123 -> error
	// step 3: insert value to the bigint struct and return
	return Bigint{
		value: "123",
	}, nil
}

// Set methods updates value
func (z *Bigint) Set(num string) error {
	// TODO
	// step 1: validation
	// step 2: clean up
	// step 3: set new value
	return nil
}

// Add ...
func Add(a, b Bigint) Bigint {
	// TODO
	// step 1:  cast - a.value -> int32,  b.value -> int32
	// step 2: add casted numbers
	// step 3: cast result to string
	// step 4: return result
	return Bigint{
		value: a.value + b.value, // FIX this
	}
}

// Sub ...
func Sub(a, b Bigint) Bigint {
	// TODO
	// step 1:  cast - a.value -> int32,  b.value -> int32
	// step 2: sub casted numbers
	// step 3: cast result to string
	// step 4: return result
	return Bigint{
		value: "123", // FIX this
	}
}

// Multiply ...
func Multiply(a, b Bigint) Bigint {
	// TODO
	// step 1:  cast - a.value -> int32,  b.value -> int32
	// step 2: multiply casted numbers
	// step 3: cast result to string
	// step 4: return result
	return Bigint{
		value: "123", // FIX this
	}
}

// Mod ...
func Mod(a, b Bigint) Bigint {
	// TODO
	// step 1:  cast - a.value -> int32,  b.value -> int32
	// step 2: mod casted numbers
	// step 3: cast result to string
	// step 4: return result
	return Bigint{
		value: "123", // FIX this
	}
}

// Abs ...
func (z Bigint) Abs() Bigint {
	// TODO
	// Abs
	return Bigint{
		value: "123", // FIX this
	}
}
