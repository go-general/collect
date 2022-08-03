package maps

import (
	"sync"
)

type syncMap[K comparable, V any] struct {
	lock sync.RWMutex
	m    map[K]V
}

func (m *syncMap[K, V]) IsEmpty() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return len(m.m) == 0
}

func (m *syncMap[K, V]) Size() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return len(m.m)
}

func (m *syncMap[K, V]) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.m = make(map[K]V)
}

func (m *syncMap[K, V]) ContainsKey(key K) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	_, ok := m.m[key]
	return ok
}

func (m *syncMap[K, V]) Get(key K) (V, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	val, ok := m.m[key]
	return val, ok
}

func (m *syncMap[K, V]) Put(key K, val V) (V, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	oldVal, ok := m.m[key]
	m.m[key] = val
	return oldVal, ok
}

func (m *syncMap[K, V]) Remove(key K) (V, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()

	oldVal, ok := m.m[key]
	if !ok {
		return oldVal, false
	}

	delete(m.m, key)
	return oldVal, true
}

func (m *syncMap[K, V]) Keys() []K {
	m.lock.RLock()
	defer m.lock.RUnlock()

	keys := make([]K, 0, len(m.m))
	for k := range m.m {
		keys = append(keys, k)
	}
	return keys
}

func (m *syncMap[K, V]) Values() []V {
	m.lock.RLock()
	defer m.lock.RUnlock()

	values := make([]V, 0, len(m.m))
	for _, v := range m.m {
		values = append(values, v)
	}
	return values
}

func (m *syncMap[K, V]) Range(f func(k K, v V) bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
