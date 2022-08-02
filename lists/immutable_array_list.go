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

func (l *immutableArrayList[T]) Clear() {
	// immutable arrayList can not do clear action
}

func (l *immutableArrayList[T]) Values() []T {
	return l.values
}

func (l *immutableArrayList[T]) Add(obj ...T) {
}

func (l *immutableArrayList[T]) AddAll(list List[T]) {
}

func (l *immutableArrayList[T]) Contains(obj T) bool {
	for idx := range l.values {
		if l.values[idx] == obj {
			return true
		}
	}
	return false
}

func (l *immutableArrayList[T]) Remove(obj T) bool {
	return false
}

func (l *immutableArrayList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= l.Size() {
		var val T
		return val, false
	}
	return l.values[index], true
}

func (l *immutableArrayList[T]) Set(index int, obj T) (T, bool) {
	if index < 0 || index >= l.Size() {
		var val T
		return val, false
	}

	old := l.values[index]
	return old, true
}

func (l *immutableArrayList[T]) Range(f func(index int, t T) bool) {
	for idx, obj := range l.values {
		f(idx, obj)
	}
}
