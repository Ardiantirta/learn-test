package struct_method_interface

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Perimeter() float64 {
	side := math.Sqrt(math.Pow(t.Base, 2) + math.Pow(t.Height, 2))
	return t.Base + t.Height + side
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
