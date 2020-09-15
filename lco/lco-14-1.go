package lco

import "math"

func cuttingRope(n int) int {
	table := make([]int, 5)
	table[0] = 1
	table[1] = 1
	table[2] = 1
	table[3] = 2
	table[4] = 4

	temp := 0
	k := 5
	for k <= n {
		table = append(table, 0)
		for i := 2; i < k; i++ {
			if k-i >= table[k-i] {
				temp = (i * (k - i))
			} else {
				temp = (i * table[k-i])
			}

			if temp > table[k] {
				table[k] = temp
			}
		}
		//fmt.Println(k, table[k])
		k++
	}

	return table[n]
}

func cuttingRope1_3(n int) int {
	table := make([]int, 3)
	table[0] = 1
	table[1] = 2
	table[2] = 4

	temp := 0
	k := 5
	i := 0
	for k <= n {
		i = (k - 3 - 2) % 3
		if k-3 >= table[i] {
			temp = (3 * (k - 3))
		} else {
			temp = (3 * table[i])
		}

		table[i] = temp % (1e9 + 7)
		k++
	}

	return table[(n-2)%3]
}

func cuttingRope1_2(n int) int {
	if n <= 3 {
		return n - 1
	}
	a, b := n/3, n%3
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	}
	if b == 1 {
		return int(math.Pow(3, float64(a)-1) * 4)
	}
	return int(math.Pow(3, float64(a)) * 2)
}
