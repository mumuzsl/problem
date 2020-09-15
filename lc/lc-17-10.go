package lc

func majorityElement(nums []int) int {
	numsLen := len(nums)
	middle := numsLen / 2
	temp := nums[0]
	count := 1

	for i := 0; i < numsLen; i++ {
		if nums[i] == temp {
			count++
		} else {
			count--
		}
		if count == 0 {
			temp = nums[i]
			count = 1
		}
	}

	count = 0
	for _, v := range nums {
		if v == temp {
			count++
		}
		if count > middle {
			return temp
		}
	}

	return -1

	//count := 0
	//var r = nums[0]
	//for _, i := range nums {
	//	if i == r {
	//		count += 1
	//	} else {
	//		count -= 1
	//	}
	//	if count == 0 {
	//		r = i
	//		count += 1
	//	}
	//}
	//
	//if count == 0 {
	//	return -1
	//}
	//return r
}
