package square

import (
	"nghiem/pack2"
)

func Area(a float64) float64 {
	return pack2.Mul(a, a)
}

func Perimeter(a float64) float64 {
	return pack2.Mul(a, 4)
}
