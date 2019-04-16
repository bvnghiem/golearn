package entities

import (
	"math"
)

type HinhTron struct {
	BanKinh float64
}

func NewHinhTron(bankinh float64) HinhTron {
	return HinhTron{bankinh}
}

func (ht HinhTron) Ten() string {
	return "Hinh tron"
}

func (ht HinhTron) ChuVi() float64 {
	return math.Pi * 2 * ht.BanKinh
}

func (ht HinhTron) DienTich() float64 {
	return math.Pi * ht.BanKinh * ht.BanKinh
}
