package sets

import (
	"github.com/go-general/collect/internal/basic"
)

type Set[T comparable] interface {
	basic.Collector

	// Values get all values in set.
	Values() []T

	// Add adds value to set.
	Add(...T)

	// Contains checks if value exists in set.
	// return true if exists, false if not exists.
	Contains(T) bool

	// Remove removes value in this set.
	// return true if value exists, false if not exists.
	Remove(T) bool

	// Merge merges all values to this set.
	Merge(...Set[T])

	// Range ranges value in set.
	Range(func(obj T) bool)

	Union(Set[T]) Set[T]

	Intersect(Set[T]) Set[T]

	Difference(Set[T]) Set[T]

	IsSubsetOf(Set[T]) bool
}

func NewHashSet[K comparable]() Set[K] {
	return &hashSet[K]{
		m: make(map[K]struct{}),
	}
}
