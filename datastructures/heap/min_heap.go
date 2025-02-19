package heap

type MinHeap []*Element

func (h MinHeap) Len() int {
	return len(h) - 1
}

// Defines the min-heap property (smallest priority first)
func (h MinHeap) less(i, j int) bool {
	return h[i].Priority < h[j].Priority
}

func (h MinHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(val *Element) {
	*h = append(*h, val)
	for i := h.Len(); i > 1 && h.less(i, i/2); {
		h.swap(i, i/2)
		i /= 2
	}
}
func (h *MinHeap) Top() *Element {
	if h.Len() == 0 {
		panic("invalid top operation. container is empty.")
	}
	return (*h)[1]
}

func (h *MinHeap) Pop() {
	if h.Len() == 0 {
		panic("invalid pop operation. container is empty.")
	}
	size := h.Len()
	h.swap(1, size)
	*h = (*h)[:size]
	size--
	for i := 1; i*2 <= size; {
		left, right := i*2, i*2+1
		next := left
		if right <= size && h.less(right, left) {
			next = right
		}
		if !h.less(next, i) {
			break
		}
		h.swap(i, next)
		i = next
	}
}

func NewMinHeap() Heap {
	return &MinHeap{{Data: "", Priority: 0}}
}
