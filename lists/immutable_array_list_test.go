package lists

import (
	"reflect"
	"testing"
)

func TestImmutableArrayList(t *testing.T) {
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
		list := NewImmutableArrayList[int](test.input...)
		got := list.Values()

		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("unexpected values for immutableArrayList, expected: %v, but got: %v", test.expected, got)
		}
	}
}
