package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwissTableBasicOperations(t *testing.T) {
	table := NewSwissTable[string, int](5)

	table.Put("one", 1)
	val, exists := table.Get("one")
	assert.True(t, exists)
	assert.Equal(t, 1, val)

	// Test Has
	assert.True(t, table.Has("one"))
	assert.False(t, table.Has("two"))

	// Test Update Value
	table.Put("one", -1)
	val, exists = table.Get("one")
	assert.True(t, exists)
	assert.Equal(t, -1, val)
}

func TestSwissTableDeletion(t *testing.T) {
	table := NewSwissTable[string, int](5)
	table.Put("two", 2)
	table.Put("three", 3)

	// Test Delete
	deleted := table.Delete("two")
	assert.True(t, deleted)
	_, exists := table.Get("two")
	assert.False(t, exists)

	// Ensure other keys are unaffected
	val, exists := table.Get("three")
	assert.True(t, exists)
	assert.Equal(t, 3, val)
}

func TestSwissTableClear(t *testing.T) {
	table := NewSwissTable[string, int](5)
	table.Put("one", 1)
	table.Put("two", 2)

	table.Clear()

	assert.False(t, table.Has("one"))
	assert.False(t, table.Has("two"))
	assert.Equal(t, 0, table.Len())
}

func TestSwissTableResize(t *testing.T) {
	table := NewSwissTable[string, int](1)

	data := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10,
	}
	for k, v := range data {
		table.Put(k, v)
	}

	// Ensure all elements are still accessible
	for k, v := range data {
		val, exists := table.Get(k)
		assert.True(t, exists)
		assert.Equal(t, v, val)

	}

}

func TestSwissTableAll(t *testing.T) {
	table := NewSwissTable[string, int](1)
	table.Put("one", 1)
	table.Put("two", 2)
	table.Put("three", 3)

	values := make(map[string]int)
	table.All(func(k string, v int) bool {
		values[k] = v
		return false
	})

	assert.Equal(t, 3, len(values))
	assert.Equal(t, 1, values["one"])
	assert.Equal(t, 2, values["two"])
	assert.Equal(t, 3, values["three"])
}
