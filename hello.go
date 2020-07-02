package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	vertex1 := Vertex{1, 2}
	fmt.Println(vertex1)
	p := &vertex1
	p.X = 33
	fmt.Println(vertex1)
	(*p).X = 36
	fmt.Println(vertex1)
}

