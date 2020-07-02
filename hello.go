package main

import (
	"fmt"
	"math/rand"

	"example.com/pi/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello GO"))
}
