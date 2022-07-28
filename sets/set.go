package sets

import (
	"github.com/go-general/collect/internal/basic"
)

type Set[T comparable] interface {
	basic.Collector[T]

	// add the specified element to this set if it is not already present
	Add(...T) bool

	// Merge the specified sets into this set.
	Merge(...Set[T]) bool

	// Returns true if this set contains the specified element.
	Contains(T) bool

	// Returns true if this set contains all of the elements of the specified collection.
	// ContainsAll(Collection<?> c) bool

	// Compares the specified object with this set for equality.
	// Equals(Object o) bool

	// Removes the specified element from this set if it is present
	Remove(T) bool

	Range(func(obj T) bool)

	// operate on the set

	// Union finds the union of this set and the specified set and returns a new set with the result.
	Union(Set[T]) Set[T]

	// Intersect finds the intersection of this set and the specified set and returns a new set with the result.
	Intersect(Set[T]) Set[T]

	// Difference finds the difference of this set and the specified set and returns a new set with the result.
	Difference(Set[T]) Set[T]

	// IsSubsetOf determines whether this set is a subset of the specified set.
	IsSubsetOf(Set[T]) bool
}
