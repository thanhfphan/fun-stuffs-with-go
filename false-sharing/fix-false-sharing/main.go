package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type S struct {
	_       [56]int64 // the place where magic happens
	Counter atomic.Int64
}

func main() {
	var sum atomic.Int64

	n := (1 << 25) // 2^25
	var wg sync.WaitGroup
	threads := 4
	counters := make([]S, threads)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			for j := 0; j < n; j++ {
				counters[idx].Counter.Add(1)
			}
			chunk := counters[idx].Counter.Load()
			sum.Add(chunk)
		}(i)

	}
	wg.Wait()
	fmt.Println("Sum: ", sum.Load())
}
