package main

import "fmt"

func main() {
	var sub Subject
	sub = &Proxy{}
	fmt.Println(sub.Do())
}

type Subject interface {
	Do() string
}

type RealSubject struct {
}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p *Proxy) Do() string {
	var res string
	res += "before,"
	res += p.real.Do()
	res += ",after"
	return res
}
