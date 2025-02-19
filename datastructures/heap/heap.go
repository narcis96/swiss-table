package heap

type Element struct {
	Data     string
	Priority int
}

type Heap interface {
	Push(x *Element)
	Top() *Element
	Pop()
	Len() int
}
