package lists

type immutableArrayList[T comparable] struct {
	values []T
}

func (l *immutableArrayList[T]) IsEmpty() bool {
	return len(l.values) == 0
}

func (l *immutableArrayList[T]) Size() int {
	return len(l.values)
}

func (l *immutableArrayList[T]) Clear() {}

func (l *immutableArrayList[T]) Values() []T {
	return l.values
}

func (l *immutableArrayList[T]) Add(obj T) bool {
	return false
}

func (l *immutableArrayList[T]) AddAll(list List[T]) bool {
	return false
}

func (l *immutableArrayList[T]) Contains(obj T) bool {
	for _, v := range l.values {
		if v == obj {
			return true
		}
	}
	return false
}

func (l *immutableArrayList[T]) Remove(obj T) bool {
	return false
}

func (l *immutableArrayList[T]) Get(index int) *T {
	if len(l.values) < index {
		return nil
	}
	return &l.values[index]
}

func (l *immutableArrayList[T]) Set(index int, obj T) *T {
	return nil
}

func (l *immutableArrayList[T]) Range(f func(index int, obj T) bool) {
	for idx, obj := range l.values {
		f(idx, obj)
	}
}
