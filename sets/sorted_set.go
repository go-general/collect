package sets

import (
	"sort"

	"github.com/go-general/collect/internal/types"
)

type sortedSet[T types.Ordered] struct {
	*hashSet[T]
}

func newSortedSet[T types.Ordered]() *sortedSet[T] {
	return &sortedSet[T]{
		hashSet: newHashSet[T](),
	}
}

func (s *sortedSet[T]) Range(f func(T) bool) {
	values := s.Values()
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	for i := 0; i < len(values); i++ {
		if !f(values[i]) {
			break
		}
	}
}

func (s *sortedSet[T]) Values() []T {
	values := s.hashSet.Values()
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	return values
}
