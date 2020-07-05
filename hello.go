package main

import "fmt"
import "time"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Printf("fibonacci: number sent: %d\n", x)
			x, y = y, x+y
			fmt.Printf("fibonacci: now x=%d, y=%d\n", x, y)
		case <-quit:
			fmt.Println("fibonacci: quit")
			return
		}
	}
}

func feeder(c, quit chan int) {
		var fibonum int
		for i := 0; i < 10; i++ {
			fmt.Println("feeder: sleep ...")
			time.Sleep(2 * time.Second)
			fmt.Println("feeder: wait ...")
			fibonum = <-c
			fmt.Printf("feeder: fibonacci => %d\n", fibonum)
		}
		fmt.Println("feeder: sleep before QUIT ...\n")
		time.Sleep(time.Second)
		fmt.Println("feeder: send QUIT via its channel\n")
		quit <- 0
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go feeder(c, quit)
	fibonacci(c, quit)
}
