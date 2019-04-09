package oop

import (
	"math"
)

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}
func distance(a, b Point) float64 {
	return math.Sqrt(float64(b.X-a.X)*float64(b.X-a.X) + float64(b.Y-a.Y)*float64(b.Y-a.Y))
}

func Perimeter(a, b, c Point) float64 {
	return distance(a, b) + distance(b, c) + distance(a, c)
}
