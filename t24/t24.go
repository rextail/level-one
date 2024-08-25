package t24

import (
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p1 Point) Distance(p2 Point) float64 {
	x := p2.x - p1.x
	y := p2.y - p1.y
	return math.Sqrt(x*x + y*y)
}
