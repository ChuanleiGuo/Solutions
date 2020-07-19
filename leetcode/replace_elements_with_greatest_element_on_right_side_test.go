package leetcode

import (
	"reflect"
	"testing"
)

func TestReplaceElements(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{17, 18, 5, 4, 6, 1},
			expected: []int{18, 6, 6, 6, 1, -1},
		},
	}

	for _, c := range cases {
		actual := replaceElements(c.input)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("err")
		}
	}
}
