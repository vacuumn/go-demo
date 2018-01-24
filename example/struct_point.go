package example

import "math"

type Point struct {
	X, Y float64 // Upper case means exported
	name string
}

func (p *Point) Scale(s float64) {
	p.X *= s; p.Y *= s; // p is explicit
}
func (p Point) Abs() float64 {
	return math.Sqrt(p.X * p.X + p.Y * p.Y)
}
