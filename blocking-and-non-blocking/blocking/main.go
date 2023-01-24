package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	threads := 8
	iterations := 1 << 22
	count := 0
	var m sync.Mutex
	var wg sync.WaitGroup

	// fast work
	for i := 0; i < threads; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for {
				m.Lock()
				if count == iterations {
					m.Unlock()
					break
				}
				count++
				m.Unlock()
			}
		}()
	}

	// slow work
	go func() {
		wg.Add(1)
		defer wg.Done()

		for {
			m.Lock()
			if count == iterations {
				m.Unlock()
				break
			}
			time.Sleep(1 * time.Microsecond)
			count++
			m.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("Count value ", count)
}
