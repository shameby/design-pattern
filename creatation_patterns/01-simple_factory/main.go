package main

import "fmt"

func main() {
	t1, t2 := NewApi(ApiTypHi), NewApi(ApiTypHello)
	fmt.Println(t1.Say("wrh"))
	fmt.Println(t2.Say("wrh"))
}

type ApiTyp int8

const (
	ApiTypHi = iota
	ApiTypHello
)
type API interface {
	Say(name string) string
}

func NewApi(t ApiTyp) API {
	switch t {
	case ApiTypHi:
		return &HiApi{}
	case ApiTypHello:
		return &HelloApi{}
	}
	return nil
}

type HiApi struct {
}

func (this_ *HiApi) Say(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
type HelloApi struct {
}

func (this_ *HelloApi) Say(name string) string {
	return fmt.Sprintf("Hello %s", name)
}