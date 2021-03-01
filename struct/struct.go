package main

import (
	"fmt"
)

type F interface {
	f()
}

type S1 struct {
	data int32
}

func (s S1) f() {
	s.data = 1
	fmt.Println("S1 Data is ", s.data)
}

type S2 struct {
	data int32
}

func (s *S2) f() {
	s.data = 2
	fmt.Println("S2 Data is ", s.data)
}

func main() {
	var s1 = S1{data: 0}
	var s2 = S2{data: 0}

	var f1 F = s1
	var f2 F = &s2

	f1.f()
	f2.f()

	fmt.Println("S1 Data is ", s1.data)
	fmt.Println("S2 Data is ", s2.data)
}
