package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	threads := 8
	iterations := 1 << 22
	var count atomic.Int32
	var wg sync.WaitGroup

	// fast work
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				expected := count.Load()
				if expected == int32(iterations) {
					return
				}
				desired := expected + 1
				count.CompareAndSwap(expected, desired)
			}
		}()
	}

	// slow work
	go func() {
		wg.Add(1)
		defer wg.Done()

		for {
			expected := count.Load()
			if expected == int32(iterations) {
				return
			}
			desired := expected + 1
			count.CompareAndSwap(expected, int32(desired))
			time.Sleep(1 * time.Microsecond)
		}
	}()

	wg.Wait()

	fmt.Println("Count value ", count.Load())
}
