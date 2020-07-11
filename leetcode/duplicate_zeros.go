package leetcode

func duplicateZeros(arr []int) {
	dupCount := 0
	length := len(arr) - 1

	for i := 0; i <= length-dupCount; i++ {
		if arr[i] == 0 {
			if i == length-dupCount {
				arr[length] = 0
				length--
				break
			}
			dupCount++
		}
	}

	last := length - dupCount

	for i := last; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+dupCount] = 0
			dupCount--
			arr[i+dupCount] = 0
		} else {
			arr[i+dupCount] = arr[i]
		}
	}
}
