package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)

	a = f  // a MyFloat implements Abser
	fmt.Printf("abs of f (MyFloat): %v\n", f.Abs())
	fmt.Printf("abs of a (interface): %v\n", a.Abs())

	v := Vertex{3, 4}
	a = &v // a *Vertex implements Abser
	fmt.Printf("abs of v (Vertex*): %v\n", v.Abs())
	fmt.Printf("abs of a (interface): %v\n", a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
