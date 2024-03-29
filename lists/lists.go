package lists

import (
	"github.com/go-general/collect/internal/basic"
)

const (
	defaultSize = 16
)

type List[T any] interface {
	basic.Collector

	// Values return all value in list.
	Values() []T

	// Add adds value to list.
	Add(t ...T)

	// AddAll adds all list values to list.
	AddAll(list List[T])

	// Contains checks if value exists in list.
	// return true if exists, false if not exists.
	Contains(t T) bool

	// Remove removes value in list.
	// return true if exists, false if not exists.
	Remove(t T) bool

	// Get gets value of given index in list.
	// return value and true if index exists, empty value and false if index not exists.
	Get(index int) (T, bool)

	// Set sets value of given index in list.
	// return old value and true if index exists, empty value and false if index not exists.
	Set(index int, t T) (T, bool)

	// Range ranges index and value in list.
	Range(func(index int, t T) bool)
}

func NewArrayList[T comparable]() List[T] {
	return &arrayList[T]{
		values: make([]T, 0, defaultSize),
	}
}

func NewArrayListWithSize[T comparable](size int) List[T] {
	return &arrayList[T]{
		values: make([]T, 0, size),
	}
}

func NewArrayListWithValues[T comparable](values ...T) List[T] {
	vals := make([]T, len(values))
	copy(vals, values)

	return &arrayList[T]{
		values: vals,
	}
}

func NewImmutableArrayList[T comparable](values ...T) List[T] {
	vals := make([]T, len(values))
	copy(vals, values)

	return &immutableArrayList[T]{
		values: vals,
	}
}

func NewSyncList[T comparable]() List[T] {
	return &syncList[T]{
		values: make([]T, 0, defaultSize),
	}
}

func NewSyncListWithSize[T comparable](size int) List[T] {
	return &syncList[T]{
		values: make([]T, 0, size),
	}
}

func NewSyncListWithValues[T comparable](values ...T) List[T] {
	vals := make([]T, len(values))
	copy(vals, values)

	return &syncList[T]{
		values: vals,
	}
}
