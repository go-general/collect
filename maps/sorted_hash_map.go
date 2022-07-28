package maps

import (
	"github.com/go-general/collect/internal/types"
	"github.com/go-general/collect/sets"
)

type SortedHashMap[K types.Ordered, V any] struct {
	*hashMap[K, V]
	keys sets.Set[K]
}

func (h *SortedHashMap[K, V]) Put(k K, v V) bool {
	if h.hashMap.Put(k, v) {
		h.keys.Add(k)
		return true
	}

	return false
}

func (h *SortedHashMap[K, V]) Merge(m ...Map[K, V]) bool {
	h.hashMap.Merge(m...)

	h.keys.Clear()
	h.hashMap.Range(func(k K, v V) bool {
		h.keys.Add(k)
		return true
	})

	return true
}

func (h *SortedHashMap[K, V]) Remove(k K) bool {
	if h.hashMap.Remove(k) {
		h.keys.Remove(k)
		return true
	}

	return false
}

func (h *SortedHashMap[K, V]) Keys() sets.Set[K] {
	return h.keys
}

func (h *SortedHashMap[K, V]) Range(f func(K, V) bool) {
	h.keys.Range(func(k K) bool {
		v := h.hashMap.Get(k)
		return f(k, *v)
	})
}

func (h *SortedHashMap[K, V]) Values() []V {
	var values = make([]V, 0, h.keys.Size())
	h.keys.Range(func(k K) bool {
		v := h.hashMap.Get(k)
		values = append(values, *v)
		return true
	})

	return values
}
