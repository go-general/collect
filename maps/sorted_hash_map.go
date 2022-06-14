package maps

import "github.com/go-general/collect/lists"

type SortedHashMap[K comparable, V any] struct {
	*HashMap[K, V]
	keys lists.List[K]
}

func (h *SortedHashMap[K, V]) Put(k K, v V) bool {
	if h.HashMap.Put(k, v) {
		h.keys.Add(k)
		return true
	}

	return false
}

func (h *SortedHashMap[K, V]) Merge(m ...Map[K, V]) bool {
	h.HashMap.Merge(m...)

	h.keys.Clear()
	h.HashMap.Range(func(k K, v V) bool {
		h.keys.Add(k)
		return true
	})

	return true
}
