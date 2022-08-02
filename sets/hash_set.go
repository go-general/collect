package sets

// hashSet is a set implementation that uses a hash map to store its elements.It implements interface Set.
type hashSet[T comparable] struct {
	m map[T]struct{}
}

func (s *hashSet[T]) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *hashSet[T]) Size() int {
	return len(s.m)
}

func (s *hashSet[T]) Clear() {
	s.m = make(map[T]struct{})
}

func (s *hashSet[T]) Values() []T {
	values := make([]T, 0, len(s.m))
	for k := range s.m {
		values = append(values, k)
	}
	return values
}

func (s *hashSet[T]) Add(t ...T) {
	for _, v := range t {
		s.m[v] = struct{}{}
	}
}

func (s *hashSet[T]) Contains(t T) bool {
	_, ok := s.m[t]
	return ok
}

func (s *hashSet[T]) Remove(t T) bool {
	_, ok := s.m[t]
	if !ok {
		return false
	}

	delete(s.m, t)
	return true
}

func (s *hashSet[T]) Range(f func(t T) bool) {
	for k := range s.m {
		if !f(k) {
			break
		}
	}
}

func (s *hashSet[T]) Merge(set ...Set[T]) Set[T] {
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

func (s *hashSet[T]) Union(set Set[T]) Set[T] {
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

func (s *hashSet[T]) Intersect(set Set[T]) Set[T] {
	intersect := NewHashSet[T]()

	set.Range(func(obj T) bool {
		if s.Contains(obj) {
			intersect.Add(obj)
		}
		return true
	})

	return intersect
}

func (s *hashSet[T]) Difference(set Set[T]) Set[T] {
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

func (s *hashSet[T]) IsSubsetOf(set Set[T]) bool {
	for k := range s.m {
		if !s.Contains(k) {
			return false
		}
	}

	return true
}
