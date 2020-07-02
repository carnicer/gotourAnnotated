package main

import "fmt"

func adder() func(int) int {
	sum := 0
	fmt.Println("adder, initialized (0)")
	return func(x int) int {
		sum0 := sum
		sum += x
		fmt.Printf("sum: %v +%v => %v\n", sum0, x, sum)
		return sum
	}
}

func main() {
	//pos, neg := adder(), adder()
	pos := adder()
	for i := 0; i < 5; i++ {
		fmt.Println(
			"i:",
			i,
			"--",
			pos(i),
			pos(-2*i),
		)
		fmt.Println("----")
	}
}
