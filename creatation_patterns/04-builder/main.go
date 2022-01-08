package main

import "fmt"

func main() {
	intB, stringB := &IntBuilder{}, &StringBuilder{}
	d1, d2 := NewDirector(intB), NewDirector(stringB)
	d1.Construct()
	d2.Construct()
	fmt.Println(intB.GetResult())
	fmt.Println(stringB.GetResult())
}

type Builder interface {
	Part1()
	Part2()
	Part3()
}

func NewDirector(builder Builder) *Director {
	return &Director{builder}
}

type Director struct {
	Builder
}

func (d *Director) Construct() {
	d.Builder.Part1()
	d.Builder.Part2()
	d.Builder.Part3()
}

type IntBuilder struct {
	result int64
}

func (this_ *IntBuilder) Part1() {
	this_.result += 1
}

func (this_ *IntBuilder) Part2() {
	this_.result += 2
}

func (this_ *IntBuilder) Part3() {
	this_.result += 3
}

func (this_ *IntBuilder) GetResult() int64 {
	return this_.result
}

type StringBuilder struct {
	result string
}

func (this_ *StringBuilder) Part1() {
	this_.result += "1"
}

func (this_ *StringBuilder) Part2() {
	this_.result += "2"
}

func (this_ *StringBuilder) Part3() {
	this_.result += "3"
}

func (this_ *StringBuilder) GetResult() string {
	return this_.result
}
