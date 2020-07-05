package leetcode

func findNumbers(nums []int) int {
	counter := 0
	for _, num := range nums {
		if nd := numDigit(num); nd%2 == 0 {
			counter++
		}
	}
	return counter
}

func numDigit(num int) int {
	c := 0
	for num > 0 {
		c++
		num /= 10
	}
	return c
}
