package main

import "fmt"

type AIN interface {
	F(int) string
}

type A struct {
	F func(int) string
}

type AA struct {
	A
}

func (a *AA) F(int2 int) string {
	return "ok"
}

func NewAA() AIN {
	a := new(AA)
	a.A.F = a.F
	return a
}

func main() {
	aa := NewAA()
	if aa.F != nil {
		fmt.Println(aa.F(0))
	}
}
