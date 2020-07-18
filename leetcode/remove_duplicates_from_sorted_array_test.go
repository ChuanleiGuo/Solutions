package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for _, c := range cases {
		actual := removeDuplicates(c.input)
		if actual != c.expected {
			t.Errorf("expected: %d, actual: %d", c.expected, actual)
		}
	}
}
