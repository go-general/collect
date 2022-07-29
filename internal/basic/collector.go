package basic

type Collector interface {
	IsEmpty() bool

	Size() int

	Clear()
}
