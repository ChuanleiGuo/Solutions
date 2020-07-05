package leetcode

func sortedSquares(A []int) []int {
	posIdx := 0
	for posIdx < len(A) && A[posIdx] < 0 {
		posIdx++
	}
	negIdx := posIdx - 1
	var res []int

	for negIdx >= 0 && posIdx < len(A) {
		if A[negIdx]*A[negIdx] < A[posIdx]*A[posIdx] {
			res = append(res, A[negIdx]*A[negIdx])
			negIdx--
		} else {
			res = append(res, A[posIdx]*A[posIdx])
			posIdx++
		}
	}

	for negIdx >= 0 {
		res = append(res, A[negIdx]*A[negIdx])
		negIdx--
	}

	for posIdx < len(A) {
		res = append(res, A[posIdx]*A[posIdx])
		posIdx++
	}
	return res
}
