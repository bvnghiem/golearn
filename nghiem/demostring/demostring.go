package demostring

import (
	"fmt"
	"sort"
	"strings"
)

func Demo1() {
	s := "abcdef"
	fmt.Println("len:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func Demo2() {
	s1 := "abc"
	s1 = strings.ToUpper(s1)
	fmt.Println(s1)
	s2 := "ACVBCVYASDA"
	s2 = strings.ToLower(s2)
	fmt.Println(s2)

}

func Demo3() {
	s1 := "computer"
	s2 := "com"
	r1 := strings.HasPrefix(s1, s2)
	fmt.Println(r1)

	s3 := "ter"
	r2 := strings.HasSuffix(s1, s3)
	fmt.Println(r2)

	s4 := "put"
	r3 := strings.Contains(s1, s4)
	fmt.Println(r3)
}

func Exer() {
	spList := []string{"zip", "book", "case", "monitor", "keyboard"}
	fmt.Println(spList)
	fmt.Println("Find 'a' in above list: ", search(spList, "a"))
	fmt.Println("Sort this list descending:")
	sortDes(spList)
	fmt.Println(spList)
}

func search(sList []string, key string) (results []string) {

	for _, s := range sList {
		if strings.Contains(s, key) {
			results = append(results, s)
		}
	}
	return results
}

func sortDes(sList []string) {
	sort.Slice(sList, func(i, j int) bool {
		return strings.Compare(sList[i], sList[j]) > 0
	})
}
