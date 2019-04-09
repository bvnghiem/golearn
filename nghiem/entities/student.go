package entities

import (
	"fmt"
)

type Student struct {
	Id    string
	Name  string
	Age   int
	Score float64
}

func NewStudent(id, name string, age int, score float64) Student {
	return Student{id, name, age, score}
}

func (s Student) ToString() string {
	return fmt.Sprintf("id: %s, name: %s, age: %d, score: %.2f",
		s.Id, s.Name, s.Age, s.Score)
}

func (s Student) SetAge1(age int) {
	s.Age = age
}
func (s *Student) SetAge2(age int) {
	s.Age = age
}
