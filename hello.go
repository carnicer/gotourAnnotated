package main

import "fmt"
import "math"

func powLimit(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g < %g (limit)\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(
		powLimit(3, 2, 10),
		powLimit(3, 3, 20),
	)
}
