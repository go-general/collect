package maps

import (
	"github.com/go-general/collect/lists"
)

type Map[K comparable, V any] interface {
	// 从此映射中移除所有映射关系
	Clear()

	// 如果此映射包含指定键的映射关系，则返回 true
	ContainsKey(o K) bool

	// 如果此映射将一个或多个键映射到指定值，则返回 true
	ContainsValue(v V) bool

	// 返回指定键所映射的值；如果此映射不包含该键的映射关系，则返回 nil
	Get(k K) any

	// 如果此映射未包含键-值映射关系，则返回 true
	IsEmpty() bool

	// 将指定的值与此映射中的指定键关联
	Put(k K, v V) bool

	// 从指定映射中将所有映射关系复制到此映射中
	PutAll(m Map[K, V]) bool

	// 如果存在一个键的映射关系，则将其从此映射中移除
	Remove(k K) bool

	// 返回此映射中的键-值映射关系数
	Size() int

	// 返回此映射中包含的键的 Set 视图
	Keys() lists.List[K]

	// 返回此映射中包含的值的 Collection 视图
	Values() lists.List[V]

	Range(func(k K, v V) bool)
}
