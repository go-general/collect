package maps

import (
	"testing"
)

func TestImmutableHashMap(t *testing.T) {
	m := NewImmutableHashMap[int, int](map[int]int{
		1: 1,
		2: 2,
		3: 3,
	})
	t.Log(m.Keys())
	t.Log(m.Values())
}
