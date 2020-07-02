package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// overflow when compiling =>
	// fmt.Printf("Small type:%T, Big type:%T\n", Small, Big)
	fmt.Printf("Small type:%T, value:%v\n", Small, Small)
	fmt.Printf("needInt(Small):%v\n", needInt(Small))
	// overflow when compiling =>
	// fmt.Printf("needInt(Big):%v\n", needInt(Big))
	fmt.Printf("needFloat(Small):%v\n", needFloat(Small))
	fmt.Printf("needFloat(Big):%v\n", needFloat(Big))
}
