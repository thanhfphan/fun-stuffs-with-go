package lock

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func testCorrectness(t *testing.T, threads, n int) (bool, bool) {
	t.Helper()
	var wg sync.WaitGroup
	var count1, count2 int
	var l Spinlock
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < n; i++ {
				l.Lock()
				count1 += 1
				count2 += 2
				l.Unlock()
			}
		}()
	}

	wg.Wait()

	c1 := threads * n
	c2 := threads * n * 2
	return count1 == c1, count2 == c2
}

func TestCorrectness(t *testing.T) {
	r := require.New(t)
	c1, c2 := testCorrectness(t, 1, 100)
	r.True(c1)
	r.True(c2)
	c3, c4 := testCorrectness(t, 4, 1000)
	r.True(c3)
	r.True(c4)
	c5, c6 := testCorrectness(t, 8, 10000)
	r.True(c5)
	r.True(c6)
}
