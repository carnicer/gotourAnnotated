package main

import "fmt"

func fibonacci() func() int {
	prev := 0
	prevprev := 0
	fmt.Println("fibonacci, initialized (0)")
	return func() int {
		fibonacci := prev + prevprev
		/*
		fmt.Printf("prevprev: %v => %v\n", prevprev, prev)
		fmt.Printf("prev: %v => %v\n", prev, fibonacci)
		fmt.Println("--")
		*/
		prevprev = prev
		// first time, prev == fibonacci == 0
		if prev == fibonacci {
			prev = 1
		} else {
			prev = fibonacci
		}
		return fibonacci
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 13; i++ {
		fmt.Println(
			"#",
			i,
			"=>",
			f(),
		)
		//fmt.Println("----")
	}
}
