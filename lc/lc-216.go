package lc

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	temp := make([]int, k)
	var dfs func(start, sum, deep int)
	dfs = func(start, sum, deep int) {
		if sum > n {
			return
		}
		if deep == k {
			if sum == n {
				ans = append(ans, make([]int, k))
				copy(ans[len(ans)-1], temp)
			}
			return
		}
		for i := start; i <= 9; i++ {
			temp[deep] = i
			dfs(start+1, sum+i, deep+1)
		}
	}
	dfs(1, 0, 0)
	return ans
}
