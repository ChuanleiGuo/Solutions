package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	maxNum := 0

	start := 0
	for start < len(nums) {
		if nums[start] != 1 {
			start++
		}

		end := start
		for end < len(nums) && nums[end] == 1 {
			end++
		}

		if end-start > maxNum {
			maxNum = end - start
		}

		start = end + 1
	}
	return maxNum
}

func findMaxConsecutiveOnes2(nums []int) int {
	maxNum := 0

	m := 0
	for _, n := range nums {
		if n == 1 {
			m++
			if m > maxNum {
				maxNum = m
			}
		} else {
			m = 0
		}
	}
	return maxNum
}
