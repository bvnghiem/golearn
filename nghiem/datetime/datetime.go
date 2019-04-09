package datetime

import (
	"fmt"
	"time"
)

func Demo1() {
	today := time.Now()
	fmt.Println(today)
	fmt.Println("Year: ", today.Year())

	fmt.Println("Month: ", today.Month())

	fmt.Println("Month: ", int(today.Month()))

	fmt.Println("Day: ", today.Day())

	fmt.Println("Week Day: ", today.Weekday())
}

func Demo2() {
	today := time.Now()

	fmt.Println(today.Format("02/01/2006 15:04:05 "))

	s := "20/10/2019"
	d, e := time.Parse("02/01/2006", s)
	if e != nil {
		fmt.Println(2)
	} else {
		fmt.Println(d.Format("02-01-2006"))
	}
}

func Demo3() {
	today := time.Now()
	fmt.Println(today.Format("02/01/2006"))
	d1 := today.AddDate(0, 0, 3)
	fmt.Println(d1.Format("02/01/2006"))
	d2 := today.Add(4 * 24 * time.Hour)
	fmt.Println(d2.Format("02/01/2006"))

	result := d2.Sub(d1).Hours()

	fmt.Println("d2 - d1 =", result)
}

func Exer() {
	s := "06/04/1989"
	fmt.Println("Today is your birthday: ", isBirthday(s))
	b, y := isOver18(s)
	fmt.Println("You are over 18:", b)
	fmt.Println("You are", y, "now")
}

func isBirthday(s string) bool {
	today := time.Now()
	bd, err := time.Parse("02/01/2006", s)
	if err != nil {
		fmt.Println(s, "is not a date value")
		return false
	} else {
		fmt.Println("Your birthday is", bd.Day(), bd.Month(), bd.Year())
		return bd.Day() == today.Day() && int(bd.Month()) == int(today.Month())
	}
}

func isOver18(s string) (bool, int) {
	bd, err := time.Parse("02/01/2006", s)
	if err != nil {
		fmt.Println(s, "is not a date value")
		return false, 0
	} else {
		today := time.Now()
		oldD := today.AddDate(-18, 0, 0)
		return bd.Year() < oldD.Year(), today.Year() - bd.Year()
	}
}
