package main

import "fmt"

type I interface {
	M()
}

type Istring struct {
	S string
}

func (t *Istring) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *Istring
	i = t
	describe(i)
	i.M()

	fmt.Println("--")

	i = &Istring{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
