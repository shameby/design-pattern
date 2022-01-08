package main

import (
	"fmt"
	"sync"
)

func main() {
	s1, s2 := GetInstance(), GetInstance()
	fmt.Println(s1 == s2)
}

var s *Singleton

type Singleton struct {
}

func GetInstance() *Singleton {
	once := sync.Once{}
	once.Do(func() {
		s = &Singleton{}
	})
	return s
}
