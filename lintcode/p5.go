package lintcode

import (
	"sort"
)

/**
 * @param n: An integer
 * @param nums: An array
 * @return: the Kth largest element
 */
func kthLargestElement(n int, nums []int) int {
	// write your code here
	numsLen := len(nums)
	sort.Ints(nums)
	return nums[numsLen-n]
}
