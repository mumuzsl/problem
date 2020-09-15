package lc

import "fmt"

func countEval(s string, result int) int {
	slen := len(s)
	dp := make([]int, slen*slen*2)
	var rec func(start, end, resutl int) int
	var ans func(v1, v2 int, op uint8) int
	ans = func(v1, v2 int, op uint8) int {
		switch op {
		case '|':
			return v1 | v2
		case '^':
			return v1 ^ v2
		default:
			return v1 & v2
		}
	}
	rec = func(start, end, re int) int {
		if start == end {
			if int(s[start]-'0') == re {
				return 1
			} else {
				return 0
			}
		}
		index := start*slen*2 + end*2 + re
		fmt.Println(start, end, re, slen, index)
		if dp[index] != -1 {
			fmt.Println("afdasf")
			return dp[index]
		}
		count := 0
		var op uint8
		for k := start; k < end; k += 2 {
			op = s[k+1]
			for i := 0; i < 2; i++ {
				if ans(i, 0, op) == re {
					count += rec(start, k, i) * rec(k+2, end, 0)
				}
				if ans(i, 1, op) == re {
					count += rec(start, k, i) * rec(k+2, end, 1)
				}
			}
		}
		dp[index] = count
		return count
	}
	for i := 0; i < slen*slen*2; i++ {
		dp[i] = -1
	}
	return rec(0, slen-1, result)
}
