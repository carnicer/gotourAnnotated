package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("-- i=%d, sleep ...\n", i)
		time.Sleep(time.Second)
		fmt.Printf("-- i=%d, channel ...\n", i)
		c <- x
		x, y = y, x+y
		fmt.Printf("-- i=%d, ====\n", i)
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	fmt.Printf("goroutine launched, c size=%d\n", len(c))
	for i := range c {
		fmt.Println(i)
	}
}
