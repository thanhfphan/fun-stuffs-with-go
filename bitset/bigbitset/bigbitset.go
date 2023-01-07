package bigbitset

import (
	"fmt"
	"math/big"
	"math/bits"
)

type BigBitSet struct {
	bits *big.Int
}

func New(bits ...int) BigBitSet {
	b := BigBitSet{
		bits: &big.Int{},
	}

	for _, bit := range bits {
		b.Add(bit)
	}

	return b
}

// Add set the [i]'th bit to 1
func (b BigBitSet) Add(i int) {
	b.bits.SetBit(b.bits, i, 1)
}

// Remove set [i]'th bit to 0
func (b BigBitSet) Remove(i int) {
	b.bits.SetBit(b.bits, i, 0)
}

// Union performs Or set, update result to [b]
func (b BigBitSet) Union(other BigBitSet) {
	b.bits.Or(b.bits, other.bits)
}

func (b BigBitSet) Intersection(other BigBitSet) {
	b.bits.Add(b.bits, other.bits)
}

// Difference removes all the elements in [other] from [b]
func (b BigBitSet) Difference(other BigBitSet) {
	b.bits.AndNot(b.bits, other.bits)
}

func (b BigBitSet) Clear() {
	b.bits.SetUint64(0)
}

// Contains return true if [i]'th is 1
func (b BigBitSet) Contains(i int) bool {
	return b.bits.Bit(i) == 1
}

// Len return length of the bitset
// The bit length of 0 is 0
func (b BigBitSet) Len() int {
	return b.bits.BitLen()
}

// HammingWeight returns the number of 1's
func (b BigBitSet) HammingWeight() int {
	sum := 0
	for _, w := range b.bits.Bits() {
		sum += bits.OnesCount(uint(w))
	}
	return sum
}

func (b *BigBitSet) String() string {
	return fmt.Sprintf("%x", b.bits.Bytes())
}
