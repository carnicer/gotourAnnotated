package main

import "fmt"

func main() {
	sum := 1
	for {
		fmt.Printf("sum:%v\n", sum)
		sum += sum
		if sum > 1000 {
			break
		}
	}
	fmt.Println(sum)
}
