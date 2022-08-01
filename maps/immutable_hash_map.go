package maps

type immutableHashMap[K comparable, V any] struct {
	m map[K]V
}

func (h *immutableHashMap[K, V]) IsEmpty() bool {
	return len(h.m) == 0
}

func (h *immutableHashMap[K, V]) Size() int {
	return len(h.m)
}

func (h *immutableHashMap[K, V]) Clear() {
}

func (h *immutableHashMap[K, V]) ContainsKey(key K) bool {
	_, ok := h.m[key]
	return ok
}

func (h *immutableHashMap[K, V]) Get(key K) (V, bool) {
	val, ok := h.m[key]
	return val, ok
}

func (h *immutableHashMap[K, V]) Put(key K, val V) (V, bool) {
	oldVal, ok := h.m[key]
	return oldVal, ok
}

func (h *immutableHashMap[K, V]) Remove(key K) (V, bool) {
	oldVal, ok := h.m[key]
	return oldVal, ok
}

func (h *immutableHashMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(h.m))
	for k := range h.m {
		keys = append(keys, k)
	}
	return keys
}

func (h *immutableHashMap[K, V]) Values() []V {
	values := make([]V, 0, len(h.m))
	for _, val := range h.m {
		values = append(values, val)
	}
	return values
}

func (h *immutableHashMap[K, V]) Range(f func(k K, v V) bool) {
	for k, v := range h.m {
		if !f(k, v) {
			break
		}
	}
}
