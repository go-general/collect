package sets

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	set := NewHashSet[int]()
	t.Log(set.IsEmpty())

	set.Add(1)
	t.Log(set.IsEmpty())

	set.Add(2)
	set.Add(3)
	t.Log(set.Size())
}
