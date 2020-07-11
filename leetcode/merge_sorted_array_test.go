package leetcode

import "testing"

func TestMerge(t *testing.T) {
	cases := []struct {
		num1     []int
		m        int
		num2     []int
		n        int
		expected []int
	}{
		{
			num1:     []int{1, 2, 3, 0, 0, 0},
			m:        3,
			num2:     []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			num1:     []int{0},
			m:        0,
			num2:     []int{1},
			n:        1,
			expected: []int{1},
		},
	}

	for _, c := range cases {
		merge(c.num1, c.m, c.num2, c.n)
		for i := range c.num1 {
			if c.num1[i] != c.expected[i] {
				t.Errorf("error")
			}
		}
	}
}
