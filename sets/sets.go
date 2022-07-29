package sets

import (
	"github.com/go-general/collect/internal/basic"
)

type Set[T comparable] interface {
	basic.Collector

	Values() []T

	Add(...T) bool

	Merge(...Set[T]) bool

	Contains(T) bool

	Remove(T) bool

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
