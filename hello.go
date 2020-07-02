package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	ss := s[:0]
	printSlice(ss)

	// Extend its length.
	ss = s[:4]
	printSlice(ss)

	// Drop its first two values.
	ss = s[2:]
	printSlice(ss)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
