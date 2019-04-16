package oop

import (
	"fmt"
	"golearn/nghiem/entities"
)

func Demo() {
	a := NewPoint(2, 3)
	b := NewPoint(8, 23)
	c := NewPoint(9, 78)
	fmt.Println("Chu vi abc: ", Perimeter(a, b, c))
}
func Demo1() {
	s1 := entities.Student{
		Id:    "s1",
		Human: entities.Human{"Student 1", "Female"},
		Score: 8.2112,
	}
	fmt.Println(s1.ToString())
}

func Demo2() {
	s := entities.NewStudent("s2", "Student 2", "Male", 5.645)
	fmt.Println(s.ToString())
}
func Demo3() {
	s1 := entities.NewStudent("s1", "Student 1", "male", 8.89)
	e1 := entities.NewEmployee("e1", "Employee 1", "female", 234.4)
	fmt.Println("========================")
	fmt.Println(s1.ToString())
	fmt.Println("========================")
	fmt.Println(e1.ToString())
}

func Demo4() {
	var animal entities.Animal
	cat := entities.Cat{}
	duck := entities.Duck{}

	animal = cat
	fmt.Println("Animal sound:", animal.Sound())
	animal = duck
	fmt.Println("Animal sound:", animal.Sound())
}

func Demo5() {
	cat1 := entities.Cat{}
	cat2 := entities.Cat{}
	duck1 := entities.Duck{}
	duck2 := entities.Duck{}
	duck3 := entities.Duck{}

	animals := []entities.Animal{cat1, duck1, duck2, cat2, duck3}

	for _, animal := range animals {
		fmt.Println("Animal sound:", animal.Sound())
	}
}

func Demo6() {
	ht := entities.NewHinhTron(3.56)
	hv := entities.NewHinhVuong(4.2)
	hinhs := []entities.Hinh{ht, hv}
	for _, hinh := range hinhs {
		fmt.Println("------------------------------------")
		fmt.Println("Ten hinh: ", hinh.Ten())
		fmt.Println("Chu vi: ", hinh.ChuVi())
		fmt.Println("Dien tich: ", hinh.DienTich())
	}
}
