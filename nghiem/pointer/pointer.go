package pointer

import (
	"fmt"
)

func Demo1() {
	a := 5
	fmt.Println("a: ", a)
	fmt.Println("addr a: ", &a)

	p := &a
	fmt.Println("addr p: ", p)
	fmt.Println("value p: ", *p)
	*p = 10
	fmt.Println("value a: ", a)

	q := &a
	fmt.Println("addr q: ", q)
	fmt.Println("value q: ", *q)
}

func Demo2() {
	a, b := 5, 10
	p, q := &a, &b

	fmt.Println("addr p: ", p)
	fmt.Println("value p: ", *p)

	fmt.Println("addr q: ", q)
	fmt.Println("value q: ", *q)
}

func Demo3() {
	a := 5
	fmt.Println("a:", a)
	changeValue1(a)
	fmt.Println("a:", a)
	changeValue2(&a)
	fmt.Println("a:", a)
}
func changeValue1(a int) {
	a = 6
}

func changeValue2(p *int) {
	*p = 7
}

func Demo4() {
	a, b := 5, 10
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	swap(&a, &b)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
func swap(p, q *int) {
	t := *q
	*q = *p
	*p = t
}
