package gid

import (
	"testing"
)

// Benchmarks Presence Update event with fake data.
func BenchmarkGenerateChan(b *testing.B) {
	shard := New(1)

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = shard.GenarateID()
	}

}

func TestID(t *testing.T) {
	shard := New(1)

	var mapID = make(map[uint64]bool)
	for i := 0; i < 1000000; i++ {
		id := shard.GenarateID()
		if _, ok := mapID[id]; ok {
			t.Errorf("the id=%d should't repeat", id)
		}
		mapID[id] = true
	}
}
