package main

import "fmt"

type I interface {
	M()
}

type Istring struct {
	S string
}

type Ivector struct {
	X, Y int
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (is Istring) M() {
	fmt.Println(is.S)
}

func (iv Ivector) M() {
	fmt.Println(iv.X, iv.Y)
}

func main() {
	var i I = Istring{"hello"}
	i.M()
	var v I = Ivector{3, 4}
	v.M()
}
