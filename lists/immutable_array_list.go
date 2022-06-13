package lists

type ImmutableArrayList[T comparable] struct {
	values []T
}

var _ List[int] = &ImmutableArrayList[int]{}

func (l *ImmutableArrayList[T]) Contains(o T) bool {
	for _, val := range l.values {
		if val == o {
			return true
		}
	}
	return false
}

func (l *ImmutableArrayList[T]) Add(e T) bool {
	return false
}

func (l *ImmutableArrayList[T]) IsEmpty() bool {
	return len(l.values) == 0
}

func (l *ImmutableArrayList[T]) Remove(o T) bool {
	return false
}

func (l *ImmutableArrayList[T]) Size() int {
	return len(l.values)
}

func (l *ImmutableArrayList[T]) Get(index int) T {
	return l.values[index]
}
