package main

import (
	"math/rand"
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
	data := make([]int, 1<<20)
	for i := 0; i < interation; i++ {
		generateData(data)
		processData(data)
	}
}

func main() {
	run()
}
