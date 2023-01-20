package lock

import (
	"runtime"
	"sync/atomic"
)

type state int32

const (
	idle   state = 0
	locked state = 1
)

type Spinlock struct {
	state int32
}

func (s *Spinlock) Lock() {
	for !atomic.CompareAndSwapInt32(&s.state, int32(idle), int32(locked)) {
		runtime.Gosched()
	}
}

func (s *Spinlock) Unlock() {
	atomic.StoreInt32(&s.state, int32(idle))
}
