package leetcode

func moveZeroes(nums []int) {
	wIdx := 0
	for rIdx := 0; rIdx < len(nums); rIdx++ {
		if nums[rIdx] != 0 {
			nums[wIdx], nums[rIdx] = nums[rIdx], nums[wIdx]
			wIdx++
		}
	}
}
