package lc

import (
	"strings"
)

func permutation(s string) []string {
	sLen := len(s)
	ans := make(map[string]int, 0)
	temp := strings.Split(s, "")

	var dfs func(x int)
	dfs = func(x int) {
		if x == sLen {
			ans[strings.Join(temp, "")] = 1
			return
		}
		for i := x; i < sLen; i++ {
			temp[i], temp[x] = temp[x], temp[i]
			dfs(x + 1)
			temp[i], temp[x] = temp[x], temp[i]
		}
	}

	dfs(0)
	r := make([]string, len(ans))
	k := 0
	for key := range ans {
		r[k] = key
		k++
	}
	return r
}
