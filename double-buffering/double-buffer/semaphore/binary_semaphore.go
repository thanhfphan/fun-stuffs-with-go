package semaphore

import (
	"runtime"
	"sync/atomic"
)

type Semaphore interface {
	Signal()
	Wait()
}

type state int32

const (
	idle      state = 0
	modifying state = 1
)

type binarySemaphore struct {
	state int32
}

func NewBinarySemaphore() Semaphore {
	return &binarySemaphore{
		state: 0,
	}
}

func (g *binarySemaphore) Wait() {
	for !atomic.CompareAndSwapInt32(&g.state, int32(idle), int32(modifying)) {
		runtime.Gosched()
	}
}

func (g *binarySemaphore) Signal() {
	for !atomic.CompareAndSwapInt32(&g.state, int32(modifying), int32(idle)) {
		runtime.Gosched()
	}
}
