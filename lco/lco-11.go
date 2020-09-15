package lco

func minArray(numbers []int) int {
	numbersLen := len(numbers)

	r := numbers[0]

	for i := 1; i < numbersLen; i++ {
		if numbers[i] < numbers[i-1] && numbers[i] < r {
			r = numbers[i]
			break
		}
	}

	return r
}
