package lists

import (
	"github.com/go-general/collect/internal/basic"
)

type List[T any] interface {
	basic.Collector[T]

	// 添加元素到集合
	Add(obj T) bool

	// 存放⼀个集合
	AddAll(list List[T]) bool

	// 查找集合中的元素
	Contains(obj T) bool

	// 删除⼀个集合中的元素
	Remove(obj T) bool

	// 根据索引取得元素
	Get(index int) *T

	// 替换元素，index为要替换元素下标 element为要替换元素
	Set(index int, obj T) *T

	Range(func(index int, obj T) bool)
}
