package maps

type fixedHashMap[K comparable, V any] struct {
	*hashMap[K, V]
	size int // size of the map
}

func (h *fixedHashMap[K, V]) Put(k K, v V) bool {
	if h.size == h.Size() {
		return false
	}

	return h.hashMap.Put(k, v)
}

func (h *fixedHashMap[K, V]) Merge(m ...Map[K, V]) bool {
	var totalSize int
	for _, m := range m {
		totalSize += m.Size()
	}

	if h.Size()+totalSize > h.size {
		return false
	}

	return h.hashMap.Merge(m...)
}
