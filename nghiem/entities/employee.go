package entities

import (
	"fmt"
)

type Employee struct {
	Id     string
	Human  Human
	Salary float32
}

func NewEmployee(id string, name string, gender string, salary float32) Employee {
	return Employee{id, Human{name, gender}, salary}
}

func (e Employee) ToString() string {
	return fmt.Sprintf("Id: %s\n%s\nSalary: %.2f", e.Id, e.Human.ToString(), e.Salary)
}
