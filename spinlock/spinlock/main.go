package main

import (
	"sync"

	"github.com/thanhfphan/fun-stuff-with-go/spinlock/spinlock/lock"
)

func main() {
	var l lock.Spinlock // Using spinlock

	maxInt := 1 << 25
	threads := 8
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for true {
				l.Lock()

				if maxInt <= 0 {
					l.Unlock()
					break
				}

				maxInt--
				l.Unlock()
			}
		}()
	}
	wg.Wait()
}
