package lco

func findRepeatNumber(nums []int) int {
	numsLen := len(nums)
	tempMap := make(map[int]int, 0)
	r := nums[0]

	for i := 0; i < numsLen; i++ {
		k := nums[i]
		tempMap[k]++
		if tempMap[k] > 1 {
			return k
		}
	}

	return r
}

func findRepeatNumber2(nums []int) int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			temp := nums[i]
			nums[i] = nums[temp]
			nums[temp] = temp
		}
	}
	return -1
}
