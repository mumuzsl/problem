package lc

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	candidatesLen := len(candidates)
	temp := make([]int, 0)
	sort.Ints(candidates)
	var dfs func(start, sum int) bool
	dfs = func(start, sum int) bool {
		if sum > target {
			return false
		}
		if sum == target {
			return true
		}
		for i := start; i < candidatesLen; i++ {
			//if i > start && candidates[i] == candidates[i-1] {
			//	continue
			//}
			for i > start && candidates[i] == candidates[i-1] {
				i++
				if i == candidatesLen {
					return false
				}
			}
			temp = append(temp, candidates[i])
			if dfs(i+1, candidates[i]+sum) {
				ans = append(ans, make([]int, len(temp)))
				copy(ans[len(ans)-1], temp)
			}
			if len(temp) >= 1 {
				temp = temp[:len(temp)-1]
			}
		}
		return false
	}
	dfs(0, 0)
	return ans
}
