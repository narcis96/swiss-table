package hashtable

type Key string

type Value struct {
	Data string
}

type Hashtable[K comparable, V any] interface {
	Put(key K, x V)
	Has(key K) bool
	Get(key K) (V, bool)
	Delete(key K) bool
	Clear()
	Len() int
	Capacity() int
	All(callback func(k K, v V) bool)
}
