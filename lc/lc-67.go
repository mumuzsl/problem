package lc

import "strconv"

func addBinary(a string, b string) string {
	carry := 0
	ans := ""
	var sum int
	var c1, c2 int8

	aLen, bLen := len(a), len(b)
	i, j := aLen-1, bLen-1
	for {
		c1, c2 = '0', '0'
		if i >= 0 {
			c1 = int8(a[i])
			i--
		}
		if j >= 0 {
			c2 = int8(b[j])
			j--
		}
		sum, carry = int((c1-'0')+(c2-'0'))+carry, 0
		if sum >= 2 {
			carry = 1
		}
		ans = strconv.Itoa(sum%2) + ans
		if i < 0 && j < 0 && carry == 0 {
			break
		}
	}
	return ans
}
