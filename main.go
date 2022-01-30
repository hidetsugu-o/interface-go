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

func main() {
	vs := []Stringfy{
		&Person{Name: "gopher", Age: 30, Gender: "男"},
		&Animal{Name: "自分ツッコミくま", Kind: "くま"},
	}

	for _, v := range vs {
		fmt.Println(v.Tostring())
	}
}
