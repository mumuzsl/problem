package lco

// 都能想到的算法，会超时，淘汰
func myPow(x float64, n int) float64 {
	var r float64
	r = 1
	for i := 0; i < n; i++ {
		r *= x
	}
	for i := 0; i > n; i-- {
		r /= x
	}
	return r
}

func myPow2(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	var r float64
	r = 1
	if n < 0 {
		x, n = 1/x, -n
	}
	for n != 0 {
		if n&1 == 1 {
			r *= x
		}
		x *= x
		n >>= 1
	}
	return r
}
