package array

import (
	"fmt"
)

func Demo1() {
	var a [5]int
	a[0] = 0
	a[1] = 1
	a[2] = 2
	a[3] = 3
	a[4] = 4

	fmt.Println(a)
	fmt.Println("len: ", len(a))
	fmt.Println("value")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println()
	for i, v := range a {
		fmt.Println("a[", i, "]=", v)
	}

	fmt.Println()
	for _, v := range a {
		fmt.Println(v)
	}
}

func Demo2() {
	var a = [5]int{3, 5, -1, 0, 4}
	names := [3]string{"name1", "name2", "name3"}

	for _, v := range a {
		fmt.Println(v)
	}
	for _, v := range names {
		fmt.Println(v)
	}
}

func Demo3() {
	a := [...]int{2, 4, 54}
	fmt.Println("len: ", len(a))
	for _, v := range a {
		fmt.Println(v)
	}
}

func display(a [5]int) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func Demo4() {

	var a = [5]int{3, 5, -1, 0, 4}
	display(a)
}

func sum(a [5]int) (ch, le, am, duong int) {
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			ch += a[i]
		} else {
			le += a[i]
		}
		if a[i] >= 0 {
			duong += a[i]
		} else {
			am += a[i]
		}
	}
	return ch, le, am, duong
}

func count(a [5]int) (ch, le, am, duong int) {
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			ch++
		} else {
			le++
		}
		if a[i] >= 0 {
			duong++
		} else {
			am++
		}
	}

	return ch, le, am, duong
}

func minMax(a [5]int) (min, max int) {
	for i := 0; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}

		if a[i] > max {
			max = a[i]
		}
	}

	return min, max

}
func sort(a [5]int) [5]int {

	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				s := a[i]
				a[i] = a[j]
				a[j] = s
			}
		}
	}
	return a
}

func Demo5() {

	var a = [5]int{3, 3, -1, -2, 4}
	fmt.Println(a)
	fmt.Println(sum(a))
	fmt.Println(count(a))
	fmt.Println(minMax(a))
	fmt.Println(sort(a))
}
