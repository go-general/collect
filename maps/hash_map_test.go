package maps

import (
	"testing"
)

func TestHashMap(t *testing.T) {
	m := NewHashMap[int, int]()
	m.Put(1, 1)
	m.Put(2, 2)
	t.Log(m.Get(3))
}
