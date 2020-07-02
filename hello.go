package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int => Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("string => %q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("default => I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
