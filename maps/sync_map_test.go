package maps

import (
	"testing"
)

func TestSyncMap(t *testing.T) {
	m := NewSyncMap[int, int]()
	t.Log(m.Keys())
	t.Log(m.Values())
}
