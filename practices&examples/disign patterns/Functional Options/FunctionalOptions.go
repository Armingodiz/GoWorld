package main

import "fmt"

type Car struct {
	color   string
	hasRoof bool
	price   int
	speed   int
}

type CarOptions func(*Car)

func SetColor(color string) CarOptions {
	return func(c *Car) {
		c.color = color
	}
}

func WithOutRoof() CarOptions {
	return func(c *Car) {
		c.hasRoof = false
	}
}

func SetPrice(price int) CarOptions {
	return func(c *Car) {
		c.price = price
	}
}

func SetSpeed(speed int) CarOptions {
	return func(c *Car) {
		c.speed = speed
	}
}

func NewCar(changes ...CarOptions) *Car {
	defaultCar := &Car{
		color:   "black",
		hasRoof: true,
		price:   2000,
		speed:   180,
	}
	for _, change := range changes {
		change(defaultCar)
	}
	return defaultCar
}

func main() {
  car := NewCar(SetColor("white"),SetSpeed(200),WithOutRoof())
  fmt.Printf("%+v\n", car)
  // &{color:white hasRoof:false price:2000 speed:200}
}
