package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	hour := t.Hour()
	fmt.Printf("time:%v, hour:%v\n", t, hour)
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

