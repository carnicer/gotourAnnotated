package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	vertex1 := Vertex{1, 2}
	fmt.Println(vertex1)
}

