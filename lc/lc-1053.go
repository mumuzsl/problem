package lc

func prevPermOpt1(A []int) []int {
	Alen := len(A)
	var max, index int
	for j := Alen - 1; j >= 1; j-- {
		if A[j-1] > A[j] {
			max, index = A[j], j
			for i := j; i < Alen; i++ {
				if A[j-1] > A[i] {
					if A[i] > max {
						max, index = A[i], i
					}
				}
			}
			A[j-1], A[index] = A[index], A[j-1]
			break
		}
	}
	return A
}
