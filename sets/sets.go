package sets

import (
	"sync"

	"github.com/go-general/collect/internal/basic"
)

type Set[T comparable] interface {
	basic.Collector

	// Values get all values in set.
	Values() []T

	// Add adds value to set.
	Add(t ...T)

	// Contains checks if value exists in set.
	// return true if exists, false if not exists.
	Contains(t T) bool

	// Remove removes value in this set.
	// return true if value exists, false if not exists.
	Remove(t T) bool

	// Range ranges value in set.
	Range(func(t T) bool)

	// Merge merges all values to this set.
	Merge(set ...Set[T]) Set[T]

	Union(set Set[T]) Set[T]

	Intersect(set Set[T]) Set[T]

	Difference(set Set[T]) Set[T]

	IsSubsetOf(set Set[T]) bool
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

func NewSyncSet[T comparable]() Set[T] {
	return &syncSet[T]{
		m:    &sync.Map{},
		size: 0,
	}
}
