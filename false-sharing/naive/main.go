package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum atomic.Int64

	n := (1 << 25) // 2^25
	var wg sync.WaitGroup
	threads := 4
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < n; j++ {
				sum.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Sum: ", sum.Load())
}
