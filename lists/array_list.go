package lists

type ArrayList[T comparable] struct {
	size   int
	values []T
}
