package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	var slice []int = primes[1 : 4]
	fmt.Println(slice)
	slice2 := primes[2 : 4]
	fmt.Println(slice2)
}
