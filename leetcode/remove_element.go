package leetcode

func removeElement(nums []int, val int) int {
	i := 0
	tail := len(nums)
	for i < tail {
		if nums[i] == val {
			nums[i] = nums[tail-1]
			tail--
		} else {
			i++
		}
	}
	return tail
}
