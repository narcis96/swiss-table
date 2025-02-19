package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeapPush(t *testing.T) {
	h := NewMinHeap()

	h.Push(&Element{Data: "task1", Priority: 3})
	h.Push(&Element{Data: "task3", Priority: 2})
	h.Push(&Element{Data: "task2", Priority: 1})

	assert.Equal(t, 1, h.Top().Priority)
}

func TestMinHeapPushDuplicatePriorities(t *testing.T) {
	h := NewMinHeap()

	for i := range 100 {
		h.Push(&Element{Data: "task1", Priority: 100 - i})
		h.Push(&Element{Data: "task2", Priority: 50 + i})
	}
	assert.Equal(t, 1, h.Top().Priority)
	assert.Equal(t, 200, h.Len())
}

func TestMinHeapPop(t *testing.T) {
	h := NewMinHeap()

	h.Push(&Element{Data: "task1", Priority: 3})
	h.Push(&Element{Data: "task2", Priority: 1})
	h.Push(&Element{Data: "task3", Priority: 2})

	h.Pop()

	assert.Equal(t, 2, h.Top().Priority)
	assert.Equal(t, 2, h.Len())
	h.Pop()

	assert.Equal(t, 3, h.Top().Priority)
	assert.Equal(t, 1, h.Len())
}

// Checks if the Top() method correctly retrieves the smallest element.
func TestMinHeapTop(t *testing.T) {
	h := NewMinHeap()
	h.Push(&Element{Data: "task1", Priority: 5})

	assert.Equal(t, 5, h.Top().Priority)
}

// Verifies that calling Pop or Top on an empty heap panics.
func TestMinHeapEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on empty heap but did not get one.")
		}
	}()
	h := NewMinHeap()
	h.Pop()
}
