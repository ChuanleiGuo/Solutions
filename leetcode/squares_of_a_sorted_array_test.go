package leetcode

import "testing"

func TestSortedSquares(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{-4, -1, 0, 3, 10},
			[]int{0, 1, 9, 16, 100},
		},
		{
			[]int{-7, -3, 2, 3, 11},
			[]int{4, 9, 9, 49, 121},
		},
	}

	for _, c := range cases {
		actual := sortedSquares(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected len: %d, actual: %d", len(c.expected), len(actual))
		}
		for i := range c.expected {
			if c.expected[i] != actual[i] {
				t.Errorf("expected: %d, actual: %d", c.expected[i], actual[i])
			}
		}
	}
}
