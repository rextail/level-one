package t26

import "testing"

func TestIsUnique(t *testing.T) {
	testcases := []struct {
		str      string
		expected bool
	}{
		{
			"abcda",
			false,
		},
		{
			"abcd",
			true,
		},
		{
			"汉汉asv",
			false,
		}, {
			"AAAA",
			false,
		},
		{
			"Aa",
			false,
		},
		{
			"Önekİişée",
			false,
		},
		{
			"😀😃😄😁😆", //😀😃 разные
			true,
		},
		{
			"😀😃😄😁😀", // Повторяющийся эмодзи 😀
			false,
		},
		{
			"éééé",
			false,
		},
		{
			"eé",
			true,
		},
	}

	for _, tc := range testcases {
		got := IsUnique(tc.str)
		if got != tc.expected {
			t.Errorf("String: %s; got: %v, expected: %v", tc.str, got, tc.expected)
		}
	}

}
