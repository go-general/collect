package maps

type FixedHashMap[K comparable, V any] struct {
	*HashMap[K, V]
	maxSize int // max size of the map
}

func (h *FixedHashMap[K, V]) Put(k K, v V) bool {
	if h.maxSize == h.Size() {
		return false
	}

	return h.HashMap.Put(k, v)
}

func (h *FixedHashMap[K, V]) Merge(m ...Map[K, V]) bool {
	var totalSize int
	for _, m := range m {
		totalSize += m.Size()
	}

	if h.Size()+totalSize > h.maxSize {
		return false
	}

	return h.HashMap.Merge(m...)
}
