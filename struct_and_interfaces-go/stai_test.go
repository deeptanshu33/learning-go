package structandinterfacesgo

import "testing"

func TestPerimeteter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != expected {
			t.Errorf("%#v, Expected %g, got %g", shape, expected, got)
		}
	}

	perimeterTests := []struct {
		name              string
		shape             Shape
		expectedPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{12.5, 15.0}, expectedPerimeter: 55.0},
		{name: "Circle", shape: Circle{5.4}, expectedPerimeter: 33.929200658769766},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			checkPerimeter(t, tt.shape, tt.expectedPerimeter)
		})
	}
}
