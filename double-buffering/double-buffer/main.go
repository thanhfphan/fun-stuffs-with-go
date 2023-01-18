package main

import (
	"math/rand"
	"sync"

	"github.com/thanhfphan/fun-stuff-with-go/double-buffering/double-buffer/semaphore"
)

func generateData(data []int) {
	for i := range data {
		data[i] = rand.Intn(100)
	}
}

func processData(data []int) {
	// dummy work
	for i := 0; i < 5; i++ {
		for i := range data {
			data[i] %= data[i] + 1
		}
	}
}

// the main thing
func run() {
	interation := 100
	data1 := make([]int, 1<<20)
	data2 := make([]int, 1<<20)
	bs := semaphore.NewBinarySemaphore()
	var w sync.WaitGroup
	w.Add(2)
	// generate data
	go func() {
		defer w.Done()
		for i := 0; i < interation; i++ {
			generateData(data1)

			// wait until processing work is done
			bs.Wait()

			// swap data
			copy(data2, data1)

			// signal to [processing data] goroutine
			bs.Signal()
		}
	}()

	// processing data
	go func() {
		defer w.Done()
		for i := 0; i < interation; i++ {
			// wait signal from [generate data] goroutine
			bs.Wait()

			processData(data2)

			// signal to [processing data] goroutine to begin the work
			bs.Signal()
		}
	}()

	w.Wait()
}

func main() {
	run()
}
