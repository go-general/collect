package sets

import (
	"testing"
)

func TestImmutableHashSet(t *testing.T) {
	set := NewImmutableHashSet[int](1, 2, 3, 4, 5)
	t.Log(set.Values())
}
