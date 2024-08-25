package t16

import (
	"slices"
	"testing"
)

func TestQuicksort(t *testing.T) {
	t.Run("test integers with standard go quicksort", func(t *testing.T) {
		testcases := []struct {
			arr      []int
			expected []int
		}{
			{
				[]int{5, 3, 2, 1, 2, 7, 6, 10, 9, 8, 13, 11, 12, 15, 14, 16},
				[]int{1, 2, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			},
			{
				[]int{3, 1, 2, 5, 4},
				[]int{1, 2, 3, 4, 5},
			},
			{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		}

		for _, tc := range testcases {
			QuickByGoTeam(tc.arr)
			if slices.Compare(tc.arr, tc.expected) != 0 {
				t.Errorf("expected %v, got %v", tc.expected, tc.arr)
			}
		}
	})
	t.Run("test integers", func(t *testing.T) {
		testcases := []struct {
			arr      []int
			expected []int
		}{
			{
				[]int{5, 3, 2, 1, 2, 7, 6, 10, 9, 8, 13, 11, 12, 15, 14, 16},
				[]int{1, 2, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			},
			{
				[]int{3, 1, 2, 5, 4},
				[]int{1, 2, 3, 4, 5},
			},
			{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		}

		for _, tc := range testcases {
			Quicksort(tc.arr, 0, len(tc.arr)-1)
			if slices.Compare(tc.arr, tc.expected) != 0 {
				t.Errorf("expected %v, got %v", tc.expected, tc.arr)
			}
		}
	})

}
