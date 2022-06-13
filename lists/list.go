package lists

type List[T any] interface {
	// 添加元素到集合
	Add(e T) bool

	// 存放⼀个集合
	// AddAll(l List[T]) bool

	// 查找集合中的元素
	Contains(o T) bool

	// 判断⼀个集合是否为空
	IsEmpty() bool

	// 删除⼀个集合中的元素
	Remove(o T) bool

	// 返回集合的长
	Size() int

	// 根据索引取得元素
	Get(index int) T

	// 替换元素，index为要替换元素下标 element为要替换元素
	// set(index int, e T) T

	Values() []T

	Range(func(index int, obj T) bool)
}
