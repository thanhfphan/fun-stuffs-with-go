package gid

import (
	"sync"
	"time"
)

// GlobalID ...
type GlobalID struct {
	epoch   time.Time
	shardID int64

	elapsedTime int64
	sequence    uint16

	mutex *sync.Mutex
}

const (
	TimeBits     int    = 41
	SequenceBits int    = 12
	ShardBits    int    = 11
	MaskSequence uint16 = 1<<SequenceBits - 1
)

// New ...
func New(shardID int64) *GlobalID {
	timeline, err := time.Parse("02/01/2006", "22/06/2022")
	if err != nil {
		panic(err)
	}

	return &GlobalID{
		epoch: timeline,
		mutex: new(sync.Mutex),
	}
}

// GenarateID return an id 64 bits
// +----------------+-------------------+--------------------+
// |    41 bits     |      12 bits      |		11 bits		 |
// |  (time bits)   |  (sequence bits)  |	 (shard bits)	 |
// +----------------+-------------------+--------------------+
func (g *GlobalID) GenarateID() uint64 {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	currentMiliseconds := time.Since(g.epoch.UTC()).Milliseconds()

	if g.elapsedTime < currentMiliseconds {
		g.elapsedTime = currentMiliseconds
		g.sequence = 0
	} else {
		g.sequence = (g.sequence + 1) & MaskSequence
		if g.sequence == 0 {
			g.elapsedTime++
			time.Sleep(time.Duration(g.elapsedTime-currentMiliseconds) * 1 * time.Millisecond)
		}
	}

	return uint64(g.elapsedTime)<<(SequenceBits+ShardBits) |
		uint64(g.sequence)<<uint64(ShardBits) |
		uint64(g.shardID)
}
