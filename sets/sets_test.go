package sets

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	tests := []struct {
		origin   *hashSet[int]
		target   *hashSet[int]
		expected *hashSet[int]
	}{
		{
			origin: &hashSet[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
			target: &hashSet[int]{
				m: map[int]struct{}{
					2: {},
				},
			},
			expected: &hashSet[int]{
				m: map[int]struct{}{
					1: {},
					2: {},
				},
			},
		},
	}

	for _, test := range tests {
		got := Union[int](test.origin, test.target)
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("unexpected values for union, expected: %v, but got: %v", test.expected, got)
		}
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		origin   *hashSet[int]
		target   *hashSet[int]
		expected *hashSet[int]
	}{
		{
			origin: &hashSet[int]{
				m: map[int]struct{}{
					1: {},
					2: {},
				},
			},
			target: &hashSet[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
			expected: &hashSet[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
		},
	}

	for _, test := range tests {
		got := Intersect[int](test.origin, test.target)
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("unexpected values for intersect, expected: %v, but got: %v", test.expected, got)
		}
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		origin   *hashSet[int]
		target   *hashSet[int]
		expected *hashSet[int]
	}{
		{
			origin: &hashSet[int]{
				m: map[int]struct{}{
					1: {},
					2: {},
				},
			},
			target: &hashSet[int]{
				m: map[int]struct{}{
					1: {},
					3: {},
				},
			},
			expected: &hashSet[int]{
				m: map[int]struct{}{
					3: {},
				},
			},
		},
	}

	for _, test := range tests {
		got := Difference[int](test.origin, test.target)
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("unexpected values for difference, expected: %v, but got: %v", test.expected, got)
		}
	}
}

func TestIsSubsetOf(t *testing.T) {
	tests := []struct {
		origin   *hashSet[int]
		target   *hashSet[int]
		expected bool
	}{
		{
			origin: &hashSet[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
			target: &hashSet[int]{
				m: map[int]struct{}{
					1: {},
					2: {},
				},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		got := IsSubsetOf[int](test.origin, test.target)
		if !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("unexpected values for isSubsetOf, expected: %v, but got: %v", test.expected, got)
		}
	}
}
