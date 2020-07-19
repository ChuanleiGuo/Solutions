package leetcode

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
		{
			input:    nil,
			expected: nil,
		},
	}

	for i, c := range cases {
		moveZeroes(c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			t.Errorf("idx: %d error", i)
		}
	}
}
