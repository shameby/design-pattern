package main

import "fmt"

func main() {
	proMgr, depMgr, generalMgr := &RequestChain{Manger:&ProjectMgr{}}, &RequestChain{Manger:&DepMgr{}}, &RequestChain{Manger:&GeneralMgr{}}
	proMgr.successor = depMgr
	depMgr.successor = generalMgr

	proMgr.HandleFeeRequest("bob", 150000)
}

type Manger interface {
	HaveRight(money int64) bool
	HandleFeeRequest(name string, money int64) bool
}

type RequestChain struct {
	Manger
	successor *RequestChain
}

func (r *RequestChain) SetSuccessor(m *RequestChain) {
	r.successor = m
}

func (r *RequestChain) HandleFeeRequest(name string, money int64) bool {
	if r.Manger.HaveRight(money) {
		return r.Manger.HandleFeeRequest(name, money)
	}
	if r.successor != nil {
		return r.successor.HandleFeeRequest(name, money)
	}
	return false
}

func (r *RequestChain) HaveRight(money int64) bool {
	return true
}

type ProjectMgr struct {}

func (*ProjectMgr) HaveRight(money int64) bool {
	return money < 500
}

func (*ProjectMgr) HandleFeeRequest(name string, money int64) bool {
	if name == "bob" {
		fmt.Printf("Project manager permit %s %d fee request.\n", name, money)
		return true
	}
	fmt.Printf("Project manager don't permit %s %d fee request.\n", name, money)
	return false
}

type DepMgr struct {}

func (*DepMgr) HaveRight(money int64) bool {
	return money < 5000
}

func (*DepMgr) HandleFeeRequest(name string, money int64) bool {
	if name == "tom" {
		fmt.Printf("Dep manager permit %s %d fee request.\n", name, money)
		return true
	}
	fmt.Printf("Dep manager don't permit %s %d fee request.\n", name, money)
	return false
}

type GeneralMgr struct {}

func (*GeneralMgr) HaveRight(money int64) bool {
	return true
}

func (*GeneralMgr) HandleFeeRequest(name string, money int64) bool {
	if name == "ada" {
		fmt.Printf("General manager permit %s %d fee request.\n", name, money)
		return true
	}
	fmt.Printf("General manager don't permit %s %d fee request.\n", name, money)
	return false
}