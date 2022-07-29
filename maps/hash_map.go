package maps

type hashMap[K comparable, V any] struct {
	m map[K]V
}

func (h *hashMap[K, V]) IsEmpty() bool {
	return len(h.m) == 0
}

func (h *hashMap[K, V]) Size() int {
	return len(h.m)
}

func (h *hashMap[K, V]) Clear() {
	h.m = make(map[K]V)
}

func (h *hashMap[K, V]) ContainsKey(key K) bool {
	_, ok := h.m[key]
	return ok
}

func (h *hashMap[K, V]) Get(key K) (V, bool) {
	val, ok := h.m[key]
	return val, ok
}

func (h *hashMap[K, V]) Put(key K, val V) {
	h.m[key] = val
}

func (h *hashMap[K, V]) Remove(key K) {
	delete(h.m, key)
}

func (h *hashMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(h.m))
	for k := range h.m {
		keys = append(keys, k)
	}
	return keys
}

func (h *hashMap[K, V]) Values() []V {
	values := make([]V, 0, len(h.m))
	for _, val := range h.m {
		values = append(values, val)
	}
	return values
}
