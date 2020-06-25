package main

import (
	"fmt"
	"math"
	"time"
)

// Add takes two integers and return the sum of them.
func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func CalculateDistance(a, b [2]float64)  float64 {
	dx := b[0] - a[0]
	dy := b[1] - a[1]
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}

func CalculateSquaredDistance(a, b [2]float64) float64 {
	dx := b[0] - a[0]
	dy := b[1] - a[1]
	return math.Pow(dx, 2) + math.Pow(dy, 2)
}

var (
	A = [2]float64{0, 6}
	B = [2]float64{8, 0}
)

func main() {
	start := time.Now()
	result := CalculateDistance(A, B)
	duration := time.Since(start)

	fmt.Printf("result: %0.f\ntime consumed by func 'CalculateDistance': %d\n", result, duration.Milliseconds())

	start = time.Now()
	result = CalculateSquaredDistance(A, B)
	duration = time.Since(start)

	fmt.Printf("result: %0.f\ntime consumed by func 'CalculateSquaredDistance': %d", result, duration.Milliseconds())
}
