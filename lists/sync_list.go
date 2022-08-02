package lists

import (
	"sync"
)

type syncList[T comparable] struct {
	lock   sync.RWMutex
	values []T
}

func (l *syncList[T]) IsEmpty() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return len(l.values) == 0
}

func (l *syncList[T]) Size() int {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return len(l.values)
}

func (l *syncList[T]) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.values = make([]T, 0, 0)
}

func (l *syncList[T]) Values() []T {
	l.lock.RLock()
	defer l.lock.RUnlock()

	return l.values
}

func (l *syncList[T]) Add(obj ...T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	for _, val := range obj {
		l.values = append(l.values, val)
	}
}

func (l *syncList[T]) AddAll(list List[T]) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.values = append(l.values, list.Values()...)
}

func (l *syncList[T]) Contains(obj T) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()

	for idx := range l.values {
		if l.values[idx] == obj {
			return true
		}
	}
	return false
}

func (l *syncList[T]) Remove(obj T) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	for idx := range l.values {
		if l.values[idx] == obj {
			l.values = append(l.values[:idx], l.values[idx+1:]...)
			return true
		}
	}
	return false
}

func (l *syncList[T]) Get(index int) (T, bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if index < 0 || index >= l.Size() {
		var val T
		return val, false
	}
	return l.values[index], true
}

func (l *syncList[T]) Set(index int, obj T) (T, bool) {
	l.lock.Lock()
	defer l.lock.Unlock()
	
	if index < 0 || index >= l.Size() {
		var val T
		return val, false
	}

	old := l.values[index]
	l.values[index] = obj
	return old, true
}

func (l *syncList[T]) Range(f func(index int, t T) bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	for idx, obj := range l.values {
		f(idx, obj)
	}
}
