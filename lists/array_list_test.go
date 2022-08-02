package lists

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	list := NewArrayList[int](10)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	t.Log(list.Values())
	list.Range(func(index int, obj int) bool {
		t.Log(index, obj)
		return true
	})
}
