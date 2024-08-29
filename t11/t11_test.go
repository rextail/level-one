package t11

import (
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	testcases := []struct {
		s1       []int
		s2       []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{3, 4},
			[]int{3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{},
			[]int{},
			[]int{},
		},
	}
	for _, tc := range testcases {
		got := Intersection(tc.s1, tc.s2)
		if slices.Compare(got, tc.expected) != 0 {
			t.Errorf("got: %v, expected: %v ", got, tc.expected)
		}
	}
}
