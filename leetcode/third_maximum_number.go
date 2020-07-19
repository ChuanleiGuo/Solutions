package leetcode

import "math"

func thirdMax(nums []int) int {
	max1 := math.MinInt32
	max2 := math.MinInt32
	max3 := math.MinInt32

	for _, n := range nums {
		if n > max1 {
			max1 = n
		}
	}

	for _, n := range nums {
		if n >= max2 && n < max1 {
			max2 = n
		}
	}

	for _, n := range nums {
		if n >= max3 && n < max1 && n < max2 {
			max3 = n
			if n == math.MinInt32 {
				return math.MinInt32
			}
		}
	}
	if max3 == math.MinInt32 || max2 == math.MinInt32 {
		return max1
	}
	return max3
}
