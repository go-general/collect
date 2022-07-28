package sets

import (
	"github.com/go-general/collect/internal/types"
)

func NewHashSet[K comparable]() Set[K] {
	return &hashSet[K]{
		m: make(map[K]bool),
	}
}

func NewSortedSet[K types.Ordered]() Set[K] {
	return &sortedSet[K]{
		hashSet: newHashSet[K](),
	}
}
