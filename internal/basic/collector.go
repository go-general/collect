package basic

type Collector[T any] interface {
	IsEmpty() bool

	Size() int

	Clear()

	Values() []T
}
