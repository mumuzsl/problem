package lco

func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b, p, x, rem := n/3-1, n%3, 1000000007, 3, 1
	for a > 0 {
		if a%2 == 1 {
			rem = (rem * x) % p
		}
		x = x * x % p
		a /= 2
	}
	if b == 0 {
		return rem * 3 % p
	}
	if b == 1 {
		return rem * 4 % p
	}
	return rem * 6 % p
}
