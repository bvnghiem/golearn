package oop

import (
	"fmt"
	"nghiem/entities"
)

func Demo1() {
	s1 := entities.Student{
		Id:    "s1",
		Name:  "Student 1",
		Age:   21,
		Score: 8.2112,
	}
	fmt.Println(s1.ToString())
}

func Demo2() {
	s := entities.NewStudent("s2", "Student 2", 23, 5.645)
	fmt.Println(s.ToString())
	s.SetAge1(25)
	fmt.Println(s.ToString())
	s.SetAge2(25)
	fmt.Println(s.ToString())
}

func Demo() {
	a := NewPoint(2, 3)
	b := NewPoint(8, 23)
	c := NewPoint(9, 78)
	fmt.Println("Chu vi abc: ", Perimeter(a, b, c))
}
