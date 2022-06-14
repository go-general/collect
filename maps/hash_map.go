package maps

import (
	"sync"

	"github.com/go-general/collect/lists"
)

// HashMap is a map with a hash function, which implements interface Map.
type HashMap[K comparable, V any] struct {
	m  map[K]V
	mu *sync.RWMutex
}

func newHashMap[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		m:  make(map[K]V),
		mu: &sync.RWMutex{},
	}
}

// 从此映射中移除所有映射关系
func (h *HashMap[K, V]) Clear() {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.m = make(map[K]V)
}

// 如果此映射包含指定键的映射关系，则返回 true
func (h *HashMap[K, V]) ContainsKey(key K) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	_, ok := h.m[key]
	return ok
}

// 如果此映射将一个或多个键映射到指定值，则返回 true
func (h *HashMap[K, V]) ContainsValue(val V) bool {
	return false
}

// 返回指定键所映射的值；如果此映射不包含该键的映射关系，则返回 nil
func (h *HashMap[K, V]) Get(key K) *V {
	h.mu.RLock()
	defer h.mu.RUnlock()

	val, ok := h.m[key]
	if !ok {
		return nil
	}

	return &val
}

// 如果此映射未包含键-值映射关系，则返回 true
func (h *HashMap[K, V]) IsEmpty() bool {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return len(h.m) == 0
}

// 将指定的值与此映射中的指定键关联
func (h *HashMap[K, V]) Put(key K, val V) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

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
func (h *HashMap[K, V]) Merge(m ...Map[K, V]) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, m := range m {
		m.Range(func(key K, val V) bool {
			h.m[key] = val
			return true
		})
	}

	return true
}

// 如果存在一个键的映射关系，则将其从此映射中移除
func (h *HashMap[K, V]) Remove(key K) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	_, ok := h.m[key]
	if !ok {
		return false
	}

	delete(h.m, key)
	return true
}

// 返回此映射中的键-值映射关系数
func (h *HashMap[K, V]) Size() int {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return len(h.m)
}

// 返回此映射中包含的键的 Set 视图
func (h *HashMap[K, V]) Keys() lists.List[K] {
	h.mu.RLock()
	defer h.mu.RUnlock()

	keys := lists.NewArrayList[K](len(h.m))
	for key := range h.m {
		keys.Add(key)
	}
	return keys
}

// 返回此映射中包含的值的 Collection 视图
func (h *HashMap[K, V]) Values() lists.List[V] {
	h.mu.RLock()
	defer h.mu.RUnlock()

	// TODO: need to implement a list which elements are any type of V
	// values := lists.NewArrayList[V](len(h.m))
	// for _, val := range h.m {
	// 	values.Add(val)
	// }
	// return values
	return nil
}

func (h *HashMap[K, V]) Range(f func(key K, val V) bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for key, val := range h.m {
		if !f(key, val) {
			break
		}
	}
}
