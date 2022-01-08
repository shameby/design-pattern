package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	p := &Parser{}
	p.Parse("1 + 2 + 3 - 2 + 5 - 1")
	fmt.Println(p.Result().Interpret())
}

type Node interface {
	Interpret() int64
}

type ValNode struct {
	val int64
}

func (n *ValNode) Interpret() int64 {
	return n.val
}

type AddNode struct {
	left, right Node
}

func (n *AddNode) Interpret() int64 {
	return n.left.Interpret() + n.right.Interpret()
}

type MinusNode struct {
	left, right Node
}

func (n *MinusNode) Interpret() int64 {
	return n.left.Interpret() - n.right.Interpret()
}

type Parser struct {
	exp   []string
	index int
	prev  Node
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")

	for p.index < len(p.exp) {
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddNode()
		case "-":
			p.prev = p.newMinusNode()
		default:
			p.prev = p.newValNode()
		}
	}
}

func (p *Parser) newAddNode() Node {
	p.index++
	return &AddNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newMinusNode() Node {
	p.index++
	return &MinusNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newValNode() Node {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValNode{
		val: int64(v),
	}
}

func (p *Parser) Result() Node {
	return p.prev
}
