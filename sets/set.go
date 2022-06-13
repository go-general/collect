package sets

import (
	"github.com/go-general/collect/lists"
)

type Set[T any] interface {
	// add the specified element to this set if it is not already present
	Add(e T) bool

	// addAll adds all of the elements in the specified collection to this set if they're not already present
	// AddAll(Collection<? extends E> c) bool

	// Removes all of the elements from this set (optional operation).
	Clear()

	// Returns true if this set contains the specified element.
	Contains(o T) bool

	// Returns true if this set contains all of the elements of the specified collection.
	// ContainsAll(Collection<?> c) bool

	// Compares the specified object with this set for equality.
	// Equals(Object o) bool

	// Returns true if this set contains no elements.
	IsEmpty() bool

	// Removes the specified element from this set if it is present
	Remove(o T) bool

	// Removes from this set all of its elements that are contained in the specified collection
	RemoveAll(c lists.List[T]) bool

	// Retains only the elements in this set that are contained in the specified collection
	// RetainAll(Collection<?> c) bool

	// Returns the number of elements in this set (its cardinality).
	Size() int

	// Returns an array containing all of the elements in this set.
	ToArray() []T

	Range(func(obj T) bool)
}
