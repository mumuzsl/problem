package lc

import "sort"

func minDifference(nums []int) int {
	N := len(nums)
	if N <= 4 {
		return 0
	}
	sort.Ints(nums)
	var r int
	limit := N - 4
	diff := nums[limit] - nums[0]
	for i := 1; i < 4; i++ {
		r = nums[i+limit] - nums[i]
		if r < diff {
			diff = r
		}
	}
	return diff
}
