package lc

func dieSimulator(n int, rollMax []int) int {
	const mod = 1e9 + 7
	dp := make([]int, 15*6*n)
	conv := func(i, j, k int) int {
		return i*15*6 + j*15 + k
	}
	var s int
	for i := 0; i < n; i++ {
		for j := 0; j < 6; j++ {
			if i == 0 {
				dp[conv(i, j, 0)] = 1
				continue
			}
			for k := 1; k < rollMax[j]; k++ {
				dp[conv(i, j, k)] = dp[conv(i-1, j, k-1)]
			}
			s = 0
			for z := 0; z < 6; z++ {
				if z == j {
					continue
				}
				for k := 0; k < 15; k++ {
					// fmt.Println(i,z,k)
					s += dp[conv(i-1, z, k)]
					s %= mod
				}
			}
			dp[conv(i, j, 0)] = s
		}
	}
	ans := 0
	for j := 0; j < 6; j++ {
		for k := 0; k < 15; k++ {
			ans += dp[conv(n-1, j, k)]
			ans %= mod
		}
	}
	return ans
}
