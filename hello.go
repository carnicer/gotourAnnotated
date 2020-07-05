package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	fmt.Println("go launched, sleep ...")
	//time.Sleep(100 * time.Millisecond)
	fmt.Println("finished sleeping")
	say("hello")
	fmt.Println("finished")
}
