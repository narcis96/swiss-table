package hashtable

import (
	"fmt"
	"hash"
	"hash/fnv"
	"sync"
)

// type hashfn func(key any, seed any) uint64

var fnv64Unsafe = fnv.New64()

func convertSeed(seed int64) []byte {
	var seedBytes [8]byte
	for i := 0; i < 8; i++ {
		seedBytes[i] = byte(seed >> (i * 8))
	}
	return seedBytes[:]
}

func HashFNV64Unsafe(key any, seed int64) uint64 {
	fnv64Unsafe.Reset()
	fnv64Unsafe.Write(convertSeed(seed))
	fnv64Unsafe.Write([]byte(fmt.Sprintf("%v", key)))
	return fnv64Unsafe.Sum64()
}

var fnvPool = sync.Pool{
	New: func() interface{} {
		return fnv.New64()
	},
}

func HashFNV64Safe(key any, seed int64) uint64 {
	h := fnvPool.Get().(hash.Hash64)
	h.Reset()
	h.Write(convertSeed(seed))
	h.Write([]byte(fmt.Sprintf("%v", key)))
	sum := h.Sum64()
	fnvPool.Put(h)
	return sum
}
