package util

import (
	"fmt"
	"reflect"
)

type TestCase[TInput, TExpected any] struct {
	Input    TInput
	Expected TExpected
}

func Test[T any](expected, output T) error {
	if reflect.DeepEqual(expected, output) {
		return nil
	}

	return fmt.Errorf("expected %v, but got %v", expected, output)
}
