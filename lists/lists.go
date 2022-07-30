package lists

import (
	"github.com/go-general/collect/internal/basic"
)

type List[T any] interface {
	basic.Collector

	// Values return all value in list.
	Values() []T

	// Add adds value to list.
	Add(obj ...T)

	// AddAll adds all list values to list.
	AddAll(list List[T])

	// Contains checks if value exists in list.
	// return true if exists, false if not exists.
	Contains(obj T) bool

	// Remove removes value in list.
	// return true if exists, false if not exists.
	Remove(obj T) bool

	// Get gets value of given index in list.
	// return value and true if index exists, empty value and false if index not exists.
	Get(index int) (T, bool)

	// Set sets value of given index in list.
	// return old value and true if index exists, empty value and false if index not exists.
	Set(index int, obj T) (T, bool)

	// Range ranges index and value in list.
	Range(func(index int, obj T) bool)
}

func NewArrayList[T comparable](size int) List[T] {
	return &arrayList[T]{
		values: make([]T, 0, size),
	}
}
