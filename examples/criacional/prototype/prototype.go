package main

import "fmt"

type Shape interface {
	Clone() Shape
	GetInfo() string
}

type Circle struct {
	Radius int
	Color  string
}

func (c *Circle) Clone() Shape {
	return &Circle{
		Radius: c.Radius,
		Color:  c.Color,
	}
}
func (c *Circle) GetInfo() string {
	return fmt.Sprintf("Císculo com radio %d e cor %s", c.Radius, c.Color)
}

func main() {
	originalCircle := &Circle{Radius: 5, Color: "Vermelho"}
	cloneCircle := originalCircle.Clone()

	fmt.Println("Original:", originalCircle.GetInfo())
	fmt.Println("Clone:", cloneCircle.GetInfo())
}
