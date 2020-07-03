package main

import "fmt"

type I interface {
	M()
}

type Istring struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (is Istring) M() {
	fmt.Println(is.S)
}

func main() {
	var i I = Istring{"hello"}
	i.M()
}
