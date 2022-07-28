package maps

import (
	"github.com/go-general/collect/internal/types"
	"github.com/go-general/collect/sets"
)

func NewHashMap[K comparable, V any]() Map[K, V] {
	return &hashMap[K, V]{
		m: make(map[K]V),
	}
}

func NewFixedHashMap[K comparable, V any](size int) Map[K, V] {
	return &FixedHashMap[K, V]{
		maxSize: size,
		hashMap: newHashMap[K, V](),
	}
}

func NewSortedHashMap[K types.Ordered, V any]() Map[K, V] {
	return &SortedHashMap[K, V]{
		hashMap: newHashMap[K, V](),
		keys:    sets.NewSortedSet[K](),
	}
}

func NewImmutableHashMap[K comparable, V any](m map[K]V) Map[K, V] {
	return &ImmutableHashMap[K, V]{
		hashMap: &hashMap[K, V]{
			m: m,
		},
	}
}
