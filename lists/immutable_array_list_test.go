package lists

import (
	"testing"
)

func TestImmutableArrayList(t *testing.T) {
	list := NewImmutableArrayList[int](1, 2, 3, 4, 5)
	t.Log(list.Size())
	t.Log(list.Contains(1))
	t.Log(list.Contains(10))
}
