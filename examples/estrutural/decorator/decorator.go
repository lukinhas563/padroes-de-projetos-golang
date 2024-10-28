package main

import "fmt"

type IProduct interface {
	GetPrice() float32
	GetName() string
}

type tshirt struct {
	name  string
	price float32
}

func NewTshirt(name string, price float32) IProduct {
	return &tshirt{
		name:  name,
		price: price,
	}
}

func (t *tshirt) GetName() string {
	return t.name
}

func (t *tshirt) GetPrice() float32 {
	return t.price
}

// Decorator
type ProductDecorator struct {
	product IProduct
}

func NewProductDecorator(product IProduct) IProduct {
	return &ProductDecorator{
		product: product,
	}
}

func (d *ProductDecorator) GetName() string {
	return "printed " + d.product.GetName()
}

func (d *ProductDecorator) GetPrice() float32 {
	return d.product.GetPrice() + 10
}

func main() {
	newProduct := NewTshirt("t-shirt", 29.99)
	newProductDecorator := NewProductDecorator(newProduct)
	newProductDecoratorTwo := NewProductDecorator(newProductDecorator)

	fmt.Println(newProductDecorator.GetName())
	fmt.Println(newProductDecorator.GetPrice())

	fmt.Println(newProductDecoratorTwo.GetName())
	fmt.Println(newProductDecoratorTwo.GetPrice())
}
