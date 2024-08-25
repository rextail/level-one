package t19

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	testcases := []struct {
		str      string
		expected string
	}{
		{
			"abcda",
			"adcba",
		},
		{
			"👽👺👽💨",
			"💨👽👺👽",
		},
		{
			"汉汉asv",
			"vsa汉汉",
		},
		{
			"😀😃😄😁😆", //😀😃 разные
			"😆😁😄😃😀",
		},
		{
			"😀  😀", // Повторяющийся эмодзи 😀
			"😀  😀",
		},
	}

	for _, tc := range testcases {
		got := ReverseString(tc.str)
		if got != tc.expected {
			t.Errorf("String: %s; got: %v, expected: %v", tc.str, got, tc.expected)
		}
	}
}
