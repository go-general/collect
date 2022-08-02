package lists

import (
	"testing"
)

func TestSyncList(t *testing.T) {
	list := NewSyncList[int]()
	t.Log(list.Values())
	list.Range(func(i int, obj int) bool {
		t.Log(i, obj)
		return true
	})
}
