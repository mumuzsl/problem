package lc

import "strconv"

func splitIntoFibonacci(S string) []int {
	sLen := len(S)
	ans := make([]int, 0)
	var dfs func(p, pre1, pre2, deep int) bool
	dfs = func(p, pre1, pre2, deep int) bool {
		if p == sLen {
			return deep >= 3
		}
		for i := 1; i <= 11; i++ {
			if p+i > sLen || (S[p] == '0' && i > 1) {
				break
			}
			num, _ := strconv.ParseInt(S[p:p+i], 10, 32)
			if deep < 2 || int(num) == pre1+pre2 {
				//ans = append(ans, int(num))
				if dfs(p+i, pre2, int(num), deep+1) {
					ans = append(ans, int(num))
					return true
				}
			}
		}
		return false
	}

	if dfs(0, 0, 0, 0) {
		ansLen := len(ans)
		for i := 0; i < ansLen/2; i++ {
			ans[i], ans[ansLen-1-i] = ans[ansLen-1-i], ans[i]
		}
		return ans
	} else {
		return []int{}
	}
}
