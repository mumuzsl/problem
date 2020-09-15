package lco

func movingCount(m int, n int, k int) int {
	// 为0表示没走过为, 1表示走过, 2表示不能走
	// 用一维数组代替二维数组来进行标记，数组默认值为0
	board := make([]int, m*n)

	//求一个数的数位和
	var sum func(k int) int
	sum = func(k int) int {
		return k/10%10 + k%10
	}

	// 递归搜索一个格子，返回包括这个格子在内可以到达的格子数
	var search func(r, c int) int
	search = func(r, c int) int {
		// 下标越界或标记不为0直接返回0 (标记不为0表示走过或不可以走)
		if (r >= m || r < 0) || (c >= n || c < 0) || board[r*n+c] != 0 {
			return 0
		}
		// 下标数位和大于k,将对应的格子标记为2并返回
		if (sum(r) + sum(c)) > k {
			board[r*n+c] = 2
			return 0
		}
		count := 1
		board[r*n+c] = 1
		count += search(r+1, c)
		//count += search(r-1, c)
		count += search(r, c+1)
		//count += search(r, c-1)
		return count
	}

	return search(0, 0)
}

func movingCount2(m int, n int, k int) int {
	// 为0表示没走过为, 1表示走过, 2表示不能走
	// 用一维数组代替二维数组来进行标记，数组默认值为0
	board := make([]int, m*n)

	//求一个数的数位和
	var sum func(k int) int
	sum = func(k int) int {
		s := 0
		for k != 0 {
			s += k % 10
			k = k / 10
		}
		return s
	}

	// 递归搜索一个格子，返回包括这个格子在内可以到达的格子数
	var search func(r, c int) int
	search = func(r, c int) int {
		// 下标越界或标记不为0直接返回0 (标记不为0表示走过或不可以走)
		if (r >= m || r < 0) || (c >= n || c < 0) || board[r*n+c] != 0 {
			return 0
		}
		count := 1
		board[r*n+c] = 1
		if r+1 < m && (sum(r+1)+sum(c)) > k {
			board[(r+1)*n+c] = 2
		} else {
			count += search(r+1, c)
		}
		if c+1 < n && (sum(r)+sum(c+1)) > k {
			board[r*n+c+1] = 2
		} else {
			count += search(r, c+1)
		}
		return count
	}

	return search(0, 0)
}
