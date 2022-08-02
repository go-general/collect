package sets

type immutableHashSet[T comparable] struct {
	m map[T]struct{}
}

func (h *immutableHashSet[T]) IsEmpty() bool {
	return len(h.m) == 0
}

func (h *immutableHashSet[T]) Size() int {
	return len(h.m)
}

func (h *immutableHashSet[T]) Clear() {
	// immutable hashSet can not do clear action
}

func (h *immutableHashSet[T]) Values() []T {
	values := make([]T, 0, len(h.m))
	for k := range h.m {
		values = append(values, k)
	}
	return values
}

func (h *immutableHashSet[T]) Add(t ...T) {
	// immutable hashSet can not do Add action
}

func (h *immutableHashSet[T]) Contains(t T) bool {
	_, ok := h.m[t]
	return ok
}

func (h *immutableHashSet[T]) Remove(t T) bool {
	_, ok := h.m[t]
	if !ok {
		return false
	}

	return true
}

func (h *immutableHashSet[T]) Range(f func(t T) bool) {
	for k := range h.m {
		if !f(k) {
			break
		}
	}
}

func (h *immutableHashSet[T]) Merge(s ...Set[T]) Set[T] {
	merged := NewHashSet[T]()

	h.Range(func(obj T) bool {
		merged.Add(obj)
		return true
	})

	for _, v := range s {
		v.Range(func(obj T) bool {
			merged.Add(obj)
			return true
		})
	}

	return merged
}

func (h *immutableHashSet[T]) Union(s Set[T]) Set[T] {
	union := NewHashSet[T]()

	s.Range(func(obj T) bool {
		union.Add(obj)
		return true
	})

	h.Range(func(obj T) bool {
		union.Add(obj)
		return true
	})

	return union
}

func (h *immutableHashSet[T]) Intersect(s Set[T]) Set[T] {
	intersect := NewHashSet[T]()

	h.Range(func(obj T) bool {
		if s.Contains(obj) {
			intersect.Add(obj)
		}
		return true
	})

	return intersect
}

func (h *immutableHashSet[T]) Difference(s Set[T]) Set[T] {
	diff := NewHashSet[T]()

	h.Range(func(obj T) bool {
		if !s.Contains(obj) {
			diff.Add(obj)
		}
		return true
	})

	s.Range(func(obj T) bool {
		if !h.Contains(obj) {
			diff.Add(obj)
		}
		return true
	})

	return diff
}

func (h *immutableHashSet[T]) IsSubsetOf(s Set[T]) bool {
	for k := range h.m {
		if !s.Contains(k) {
			return false
		}
	}

	return true
}
