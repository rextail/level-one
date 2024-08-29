package t8

import (
	"testing"
)

func TestSetBit(t *testing.T) {
	testcases := []struct {
		num      int64
		i        int64
		expected int64
	}{
		{
			1,
			2,
			5,
		},
		{
			2,
			2,
			6,
		},
	}
	for _, tc := range testcases {
		SetBit(&tc.num, tc.i)
		if tc.num != tc.expected {
			t.Errorf("got %v, expected %v", tc.num, tc.expected)
		}
	}
}
