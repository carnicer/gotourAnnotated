package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	b2 := b[:2]
	printSlice("b2", b2)

	b25 := b2[2:5]
	printSlice("b25", b25)

	b = b[:6]
	b26 := b2[2:6]
	printSlice("b26", b26)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
