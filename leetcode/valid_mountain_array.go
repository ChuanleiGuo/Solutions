package leetcode

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	top := 0
	for ; top+1 < len(arr) && arr[top] < arr[top+1]; top++ {

	}

	if top == 0 || top+1 == len(arr) {
		return false
	}

	for ; top+1 < len(arr) && arr[top] > arr[top+1]; top++ {

	}

	return top+1 == len(arr)
}
