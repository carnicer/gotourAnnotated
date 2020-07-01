package main

import (
    "fmt"

    "example.com/pi/hello/morestrings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println("hello world")
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello GO"))
}

