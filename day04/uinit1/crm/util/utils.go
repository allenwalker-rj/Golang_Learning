package util

func Sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Sub(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum -= num
	}
	return sum
}
