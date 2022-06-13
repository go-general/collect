package lists

func NewArrayList[T comparable](size int) List[T] {
	return &ArrayList[T]{
		values: make([]T, 0, size),
	}
}

func NewImmutableArrayList[T comparable](values ...T) List[T] {
	vals := make([]T, 0, len(values))
	for _, v := range values {
		vals = append(vals, v)
	}
	return &ImmutableArrayList[T]{
		values: vals,
	}
}
