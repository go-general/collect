package sets

import (
	"reflect"
	"testing"
)

func TestHashSet(t *testing.T) {
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
		set := NewHashSet[int]()
		for _, d := range test.input {
			set.Add(d)
		}
		got := set.Values()

		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("unexpected values for set, expected: %v, but got: %v", test.expected, got)
		}
	}
}
