package sets

import (
	"reflect"
	"sort"
	"testing"
)

func TestSyncSet(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		set := NewSyncSet[int]()
		for _, d := range test.input {
			set.Add(d)
		}

		got := set.Values()
		sort.Ints(got)
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("unexpected values for syncSet, expected: %v, but got: %v", test.expected, set.Values())
		}
	}
}
