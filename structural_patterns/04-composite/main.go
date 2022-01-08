package main

import "fmt"

const (
	LeafNode = iota
	CompositeNode
)

func main() {
	root := NewComponent(CompositeNode, "root")

	c1, c2, c3 := NewComponent(CompositeNode, "c1"),  NewComponent(CompositeNode, "c2"),  NewComponent(CompositeNode, "c3")
	l1, l2, l3 := NewComponent(LeafNode, "l1"),  NewComponent(LeafNode, "l2"),  NewComponent(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c3.AddChild(l3)

	root.Print("-")
}

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = &Leaf{}
	case CompositeNode:
		c = &Composite{
			children: []Component{},
		}
	default:
		c = &Leaf{}
	}
	c.SetName(name)
	return c
}

type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(p Component) {
	c.parent = p
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(cp Component) {}

func (c *component) Print(string) {}

type Leaf struct {
	component
}

func (l *Leaf) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, l.Name())
}

type Composite struct {
	component
	children []Component
}

func (c *Composite) AddChild(cp Component) {
	cp.SetParent(c)
	c.children = append(c.children, cp)
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, c.Name())
	pre += " "
	for _, child := range c.children {
		child.Print(pre)
	}
}
