package lists

import (
	"reflect"
	"testing"
)

func TestSyncList(t *testing.T) {
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
		list := NewSyncList[int]()
		for _, d := range test.input {
			list.Add(d)
		}
		got := list.Values()

		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("unexpected values for syncList, expected: %v, but got: %v", test.expected, got)
		}
	}
}
