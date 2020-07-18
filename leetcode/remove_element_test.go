package leetcode

import "testing"

func TestRemoteElement(t *testing.T) {
	cases := []struct {
		nums     []int
		val      int
		expected int
	}{
		{
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
		},
		{
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			expected: 5,
		},
	}

	for _, c := range cases {
		actual := removeElement(c.nums, c.val)
		if c.expected != actual {
			t.Errorf("actual: %d, expected: %d", actual, c.expected)
		}
	}
}
