package lc

func massage(nums []int) int {
	numsLen := len(nums)

	if numsLen == 0 {
		return 0
	}

	if numsLen == 1 {
		return nums[0]
	}

	dp := make([]int, 3)
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])
	for i := 2; i < numsLen; i++ {
		dp[i%3] = Max(dp[(i-1)%3], dp[(i-2)%3]+nums[i])
	}
	return dp[(numsLen-1)%3]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
