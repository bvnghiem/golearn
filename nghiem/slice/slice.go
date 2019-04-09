package slice

import (
	"fmt"
	"sort"
)

func Demo1() {

	var a = [5]int{8, 5, -1, -2, 4}
	fmt.Println(a)
	s1 := a[1:4]
	fmt.Println(s1)
	s1[0] = 9
	fmt.Println(a)

	s2 := a[:3]
	fmt.Println(s2)
	s3 := a[1:]
	fmt.Println(s3)
	s4 := a[:]
	fmt.Println(s4)
}

func Demo2() {
	a := []int{3}
	fmt.Println(a)
	fmt.Println(cap(a))
	a = append(a, 0)
	fmt.Println(a)
	fmt.Println(cap(a))
	a = append(a, 12, 7)
	fmt.Println(a)
	fmt.Println(cap(a))
	b := []int{-5, 6}
	a = append(a, b...)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}

func Demo3() {

	var a = [5]int{8, 5, -1, -2, 4}
	fmt.Println("a:", a)
	s := a[0:2]
	fmt.Println("s:", s)
	fmt.Println("len: ", len(s), ", cap: ", cap(s))

	for i := 1; i <= 10; i++ {
		s = append(s, i*s[0])
		fmt.Println("s:", s)
		fmt.Println("len: ", len(s), ", cap: ", cap(s))
	}
}

func Demo4() {
	s := make([]int, 1)
	fmt.Println("s:", s)
	fmt.Println("len: ", len(s), ", cap: ", cap(s))
	s[0] = 3
	for i := 1; i <= 10; i++ {
		s = append(s, i*s[0])
		fmt.Println("s:", s)
		fmt.Println("len: ", len(s), ", cap: ", cap(s))
	}
}

func Demo5() {
	a := []int{1, 2, 3}
	display(a)

	b := [5]int{1, 2, 3, 4, 5}
	display(b[:])

	changeValue1(b)
	display(b[:])

	changeValue2(b[:])
	display(b[:])
}

func display(a []int) {
	for _, v := range a {
		fmt.Printf("%5d", v)
	}
	fmt.Println()
	fmt.Println("---------------------------")
}

func changeValue1(a [5]int) {
	a[2] = 30
}

func changeValue2(a []int) {
	a[2] = 30
}

func Demo6() {

	a := []int{10, -2, 32, 14, 5}
	display(a)
	sort.Ints(a)
	display(a)
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	display(a)
}
