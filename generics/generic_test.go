package main

import (
	"fmt"
	"testing"
)

func Print[T any](value T) {
	fmt.Println(value)
}

func Print2[T int](value T) {
	fmt.Println(value)
}

func Print3[T func(string) string](value T, str string) string {
	return value(str)
}

func TestAny(t *testing.T) {
	Print("a")
	Print('a')
	Print(123)
	Print(func() {})
	Print2(456)

	var func1 = func(str string) string {
		return str + "!!!"
	}

	t.Log(Print3(func1, "bbb"))
}
