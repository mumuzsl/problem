package lc

func combine(n int, k int) [][]int {
	flag := make([]bool, n)
	ans := make([][]int, 0)
	temp := make([]int, k)
	var dfs func(start, deep int)
	dfs = func(start, deep int) {
		if deep == k {
			ans = append(ans, make([]int, k))
			copy(ans[len(ans)-1], temp)
			return
		}
		temp[deep] = start + 1
		flag[start] = true
		for i := start; i < n; i++ {
			if !flag[i] {
				//flag[i] = true
				//temp[deep] = i + 1
				dfs(i+1, deep+1)
				//flag[i] = false
			}
		}
		flag[start] = false
	}
	dfs(0, 0)
	return ans
}
