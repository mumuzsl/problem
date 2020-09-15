package lco

func fib(n int) int {
	f := []int{0, 1}

	for i := 0; i < n; i++ {
		f[0], f[1] = f[1], (f[1] + f[0]) % (1e9 + 7)
	}

	return f[0]
}
