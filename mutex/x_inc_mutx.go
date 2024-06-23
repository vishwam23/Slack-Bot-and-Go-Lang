package main

import (
	"fmt"
	"sync"
)

var (
	x  = 0
	mu sync.Mutex
	wg sync.WaitGroup
)

func incrementX() {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	x++
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementX()
	}

	wg.Wait()

	fmt.Println("Final value of x:", x)
}
