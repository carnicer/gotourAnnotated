package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Printf("%d items, sum=%v\n", len(s), sum)
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	s1 := s[: len(s)/2]
	fmt.Println(s1)
	go sum(s1, c)

	s2 := s[len(s)/2 :]
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
