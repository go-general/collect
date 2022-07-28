package sets

// hashSet is a set implementation that uses a hash map to store its elements.It implements interface Set.
type hashSet[T comparable] struct {
	m map[T]struct{}
}

func (h *hashSet[T]) IsEmpty() bool {
	return len(h.m) == 0
}

func (h *hashSet[T]) Size() int {
	return len(h.m)
}

func (h *hashSet[T]) Clear() {
	h.m = make(map[T]struct{})
}

func (h *hashSet[T]) Values() []T {
	values := make([]T, 0, len(h.m))
	for k := range h.m {
		values = append(values, k)
	}

	return values
}

func (h *hashSet[T]) Add(t ...T) bool {
	for _, v := range t {
		if _, ok := h.m[v]; !ok {
			h.m[v] = struct{}{}
		}
	}

	return true
}

func (h *hashSet[T]) Merge(s ...Set[T]) bool {
	for _, v := range s {
		v.Range(func(t T) bool {
			h.m[t] = struct{}{}
			return true
		})
	}

	return true
}

func (h *hashSet[T]) Contains(t T) bool {
	_, ok := h.m[t]
	return ok
}

func (h *hashSet[T]) Remove(t T) bool {
	_, ok := h.m[t]
	if !ok {
		return false
	}

	delete(h.m, t)
	return true
}

func (h *hashSet[T]) Range(f func(obj T) bool) {
	for k := range h.m {
		if !f(k) {
			break
		}
	}
}

func (h *hashSet[T]) Union(s Set[T]) Set[T] {
	union := NewHashSet[T]()
	s.Range(func(obj T) bool {
		union.Add(obj)
		return true
	})

	for k := range h.m {
		union.Add(k)
	}

	return union
}

func (h *hashSet[T]) Intersect(s Set[T]) Set[T] {
	intersect := NewHashSet[T]()
	h.Range(func(obj T) bool {
		if s.Contains(obj) {
			intersect.Add(obj)
		}
		return true
	})

	return intersect
}

func (h *hashSet[T]) Difference(s Set[T]) Set[T] {
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

func (h *hashSet[T]) IsSubsetOf(s Set[T]) bool {
	for k := range h.m {
		if !s.Contains(k) {
			return false
		}
	}

	return true
}
