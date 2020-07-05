package leetcode

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 1, 0, 1, 1, 1},
			expected: 3,
		},
		{
			input:    []int{1, 1, 1},
			expected: 3,
		},
		{
			input:    []int{1, 1, 1, 1, 0, 1, 1, 1},
			expected: 4,
		},
		{
			input:    []int{},
			expected: 0,
		},
		{
			input:    []int{1, 0, 1, 1, 0, 1},
			expected: 2,
		},
	}

	for _, c := range cases {
		if actual := findMaxConsecutiveOnes(c.input); actual != c.expected {
			t.Errorf("expected: %d, actual: %d", c.expected, actual)
		}
	}
}
