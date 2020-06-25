package main

import "testing"

func TestCalculateDistance(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, expected float64) {
		t.Helper()

		if got != expected {
			t.Errorf(`expected %f but got %f`, expected, got)
		}
	}

	t.Run("a: [0, 3]; b: [4, 0]", func(t *testing.T) {
		a := [2]float64{0, 3}
		b := [2]float64{4, 0}

		got := CalculateDistance(a, b)
		expected := float64(5)

		assertCorrectMessage(t, got, expected)
	})

	t.Run("a: [0, 6]; b: [8, 0]", func(t *testing.T) {
		a := [2]float64{0, 6}
		b := [2]float64{8, 0}

		got := CalculateDistance(a, b)
		expected := float64(10)

		assertCorrectMessage(t, got, expected)
	})
}

func BenchmarkCalculateDistance(b *testing.B) {
	x := [2]float64{0, 6}
	y := [2]float64{8, 0}

	for i := 0; i < b.N; i++ {
		CalculateDistance(x, y)
	}
}
