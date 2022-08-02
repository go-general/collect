package maps

import (
	"sync"
)

type syncMap[K any, V any] struct {
	m    *sync.Map
	size int
}

func (m *syncMap[K, V]) IsEmpty() bool {
	return m.size == 0
}

func (m *syncMap[K, V]) Size() int {
	return m.size
}

func (m *syncMap[K, V]) Clear() {
	m.m = &sync.Map{}
	m.size = 0
}

func (m *syncMap[K, V]) ContainsKey(key K) bool {
	_, ok := m.m.Load(key)
	return ok
}

func (m *syncMap[K, V]) Get(key K) (V, bool) {
	val, ok := m.m.Load(key)
	return val, ok
}

func (m *syncMap[K, V]) Put(key K, val V) (V, bool) {
	oldVal, ok := m.m.Load(key)
	m.m.Store(key, val)
	return oldVal, ok
}

func (m *syncMap[K, V]) Remove(key K) (V, bool) {
	oldVal, ok := m.m.LoadAndDelete(key)
	m.size--
	return oldVal, ok
}

func (m *syncMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.size)
	m.m.Range(func(k, v any) bool {
		keys = append(keys, k)
		return true
	})
	return keys
}

func (m *syncMap[K, V]) Values() []V {
	values := make([]V, 0, m.size)
	m.m.Range(func(k, v any) bool {
		values = append(values, v)
		return true
	})
	return values
}

func (m *syncMap[K, V]) Range(f func(k K, v V) bool) {
	m.m.Range(func(k, v any) bool {
		return f(k, v)
	})
}
