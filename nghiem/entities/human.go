package entities

import (
	"fmt"
)

type Human struct {
	Name   string
	Gender string
}

func (human Human) ToString() string {
	return fmt.Sprintf("Name: %s\nGender: %s", human.Name, human.Gender)
}
