package maps

type ImmutableHashMap[K comparable, V any] struct {
	*hashMap[K, V]
}

func (i *ImmutableHashMap[K, V]) Put(_ K, _ V) bool {
	return false
}

func (i *ImmutableHashMap[K, V]) Merge(_ ...Map[K, V]) bool {
	return false
}

func (i *ImmutableHashMap[K, V]) Clear() {}

func (i *ImmutableHashMap[K, V]) Remove(_ K) bool {
	return false
}
