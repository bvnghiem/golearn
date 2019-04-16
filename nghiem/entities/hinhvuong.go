package entities

type HinhVuong struct {
	Canh float64
}

func NewHinhVuong(canh float64) HinhVuong {
	return HinhVuong{canh}
}

func (hv HinhVuong) Ten() string {
	return "Hinh vuong"
}

func (hv HinhVuong) ChuVi() float64 {
	return 4 * hv.Canh
}

func (hv HinhVuong) DienTich() float64 {
	return hv.Canh * hv.Canh
}
