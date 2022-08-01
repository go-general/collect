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

	// Range ranges value in set.
	Range(func(obj T) bool)

	// Merge merges all values to this set.
	Merge(...Set[T]) Set[T]

	Union(Set[T]) Set[T]

	Intersect(Set[T]) Set[T]

	Difference(Set[T]) Set[T]

	IsSubsetOf(Set[T]) bool
}

func NewHashSet[T comparable]() Set[T] {
	return &hashSet[T]{
		m: make(map[T]struct{}),
	}
}

func NewImmutableHashSet[T comparable](values ...T) Set[T] {
	m := make(map[T]struct{}, len(values))
	for _, val := range values {
		m[val] = struct{}{}
	}

	return &hashSet[T]{
		m: m,
	}
}
