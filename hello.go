package main

import "fmt"
import "math"

func powLimit(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		powLimit(3, 2, 10),
		powLimit(3, 3, 20),
	)
}
