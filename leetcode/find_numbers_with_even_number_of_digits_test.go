package leetcode

import "testing"

func TestFindNumbers(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			[]int{12, 345, 2, 6, 7896},
			2,
		},
		{
			[]int{12, 3456, 2, 6, 7896},
			3,
		},
		{
			[]int{},
			0,
		},
	}

	for _, c := range cases {
		if actual := findNumbers(c.input); actual != c.expected {
			t.Errorf("expected: %d, actual: %d", c.expected, actual)
		}
	}
}
