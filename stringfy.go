package main

import (
	"fmt"
)

type Stringfy interface {
	Tostring() string
}

// 人間
type Person struct {
	Name   string
	Age    int
	Gender string
}

func (person *Person) Tostring() string {
	return fmt.Sprintf("%s(%d歳:%s)", person.Name, person.Age, person.Gender)
}

// 動物
type Animal struct {
	Name string
	Kind string
}

func (animal *Animal) Tostring() string {
	return fmt.Sprintf("%s[%s]", animal.Name, animal.Kind)
}
