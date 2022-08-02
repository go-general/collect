package maps

type immutableHashMap[K comparable, V any] struct {
	m map[K]V
}

func (m *immutableHashMap[K, V]) IsEmpty() bool {
	return len(m.m) == 0
}

func (m *immutableHashMap[K, V]) Size() int {
	return len(m.m)
}

func (m *immutableHashMap[K, V]) Clear() {
	// immutable hashMap can not do clear action
}

func (m *immutableHashMap[K, V]) ContainsKey(key K) bool {
	_, ok := m.m[key]
	return ok
}

func (m *immutableHashMap[K, V]) Get(key K) (V, bool) {
	val, ok := m.m[key]
	return val, ok
}

func (m *immutableHashMap[K, V]) Put(key K, val V) (V, bool) {
	oldVal, ok := m.m[key]
	return oldVal, ok
}

func (m *immutableHashMap[K, V]) Remove(key K) (V, bool) {
	oldVal, ok := m.m[key]
	return oldVal, ok
}

func (m *immutableHashMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}

func (m *immutableHashMap[K, V]) Values() []V {
	values := make([]V, 0, len(m.m))
	for _, val := range m.m {
		values = append(values, val)
	}
	return values
}

func (m *immutableHashMap[K, V]) Range(f func(k K, v V) bool) {
	for k, v := range m.m {
		if !f(k, v) {
			break
		}
	}
}
