package lists

import (
	"testing"
)

func TestArrayList(t *testing.T) {
	var list List[int]
	list = NewArrayList[int](10)
	list.Add(1)
	list.Add(2)
	list.Add(3)
	t.Log(list.Size())
}
