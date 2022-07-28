package maps

import (
	"github.com/go-general/collect/internal/types"
	"github.com/go-general/collect/sets"
)

type Map[K comparable, V any] interface {
	// 从此映射中移除所有映射关系
	Clear()

	// 如果此映射包含指定键的映射关系，则返回 true
	ContainsKey(K) bool

	// 如果此映射将一个或多个键映射到指定值，则返回 true
	ContainsValue(V) bool

	// 返回指定键所映射的值；如果此映射不包含该键的映射关系，则返回 nil
	Get(K) *V

	// 如果此映射未包含键-值映射关系，则返回 true
	IsEmpty() bool

	// 将指定的值与此映射中的指定键关联
	Put(K, V) bool

	// 从指定映射中将所有映射关系复制到此映射中
	Merge(...Map[K, V]) bool

	// 如果存在一个键的映射关系，则将其从此映射中移除
	Remove(K) bool

	// 返回此映射中的键-值映射关系数
	Size() int

	// 返回此映射中包含的键的 Set 视图
	Keys() sets.Set[K]

	// 返回此映射中包含的值的 Collection 视图
	// TODO: use a list instead of a slice
	Values() []V

	Range(func(K, V) bool)
}

func NewHashMap[K comparable, V any]() Map[K, V] {
	return &hashMap[K, V]{
		m: make(map[K]V),
	}
}

func NewFixedHashMap[K comparable, V any](size int) Map[K, V] {
	return &fixedHashMap[K, V]{
		size: size,
		hashMap: &hashMap[K, V]{
			m: make(map[K]V),
		},
	}
}

func NewSortedHashMap[K types.Ordered, V any]() Map[K, V] {
	return &SortedHashMap[K, V]{
		hashMap: &hashMap[K, V]{
			m: make(map[K]V),
		},
		keys: sets.NewSortedSet[K](),
	}
}

func NewImmutableHashMap[K comparable, V any](m map[K]V) Map[K, V] {
	return &ImmutableHashMap[K, V]{
		hashMap: &hashMap[K, V]{
			m: m,
		},
	}
}
