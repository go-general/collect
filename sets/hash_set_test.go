package sets

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	set := NewHashSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	t.Log(set.Size())
}
