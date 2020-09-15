package lc

func minDeletionSize(A []string) int {
	ans := 0
	n := len(A[0])
	Alen := len(A)
	for i := 0; i < n; i++ {
		for j := 1; j < Alen; j++ {
			if A[j-1][i] > A[j][i] {
				ans++
				break
			}
		}
	}
	return ans
}
