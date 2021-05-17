package dumphash

type Bucket struct {
	Value     int
	IsDeleted bool
}

type MyHashMap struct {
	Values []*Bucket
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	vals := make([]*Bucket, 1000)
	return MyHashMap{
		Values: vals,
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	for key > len(this.Values) {
		tmp := make([]*Bucket, len(this.Values)*2)
		for i, val := range this.Values {
			tmp[i] = val
		}
		this.Values = tmp
	}
	if this.Values[key] != nil {
		this.Values[key].Value = value
		this.Values[key].IsDeleted = false
	} else {
		this.Values[key] = &Bucket{
			Value:     value,
			IsDeleted: false,
		}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if key > len(this.Values) {
		return -1
	}

	if this.Values[key] != nil && this.Values[key].IsDeleted == false {
		return this.Values[key].Value
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	if key > len(this.Values) {
		return
	}
	if this.Values[key] != nil {
		this.Values[key].IsDeleted = true
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
