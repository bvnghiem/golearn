package main

import (
	"fmt"
	"golearn/nghiem/oop"
	"strconv"
)

func demo1() {
	var age int
	var price float32
	var fullName string
	var status bool
	var key rune

	age = 30
	price = 5.1
	fullName = "Bui Van Nghiem"
	status = true
	key = 'D'
	fmt.Println("demo1")
	fmt.Println("================================================================")
	fmt.Println("Full name: ", fullName, ", age: ", age, ", price: ", price,
		", status: ", status)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %1.1f\n", price)
	fmt.Printf("status: %t\n", status)
	fmt.Printf("key: %c\n", key)
	fmt.Printf("key: %d\n", key)

	fmt.Println("================================================================")
}

func demo2() {

	var age = 30
	var price = 4.5
	var fullName = "BVNghiem"
	var status = false
	var key = 'C'

	fmt.Println("============================= Demo 2 ==============================")
	fmt.Println("Full name: ", fullName, ", age: ", age, ", price: ", price,
		", status: ", status)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %1.1f\n", price)
	fmt.Printf("status: %t\n", status)
	fmt.Printf("key: %c\n", key)
	fmt.Printf("key: %d\n", key)
	fmt.Println("================================================================")
}

func demo3() {
	var (
		age      = 20
		fullName = "bvnghiem"
		price    = 3.4
		status   = true
		key      = 'E'
	)
	fmt.Println("============================= Demo 3 ==============================")
	fmt.Println("Full name: ", fullName, ", age: ", age, ", price: ", price,
		", status: ", status)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %1.1f\n", price)
	fmt.Printf("status: %t\n", status)
	fmt.Printf("key: %c\n", key)
	fmt.Printf("key: %d\n", key)
	fmt.Println("================================================================")

}
func demo4() {
	var a1, a2, a3 = 4, 10, -2

	fmt.Println("============================= Demo 4 ==============================")
	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)
	fmt.Println("a3: ", a3)
	fmt.Println("================================================================")

}

func demo5() {
	a, b, c := 3, "test", true
	fmt.Println("============================= Demo 4 ==============================")
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("================================================================")
}

func demo6() {
	var a = 30
	var b float32 = 5.6
	var result = float32(a) + b
	fmt.Println("result: ", result)
	firstName, lastName := "Nghiem", "Bui"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)
	var s = "age: "
	var result2 = s + fmt.Sprintf("%d", a)
	fmt.Println(result2)
}

func demo7() {
	var a = 5
	if a > 3 {
		fmt.Println(" a > 3")
	} else {
		fmt.Println("a <= 3")
	}

	if b := 2; b%2 != 0 {
		fmt.Println(" b is odd")
	} else {
		fmt.Println("b is even")
	}
}

func demo8() {
	var menu = 2
	if menu == 1 {
		fmt.Println("menu 1 is selected")
	} else if menu == 2 {
		fmt.Println("menu 2 is selected")
	} else {
		fmt.Println("Invalid")
	}
}

func demo9() {
	var menu = 2
	switch menu {
	case 1:
		fmt.Println("menu 1 is selected")
	case 2:
		fmt.Println("menu 2 is selected")
	case 3:
		fmt.Println("menu 3 is selected")
	default:
		fmt.Println("Invalid")
	}

}
func demo10() {
	var a = 1
	switch a {
	case 1, 2, 3:
		fmt.Printf("menu %d is selected\n", a)
	default:
		fmt.Println("Invalid")
	}

}
func demo11() {
	var a = 12
	switch {
	case a >= 1 && a <= 10:
		fmt.Printf("1 >= %d <=10\n", a)
	case a > 10 && a <= 20:
		fmt.Printf("10 > %d <=20\n", a)
	default:
		fmt.Printf("%d > 20\n", a)
	}
}
func demo12() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%5d", i)
	}

}

func demo13() {
	var i = 1
	for i <= 10 {
		fmt.Printf("%5d", i)
		i++
	}
}

func demo14() {
	var i = 1
	for {
		fmt.Printf("%5d", i)
		i++
		if i > 10 {
			break
		}

	}
}

func Add1(a int, b int) int {
	return a + b
}

func Add2(a, b int) int {
	return a + b
}

func Add3(a, b int) (s int) {
	s = a + b
	return a + b
}

func cal1(a, b int) (int, int, int, float32) {
	return a + b, a - b, a * b, float32(a) / float32(b)
}

func cal2(a, b int) (r1 int, r2 int, r3 int, r4 float32) {
	r1 = a + b
	r2 = a - b
	r3 = a * b
	r4 = float32(a) / float32(b)
	return r1, r2, r3, r4

}

func parseBool(s string) {
	v, e := strconv.ParseBool(s)
	if e != nil {
		fmt.Println(s, "is not a bool value")
	} else {
		fmt.Println("value: ", v)
	}
}

func cal3(a, b int) (cv int, dt int) {
	dt = a * b
	cv = 2 * (a + b)
	return cv, dt
}

func cal4(m, n int) (se int, so int) {
	se = 0
	so = 0
	for m < n {
		if m%2 == 0 {
			se += m
		} else {
			so += m
		}
		m++
	}
	return se, so
}

func cal5(m, n int) (ce int, co int) {
	ce = 0
	co = 0
	for m < n {
		if m%2 == 0 {
			ce++
		} else {
			co++
		}
		m++
	}
	return ce, co
}

func main() {
	oop.Demo6()
}
