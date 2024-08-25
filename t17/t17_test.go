package t17

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	got := BinarySearch(arr, 3)
	expected := 2
	if got != expected {
		t.Errorf("got %d, expected %d", got, expected)
	}
}
