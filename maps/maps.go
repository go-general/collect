package maps

import (
	"github.com/go-general/collect/internal/basic"
)

type Map[K comparable, V any] interface {
	basic.Collector

	// ContainsKey checks if key exists in map.
	// return true if exists, false if not.
	ContainsKey(K) bool

	// Get gets value of key in map.
	// return value and true if exists, empty value and false if not exists.
	Get(K) (V, bool)

	// Put sets key and value to map.
	// return old value and true if K exists, empty value and false if K not exists.
	Put(K, V) (V, bool)

	// Remove removes the key in map.
	// return old value and true if K exists, empty value and false if K not exists.
	Remove(K) (V, bool)

	// Keys get all keys in map.
	Keys() []K

	// Values get all values in map.
	Values() []V

	// Range ranges key and value in map.
	Range(f func(k K, v V) bool)
}

func NewHashMap[K comparable, V any]() Map[K, V] {
	return &hashMap[K, V]{
		m: make(map[K]V),
	}
}
