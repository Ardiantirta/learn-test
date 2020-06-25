package main

import (
	"fmt"
	"testing"
)

func TestSubtracter(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, expected int) {
		t.Helper()

		if got != expected {
			t.Errorf(`expected '%d' but got '%d'`, expected, got)
		}
	}

	t.Run(`adding 1 and 5`, func(t *testing.T) {
		got := Subtract(1, 5)
		expected := -4

		assertCorrectMessage(t, got, expected)
	})

	t.Run(`adding 59 and 26`, func(t *testing.T) {
		got := Subtract(59, 26)
		expected := 33

		assertCorrectMessage(t, got, expected)
	})
}

func ExampleSubtract() {
	result := Subtract(59, 26)
	fmt.Println(result)
	// Output: 33
}
