package lists

import (
	"testing"
)

func TestImmutableArrayList(t *testing.T) {
	list := NewImmutableArrayList[int](1, 2, 3, 4, 5)
	t.Log(list.Values())
	list.Range(func(i int, obj int) bool {
		t.Log(i, obj)
		return true
	})
}
