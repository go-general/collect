package sets

import (
	"github.com/go-general/collect/internal/basic"
)

const (
	defaultSize = 16
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
}

func NewHashSet[T comparable]() Set[T] {
	return &hashSet[T]{
		m: make(map[T]struct{}, defaultSize),
	}
}

func NewHashSetWithSize[T comparable](size int) Set[T] {
	return &hashSet[T]{
		m: make(map[T]struct{}, size),
	}
}

func NewHashSetWithValues[T comparable](values ...T) Set[T] {
	m := make(map[T]struct{}, len(values))
	for _, val := range values {
		m[val] = struct{}{}
	}

	return &hashSet[T]{
		m: m,
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
		m: make(map[T]struct{}, defaultSize),
	}
}

func NewSyncSetWithSize[T comparable](size int) Set[T] {
	return &syncSet[T]{
		m: make(map[T]struct{}, size),
	}
}

func NewSyncSetWithValues[T comparable](values ...T) Set[T] {
	m := make(map[T]struct{}, len(values))
	for _, val := range values {
		m[val] = struct{}{}
	}

	return &syncSet[T]{
		m: m,
	}
}

// Union combines all objects in origin and target set.
func Union[T comparable](origin, target Set[T]) Set[T] {
	union := NewHashSet[T]()

	if origin != nil {
		origin.Range(func(obj T) bool {
			union.Add(obj)
			return true
		})
	}

	if target != nil {
		target.Range(func(obj T) bool {
			union.Add(obj)
			return true
		})
	}

	return union
}

// Intersect chooses same objects in both origin and target set.
func Intersect[T comparable](origin, target Set[T]) Set[T] {
	intersect := NewHashSet[T]()
	if origin == nil || target == nil {
		return intersect
	}

	origin.Range(func(obj T) bool {
		if target.Contains(obj) {
			intersect.Add(obj)
		}
		return true
	})

	return intersect
}

// Difference chooses objects in target set but not in origin set.
func Difference[T comparable](origin, target Set[T]) Set[T] {
	diff := NewHashSet[T]()
	if target == nil {
		return diff
	}

	target.Range(func(obj T) bool {
		if origin == nil {
			diff.Add(obj)
		} else if !origin.Contains(obj) {
			diff.Add(obj)
		}
		return true
	})

	return diff
}

// IsSubsetOf checks if all objects in origin set contains in target set.
func IsSubsetOf[T comparable](origin, target Set[T]) bool {
	if origin == nil {
		return true
	}

	if target == nil {
		return false
	}

	isSubSet := true
	origin.Range(func(obj T) bool {
		if target.Contains(obj) {
			return true
		}
		isSubSet = false
		return false
	})

	return isSubSet
}
