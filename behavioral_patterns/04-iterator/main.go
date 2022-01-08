package main

import "fmt"

func main() {
	var agg Aggregate
	agg = &Numbers{1, 10}
	iter := agg.Iterator()
	for iter.First(); !iter.IsDone(); {
		fmt.Println(iter.Next())
	}
}

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

type NumbersIterator struct {
	numbers *Numbers
	next    int
}

func (ni *NumbersIterator) First() {
	ni.next = ni.numbers.start
}

func (ni *NumbersIterator) IsDone() bool {
	return ni.next > ni.numbers.end
}

func (ni *NumbersIterator) Next() interface{} {
	if !ni.IsDone() {
		next := ni.next
		ni.next++
		return next
	}
	return nil
}

type Numbers struct {
	start, end int
}

func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		next:    n.start,
	}
}

type Aggregate interface {
	Iterator() Iterator
}
