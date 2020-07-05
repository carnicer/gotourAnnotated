package main

import (
	"fmt"
	"time"
)

func main() {
	boomSet := false
	ticks := 0
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			ticks++
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			boomSet = true
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
		if boomSet { break }
	}
	fmt.Println(ticks)
}
