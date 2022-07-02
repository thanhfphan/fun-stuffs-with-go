package gid

import "time"

// GlobalID ...
type GlobalID struct {
	epoch        time.Time
	timeBits     int
	sequenceBits int
	shardBits    int
	shardID      int64
}

// New ...
func New(shardID int64) *GlobalID {
	timeline, err := time.Parse("02/01/2006", "22/06/2022")
	if err != nil {
		panic(err)
	}

	return &GlobalID{
		timeBits:     41,
		sequenceBits: 12,
		shardBits:    11,
		epoch:        timeline,
	}
}

// Genarate return an id 64 bits
// ----------------------------------------------------------
// |	41 bits		|		12 bits		|		11 bits		|
// |  (time bits)	|  (sequence bits)	|	 (shard bits)	|
// ---------------------------------------------------------
func (g *GlobalID) Genarate() int64 {
	currentMiliseconds := time.Since(g.epoch.UTC()).Milliseconds()

	sequence := 1<<g.sequenceBits - 1 //TODO: hard code get last sequence, need to find a away get unique number each milisecond(use mutex lock or atomic package)

	return currentMiliseconds<<(g.sequenceBits+g.shardBits) | int64(sequence)<<int64(g.shardBits) | g.shardID
}
