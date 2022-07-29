package maps

import (
	"github.com/go-general/collect/internal/basic"
)

type Map[K comparable, V any] interface {
	basic.Collector

	ContainsKey(K) bool

	Get(K) (V, bool)

	Put(K, V)

	Remove(K)

	Keys() []K

	Values() []V
}

func NewHashMap[K comparable, V any]() Map[K, V] {
	return &hashMap[K, V]{
		m: make(map[K]V),
	}
}
