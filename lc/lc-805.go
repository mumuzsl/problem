package lc

import (
	"sort"
)

func splitArraySameAverage(A []int) bool {
	N := len(A)
	sum := 0
	for i := 0; i < N; i++ {
		sum += A[i]
	}
	var dfs func(begin, step, target int) bool
	dfs = func(begin, step, target int) bool {
		if step == 0 {
			return target == 0
		}
		if target < step*A[begin] {
			return false
		}
		for i := begin; i <= N-step; i++ {
			if i > begin && A[i] == A[i-1] {
				continue
			}
			if dfs(i+1, step-1, target-A[i]) {
				return true
			}
		}
		return false
	}
	sort.Ints(A)
	for i := 1; i <= N/2; i++ {
		if sum*i%N == 0 && dfs(0, i, sum*i/N) {
			return true
		}
	}
	return false
}
