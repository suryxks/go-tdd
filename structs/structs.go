package structs

import "math"

type Rectangle struct {
	width  float64
	height float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	base   float64
	height float64
}
type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
func (r Rectangle) Area() float64 {
	return r.height * r.width
}
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.height + rectangle.width)
}
