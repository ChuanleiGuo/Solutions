package leetcode

import "testing"

func TestHeightChecker(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 1, 4, 2, 1, 3},
			expected: 3,
		},
		{
			input:    []int{5, 1, 2, 3, 4},
			expected: 5,
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: 0,
		},
	}

	for i, c := range cases {
		actual := heightChecker(c.input)
		if actual != c.expected {
			t.Errorf("idx: %d, actual: %d, expected: %d", i, actual, c.expected)
		}
	}
}
