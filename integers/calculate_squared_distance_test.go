package main

import "testing"

func TestCalculateSquaredDistance(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, expected float64) {
		t.Helper()

		if got != expected {
			t.Errorf(`expected %f but got %f`, expected, got)
		}
	}

	t.Run("a: [0, 3]; b: [4, 0]", func(t *testing.T) {
		a := [2]float64{0, 3}
		b := [2]float64{4, 0}

		got := CalculateSquaredDistance(a, b)
		expected := float64(25)

		assertCorrectMessage(t, got, expected)
	})

	t.Run("a: [0, 6]; b: [8, 0]", func(t *testing.T) {
		a := [2]float64{6, 0}
		b := [2]float64{0, 8}

		got := CalculateSquaredDistance(a, b)
		expected := float64(100)

		assertCorrectMessage(t, got, expected)
	})
}

func BenchmarkCalculateSquaredDistance(b *testing.B) {
	x := [2]float64{0, 6}
	y := [2]float64{8, 0}

	for i := 0; i < b.N; i++ {
		CalculateSquaredDistance(x, y)
	}
}
