package lists

type FixedArrayList[T comparable] struct {
	size   int
	values []T
}
