package lists

type arrayList[T comparable] struct {
	values []T
}

func (l *arrayList[T]) IsEmpty() bool {
	return len(l.values) == 0
}

func (l *arrayList[T]) Size() int {
	return len(l.values)
}

func (l *arrayList[T]) Clear() {
	l.values = make([]T, 0, 0)
}

func (l *arrayList[T]) Values() []T {
	return l.values
}

func (l *arrayList[T]) Add(obj ...T) {
	for _, val := range obj {
		l.values = append(l.values, val)
	}
}

func (l *arrayList[T]) AddAll(list List[T]) {
	l.values = append(l.values, list.Values()...)
}

func (l *arrayList[T]) Contains(obj T) bool {
	for idx := range l.values {
		if l.values[idx] == obj {
			return true
		}
	}
	return false
}

func (l *arrayList[T]) Remove(obj T) bool {
	for idx := range l.values {
		if l.values[idx] == obj {
			l.values = append(l.values[:idx], l.values[idx+1:]...)
			return true
		}
	}
	return false
}

func (l *arrayList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= l.Size() {
		var val T
		return val, false
	}
	return l.values[index], true
}

func (l *arrayList[T]) Set(index int, obj T) (T, bool) {
	if index < 0 || index >= l.Size() {
		var val T
		return val, false
	}

	old := l.values[index]
	l.values[index] = obj
	return old, true
}

func (l *arrayList[T]) Range(f func(index int, t T) bool) {
	for idx, obj := range l.values {
		f(idx, obj)
	}
}
