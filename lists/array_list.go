package lists

type ArrayList[T comparable] struct {
	values []T
}

func (l *ArrayList[T]) IsEmpty() bool {
	return len(l.values) == 0
}

func (l *ArrayList[T]) Size() int {
	return len(l.values)
}

func (l *ArrayList[T]) Clear() {
	l.values = make([]T, 0, 0)
}

func (l *ArrayList[T]) Values() []T {
	return l.values
}

func (l *ArrayList[T]) Add(obj T) bool {
	l.values = append(l.values, obj)
	return true
}

func (l *ArrayList[T]) AddAll(list List[T]) bool {
	l.values = append(l.values, list.Values()...)
	return true
}

func (l *ArrayList[T]) Contains(obj T) bool {
	for _, v := range l.values {
		if v == obj {
			return true
		}
	}
	return false
}

func (l *ArrayList[T]) Remove(obj T) bool {
	for idx, v := range l.values {
		if v == obj {
			l.values = append(l.values[:idx], l.values[idx+1:]...)
			return true
		}
	}
	return false
}

func (l *ArrayList[T]) Get(index int) *T {
	if len(l.values) < index {
		return nil
	}
	return &l.values[index]
}

func (l *ArrayList[T]) Set(index int, obj T) *T {
	if len(l.values) < index {
		return nil
	}
	old := l.values[index]
	l.values[index] = obj
	return &old
}

func (l *ArrayList[T]) Range(f func(index int, obj T) bool) {
	for idx, obj := range l.values {
		f(idx, obj)
	}
}
