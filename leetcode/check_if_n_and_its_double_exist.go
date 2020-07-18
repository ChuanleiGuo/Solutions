package leetcode

func checkIfExist(arr []int) bool {
	intSet := make(map[int]bool, len(arr))
	for _, v := range arr {
		if exist := intSet[v*2]; exist && v*2 != v {
			return true
		}
		if exist := intSet[v/2]; exist && v%2 == 0 {
			return true
		}
		intSet[v] = true
	}
	return false
}
