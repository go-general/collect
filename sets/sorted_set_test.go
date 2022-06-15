package sets

import (
	"reflect"
	"testing"

	"github.com/go-general/collect/types"
)

func Test_sortedSet_Values(t *testing.T) {
	type fields[T types.Ordered] struct {
		args []T
		want []T
	}
	tests := []struct {
		name   string
		fields fields[string]
	}{
		{
			name: "test",
			fields: fields[string]{
				args: []string{"f", "e", "d", "c", "b", "a"},
				want: []string{"a", "b", "c", "d", "e", "f"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newSortedSet[string]()
			for _, v := range tt.fields.args {
				s.Add(v)
			}

			if got := s.Values(); !reflect.DeepEqual(got, tt.fields.want) {
				t.Errorf("Values() = %v, want %v", got, tt.fields)
			}
		})
	}

	tests2 := []struct {
		name   string
		fields fields[int]
	}{
		{
			name: "test",
			fields: fields[int]{
				args: []int{1, 4, 3, 2, 5, 6},
				want: []int{1, 2, 3, 4, 5, 6},
			},
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			s := newSortedSet[int]()
			for _, v := range tt.fields.args {
				s.Add(v)
			}

			if got := s.Values(); !reflect.DeepEqual(got, tt.fields.want) {
				t.Errorf("Values() = %v, want %v", got, tt.fields)
			}
		})
	}
}
