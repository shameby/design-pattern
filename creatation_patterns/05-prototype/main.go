package main

import "fmt"

func main() {
	mgr := NewProtoTypMgr()
	t1 := &TestTyp{
		name: "wrh",
	}
	mgr.Set("t1", t1)

	t2 := mgr.Get("t1").Clone()
	fmt.Println(t2.(*TestTyp).name)
}

type Cloneable interface {
	Clone() Cloneable
}

type ProtoTypeManager struct {
	prototypes map[string]Cloneable
}

func NewProtoTypMgr() *ProtoTypeManager {
	return &ProtoTypeManager{
		prototypes: map[string]Cloneable{},
	}
}

func (this_ *ProtoTypeManager) Get(name string) Cloneable {
	return this_.prototypes[name]
}

func (this_ *ProtoTypeManager) Set(name string, typ Cloneable) {
	this_.prototypes[name] = typ
}

type TestTyp struct {
	name string
}

func (t *TestTyp) Clone() Cloneable {
	tc := *t
	return &tc
}

