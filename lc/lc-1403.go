package lc

func minSubsequence(nums []int) []int {
	var key, j int
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		key = nums[i]
		j = i - 1
		for j >= 0 && nums[j] < key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
	ans := make([]int, 0)
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp += nums[i]
		ans = append(ans, nums[i])
		if temp > sum-temp {
			break
		}
	}
	return ans
}
