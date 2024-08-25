package t23

import (
	"reflect"
	"testing"
)

func TestDeleteByIndex(t *testing.T) {
	testcases := []struct {
		index    int
		arr      []int
		expected []int
	}{
		{
			-1,
			[]int{1, 3, 4, 5},
			[]int{1, 3, 4, 5},
		},
		{
			0,
			[]int{1, 3, 4, 5},
			[]int{3, 4, 5},
		},
		{
			1,
			[]int{1, 3, 4, 5},
			[]int{1, 4, 5},
		},
		{
			2,
			[]int{1, 3, 4, 5},
			[]int{1, 3, 5},
		},
		{
			3,
			[]int{1, 3, 4, 5},
			[]int{1, 3, 4},
		},
		// Большой массив: удаление первого элемента
		{
			0,
			func() []int {
				arr := make([]int, 100000)
				for i := range arr {
					arr[i] = i + 1
				}
				return arr
			}(),
			func() []int {
				arr := make([]int, 99999)
				for i := range arr {
					arr[i] = i + 2
				}
				return arr
			}(),
		},
		// Большой массив: удаление последнего элемента
		{
			99999,
			func() []int {
				arr := make([]int, 100000)
				for i := range arr {
					arr[i] = i + 1
				}
				return arr
			}(),
			func() []int {
				arr := make([]int, 99999)
				for i := range arr {
					arr[i] = i + 1
				}
				return arr
			}(),
		},
	}
	t.Run("test with old base array", func(t *testing.T) {
		for _, tc := range testcases {
			arr := make([]int, len(tc.arr))
			copy(arr, tc.arr)
			got := DelWithOldBase(arr, tc.index)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		}
	})
	t.Run("test with new base array", func(t *testing.T) {
		for _, tc := range testcases {
			got := DelWithNewBase(tc.arr, tc.index)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		}
	})

}
