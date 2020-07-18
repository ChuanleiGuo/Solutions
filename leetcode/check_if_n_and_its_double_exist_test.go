package leetcode

import "testing"

func TestCheckIfExist(t *testing.T) {
	cases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{10, 2, 5, 3},
			expected: true,
		},
		{
			input:    []int{7, 1, 14, 11},
			expected: true,
		},
		{
			input:    []int{3, 1, 7, 11},
			expected: false,
		},
		{
			input:    []int{-2, 0, 10, -19, 4, 6, -8},
			expected: false,
		},
	}

	for _, c := range cases {
		actual := checkIfExist(c.input)
		if actual != c.expected {
			t.Errorf("expected: %v, actual: %v", c.expected, actual)
		}
	}
}
