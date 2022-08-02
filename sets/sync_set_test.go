package sets

import (
	"testing"
)

func TestSyncSet(t *testing.T) {
	set := NewSyncSet[int]()
	t.Log(set.IsEmpty())

	set.Add(1)
	t.Log(set.IsEmpty())

	set.Add(2)
	set.Add(3)
	t.Log(set.Size())
}
