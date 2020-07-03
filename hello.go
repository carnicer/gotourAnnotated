package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	fmt.Println("describe i (type &T) ...")
	describe(i)
	fmt.Println("i.M: ...")
	i.M()

	fmt.Println("--")

	i = F(math.Pi)
	fmt.Println("describe i (type F) ...")
	describe(i)
	fmt.Println("i.M: ...")
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
