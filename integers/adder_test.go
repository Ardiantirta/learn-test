package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, expected int) {
		t.Helper()

		if got != expected {
			t.Errorf(`expected '%d' but got '%d'`, expected, got)
		}
	}

	t.Run(`adding 1 and 5`, func(t *testing.T) {
		got := Add(1, 5)
		expected := 6

		assertCorrectMessage(t, got, expected)
	})

	t.Run(`adding 59 and 26`, func(t *testing.T) {
		got := Add(59, 26)
		expected := 85

		assertCorrectMessage(t, got, expected)
	})
}

func ExampleAdd() {
	sum := Add(59, 26)
	fmt.Println(sum)
	// Output: 85
}
