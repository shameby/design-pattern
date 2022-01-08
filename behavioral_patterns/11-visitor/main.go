package main

import "fmt"

func main() {
	ec, ic := &EnterpriseCustomer{"tencent"}, &IndividualCustomer{"wrh"}
	sv, av := &ServiceRequestVisitor{}, &AnalysisVisitor{}
	cc := &CustomerCol{}
	cc.Add(ec)
	cc.Add(ic)
	cc.Accept(sv)
	cc.Accept(av)
}

type CustomerCol struct {
	customers []Customer
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol) Accept(visitor Visitor) {
	for _, cus := range c.customers {
		cus.Accept(visitor)
	}
}

type Visitor interface {
	Visit(Customer)
}

type Customer interface {
	Accept(Visitor)
}

type EnterpriseCustomer struct {
	name string
}

func (e *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(e)
}

type IndividualCustomer struct {
	name string
}

func (e *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(e)
}

type ServiceRequestVisitor struct {}

func (v *ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s \n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s \n", c.name)
	}
}

type AnalysisVisitor struct {}

func (v *AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s \n", c.name)
	case *IndividualCustomer:
		fmt.Printf("will not analysis individual customer %s \n", c.name)
	}
}