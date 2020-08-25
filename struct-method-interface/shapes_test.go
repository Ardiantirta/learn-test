package struct_method_interface

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{"Rectangle", Rectangle{10.0, 10.0}, 40.0},
		{"Circle", Circle{10.0}, 62.83185307179586},
		{"Triangle", Triangle{Base: 3.0, Height: 4.0}, 12.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Rectangle", Rectangle{Width: 10.0, Height: 10.0}, 100.0},
		{"Circle", Circle{Radius: 10.0}, 314.1592653589793},
		{"Triangle", Triangle{Base: 10.0, Height: 10.0}, 50.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
