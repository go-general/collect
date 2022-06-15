package sets

import (
	"reflect"
	"testing"
)

func TestHashSet_Difference(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		hs   []T
		args []T
		want []T
	}
	tests := []testCase[string]{
		{
			name: "test",
			hs:   []string{"a", "b", "c", "d", "e", "f"},
			args: []string{"b", "c", "d", "e"},
			want: []string{"a", "f"},
		},
		{
			name: "test2",
			hs:   []string{"1", "2", "3"},
			args: []string{"2", "3", "4"},
			want: []string{"1", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hs := newHashSetFromSlice[string](tt.hs)
			target := newHashSetFromSlice[string](tt.args)
			got := hs.Difference(target).Values() // todo need sort values
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
