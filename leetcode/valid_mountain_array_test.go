package leetcode

import "testing"

func TestValidMountainArray(t *testing.T) {
	cases := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{2, 1},
			expected: false,
		},
		{
			input:    []int{3, 5, 5},
			expected: false,
		},
		{
			input:    []int{0, 3, 2, 1},
			expected: true,
		},
		{
			input:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: false,
		},
	}

	for i, c := range cases {
		actual := validMountainArray(c.input)
		if actual != c.expected {
			t.Errorf("idx: %d, actual: %v, expected: %v", i, actual, c.expected)
		}
	}
}
