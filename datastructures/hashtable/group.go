package hashtable

const (
	groupSize         = 8
	avgGroupLoadLimit = 6
)

const (
	h2Mask      int8 = 0x7f // 0b0111_1111
	emptySlot   int8 = -128 // 0b1000_0000
	deletedSlot int8 = -2   // 0b1111_1110
)

// group is a group of 16 key-value pairs
type group[K comparable, V any] struct {
	// metadata is the h2 metadata array for a group.
	// find operations first probe the controls bytes
	// to filter candidates before matching keys
	ctrl   [groupSize]int8
	keys   [groupSize]K
	values [groupSize]V
}

// (int) the position of the match or -1 if not found
// (bool) true if we reached an empty slot and false otherwise
func (g group[K, V]) match(key K, h2 int8) (int, bool) {
	for i := 0; i < groupSize; i++ {
		if g.ctrl[i] == h2 && g.keys[i] == key {
			return i, false
		}
		if g.ctrl[i] == emptySlot {
			return i, true
		}
	}
	return -1, false
}
