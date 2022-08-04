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
