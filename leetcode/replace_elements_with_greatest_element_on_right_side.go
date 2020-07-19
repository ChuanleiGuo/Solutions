package leetcode

func replaceElements(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	rMax := -1
	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		arr[i] = rMax
		if v >= rMax {
			rMax = v
		}
	}
	return arr
}
