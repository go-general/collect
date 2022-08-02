package sets

import (
	"sync"
)

type syncSet[T comparable] struct {
	m    *sync.Map
	size int
}

func (s *syncSet[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *syncSet[T]) Size() int {
	return s.size
}

func (s *syncSet[T]) Clear() {
	s.m = &sync.Map{}
	s.size = 0
}

func (s *syncSet[T]) Values() []T {
	values := make([]T, 0, s.size)
	s.m.Range(func(k, v any) bool {
		values = append(values, k)
		return true
	})
	return values
}

func (s *syncSet[T]) Add(t ...T) {
	s.m.Store(t, struct{}{})
	s.size++
}

func (s *syncSet[T]) Contains(t T) bool {
	_, ok := s.m.Load(t)
	return ok
}

func (s *syncSet[T]) Remove(t T) bool {
	_, ok := s.m.LoadAndDelete(t)
	if !ok {
		return false
	}

	s.size--
	return true
}

func (s *syncSet[T]) Range(f func(t T) bool) {
	s.m.Range(func(k, v any) bool {
		return f(k)
	})
}

func (s *syncSet[T]) Merge(set ...Set[T]) Set[T] {
	merged := NewHashSet[T]()

	s.Range(func(obj T) bool {
		merged.Add(obj)
		return true
	})

	for _, v := range set {
		v.Range(func(obj T) bool {
			merged.Add(obj)
			return true
		})
	}

	return merged
}

func (s *syncSet[T]) Union(set Set[T]) Set[T] {
	union := NewHashSet[T]()

	s.Range(func(obj T) bool {
		union.Add(obj)
		return true
	})

	set.Range(func(obj T) bool {
		union.Add(obj)
		return true
	})

	return union
}

func (s *syncSet[T]) Intersect(set Set[T]) Set[T] {
	intersect := NewHashSet[T]()

	s.Range(func(obj T) bool {
		if s.Contains(obj) {
			intersect.Add(obj)
		}
		return true
	})

	return intersect
}

func (s *syncSet[T]) Difference(set Set[T]) Set[T] {
	diff := NewHashSet[T]()

	s.Range(func(obj T) bool {
		if !s.Contains(obj) {
			diff.Add(obj)
		}
		return true
	})

	set.Range(func(obj T) bool {
		if !s.Contains(obj) {
			diff.Add(obj)
		}
		return true
	})

	return diff
}

func (s *syncSet[T]) IsSubsetOf(set Set[T]) bool {
	s.m.Range(func(k, v any) bool {
		if !s.Contains(k) {
			return false
		}
		return true
	})

	return true
}
