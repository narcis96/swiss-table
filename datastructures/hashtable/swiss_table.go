package hashtable

import (
	"time"
)

const (
	maxLoadFactor = 0.75
)

type SwissTable[K comparable, V any] struct {
	groups  []group[K, V]
	cnt     uint32
	seed    int64
	removed uint32
}

func exceedsLoadFactor(cnt, removed, numGroups uint32) bool {
	slots := float32(numGroups * groupSize)
	if float32(cnt-removed)/slots > maxLoadFactor {
		return true
	}

	if (cnt + removed) > numGroups*avgGroupLoadLimit {
		return true
	}
	return false
}

func nextSize(cnt int) int {
	// if m.dead >= (m.resident / 2) {
	// 	return uint32(len(m.groups))
	// }
	return cnt * 2
}

func NewSwissTable[K comparable, V any](capacity uint32) Hashtable[K, V] {
	m := &SwissTable[K, V]{}
	m.clearAndResize(int(capacity/groupSize + 1))
	return m
}

// Returns true if |key| is present in |m|.
func (m *SwissTable[K, V]) Has(key K) bool {
	groupIndex, matchPos := m.probing(key)
	return m.groups[groupIndex].ctrl[matchPos] != emptySlot
}

// Returns the |value| mapped by |key| if one exists.
func (m *SwissTable[K, V]) Get(key K) (V, bool) {
	groupIndex, matchPos := m.probing(key)
	if m.groups[groupIndex].ctrl[matchPos] != emptySlot {
		return m.groups[groupIndex].values[matchPos], true
	}
	var value V
	return value, false
}

func (m *SwissTable[K, V]) Put(key K, value V) {
	if exceedsLoadFactor(m.cnt, m.removed, uint32(len(m.groups))) {
		m.rehash(nextSize(len(m.groups)))
	}
	groupIndex, matchPos := m.probing(key)
	if m.groups[groupIndex].ctrl[matchPos] == emptySlot {
		m.cnt++
	}
	m.groups[groupIndex].keys[matchPos] = key
	m.groups[groupIndex].values[matchPos] = value
	m.groups[groupIndex].ctrl[matchPos] = h2(HashFNV64Safe(key, m.seed))
}

// Delete attempts to remove |key|, returns true successful.
func (m *SwissTable[K, V]) Delete(key K) bool {
	groupIndex, matchPos := m.probing(key)
	if m.groups[groupIndex].ctrl[matchPos] == emptySlot {
		return false
	}
	var k K
	var v V
	m.groups[groupIndex].keys[matchPos] = k
	m.groups[groupIndex].values[matchPos] = v
	m.groups[groupIndex].ctrl[matchPos] = deletedSlot
	m.removed++
	return true
}

func (m *SwissTable[K, V]) Clear() {
	m.clearAndResize(len(m.groups)/2 + 1)
}

func (m *SwissTable[K, V]) All(callback func(k K, v V) bool) {
	// take a consistent view of the table in case
	groups := m.groups
	for i := range groups {
		g := groups[i]
		for index := range groupSize {
			c := g.ctrl[index]
			if c == emptySlot || c == deletedSlot {
				continue
			}
			k, v := g.keys[index], g.values[index]
			if stop := callback(k, v); stop {
				return
			}
		}
	}

}

// Count returns the number of elements in the Map.
func (m *SwissTable[K, V]) Len() int {
	return int(m.cnt - m.removed)
}

// How many elements the map can accomodate without resizing.
func (m *SwissTable[K, V]) Capacity() int {
	return len(m.groups) * groupSize
}

func (m *SwissTable[K, V]) probing(key K) (int, int) {
	h1, h2 := splitHash(HashFNV64Safe(key, m.seed))
	numGroups := uint64(len(m.groups))
	groupIndex := h1 % numGroups

	for {
		gr := m.groups[groupIndex]
		matchPos, emptySlotFound := gr.match(key, h2)
		if matchPos >= 0 || emptySlotFound {
			return int(groupIndex), matchPos
		}
		groupIndex = (groupIndex + 1) % numGroups
	}
}

func (m *SwissTable[K, V]) rehash(n int) {
	groups := m.groups
	m.clearAndResize(n)
	for i := range groups {
		g := groups[i]
		for index := range groupSize {
			c := g.ctrl[index]
			if c == emptySlot || c == deletedSlot {
				continue
			}
			m.Put(g.keys[index], g.values[index])
		}
	}
}

func (m *SwissTable[K, V]) clearAndResize(numGroups int) {
	m.groups = make([]group[K, V], numGroups)
	for i := range m.groups {
		for index := range groupSize {
			m.groups[i].ctrl[index] = emptySlot
		}
	}
	m.removed, m.cnt, m.seed = 0, 0, time.Now().UnixMilli()+int64(numGroups)
}

// returns the 57 most significant bits of the hash
func h1(h uint64) uint64 { return h >> 7 }

// returns the 7 least significant bits of the hash
func h2(h uint64) int8 { return int8(h) & h2Mask }

func splitHash(h uint64) (uint64, int8) {
	return h1(h), h2(h)
}
