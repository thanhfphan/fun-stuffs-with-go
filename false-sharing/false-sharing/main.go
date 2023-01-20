package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum atomic.Int64
	counters := make([]atomic.Int64, 4)

	n := (1 << 25) // 2^25
	var wg sync.WaitGroup
	threads := 4
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			for j := 0; j < n; j++ {
				counters[idx].Add(1)
			}
			chunk := counters[idx].Load()
			sum.Add(chunk)
		}(i)

	}
	wg.Wait()
	fmt.Println("Sum: ", sum.Load())
}
