package lco

import "math"

func printNumbers(n int) []int {
	m := math.Pow10(n)
	result := make([]int, int(m-1))
	for i := 1; i < int(m); i++ {
		result[i-1] = i
	}
	return result
}
