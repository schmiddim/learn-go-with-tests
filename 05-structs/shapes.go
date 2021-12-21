package shapes

import "math"

type Shape interface {
	Area() float64 //Everything that implements method Area() has the interface
}

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

func (rectangle Rectangle) Perimeter() (result float64) {
	result = (rectangle.Width + rectangle.Height) * 2
	return
}

func (rectangle Rectangle) Area() (result float64) {

	result = rectangle.Height * rectangle.Width
	return
}

func (circle Circle) Area() (result float64) {

	result = circle.Radius * circle.Radius * math.Pi
	return
}
