package bigint_pkg

import "testing"

func TestAdd(t1 *testing.T) {
	tests := []struct {
		name     string
		num1     Bigint
		num2     Bigint
		expected string
	}{
		{name: "zero values", num1: Bigint{Value: "0"}, num2: Bigint{Value: "0"}, expected: "0"},
		{name: "positive values", num1: Bigint{Value: "123"}, num2: Bigint{Value: "321"}, expected: "444"},
		{name: "negative values", num1: Bigint{Value: "-123"}, num2: Bigint{Value: "-321"}, expected: "-444"},
		{name: "mixed values 1", num1: Bigint{Value: "-123"}, num2: Bigint{Value: "323"}, expected: "200"},
		{name: "mixed values 2", num1: Bigint{Value: "123"}, num2: Bigint{Value: "-323"}, expected: "-200"},
		{name: "reminder values", num1: Bigint{Value: "999"}, num2: Bigint{Value: "123"}, expected: "1122"},
		{name: "big num", num1: Bigint{Value: "312312312331231231233123123123312312312331231231233123123123"}, num2: Bigint{Value: "1"}, expected: "312312312331231231233123123123312312312331231231233123123124"},
	}

	for _, arg := range tests {
		t1.Run(arg.name, func(t2 *testing.T) {
			res := Add(arg.num1, arg.num2)
			if res.Value != arg.expected {
				t2.Errorf("got %v but expected %v, input values: %v, %v", res.Value, arg.expected, arg.num1.Value, arg.num2.Value)
			}
		})
	}

	t1.Log("Finished")
}

func TestSub(t1 *testing.T) {
	tests := []struct {
		name     string
		num1     Bigint
		num2     Bigint
		expected string
	}{
		{name: "zero values", num1: Bigint{Value: "0"}, num2: Bigint{Value: "0"}, expected: "0"},
		{name: "positive values", num1: Bigint{Value: "123"}, num2: Bigint{Value: "321"}, expected: "-198"},
		{name: "negative values", num1: Bigint{Value: "-123"}, num2: Bigint{Value: "-321"}, expected: "198"},
		{name: "mixed values 1", num1: Bigint{Value: "-123"}, num2: Bigint{Value: "323"}, expected: "-446"},
		{name: "mixed values 2", num1: Bigint{Value: "123"}, num2: Bigint{Value: "-323"}, expected: "446"},
		{name: "reminder values", num1: Bigint{Value: "999"}, num2: Bigint{Value: "123"}, expected: "876"},
		{name: "big num", num1: Bigint{Value: "312312312331231231233123123123312312312331231231233123123123"}, num2: Bigint{Value: "1"}, expected: "312312312331231231233123123123312312312331231231233123123122"},
	}

	for _, arg := range tests {
		t1.Run(arg.name, func(t2 *testing.T) {
			res := Sub(arg.num1, arg.num2)
			if res.Value != arg.expected {
				t2.Errorf("got %v but expected %v, input values: %v, %v", res.Value, arg.expected, arg.num1.Value, arg.num2.Value)
			}
		})
	}

	t1.Log("Finished")
}
