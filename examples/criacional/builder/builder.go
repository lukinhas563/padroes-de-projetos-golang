package main

import "fmt"

type Car struct {
	Engine string
	Wheels int
	Color  string
}

type CarBuilder struct {
	car Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (b *CarBuilder) SetEngine(engine string) *CarBuilder {
	b.car.Engine = engine
	return b
}

func (b *CarBuilder) SetWheels(wheels int) *CarBuilder {
	b.car.Wheels = wheels
	return b
}

func (b *CarBuilder) SetColor(color string) *CarBuilder {
	b.car.Color = color
	return b
}

func (b *CarBuilder) Build() Car {
	return b.car
}

func main() {
	builder := NewCarBuilder()

	car := builder.SetEngine("V8").
		SetWheels(4).
		SetColor("Vermelho").
		Build()

	fmt.Printf("Carro criado: %+v\n", car)
}
