package sets

import (
	"sync"
)

type syncSet[T comparable] struct {
	lock sync.RWMutex
	m    map[T]struct{}
}

func (s *syncSet[T]) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.m) == 0
}

func (s *syncSet[T]) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.m)
}

func (s *syncSet[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.m = make(map[T]struct{})
}

func (s *syncSet[T]) Values() []T {
	s.lock.RLock()
	defer s.lock.RUnlock()

	values := make([]T, 0, len(s.m))
	for k := range s.m {
		values = append(values, k)
	}
	return values
}

func (s *syncSet[T]) Add(t ...T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, v := range t {
		s.m[v] = struct{}{}
	}
}

func (s *syncSet[T]) Contains(t T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	_, ok := s.m[t]
	return ok
}

func (s *syncSet[T]) Remove(t T) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.m[t]
	if !ok {
		return false
	}

	delete(s.m, t)
	return ok
}

func (s *syncSet[T]) Range(f func(t T) bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for k := range s.m {
		if !f(k) {
			return
		}
	}
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
	s.Range(func(k T) bool {
		if !s.Contains(k) {
			return false
		}
		return true
	})

	return true
}
