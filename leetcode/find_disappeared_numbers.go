package leetcode

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		j := absInt(nums[i]) - 1
		nums[j] = absInt(nums[j]) * -1
	}

	var res []int
	for i, n := range nums {
		if n > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func absInt(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
