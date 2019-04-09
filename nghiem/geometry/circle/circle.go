package circle

import (
	"math"
)

func Area(rad float64) float64 {
	return math.Pi * rad * rad
}

func Perimeter(rad float64) float64 {
	return 2 * math.Pi * rad
}
