package main

import "fmt"

func main() {
	var c Component = &ConcreteComponent{}

	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 5)

	fmt.Println(c.Calc())
}

type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 1
}

type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		c,
		num,
	}
}

func (ad *AddDecorator) Calc() int {
	return ad.Component.Calc() + ad.num
}

type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		c,
		num,
	}
}

func (md *MulDecorator) Calc() int {
	return md.Component.Calc() * md.num
}