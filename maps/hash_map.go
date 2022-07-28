package maps

import (
	"github.com/go-general/collect/sets"
)

// hashMap is a map with a hash function, which implements interface Map.
type hashMap[K comparable, V any] struct {
	m map[K]V
}

// 从此映射中移除所有映射关系
func (h *hashMap[K, V]) Clear() {
	h.m = make(map[K]V)
}

// 如果此映射包含指定键的映射关系，则返回 true
func (h *hashMap[K, V]) ContainsKey(key K) bool {
	_, ok := h.m[key]
	return ok
}

// 如果此映射将一个或多个键映射到指定值，则返回 true
func (h *hashMap[K, V]) ContainsValue(val V) bool {
	return false
}

// 返回指定键所映射的值；如果此映射不包含该键的映射关系，则返回 nil
func (h *hashMap[K, V]) Get(key K) *V {
	val, ok := h.m[key]
	if !ok {
		return nil
	}

	return &val
}

// 如果此映射未包含键-值映射关系，则返回 true
func (h *hashMap[K, V]) IsEmpty() bool {
	return len(h.m) == 0
}

// 将指定的值与此映射中的指定键关联
func (h *hashMap[K, V]) Put(key K, val V) bool {
	// check if the key is already in the map
	_, ok := h.m[key]
	if ok {
		h.m[key] = val
		return false
	}

	h.m[key] = val
	return true
}

// 从指定映射中将所有映射关系复制到此映射中
func (h *hashMap[K, V]) Merge(m ...Map[K, V]) bool {
	for _, m := range m {
		m.Range(func(key K, val V) bool {
			h.m[key] = val
			return true
		})
	}

	return true
}

// 如果存在一个键的映射关系，则将其从此映射中移除
func (h *hashMap[K, V]) Remove(key K) bool {
	_, ok := h.m[key]
	if !ok {
		return false
	}

	delete(h.m, key)
	return true
}

// 返回此映射中的键-值映射关系数
func (h *hashMap[K, V]) Size() int {
	return len(h.m)
}

// 返回此映射中包含的键的 Set 视图
func (h *hashMap[K, V]) Keys() sets.Set[K] {
	keys := sets.NewHashSet[K]()
	for key := range h.m {
		keys.Add(key)
	}
	return keys
}

// 返回此映射中包含的值的 Collection 视图
func (h *hashMap[K, V]) Values() []V {
	// TODO: need to implement a list which elements are any type of V
	values := make([]V, 0, len(h.m))
	for _, val := range h.m {
		values = append(values, val)
	}

	return values
}

func (h *hashMap[K, V]) Range(f func(key K, val V) bool) {
	for key, val := range h.m {
		if !f(key, val) {
			break
		}
	}
}
