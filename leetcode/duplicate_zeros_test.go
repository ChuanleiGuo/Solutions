package leetcode

import "testing"

func TestDuplicateZeros(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 0, 2, 3, 0, 4, 5, 0},
			expected: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{1, 0, 2, 0, 0},
			expected: []int{1, 0, 0, 2, 0},
		},
	}

	for _, c := range cases {
		duplicateZeros(c.input)
		if len(c.input) != len(c.expected) {
			t.Errorf("err")
		}
		for i := range c.input {
			if c.input[i] != c.expected[i] {
				t.Errorf("err")
			}
		}
	}
}
