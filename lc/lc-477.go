package lc

func totalHammingDistance(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sum := 0
	cnt := make([]int, 32)
	for _, num := range nums {
		for i := 0; num > 0; i++ {
			cnt[i] += (num & 1)
			num >>= 1
		}
	}
	for _, k := range cnt {
		sum += k * (n - k)
	}
	return sum
}

func HammingDistance(a, b int) int {
	t, count := a^b, 0
	for t != 0 {
		t = t & (t - 1)
		count++
	}
	return count
}
