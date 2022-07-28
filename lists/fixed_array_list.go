package lists

type fixedArrayList[T comparable] struct {
	size   int
	values []T
}
