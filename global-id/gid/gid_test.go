package gid

import (
	"testing"
)

// Benchmarks Presence Update event with fake data.
func BenchmarkGenerateChan(b *testing.B) {
	shard := New(1)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = shard.Genarate()
	}

}

func TestID(t *testing.T) {
	shard := New(1)

	var mapID = make(map[int64]bool)
	for i := 0; i < 100000; i++ {
		id := shard.Genarate()
		if _, ok := mapID[id]; ok {
			t.Errorf("the id=%d should't repeat", id)
		}
		mapID[id] = true
	}
}
