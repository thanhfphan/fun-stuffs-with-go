package src

type Deque[T any] interface {
	PushLeft(T) bool
	PushRight(T) bool

	PopLeft() (T, bool)
	PopRight() (T, bool)

	PeekLeft() (T, bool)
	PeekRight() (T, bool)

	Len() int
	IsEmpty() bool
}

func NewDeque[T any](initSize int) Deque[T] {
	return &sliceDeque[T]{
		data:  make([]T, initSize),
		right: 1,
	}
}

type sliceDeque[T any] struct {
	size, left, right int
	data              []T
}

func (d *sliceDeque[T]) PushLeft(element T) bool {
	d.data[d.left] = element
	d.size++
	d.left--
	if d.left < 0 {
		d.left = len(d.data) - 1 // Wrap around
	}

	d.resize()
	return true
}

func (d *sliceDeque[T]) PushRight(element T) bool {
	d.data[d.right] = element
	d.size++
	d.right++
	d.right %= len(d.data)

	d.resize()
	return true
}

func (d *sliceDeque[T]) PopLeft() (T, bool) {
	if d.size == 0 {
		return *new(T), false
	}

	idx := d.leftmostIdx()
	elt := d.data[idx]
	d.data[idx] = *new(T) // prevent memory leak
	d.size--
	d.left++
	d.left %= len(d.data)

	return elt, true
}

func (d *sliceDeque[T]) PopRight() (T, bool) {
	if d.size == 0 {
		return *new(T), false
	}

	idx := d.rightmostIdx()
	elt := d.data[idx]
	d.size--
	d.right--
	if d.right < 0 {
		d.right = len(d.data) - 1
	}

	return elt, true
}

func (d *sliceDeque[T]) PeekLeft() (T, bool) {
	if d.size == 0 {
		return *new(T), false
	}

	idx := d.leftmostIdx()
	return d.data[idx], true
}

func (d *sliceDeque[T]) PeekRight() (T, bool) {
	if d.size == 0 {
		return *new(T), false
	}

	idx := d.rightmostIdx()
	return d.data[idx], true
}

func (d *sliceDeque[T]) Len() int {
	return d.size
}

func (d *sliceDeque[T]) IsEmpty() bool {
	return d.size == 0
}

func (d *sliceDeque[T]) leftmostIdx() int {
	if d.left == len(d.data)-1 {
		return 0
	}

	return d.left + 1
}

func (d *sliceDeque[T]) rightmostIdx() int {
	if d.right == 0 {
		return len(d.data) - 1
	}

	return d.right - 1
}

func (d *sliceDeque[T]) resize() {
	if d.size != len(d.data) {
		return
	}

	newData := make([]T, d.size*2)
	leftMostIdx := d.leftmostIdx()
	copy(newData, d.data[leftMostIdx:])
	numCopied := len(d.data) - leftMostIdx
	copy(newData[numCopied:], d.data[:d.right])
	d.data = newData
	d.left = len(d.data) - 1
	d.right = d.size
}
