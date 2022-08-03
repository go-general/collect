package sets

import (
	"reflect"
	"testing"
)

func TestImmutableHashSet(t *testing.T) {
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
		set := NewImmutableHashSet[int](test.input...)
		got := set.Values()

		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("unexpected values for immutableHashSet, expected: %v, but got: %v", test.expected, got)
		}
	}
}
