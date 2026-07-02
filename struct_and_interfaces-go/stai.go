package structandinterfacesgo

import "math"

type Shape interface {
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Perimeter() float64 {
	perimeter := 2 * (r.height + r.width)
	return perimeter
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
