package maps

import (
	"reflect"
	"sort"
	"testing"
)

func TestHashMap(t *testing.T) {
	tests := []struct {
		input          map[int]int
		expectedKeys   []int
		expectedValues []int
	}{
		{
			input:          map[int]int{1: 1, 2: 2, 3: 3},
			expectedKeys:   []int{1, 2, 3},
			expectedValues: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		m := NewHashMap[int, int]()
		for k, v := range test.input {
			m.Put(k, v)
		}

		keys := m.Keys()
		sort.Ints(keys)
		if !reflect.DeepEqual(keys, test.expectedKeys) {
			t.Fatalf("unexpected keys for hashMap, expected: %v, but got: %v", test.expectedKeys, m.Keys())
		}

		values := m.Values()
		sort.Ints(values)
		if !reflect.DeepEqual(values, test.expectedValues) {
			t.Fatalf("unexpected values for hashMap, expected: %v, but got: %v", test.expectedValues, m.Values())
		}
	}
}
