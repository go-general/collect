package lists

type ImmutableArrayList[T comparable] struct {
	values []T
}

func (l *ImmutableArrayList[T]) IsEmpty() bool {
	return len(l.values) == 0
}

func (l *ImmutableArrayList[T]) Size() int {
	return len(l.values)
}

func (l *ImmutableArrayList[T]) Clear() {}

func (l *ImmutableArrayList[T]) Values() []T {
	return l.values
}

func (l *ImmutableArrayList[T]) Add(obj T) bool {
	return false
}

func (l *ImmutableArrayList[T]) AddAll(list List[T]) bool {
	return false
}

func (l *ImmutableArrayList[T]) Contains(obj T) bool {
	for _, v := range l.values {
		if v == obj {
			return true
		}
	}
	return false
}

func (l *ImmutableArrayList[T]) Remove(obj T) bool {
	return false
}

func (l *ImmutableArrayList[T]) Get(index int) *T {
	if len(l.values) < index {
		return nil
	}
	return &l.values[index]
}

func (l *ImmutableArrayList[T]) Set(index int, obj T) *T {
	return nil
}

func (l *ImmutableArrayList[T]) Range(f func(index int, obj T) bool) {
	for idx, obj := range l.values {
		f(idx, obj)
	}
}
