package main

import "fmt"

func sum(s []int, c chan int) {
	size := len(s)
	sum := 0
	for i, v := range s {
		sum += v
		fmt.Printf("sum %d/%d items\n", i, size)
	}
	fmt.Printf("%d items, sum=%v\n", size, sum)
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	delimiterPos := 2

	s1 := s[: delimiterPos]
	fmt.Println(s1)
	go sum(s1, c)

	s2 := s[delimiterPos :]
	fmt.Println(s2)
	go sum(s2, c)

	/*
	x, y := <-c, <-c // receive from c
	*/
	waitMsg := "waiting for channel ..."
	fmt.Println(waitMsg)
	x := <-c
	fmt.Println(waitMsg)
	y := <-c

	fmt.Println("--")
	fmt.Printf("sum1=%d, sum2=%d\n", x, y)
}
