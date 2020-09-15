package lc

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 0 || len(nums2) == 0 {
		return
	}
	var key, j int
	for i := m; i < m+n; i++ {
		key = nums2[i-m]
		j = i - 1
		for j >= 0 && nums1[j] > key {
			nums1[j+1] = nums1[j]
			j--
		}
		nums1[j+1] = key
	}
}
