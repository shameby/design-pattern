package main

import "fmt"

func main() {
	s := &Subject{observers: []Observer{}}
	s.Attach(&Reader{"wrh"})
	s.Attach(&Reader{"yby"})
	s.Attach(&Reader{"wxh"})
	s.Attach(&Reader{"wbd"})
	s.UpdateCtx("I'm coming")
}

type Subject struct {
	observers []Observer
	context   string
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateCtx(ctx string) {
	s.context = ctx
	s.notify()
}

type Observer interface {
	Update(*Subject)
}

type Reader struct {
	name string
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s recieve msg %s \n", r.name, s.context)
}