package leetcode

import "sort"

func heightChecker(heights []int) int {
	newH := make([]int, len(heights))
	copy(newH, heights)
	sort.Ints(newH)

	diff := 0
	for i := range heights {
		if heights[i] != newH[i] {
			diff++
		}
	}
	return diff
}
