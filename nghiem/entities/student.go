package entities

import "fmt"

type Student struct {
	Id    string
	Human Human
	Score float32
}

func NewStudent(id, name string, gender string, score float32) Student {
	return Student{id, Human{name, gender}, score}
}

func (s Student) ToString() string {
	return fmt.Sprintf("Id: %s\n%s\nScore: %.2f", s.Id, s.Human.ToString(), s.Score)
}
