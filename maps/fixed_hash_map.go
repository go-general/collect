package maps

type FixedHashMap[K comparable, V any] struct {
	*hashMap[K, V]
	maxSize int // max size of the map
}

func (h *FixedHashMap[K, V]) Put(k K, v V) bool {
	if h.maxSize == h.Size() {
		return false
	}

	return h.hashMap.Put(k, v)
}

func (h *FixedHashMap[K, V]) Merge(m ...Map[K, V]) bool {
	var totalSize int
	for _, m := range m {
		totalSize += m.Size()
	}

	if h.Size()+totalSize > h.maxSize {
		return false
	}

	return h.hashMap.Merge(m...)
}
