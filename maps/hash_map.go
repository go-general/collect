package maps

type hashMap[K comparable, V any] struct {
	m map[K]V
}

func (m *hashMap[K, V]) IsEmpty() bool {
	return len(m.m) == 0
}

func (m *hashMap[K, V]) Size() int {
	return len(m.m)
}

func (m *hashMap[K, V]) Clear() {
	m.m = make(map[K]V)
}

func (m *hashMap[K, V]) ContainsKey(key K) bool {
	_, ok := m.m[key]
	return ok
}

func (m *hashMap[K, V]) Get(key K) (V, bool) {
	val, ok := m.m[key]
	return val, ok
}

func (m *hashMap[K, V]) Put(key K, val V) (V, bool) {
	oldVal, ok := m.m[key]
	m.m[key] = val
	return oldVal, ok
}

func (m *hashMap[K, V]) Remove(key K) (V, bool) {
	oldVal, ok := m.m[key]
	delete(m.m, key)
	return oldVal, ok
}

func (m *hashMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}

func (m *hashMap[K, V]) Values() []V {
	values := make([]V, 0, len(m.m))
	for _, val := range m.m {
		values = append(values, val)
	}
	return values
}

func (m *hashMap[K, V]) Range(f func(k K, v V) bool) {
	for k, v := range m.m {
		if !f(k, v) {
			break
		}
	}
}
