package t20

import "testing"

func TestReverseWords(t *testing.T) {
	testcases := []struct {
		str      string
		expected string
	}{
		{
			"abc da",
			"da abc",
		},
		{
			"👽 👺 👽 💨",
			"💨 👽 👺 👽",
		},
		{
			"汉 汉 a s v",
			"v s a 汉 汉",
		},
		{
			"  汉 汉 a s v",
			"v s a 汉 汉  ",
		},
		{
			"😀 😀", // Повторяющийся эмодзи 😀
			"😀 😀",
		},
		{
			"Улыбок тебе дед Макар",
			"Макар дед тебе Улыбок",
		},
	}

	for _, tc := range testcases {
		got := ReverseWords(tc.str)
		if got != tc.expected {
			t.Errorf("String: %s; got: %v, expected: %v", tc.str, got, tc.expected)
		}
	}
}
