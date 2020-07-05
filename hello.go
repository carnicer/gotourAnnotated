package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, wg* sync.WaitGroup) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
	fmt.Println("--")

	defer wg.Done()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	var waitGroup sync.WaitGroup

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		waitGroup.Add(1)
		go c.Inc("somekey", &waitGroup)
	}

	defer fmt.Println("no wait (deferred):", c.Value("somekey"))
	waitGroup.Wait()
	fmt.Println("after wait:", c.Value("somekey"))
}
