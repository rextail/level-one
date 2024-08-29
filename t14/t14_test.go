package t14

import "testing"

func TestDetectType(t *testing.T) {
	testcases := []struct {
		val      any
		expected string
	}{
		{
			"str",
			"string",
		},
		{
			15,
			"int",
		},
		{
			make(chan struct{}),
			"chan struct{}",
		},
	}
	for _, tc := range testcases {
		got := DetectType(tc.val)
		if got != tc.expected {
			t.Errorf("got: %s, expected: %s", got, tc.expected)
		}
	}
}
