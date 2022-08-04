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
