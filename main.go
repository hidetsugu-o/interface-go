package main

import "fmt"

type Dog struct{}

func (Dog) Speak() { fmt.Println("bowwow") }
func (Dog) Run()   { fmt.Println("yay, I'm running") }

type Cat struct{}

func (Cat) Speak() { fmt.Println("mewmew") }
func (Cat) Run()   { fmt.Println("shut up") }

type Robot struct{}

func (Robot) Run() { fmt.Println("geek geek") }

type Runner interface {
	Run()
}

func main() {
	Race([]Runner{Robot{}, Dog{}, Cat{}})
}

func Race(runners []Runner) {
	for _, runner := range runners {
		runner.Run()
	}
}
