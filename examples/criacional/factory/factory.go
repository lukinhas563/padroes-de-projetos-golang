package main

import "fmt"

type IProduct interface {
	sayHi()
}

type person struct {
	name string
}

func (p *person) sayHi() {
	fmt.Println("Hi")
}

func NewPerson(name string) IProduct {
	return &person{
		name: name,
	}
}

func main() {
	person := NewPerson("Pessoa 1")

	person.sayHi()
}
