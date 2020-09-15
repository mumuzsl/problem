package lc

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	states := []map[int8]int{
		{' ': 0, 's': 1, 'd': 2, '?': 3},
		{'d': 2, ' ': 3, 's': 3, '?': 3},
		{'d': 2, ' ': 3, 's': 3, '?': 3},
		{'d': 3, ' ': 3, 's': 3, '?': 3},
	}
	sign, ans := 0, 0
	p := 0
	var t int8
	var r int
	for _, c := range str {
		t = '?'
		if c >= '0' && c <= '9' {
			t = 'd'
		} else if c == '+' || c == '-' {
			t = 's'
		} else if c == ' ' {
			t = int8(c)
		}
		p = states[p][t]
		if p == 2 {
			if ans != 0 {
				if sign == 0 {
					r = math.MaxInt32 / ans
					fmt.Println(r)
					if r < 10 || math.MaxInt32-ans*10 < int(c-'0') {
						ans = math.MaxInt32
						return ans
					}
				} else {
					r = math.MinInt32 / (0 - ans)
					fmt.Println(r, math.MinInt32+ans*10 > 0-int(c-'0'))
					if r < 10 || math.MinInt32+ans*10 > 0-int(c-'0') {
						ans = math.MinInt32
						return ans
					}
				}
			}
			ans = ans*10 + int(c-'0')
			fmt.Println(ans, c-'0')
		} else if p == 1 {
			if c == '-' {
				sign = 1
			} else {
				sign = 0
			}
		}
	}

	if sign == 1 {
		return 0 - ans
	}
	return ans
}
