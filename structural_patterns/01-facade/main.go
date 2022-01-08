package main

import "fmt"

func main() {
	a := NewAPI()
	fmt.Println(a.Test())
}

func NewAPI() API {
	return &apiIml{
		a: &aModuleIml{},
		b: &bModuleIml{},
	}
}

type API interface {
	Test() string
}

type apiIml struct {
	a AModuleApi
	b BModuleApi
}

func (this_ *apiIml) Test() string {
	return fmt.Sprintf("%s, %s", this_.a.TestA(), this_.b.TestB())
}

type AModuleApi interface {
	TestA() string
}
type aModuleIml struct{}

func (a *aModuleIml) TestA() string {
	return "a is running"
}

type BModuleApi interface {
	TestB() string
}

type bModuleIml struct{}

func (b *bModuleIml) TestB() string {
	return "b is running"
}
