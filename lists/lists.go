package lists

import (
	"github.com/go-general/collect/internal/basic"
)

type List[T any] interface {
	basic.Collector

	Values() []T

	Add(obj T) bool

	AddAll(list List[T]) bool

	Contains(obj T) bool

	Remove(obj T) bool

	Get(index int) (T, bool)

	Set(index int, obj T) (T, bool)

	Range(func(index int, obj T) bool)
}

func NewArrayList[T comparable](size int) List[T] {
	return &arrayList[T]{
		values: make([]T, 0, size),
	}
}
