package maps

import (
	"sync"

	"github.com/go-general/collect/lists"
)

func NewHashMap[K comparable, V any]() Map[K, V] {
	return &HashMap[K, V]{
		m:  make(map[K]V),
		mu: &sync.RWMutex{},
	}
}

func NewFixedHashMap[K comparable, V any](size int) Map[K, V] {
	return &FixedHashMap[K, V]{
		maxSize: size,
		HashMap: newHashMap[K, V](),
	}
}

func NewSortedHashMap[K comparable, V any]() Map[K, V] {
	return &SortedHashMap[K, V]{
		HashMap: newHashMap[K, V](),
		keys:    lists.NewArrayList[K](0),
	}
}
