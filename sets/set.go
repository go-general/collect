package sets

import (
	"github.com/go-general/collect"
)

type Set[T comparable] interface {
	collect.Collector[T]

	// add the specified element to this set if it is not already present
	Add(obj T) bool

	// addAll adds all of the elements in the specified collection to this set if they're not already present
	AddAll(set Set[T]) bool

	// Returns true if this set contains the specified element.
	Contains(obj T) bool

	// Returns true if this set contains all of the elements of the specified collection.
	// ContainsAll(Collection<?> c) bool

	// Compares the specified object with this set for equality.
	// Equals(Object o) bool

	// Removes the specified element from this set if it is present
	Remove(obj T) bool

	Range(func(obj T) bool)
}
