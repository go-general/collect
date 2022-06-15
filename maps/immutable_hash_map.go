package maps

type ImmutableHashMap[K comparable, V any] struct {
	*hashMap[K, V]
}

func newImmutableHashMap[K comparable, V any](m map[K]V) *ImmutableHashMap[K, V] {
	return &ImmutableHashMap[K, V]{
		hashMap: &hashMap[K, V]{
			m: m,
		},
	}
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
