package main

import "fmt"

func main() {
	innerIfc := NewInnerIfc()
	t := NewTar(innerIfc)
	fmt.Println(t.Request())
}

type InnerIfc interface {
	SpecificRequest() string
}

type innerStruct struct{}

func (is *innerStruct) SpecificRequest() string {
	return "inner method"
}

type adapter struct {
	InnerIfc
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}

func NewInnerIfc() InnerIfc {
	return &innerStruct{}
}

type Target interface {
	Request() string
}

func NewTar(adp InnerIfc) Target {
	return &adapter{
		adp,
	}
}
